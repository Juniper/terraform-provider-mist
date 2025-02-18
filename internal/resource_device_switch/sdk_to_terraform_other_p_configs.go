package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func otherIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosOtherIpConfig) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var evpnAnycast basetypes.BoolValue
		var ip basetypes.StringValue
		var ip6 basetypes.StringValue
		var netmask basetypes.StringValue
		var netmask6 basetypes.StringValue
		var type4 basetypes.StringValue
		var type6 basetypes.StringValue

		if d.EvpnAnycast != nil {
			evpnAnycast = types.BoolValue(*d.EvpnAnycast)
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

		dataMapValue := map[string]attr.Value{
			"evpn_anycast": evpnAnycast,
			"ip":           ip,
			"ip6":          ip6,
			"netmask":      netmask,
			"netmask6":     netmask6,
			"type":         type4,
			"type6":        type6,
		}
		data, e := NewOtherIpConfigsValue(OtherIpConfigsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := OtherIpConfigsValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
