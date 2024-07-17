package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_rftemplate"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgRfTemplateResource{}
	_ resource.ResourceWithConfigure = &orgRfTemplateResource{}
)

func NewOrgRfTemplate() resource.Resource {
	return &orgRfTemplateResource{}
}

type orgRfTemplateResource struct {
	client mistapi.ClientInterface
}

func (r *orgRfTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist RfTemplate client")
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
func (r *orgRfTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_rftemplate"
}

func (r *orgRfTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This manages the RF Templates." +
			"The RF Templates can be used to define Wireless Access Points radio configuration, and can be assigned to the sites",
		Attributes: resource_org_rftemplate.OrgRftemplateResourceSchema(ctx).Attributes,
	}
}

func (r *orgRfTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting RfTemplate Create")
	var plan, state resource_org_rftemplate.OrgRftemplateModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	rftemplate, diags := resource_org_rftemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	data, err := r.client.OrgsRFTemplates().CreateOrgRfTemplate(ctx, orgId, rftemplate)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating RfTemplate",
			"Could not create RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_rftemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgRfTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_rftemplate.OrgRftemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	rftemplateId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting RfTemplate Read: rftemplate_id "+state.Id.ValueString())
	data, err := r.client.OrgsRFTemplates().GetOrgRfTemplate(ctx, orgId, rftemplateId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting RfTemplate",
			"Could not get RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_rftemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgRfTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_rftemplate.OrgRftemplateModel
	tflog.Info(ctx, "Starting RfTemplate Update")

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

	rftemplate, diags := resource_org_rftemplate.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	rftemplateId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting RfTemplate Update for RfTemplate "+state.Id.ValueString())
	data, err := r.client.OrgsRFTemplates().UpdateOrgRfTemplate(ctx, orgId, rftemplateId, rftemplate)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating RfTemplate",
			"Could not update RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_rftemplate.SdkToTerraform(ctx, data.Data)
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

func (r *orgRfTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_rftemplate.OrgRftemplateModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	rftemplateId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting RfTemplate Delete: rftemplate_id "+state.Id.ValueString())
	_, err := r.client.OrgsRFTemplates().DeleteOrgRfTemplate(ctx, orgId, rftemplateId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting RfTemplate",
			"Could not delete RfTemplate, unexpected error: "+err.Error(),
		)
		return
	}
}
