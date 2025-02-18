package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bandsTerraformToSdk(plan basetypes.ListValue) []models.Dot11BandEnum {

	var dataList []models.Dot11BandEnum
	for _, v := range plan.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(basetypes.StringValue)
		data := models.Dot11BandEnum(vPlan.ValueString())
		dataList = append(dataList, data)
	}

	return dataList
}
