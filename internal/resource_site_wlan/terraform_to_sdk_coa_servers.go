package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServerTerraformToSdk(plan basetypes.ListValue) []models.CoaServer {

	var dataList []models.CoaServer
	for _, v := range plan.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(CoaServersValue)
		data := models.CoaServer{}
		if vPlan.Ip.ValueStringPointer() != nil {
			data.Ip = vPlan.Ip.ValueString()
		}
		if vPlan.Secret.ValueStringPointer() != nil {
			data.Secret = vPlan.Secret.ValueString()
		}
		if vPlan.DisableEventTimestampCheck.ValueBoolPointer() != nil {
			data.DisableEventTimestampCheck = vPlan.DisableEventTimestampCheck.ValueBoolPointer()
		}
		if vPlan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = vPlan.Enabled.ValueBoolPointer()
		}
		if vPlan.Port.ValueStringPointer() != nil {
			data.Port = models.ToPointer(models.RadiusCoaPortContainer.FromString(vPlan.Port.ValueString()))
		}
		dataList = append(dataList, data)
	}

	return dataList
}
