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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgDeviceprofileApResource{}
	_ resource.ResourceWithConfigure   = &orgDeviceprofileApResource{}
	_ resource.ResourceWithImportState = &orgDeviceprofileApResource{}
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
			"AP Device profiles for aps are used to specify a configuration that can be applied to a select set of aps from any site in the organization. " +
			"They allow for efficient application of configurations based on ap groups, wlan groups, RF settings, and sites. " +
			"Device profiles enable various use cases such as activating ethernet passthrough, applying different rf settings, applying mesh configuration, " +
			"activating specific features like esl or vble, and more.",
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
			"Invalid \"org_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
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
			"Invalid \"org_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Read: deviceprofile_ap_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsDeviceProfiles().GetOrgDeviceProfile(ctx, orgId, deviceprofile_apId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting DeviceprofileAp",
			"Could not get DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
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
			"Invalid \"org_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
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
			"Invalid \"org_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_apId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_ap_id\" value for \"org_deviceprofile_ap\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileAp Delete: deviceprofile_ap_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsDeviceProfiles().DeleteOrgDeviceProfile(ctx, orgId, deviceprofile_apId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DeviceprofileAp",
			"Could not delete DeviceprofileAp, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgDeviceprofileApResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
