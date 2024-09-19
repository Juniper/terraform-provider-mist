package provider

import (
	"context"
	"fmt"
	"strings"

	mist_api_error "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway_cluster"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceGatewayClusterResource{}
	_ resource.ResourceWithConfigure   = &deviceGatewayClusterResource{}
	_ resource.ResourceWithImportState = &deviceGatewayClusterResource{}
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
		MarkdownDescription: docCategoryDevices + "This resource can be used to form or delete a Gateway" +
			" Clusters. It can be used with two Gateways assigned to the same site.\n" +
			"Once the Cluster is formed, it can be create just like a Gateway with the `mist_device_gateway` resource:\n" +
			"1. Claim the gateways and assign them to a site with the `mist_org_inventory` resource\n" +
			"2. Form the Cluster with the `mist_device_gateway_cluster` resource by providing the `site_id` and the two nodes " +
			"MAC Addresses (the first in the list will be the node0)\n" +
			"3. Configure the Cluster with the `mist_device_gateway` resource\n\n" +
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
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	// Generate the id of the cluster based on the node0 MAC Address
	var nodes []resource_device_gateway_cluster.NodesValue
	plan.Nodes.ElementsAs(ctx, &nodes, false)
	mac := "00000000-0000-0000-1000-" + nodes[0].Mac.ValueString()
	id, err := uuid.Parse(mac)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", id, err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceGatewayCluster Create on Site "+plan.SiteId.ValueString()+" for device "+plan.Id.ValueString())
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteId, id, device_gateway_cluster)

	api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
	if api_err != "" {
		resp.Diagnostics.AddError(
			"Error creating \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to create the Gateway Cluster. %s", api_err),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, id, &data.Data)
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

	tflog.Info(ctx, "Starting DeviceGatewayCluster Read: device_gateway_cluster_id "+state.Id.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	id, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.SitesDevicesWANCluster().GetSiteDeviceHaClusterNode(ctx, siteId, id)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting \"mist_device_gateway_cluster\" resource",
			"Unable to get the Gateway Cluster, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, id, &httpr.Data)
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

	tflog.Info(ctx, "Starting DeviceGatewayCluster Update for DeviceGatewayCluster "+state.Id.ValueString())

	siteIdState, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	idState, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}
	// if the device Id or the Nodes changed, delete the cluster then recreate it
	// otherwise it means it's only the site id that changed, and there is no need to recreate the cluster
	if !plan.Id.Equal(state.Id) || !plan.Nodes.Equal(state.Nodes) {
		_, err = r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteIdState, idState)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting  \"mist_device_gateway_cluster\" resource to apply the cluster changes",
				"Unable to delete Gateway Cluster, unexpected error: "+err.Error(),
			)
			return
		}
	}

	siteIdPlan, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	idPlan, err := uuid.Parse(plan.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", plan.Id.ValueString(), err.Error()),
		)
		return
	}
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteIdPlan, idPlan, device_gateway_cluster)
	if err != nil {
		api_err := mist_api_error.ProcessApiError(ctx, data.Response.StatusCode, data.Response.Body, err)
		if api_err != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_device_gateway_cluster\" resource",
				fmt.Sprintf("Unable to update the Gateway Cluster. %s", api_err),
			)
			return
		}
	}

	state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteIdPlan, idPlan, &data.Data)
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

	tflog.Info(ctx, "Starting DeviceGatewayCluster Delete: device_gateway_cluster_id "+state.Id.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	id, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteId, id)
	api_err := mist_api_error.ProcessApiError(ctx, data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && api_err != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to delete the Gateway Cluster. %s", api_err),
		)
		return
	}
}

func (r *deviceGatewayClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			"import \"id\" format must be \"{site_id}.{cluster_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{cluster_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to parse the the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{cluster_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
