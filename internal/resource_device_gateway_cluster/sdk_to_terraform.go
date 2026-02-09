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
	var nodesList []NodesValue
	var diags diag.Diagnostics
	for _, item := range data.Nodes {
		dataMap := map[string]attr.Value{
			"mac": types.StringValue(item.Mac),
		}
		nodeMac, err := NewNodesValue(NodesValue{}.AttributeTypes(ctx), dataMap)
		diags.Append(err...)
		nodesList = append(nodesList, nodeMac)
	}

	nodes, err := types.ListValueFrom(ctx, NodesValue{}.Type(ctx), nodesList)
	diags.Append(err...)

	result := DeviceGatewayClusterModel{
		Id:     types.StringValue(deviceId.String()),
		SiteId: types.StringValue(siteId.String()),
		Nodes:  nodes,
	}

	return result, diags
}
