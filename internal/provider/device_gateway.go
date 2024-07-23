package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &deviceGatewayResource{}
	_ resource.ResourceWithConfigure = &deviceGatewayResource{}
)

func NewDeviceGatewayResource() resource.Resource {
	return &deviceGatewayResource{}
}

type deviceGatewayResource struct {
	client mistapi.ClientInterface
}

func (r *deviceGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceGateway client")
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
func (r *deviceGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway"
}

func (r *deviceGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Gateway configuration." +
			"It can be used to define specific configuration at the device level or to override Org Gateway template settings.",
		Attributes: resource_device_gateway.DeviceGatewayResourceSchema(ctx).Attributes,
	}
}

func (r *deviceGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceGateway Create")
	var plan, state resource_device_gateway.DeviceGatewayModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway, diags := resource_device_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway site_id from plan",
			"Could not get device_gateway site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway device_id from plan",
			"Could not get device_gateway device_id, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_gateway",
			"Could not create device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)
	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
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

func (r *deviceGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_gateway.DeviceGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Read: device_gateway_id "+state.DeviceId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway site_id from state",
			"Could not get device_gateway site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway device_id from state",
			"Could not get device_gateway device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevices().GetSiteDevice(ctx, siteId, deviceId)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway",
			"Could not get device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)

	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_gateway.DeviceGatewayModel
	tflog.Info(ctx, "Starting DeviceGateway Update")

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

	device_gateway, diags := resource_device_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Update for DeviceGateway "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway site_id from state",
			"Could not get device_gateway site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway device_id from state",
			"Could not get device_gateway device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_gateway",
			"Could not update device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)
	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
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

func (r *deviceGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_gateway.DeviceGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway, e := resource_device_gateway.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Delete: device_gateway_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway site_id from state",
			"Could not get device_gateway site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway device_id from state",
			"Could not get device_gateway device_id, unexpected error: "+err.Error(),
		)
		return
	}

	httpr, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)
	if httpr.Response.StatusCode != 404 && httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device_gateway",
			"Could not delete device_gateway, unexpected error: "+err.Error(),
		)
		return
	}
}
