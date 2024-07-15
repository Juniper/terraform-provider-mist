package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_nactag"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgNacTagResource{}
	_ resource.ResourceWithConfigure = &orgNacTagResource{}
)

func NewOrgNacTag() resource.Resource {
	return &orgNacTagResource{}
}

type orgNacTagResource struct {
	client mistapi.ClientInterface
}

func (r *orgNacTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist NacTag client")
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
func (r *orgNacTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_nactag"
}

func (r *orgNacTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_nactag.OrgNactagResourceSchema(ctx)
}

func (r *orgNacTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting NacTag Create")
	var plan, state resource_org_nactag.OrgNactagModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	nactag, diags := resource_org_nactag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsNACTags().CreateOrgNacTag(ctx, orgId, &nactag)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating NacTag",
			"Could not create NacTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nactag.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_nactag.OrgNactagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	nactagId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting NacTag Read: nactag_id "+state.Id.ValueString())
	data, err := r.client.OrgsNACTags().GetOrgNacTag(ctx, orgId, nactagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting NacTag",
			"Could not get NacTag, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_nactag.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_nactag.OrgNactagModel
	tflog.Info(ctx, "Starting NacTag Update")

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

	nactag, diags := resource_org_nactag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	nactagId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting NacTag Update for NacTag "+state.Id.ValueString())
	data, err := r.client.OrgsNACTags().
		UpdateOrgNacTag(ctx, orgId, nactagId, &nactag)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating NacTag",
			"Could not update NacTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_nactag.SdkToTerraform(ctx, data.Data)
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

func (r *orgNacTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_nactag.OrgNactagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(state.OrgId.ValueString())
	nactagId := uuid.MustParse(state.Id.ValueString())
	tflog.Info(ctx, "Starting NacTag Delete: nactag_id "+state.Id.ValueString())
	_, err := r.client.OrgsNACTags().DeleteOrgNacTag(ctx, orgId, nactagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting NacTag",
			"Could not delete NacTag, unexpected error: "+err.Error(),
		)
		return
	}
}
