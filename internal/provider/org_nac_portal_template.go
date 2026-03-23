package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nac_portal_template"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNacPortalTemplateResource{}
	_ resource.ResourceWithConfigure = &orgNacPortalTemplateResource{}
)

func NewOrgNacPortalTemplate() resource.Resource {
	return &orgNacPortalTemplateResource{}
}

type orgNacPortalTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgNacPortalTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org NAC Portal Template client")
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

func (r *orgNacPortalTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nac_portal_template"
}

func (r *orgNacPortalTemplateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryNac + "This resource is used customize the NAC Portal.\n\n" +
			"The NAC Portal Template can be used to define:\n" +
			"* The portal alignment, color scheme, and logo\n" +
			"* Whether to display \"Powered by Juniper Mist\" footer\n\n" +
			"**Notes:**\n" +
			"* There is no feedback from the API, so there is no possibility to validate the changes. " +
			"The resource states is directly generated based on the resource plan." +
			"* There is no option to delete or revert the changes. Deleting the resource will just remove it from the states. " +
			"Once removed, it is possible to create a new one. It will replace the previous template",
		Attributes: resource_org_nac_portal_template.OrgNacPortalTemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgNacPortalTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org NAC Portal Template Create")
	var plan, state resource_org_nac_portal_template.OrgNacPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(plan.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_nac_portal_template.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsNACPortals().UpdateOrgNacPortalTemplate(ctx, orgId, nacPortalId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_nac_portal_template\" resource",
			"Unable to create the NAC Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.OrgId = plan.OrgId
	state.NacportalId = plan.NacportalId
	state.Alignment = plan.Alignment
	state.Color = plan.Color
	state.Logo = plan.Logo
	state.PoweredBy = plan.PoweredBy

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgNacPortalTemplateResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {

}

func (r *orgNacPortalTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "Starting Org NAC Portal Template Update")
	var state, plan resource_org_nac_portal_template.OrgNacPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(plan.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_nac_portal_template.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsNACPortals().UpdateOrgNacPortalTemplate(ctx, orgId, nacPortalId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_nac_portal_template\" resource",
			"Unable to update the NAC Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.OrgId = plan.OrgId
	state.NacportalId = plan.NacportalId
	state.Alignment = plan.Alignment
	state.Color = plan.Color
	state.Logo = plan.Logo
	state.PoweredBy = plan.PoweredBy

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgNacPortalTemplateResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_nac_portal_template.OrgNacPortalTemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	nacPortalId, err := uuid.Parse(state.NacportalId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"nacportal_id\" value for \"mist_org_nac_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.NacportalId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_nac_portal_template.DeleteTerraformToSdk()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsNACPortals().UpdateOrgNacPortalTemplate(ctx, orgId, nacPortalId, template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_nac_portal_template\" resource",
			"Unable to delete the NAC Portal Template, unexpected error: "+err.Error(),
		)
		return
	}
}
