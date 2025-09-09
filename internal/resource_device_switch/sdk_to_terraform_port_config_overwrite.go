package resource_device_switch

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func portConfigOverwriteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchPortConfigOverwrite) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)

	for k, d := range m {

		var description basetypes.StringValue
		var disabled basetypes.BoolValue
		var duplex basetypes.StringValue
		var macLimit basetypes.StringValue
		var poeDisabled basetypes.BoolValue
		var portNetwork basetypes.StringValue
		var speed basetypes.StringValue

		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.MacLimit != nil {
			macLimit = mistutils.SwitchPortOverwriteUsageMacLimitAsString(d.MacLimit)
		}
		if d.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortNetwork != nil {
			portNetwork = types.StringValue(*d.PortNetwork)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		dataMapValue := map[string]attr.Value{
			"description":  description,
			"disabled":     disabled,
			"duplex":       duplex,
			"mac_limit":    macLimit,
			"poe_disabled": poeDisabled,
			"port_network": portNetwork,
			"speed":        speed,
		}
		data, e := NewPortConfigOverwriteValue(PortConfigOverwriteValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}
	r, e := types.MapValueFrom(ctx, PortConfigOverwriteValue{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return r
}
