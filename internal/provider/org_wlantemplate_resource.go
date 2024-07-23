package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlantemplate"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWlanTemplateResource{}
	_ resource.ResourceWithConfigure = &orgWlanTemplateResource{}
)

func NewOrgWlanTemplate() resource.Resource {
	return &orgWlanTemplateResource{}
}

type orgWlanTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgWlanTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WlanTemplate client")
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
func (r *orgWlanTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wlantemplate"
}

func (r *orgWlanTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategorySite + "This resource manages the Org WLAN Templates." +
			"A WLAN template is a collection of WLANs, tunneling policies, and wxlan policies. " +
			"It is used to create and manage wlan configurations at an organizational level. " +
			"WLAN templates allow for modular, scalable, and easy-to-manage configuration of ssids and their application to specific sites, " +
			"site groups, or ap device profiles. " +
			"They are valuable for automating configuration across multiple sites and profiles, making it easier to scale efficiently.",
		Attributes: resource_org_wlantemplate.OrgWlantemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgWlanTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WlanTemplate Create")
	var plan, state resource_org_wlantemplate.OrgWlantemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wlantemplate, diags := resource_org_wlantemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	data, err := r.client.OrgsWLANTemplates().CreateOrgTemplate(ctx, orgId, wlantemplate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WlanTemplate",
			"Could not create WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wlantemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgWlanTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_wlantemplate.OrgWlantemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WlanTemplate Read: wlantemplate_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wlantemplateId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsWLANTemplates().GetOrgTemplate(ctx, orgId, wlantemplateId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WlanTemplate",
			"Could not get WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_wlantemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgWlanTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_wlantemplate.OrgWlantemplateModel
	tflog.Info(ctx, "Starting WlanTemplate Update")

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

	wlantemplate, diags := resource_org_wlantemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WlanTemplate Update for WlanTemplate "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wlantemplateId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsWLANTemplates().UpdateOrgTemplate(ctx, orgId, wlantemplateId, wlantemplate)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WlanTemplate",
			"Could not update WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wlantemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgWlanTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wlantemplate.OrgWlantemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WlanTemplate Delete: wlantemplate_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wlantemplateId := uuid.MustParse(state.Id.ValueString())
	httpr, err := r.client.OrgsWLANTemplates().DeleteOrgTemplate(ctx, orgId, wlantemplateId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WlanTemplate",
			"Could not delete WlanTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
