package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxedge"

	"github.com/tmunzer/mistapi-go/mistapi"

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
		if err != nil {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				"Unable to claim the MxEdge: "+err.Error(),
			)
			return
		}

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

		// Extract the device ID from the claim response
		if len(claimResponse.Data.InventoryAdded) == 0 {
			resp.Diagnostics.AddError(
				"Error claiming \"mist_org_mxedge\" resource",
				"No devices were added in the claim response",
			)
			return
		}

		addedDevice := claimResponse.Data.InventoryAdded[0]
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

		// Now retrieve the claimed device details using ListOrgMxEdges
		// Paginate through all results to find the device
		tflog.Info(ctx, "Retrieving claimed MxEdge details: "+mxedgeId.String())
		var limit = 1000
		var page = 1
		var total = 0
		var found bool
		forSite := models.MxedgeForSiteEnum_ANY

		for limit*page <= total+limit && !found {
			listResp, listErr := r.client.OrgsMxEdges().ListOrgMxEdges(ctx, orgId, &forSite, &limit, &page)
			if listErr != nil {
				resp.Diagnostics.AddError(
					"Error retrieving claimed \"mist_org_mxedge\" resource",
					fmt.Sprintf("Unable to list MxEdges: %s", listErr.Error()),
				)
				return
			}
			if listResp.Response == nil {
				resp.Diagnostics.AddError(
					"Error retrieving claimed \"mist_org_mxedge\" resource",
					"API response is nil",
				)
				return
			}
			if listResp.Response.StatusCode != 200 {
				apiErr := mistapierror.ProcessApiError(listResp.Response.StatusCode, listResp.Response.Body, listErr)
				resp.Diagnostics.AddError(
					"Error retrieving claimed \"mist_org_mxedge\" resource",
					fmt.Sprintf("Unable to list MxEdges: %s", apiErr),
				)
				return
			}

			// Extract pagination info from headers on first page
			if page == 1 {
				limitHeader := listResp.Response.Header.Get("X-Page-Limit")
				totalHeader := listResp.Response.Header.Get("X-Page-Total")
				if limitHeader != "" {
					if parsedLimit, err := strconv.Atoi(limitHeader); err == nil {
						limit = parsedLimit
					}
				}
				if totalHeader != "" {
					if parsedTotal, err := strconv.Atoi(totalHeader); err == nil {
						total = parsedTotal
					}
				}
			}

			// Search for the device in this page
			for _, device := range listResp.Data {
				if device.Id != nil && *device.Id == mxedgeId {
					mistMxedge = device
					found = true
					break
				}
			}

			page++
		}

		if !found {
			resp.Diagnostics.AddError(
				"Error retrieving claimed \"mist_org_mxedge\" resource",
				fmt.Sprintf("Device %s not found in organization after claiming (searched %d devices)", mxedgeId.String(), total),
			)
			return
		}

		// If additional fields from plan need to be updated (name, site_id, etc.), do it now
		if !plan.Name.IsNull() || !plan.SiteId.IsNull() || !plan.MxclusterId.IsNull() {
			updateMxedge, diags := resource_org_mxedge.TerraformToSdk(ctx, &plan)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}

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

			body, _ := io.ReadAll(data.Response.Body)
			json.Unmarshal(body, &mistMxedge)
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

		body, _ := io.ReadAll(data.Response.Body)
		json.Unmarshal(body, &mistMxedge)
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

	// Use ListOrgMxEdges with forSite=ANY to get both org-level and site-assigned MxEdges
	// Implement pagination to handle orgs with more than 1000 devices
	var limit = 1000
	var page = 0
	var total = 9999
	var found bool
	var mistMxedge models.Mxedge
	forSite := models.MxedgeForSiteEnum_ANY

	for limit*page < total && !found {
		page += 1
		tflog.Debug(ctx, fmt.Sprintf("Searching for MxEdge in list page %d (limit=%d, total=%d)", page, limit, total))

		listResp, listErr := r.client.OrgsMxEdges().ListOrgMxEdges(ctx, orgId, &forSite, &limit, &page)
		if listErr != nil {
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to list MxEdges: %s", listErr.Error()),
			)
			return
		}
		if listResp.Response == nil {
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" resource",
				"List API response is nil",
			)
			return
		}
		if listResp.Response.StatusCode != 200 {
			apiErr := mistapierror.ProcessApiError(listResp.Response.StatusCode, listResp.Response.Body, listErr)
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" resource",
				fmt.Sprintf("Unable to list MxEdges: %s", apiErr),
			)
			return
		}

		// Extract pagination info from headers
		limitString := listResp.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limitString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Limit value into int, unexpected error: "+err.Error(),
			)
			return
		}

		totalString := listResp.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(totalString); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Unable to convert the X-Page-Total value into int, unexpected error: "+err.Error(),
			)
			return
		}

		// Search for the device in this page
		for _, device := range listResp.Data {
			if device.Id != nil && *device.Id == mxedgeId {
				mistMxedge = device
				found = true
				tflog.Debug(ctx, fmt.Sprintf("Found MxEdge in list on page %d", page))
				break
			}
		}
	}

	if !found {
		tflog.Warn(ctx, fmt.Sprintf("MxEdge not found in list after checking %d pages, removing from state", page))
		resp.State.RemoveResource(ctx)
		return
	}

	// Preserve org_id and claim_code from existing state before overwriting
	existingOrgId := state.OrgId
	existingClaimCode := state.Magic

	state, diags = resource_org_mxedge.SdkToTerraform(ctx, &mistMxedge)
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
					body, _ := io.ReadAll(unassignResponse.Response.Body)
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

	body, _ := io.ReadAll(data.Response.Body)
	mistMxedge := models.Mxedge{}
	json.Unmarshal(body, &mistMxedge)

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
					body, _ := io.ReadAll(assignResponse.Response.Body)
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

		// Refresh state after assignment
		// Use ListOrgMxEdges since GetOrgMxEdge doesn't work for site-assigned devices
		// Paginate through all results to find the device
		var limit = 1000
		var page = 1
		var total = 0
		var found bool
		forSite := models.MxedgeForSiteEnum_ANY

		for limit*page <= total+limit && !found {
			listResp, listErr := r.client.OrgsMxEdges().ListOrgMxEdges(ctx, orgId, &forSite, &limit, &page)
			if listErr != nil {
				resp.Diagnostics.AddError(
					"Error retrieving \"mist_org_mxedge\" after site assignment",
					fmt.Sprintf("Unable to list MxEdges: %s", listErr.Error()),
				)
				return
			}
			if listResp.Response == nil {
				resp.Diagnostics.AddError(
					"Error retrieving \"mist_org_mxedge\" after site assignment",
					"List API response is nil",
				)
				return
			}
			if listResp.Response.StatusCode != 200 {
				apiErr := mistapierror.ProcessApiError(listResp.Response.StatusCode, listResp.Response.Body, listErr)
				resp.Diagnostics.AddError(
					"Error retrieving \"mist_org_mxedge\" after site assignment",
					fmt.Sprintf("Unable to list MxEdges: %s", apiErr),
				)
				return
			}

			// Extract pagination info from headers on first page
			if page == 1 {
				limitHeader := listResp.Response.Header.Get("X-Page-Limit")
				if parsedLimit, err := strconv.Atoi(limitHeader); err != nil {
					resp.Diagnostics.AddError(
						"Error extracting HTTP Response Headers",
						"Unable to convert the X-Page-Limit value into int, unexpected error: "+err.Error(),
					)
					return
				} else {
					limit = parsedLimit
				}

				totalHeader := listResp.Response.Header.Get("X-Page-Total")
				if parsedTotal, err := strconv.Atoi(totalHeader); err != nil {
					resp.Diagnostics.AddError(
						"Error extracting HTTP Response Headers",
						"Unable to convert the X-Page-Total value into int, unexpected error: "+err.Error(),
					)
					return
				} else {
					total = parsedTotal
				}
			}

			// Search for the device in this page
			for _, device := range listResp.Data {
				if device.Id != nil && *device.Id == mxedgeId {
					mistMxedge = device
					found = true
					break
				}
			}

			page++
		}

		if !found {
			resp.Diagnostics.AddError(
				"Error retrieving \"mist_org_mxedge\" after site assignment",
				fmt.Sprintf("Device %s not found in organization after assignment (searched %d devices)", mxedgeId.String(), total),
			)
			return
		}
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
