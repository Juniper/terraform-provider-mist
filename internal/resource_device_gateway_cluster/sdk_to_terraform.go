package resource_device_gateway_cluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, mistSiteId uuid.UUID, mistDeviceId uuid.UUID, data *models.GatewayCluster) (DeviceGatewayClusterModel, diag.Diagnostics) {
	var state DeviceGatewayClusterModel
	var diags diag.Diagnostics

	var id = types.StringValue(mistDeviceId.String())
	var nodes = types.ListNull(NodesValue{}.Type(ctx))
	var siteId = types.StringValue(mistSiteId.String())

	var nodesList []NodesValue
	for _, d := range data.Nodes {

		mac := types.StringValue(d.Mac)

		dataMapValue := map[string]attr.Value{
			"mac": mac,
		}
		nodeMac, e := NewNodesValue(NodesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		nodesList = append(nodesList, nodeMac)
	}

	datalistType := NodesValue{}.Type(ctx)
	nodes, e := types.ListValueFrom(ctx, datalistType, nodesList)
	diags.Append(e...)

	state.Id = id
	state.Nodes = nodes
	state.SiteId = siteId

	return state, diags
}
