package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func otherIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosOtherIpConfig) basetypes.MapValue {
	tflog.Debug(ctx, "ipConfigsSdkToTerraform")

	state_value_map := make(map[string]attr.Value)
	for k, d := range m {

		var evpn_anycast basetypes.BoolValue
		var ip basetypes.StringValue
		var ip6 basetypes.StringValue
		var netmask basetypes.StringValue
		var netmask6 basetypes.StringValue
		var type4 basetypes.StringValue
		var type6 basetypes.StringValue

		if d.EvpnAnycast != nil {
			evpn_anycast = types.BoolValue(*d.EvpnAnycast)
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Ip6 != nil {
			ip6 = types.StringValue(*d.Ip6)
		}
		if d.Netmask != nil {
			netmask = types.StringValue(*d.Netmask)
		}
		if d.Netmask6 != nil {
			netmask6 = types.StringValue(*d.Netmask6)
		}
		if d.Type != nil {
			type4 = types.StringValue(string(*d.Type))
		}
		if d.Type6 != nil {
			type6 = types.StringValue(string(*d.Type6))
		}

		data_map_attr_type := OtherIpConfigsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"evpn_anycast": evpn_anycast,
			"ip":           ip,
			"ip6":          ip6,
			"netmask":      netmask,
			"netmask6":     netmask6,
			"type":         type4,
			"type6":        type6,
		}
		data, e := NewOtherIpConfigsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := OtherIpConfigsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
