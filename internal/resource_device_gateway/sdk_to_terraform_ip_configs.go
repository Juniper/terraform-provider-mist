package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func ipConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayIpConfigProperty) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var ip basetypes.StringValue
		var netmask basetypes.StringValue
		var secondaryIps = types.ListNull(types.StringType)
		var typeIp basetypes.StringValue

		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Netmask != nil {
			netmask = types.StringValue(*d.Netmask)
		}
		if d.SecondaryIps != nil {
			secondaryIps = mistutils.ListOfStringSdkToTerraform(d.SecondaryIps)
		}
		if d.Type != nil {
			typeIp = types.StringValue(string(*d.Type))
		}

		dataMapValue := map[string]attr.Value{
			"ip":            ip,
			"netmask":       netmask,
			"secondary_ips": secondaryIps,
			"type":          typeIp,
		}
		data, e := NewIpConfigsValue(IpConfigsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := IpConfigsValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
