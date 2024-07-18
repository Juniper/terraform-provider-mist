package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_site_wxrule"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &siteWxRuleResource{}
	_ resource.ResourceWithConfigure = &siteWxRuleResource{}
)

func NewSiteWxRule() resource.Resource {
	return &siteWxRuleResource{}
}

type siteWxRuleResource struct {
	client mistapi.ClientInterface
}

func (r *siteWxRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WxRule client")
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
func (r *siteWxRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_wxrule"
}

func (r *siteWxRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource manages the Site WxRules (WLAN policies)." +
			"A WxLAN policy is a set of rules and settings that can be applied to devices in a network to determine " +
			"how they are treated. it provides support for access policies, network segmentation, role-based policies, " +
			"micro-segmentation, and least privilege. " +
			"WxLAN policies are used to allow or deny specific users from accessing specific resources in a wireless network.",
		Attributes: resource_site_wxrule.SiteWxruleResourceSchema(ctx).Attributes,
	}
}

func (r *siteWxRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxRule Create")
	var plan, state resource_site_wxrule.SiteWxruleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(plan.SiteId.ValueString())
	wxrule, diags := resource_site_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.SitesWxRules().CreateSiteWxRule(ctx, siteId, wxrule)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating WxRule",
			"Could not create WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_site_wxrule.SiteWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting WxRule Read: wxrule_id "+state.Id.ValueString())
	data, err := r.client.SitesWxRules().GetSiteWxRule(ctx, siteId, wxruleId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WxRule",
			"Could not get WxRule, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_site_wxrule.SiteWxruleModel
	tflog.Info(ctx, "Starting WxRule Update")

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

	wxrule, diags := resource_site_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting WxRule Update for WxRule "+state.Id.ValueString())
	data, err := r.client.SitesWxRules().UpdateSiteWxRule(ctx, siteId, wxruleId, wxrule)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WxRule",
			"Could not update WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_site_wxrule.SdkToTerraform(ctx, data.Data)
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

func (r *siteWxRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_site_wxrule.SiteWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId := uuid.MustParse(state.SiteId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting WxRule Delete: wxrule_id "+state.Id.ValueString())
	_, err := r.client.SitesWxRules().DeleteSiteWxRule(ctx, siteId, wxruleId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WxRule",
			"Could not delete WxRule, unexpected error: "+err.Error(),
		)
		return
	}
}
