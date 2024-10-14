package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_switch"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceSwitchResource{}
	_ resource.ResourceWithConfigure   = &deviceSwitchResource{}
	_ resource.ResourceWithImportState = &deviceSwitchResource{}
)

func NewDeviceSwitchResource() resource.Resource {
	return &deviceSwitchResource{}
}

type deviceSwitchResource struct {
	client mistapi.ClientInterface
}

func (r *deviceSwitchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceSwitch client")
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
func (r *deviceSwitchResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_switch"
}

func (r *deviceSwitchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Switch configuration.\n" +
			"It can be used to define specific configuration at the device level or to override Org/Site Network template settings.",
		Attributes: resource_device_switch.DeviceSwitchResourceSchema(ctx).Attributes,
	}
}

func (r *deviceSwitchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceSwitch Create")
	var plan, state resource_device_switch.DeviceSwitchModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_switch, diags := resource_device_switch.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_switch)

	if data.Response.StatusCode != 200 {
		api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
		if api_err != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_device_switch\" resource",
				fmt.Sprintf("Unable to create the Switch. %s", api_err),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_switch := models.DeviceSwitch{}
	json.Unmarshal(body, &mist_switch)
	state, diags = resource_device_switch.SdkToTerraform(ctx, &mist_switch)
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

func (r *deviceSwitchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_switch.DeviceSwitchModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Read: device_switch_id "+state.DeviceId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.SitesDevices().GetSiteDevice(ctx, siteId, deviceId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_device_switch\" resource",
			"Unable to get the Switch, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
	mist_switch := models.DeviceSwitch{}
	json.Unmarshal(body, &mist_switch)

	state, diags = resource_device_switch.SdkToTerraform(ctx, &mist_switch)
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

func (r *deviceSwitchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_switch.DeviceSwitchModel
	tflog.Info(ctx, "Starting DeviceSwitch Update")

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

	device_switch, diags := resource_device_switch.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Update for DeviceSwitch "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_switch)

	if data.Response.StatusCode != 200 {
		api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
		if api_err != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_device_switch\" resource",
				fmt.Sprintf("Unable to update the Switch. %s", api_err),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_switch := models.DeviceSwitch{}
	json.Unmarshal(body, &mist_switch)
	state, diags = resource_device_switch.SdkToTerraform(ctx, &mist_switch)
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

func (r *deviceSwitchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_switch.DeviceSwitchModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_switch, e := resource_device_switch.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceSwitch Delete: device_switch_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_switch)
	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if data.Response.StatusCode != 200 && data.Response.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to delete the Switch. %s", api_err),
		)
		return
	}
}

func (r *deviceSwitchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_switch\" resource",
			"Import \"id\" format must be \"{site_id}.{device_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_switch\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), importIds[1])...)
}
