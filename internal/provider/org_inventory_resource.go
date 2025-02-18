package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_inventory"

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
		MarkdownDescription: docCategoryDevices + "This resource manages the Org Inventory.\n\n" +
			"It can be used to claim, unclaim, assign, unassign, reassign devices.\n\n" +
			"->Removing a device from the `devices` list or `inventory` map will NOT release it unless `unclaim_when_destroyed` is set to `true`\n\n" +
			"!> The `devices` attribute (List) is deprecated and is replaced by the `inventory` attribute (Map) as " +
			"it can generate \"inconsistent result after apply\" errors. If this happens, it is required to force a refresh of the " +
			"state to synchronise the new list.  \n" +
			"The `devices` attribute will generate inconsistent result after apply " +
			"when a device other than the last one is removed from the list or " +
			"when a device is added somewhere other than the end of the list",
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
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	/////////////////////// Update
	diags = r.updateInventory(ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)

	/////////////////////// Sync, required to get missing devices info (MAC, Serial, ...)
	state, diags = r.refreshInventory(ctx, &orgId, &plan)
	resp.Diagnostics.Append(diags...)

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
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	state, diags = r.refreshInventory(ctx, &orgId, &comp)
	resp.Diagnostics.Append(diags...)

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
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	/////////////////////// Update
	diags = r.updateInventory(ctx, &orgId, &plan, &state)
	resp.Diagnostics.Append(diags...)

	/////////////////////// Sync, required to get missing devices info (MAC, Serial, ...)
	state, diags = r.refreshInventory(ctx, &orgId, &plan)
	resp.Diagnostics.Append(diags...)

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

	// var serials []string
	// for _, v := range state.Devices.Elements() {
	// 	var vi interface{} = v
	// 	var dev_state = vi.(resource_org_inventory.DevicesValue)
	// 	serials = append(serials, dev_state.Serial.ValueString())
	// }
	serials, diags := resource_org_inventory.DeleteOrgInventory(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
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
			"Unable to unclaim the devices, unexpected error: "+err.Error(),
		)
	}
	tflog.Debug(ctx, "response for API Call to unclaim devices:", map[string]interface{}{
		"Error":   strings.Join(unclaim_response.Data.Error, ", "),
		"Reason":  strings.Join(unclaim_response.Data.Reason, ", "),
		"Success": strings.Join(unclaim_response.Data.Success, ", "),
	})
}

func (r *orgInventoryResource) updateInventory(ctx context.Context, orgId *uuid.UUID, plan *resource_org_inventory.OrgInventoryModel, state *resource_org_inventory.OrgInventoryModel) diag.Diagnostics {
	var diags diag.Diagnostics

	claim, unclaim, unassign, assign_claim, assign, e := resource_org_inventory.TerraformToSdk(ctx, state, plan)
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
			site_id, ok := assign_claim[strings.ToUpper(v.Magic)]
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

func (r *orgInventoryResource) refreshInventory(ctx context.Context, orgId *uuid.UUID, ref_inventory *resource_org_inventory.OrgInventoryModel) (resource_org_inventory.OrgInventoryModel, diag.Diagnostics) {
	var state resource_org_inventory.OrgInventoryModel
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
	var page int = 0
	var total int = 9999
	var elements []models.Inventory

	tflog.Info(ctx, "Starting Inventory Read: org_id  "+orgId.String())

	for limit*page < total {
		page += 1
		tflog.Debug(ctx, "Pagination Info", map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		})
		// Read API call logic
		data, err := r.client.OrgsInventory().GetOrgInventory(ctx, *orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
		if err != nil {
			diags.AddError(
				"Error refreshing Inventory",
				"Unable to get the Inventory, unexpected error: "+err.Error(),
			)
			return state, diags
		}

		limit_string := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limit_string); err != nil {
			diags.AddError(
				"Error refreshing Inventory",
				"Unable to convert the X-Page-Limit value into int, unexcpected error: "+err.Error(),
			)
			return state, diags
		}

		total_string := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(total_string); err != nil {
			diags.AddError(
				"Error refreshing Inventory",
				"Unable to convert the X-Page-Total value into int, unexcpected error: "+err.Error(),
			)
			return state, diags
		}

		elements = append(elements, data.Data...)
	}

	state, e := resource_org_inventory.SdkToTerraform(ctx, orgId.String(), &elements, ref_inventory)
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
			fmt.Sprintf("Unable to claim the devices %s. API Response: %s", claim_code, reason),
		)
	}
}

