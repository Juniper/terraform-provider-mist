package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		var ae_disable_lacp basetypes.BoolValue
		var ae_idx basetypes.Int64Value
		var ae_lacp_slow basetypes.BoolValue
		var aggregated basetypes.BoolValue
		var critical basetypes.BoolValue
		var description basetypes.StringValue
		var disable_autoneg basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamic_usage basetypes.StringValue
		var esilag basetypes.BoolValue
		var mtu basetypes.Int64Value
		var no_local_overwrite basetypes.BoolValue
		var poe_disabled basetypes.BoolValue
		var speed basetypes.StringValue
		var usage basetypes.StringValue = types.StringValue(d.Usage)

		if d.AeDisableLacp != nil {
			ae_disable_lacp = types.BoolValue(*d.AeDisableLacp)
		}
		if d.AeIdx != nil {
			ae_idx = types.Int64Value(int64(*d.AeIdx))
		}
		if d.AeLacpSlow != nil {
			ae_lacp_slow = types.BoolValue(*d.AeLacpSlow)
		}
		if d.Aggregated != nil {
			aggregated = types.BoolValue(*d.Aggregated)
		}
		if d.Critical != nil {
			critical = types.BoolValue(*d.Critical)
		}
		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			disable_autoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicUsage.Value() != nil {
			dynamic_usage = types.StringValue(*d.DynamicUsage.Value())
		}
		if d.Esilag != nil {
			esilag = types.BoolValue(*d.Esilag)
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.NoLocalOverwrite != nil {
			no_local_overwrite = types.BoolValue(*d.NoLocalOverwrite)
		}
		if d.PoeDisabled != nil {
			poe_disabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		data_map_attr_type := PortConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ae_disable_lacp":    ae_disable_lacp,
			"ae_idx":             ae_idx,
			"ae_lacp_slow":       ae_lacp_slow,
			"aggregated":         aggregated,
			"critical":           critical,
			"description":        description,
			"disable_autoneg":    disable_autoneg,
			"duplex":             duplex,
			"dynamic_usage":      dynamic_usage,
			"esilag":             esilag,
			"mtu":                mtu,
			"no_local_overwrite": no_local_overwrite,
			"poe_disabled":       poe_disabled,
			"speed":              speed,
			"usage":              usage,
		}
		data, e := NewPortConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_item_value[k] = data
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
