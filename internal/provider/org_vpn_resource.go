package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_vpn"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgVpnResource{}
	_ resource.ResourceWithConfigure   = &orgVpnResource{}
	_ resource.ResourceWithImportState = &orgVpnResource{}
)

func NewOrgVpn() resource.Resource {
	return &orgVpnResource{}
}

type orgVpnResource struct {
	client mistapi.ClientInterface
}

func (r *orgVpnResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Vpn client")
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

func (r *orgVpnResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_vpn"
}

func (r *orgVpnResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages the Org VPN.",
		Attributes:          resource_org_vpn.OrgVpnResourceSchema(ctx).Attributes,
	}
}

func (r *orgVpnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Vpn Create")
	var plan, state resource_org_vpn.OrgVpnModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	vpn, diags := resource_org_vpn.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsVPNs().CreateOrgVpns(ctx, orgId, vpn)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Vpn",
			"Could not create Vpn, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_vpn.SdkToTerraform(ctx, &data.Data)
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

func (r *orgVpnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_vpn.OrgVpnModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Vpn Read: vpn_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	vpnId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	httpr, err := r.client.OrgsVPNs().GetOrgVpn(ctx, orgId, vpnId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Vpn",
			"Could not get Vpn, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_vpn.SdkToTerraform(ctx, &httpr.Data)
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

func (r *orgVpnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_vpn.OrgVpnModel
	tflog.Info(ctx, "Starting Vpn Update")

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

	vpn, diags := resource_org_vpn.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Vpn Update for Vpn "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	vpnId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	data, err := r.client.OrgsVPNs().UpdateOrgVpn(ctx, orgId, vpnId, vpn)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Vpn",
			"Could not update Vpn, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_vpn.SdkToTerraform(ctx, &data.Data)
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

func (r *orgVpnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_vpn.OrgVpnModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Vpn Delete: vpn_id "+state.Id.ValueString())
	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	vpnId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_vpn\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}
	httpr, err := r.client.OrgsVPNs().DeleteOrgVpn(ctx, orgId, vpnId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Vpn",
			"Could not delete Vpn, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgVpnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org\" resource",
			"Could not parse the UUID: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
