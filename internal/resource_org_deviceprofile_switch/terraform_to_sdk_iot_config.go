package resource_org_deviceprofile_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func iotConfigTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchIotPort {
	data := make(map[string]models.SwitchIotPort)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(IotConfigValue)
		item := models.SwitchIotPort{}

		if vPlan.AlarmClass.ValueStringPointer() != nil {
			item.AlarmClass = models.ToPointer(models.SwitchIotPortAlarmClassEnum(vPlan.AlarmClass.ValueString()))
		}
		if vPlan.Enabled.ValueBoolPointer() != nil {
			item.Enabled = vPlan.Enabled.ValueBoolPointer()
		}
		if vPlan.InputSrc.ValueStringPointer() != nil {
			item.InputSrc = models.ToPointer(models.SwitchIotPortInputSrcEnum(vPlan.InputSrc.ValueString()))
		}
		if vPlan.Name.ValueStringPointer() != nil {
			item.Name = vPlan.Name.ValueStringPointer()
		}

		data[k] = item
	}
	return data
}
