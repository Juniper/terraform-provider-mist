package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
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
func (r *orgApitokenResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_apitoken"
}

func (r *orgApitokenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This resource manages Org API Tokens.",
		Attributes:          resource_org_apitoken.OrgApitokenResourceSchema(ctx).Attributes,
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

	apitoken, diags := resource_org_apitoken.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsAPITokens().CreateOrgApiToken(ctx, orgId, apitoken)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to create the API Token. %s", api_err),
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

func (r *orgApitokenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	apitoken, diags := resource_org_apitoken.TerraformToSdk(ctx, &plan)
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

	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to update the API Token. %s", api_err),
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

func (r *orgApitokenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

	tflog.Info(ctx, "Starting Apitoken Delete: apitoken_id "+state.Id.ValueString())
	data, err := r.client.OrgsAPITokens().DeleteOrgApiToken(ctx, orgId, apitokenId)
	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to delete the API Token. %s", api_err),
		)
		return
	}
}

func (r *orgApitokenResource) readApiToken(ctx context.Context, diags *diag.Diagnostics, org_id basetypes.StringValue, apitoken_id basetypes.StringValue, state resource_org_apitoken.OrgApitokenModel) *resource_org_apitoken.OrgApitokenModel {
	tflog.Info(ctx, "Starting Apitoken Read: apitoken_id "+apitoken_id.ValueString())
	orgId, err := uuid.Parse(org_id.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", org_id.ValueString(), err.Error()),
		)
		return nil
	}

	apitokenId, err := uuid.Parse(apitoken_id.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", apitoken_id.ValueString(), err.Error()),
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

// func (r *orgApitokenResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

// 	importIds := strings.Split(req.ID, ".")
// 	if len(importIds) != 2 {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
// 			"import \"id\" format must be \"{org_id}.{apitoken_id}\"",
// 		)
// 		return
// 	}
// 	_, err := uuid.Parse(importIds[0])
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"org_id\" value for \"mist_org_apitoken\" resource",
// 			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{apitoken_id}\"", importIds[0], err.Error()),
// 		)
// 		return
// 	}
// 	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

// 	_, err = uuid.Parse(importIds[1])
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"id\" value for \"mist_org_apitoken\" resource",
// 			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{apitoken_id}\"", importIds[1], err.Error()),
// 		)
// 		return
// 	}
// 	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
// }
