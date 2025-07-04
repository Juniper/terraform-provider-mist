package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_servicepolicy"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgServicepolicyResource{}
	_ resource.ResourceWithConfigure   = &orgServicepolicyResource{}
	_ resource.ResourceWithImportState = &orgServicepolicyResource{}
)

func NewOrgServicepolicyResource() resource.Resource {
	return &orgServicepolicyResource{}
}

type orgServicepolicyResource struct {
	client mistapi.ClientInterface
}

func (r *orgServicepolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *orgServicepolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_servicepolicy"
}
func (r *orgServicepolicyResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages WAN Assurance Service Policies (Application Policies).\n\n" +
			"The Service Policies can be used in the `service_policies` object by referencing the Service Policy ID as the `servicepolicy_id` in:\n" +
			"* the Gateway configuration (`mist_device_gateway.service_policies`)\n" +
			"* the Gateway Templates (`mist_org_gatewaytemplate.service_policies`)\n" +
			"* the HUB Profiles (`mist_org_deviceprofile_gateway.service_policies`)\n" +
			"They can be used to manage common policies between multiples configurations",
		Attributes: resource_org_servicepolicy.OrgServicepolicyResourceSchema(ctx).Attributes,
	}
}

func (r *orgServicepolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgServicepolicy Create")
	var plan, state resource_org_servicepolicy.OrgServicepolicyModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	servicepolicy, diags := resource_org_servicepolicy.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting OrgServicepolicy Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsServicePolicies().CreateOrgServicePolicy(ctx, orgId, &servicepolicy)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to create the Service Policy. %s", apiErr),
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

func (r *orgServicepolicyResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_servicepolicy.OrgServicepolicyModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Read: servicepolicy_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	servicepolicyId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.OrgsServicePolicies().GetOrgServicePolicy(ctx, orgId, servicepolicyId)
	if httpr.Response != nil && httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_servicepolicy\" resource",
			"Unable to get the Service Policy, unexpected error: "+err.Error(),
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

func (r *orgServicepolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

	servicepolicy, diags := resource_org_servicepolicy.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Update for OrgServicepolicy "+plan.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	servicepolicyId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsServicePolicies().UpdateOrgServicePolicy(ctx, orgId, servicepolicyId, &servicepolicy)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to update the Service Policy. %s", apiErr),
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

func (r *orgServicepolicyResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_servicepolicy.OrgServicepolicyModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgServicepolicy Delete: servicepolicy_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	servicepolicyId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsServicePolicies().DeleteOrgServicePolicy(ctx, orgId, servicepolicyId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_servicepolicy\" resource",
				fmt.Sprintf("Unable to delete the Service Policy. %s", apiErr),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_servicepolicy\" resource",
			"Unable to delete the Service Policy, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *orgServicepolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_servicepolicy\" resource",
			"import \"id\" format must be \"{org_id}.{servicepolicy_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{servicepolicy_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_servicepolicy\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{servicepolicy_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
