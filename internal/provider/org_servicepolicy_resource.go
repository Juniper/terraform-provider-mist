package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_servicepolicy"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgOrgServicepolicyResource{}
	_ resource.ResourceWithConfigure   = &orgOrgServicepolicyResource{}
	_ resource.ResourceWithImportState = &orgOrgServicepolicyResource{}
)

func NewOrgServicepolicyResource() resource.Resource {
	return &orgOrgServicepolicyResource{}
}

type orgOrgServicepolicyResource struct {
	client mistapi.ClientInterface
}

func (r *orgOrgServicepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgServicepolicy client")
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

func (r *orgOrgServicepolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_servicepolicy"
}
func (r *orgOrgServicepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages WAN Assurance Service Policies (Application Policiess)." +
			"The Service Policies are used in the `servicepolicy_policies` from the Gateway configuration and Gateway templates." +
			"They can be used to manage common policies betweeen multiples configurations",
		Attributes: resource_org_servicepolicy.OrgServicepolicyResourceSchema(ctx).Attributes,
	}
}

func (r *orgOrgServicepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgServicepolicy Create")
	var plan, state resource_org_servicepolicy.OrgServicepolicyModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	servicepolicy, diags := resource_org_servicepolicy.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId := uuid.MustParse(plan.OrgId.ValueString())
	tflog.Info(ctx, "Starting OrgServicepolicy Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsServicePolicies().CreateOrgServicePolicy(ctx, orgId, &servicepolicy)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating servicepolicy",
			"Could not create servicepolicy, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_servicepolicy.SdkToTerraform(ctx, &data.Data)
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

func (r *orgOrgServicepolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_servicepolicy.OrgServicepolicyModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Read: servicepolicy_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	servicepolicyId := uuid.MustParse(state.Id.ValueString())
	httpr, err := r.client.OrgsServicePolicies().GetOrgServicePolicy(ctx, orgId, servicepolicyId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting servicepolicy",
			"Could not get servicepolicy, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_servicepolicy.SdkToTerraform(ctx, &httpr.Data)
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

func (r *orgOrgServicepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_servicepolicy.OrgServicepolicyModel
	tflog.Info(ctx, "Starting OrgServicepolicy Update")

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

	servicepolicy, diags := resource_org_servicepolicy.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Update for OrgServicepolicy "+plan.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	servicepolicyId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsServicePolicies().UpdateOrgServicePolicy(ctx, orgId, servicepolicyId, &servicepolicy)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating servicepolicy",
			"Could not update servicepolicy, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_servicepolicy.SdkToTerraform(ctx, &data.Data)
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

func (r *orgOrgServicepolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_servicepolicy.OrgServicepolicyModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Delete: servicepolicy_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	servicepolicyId := uuid.MustParse(state.Id.ValueString())
	httpr, err := r.client.OrgsServicePolicies().DeleteOrgServicePolicy(ctx, orgId, servicepolicyId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting servicepolicy",
			"Could not delete servicepolicy, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgOrgServicepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org id from import",
			"Could not get org id, unexpected error: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
