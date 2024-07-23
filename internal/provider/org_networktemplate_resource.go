package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_networktemplate"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNetworkTemplateResource{}
	_ resource.ResourceWithConfigure = &orgNetworkTemplateResource{}
)

func NewOrgNetworkTemplate() resource.Resource {
	return &orgNetworkTemplateResource{}
}

type orgNetworkTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgNetworkTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NetworkTemplate client")
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
func (r *orgNetworkTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_networktemplate"
}

func (r *orgNetworkTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWired + "This resource managed the Org Network Templates (Switch templates)." +
			"A network template is a predefined configuration that provides a consistent and reusable set of network settings for devices within an organization. " +
			"It includes various parameters such as ip addressing, vlan configurations, routing protocols, security policies, and other network-specific settings. " +
			"Network templates simplify the deployment and management of switches by ensuring consistent configurations across multiple devices and sites. " +
			"They help enforce standardization, reduce human error, and streamline troubleshooting and maintenance tasks.",
		Attributes: resource_org_networktemplate.OrgNetworktemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgNetworkTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NetworkTemplate Create")
	var plan, state resource_org_networktemplate.OrgNetworktemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	networktemplate, diags := resource_org_networktemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsNetworkTemplates().CreateOrgNetworkTemplate(ctx, orgId, &networktemplate)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating NetworkTemplate",
			"Could not create NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_networktemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgNetworkTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_networktemplate.OrgNetworktemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NetworkTemplate Read: networktemplate_id "+state.Id.ValueString())

	orgId := uuid.MustParse(state.OrgId.ValueString())
	templateId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsNetworkTemplates().GetOrgNetworkTemplate(ctx, orgId, templateId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NetworkTemplate",
			"Could not get NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_networktemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgNetworkTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_networktemplate.OrgNetworktemplateModel
	tflog.Info(ctx, "Starting NetworkTemplate Update")

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

	orgId := uuid.MustParse(state.OrgId.ValueString())
	templateId := uuid.MustParse(state.Id.ValueString())
	networktemplate, diags := resource_org_networktemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting NetworkTemplate Update for NetworkTemplate "+state.Id.ValueString())
	data, err := r.client.OrgsNetworkTemplates().UpdateOrgNetworkTemplates(ctx, orgId, templateId, &networktemplate)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating NetworkTemplate",
			"Could not update NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_networktemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgNetworkTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_networktemplate.OrgNetworktemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	templateId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting NetworkTemplate Delete: networktemplate_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsNetworkTemplates().DeleteOrgNetworkTemplate(ctx, orgId, templateId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting NetworkTemplate",
			"Could not delete NetworkTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
