package resource_device_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func NetworksTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchNetwork {
	data := make(map[string]models.SwitchNetwork)
	for vlanName, vlanDataAttr := range d.Elements() {
		var vlanDataInterface interface{} = vlanDataAttr
		netPlan := vlanDataInterface.(NetworksValue)

		netData := models.SwitchNetwork{}
		if netPlan.VlanId.ValueStringPointer() != nil {
			netData.VlanId = models.VlanIdWithVariableContainer.FromString(netPlan.VlanId.ValueString())
		}
		if netPlan.Subnet.ValueStringPointer() != nil {
			netData.Subnet = models.ToPointer(netPlan.Subnet.ValueString())
		}
		if netPlan.Subnet6.ValueStringPointer() != nil {
			netData.Subnet6 = models.ToPointer(netPlan.Subnet6.ValueString())
		}
		if netPlan.Gateway.ValueStringPointer() != nil {
			netData.Gateway = models.ToPointer(netPlan.Gateway.ValueString())
		}
		if netPlan.Gateway6.ValueStringPointer() != nil {
			netData.Gateway6 = models.ToPointer(netPlan.Gateway6.ValueString())
		}
		if netPlan.Isolation.ValueBoolPointer() != nil {
			netData.Isolation = models.ToPointer(netPlan.Isolation.ValueBool())
		}
		if netPlan.IsolationVlanId.ValueStringPointer() != nil {
			netData.IsolationVlanId = models.ToPointer(netPlan.IsolationVlanId.ValueString())
		}
		data[vlanName] = netData
	}
	return data
}
