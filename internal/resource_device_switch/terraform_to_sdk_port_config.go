package resource_device_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portConfigTerraformToSdk(d basetypes.MapValue) map[string]models.JunosPortConfig {

	data := make(map[string]models.JunosPortConfig)
	for k, v := range d.Elements() {
		var planInterface interface{} = v
		planObj := planInterface.(PortConfigValue)
		itemObj := models.JunosPortConfig{}
		itemObj.Usage = planObj.Usage.ValueString()
		if planObj.AeDisableLacp.ValueBoolPointer() != nil {
			itemObj.AeDisableLacp = models.ToPointer(planObj.AeDisableLacp.ValueBool())
		}
		if planObj.AeIdx.ValueInt64Pointer() != nil {
			itemObj.AeIdx = models.ToPointer(int(planObj.AeIdx.ValueInt64()))
		}
		if planObj.AeLacpSlow.ValueBoolPointer() != nil {
			itemObj.AeLacpSlow = models.ToPointer(planObj.AeLacpSlow.ValueBool())
		}
		if planObj.Aggregated.ValueBoolPointer() != nil {
			itemObj.Aggregated = models.ToPointer(planObj.Aggregated.ValueBool())
		}
		if planObj.Critical.ValueBoolPointer() != nil {
			itemObj.Critical = models.ToPointer(planObj.Critical.ValueBool())
		}
		if planObj.Description.ValueStringPointer() != nil {
			itemObj.Description = models.ToPointer(planObj.Description.ValueString())
		}
		if planObj.DisableAutoneg.ValueBoolPointer() != nil {
			itemObj.DisableAutoneg = models.ToPointer(planObj.DisableAutoneg.ValueBool())
		}
		if planObj.Duplex.ValueStringPointer() != nil {
			itemObj.Duplex = models.ToPointer(models.JunosPortConfigDuplexEnum(planObj.Duplex.ValueString()))
		}
		if planObj.DynamicUsage.ValueStringPointer() != nil {
			itemObj.DynamicUsage = models.NewOptional(models.ToPointer(planObj.DynamicUsage.ValueString()))
		}
		if planObj.Esilag.ValueBoolPointer() != nil {
			itemObj.Esilag = models.ToPointer(planObj.Esilag.ValueBool())
		}
		if planObj.Mtu.ValueInt64Pointer() != nil {
			itemObj.Mtu = models.ToPointer(int(planObj.Mtu.ValueInt64()))
		}
		if planObj.NoLocalOverwrite.ValueBoolPointer() != nil {
			itemObj.NoLocalOverwrite = models.ToPointer(planObj.NoLocalOverwrite.ValueBool())
		}
		if planObj.PoeDisabled.ValueBoolPointer() != nil {
			itemObj.PoeDisabled = models.ToPointer(planObj.PoeDisabled.ValueBool())
		}
		if planObj.Speed.ValueStringPointer() != nil {
			itemObj.Speed = models.ToPointer(models.JunosPortConfigSpeedEnum(planObj.Speed.ValueString()))
		}
		if planObj.Usage.ValueStringPointer() != nil {
			itemObj.Usage = *planObj.Usage.ValueStringPointer()
		}
		data[k] = itemObj
	}
	return data
}
