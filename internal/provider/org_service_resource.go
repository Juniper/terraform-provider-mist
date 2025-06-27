package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_org_service"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgServiceResource{}
	_ resource.ResourceWithConfigure   = &orgServiceResource{}
	_ resource.ResourceWithImportState = &orgServiceResource{}
)

func NewOrgServiceResource() resource.Resource {
	return &orgServiceResource{}
}

type orgServiceResource struct {
	client mistapi.ClientInterface
}

func (r *orgServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist OrgService client")
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

func (r *orgServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_service"
}
func (r *orgServiceResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages WAN Assurance Services (Applications).\n\n" +
			"A Service is used to define a Custom Application that can be used in the `service_policies`. These Services can be referenced by their name in\n" +
			"* the Service Policies (`mist_org_servicepolicy.services`)\n" +
			"* the Gateway configuration (`mist_device_gateway.service_policies.services`)\n" +
			"* the Gateway Templates (`mist_org_gatewaytemplate.service_policies.services`)\n" +
			"* the HUB Profiles (`mist_org_deviceprofile_gateway.service_policies.services`)\n",
		Attributes: resource_org_service.OrgServiceResourceSchema(ctx).Attributes,
	}
}

func (r *orgServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgService Create")
	var plan, state resource_org_service.OrgServiceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	service, diags := resource_org_service.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}
	tflog.Info(ctx, "Starting OrgService Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsServices().CreateOrgService(ctx, orgId, &service)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_org_service\" resource",
			fmt.Sprintf("Unable to create the Service. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_service.SdkToTerraform(ctx, &data.Data)
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

func (r *orgServiceResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_service.OrgServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Read: service_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	serviceId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	httpr, err := r.client.OrgsServices().GetOrgService(ctx, orgId, serviceId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_org_service\" resource",
			"Unable to get the Service, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_service.SdkToTerraform(ctx, &httpr.Data)
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

func (r *orgServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_service.OrgServiceModel
	tflog.Info(ctx, "Starting OrgService Update")

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

	service, diags := resource_org_service.TerraformToSdk(&plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Update for OrgService "+plan.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	serviceId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.OrgsServices().UpdateOrgService(ctx, orgId, serviceId, &service)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
	if apiErr != "" {
		resp.Diagnostics.AddError(
			"Error updating \"mist_org_service\" resource",
			fmt.Sprintf("Unable to update the Service. %s", apiErr),
		)
		return
	}

	state, diags = resource_org_service.SdkToTerraform(ctx, &data.Data)
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

func (r *orgServiceResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_service.OrgServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Delete: service_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	serviceId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.OrgsServices().DeleteOrgService(ctx, orgId, serviceId)
	if data != nil {
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		if data.StatusCode != 404 && apiErr != "" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_org_service\" resource",
				fmt.Sprintf("Unable to delete the Service. %s", apiErr),
			)
			return
		}
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_org_service\" resource",
			"Unable to delete the Service, unexpected error: "+err.Error(),
		)
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *orgServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_service\" resource",
			"import \"id\" format must be \"{org_id}.{service_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{service_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_org_service\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{service_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
