package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func NetworksTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchNetwork {
	data := make(map[string]models.SwitchNetwork)
	for vlan_name, vlan_data_attr := range d.Elements() {
		var vlan_data_interface interface{} = vlan_data_attr
		net_plan := vlan_data_interface.(NetworksValue)

		net_data := models.SwitchNetwork{}
		if net_plan.VlanId.ValueInt64Pointer() != nil {
			net_data.VlanId = int(net_plan.VlanId.ValueInt64())
		}
		if net_plan.Subnet.ValueStringPointer() != nil {
			net_data.Subnet = models.ToPointer(net_plan.Subnet.ValueString())
		}
		if net_plan.Isolation.ValueBoolPointer() != nil {
			net_data.Isolation = models.ToPointer(net_plan.Isolation.ValueBool())
		}
		if net_plan.IsolationVlanId.ValueStringPointer() != nil {
			net_data.IsolationVlanId = models.ToPointer(net_plan.IsolationVlanId.ValueString())
		}
		data[vlan_name] = net_data
	}
	return data
}
