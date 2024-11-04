package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func localPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosLocalPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		var critical basetypes.BoolValue
		var description basetypes.StringValue
		var disable_autoneg basetypes.BoolValue
		var duplex basetypes.StringValue
		var mtu basetypes.Int64Value
		var poe_disabled basetypes.BoolValue
		var speed basetypes.StringValue
		var usage basetypes.StringValue = types.StringValue(d.Usage)

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
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.PoeDisabled != nil {
			poe_disabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		data_map_attr_type := PortConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"critical":        critical,
			"description":     description,
			"disable_autoneg": disable_autoneg,
			"duplex":          duplex,
			"mtu":             mtu,
			"poe_disabled":    poe_disabled,
			"speed":           speed,
			"usage":           usage,
		}
		data, e := NewPortConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_item_value[k] = data
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
