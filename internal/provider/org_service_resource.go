package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"terraform-provider-mist/internal/resource_org_service"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgOrgServiceResource{}
	_ resource.ResourceWithConfigure = &orgOrgServiceResource{}
)

func NewOrgServiceResource() resource.Resource {
	return &orgOrgServiceResource{}
}

type orgOrgServiceResource struct {
	client mistapi.ClientInterface
}

func (r *orgOrgServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *orgOrgServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_service"
}
func (r *orgOrgServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_service.OrgServiceResourceSchema(ctx)
}

func (r *orgOrgServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting OrgService Create")
	var plan, state resource_org_service.OrgServiceModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	service, diags := resource_org_service.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	orgId := uuid.MustParse(plan.OrgId.ValueString())
	tflog.Info(ctx, "Starting OrgService Create for Org "+plan.OrgId.ValueString())
	data, err := r.client.OrgsServices().CreateOrgService(ctx, orgId, &service)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating service",
			"Could not create service, unexpected error: "+err.Error(),
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

func (r *orgOrgServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_service.OrgServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Read: service_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	serviceId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsServices().GetOrgService(ctx, orgId, serviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting service",
			"Could not get service, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_service.SdkToTerraform(ctx, &data.Data)
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

func (r *orgOrgServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

	service, diags := resource_org_service.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Update for OrgService "+plan.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	serviceId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsServices().UpdateOrgService(ctx, orgId, serviceId, &service)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating service",
			"Could not update service, unexpected error: "+err.Error(),
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

func (r *orgOrgServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_service.OrgServiceModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting OrgService Delete: service_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	serviceId := uuid.MustParse(state.Id.ValueString())
	_, err := r.client.OrgsServices().DeleteOrgService(ctx, orgId, serviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting service",
			"Could not delete service, unexpected error: "+err.Error(),
		)
		return
	}
}
