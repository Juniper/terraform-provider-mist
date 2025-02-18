package resource_org_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func opticPortConfigTerraformToSdk(d basetypes.MapValue) map[string]models.OpticPortConfigPort {
	dataMap := make(map[string]models.OpticPortConfigPort)
	for k, dataAttr := range d.Elements() {
		var dataInterface interface{} = dataAttr
		plan := dataInterface.(OpticPortConfigValue)

		data := models.OpticPortConfigPort{}
		if plan.Channelized.ValueBoolPointer() != nil {
			data.Channelized = plan.Channelized.ValueBoolPointer()
		}
		if plan.Speed.ValueStringPointer() != nil {
			data.Speed = models.ToPointer(plan.Speed.ValueString())
		}
		dataMap[k] = data
	}
	return dataMap
}
