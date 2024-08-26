package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_psk"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgPskResource{}
	_ resource.ResourceWithConfigure   = &orgPskResource{}
	_ resource.ResourceWithImportState = &orgPskResource{}
)

func NewOrgPsk() resource.Resource {
	return &orgPskResource{}
}

type orgPskResource struct {
	client mistapi.ClientInterface
}

func (r *orgPskResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Psk client")
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
func (r *orgPskResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_psk"
}

func (r *orgPskResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This data source provides the list of Org PSKs.",
		Attributes:          resource_org_psk.OrgPskResourceSchema(ctx).Attributes,
	}
}

func (r *orgPskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Psk Create")
	var plan, state resource_org_psk.OrgPskModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	psk, diags := resource_org_psk.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var upsert bool = false
	data, err := r.client.OrgsPsks().CreateOrgPsk(ctx, orgId, &upsert, &psk)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Psk",
			"Could not create Psk, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_psk.SdkToTerraform(ctx, &data.Data)
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

func (r *orgPskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_psk.OrgPskModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	pskId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Psk Read: psk_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsPsks().GetOrgPsk(ctx, orgId, pskId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Psk",
			"Could not get Psk, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_psk.SdkToTerraform(ctx, &httpr.Data)
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

func (r *orgPskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_psk.OrgPskModel
	tflog.Info(ctx, "Starting Psk Update")

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

	psk, diags := resource_org_psk.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	pskId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Psk Update for Psk "+state.Id.ValueString())
	data, err := r.client.OrgsPsks().
		UpdateOrgPsk(ctx, orgId, pskId, &psk)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Psk",
			"Could not update Psk, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_psk.SdkToTerraform(ctx, &data.Data)
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

func (r *orgPskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_psk.OrgPskModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	pskId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Psk Delete: psk_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsPsks().DeleteOrgPsk(ctx, orgId, pskId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Psk",
			"Could not delete Psk, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgPskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_psk\" resource",
			"import \"id\" format must be \"{org_id}.{psk_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{psk_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_psk\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{psk_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
