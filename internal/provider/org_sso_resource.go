package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_sso"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgSsoResource{}
	_ resource.ResourceWithConfigure   = &orgSsoResource{}
	_ resource.ResourceWithImportState = &orgSsoResource{}
)

func NewOrgSso() resource.Resource {
	return &orgSsoResource{}
}

type orgSsoResource struct {
	client mistapi.ClientInterface
}

func (r *orgSsoResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist SSO client")
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
func (r *orgSsoResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_sso"
}

func (r *orgSsoResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource manages Org SSO Configuration.\n\n" +
			"Org SSO, or Single Sign-On, is a method of authentication that allows users to securely log in to multiple applications " +
			"and websites with a single set of login credentials.  \n" +
			"It involves integrating the Org portal with an Identity Provider (IdP) using the Security Assertion Markup Language (SAML) framework.  \n" +
			"This enables users to authenticate themselves through their corporate IdP, eliminating the need to remember separate " +
			"passwords or enter credentials each time they access the Org portal.",
		Attributes: resource_org_sso.OrgSsoResourceSchema(ctx).Attributes,
	}
}

func (r *orgSsoResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Sso Create")
	var plan, state resource_org_sso.OrgSsoModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	sso, diags := resource_org_sso.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsSSO().CreateOrgSso(ctx, orgId, sso)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to create the Org SSO. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_sso.SdkToTerraform(&data.Data)
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

func (r *orgSsoResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_sso.OrgSsoModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Sso Read: sso_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	ssoId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		diags.AddError(
			"Invalid \"id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsSSO().GetOrgSso(ctx, orgId, ssoId)
	if data.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		diags.AddError(
			"Error getting \"mist_org_sso\" resource",
			"Unable to get the Org SSO, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_sso.SdkToTerraform(&data.Data)
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

func (r *orgSsoResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_sso.OrgSsoModel
	tflog.Info(ctx, "Starting Sso Update")

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

	sso, diags := resource_org_sso.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	ssoId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Sso Update for Sso "+state.Id.ValueString())
	data, err := r.client.OrgsSSO().UpdateOrgSso(ctx, orgId, ssoId, sso)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to update the Org SSO. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_sso.SdkToTerraform(&data.Data)
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

func (r *orgSsoResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_sso.OrgSsoModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	ssoId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsSSO().DeleteOrgSso(ctx, orgId, ssoId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_sso\" resource",
				fmt.Sprintf("Unable to delete the Org SSO. %s", apiErr),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_sso\" resource",
			"Unable to delete the Org SSO, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *orgSsoResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso\" resource",
			"import \"id\" format must be \"{org_id}.{sso_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{sso_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_sso\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{sso_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
