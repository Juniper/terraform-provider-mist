package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_ap"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgDeviceprofileApResource{}
	_ resource.ResourceWithConfigure = &orgDeviceprofileApResource{}
)

func NewOrgDeviceprofileAp() resource.Resource {
	return &orgDeviceprofileApResource{}
}

type orgDeviceprofileApResource struct {
	client mistapi.ClientInterface
}

func (r *orgDeviceprofileApResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceprofileAp client")
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
func (r *orgDeviceprofileApResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofile_ap"
}

func (r *orgDeviceprofileApResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the AP Device Profiles." +
			"An AP Device Profile can be used to define common configuration accross multiple Wireless Access Points" +
			"and is assigned to one or multiple APs as a deviceprofile with the `mist_org_deviceprofile_assign` resource",
		Attributes: resource_org_deviceprofile_ap.OrgDeviceprofileApResourceSchema(ctx).Attributes,
	}
}

func (r *orgDeviceprofileApResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceprofileAp Create")
	var plan, state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceprofile_ap, diags := resource_org_deviceprofile_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap org_id from plan",
			"Could not get org_deviceprofile_ap org_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsDeviceProfiles().CreateOrgDeviceProfiles(ctx, orgId, &deviceprofile_ap)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating DeviceprofileAp",
			"Could not create DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_deviceprofile_ap := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mist_deviceprofile_ap)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mist_deviceprofile_ap)
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

func (r *orgDeviceprofileApResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap org_id from state",
			"Could not get org_deviceprofile_ap org_id, unexpected error: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap deviceprofile_ap_id from state",
			"Could not get org_deviceprofile_ap deviceprofile_ap_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Read: deviceprofile_ap_id "+state.Id.ValueString())
	data, err := r.client.OrgsDeviceProfiles().GetOrgDeviceProfile(ctx, orgId, deviceprofile_apId)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting DeviceprofileAp",
			"Could not get DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_deviceprofile_ap := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mist_deviceprofile_ap)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mist_deviceprofile_ap)
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

func (r *orgDeviceprofileApResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_deviceprofile_ap.OrgDeviceprofileApModel
	tflog.Info(ctx, "Starting DeviceprofileAp Update")

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

	deviceprofile_ap, diags := resource_org_deviceprofile_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap org_id from state",
			"Could not get org_deviceprofile_ap org_id, unexpected error: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap deviceprofile_ap_id from state",
			"Could not get org_deviceprofile_ap deviceprofile_ap_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Update for DeviceprofileAp "+state.Id.ValueString())
	data, err := r.client.OrgsDeviceProfiles().UpdateOrgDeviceProfile(ctx, orgId, deviceprofile_apId, &deviceprofile_ap)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error updating DeviceprofileAp",
			"Could not update DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_deviceprofile_ap := models.DeviceprofileAp{}
	err = json.Unmarshal(body, &mist_deviceprofile_ap)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_ap.SdkToTerraform(ctx, &mist_deviceprofile_ap)
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

func (r *orgDeviceprofileApResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_deviceprofile_ap.OrgDeviceprofileApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap org_id from state",
			"Could not get org_deviceprofile_ap org_id, unexpected error: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_deviceprofile_ap deviceprofile_ap_id from state",
			"Could not get org_deviceprofile_ap deviceprofile_ap_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Delete: deviceprofile_ap_id "+state.Id.ValueString())
	_, err = r.client.OrgsDeviceProfiles().DeleteOrgDeviceProfile(ctx, orgId, deviceprofile_apId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DeviceprofileAp",
			"Could not delete DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}
}
