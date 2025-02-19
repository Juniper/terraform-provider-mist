package provider

import (
	"context"
	"fmt"
	"strings"

	mistapierror "github.com/Juniper/terraform-provider-mist/internal/commons/api_response_error"
	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway_cluster"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

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
func (r *deviceGatewayClusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway_cluster"
}

func (r *deviceGatewayClusterResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource can be used to form or delete a Gateway Clusters.\n\n" +
			"A Gateway Cluster can be formed with two Gateways assigned to the same site. " +
			"Once the Cluster is formed, it can be configured just like a Gateway with the `mist_device_gateway` resource:\n" +
			"1. Claim the gateways and assign them to the same site with the `mist_org_inventory` resource\n" +
			"2. Form the Cluster with the `mist_device_gateway_cluster` resource by providing the `site_id` and the MAC Addresses " +
			"of two nodes (the first in the list will be the node0)\n" +
			"3. Configure the Cluster with the `mist_device_gateway` resource\n\n" +
			"Please check the [SRX Juniper Documentation](https://www.juniper.net/documentation/us/en/software/mist/mist-wan/topics/topic-map/srx-high-availability-configuration.html)" +
			"or the [SSR Juniper Documentation](https://www.juniper.net/documentation/us/en/software/mist/mist-wan/topics/topic-map/ssr-high-availability-configuration.html)" +
			" first to validate the cabling between the Gateways\n\n" +
			"~> Both gateways must belong to the same site when creating the Gateway Cluster",
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

	deviceGatewayCluster, diags := resource_device_gateway_cluster.TerraformToSdk(&plan)
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
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteId, id, deviceGatewayCluster)

	apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)

	if apiErr != "" {
		PlannedNode0Mac := nodes[0].Mac.ValueString()
		PlannedNode1Mac := nodes[1].Mac.ValueString()
		// if the cluster already exists and the primary node is the same
		// this can be detected because a cluster already exists with the node0 MAC Address UUID
		if strings.Contains(apiErr, "already belong to a ha cluster") {
			existingGatewayCluster, sameNodes := r.checkClusterNodes(ctx, siteId, PlannedNode0Mac, PlannedNode1Mac)
			// same nodes, same order
			if existingGatewayCluster != nil && sameNodes {
				state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, id, existingGatewayCluster)
				// same node0, other node1
			} else if existingGatewayCluster != nil {
				mistNode1Mac := existingGatewayCluster.Nodes[1].Mac
				resp.Diagnostics.AddError(
					"Error creating \"mist_device_gateway_cluster\" resource",
					fmt.Sprintf("The primary node %s already belong to a ha cluster with a different secondary node %s. Got %s", PlannedNode0Mac, mistNode1Mac, PlannedNode1Mac),
				)
				return
				// other error
			} else {
				resp.Diagnostics.AddError(
					"Error creating \"mist_device_gateway_cluster\" resource",
					fmt.Sprintf("Unable to create the Gateway Cluster. %s", apiErr),
				)
				return
			}
			// if the cluster already exists, and we are suspecting a node invertion
			// this can be detected because the node0 is assigned to the site but there is no device/cluster with its MAC Address
		} else if strings.Contains(apiErr, "resource not found") {
			existingGatewayCluster, sameNodes := r.checkClusterNodes(ctx, siteId, PlannedNode1Mac, PlannedNode0Mac)
			// same nodes, inverted order
			if existingGatewayCluster != nil && sameNodes {
				resp.Diagnostics.AddError(
					"Error creating \"mist_device_gateway_cluster\" resource",
					fmt.Sprintf("The nodes %s and %s already belong to the same Cluster, but the nodes are inverted", PlannedNode0Mac, PlannedNode1Mac),
				)
				return
				// node1 already belongs to a cluster with another node
			} else if existingGatewayCluster != nil {
				mistOtherNodeMac := existingGatewayCluster.Nodes[1].Mac
				resp.Diagnostics.AddError(
					"Error creating \"mist_device_gateway_cluster\" resource",
					fmt.Sprintf("The node %s already belong to a ha cluster with a different node %s. Got %s", PlannedNode1Mac, mistOtherNodeMac, PlannedNode0Mac),
				)
				return
				// other error
			} else {
				resp.Diagnostics.AddError(
					"Error creating \"mist_device_gateway_cluster\" resource",
					fmt.Sprintf(
						"Unable to create the Gateway Cluster. %s."+
							"\nEither the node %s is not assigned to the correct site, either is already belongs to another cluster.",
						apiErr, PlannedNode0Mac),
				)
				return
			}
			// for any other error
		} else {
			resp.Diagnostics.AddError(
				"Error creating \"mist_device_gateway_cluster\" resource",
				fmt.Sprintf("Unable to create the Gateway Cluster. %s", apiErr),
			)
			return
		}
	} else {
		state, diags = resource_device_gateway_cluster.SdkToTerraform(ctx, siteId, id, &data.Data)
	}

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

