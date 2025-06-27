package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_ap"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceApResource{}
	_ resource.ResourceWithConfigure   = &deviceApResource{}
	_ resource.ResourceWithImportState = &deviceApResource{}
)

func NewDeviceApResource() resource.Resource {
	return &deviceApResource{}
}

type deviceApResource struct {
	client mistapi.ClientInterface
}

func (r *deviceApResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceAp Resource client")
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
func (r *deviceApResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_ap"
}

func (r *deviceApResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Wireless Access Point configuration.\n\n" +
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

	deviceAp, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_ap\" resource",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_ap\" resource",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &deviceAp)

	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_device_ap\" resource",
				fmt.Sprintf("Unable to create the Wireless Access Point. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistAp := models.DeviceAp{}
	json.Unmarshal(body, &mistAp)
	state, diags = resource_device_ap.SdkToTerraform(ctx, &mistAp)
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

func (r *deviceApResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
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
			"Invalid \"site_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_ap\" resource",
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
			"Error getting \"mist_device_ap\" resource",
			"Unable to get the Wireless Access Point, unexpected error: "+err.Error(),
		)
		return
	}
	body, _ := io.ReadAll(httpr.Response.Body)
	mistAp := models.DeviceAp{}
	json.Unmarshal(body, &mistAp)

	state, diags = resource_device_ap.SdkToTerraform(ctx, &mistAp)
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

	deviceAp, diags := resource_device_ap.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Update for DeviceAp "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &deviceAp)

	if data.Response.StatusCode != 200 {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error updating \"mist_device_ap\" resource",
				fmt.Sprintf("Unable to update the Wireless Access Point. %s", apiErr),
			)
			return
		}
	}

	body, _ := io.ReadAll(data.Response.Body)
	mistAp := models.DeviceAp{}
	json.Unmarshal(body, &mistAp)
	state, diags = resource_device_ap.SdkToTerraform(ctx, &mistAp)
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

func (r *deviceApResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_ap.DeviceApModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deviceAp, e := resource_device_ap.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceAp Delete: device_ap_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &deviceAp)
	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if data.Response.StatusCode != 200 && data.Response.StatusCode != 404 && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to delete the Wireless Access Point. %s", apiErr),
		)
		return
	}
}

func (r *deviceApResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Import \"id\" format must be \"{site_id}.{device_id}\", got %s", req.ID),
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_ap\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("device_id"), importIds[1])...)
}
