package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func ipConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayIpConfigProperty) basetypes.MapValue {
	tflog.Debug(ctx, "ipConfigsSdkToTerraform")

	state_value_map := make(map[string]attr.Value)
	for k, d := range m {

		var ip basetypes.StringValue
		var netmask basetypes.StringValue
		var secondary_ips basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var type_ip basetypes.StringValue

		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Netmask != nil {
			netmask = types.StringValue(*d.Netmask)
		}
		if d.SecondaryIps != nil {
			secondary_ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.SecondaryIps)
		}
		if d.Type != nil {
			type_ip = types.StringValue(string(*d.Type))
		}

		data_map_attr_type := IpConfigsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ip":            ip,
			"netmask":       netmask,
			"secondary_ips": secondary_ips,
			"type":          type_ip,
		}
		data, e := NewIpConfigsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := IpConfigsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
