package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway_cluster"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &deviceGatewayClusterResource{}
	_ resource.ResourceWithConfigure = &deviceGatewayClusterResource{}
)

func NewDeviceGatewayClusterResource() resource.Resource {
	return &deviceGatewayClusterResource{}
}

type deviceGatewayClusterResource struct {
	client mistapi.ClientInterface
}

func (r *deviceGatewayClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceGatewayCluster client")
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
func (r *deviceGatewayClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway_cluster"
}

func (r *deviceGatewayClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Gateway Clusters." +
			"It can be used to form or unset a cluster with two Gateways assigned to the same site." +
			"Please check the Juniper Documentation first to validate the cabling between the Gateways",
		Attributes: resource_device_gateway_cluster.DeviceGatewayClusterResourceSchema(ctx).Attributes,
	}
}

func (r *deviceGatewayClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceGatewayCluster Create")
	var plan, state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway_cluster, diags := resource_device_gateway_cluster.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from plan",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from plan",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, "Starting DeviceGatewayCluster Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteId, deviceId, device_gateway_cluster)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_gateway_cluster",
			"Could not create device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, deviceId, &data.Data)
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

func (r *deviceGatewayClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Read: device_gateway_cluster_id "+state.DeviceId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}

	data, err := r.client.SitesDevicesWANCluster().GetSiteDeviceHaClusterNode(ctx, siteId, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster",
			"Could not get device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, deviceId, &data.Data)
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

func (r *deviceGatewayClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_gateway_cluster.DeviceGatewayClusterModel
	tflog.Info(ctx, "Starting DeviceGatewayCluster Update")

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

	device_gateway_cluster, diags := resource_device_gateway_cluster.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Update for DeviceGatewayCluster "+state.DeviceId.ValueString())

	siteIdState, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceIdState, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}
	// if the device Id or the Nodes changed, delete the cluster then recreate it
	// otherwise it means it's only the site id that changed, and there is no need to recreate the cluster
	if !plan.DeviceId.Equal(state.DeviceId) || !plan.Nodes.Equal(state.Nodes) {
		_, err = r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteIdState, deviceIdState)
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device_gateway_cluster to apply the cluster changes",
			"Could not delete device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	siteIdPlan, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceIdPlan, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteIdPlan, deviceIdPlan, device_gateway_cluster)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_gateway_cluster",
			"Could not update device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteIdPlan, deviceIdPlan, &data.Data)
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

func (r *deviceGatewayClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_gateway_cluster.DeviceGatewayClusterModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Delete: device_gateway_cluster_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster site_id from state",
			"Could not get device_gateway_cluster site_id, unexpected error: "+err.Error(),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway_cluster device_id from state",
			"Could not get device_gateway_cluster device_id, unexpected error: "+err.Error(),
		)
		return
	}

	_, err = r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteId, deviceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device_gateway_cluster",
			"Could not delete device_gateway_cluster, unexpected error: "+err.Error(),
		)
		return
	}
}
