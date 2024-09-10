package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgInventoryResource{}
	_ resource.ResourceWithConfigure   = &orgInventoryResource{}
	_ resource.ResourceWithImportState = &orgInventoryResource{}
)

func NewOrgInventory() resource.Resource {
	return &orgInventoryResource{}
}

type orgInventoryResource struct {
	client mistapi.ClientInterface
}

func (r *orgInventoryResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Inventory client")
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
func (r *orgInventoryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_inventory"
}

func (r *orgInventoryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Org inventory." +
			"It can be used to claim, unclaim, assign, unassign, reassign devices",
		Attributes: resource_org_inventory.OrgInventoryResourceSchema(ctx).Attributes,
	}
}

func (r *orgInventoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Inventory Create")
	var plan, state resource_org_inventory.OrgInventoryModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	/////////////////////// Update
	diags = r.updateInventory(ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	/////////////////////// Sync, required to get missing devices info (MAC, Serial, ...)
	state, diags = r.refreshInventory(ctx, &orgId, &plan)
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

func (r *orgInventoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state, comp resource_org_inventory.OrgInventoryModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var orgIdString string
	// if it's an "Import" process, keep the "comp" var empty and remove the "import." prefic
	// otherwise, set the comp var with the current state
	// the comp var is used to only report information about the devices managed by TF. This
	// is currently the only way to be sure to not delete them from the Org
	if strings.HasPrefix(state.OrgId.ValueString(), "import") {
		orgIdString = strings.Split(state.OrgId.ValueString(), ".")[1]
	} else {
		orgIdString = state.OrgId.ValueString()
		comp = state
	}

	orgId, err := uuid.Parse(orgIdString)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	state, diags = r.refreshInventory(ctx, &orgId, &comp)
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

func (r *orgInventoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_inventory.OrgInventoryModel
	tflog.Info(ctx, "Starting Inventory Update")

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
			"Invalid \"org_id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	/////////////////////// Update
	diags = r.updateInventory(ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	/////////////////////// Sync, required to get missing devices info (MAC, Serial, ...)
	state, diags = r.refreshInventory(ctx, &orgId, &plan)
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

func (r *orgInventoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_inventory.OrgInventoryModel
	tflog.Info(ctx, "Starting Inventory Delete: wxtag_id ")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var serials []string
	for _, v := range state.Devices.Elements() {
		var vi interface{} = v
		var dev_state = vi.(resource_org_inventory.DevicesValue)
		serials = append(serials, dev_state.Serial.ValueString())
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	unclaim_body := models.InventoryUpdate{}
	unclaim_body.Op = models.InventoryUpdateOperationEnum_DELETE
	unclaim_body.Serials = serials
	unclaim_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unclaim_body)
	if unclaim_response.Response.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error Unclaiming Devices from the Org Inventory",
			"Could not Unclaim devices, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Debug(ctx, "response for API Call to unclaim devices:", map[string]interface{}{
		"Error":   strings.Join(unclaim_response.Data.Error, ", "),
		"Reason":  strings.Join(unclaim_response.Data.Reason, ", "),
		"Success": strings.Join(unclaim_response.Data.Success, ", "),
	})
}

func (r *orgInventoryResource) updateInventory(ctx context.Context, orgId *uuid.UUID, plan *resource_org_inventory.OrgInventoryModel, state *resource_org_inventory.OrgInventoryModel) diag.Diagnostics {
	var diags diag.Diagnostics

	claim, unclaim, unassign, assign_claim, assign, e := resource_org_inventory.TerraformToSdk(ctx, &plan.Devices, &state.Devices)
	if e != nil {
		diags.Append(e...)
		return diags
	}

	tflog.Debug(ctx, "updateInventory", map[string]interface{}{
		"claim":    strings.Join(claim, ", "),
		"unclaim":  strings.Join(unclaim, ", "),
		"unassign": strings.Join(unassign, ", "),
		"assign":   len(assign),
	})

	/////////////////////// CLAIM
	if len(claim) > 0 {
		inventory_added := r.claimDevices(ctx, *orgId, claim, &diags)
		if diags.HasError() {
			return diags
		}
		for _, v := range inventory_added {
			site_id, ok := assign_claim[v.Magic]
			if ok {
				assign[site_id] = append(assign[site_id], v.Mac)
			}
		}
	}
	/////////////////////// UNCLAIM
	if len(unclaim) > 0 {
		r.unclaimDevices(ctx, *orgId, unclaim, &diags)
	}
	/////////////////////// UNASSIGN
	if len(unassign) > 0 {
		r.unassignDevices(ctx, *orgId, unassign, &diags)
	}
	/////////////////////// ASSIGN
	if len(assign) > 0 {
		r.assignDevices(ctx, *orgId, assign, &diags)
	}
	return diags
}

func (r *orgInventoryResource) refreshInventory(ctx context.Context, orgId *uuid.UUID, plan *resource_org_inventory.OrgInventoryModel) (resource_org_inventory.OrgInventoryModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	tflog.Info(ctx, "Starting Inventory state refresh: org_id  "+orgId.String())
	var serial string
	var model string
	var mType models.DeviceTypeEnum
	var mac string
	var siteId string
	var vcMac string
	var vc bool = true
	var unassigned bool
	var limit int = 1000
	var page int
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+orgId.String())
	data, err := r.client.OrgsInventory().GetOrgInventory(ctx, *orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
	if err != nil {
		diags.AddError(
			"Error refreshing Inventory",
			"Could not get Inventory, unexpected error: "+err.Error(),
		)
	}
	tflog.Debug(ctx, "refreshInventory", map[string]interface{}{"inventory": data.Data})
	state, e := resource_org_inventory.SdkToTerraform(ctx, orgId.String(), data.Data, plan)
	diags.Append(e...)

	return state, diags
}

func logResponseInventory(ctx context.Context, message string, response models.ResponseInventory) {

	tflog.Debug(ctx, message, map[string]interface{}{
		"Error":                strings.Join(response.Error, ", "),
		"added":                strings.Join(response.Added, ", "),
		"duplicated":           strings.Join(response.Duplicated, ", "),
		"inventory added":      response.InventoryAdded,
		"inventory duplicated": response.InventoryDuplicated,
		"reason":               response.Reason,
	})
}

func processResponseInventoryError(response models.ResponseInventory, diags *diag.Diagnostics) {
	for i, claim_code := range response.Error {
		reason := response.Reason[i]

		diags.AddError(
			"Error Claiming Devices to the Org Inventory",
			fmt.Sprintf("Could not Claim devices %s. API Response: %s", claim_code, reason),
		)
	}
}

func (r *orgInventoryResource) claimDevices(ctx context.Context, orgId uuid.UUID, claim []string, diags *diag.Diagnostics) []models.ResponseInventoryInventoryAddedItems {

	tflog.Info(ctx, "Starting to Claim devices")
	claim_response, err := r.client.OrgsInventory().AddOrgInventory(ctx, orgId, claim)

	if claim_response.Response.StatusCode == 400 {
		bodyBytes, _ := io.ReadAll(claim_response.Response.Body)
		error_response := models.ResponseInventory{}
		json.Unmarshal(bodyBytes, &error_response)

		logResponseInventory(ctx, "Error response for API Call to claim devices:", error_response)
		processResponseInventoryError(error_response, diags)

		return nil
	} else if err != nil {
		diags.AddError(
			"Error Claiming Devices to the Org Inventory",
			"Could not Claim devices, unexpected error: "+err.Error(),
		)
		return nil
	}

	logResponseInventory(ctx, "Success response for API Call to claim devices:", claim_response.Data)

	if len(claim_response.Data.Duplicated) > 0 {
		for _, claim_code := range claim_response.Data.Duplicated {
			diags.AddWarning("Duplicated Device", fmt.Sprintf("Device %s was already claimed. It has been added to the Inventory state.", claim_code))
		}
	}
	processResponseInventoryError(claim_response.Data, diags)
	return claim_response.Data.InventoryAdded
}

func (r *orgInventoryResource) unclaimDevices(ctx context.Context, orgId uuid.UUID, unclaim []string, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "Starting to Unclaim devicesdevices: ", map[string]interface{}{"macs": strings.Join(unclaim, ", ")})

	unclaim_body := models.InventoryUpdate{}
	unclaim_body.Op = models.InventoryUpdateOperationEnum_DELETE
	unclaim_body.Serials = unclaim
	unclaim_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unclaim_body)
	if err != nil {
		diags.AddError(
			"Error Unclaiming Devices from the Org Inventory",
			"Could not Unclaim devices, unexpected error: "+err.Error(),
		)
	}
	tflog.Debug(ctx, "response for API Call to unclaim devices:", map[string]interface{}{
		"Error":   strings.Join(unclaim_response.Data.Error, ", "),
		"Reason":  strings.Join(unclaim_response.Data.Reason, ", "),
		"Success": strings.Join(unclaim_response.Data.Success, ", "),
	})
}

func (r *orgInventoryResource) unassignDevices(ctx context.Context, orgId uuid.UUID, unassign []string, diags *diag.Diagnostics) {
	tflog.Debug(ctx, "Starting to Unassign devices: ", map[string]interface{}{"macs": strings.Join(unassign, ", ")})

	unassign_body := models.InventoryUpdate{}
	unassign_body.Op = models.InventoryUpdateOperationEnum_UNASSIGN
	unassign_body.Macs = unassign
	unassign_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &unassign_body)
	tflog.Debug(ctx, "response for API Call to claim devices:", map[string]interface{}{
		"Error":   strings.Join(unassign_response.Data.Error, ", "),
		"Reason":  strings.Join(unassign_response.Data.Reason, ", "),
		"Success": strings.Join(unassign_response.Data.Success, ", "),
	})

	if err != nil {
		diags.AddError(
			"Error Unassigning Devices from the Org Inventory",
			"Could not Unassign devices, unexpected error: "+err.Error(),
		)
	}
}