func (r *orgInventoryResource) claimDevices(ctx context.Context, orgId uuid.UUID, claim []string, diags *diag.Diagnostics) []models.ResponseInventoryInventoryAddedItems {

	tflog.Info(ctx, "Starting to Claim devices")
	claim_response, err := r.client.OrgsInventory().AddOrgInventory(ctx, orgId, claim)

	api_err := mist_api_error.ProcessInventoryApiError(ctx, "claim", claim_response.Response.StatusCode, claim_response.Response.Body, err)
	if len(api_err) > 0 {
		for _, err_value := range api_err {
			diags.AddError(
				"Error Claiming Devices to the Org Inventory",
				err_value,
			)
		}
	} else if err != nil {
		diags.AddError(
			"Error Claiming Devices to the Org Inventory",
			"Unable to claim the devices, unexpected error: "+err.Error(),
		)
		return nil
	}

	logResponseInventory(ctx, "Success response for API Call to claim devices:", claim_response.Data)

	if len(claim_response.Data.InventoryDuplicated) > 0 {
		for _, duplicatedDevice := range claim_response.Data.InventoryDuplicated {
			diags.AddWarning("Duplicated Device", fmt.Sprintf("Device %s was already claimed (MAC: %s, Serial: %s, Model: %s). It has been imported into the Inventory state.", duplicatedDevice.Magic, duplicatedDevice.Mac, duplicatedDevice.Serial, duplicatedDevice.Model))
			var tmp models.ResponseInventoryInventoryAddedItems
			tmp.Mac = duplicatedDevice.Mac
			tmp.Magic = duplicatedDevice.Magic
			tmp.Model = duplicatedDevice.Model
			tmp.Serial = duplicatedDevice.Serial
			tmp.Type = duplicatedDevice.Type
			claim_response.Data.InventoryAdded = append(claim_response.Data.InventoryAdded, tmp)
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

	api_err := mist_api_error.ProcessInventoryApiError(ctx, "unclaim", unclaim_response.Response.StatusCode, unclaim_response.Response.Body, err)
	if len(api_err) > 0 {
		for _, err_value := range api_err {
			diags.AddError(
				"Error when releasing Devices from the Org Inventory",
				err_value,
			)
		}
	} else if err != nil {
		diags.AddError(
			"Error when releasing Devices from the Org Inventory",
			"Unable to unclaim the devices, unexpected error: "+err.Error(),
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

	api_err := mist_api_error.ProcessInventoryApiError(ctx, "unassign", unassign_response.Response.StatusCode, unassign_response.Response.Body, err)
	if len(api_err) > 0 {
		for _, err_value := range api_err {
			diags.AddError(
				"Error when unassigning Devices from a Site to the Org Inventory",
				err_value,
			)
		}
	} else if err != nil {
		diags.AddError(
			"Error when unassigning Devices from a Site to the Org Inventory",
			"Unable to unassign the devices, unexpected error: "+err.Error(),
		)
	}

	tflog.Debug(ctx, "response for API Call to claim devices:", map[string]interface{}{
		"Error":   strings.Join(unassign_response.Data.Error, ", "),
		"Reason":  strings.Join(unassign_response.Data.Reason, ", "),
		"Success": strings.Join(unassign_response.Data.Success, ", "),
	})
}

func (r *orgInventoryResource) assignDevices(ctx context.Context, orgId uuid.UUID, assign map[string][]string, diags *diag.Diagnostics) {
	for k, v := range assign {
		tflog.Debug(ctx, "Starting to Assign devices to site "+k+": ", map[string]interface{}{"macs": strings.Join(v, ", ")})

		tflog.Info(ctx, "devices "+strings.Join(assign[k], ", ")+" to "+k)
		siteId, err := uuid.Parse(k)
		if err != nil && k != "" {
			diags.AddError(
				"Invalid \"site_id\" value for \"org_inventory\" resource",
				fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", siteId.String(), err.Error()),
			)
		} else {
			body := models.InventoryUpdate{
				DisableAutoConfig: models.ToPointer(false),
				Macs:              v,
				Managed:           models.ToPointer(true),
				NoReassign:        models.ToPointer(false),
				Op:                models.InventoryUpdateOperationEnum("assign"),
				SiteId:            models.ToPointer(siteId),
			}
			assign_response, err := r.client.OrgsInventory().UpdateOrgInventoryAssignment(ctx, orgId, &body)

			api_err := mist_api_error.ProcessInventoryApiError(ctx, "assign", assign_response.Response.StatusCode, assign_response.Response.Body, err)
			if len(api_err) > 0 {
				for _, err_value := range api_err {
					diags.AddError(
						"Error when assigning Devices from the Org Inventory to a Site",
						err_value,
					)
				}
			} else if err != nil {
				diags.AddError(
					"Error when assigning Devices from the Org Inventory to a Site",
					"Unable to assign the devices, unexpected error: "+err.Error(),
				)
			}

			tflog.Debug(ctx, "response for API Call to assign devices:", map[string]interface{}{
				"Error":   strings.Join(assign_response.Data.Error, ", "),
				"Reason":  strings.Join(assign_response.Data.Reason, ", "),
				"Success": strings.Join(assign_response.Data.Success, ", "),
			})
		}
	}
}

func (r *orgInventoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_inventory\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" must be a valid Org Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), "import."+req.ID)...)

}
