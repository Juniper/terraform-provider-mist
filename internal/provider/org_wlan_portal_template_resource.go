package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlan_portal_template"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWlanPortalTemplateResource{}
	_ resource.ResourceWithConfigure = &orgWlanPortalTemplateResource{}
)

func NewOrgWlanPortalTemplate() resource.Resource {
	return &orgWlanPortalTemplateResource{}
}

type orgWlanPortalTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgWlanPortalTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org WLAN Portal Template client")
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
func (r *orgWlanPortalTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wlan_portal_template"
}

func (r *orgWlanPortalTemplateResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource is used customize the WLAN Guest Portal.\n\n" +
			"The WLAN Portal Template can be used to define:\n" +
			"* Guest Authentication methods and parameters (access duration, ...)\n" +
			"* Default values of the text fields and labels on the portal\n" +
			"* Values of the text fields and labels based on the User Agent (`locales` property)\n\n" +
			"**Notes:**\n" +
			"* There is no feedback from the API, so there is no possibility to validate the changes. " +
			"The resource states is directly generated based on the resource plan." +
			"* There is no option to delete or revert the changes. Deleting the resource will just remove it from the states. " +
			"Once removed, it is possible to create a new one. It will replace the previous template",
		Attributes: resource_org_wlan_portal_template.OrgWlanPortalTemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgWlanPortalTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org WLAN Portal Template Create")
	var plan, state resource_org_wlan_portal_template.OrgWlanPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_wlan_portal_template.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsWlans().UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_wlan_portal_template\" resource",
			"Unable to create the WLAN Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.OrgId = plan.OrgId
	state.WlanId = plan.WlanId
	state.PortalTemplate = plan.PortalTemplate

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWlanPortalTemplateResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {

}

func (r *orgWlanPortalTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "Starting Org WLAN Portal Template Update")
	var state, plan resource_org_wlan_portal_template.OrgWlanPortalTemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_wlan_portal_template.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsWlans().UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, &template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_wlan_portal_template\" resource",
			"Unable to create the WLAN Portal Template, unexpected error: "+err.Error(),
		)
		return
	}

	state.OrgId = plan.OrgId
	state.WlanId = plan.WlanId
	state.PortalTemplate = plan.PortalTemplate

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgWlanPortalTemplateResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wlan_portal_template.OrgWlanPortalTemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(state.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"mist_org_wlan_portal_template\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	template, diags := resource_org_wlan_portal_template.DeleteTerraformToSdk()
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err = r.client.OrgsWlans().UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, template)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_wlan_portal_template\" resource",
			"Unable to delete the WLAN Portal Template, unexpected error: "+err.Error(),
		)
		return
	}
}