func (r *orgInventoryResource) assignDevices(ctx context.Context, orgId uuid.UUID, assign map[string][]string, diags *diag.Diagnostics) {
	for k, v := range assign {
		tflog.Debug(ctx, "Starting to Assign devices to site "+k+": ", map[string]interface{}{"macs": strings.Join(v, ", ")})

		assign_body := models.InventoryUpdate{}
		assign_body.Op = models.InventoryUpdateOperationEnum_ASSIGN
		assign_body.Macs = v
		assign_body.DisableAutoConfig = types.BoolValue(false).ValueBoolPointer()
		assign_body.Managed = types.BoolValue(true).ValueBoolPointer()
		tflog.Info(ctx, "devices "+strings.Join(assign[k], ", ")+" to "+k)
		siteId, err := uuid.Parse(k)
		if err != nil {
			diags.AddError(
				"Invalid \"site_id\" value for \"org_inventory\" resource",
				fmt.Sprintf("Could not parse the UUID \"%s\": %s", orgId.String(), err.Error()),
			)
		}
		assign_body.SiteId = models.ToPointer(siteId)

		assign_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &assign_body)
		if err != nil {
			diags.AddError(
				"Error Assigning Devices to the Org Inventory",
				"Could not Assign devices, unexpected error: "+err.Error(),
			)
		}
		tflog.Debug(ctx, "response for API Call to assign devices:", map[string]interface{}{
			"Error":   strings.Join(assign_response.Data.Error, ", "),
			"Reason":  strings.Join(assign_response.Data.Reason, ", "),
			"Success": strings.Join(assign_response.Data.Success, ", "),
		})
	}
}

func (r *orgInventoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" must be a valid Org Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), "import."+req.ID)...)

}
