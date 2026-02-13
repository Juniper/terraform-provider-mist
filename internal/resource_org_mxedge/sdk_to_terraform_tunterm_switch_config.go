package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermSwitchConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermSwitchConfigs) basetypes.MapValue {

	state_value_map_type := TuntermSwitchConfigValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d.AdditionalProperties {
		var portVlanId types.Int64
		var vlanIds = types.ListNull(types.StringType)

		if v.PortVlanId != nil {
			portVlanId = types.Int64Value(int64(*v.PortVlanId))
		}
		if v.VlanIds != nil {
			vlanIds_list := make([]attr.Value, len(v.VlanIds))
			for i, vlanId := range v.VlanIds {
				vlanIdStr := vlanId.String()
				vlanIds_list[i] = types.StringValue(vlanIdStr)
			}
			vlanIds_result, e := types.ListValue(types.StringType, vlanIds_list)
			diags.Append(e...)
			vlanIds = vlanIds_result
		}

		data_map_attr_type := TuntermSwitchConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_vlan_id": portVlanId,
			"vlan_ids":     vlanIds,
		}
		data, e := NewTuntermSwitchConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}

	state_result, e := types.MapValueFrom(ctx, state_value_map_type, state_value_map)
	diags.Append(e...)
	return state_result
}
