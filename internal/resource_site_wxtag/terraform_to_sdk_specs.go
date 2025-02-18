package resource_site_wxtag

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func specsTerraformToSdk(plan basetypes.ListValue) []models.WxlanTagSpec {
	var dataList []models.WxlanTagSpec
	for _, v := range plan.Elements() {
		var vInterface interface{} = v
		p := vInterface.(SpecsValue)

		data := models.WxlanTagSpec{
			PortRange: models.ToPointer(p.PortRange.ValueString()),
			Protocol:  models.ToPointer(p.Protocol.ValueString()),
			Subnets:   misttransform.ListOfStringTerraformToSdk(p.Subnets),
		}

		dataList = append(dataList, data)
	}
	return dataList

}
