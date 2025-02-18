package resource_device_gateway_cluster

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *DeviceGatewayClusterModel) (*models.GatewayCluster, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.GatewayCluster{}

	var nodes []models.GatewayClusterNode
	for _, d := range plan.Nodes.Elements() {
		var vInterface interface{} = d
		pNode := vInterface.(NodesValue)
		dNode := models.GatewayClusterNode{}

		dNode.Mac = pNode.Mac.ValueString()

		nodes = append(nodes, dNode)
	}
	data.Nodes = nodes

	return &data, diags
}
