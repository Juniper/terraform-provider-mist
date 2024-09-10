package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wlan_portal_template"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWlanPortalTemplateResource{}
	_ resource.ResourceWithConfigure = &siteWlanPortalTemplateResource{}
)

func NewSiteWlanPortalTemplate() resource.Resource {
	return &siteWlanPortalTemplateResource{}
}

type siteWlanPortalTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *siteWlanPortalTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Site WLAN Portal Template client")
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
func (r *siteWlanPortalTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wlan_portal_template"
}

func (r *siteWlanPortalTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource is used customize the WLAN Guest Portal." +
			"The WLAN Portal Template can be used to define:\n" +
			"* Guest Authentication methods and parameters (access duration, ...)\n" +
			"* Default values of the text fields and labels on the portal\n" +
			"* Values of the text fields and labels based on the User Agent (`locales` property)\n\n" +
			"**Notes:**\n" +
			"* There is no feedback from the API, so there is no possibility to validate the changes. " +
			"The resource states is directly generated based on the resource plan." +
			"* There is no option to delete or revert the changes. Deleting the resource will just remove it from the states. " +
			"Once removed, it is possible to create a new one. It will replace the previous template",
		Attributes: resource_site_wlan_portal_template.SiteWlanPortalTemplateResourceSchema(ctx).Attributes,
	}
}

func (r *siteWlanPortalTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Site WLAN Portal Template Create")
	var plan, state resource_site_wlan_portal_template.SiteWlanPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_site_wlan_portal_template.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.SitesWlans().UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site_wlan_portal_template\" resource",
			"Unable to create the WLAN Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.SiteId = plan.SiteId
	state.WlanId = plan.WlanId
	state.PortalTemplate = plan.PortalTemplate

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *siteWlanPortalTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *siteWlanPortalTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "Starting Site WLAN Portal Template Update")
	var state, plan resource_site_wlan_portal_template.SiteWlanPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_site_wlan_portal_template.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.SitesWlans().UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating \"mist_site_wlan_portal_template\" resource",
			"Unable to create the WLAN Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.SiteId = plan.SiteId
	state.WlanId = plan.WlanId
	state.PortalTemplate = plan.PortalTemplate

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteWlanPortalTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wlan_portal_template.SiteWlanPortalTemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(state.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_site_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_site_wlan_portal_template.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpr, err := r.client.SitesWlans().UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, template)
	if httpr.Response.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_site_wlan_portal_template\" resource",
			"Unable to delete the WLAN Portal Temaplate, unexpected error: "+err.Error(),
		)
		return
	}
}
