package resource_device_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portConfigOverwriteTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchPortConfigOverwrite {

	data := make(map[string]models.SwitchPortConfigOverwrite)
	for k, v := range d.Elements() {
		var planInterface interface{} = v
		planObj := planInterface.(PortConfigOverwriteValue)
		itemObj := models.SwitchPortConfigOverwrite{}
		if planObj.Description.ValueStringPointer() != nil {
			itemObj.Description = planObj.Description.ValueStringPointer()
		}
		if planObj.Disabled.ValueBoolPointer() != nil {
			itemObj.Disabled = planObj.Disabled.ValueBoolPointer()
		}
		if planObj.Duplex.ValueStringPointer() != nil {
			itemObj.Duplex = models.ToPointer(models.SwitchPortUsageDuplexOverwriteEnum(planObj.Duplex.ValueString()))
		}
		if planObj.MacLimit.ValueStringPointer() != nil {
			itemObj.MacLimit = models.ToPointer(models.SwitchPortUsageMacLimitOverwriteContainer.FromString(planObj.MacLimit.ValueString()))
		}
		if planObj.PoeDisabled.ValueBoolPointer() != nil {
			itemObj.PoeDisabled = models.ToPointer(planObj.PoeDisabled.ValueBool())
		}
		if planObj.PortNetwork.ValueStringPointer() != nil {
			itemObj.PortNetwork = models.ToPointer(planObj.PortNetwork.ValueString())
		}
		if planObj.Speed.ValueStringPointer() != nil {
			itemObj.Speed = models.ToPointer(models.SwitchPortUsageSpeedOverwriteEnum(planObj.Speed.ValueString()))
		}
		data[k] = itemObj
	}
	return data
}
