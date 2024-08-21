package provider

import (
	"context"
	"fmt"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_network"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgNetworkResource{}
	_ resource.ResourceWithConfigure   = &orgNetworkResource{}
	_ resource.ResourceWithImportState = &orgNetworkResource{}
)

func NewOrgNetworkResource() resource.Resource {
	return &orgNetworkResource{}
}

type orgNetworkResource struct {
	client mistapi.ClientInterface
}

func (r *orgNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Network client")
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
func (r *orgNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_network"
}

func (r *orgNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWan + "This resource manages the WAN Assurance Networks." +
			"The Networks are used in the `service_policies` from the Gateway configuration and Gateway templates ",
		Attributes: resource_org_network.OrgNetworkResourceSchema(ctx).Attributes,
	}
}

func (r *orgNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Network Create")
	var plan, state resource_org_network.OrgNetworkModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network orgId from plan",
			"Could not get network orgId, unexpected error: "+err.Error(),
		)
		return
	}

	network, diags := resource_org_network.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Create for Org "+plan.OrgId.String())
	data, err := r.client.OrgsNetworks().CreateOrgNetwork(ctx, orgId, network)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating network",
			"Could not create network, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_network.SdkToTerraform(ctx, data.Data)
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

func (r *orgNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_network.OrgNetworkModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network orgId from plan",
			"Could not get network orgId, unexpected error: "+err.Error(),
		)
		return
	}
	networkId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network networkId from plan",
			"Could not get network networkId, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting Network Read: network_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsNetworks().GetOrgNetwork(ctx, orgId, networkId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network",
			"Could not get network, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_network.SdkToTerraform(ctx, httpr.Data)
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

func (r *orgNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_network.OrgNetworkModel
	tflog.Info(ctx, "Starting Network Update")

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
			"Error getting network orgId from plan",
			"Could not get network orgId, unexpected error: "+err.Error(),
		)
		return
	}
	networkId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network networkId from plan",
			"Could not get network networkId, unexpected error: "+err.Error(),
		)
		return
	}

	network, diags := resource_org_network.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting Network Update for Network "+state.Id.ValueString())
	data, err := r.client.OrgsNetworks().UpdateOrgNetwork(ctx, orgId, networkId, network)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating network",
			"Could not update network, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_network.SdkToTerraform(ctx, data.Data)
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

func (r *orgNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_network.OrgNetworkModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network orgId from plan",
			"Could not get network orgId, unexpected error: "+err.Error(),
		)
		return
	}
	networkId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting network networkId from plan",
			"Could not get network networkId, unexpected error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "Starting Network Delete: network_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsNetworks().DeleteOrgNetwork(ctx, orgId, networkId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting network",
			"Could not delete network, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	_, err := uuid.Parse(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting org id from import",
			"Could not get org id, unexpected error: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
