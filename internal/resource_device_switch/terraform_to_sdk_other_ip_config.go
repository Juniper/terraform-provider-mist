package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func otherIpConfigTerraformToSdk(d basetypes.MapValue) map[string]models.JunosOtherIpConfig {

	dataMap := make(map[string]models.JunosOtherIpConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(OtherIpConfigsValue)
		data := models.JunosOtherIpConfig{}

		if plan.EvpnAnycast.ValueBoolPointer() != nil {
			data.EvpnAnycast = plan.EvpnAnycast.ValueBoolPointer()
		}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = plan.Ip.ValueStringPointer()
		}
		if plan.Ip6.ValueStringPointer() != nil {
			data.Ip6 = plan.Ip6.ValueStringPointer()
		}
		if plan.Netmask.ValueStringPointer() != nil {
			data.Netmask = plan.Netmask.ValueStringPointer()
		}
		if plan.Netmask6.ValueStringPointer() != nil {
			data.Netmask6 = plan.Netmask6.ValueStringPointer()
		}
		if plan.OtherIpConfigsType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.IpTypeEnum(plan.OtherIpConfigsType.ValueString()))
		}
		if plan.Type6.ValueStringPointer() != nil {
			data.Type6 = models.ToPointer(models.IpType6Enum(plan.Type6.ValueString()))
		}
		dataMap[k] = data
	}
	return dataMap
}
