package provider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxedge"

	"github.com/tmunzer/mistapi-go/mistapi"
	sdkerrors "github.com/tmunzer/mistapi-go/mistapi/errors"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgMxedgeResource{}
	_ resource.ResourceWithConfigure   = &orgMxedgeResource{}
	_ resource.ResourceWithImportState = &orgMxedgeResource{}
)

func NewOrgMxedgeResource() resource.Resource {
	return &orgMxedgeResource{}
}

type orgMxedgeResource struct {
	client mistapi.ClientInterface
}

func (r *orgMxedgeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgMxedge Resource client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *orgMxedgeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_mxedge"
}

func (r *orgMxedgeResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages MxEdge devices in the Mist Organization.\n\n" +
			"MxEdge is a multi-service edge platform that provides tunneling, switching, and other network services.",
		Attributes: resource_org_mxedge.OrgMxedgeResourceSchema(ctx).Attributes,
	}
}

// getMxEdge retrieves an MxEdge by ID. GetOrgMxEdge returns 404 for site-assigned
// devices (the org-level endpoint only sees unassigned devices), so on 404 it falls
// back to ListOrgMxEdges(for_site=any) and searches by ID.
// Returns (device, true) if found, (nil, true) if definitively not found (caller
// should remove from state), or (nil, false) on a real API error.
func (r *orgMxedgeResource) getMxEdge(ctx context.Context, orgId, mxedgeId uuid.UUID) (*models.Mxedge, bool, error) {
	getResp, err := r.client.OrgsMxEdges().GetOrgMxEdge(ctx, orgId, mxedgeId)
	if err == nil {
		return &getResp.Data, true, nil
	}
	// Only fall back on 404; surface all other errors.
	var http404 *sdkerrors.ResponseHttp404
	if !errors.As(err, &http404) &&
		!(getResp.Response != nil && getResp.Response.StatusCode == 404) {
		return nil, false, err
	}
	tflog.Warn(ctx, fmt.Sprintf("GetOrgMxEdge 404 for %s, falling back to ListOrgMxEdges(for_site=any)", mxedgeId))
	forSiteAny := models.MxedgeForSiteEnum_ANY
	listResp, listErr := r.client.OrgsMxEdges().ListOrgMxEdges(ctx, orgId, &forSiteAny, nil, nil)
	if listErr != nil {
		return nil, false, fmt.Errorf("GetOrgMxEdge returned 404 and list fallback failed: %w", listErr)
	}
	for i := range listResp.Data {
		if listResp.Data[i].Id != nil && *listResp.Data[i].Id == mxedgeId {
			return &listResp.Data[i], true, nil
		}
	}
	// Not found in org at all.
	return nil, true, nil
}

