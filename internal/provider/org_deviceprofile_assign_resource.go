package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_assign"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgDeviceprofileAssignResource{}
	_ resource.ResourceWithConfigure   = &orgDeviceprofileAssignResource{}
	_ resource.ResourceWithImportState = &orgDeviceprofileAssignResource{}
)

func NewOrgDeviceprofileAssign() resource.Resource {
	return &orgDeviceprofileAssignResource{}
}

type orgDeviceprofileAssignResource struct {
	client mistapi.ClientInterface
}

func (r *orgDeviceprofileAssignResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceprofileAssign client")
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
func (r *orgDeviceprofileAssignResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofile_assign"
}

func (r *orgDeviceprofileAssignResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resourceis used to assign/unassign a device profile to one or multiple devices" +
			"The `mist_org_deviceprofile_gateway` resource can be assigned to Gateways" +
			"The `mist_org_deviceprofile_ap` resource can be assigned the Wireless Access Points",
		Attributes: resource_org_deviceprofile_assign.OrgDeviceprofileAssignResourceSchema(ctx).Attributes,
	}
}

func (r *orgDeviceprofileAssignResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceprofileAssign Create")
	var plan, state resource_org_deviceprofile_assign.OrgDeviceprofileAssignModel
	var macs_assigned_success, macs_unassigned_success []string

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	deviceprofileId, err := uuid.Parse(plan.DeviceprofileId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	macs_to_assign, _, diags := resource_org_deviceprofile_assign.TerraformToSdk(ctx, plan.Macs, state.Macs)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if len(macs_to_assign.Macs) > 0 {
		macs_assigned_success, err = r.assign(ctx, orgId, deviceprofileId, macs_to_assign)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Unassigning devices to Deviceprofile",
				"Unable to unassign the devices to Deviceprofile, unexpected error: "+err.Error(),
			)
			return
		}
	}

	new_macs_state, diags := resource_org_deviceprofile_assign.SdkToTerraform(ctx, &state.Macs, macs_assigned_success, macs_unassigned_success)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Macs = *new_macs_state
	state.DeviceprofileId = plan.DeviceprofileId
	state.OrgId = plan.OrgId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgDeviceprofileAssignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_deviceprofile_assign.OrgDeviceprofileAssignModel

	// TODO: really read the device info to get the profile assignements
	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileId, err := uuid.Parse(state.DeviceprofileId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceprofileId.ValueString(), err.Error()),
		)
		return
	}

	var serial string
	var model string
	var mType models.DeviceTypeEnum
	var mac string
	var siteId string
	var vcMac string
	var vc bool = false
	var unassigned bool
	var limit int = 1000
	var page int
	tflog.Info(ctx, "Starting Inventory Read: org_id  "+orgId.String())
	data, err := r.client.OrgsInventory().GetOrgInventory(ctx, orgId, &serial, &model, &mType, &mac, &siteId, &vcMac, &vc, &unassigned, &limit, &page)
	if err != nil {
		diags.AddError(
			"Error refreshing Inventory",
			"Unable to get the Inventory, unexpected error: "+err.Error(),
		)
	}

	var macs []types.String
	for _, dev := range data.Data {
		if dev.DeviceprofileId.Value() != nil && *dev.DeviceprofileId.Value() == deviceprofileId.String() {
			macs = append(macs, types.StringValue(*dev.Mac))
		}
	}
	tmp, e := types.ListValueFrom(ctx, types.StringType, macs)
	if e != nil {
		resp.Diagnostics.Append(e...)
	} else {
		state.Macs = tmp
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgDeviceprofileAssignResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_deviceprofile_assign.OrgDeviceprofileAssignModel
	var macs_assigned_success, macs_unassigned_success []string
	tflog.Info(ctx, "Starting DeviceprofileAssign Update")

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

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileId, err := uuid.Parse(state.DeviceprofileId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceprofileId.ValueString(), err.Error()),
		)
		return
	}

	macs_to_assign, macs_to_unassign, e := resource_org_deviceprofile_assign.TerraformToSdk(ctx, plan.Macs, state.Macs)
	if e != nil {
		diags.Append(e...)
		return
	}

	if len(macs_to_assign.Macs) > 0 {
		macs_assigned_success, err = r.assign(ctx, orgId, deviceprofileId, macs_to_assign)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Assigning devices to Deviceprofile",
				"Unable to Assign the Device Profile to the Devices, unexpected error: "+err.Error(),
			)
			return
		}
	}

	if len(macs_to_unassign.Macs) > 0 {
		macs_unassigned_success, err = r.unassign(ctx, orgId, deviceprofileId, macs_to_unassign)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Unassigning devices to Deviceprofile",
				"Unable to unassign the Device Profile from the Devices, unexpected error: "+err.Error(),
			)
			return
		}
	}

	new_macs_state, diags := resource_org_deviceprofile_assign.SdkToTerraform(ctx, &state.Macs, macs_assigned_success, macs_unassigned_success)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Macs = *new_macs_state
	state.DeviceprofileId = plan.DeviceprofileId
	state.OrgId = plan.OrgId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgDeviceprofileAssignResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_deviceprofile_assign.OrgDeviceprofileAssignModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	deviceprofileId, err := uuid.Parse(state.DeviceprofileId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceprofileId.ValueString(), err.Error()),
		)
		return
	}

	plan_macs := types.ListNull(types.StringType)
	_, macs_to_unassign, e := resource_org_deviceprofile_assign.TerraformToSdk(ctx, plan_macs, state.Macs)
	if e != nil {
		diags.Append(e...)
		return
	}

	_, err = r.unassign(ctx, orgId, deviceprofileId, macs_to_unassign)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_deviceprofile_assign\" resource",
			"Unable to delete the Deviceprofile Assignment, unexpected error: "+err.Error(),
		)
		return
	}

}

func (r *orgDeviceprofileAssignResource) assign(ctx context.Context, orgId uuid.UUID, deviceprofileId uuid.UUID, macs models.MacAddresses) ([]string, error) {
	if len(macs.Macs) > 0 {
		tflog.Info(ctx, "Assigning "+string(len(macs.Macs))+"to  deviceprofile_id "+deviceprofileId.String())
		data, err := r.client.OrgsDeviceProfiles().AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, &macs)

		return data.Data.Success, err

	} else {
		return nil, nil
	}
}

func (r *orgDeviceprofileAssignResource) unassign(ctx context.Context, orgId uuid.UUID, deviceprofileId uuid.UUID, macs models.MacAddresses) ([]string, error) {
	if len(macs.Macs) > 0 {
		tflog.Info(ctx, "Unassigning "+string(len(macs.Macs))+"to  deviceprofile_id "+deviceprofileId.String())
		data, err := r.client.OrgsDeviceProfiles().UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, &macs)

		return data.Data.Success, err

	} else {
		return nil, nil
	}
}

func (r *orgDeviceprofileAssignResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_deviceprofile_assign\" resource",
			"import \"id\" format must be \"{org_id}.{deviceprofile_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{deviceprofile_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_deviceprofile_assign\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{deviceprofile_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("deviceprofile_id"), importIds[1])...)
}
