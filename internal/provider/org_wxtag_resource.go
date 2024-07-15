package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"terraform-provider-mist/internal/resource_org_wxtag"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWxTagResource{}
	_ resource.ResourceWithConfigure = &orgWxTagResource{}
)

func NewOrgWxTag() resource.Resource {
	return &orgWxTagResource{}
}

type orgWxTagResource struct {
	client mistapi.ClientInterface
}

func (r *orgWxTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist WxTag client")
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
func (r *orgWxTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wxtag"
}

func (r *orgWxTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resource_org_wxtag.OrgWxtagResourceSchema(ctx)
}

func (r *orgWxTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting WxTag Create")
	var plan, state resource_org_wxtag.OrgWxtagModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wxtag, diags := resource_org_wxtag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId := uuid.MustParse(plan.OrgId.ValueString())
	data, err := r.client.OrgsWxTags().CreateOrgWxTag(ctx, orgId, wxtag)
	if err != nil {
		//url, _ := httpr.Location()
		resp.Diagnostics.AddError(
			"Error creating WxTag",
			"Could not create WxTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *orgWxTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_wxtag.OrgWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Read: wxtag_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsWxTags().GetOrgWxTag(ctx, orgId, wxtagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting WxTag",
			"Could not get WxTag, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *orgWxTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_wxtag.OrgWxtagModel
	tflog.Info(ctx, "Starting WxTag Update")

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

	wxtag, diags := resource_org_wxtag.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Update for WxTag "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())
	data, err := r.client.OrgsWxTags().UpdateOrgWxTag(ctx, orgId, wxtagId, wxtag)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WxTag",
			"Could not update WxTag, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_wxtag.SdkToTerraform(ctx, data.Data)
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

func (r *orgWxTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wxtag.OrgWxtagModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting WxTag Delete: wxtag_id "+state.Id.ValueString())
	orgId := uuid.MustParse(state.OrgId.ValueString())
	wxtagId := uuid.MustParse(state.Id.ValueString())
	_, err := r.client.OrgsWxTags().DeleteOrgWxTag(ctx, orgId, wxtagId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WxTag",
			"Could not delete WxTag, unexpected error: "+err.Error(),
		)
		return
	}
}
