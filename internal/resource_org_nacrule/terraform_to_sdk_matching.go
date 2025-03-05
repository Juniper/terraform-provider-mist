package resource_org_nacrule

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func matchingPortTypesTerraformToSdk(d basetypes.ListValue) (data []models.NacRuleMatchingPortTypeEnum) {

	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(basetypes.StringValue)
		dataItem := models.NacRuleMatchingPortTypeEnum(plan.ValueString())
		data = append(data, dataItem)
	}
	return data
}

func matchingTerraformToSdk(d MatchingValue) (data *models.NacRuleMatching) {

	if !d.AuthType.IsNull() && !d.AuthType.IsUnknown() {
		data.AuthType = models.ToPointer(models.NacAuthTypeEnum(d.AuthType.ValueString()))
	}
	if !d.Family.IsNull() && !d.Family.IsUnknown() {
		data.Family = misttransform.ListOfStringTerraformToSdk(d.Family)
	}
	if !d.Mfg.IsNull() && !d.Mfg.IsUnknown() {
		data.Mfg = misttransform.ListOfStringTerraformToSdk(d.Mfg)
	}
	if !d.Model.IsNull() && !d.Model.IsUnknown() {
		data.Model = misttransform.ListOfStringTerraformToSdk(d.Model)
	}
	if !d.Nactags.IsNull() && !d.Nactags.IsUnknown() {
		data.Nactags = misttransform.ListOfStringTerraformToSdk(d.Nactags)
	}
	if !d.OsType.IsNull() && !d.OsType.IsUnknown() {
		data.OsType = misttransform.ListOfStringTerraformToSdk(d.OsType)
	}
	if !d.PortTypes.IsNull() && !d.PortTypes.IsUnknown() {
		data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	}
	if !d.SiteIds.IsNull() && !d.SiteIds.IsUnknown() {
		data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(d.SitegroupIds)
	}
	if !d.Vendor.IsNull() && !d.Vendor.IsUnknown() {
		data.Vendor = misttransform.ListOfStringTerraformToSdk(d.Vendor)
	}

	return data
}

func notMatchingTerraformToSdk(d NotMatchingValue) (data *models.NacRuleMatching) {

	if !d.AuthType.IsNull() && !d.AuthType.IsUnknown() {
		data.AuthType = models.ToPointer(models.NacAuthTypeEnum(d.AuthType.ValueString()))
	}
	if !d.Family.IsNull() && !d.Family.IsUnknown() {
		data.Family = misttransform.ListOfStringTerraformToSdk(d.Family)
	}
	if !d.Mfg.IsNull() && !d.Mfg.IsUnknown() {
		data.Mfg = misttransform.ListOfStringTerraformToSdk(d.Mfg)
	}
	if !d.Model.IsNull() && !d.Model.IsUnknown() {
		data.Model = misttransform.ListOfStringTerraformToSdk(d.Model)
	}
	if !d.Nactags.IsNull() && !d.Nactags.IsUnknown() {
		data.Nactags = misttransform.ListOfStringTerraformToSdk(d.Nactags)
	}
	if !d.OsType.IsNull() && !d.OsType.IsUnknown() {
		data.OsType = misttransform.ListOfStringTerraformToSdk(d.OsType)
	}
	if !d.PortTypes.IsNull() && !d.PortTypes.IsUnknown() {
		data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	}
	if !d.SiteIds.IsNull() && !d.SiteIds.IsUnknown() {
		data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(d.SitegroupIds)
	}
	if !d.Vendor.IsNull() && !d.Vendor.IsUnknown() {
		data.Vendor = misttransform.ListOfStringTerraformToSdk(d.Vendor)
	}

	return data
}
