package resource_device_gateway_cluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *DeviceGatewayClusterModel) (*models.GatewayCluster, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.GatewayCluster{}

	var nodes []models.GatewayClusterNode
	for _, d := range plan.Nodes.Elements() {
		var v_interface interface{} = d
		p_node := v_interface.(NodesValue)
		d_node := models.GatewayClusterNode{}

		d_node.Mac = p_node.Mac.ValueString()

		nodes = append(nodes, d_node)
	}
	data.Nodes = nodes

	return &data, diags
}
