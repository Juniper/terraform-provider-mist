package resource_device_gateway_cluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, siteId uuid.UUID, deviceId uuid.UUID, data *models.GatewayCluster) (DeviceGatewayClusterModel, diag.Diagnostics) {
	var state DeviceGatewayClusterModel
	var diags diag.Diagnostics

	var device_id types.String = types.StringValue(deviceId.String())
	var nodes types.List = types.ListNull(NodesValue{}.Type(ctx))
	var site_id types.String = types.StringValue(siteId.String())

	var nodes_list []NodesValue
	for _, d := range data.Nodes {

		mac := types.StringValue(d.Mac)

		data_map_attr_type := NodesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"mac": mac,
		}
		node_mac, e := NewNodesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		nodes_list = append(nodes_list, node_mac)
	}

	data_list_type := NodesValue{}.Type(ctx)
	nodes, e := types.ListValueFrom(ctx, data_list_type, nodes_list)
	diags.Append(e...)

	state.DeviceId = device_id
	state.Nodes = nodes
	state.SiteId = site_id

	return state, diags
}
