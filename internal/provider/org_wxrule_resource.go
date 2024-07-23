package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wxrule"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWxRuleResource{}
	_ resource.ResourceWithConfigure = &orgWxRuleResource{}
)

func NewOrgWxRule() resource.Resource {
	return &orgWxRuleResource{}
}

type orgWxRuleResource struct {
	client mistapi.ClientInterface
}

func (r *orgWxRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *orgWxRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wxrule"
}

func (r *orgWxRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource manages the Org WxRules (WxLAN policies)." +
			"A WxLAN policy is a set of rules and settings that can be applied to devices in a network to determine " +
			"how they are treated. it provides support for access policies, network segmentation, role-based policies, " +
			"micro-segmentation, and least privilege. " +
			"WxLAN policies are used to allow or deny specific users from accessing specific resources in a wireless network.",
		Attributes: resource_org_wxrule.OrgWxruleResourceSchema(ctx).Attributes,
	}
}

func (r *orgWxRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxRule Create")
	var plan, state resource_org_wxrule.OrgWxruleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wxrule, diags := resource_org_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	data, err := r.client.OrgsWxRules().CreateOrgWxRule(ctx, orgId, wxrule)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating WxRule",
			"Could not create WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wxrule.SdkToTerraform(ctx, &data.Data)
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

func (r *orgWxRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_wxrule.OrgWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxRule Read: wxrule_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsWxRules().GetOrgWxRule(ctx, orgId, wxruleId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WxRule",
			"Could not get WxRule, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_wxrule.SdkToTerraform(ctx, &data.Data)
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

func (r *orgWxRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_wxrule.OrgWxruleModel
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

	wxrule, diags := resource_org_wxrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting WxRule Update for WxRule "+state.Id.ValueString())
	data, err := r.client.OrgsWxRules().UpdateOrgWxRule(ctx, orgId, wxruleId, wxrule)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WxRule",
			"Could not update WxRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wxrule.SdkToTerraform(ctx, &data.Data)
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

func (r *orgWxRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wxrule.OrgWxruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxruleId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting WxRule Delete: wxrule_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsWxRules().DeleteOrgWxRule(ctx, orgId, wxruleId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WxRule",
			"Could not delete WxRule, unexpected error: "+err.Error(),
		)
		return
	}
}
