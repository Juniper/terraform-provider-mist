package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicPskSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicPsk) DynamicPskValue {
	var default_psk basetypes.StringValue
	var default_vlan_id basetypes.Int64Value
	var enabled basetypes.BoolValue
	var force_lookup basetypes.BoolValue
	var source basetypes.StringValue
	var vlan_ids basetypes.ListValue = types.ListNull(types.Int64Type)

	if d != nil && d.DefaultPsk != nil {
		default_psk = types.StringValue(*d.DefaultPsk)
	}
	if d != nil && d.DefaultVlanId.Value() != nil {
		default_vlan_id = types.Int64Value(int64(*d.DefaultVlanId.Value()))
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.ForceLookup != nil {
		force_lookup = types.BoolValue(*d.ForceLookup)
	}
	if d != nil && d.Source != nil {
		source = types.StringValue(string(*d.Source))
	}
	if d != nil && d.VlanIds != nil {
		vlan_ids = vlanIdsSkToTerraform(ctx, diags, d.VlanIds)
	}

	data_map_attr_type := DynamicPskValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"default_psk":     default_psk,
		"default_vlan_id": default_vlan_id,
		"enabled":         enabled,
		"force_lookup":    force_lookup,
		"source":          source,
		"vlan_ids":        vlan_ids,
	}
	data, e := NewDynamicPskValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
