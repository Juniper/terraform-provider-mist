package provider

import (
	"context"
	"fmt"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_image"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &deviceImageResource{}
	_ resource.ResourceWithConfigure = &deviceImageResource{}
)

func NewDeviceImage() resource.Resource {
	return &deviceImageResource{}
}

type deviceImageResource struct {
	client mistapi.ClientInterface
}

func (r *deviceImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org WLAN Device Image client")
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
func (r *deviceImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_image"
}

func (r *deviceImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource is used to upload a Device picture.\n" +
			"This resource can be used to add a picture to a Wireless Access point, a Switch or " +
			"a Gateway. A Maximum of 3 pictures can be uploaded.",
		Attributes: resource_device_image.DeviceImageResourceSchema(ctx).Attributes,
	}
}

func (r *deviceImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org WLAN Device Image Create")
	var plan, state resource_device_image.DeviceImageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	imageNumber := int(plan.ImageNumber.ValueInt64())

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"file\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json string = ""

	data, err := r.client.SitesDevices().AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber, file, &json)

	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_device_image\" resource",
			fmt.Sprintf("Unable to create the Device Image. %s", api_err),
		)
		return
	}

	state.File = plan.File
	state.SiteId = plan.SiteId
	state.DeviceId = plan.DeviceId
	state.ImageNumber = plan.ImageNumber

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *deviceImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *deviceImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_image.DeviceImageModel
	tflog.Info(ctx, "Starting Org WLAN Device Image Update")

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	imageNumber := int(plan.ImageNumber.ValueInt64())

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"file\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json string = ""

	data, err := r.client.SitesDevices().AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber, file, &json)

	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_device_image\" resource",
			fmt.Sprintf("Unable to update the Device Image. %s", api_err),
		)
		return
	}

	state.File = plan.File
	state.SiteId = plan.SiteId
	state.DeviceId = plan.DeviceId
	state.ImageNumber = plan.ImageNumber

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *deviceImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_image.DeviceImageModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_device_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	imageNumber := int(state.ImageNumber.ValueInt64())

	httpr, err := r.client.SitesDevices().DeleteSiteDeviceImage(ctx, siteId, deviceId, imageNumber)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_device_image\" resource",
			"Could not delete Device Image, unexpected error: "+err.Error(),
		)
		return
	}
}
