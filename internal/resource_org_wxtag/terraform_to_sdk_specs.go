package resource_org_wxtag

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
			Subnets:   mistutils.ListOfStringTerraformToSdk(p.Subnets),
		}

		dataList = append(dataList, data)
	}
	return dataList

}
