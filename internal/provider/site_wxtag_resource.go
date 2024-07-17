package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wxtag"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWxTagResource{}
	_ resource.ResourceWithConfigure = &siteWxTagResource{}
)

func NewSiteWxTag() resource.Resource {
	return &siteWxTagResource{}
}

type siteWxTagResource struct {
	client mistapi.ClientInterface
}

func (r *siteWxTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WxTag client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *models.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *siteWxTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wxtag"
}

func (r *siteWxTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource manages the Site Wlan tags (labels)." +
			"The tags can be used " +
			"  * within the WxRules to create filtering rules, or assign specific VLAN" +
			"  * in the WLANs configuration to assign a WLAN to specific APs" +
			"  * to identify unknown application used by Wi-Fi clients",
		Attributes: resource_site_wxtag.SiteWxtagResourceSchema(ctx).Attributes,
	}
}

func (r *siteWxTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxTag Create")
	var plan, state resource_site_wxtag.SiteWxtagModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(plan.SiteId.ValueString())
	wxtag, diags := resource_site_wxtag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.SitesWxTags().CreateSiteWxTag(ctx, siteId, wxtag)
	//.SitesWxTagsAPI().CreateSiteWxTag(ctx, plan.SiteId.ValueString()).WxlanTag(*wxtag)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error updating WxTag",
			"Could not update WxTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_wxtag.SiteWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())

	tflog.Info(ctx, "Starting WxTag Read: wxtag_id "+state.Id.ValueString())
	data, err := r.client.SitesWxTags().GetSiteWxTag(ctx, siteId, wxtagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WxTag",
			"Could not get WxTag, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_wxtag.SiteWxtagModel
	tflog.Info(ctx, "Starting WxTag Update")

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

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())
	wxtag, diags := resource_site_wxtag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Update for WxTag "+state.Id.ValueString())
	data, err := r.client.SitesWxTags().UpdateSiteWxTag(ctx, siteId, wxtagId, wxtag)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WxTag",
			"Could not update WxTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wxtag.SiteWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())

	tflog.Info(ctx, "Starting WxTag Delete: wxtag_id "+state.Id.ValueString())
	_, err := r.client.SitesWxTags().DeleteSiteWxTag(ctx, siteId, wxtagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WxTag",
			"Could not delete WxTag, unexpected error: "+err.Error(),
		)
		return
	}
}
