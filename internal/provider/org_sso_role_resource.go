package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sso_role"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgSsoRoleResource{}
	_ resource.ResourceWithConfigure = &orgSsoRoleResource{}
	// _ resource.ResourceWithImportState = &orgSsoRoleResource{}
)

func NewOrgSsoRole() resource.Resource {
	return &orgSsoRoleResource{}
}

type orgSsoRoleResource struct {
	client mistapi.ClientInterface
}

func (r *orgSsoRoleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SSO Role client")
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
func (r *orgSsoRoleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sso_role"
}

func (r *orgSsoRoleResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource manages Org SSO Roles for Admin Authentication.\n\n" +
			"SSO roles refer to the different functions assigned to users within a Single Sign-On (SSO) system.  \n" +
			"These roles determine the tasks and actions that users can perform within the SSO system. " +
			"There are typically predefined roles and custom roles in an SSO system.  \n" +
			"Roles in SSO provide a well-defined separation of responsibility and visibility, allowing for granular-level access control on SSO objects.",
		Attributes: resource_org_sso_role.OrgSsoRoleResourceSchema(ctx).Attributes,
	}
}

func (r *orgSsoRoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting SsoRole Create")
	var plan, state resource_org_sso_role.OrgSsoRoleModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	ssoRole, diags := resource_org_sso_role.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsSSORoles().CreateOrgSsoRole(ctx, orgId, ssoRole)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to create the Org SSO Role. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_sso_role.SdkToTerraform(ctx, data.Data)
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

func (r *orgSsoRoleResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_sso_role.OrgSsoRoleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting SsoRole Read: sso_role_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	ssoRoleId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsSSORoles().GetOrgSsoRole(ctx, orgId, ssoRoleId)
	if data.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		diags.AddError(
			"Error getting \"mist_org_sso_role\" resource",
			"Unable to get the Org SSO Role, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sso_role.SdkToTerraform(ctx, data.Data)
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

func (r *orgSsoRoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_sso_role.OrgSsoRoleModel
	tflog.Info(ctx, "Starting SsoRole Update")

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

	ssoRole, diags := resource_org_sso_role.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	ssoRoleid, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting SsoRole Update for SsoRole "+state.Id.ValueString())
	data, err := r.client.OrgsSSORoles().UpdateOrgSsoRole(ctx, orgId, ssoRoleid, ssoRole)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to update the Org SSO Role. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_sso_role.SdkToTerraform(ctx, data.Data)
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

func (r *orgSsoRoleResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_sso_role.OrgSsoRoleModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	ssoRoleid, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting SsoRole Delete: sso_role_id "+state.Id.ValueString())
	data, err := r.client.OrgsSSORoles().DeleteOrgSsoRole(ctx, orgId, ssoRoleid)
	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_sso_role\" resource",
			fmt.Sprintf("Unable to delete the Org SSO Role. %s", apiErr),
		)
		return
	}
}

// func (r *orgSsoRoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

// 	importIds := strings.Split(req.ID, ".")
// 	if len(importIds) != 2 {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"id\" value for \"mist_org_sso_role\" resource",
// 			"import \"id\" format must be \"{org_id}.{sso_role_id}\"",
// 		)
// 		return
// 	}
// 	_, err := uuid.Parse(importIds[0])
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"org_id\" value for \"mist_org_sso_role\" resource",
// 			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{sso_role_id}\"", importIds[0], err.Error()),
// 		)
// 		return
// 	}
// 	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

// 	_, err = uuid.Parse(importIds[1])
// 	if err != nil {
// 		resp.Diagnostics.AddError(
// 			"Invalid \"id\" value for \"mist_org_sso_role\" resource",
// 			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{sso_role_id}\"", importIds[1], err.Error()),
// 		)
// 		return
// 	}
// 	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
// }