func (r *orgMxedgeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgMxedge Create")
	var plan, state resource_org_mxedge.OrgMxedgeModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	var mistMxedge models.Mxedge
	var mxedgeId uuid.UUID

	// Check if claim_code is provided - if so, claim the device
	if !plan.Magic.IsNull() && plan.Magic.ValueString() != "" {
		tflog.Info(ctx, "Claiming OrgMxedge with claim_code in Org "+plan.OrgId.ValueString())
		claimCodes := []string{plan.Magic.ValueString()}

		claimResponse, err := r.client.OrgsInventory().AddOrgInventory(ctx, orgId, claimCodes)

		// The SDK returns a ResponseInventoryError (HTTP 400) when any claim code
		// fails. Extract the per-code reasons for a useful error message, but also
		// allow through cases where the device was already claimed (duplicated).
		var inventoryAdded []models.ResponseInventoryInventoryAddedItems
		if err != nil {
			var invErr *sdkerrors.ResponseInventoryError
			if errors.As(err, &invErr) {
				if len(invErr.MError) > 0 {
					for i, code := range invErr.MError {
						reason := ""
						if i < len(invErr.Reason) {
							reason = invErr.Reason[i]
						}
						resp.Diagnostics.AddError(
							"Error claiming \"mist_org_mxedge\" resource",
							fmt.Sprintf("Unable to claim the MxEdge with claim code %q: %s", code, reason),
						)
					}
					return
				}
				// No per-code errors — device may be duplicated (already claimed)
				for _, dup := range invErr.InventoryDuplicated {
					resp.Diagnostics.AddWarning(
						"MxEdge already claimed",
						fmt.Sprintf("Device with claim code %q (MAC: %s) was already in the org inventory and has been imported.", dup.Magic, dup.Mac),
					)
					var tmp models.ResponseInventoryInventoryAddedItems
					tmp.Mac = dup.Mac
					tmp.Magic = dup.Magic
					tmp.Model = dup.Model
					tmp.Serial = dup.Serial
					tmp.Type = dup.Type
					inventoryAdded = append(inventoryAdded, tmp)
				}
				inventoryAdded = append(inventoryAdded, invErr.InventoryAdded...)
				if len(inventoryAdded) == 0 {
					resp.Diagnostics.AddError(
						"Error claiming \"mist_org_mxedge\" resource",
						"Claim returned no added or duplicated devices: "+err.Error(),
					)
					return
				}
			} else {
				resp.Diagnostics.AddError(
					"Error claiming \"mist_org_mxedge\" resource",
					"Unable to claim the MxEdge: "+err.Error(),
				)
				return
			}
		} else {
			if claimResponse.Response == nil {
				resp.Diagnostics.AddError(
					"Error claiming \"mist_org_mxedge\" resource",
					"API response is nil",
				)
				return
			}
			if claimResponse.Response.StatusCode != 200 {
				apiErr := mistapierror.ProcessApiError(claimResponse.Response.StatusCode, claimResponse.Response.Body, err)
				if apiErr != "" {
					resp.Diagnostics.AddError(
						"Error claiming \"mist_org_mxedge\" resource",
						fmt.Sprintf("Unable to claim the MxEdge. %s", apiErr),
					)
					return
				}
			}
			// Handle duplicated devices (already in org) the same as errors.As path above.
			for _, dup := range claimResponse.Data.InventoryDuplicated {
				resp.Diagnostics.AddWarning(
					"MxEdge already claimed",
					fmt.Sprintf("Device with claim code %q (MAC: %s) was already in the org inventory and has been imported.", dup.Magic, dup.Mac),
				)
				var tmp models.ResponseInventoryInventoryAddedItems
				tmp.Mac = dup.Mac
				tmp.Magic = dup.Magic
				tmp.Model = dup.Model
				tmp.Serial = dup.Serial
				tmp.Type = dup.Type
				inventoryAdded = append(inventoryAdded, tmp)
			}
			inventoryAdded = append(inventoryAdded, claimResponse.Data.InventoryAdded...)
		}

		// Extract the device ID from the claim response
		if len(inventoryAdded) == 0 {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				"No devices were added in the claim response",
			)
			return
		}

		addedDevice := inventoryAdded[0]
		if addedDevice.Mac == "" {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				"Claimed device has no MAC address",
			)
			return
		}

		// Construct device ID from MAC address: 00000000-0000-0000-1000-{MAC}
		cleanMac := strings.ReplaceAll(addedDevice.Mac, ":", "")
		cleanMac = strings.ReplaceAll(cleanMac, "-", "")
		cleanMac = strings.ToLower(cleanMac)

		if len(cleanMac) != 12 {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				fmt.Sprintf("Invalid MAC address length: expected 12, got %d", len(cleanMac)),
			)
			return
		}

		uuidStr := fmt.Sprintf("00000000-0000-0000-1000-%s", cleanMac)
		mxedgeId, err = uuid.Parse(uuidStr)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				"Failed to construct device UUID from MAC: "+err.Error(),
			)
			return
		}

		tflog.Info(ctx, "Retrieving claimed MxEdge details: "+mxedgeId.String())
		claimedDevice, found, getMxEdgeErr := r.getMxEdge(ctx, orgId, mxedgeId)
		if getMxEdgeErr != nil {
			resp.Diagnostics.AddError(
				"Error retrieving claimed \"mist_org_mxedge\" resource",
				getMxEdgeErr.Error(),
			)
			return
		}
		if !found || claimedDevice == nil {
			resp.Diagnostics.AddError(
				"Error retrieving claimed \"mist_org_mxedge\" resource",
				fmt.Sprintf("MxEdge %s not found in org after claiming", mxedgeId),
			)
			return
		}
		mistMxedge = *claimedDevice

		// If additional fields from plan need to be updated (name, mxcluster_id, etc.), do it now.
		// Site assignment is intentionally excluded here and handled by AssignOrgMxEdgeToSite below.
		if (!plan.Name.IsNull() && !plan.Name.IsUnknown()) || (!plan.MxclusterId.IsNull() && !plan.MxclusterId.IsUnknown()) {
			updateMxedge, diags := resource_org_mxedge.TerraformToSdk(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			// Strip site fields — sending site_id in UpdateOrgMxEdge causes the Mist API to
			// perform an inline assignment, which would conflict with the dedicated assign call below.
			updateMxedge.SiteId = nil
			updateMxedge.ForSite = nil

			tflog.Info(ctx, "Updating claimed MxEdge with additional fields")
			data, err := r.client.OrgsMxEdges().UpdateOrgMxEdge(ctx, orgId, mxedgeId, updateMxedge)
			if data.Response == nil {
				resp.Diagnostics.AddError(
					"Error updating claimed \"mist_org_mxedge\" resource",
					"API response is nil",
				)
				return
			}
			if data.Response.StatusCode != 200 {
				apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
				if apiErr != "" {
					resp.Diagnostics.AddError(
						"Error updating claimed \"mist_org_mxedge\" resource",
						fmt.Sprintf("Unable to update the claimed MxEdge. %s", apiErr),
					)
					return
				}
			}

			body, err := io.ReadAll(data.Response.Body)
			if err != nil {
				resp.Diagnostics.AddError("Unable to read API response body", err.Error())
				return
			}
			if err = json.Unmarshal(body, &mistMxedge); err != nil {
				resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
				return
			}
		}

	} else {
		// Create MxEdge using name and model
		tflog.Info(ctx, "Creating OrgMxedge with name/model in Org "+plan.OrgId.ValueString())

		mxedge, diags := resource_org_mxedge.TerraformToSdk(ctx, &plan)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

		data, err := r.client.OrgsMxEdges().CreateOrgMxEdge(ctx, orgId, mxedge)
		if data.Response == nil {
			resp.Diagnostics.AddError(
				"Error creating \"mist_org_mxedge\" resource",
				"API response is nil",
			)
			return
		}
		if data.Response.StatusCode != 200 {
			apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
			if apiErr != "" {
				resp.Diagnostics.AddError(
					"Error creating \"mist_org_mxedge\" resource",
					fmt.Sprintf("Unable to create the MxEdge. %s", apiErr),
				)
				return
			}
		}

		body, err := io.ReadAll(data.Response.Body)
		if err != nil {
			resp.Diagnostics.AddError("Unable to read API response body", err.Error())
			return
		}
		if err = json.Unmarshal(body, &mistMxedge); err != nil {
			resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
			return
		}
	}

	// The Mist API ignores for_site in the Create/Update body; it is only set
	// via AssignOrgMxEdgeToSite. If site_id is specified, assign now so that
	// for_site is correctly reflected in the state after creation.
	if !plan.SiteId.IsNull() && !plan.SiteId.IsUnknown() && plan.SiteId.ValueString() != "" {
		if mistMxedge.Id == nil {
			resp.Diagnostics.AddError(
				"Error assigning \"mist_org_mxedge\" to site after create",
				"Created MxEdge has no ID",
			)
			return
		}
		createdId := *mistMxedge.Id
		siteId, err := uuid.Parse(plan.SiteId.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid \"site_id\" value for \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to parse the UUID %q: %s", plan.SiteId.ValueString(), err.Error()),
			)
			return
		}
		tflog.Info(ctx, fmt.Sprintf("Assigning newly created MxEdge %s to site %s", createdId, siteId))
		assignResp, err := r.client.OrgsMxEdges().AssignOrgMxEdgeToSite(ctx, orgId, &models.MxedgesAssign{
			MxedgeIds: []uuid.UUID{createdId},
			SiteId:    siteId,
		})
		if err != nil || assignResp.Response == nil || assignResp.Response.StatusCode != 200 {
			apiErr := ""
			if assignResp.Response != nil {
				apiErr = mistapierror.ProcessApiError(assignResp.Response.StatusCode, assignResp.Response.Body, err)
			}
			if apiErr == "" && err != nil {
				apiErr = err.Error()
			}
			resp.Diagnostics.AddError(
				"Error assigning \"mist_org_mxedge\" to site after create",
				fmt.Sprintf("Unable to assign the MxEdge to site: %s", apiErr),
			)
			return
		}
		// Assign succeeded — set for_site = true and site_id directly on the local struct.
		// GetOrgMxEdge returns 404 for VM MxEdges so we avoid that call here.
		mistMxedge.ForSite = models.ToPointer(true)
		mistMxedge.SiteId = &siteId
	}

	state, diags = resource_org_mxedge.SdkToTerraform(ctx, &mistMxedge)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve claim_code from plan if it was used for claiming
	// The API doesn't return this value, but we need to keep it in state for consistency
	if !plan.Magic.IsNull() && plan.Magic.ValueString() != "" {
		state.Magic = plan.Magic
	}

	// Preserve mxcluster_id when the plan explicitly set it to "" and the API returns nil.
	// API returns nil when no cluster is assigned; "" and nil are equivalent on the API side.
	if state.MxclusterId.IsNull() && !plan.MxclusterId.IsNull() && !plan.MxclusterId.IsUnknown() {
		state.MxclusterId = plan.MxclusterId
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_mxedge.OrgMxedgeModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxedge Read: mxedge_id "+state.Id.ValueString())
	tflog.Debug(ctx, fmt.Sprintf("Reading MxEdge with org_id=%s, mxedge_id=%s", state.OrgId.ValueString(), state.Id.ValueString()))

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxedgeId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	foundDevice, found, getMxEdgeErr := r.getMxEdge(ctx, orgId, mxedgeId)
	if getMxEdgeErr != nil {
		resp.Diagnostics.AddError(
			"Error reading \"mist_org_mxedge\" resource",
			getMxEdgeErr.Error(),
		)
		return
	}
	if !found || foundDevice == nil {
		tflog.Warn(ctx, fmt.Sprintf("MxEdge %s not found in org, removing from state", mxedgeId))
		resp.State.RemoveResource(ctx)
		return
	}

	// Preserve org_id, claim_code and mxcluster_id from existing state before overwriting
	existingOrgId := state.OrgId
	existingClaimCode := state.Magic
	existingMxclusterId := state.MxclusterId

	state, diags = resource_org_mxedge.SdkToTerraform(ctx, foundDevice)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Restore org_id - API doesn't return this value
	state.OrgId = existingOrgId

	// Restore claim_code if it was previously set
	// The API doesn't return this value, but we need to keep it in state for consistency
	if !existingClaimCode.IsNull() && existingClaimCode.ValueString() != "" {
		state.Magic = existingClaimCode
	}

	// Restore mxcluster_id when API returns nil (no cluster) but state had an explicit empty string.
	// Both nil and zero-UUID mean "no cluster" on the API side; preserve the state value to avoid drift.
	if state.MxclusterId.IsNull() && !existingMxclusterId.IsNull() && !existingMxclusterId.IsUnknown() {
		state.MxclusterId = existingMxclusterId
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_mxedge.OrgMxedgeModel
	tflog.Info(ctx, "Starting OrgMxedge Update")

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxedge Update for MxEdge "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxedgeId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	// Determine current and planned site assignments
	currentSiteId := ""
	planSiteId := ""
	siteIdIsUnknown := plan.SiteId.IsUnknown()

	if !state.SiteId.IsNull() {
		currentSiteId = state.SiteId.ValueString()
	}
	if !plan.SiteId.IsNull() && !plan.SiteId.IsUnknown() {
		planSiteId = plan.SiteId.ValueString()
	}

	// Step 1: If device is currently site-assigned AND we know the final site_id state,
	// temporarily unassign it BEFORE updating (UpdateOrgMxEdge does not work on site-assigned devices)
	// Skip unassign if site_id is Unknown to preserve current assignment
	if currentSiteId != "" && !siteIdIsUnknown {
		tflog.Info(ctx, fmt.Sprintf("Temporarily unassigning MxEdge from site %s to allow attribute updates", currentSiteId))

		unassignBody := models.MxedgesUnassign{
			MxedgeIds: []uuid.UUID{mxedgeId},
		}

		unassignResponse, err := r.client.OrgsMxEdges().UnassignOrgMxEdgeFromSite(ctx, orgId, &unassignBody)
		if err != nil || unassignResponse.Response == nil || unassignResponse.Response.StatusCode != 200 {
			if unassignResponse.Response != nil {
				apiErr := mistapierror.ProcessApiError(unassignResponse.Response.StatusCode, unassignResponse.Response.Body, err)
				if apiErr != "" {
					resp.Diagnostics.AddError(
						"Error unassigning \"mist_org_mxedge\" from site",
						fmt.Sprintf("Unable to unassign the MxEdge from site. %s", apiErr),
					)
				} else {
					// Fallback when ProcessApiError returns empty string
					body, readErr := io.ReadAll(unassignResponse.Response.Body)
					if readErr != nil {
						body = []byte(fmt.Sprintf("(unable to read response body: %s)", readErr.Error()))
					}
					resp.Diagnostics.AddError(
						"Error unassigning \"mist_org_mxedge\" from site",
						fmt.Sprintf("Unable to unassign the MxEdge from site. HTTP %d: %s", unassignResponse.Response.StatusCode, string(body)),
					)
				}
			} else {
				resp.Diagnostics.AddError(
					"Error unassigning \"mist_org_mxedge\" from site",
					"API response is nil",
				)
			}
			return
		}
	}

	// Step 2: Update the MxEdge attributes (name, model, etc.)
	mxedge, diags := resource_org_mxedge.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Strip site fields — sending site_id in UpdateOrgMxEdge causes the Mist API to
	// perform an inline assignment, which would conflict with the dedicated
	// AssignOrgMxEdgeToSite call in Step 3.
	mxedge.SiteId = nil
	mxedge.ForSite = nil

	data, err := r.client.OrgsMxEdges().UpdateOrgMxEdge(ctx, orgId, mxedgeId, mxedge)
	if data.Response == nil {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_mxedge\" resource",
			"API response is nil",
		)
		return
	}
	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to update the MxEdge. %s", apiErr),
			)
			return
		}
	}

	body, err := io.ReadAll(data.Response.Body)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read API response body", err.Error())
		return
	}
	mistMxedge := models.Mxedge{}
	if err = json.Unmarshal(body, &mistMxedge); err != nil {
		resp.Diagnostics.AddError("Unable to unmarshal API response", err.Error())
		return
	}

	// Step 3: If plan includes site assignment, assign/re-assign device to site AFTER updating
	// Skip reassignment if site_id is Unknown (preserves current state)
	if planSiteId != "" && !siteIdIsUnknown {
		tflog.Info(ctx, fmt.Sprintf("Assigning MxEdge to site: %s", planSiteId))
		siteId, err := uuid.Parse(planSiteId)
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid \"site_id\" value for \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to parse the UUID \"%s\": %s", planSiteId, err.Error()),
			)
			return
		}

		assignBody := models.MxedgesAssign{
			MxedgeIds: []uuid.UUID{mxedgeId},
			SiteId:    siteId,
		}

		assignResponse, err := r.client.OrgsMxEdges().AssignOrgMxEdgeToSite(ctx, orgId, &assignBody)
		if err != nil || assignResponse.Response == nil || assignResponse.Response.StatusCode != 200 {
			if assignResponse.Response != nil {
				apiErr := mistapierror.ProcessApiError(assignResponse.Response.StatusCode, assignResponse.Response.Body, err)
				if apiErr != "" {
					resp.Diagnostics.AddError(
						"Error assigning \"mist_org_mxedge\" to site",
						fmt.Sprintf("Unable to assign the MxEdge to site. %s", apiErr),
					)
				} else {
					// Fallback when ProcessApiError returns empty string
					body, readErr := io.ReadAll(assignResponse.Response.Body)
					if readErr != nil {
						body = []byte(fmt.Sprintf("(unable to read response body: %s)", readErr.Error()))
					}
					resp.Diagnostics.AddError(
						"Error assigning \"mist_org_mxedge\" to site",
						fmt.Sprintf("Unable to assign the MxEdge to site. HTTP %d: %s", assignResponse.Response.StatusCode, string(body)),
					)
				}
			} else {
				resp.Diagnostics.AddError(
					"Error assigning \"mist_org_mxedge\" to site",
					"API response is nil",
				)
			}
			return
		}

		// Refresh state after assignment using getMxEdge to handle site-assigned devices.
		postAssignDevice, found, getMxEdgeErr := r.getMxEdge(ctx, orgId, mxedgeId)
		if getMxEdgeErr != nil {
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" after site assignment",
				getMxEdgeErr.Error(),
			)
			return
		}
		if !found || postAssignDevice == nil {
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" after site assignment",
				fmt.Sprintf("MxEdge %s not found after site assignment", mxedgeId),
			)
			return
		}
		mistMxedge = *postAssignDevice
	}

	// Convert API response to Terraform state
	state, diags = resource_org_mxedge.SdkToTerraform(ctx, &mistMxedge)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Restore claim_code if it was previously set
	// The API doesn't return this value, but we need to keep it in state for consistency
	if !plan.Magic.IsNull() && plan.Magic.ValueString() != "" {
		state.Magic = plan.Magic
	}

	// Preserve mxcluster_id when the plan explicitly set it to "" and the API returns nil.
	// API returns nil when no cluster is assigned; "" and nil are equivalent on the API side.
	if state.MxclusterId.IsNull() && !plan.MxclusterId.IsNull() && !plan.MxclusterId.IsUnknown() {
		state.MxclusterId = plan.MxclusterId
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_mxedge.OrgMxedgeModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgMxedge Delete: mxedge_id "+state.Id.ValueString())

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	mxedgeId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	// If device is currently site-assigned, unassign it before deletion
	if !state.SiteId.IsNull() && state.SiteId.ValueString() != "" {
		tflog.Info(ctx, fmt.Sprintf("Unassigning MxEdge from site %s before deletion", state.SiteId.ValueString()))

		unassignBody := models.MxedgesUnassign{
			MxedgeIds: []uuid.UUID{mxedgeId},
		}

		unassignResponse, err := r.client.OrgsMxEdges().UnassignOrgMxEdgeFromSite(ctx, orgId, &unassignBody)
		if err != nil || unassignResponse.Response == nil || unassignResponse.Response.StatusCode != 200 {
			if unassignResponse.Response != nil {
				apiErr := mistapierror.ProcessApiError(unassignResponse.Response.StatusCode, unassignResponse.Response.Body, err)
				if apiErr != "" {
					resp.Diagnostics.AddError(
						"Error unassigning \"mist_org_mxedge\" from site before deletion",
						fmt.Sprintf("Unable to unassign the MxEdge from site. %s", apiErr),
					)
				} else {
					// Fallback when ProcessApiError returns empty string
					body, readErr := io.ReadAll(unassignResponse.Response.Body)
					if readErr != nil {
						body = []byte(fmt.Sprintf("(unable to read response body: %s)", readErr.Error()))
					}
					resp.Diagnostics.AddError(
						"Error unassigning \"mist_org_mxedge\" from site before deletion",
						fmt.Sprintf("Unable to unassign the MxEdge from site. HTTP %d: %s", unassignResponse.Response.StatusCode, string(body)),
					)
				}
			} else {
				resp.Diagnostics.AddError(
					"Error unassigning \"mist_org_mxedge\" from site before deletion",
					"API response is nil",
				)
			}
			return
		}
	}

	httpr, err := r.client.OrgsMxEdges().DeleteOrgMxEdge(ctx, orgId, mxedgeId)
	if httpr.StatusCode != 200 && httpr.StatusCode != 404 {
		apiErr := mistapierror.ProcessApiError(httpr.StatusCode, httpr.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to delete the MxEdge. %s", apiErr),
			)
			return
		}
	}
}

func (r *orgMxedgeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Import \"id\" format must be \"{org_id}.{mxedge_id}\", got %s", req.ID),
		)
		return
	}

	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxedge_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{mxedge_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
