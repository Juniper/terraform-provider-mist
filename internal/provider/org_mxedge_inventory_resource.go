package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_mxedge_inventory"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgMxedgeInventoryResource{}
	_ resource.ResourceWithConfigure   = &orgMxedgeInventoryResource{}
	_ resource.ResourceWithImportState = &orgMxedgeInventoryResource{}
)

func NewOrgMxedgeInventory() resource.Resource {
	return &orgMxedgeInventoryResource{}
}

type orgMxedgeInventoryResource struct {
	client mistapi.ClientInterface
}

func (r *orgMxedgeInventoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist MxEdge Inventory client")
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

func (r *orgMxedgeInventoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_mxedge_inventory"
}

func (r *orgMxedgeInventoryResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Org MxEdge Inventory.\n\n" +
			"It can be used to claim, assign, unassign, and reassign MxEdges.\n\n" +
			"->Removing an MxEdge from the `mxedges` map will NOT release it from the organization",
		Attributes: resource_org_mxedge_inventory.OrgMxedgeInventoryResourceSchema(ctx).Attributes,
	}
}

func (r *orgMxedgeInventoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting MxEdge Inventory Create")
	var plan, state resource_org_mxedge_inventory.OrgMxedgeInventoryModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge_inventory\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	/////////////////////// Update
	r.updateInventory(&diags, ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	/////////////////////// Sync, required to get missing MxEdge info (ID, Model, Name, ...)
	state = r.refreshInventory(&diags, ctx, &orgId, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeInventoryResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state, comp resource_org_mxedge_inventory.OrgMxedgeInventoryModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var orgIdString string
	// if it's an "Import" process, keep the "comp" var empty and remove the "import." prefix
	// otherwise, set the comp var with the current state
	// the comp var is used to only report information about the MxEdges managed by TF
	if strings.HasPrefix(state.OrgId.ValueString(), "import") {
		orgIdString = strings.Split(state.OrgId.ValueString(), ".")[1]
	} else {
		orgIdString = state.OrgId.ValueString()
		comp = state
	}

	orgId, err := uuid.Parse(orgIdString)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge_inventory\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	state = r.refreshInventory(&diags, ctx, &orgId, &comp)
	resp.Diagnostics.Append(diags...)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeInventoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_mxedge_inventory.OrgMxedgeInventoryModel
	tflog.Info(ctx, "Starting MxEdge Inventory Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_mxedge_inventory\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	/////////////////////// Update
	r.updateInventory(&diags, ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	/////////////////////// Sync, required to get missing MxEdge info (ID, Model, Name, ...)
	state = r.refreshInventory(&diags, ctx, &orgId, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgMxedgeInventoryResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_mxedge_inventory.OrgMxedgeInventoryModel
	tflog.Info(ctx, "Starting MxEdge Inventory Delete")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Note: MxEdges are NOT unclaimed when the resource is destroyed
	// They remain in the organization inventory and are simply unassigned from sites
	tflog.Info(ctx, "MxEdge Inventory resource destroyed. MxEdges remain in the organization and must be manually unclaimed if desired.")
}

func (r *orgMxedgeInventoryResource) updateInventory(
	diags *diag.Diagnostics,
	ctx context.Context,
	orgId *uuid.UUID,
	plan *resource_org_mxedge_inventory.OrgMxedgeInventoryModel,
	state *resource_org_mxedge_inventory.OrgMxedgeInventoryModel,
) {

	claim, unassign, assignClaim, assign, e := resource_org_mxedge_inventory.TerraformToSdk(state, plan)
	if e != nil {
		diags.Append(e...)
		return
	}

	tflog.Debug(ctx, "updateInventory", map[string]interface{}{
		"claim":    strings.Join(claim, ", "),
		"unassign": strings.Join(unassign, ", "),
		"assign":   len(assign),
	})

	/////////////////////// CLAIM
	if len(claim) > 0 {
		r.claimMxedges(diags, ctx, *orgId, claim)
		if diags.HasError() {
			return
		}
		// After claiming, handle assignments for newly claimed MxEdges
		// We need to convert claim codes to MxEdge IDs for assignment
		for claimCode, siteId := range assignClaim {
			// Add to assign map - the refresh will resolve claim codes to MxEdge IDs
			assign[siteId] = append(assign[siteId], claimCode)
		}
		tflog.Debug(ctx, "MxEdges claimed, assignments added to assign map")
	}
	/////////////////////// UNASSIGN
	if len(unassign) > 0 {
		r.unassignMxedges(diags, ctx, *orgId, unassign)
	}
	/////////////////////// ASSIGN
	if len(assign) > 0 {
		r.assignMxedges(diags, ctx, *orgId, assign)
	}
}

func (r *orgMxedgeInventoryResource) refreshInventory(
	diags *diag.Diagnostics,
	ctx context.Context,
	orgId *uuid.UUID,
	refInventory *resource_org_mxedge_inventory.OrgMxedgeInventoryModel,
) (state resource_org_mxedge_inventory.OrgMxedgeInventoryModel) {

	tflog.Info(ctx, "Starting MxEdge Inventory state refresh: org_id "+orgId.String())

	var limit = 1000
	var page = 0
	var total = 9999
	var elements []models.Mxedge
	var forSite *models.MxedgeForSiteEnum

	tflog.Info(ctx, "Starting MxEdge Inventory Read: org_id "+orgId.String())

	for limit*page < total {
		page += 1
		tflog.Debug(ctx, "Pagination Info", map[string]any{
			"page":  page,
			"limit": limit,
			"total": total,
		})

		data, err := r.client.OrgsMxEdges().ListOrgMxEdges(ctx, *orgId, forSite, &limit, &page)
		if err != nil {
			diags.AddError(
				"Error refreshing MxEdge Inventory",
				"Unable to get the MxEdge Inventory, unexpected error: "+err.Error(),
			)
			return state
		}

		limitString := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limitString); err != nil {
			diags.AddError(
				"Error refreshing MxEdge Inventory",
				"Unable to convert the X-Page-Limit value into int, unexpected error: "+err.Error(),
			)
			return state
		}

		totalString := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(totalString); err != nil {
			diags.AddError(
				"Error refreshing MxEdge Inventory",
				"Unable to convert the X-Page-Total value into int, unexpected error: "+err.Error(),
			)
			return state
		}

		elements = append(elements, data.Data...)
	}

	state, e := resource_org_mxedge_inventory.SdkToTerraform(ctx, orgId.String(), &elements, refInventory)
	diags.Append(e...)

	return state
}

func (r *orgMxedgeInventoryResource) claimMxedges(
	diags *diag.Diagnostics,
	ctx context.Context,
	orgId uuid.UUID,
	claim []string,
) {

	tflog.Info(ctx, "Starting to Claim MxEdges")

	// Claim all MxEdges in a single call
	if len(claim) > 0 {
		claimResponse, err := r.client.OrgsMxEdges().ClaimOrgMxEdge(ctx, orgId, claim)

		if claimResponse.Response.StatusCode != 200 {
			apiErr := mistapierror.ProcessApiError(claimResponse.Response.StatusCode, claimResponse.Response.Body, err)
			if apiErr != "" {
				diags.AddError(
					"Error Claiming MxEdges to the Org Inventory",
					fmt.Sprintf("Unable to claim MxEdges %v. %s", claim, apiErr),
				)
			}
		} else if err != nil {
			diags.AddError(
				"Error Claiming MxEdges to the Org Inventory",
				"Unable to claim MxEdges, unexpected error: "+err.Error(),
			)
		} else {
			tflog.Debug(ctx, "Successfully claimed MxEdges", map[string]interface{}{
				"claim_codes": claim,
			})
		}
	}
}

func (r *orgMxedgeInventoryResource) unassignMxedges(
	diags *diag.Diagnostics,
	ctx context.Context,
	orgId uuid.UUID,
	unassign []string,
) {
	tflog.Debug(ctx, "Starting to Unassign MxEdges: ", map[string]interface{}{"mxedge_ids": strings.Join(unassign, ", ")})

	unassignBody := models.MxedgesUnassign{
		MxedgeIds: []uuid.UUID{},
	}

	for _, idStr := range unassign {
		if id, err := uuid.Parse(idStr); err == nil {
			unassignBody.MxedgeIds = append(unassignBody.MxedgeIds, id)
		}
	}

	unassignResponse, err := r.client.OrgsMxEdges().UnassignOrgMxEdgeFromSite(ctx, orgId, &unassignBody)

	if unassignResponse.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(unassignResponse.Response.StatusCode, unassignResponse.Response.Body, err)
		if apiErr != "" {
			diags.AddError(
				"Error when unassigning MxEdges from Sites",
				apiErr,
			)
		}
	} else if err != nil {
		diags.AddError(
			"Error when unassigning MxEdges from Sites",
			"Unable to unassign the MxEdges, unexpected error: "+err.Error(),
		)
	}

	tflog.Debug(ctx, "Response for API Call to unassign MxEdges:", map[string]interface{}{
		"success": unassignResponse.Data.Success,
	})
}

func (r *orgMxedgeInventoryResource) assignMxedges(
	diags *diag.Diagnostics,
	ctx context.Context,
	orgId uuid.UUID,
	assign map[string][]string,
) {
	for siteIdStr, mxedgeIds := range assign {
		tflog.Debug(ctx, "Starting to Assign MxEdges to site "+siteIdStr+": ", map[string]interface{}{"mxedge_ids": strings.Join(mxedgeIds, ", ")})

		siteId, e := uuid.Parse(siteIdStr)
		if e != nil && siteIdStr != "" {
			diags.AddError(
				"Invalid \"site_id\" value for \"mist_org_mxedge_inventory\" resource",
				fmt.Sprintf("Unable to parse the UUID \"%s\": %s", siteIdStr, e.Error()),
			)
			continue
		}

		assignBody := models.MxedgesAssign{
			MxedgeIds: []uuid.UUID{},
			SiteId:    siteId,
		}

		for _, idStr := range mxedgeIds {
			if id, err := uuid.Parse(idStr); err == nil {
				assignBody.MxedgeIds = append(assignBody.MxedgeIds, id)
			}
		}

		assignResponse, err := r.client.OrgsMxEdges().AssignOrgMxEdgeToSite(ctx, orgId, &assignBody)

		if assignResponse.Response.StatusCode != 200 {
			apiErr := mistapierror.ProcessApiError(assignResponse.Response.StatusCode, assignResponse.Response.Body, err)
			if apiErr != "" {
				diags.AddError(
					"Error when assigning MxEdges to a Site",
					apiErr,
				)
			}
		} else if err != nil {
			diags.AddError(
				"Error when assigning MxEdges to a Site",
				"Unable to assign the MxEdges, unexpected error: "+err.Error(),
			)
		}

		tflog.Debug(ctx, "Response for API Call to assign MxEdges:", map[string]interface{}{
			"success": assignResponse.Data.Success,
		})
	}
}

func (r *orgMxedgeInventoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_mxedge_inventory\" resource",
			fmt.Sprintf("Unable to parse the UUID \"%s\": %s. Import \"id\" must be a valid Org Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), "import."+req.ID)...)
}
