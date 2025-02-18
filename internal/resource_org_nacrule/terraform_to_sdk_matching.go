package resource_org_nacrule

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func matchingPortTypesTerraformToSdk(d basetypes.ListValue) []models.NacRuleMatchingPortTypeEnum {

	var data []models.NacRuleMatchingPortTypeEnum
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(basetypes.StringValue)
		dataItem := models.NacRuleMatchingPortTypeEnum(plan.ValueString())
		data = append(data, dataItem)
	}
	return data
}

func matchingTerraformToSdk(d MatchingValue) *models.NacRuleMatching {

	data := models.NacRuleMatching{}

	data.AuthType = models.ToPointer(models.NacRuleMatchingAuthTypeEnum(d.AuthType.ValueString()))
	data.Nactags = misttransform.ListOfStringTerraformToSdk(d.Nactags)
	data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(d.SitegroupIds)
	data.Vendor = misttransform.ListOfStringTerraformToSdk(d.Vendor)

	return &data
}

func notMatchingTerraformToSdk(d NotMatchingValue) *models.NacRuleMatching {

	data := models.NacRuleMatching{}

	data.AuthType = models.ToPointer(models.NacRuleMatchingAuthTypeEnum(d.AuthType.ValueString()))
	data.Nactags = misttransform.ListOfStringTerraformToSdk(d.Nactags)
	data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(d.SitegroupIds)
	data.Vendor = misttransform.ListOfStringTerraformToSdk(d.Vendor)

	return &data
}
