package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_gatewaytemplate"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgGatewaytemplateResource{}
	_ resource.ResourceWithConfigure = &orgGatewaytemplateResource{}
)

func NewOrgGatewayTemplate() resource.Resource {
	return &orgGatewaytemplateResource{}
}

type orgGatewaytemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgGatewaytemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist GatewayTemplate client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *models.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *orgGatewaytemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_gatewaytemplate"
}

func (r *orgGatewaytemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages the Gateway Templates." +
			"A Gateway template is used to define the static ip address and subnet mask of the hub device, along with the gateway. " +
			"It also allows for the selection of options such as enabling source nat and overriding the public ip for the hub if needed. " +
			"the endpoint selected in the gateway template ties the hub and spoke devices together and creates the auto-vpn tunnel.",
		Attributes: resource_org_gatewaytemplate.OrgGatewaytemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgGatewaytemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting GatewayTemplate Create")
	var plan, state resource_org_gatewaytemplate.OrgGatewaytemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate org_id from plan",
			"Could not get mist_gatewaytemplate org_id, unexpected error: "+err.Error(),
		)
		return
	}
	gatewaytemplate, diags := resource_org_gatewaytemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsGatewayTemplates().CreateOrgGatewayTemplate(ctx, orgId, gatewaytemplate)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating GatewayTemplate",
			"Could not create GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_gatewaytemplate.SdkToTerraform(ctx, &data.Data)
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

func (r *orgGatewaytemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_gatewaytemplate.OrgGatewaytemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate org_id from state",
			"Could not get mist_gatewaytemplate org_id, unexpected error: "+err.Error(),
		)
		return
	}

	templateId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate gatewaytemplate_id from state",
			"Could not get mist_gatewaytemplate gatewaytemplate_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting GatewayTemplate Read: gatewaytemplate_id "+state.Id.ValueString())
	data, err := r.client.OrgsGatewayTemplates().GetOrgGatewayTemplate(ctx, orgId, templateId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting GatewayTemplate",
			"Could not get GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_gatewaytemplate.SdkToTerraform(ctx, &data.Data)
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

func (r *orgGatewaytemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_gatewaytemplate.OrgGatewaytemplateModel
	tflog.Info(ctx, "Starting GatewayTemplate Update")

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

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate org_id from state",
			"Could not get mist_gatewaytemplate org_id, unexpected error: "+err.Error(),
		)
		return
	}
	templateId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate gatewaytemplate_id from state",
			"Could not get mist_gatewaytemplate gatewaytemplate_id, unexpected error: "+err.Error(),
		)
		return
	}
	gatewaytemplate, diags := resource_org_gatewaytemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting GatewayTemplate Update for GatewayTemplate "+state.Id.ValueString())
	data, err := r.client.OrgsGatewayTemplates().
		UpdateOrgGatewayTemplate(ctx, orgId, templateId, gatewaytemplate)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating GatewayTemplate",
			"Could not update GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_gatewaytemplate.SdkToTerraform(ctx, &data.Data)
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

func (r *orgGatewaytemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_gatewaytemplate.OrgGatewaytemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate org_id from state",
			"Could not get mist_gatewaytemplate org_id, unexpected error: "+err.Error(),
		)
		return
	}
	templateId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting mist_gatewaytemplate gatewaytemplate_id from state",
			"Could not get mist_gatewaytemplate gatewaytemplate_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting GatewayTemplate Delete: gatewaytemplate_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsGatewayTemplates().DeleteOrgGatewayTemplate(ctx, orgId, templateId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting GatewayTemplate",
			"Could not delete GatewayTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
