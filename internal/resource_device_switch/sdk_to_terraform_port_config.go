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
	mapItemValue := make(map[string]attr.Value)

	for k, d := range m {

		var aeDisableLacp basetypes.BoolValue
		var aeIdx basetypes.Int64Value
		var aeLacpSlow basetypes.BoolValue
		var aggregated basetypes.BoolValue
		var critical basetypes.BoolValue
		var description basetypes.StringValue
		var disableAutoneg basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamicUsage basetypes.StringValue
		var esilag basetypes.BoolValue
		var mtu basetypes.Int64Value
		var noLocalOverwrite basetypes.BoolValue
		var poeDisabled basetypes.BoolValue
		var speed basetypes.StringValue
		var usage = types.StringValue(d.Usage)

		if d.AeDisableLacp != nil {
			aeDisableLacp = types.BoolValue(*d.AeDisableLacp)
		}
		if d.AeIdx != nil {
			aeIdx = types.Int64Value(int64(*d.AeIdx))
		}
		if d.AeLacpSlow != nil {
			aeLacpSlow = types.BoolValue(*d.AeLacpSlow)
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
			disableAutoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicUsage.Value() != nil {
			dynamicUsage = types.StringValue(*d.DynamicUsage.Value())
		}
		if d.Esilag != nil {
			esilag = types.BoolValue(*d.Esilag)
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.NoLocalOverwrite != nil {
			noLocalOverwrite = types.BoolValue(*d.NoLocalOverwrite)
		}
		if d.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		dataMapValue := map[string]attr.Value{
			"ae_disable_lacp":    aeDisableLacp,
			"ae_idx":             aeIdx,
			"ae_lacp_slow":       aeLacpSlow,
			"aggregated":         aggregated,
			"critical":           critical,
			"description":        description,
			"disable_autoneg":    disableAutoneg,
			"duplex":             duplex,
			"dynamic_usage":      dynamicUsage,
			"esilag":             esilag,
			"mtu":                mtu,
			"no_local_overwrite": noLocalOverwrite,
			"poe_disabled":       poeDisabled,
			"speed":              speed,
			"usage":              usage,
		}
		data, e := NewPortConfigValue(PortConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}
	r, e := types.MapValueFrom(ctx, PortConfigValue{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return r
}
