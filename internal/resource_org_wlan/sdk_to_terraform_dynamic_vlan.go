package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicVlan) DynamicVlanValue {

	var default_vlan_id basetypes.Int64Value
	var enabled basetypes.BoolValue
	var local_vlan_ids basetypes.ListValue = types.ListNull(types.Int64Type)
	var type_dynamic_vlan basetypes.StringValue
	var vlans basetypes.MapValue = types.MapNull(types.StringType)

	if d != nil && d.DefaultVlanId.Value() != nil {
		default_vlan_id = types.Int64Value(int64(*d.DefaultVlanId.Value()))
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.LocalVlanIds != nil {
		local_vlan_ids = vlanIdsSkToTerraform(ctx, diags, d.LocalVlanIds)
	}
	if d != nil && d.Type != nil {
		type_dynamic_vlan = types.StringValue(string(*d.Type))
	}
	if d != nil && d.Vlans != nil {
		vlans_attr := make(map[string]attr.Value)
		for k, v := range d.Vlans {
			vlans_attr[k] = types.StringValue(string(v))
		}
		vlans = types.MapValueMust(basetypes.StringType{}, vlans_attr)
	}

	data_map_attr_type := DynamicVlanValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"default_vlan_id": default_vlan_id,
		"enabled":         enabled,
		"local_vlan_ids":  local_vlan_ids,
		"type":            type_dynamic_vlan,
		"vlans":           vlans,
	}
	data, e := NewDynamicVlanValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
