package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_deviceprofile_gateway"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgDeviceprofileGatewayResource{}
	_ resource.ResourceWithConfigure   = &orgDeviceprofileGatewayResource{}
	_ resource.ResourceWithImportState = &orgDeviceprofileGatewayResource{}
)

func NewOrgDeviceprofileGateway() resource.Resource {
	return &orgDeviceprofileGatewayResource{}
}

type orgDeviceprofileGatewayResource struct {
	client mistapi.ClientInterface
}

func (r *orgDeviceprofileGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceprofileGateway client")
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
func (r *orgDeviceprofileGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofile_gateway"
}

func (r *orgDeviceprofileGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Gateway Device Profiles (HUB Profiles)." +
			"A HUB profile is a configuration profile that automates the creation of overlay networks and defines the attributes of a hub device in a network. " +
			"It includes settings for wan interfaces, lan interfaces, dns servers, traffic steering preferences, application policies, and routing options. " +
			"HUB profiles are used to create consistent configurations for hub devices and ensure efficient connectivity between hubs and spokes in a network.",
		Attributes: resource_org_deviceprofile_gateway.OrgDeviceprofileGatewayResourceSchema(ctx).Attributes,
	}
}

func (r *orgDeviceprofileGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceprofileGateway Create")
	var plan, state resource_org_deviceprofile_gateway.OrgDeviceprofileGatewayModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceprofile_gateway, diags := resource_org_deviceprofile_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsDeviceProfiles().CreateOrgDeviceProfiles(ctx, orgId, &deviceprofile_gateway)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating DeviceprofileGateway",
			"Could not create DeviceprofileGateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_deviceprofile_gateway := models.DeviceprofileGateway{}
	err = json.Unmarshal(body, &mist_deviceprofile_gateway)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_gateway.SdkToTerraform(ctx, &mist_deviceprofile_gateway)
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

func (r *orgDeviceprofileGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_deviceprofile_gateway.OrgDeviceprofileGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_gatewayId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_gateway_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileGateway Read: deviceprofile_gateway_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsDeviceProfiles().GetOrgDeviceProfile(ctx, orgId, deviceprofile_gatewayId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting DeviceprofileGateway",
			"Could not get DeviceprofileGateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
	mist_deviceprofile_gateway := models.DeviceprofileGateway{}
	err = json.Unmarshal(body, &mist_deviceprofile_gateway)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_gateway.SdkToTerraform(ctx, &mist_deviceprofile_gateway)
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

func (r *orgDeviceprofileGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_deviceprofile_gateway.OrgDeviceprofileGatewayModel
	tflog.Info(ctx, "Starting DeviceprofileGateway Update")

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

	deviceprofile_gateway, diags := resource_org_deviceprofile_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_gatewayId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_gateway_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileGateway Update for DeviceprofileGateway "+state.Id.ValueString())
	data, err := r.client.OrgsDeviceProfiles().UpdateOrgDeviceProfile(ctx, orgId, deviceprofile_gatewayId, &deviceprofile_gateway)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error updating DeviceprofileGateway",
			"Could not update DeviceprofileGateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_deviceprofile_gateway := models.DeviceprofileGateway{}
	err = json.Unmarshal(body, &mist_deviceprofile_gateway)
	if err != nil {
		resp.Diagnostics.AddError("Unable to unMarshal API response", err.Error())
		return
	}

	state, diags = resource_org_deviceprofile_gateway.SdkToTerraform(ctx, &mist_deviceprofile_gateway)
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

func (r *orgDeviceprofileGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_deviceprofile_gateway.OrgDeviceprofileGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	deviceprofile_gatewayId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"deviceprofile_gateway_id\" value for \"org_deviceprofile_gateway\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceprofileGateway Delete: deviceprofile_gateway_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsDeviceProfiles().DeleteOrgDeviceProfile(ctx, orgId, deviceprofile_gatewayId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting DeviceprofileGateway",
			"Could not delete DeviceprofileGateway, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgDeviceprofileGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

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
