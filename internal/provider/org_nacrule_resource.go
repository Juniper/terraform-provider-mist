package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"terraform-provider-mist/internal/resource_org_nacrule"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNacRuleResource{}
	_ resource.ResourceWithConfigure = &orgNacRuleResource{}
)

func NewOrgNacRule() resource.Resource {
	return &orgNacRuleResource{}
}

type orgNacRuleResource struct {
	client mistapi.ClientInterface
}

func (r *orgNacRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NacRule client")
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
func (r *orgNacRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nacrule"
}

func (r *orgNacRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_nacrule.OrgNacruleResourceSchema(ctx)
}

func (r *orgNacRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NacRule Create")
	var plan, state resource_org_nacrule.OrgNacruleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	nacrule, diags := resource_org_nacrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule org_id from plan",
			"Could not get org_nacrule org_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsNACRules().CreateOrgNacRule(ctx, orgId, &nacrule)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating NacRule",
			"Could not create NacRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_nacrule.OrgNacruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule org_id from state",
			"Could not get org_nacrule org_id, unexpected error: "+err.Error(),
		)
		return
	}
	nacruleId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule nacrule_id from state",
			"Could not get org_nacrule nacrule_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting NacRule Read: nacrule_id "+state.Id.ValueString())
	data, err := r.client.OrgsNACRules().GetOrgNacRule(ctx, orgId, nacruleId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NacRule",
			"Could not get NacRule, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_nacrule.OrgNacruleModel
	tflog.Info(ctx, "Starting NacRule Update")

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

	nacrule, diags := resource_org_nacrule.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule org_id from state",
			"Could not get org_nacrule org_id, unexpected error: "+err.Error(),
		)
		return
	}
	nacruleId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule nacrule_id from state",
			"Could not get org_nacrule nacrule_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting NacRule Update for NacRule "+state.Id.ValueString())
	data, err := r.client.OrgsNACRules().
		UpdateOrgNacRule(ctx, orgId, nacruleId, &nacrule)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating NacRule",
			"Could not update NacRule, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nacrule.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_nacrule.OrgNacruleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule org_id from state",
			"Could not get org_nacrule org_id, unexpected error: "+err.Error(),
		)
		return
	}
	nacruleId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org_nacrule nacrule_id from state",
			"Could not get org_nacrule nacrule_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting NacRule Delete: nacrule_id "+state.Id.ValueString())
	_, err = r.client.OrgsNACRules().DeleteOrgNacRule(ctx, orgId, nacruleId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting NacRule",
			"Could not delete NacRule, unexpected error: "+err.Error(),
		)
		return
	}
}
