package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func NetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchNetwork) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {

		var isolation basetypes.BoolValue
		var isolation_vlan_id basetypes.StringValue
		var subnet basetypes.StringValue
		var vlan_id basetypes.StringValue

		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		if d.IsolationVlanId != nil {
			isolation_vlan_id = types.StringValue(*d.IsolationVlanId)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		vlan_id = types.StringValue(d.VlanId.String())

		data_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"isolation":         isolation,
			"isolation_vlan_id": isolation_vlan_id,
			"subnet":            subnet,
			"vlan_id":           vlan_id,
		}
		data, e := NewNetworksValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := NetworksValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
