package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Juniper/terraform-provider-mist/internal/resource_device_ap"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &deviceApResource{}
	_ resource.ResourceWithConfigure = &deviceApResource{}
)

func NewDeviceApResource() resource.Resource {
	return &deviceApResource{}
}

type deviceApResource struct {
	client mistapi.ClientInterface
}

func (r *deviceApResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceAp client")
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
func (r *deviceApResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ap"
}

func (r *deviceApResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Wireless Access Point configuration." +
			"It can be used to define specific configuration at the device level or to override AP Device Profile (`mist_org_deviceprofile_ap`).",
		Attributes: resource_device_ap.DeviceApResourceSchema(ctx).Attributes,
	}
}

func (r *deviceApResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceAp Create")
	var plan, state resource_device_ap.DeviceApModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_ap, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap site_id from plan",
			"Could not get device_ap site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap device_id from plan",
			"Could not get device_ap device_id, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_ap)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_ap",
			"Could not create device_ap, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_ap := models.DeviceAp{}
	json.Unmarshal(body, &mist_ap)
	state, diags = resource_device_ap.SdkToTerraform(ctx, &mist_ap)
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

func (r *deviceApResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_ap.DeviceApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Read: device_ap_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap site_id from state",
			"Could not get device_ap site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap device_id from state",
			"Could not get device_ap device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevices().GetSiteDevice(ctx, siteId, deviceId)
	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap",
			"Could not get device_ap, unexpected error: "+err.Error(),
		)
		return
	}
	body, _ := io.ReadAll(data.Response.Body)
	mist_ap := models.DeviceAp{}
	json.Unmarshal(body, &mist_ap)

	state, diags = resource_device_ap.SdkToTerraform(ctx, &mist_ap)
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

func (r *deviceApResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_ap.DeviceApModel
	tflog.Info(ctx, "Starting DeviceAp Update")

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

	device_ap, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Update for DeviceAp "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap site_id from state",
			"Could not get device_ap site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap device_id from state",
			"Could not get device_ap device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_ap)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_ap",
			"Could not update device_ap, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_ap := models.DeviceAp{}
	json.Unmarshal(body, &mist_ap)
	state, diags = resource_device_ap.SdkToTerraform(ctx, &mist_ap)
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

func (r *deviceApResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_ap.DeviceApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_ap, e := resource_device_ap.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Delete: device_ap_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap site_id from state",
			"Could not get device_ap site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_ap device_id from state",
			"Could not get device_ap device_id, unexpected error: "+err.Error(),
		)
		return
	}

	httpr, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_ap)
	if httpr.Response.StatusCode != 404 && httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device_ap",
			"Could not delete device_ap, unexpected error: "+err.Error(),
		)
		return
	}
}