func (r *deviceGatewayClusterResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
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
	if httpr.Response.StatusCode == 200 && len(httpr.Data.Nodes) == 0 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode == 404 {
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

	deviceGatewayCluster, diags := resource_device_gateway_cluster.TerraformToSdk(&plan)
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
		data, err := r.client.SitesDevicesWANCluster().DeleteSiteDeviceHaCluster(ctx, siteIdState, idState)
		apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
		// if the device is not found or if the device is not a cluster, skipping the error
		if data.StatusCode != 404 && apiErr != "not a ha cluster" {
			resp.Diagnostics.AddError(
				"Error deleting \"mist_device_gateway_cluster\" resource",
				fmt.Sprintf("Unable to delete the Gateway Cluster. %s", apiErr),
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
	data, err := r.client.SitesDevicesWANCluster().CreateSiteDeviceHaCluster(ctx, siteIdPlan, idPlan, deviceGatewayCluster)
	if err != nil {
		apiErr := mistapierror.ProcessApiError(data.Response.StatusCode, data.Response.Body, err)
		if apiErr != "" {
			resp.Diagnostics.AddError(
				"Error creating \"mist_device_gateway_cluster\" resource",
				fmt.Sprintf("Unable to update the Gateway Cluster. %s", apiErr),
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

func (r *deviceGatewayClusterResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
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
	apiErr := mistapierror.ProcessApiError(data.StatusCode, data.Body, err)
	if data.StatusCode != 404 && apiErr != "not a ha cluster" && apiErr != "" {
		resp.Diagnostics.AddError(
			"Error deleting \"mist_device_gateway_cluster\" resource",
			fmt.Sprintf("Unable to delete the Gateway Cluster. %s", apiErr),
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

// Additional functions

func (r *deviceGatewayClusterResource) checkClusterNodes(
	ctx context.Context,
	siteId uuid.UUID,
	node0Mac string,
	node1Mac string,
) (*models.GatewayCluster, bool) {
	/*
		function used when the provider tries to create the cluster and get an error response because the node already
		belongs to another cluster. It will retrieve the MAC Address of the current cluster nodes, and compare them to the plan

		parameters:
			ctx: context.Context
			siteId: uuid.UUID
				planned siteId of the cluster
			node0Mac: string
				MAC Address of the planned node0
			node1Mac: string
				MAC Address of the planned node1

		returns:
			*models.GatewayCluster
				If exact same nodes (and not inverted), returns the current gateway cluster response data
			bool
				if the current nodes (in Mist) are the planned nodes
	*/
	clusterId, e := uuid.Parse("00000000-0000-0000-1000-" + node0Mac)
	if e == nil {
		return r.checkClusterNode(ctx, siteId, clusterId, node0Mac, node1Mac)
	}
	return nil, false

}

func (r *deviceGatewayClusterResource) checkClusterNode(
	ctx context.Context,
	siteId uuid.UUID,
	clusterId uuid.UUID,
	primaryNodeMac string,
	secondaryNodeMac string,
) (*models.GatewayCluster, bool) {
	httpr, err := r.client.SitesDevicesWANCluster().GetSiteDeviceHaClusterNode(ctx, siteId, clusterId)

	if err == nil && len(httpr.ApiResponse.Data.Nodes) > 0 {
		sameNodes := true
		for _, node := range httpr.ApiResponse.Data.Nodes {
			if node.Mac != primaryNodeMac && node.Mac != secondaryNodeMac {
				sameNodes = false
			}
		}
		return &httpr.Data, sameNodes
	}
	return nil, false

}
