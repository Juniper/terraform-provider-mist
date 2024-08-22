package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_site_setting"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &siteSettingResource{}
	_ resource.ResourceWithConfigure   = &siteSettingResource{}
	_ resource.ResourceWithImportState = &siteSettingResource{}
)

func NewSiteSettingResource() resource.Resource {
	return &siteSettingResource{}
}

type siteSettingResource struct {
	client mistapi.ClientInterface
}

func (r *siteSettingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SiteSetting client")
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
func (r *siteSettingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_setting"
}

func (r *siteSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This resource manages the Site Settings." +
			"The Site Settings can used to customize the Site configuration and assign Site Variables (Sites Variables can be reused in configuration templates) ",
		Attributes: resource_site_setting.SiteSettingResourceSchema(ctx).Attributes,
	}
}

func (r *siteSettingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting SiteSetting Create")
	var plan, state resource_site_setting.SiteSettingModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteSetting, diags := resource_site_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Starting SiteSetting Create for Site "+plan.SiteId.ValueString())
	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_setting\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.SitesSetting().UpdateSiteSettings(ctx, siteId, siteSetting)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating siteSetting",
			"Could not create siteSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_setting.SdkToTerraform(ctx, &data.Data)
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

func (r *siteSettingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_setting.SiteSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting SiteSetting Read: siteSetting_id "+state.SiteId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_setting\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.SitesSetting().GetSiteSetting(ctx, siteId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting siteSetting",
			"Could not get siteSetting, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_setting.SdkToTerraform(ctx, &httpr.Data)
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

func (r *siteSettingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_setting.SiteSettingModel
	tflog.Info(ctx, "Starting SiteSetting Update")

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

	siteSetting, diags := resource_site_setting.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting SiteSetting Update for SiteSetting "+state.SiteId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_setting\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.SitesSetting().UpdateSiteSettings(ctx, siteId, siteSetting)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating siteSetting",
			"Could not update siteSetting, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_setting.SdkToTerraform(ctx, &data.Data)
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

func (r *siteSettingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_setting.SiteSettingModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting SiteSetting Delete: siteSetting_id "+state.SiteId.ValueString())
	siteSetting, e := resource_site_setting.DeleteTerraformToSdk(ctx)
	diags.Append(e...)
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"site_setting\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.SitesSetting().UpdateSiteSettings(ctx, siteId, siteSetting)
	if httpr.Response.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting siteSetting",
			"Could not delete siteSetting, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *siteSettingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"site_setting\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" must be a valid Site Id.", req.ID, err.Error()),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
