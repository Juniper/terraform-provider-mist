package provider

import (
	"context"
	"fmt"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wlan_portal_image"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWlanPortalImageResource{}
	_ resource.ResourceWithConfigure = &siteWlanPortalImageResource{}
)

func NewSiteWlanPortalImage() resource.Resource {
	return &siteWlanPortalImageResource{}
}

type siteWlanPortalImageResource struct {
	client mistapi.ClientInterface
}

func (r *siteWlanPortalImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org WLAN Portal Image client")
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
func (r *siteWlanPortalImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wlan_portal_image"
}

func (r *siteWlanPortalImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource is used to upload a WLAN Captive Web Portal background image.",
		Attributes:          resource_site_wlan_portal_image.SiteWlanPortalImageResourceSchema(ctx).Attributes,
	}
}

func (r *siteWlanPortalImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org WLAN Portal Image Create")
	var plan, state resource_site_wlan_portal_image.SiteWlanPortalImageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"file\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json string = ""

	data, err := r.client.SitesWlans().UploadSiteWlanPortalImage(ctx, siteId, wlanId, file, &json)

	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Unable to create the Portal Image. %s", api_err),
		)
		return
	}

	state.File = plan.File
	state.SiteId = plan.SiteId
	state.WlanId = plan.WlanId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *siteWlanPortalImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *siteWlanPortalImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_wlan_portal_image.SiteWlanPortalImageModel
	tflog.Info(ctx, "Starting Org WLAN Portal Image Update")

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	file, err := models.GetFile(plan.File.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"file\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", plan.File.ValueString(), err.Error()),
		)
		return
	}
	var json string = ""

	data, err := r.client.SitesWlans().UploadSiteWlanPortalImage(ctx, siteId, wlanId, file, &json)

	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Unable to update the Portal Image. %s", api_err),
		)
		return
	}

	state.File = plan.File
	state.SiteId = plan.SiteId
	state.WlanId = plan.WlanId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *siteWlanPortalImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wlan_portal_image.SiteWlanPortalImageModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(state.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesWlans().DeleteSiteWlanPortalImage(ctx, siteId, wlanId)
	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_site_wlan_portal_image\" resource",
			fmt.Sprintf("Unable to delete the Portal Image. %s", api_err),
		)
		return
	}
}
