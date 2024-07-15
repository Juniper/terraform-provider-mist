package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func otherIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.JunosOtherIpConfig {
	tflog.Debug(ctx, "otherIpConfigTerraformToSdk")

	data_map := make(map[string]models.JunosOtherIpConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OtherIpConfigsValue)
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
		data_map[k] = data
	}
	return data_map
}
