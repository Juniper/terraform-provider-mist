package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vrrpGroupsTerraformToSdk(d basetypes.MapValue) map[string]models.VrrpConfigGroup {
	dataMap := make(map[string]models.VrrpConfigGroup)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(GroupsValue)
		data := models.VrrpConfigGroup{}

		if plan.Preempt.ValueBoolPointer() != nil {
			data.Preempt = models.ToPointer(plan.Preempt.ValueBool())
		}
		if plan.Priority.ValueInt64Pointer() != nil {
			data.Priority = models.ToPointer(int(plan.Priority.ValueInt64()))
		}

		dataMap[k] = data
	}
	return dataMap
}

func vrrpTerraformToSdk(d VrrpConfigValue) *models.VrrpConfig {

	data := models.VrrpConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Groups.IsNull() && !d.Groups.IsUnknown() {
		data.Groups = vrrpGroupsTerraformToSdk(d.Groups)
	}

	return &data
}
