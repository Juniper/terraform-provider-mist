package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_apitoken"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgApitokenResource{}
	_ resource.ResourceWithConfigure = &orgApitokenResource{}
)

func NewOrgApitoken() resource.Resource {
	return &orgApitokenResource{}
}

type orgApitokenResource struct {
	client mistapi.ClientInterface
}

func (r *orgApitokenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Apitoken client")
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
func (r *orgApitokenResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_apitoken"
}

func (r *orgApitokenResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource manages Org API Tokens.\n\n" +
			"An Org API token is a unique identifier used by an application to authenticate and access the Mist APIs. " +
			"These tokens are used to authenticate requests made to the API server and ensure secure access to the API. " +
			"They are not bound to any specific user and provide access to the organization as a whole. \n" +
			"Organization tokens support different privileges and can only be used for the specific organization they are generated for.\n" +
			"Rate limiting is done on an individual token basis, so if one token reaches its rate limit, it does not impact other tokens.",
		Attributes: resource_org_apitoken.OrgApitokenResourceSchema(ctx).Attributes,
	}
}

func (r *orgApitokenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Apitoken Create")
	var plan, state resource_org_apitoken.OrgApitokenModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	apitoken, diags := resource_org_apitoken.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsAPITokens().CreateOrgApiToken(ctx, orgId, apitoken)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to create the API Token. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_apitoken.SdkToTerraform(ctx, data.Data, nil)

	state = *r.readApiToken(ctx, &diags, state.OrgId, state.Id, state)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgApitokenResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_apitoken.OrgApitokenModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state = *r.readApiToken(ctx, &diags, state.OrgId, state.Id, state)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgApitokenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_apitoken.OrgApitokenModel
	tflog.Info(ctx, "Starting Apitoken Update")

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

	if !state.SrcIps.Equal(plan.SrcIps) {
		resp.Diagnostics.AddError(
			"Invalid \"src_ips\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("The value of \"src_ips\" cannot be changed once the API Token is create. Please create a new API Token. Planned value: %s, State value %s", plan.SrcIps.String(), state.SrcIps.String()),
		)
		return
	}

	apitoken, diags := resource_org_apitoken.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	apitokenId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Apitoken Update for Apitoken "+state.Id.ValueString())
	data, err := r.client.OrgsAPITokens().UpdateOrgApiToken(ctx, orgId, apitokenId, apitoken)

	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to update the API Token. %s", apiErr),
		)
		return
	}

	state = *r.readApiToken(ctx, &diags, state.OrgId, state.Id, state)
	if resp.Diagnostics.HasError() {
		return
	}

	// state, diags = resource_org_apitoken.SdkToTerraform(ctx, data.Data, &state)
	// resp.Diagnostics.Append(diags...)
	// if resp.Diagnostics.HasError() {
	// 	return
	// }

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgApitokenResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_apitoken.OrgApitokenModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	apitokenId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsAPITokens().DeleteOrgApiToken(ctx, orgId, apitokenId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_apitoken\" resource",
				fmt.Sprintf("Unable to delete the API Token. %s", apiErr),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_apitoken\" resource",
			"Unable to delete the API Token, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *orgApitokenResource) readApiToken(ctx context.Context, diags *diag.Diagnostics, mistOrgId basetypes.StringValue, mistApitokenId basetypes.StringValue, state resource_org_apitoken.OrgApitokenModel) *resource_org_apitoken.OrgApitokenModel {
	tflog.Info(ctx, "Starting Apitoken Read: apitoken_id "+mistApitokenId.ValueString())
	orgId, err := uuid.Parse(mistOrgId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", mistOrgId.ValueString(), err.Error()),
		)
		return nil
	}

	apitokenId, err := uuid.Parse(mistApitokenId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", mistApitokenId.ValueString(), err.Error()),
		)
		return nil
	}

	httpr, err := r.client.OrgsAPITokens().GetOrgApiToken(ctx, orgId, apitokenId)
	if httpr.Response.StatusCode == 404 {
		return nil
	} else if err != nil {
		diags.AddError(
			"Error getting \"mist_org_apitoken\" resource",
			"Unable to get the API Token, unexpected error: "+err.Error(),
		)
		return nil
	}

	state, e := resource_org_apitoken.SdkToTerraform(ctx, httpr.Data, &state)
	diags.Append(e...)
	return &state
}
