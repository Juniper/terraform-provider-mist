package resource_org_networktemplate

import (
	"context"

	mist_api "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
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
		var gateway basetypes.StringValue
		var gateway6 basetypes.StringValue
		var subnet basetypes.StringValue
		var subnet6 basetypes.StringValue
		var vlan_id basetypes.StringValue

		if d.Isolation != nil {
			isolation = types.BoolValue(*d.Isolation)
		}
		if d.IsolationVlanId != nil {
			isolation_vlan_id = types.StringValue(*d.IsolationVlanId)
		}
		if d.Gateway != nil {
			gateway = types.StringValue(*d.Gateway)
		}
		if d.Gateway6 != nil {
			gateway6 = types.StringValue(*d.Gateway6)
		}
		if d.Subnet != nil {
			subnet = types.StringValue(*d.Subnet)
		}
		if d.Subnet6 != nil {
			subnet6 = types.StringValue(*d.Subnet6)
		}
		vlan_id = mist_api.VlanAsString(d.VlanId)

		data_map_attr_type := NetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"isolation":         isolation,
			"isolation_vlan_id": isolation_vlan_id,
			"gateway":           gateway,
			"gateway6":          gateway6,
			"subnet":            subnet,
			"subnet6":           subnet6,
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
