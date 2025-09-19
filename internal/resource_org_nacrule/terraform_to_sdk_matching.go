package resource_org_nacrule

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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

func matchingTerraformToSdk(d MatchingValue) *models.NacRuleMatching {
	data := models.NacRuleMatching{}

	if !d.AuthType.IsNull() && !d.AuthType.IsUnknown() {
		data.AuthType = models.ToPointer(models.NacAuthTypeEnum(d.AuthType.ValueString()))
	}
	if !d.Family.IsNull() && !d.Family.IsUnknown() {
		data.Family = mistutils.ListOfStringTerraformToSdk(d.Family)
	}
	if !d.Mfg.IsNull() && !d.Mfg.IsUnknown() {
		data.Mfg = mistutils.ListOfStringTerraformToSdk(d.Mfg)
	}
	if !d.Model.IsNull() && !d.Model.IsUnknown() {
		data.Model = mistutils.ListOfStringTerraformToSdk(d.Model)
	}
	if !d.Nactags.IsNull() && !d.Nactags.IsUnknown() {
		data.Nactags = mistutils.ListOfStringTerraformToSdk(d.Nactags)
	}
	if !d.OsType.IsNull() && !d.OsType.IsUnknown() {
		data.OsType = mistutils.ListOfStringTerraformToSdk(d.OsType)
	}
	if !d.PortTypes.IsNull() && !d.PortTypes.IsUnknown() {
		data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	}
	if !d.SiteIds.IsNull() && !d.SiteIds.IsUnknown() {
		data.SiteIds = mistutils.ListOfUuidTerraformToSdk(d.SiteIds)
	}
	if !d.SitegroupIds.IsNull() && !d.SitegroupIds.IsUnknown() {
		data.SitegroupIds = mistutils.ListOfUuidTerraformToSdk(d.SitegroupIds)
	}
	if !d.Vendor.IsNull() && !d.Vendor.IsUnknown() {
		data.Vendor = mistutils.ListOfStringTerraformToSdk(d.Vendor)
	}

	return &data
}

func notMatchingTerraformToSdk(d NotMatchingValue) *models.NacRuleMatching {
	data := models.NacRuleMatching{}

	if !d.AuthType.IsNull() && !d.AuthType.IsUnknown() {
		data.AuthType = models.ToPointer(models.NacAuthTypeEnum(d.AuthType.ValueString()))
	}
	if !d.Family.IsNull() && !d.Family.IsUnknown() {
		data.Family = mistutils.ListOfStringTerraformToSdk(d.Family)
	}
	if !d.Mfg.IsNull() && !d.Mfg.IsUnknown() {
		data.Mfg = mistutils.ListOfStringTerraformToSdk(d.Mfg)
	}
	if !d.Model.IsNull() && !d.Model.IsUnknown() {
		data.Model = mistutils.ListOfStringTerraformToSdk(d.Model)
	}
	if !d.Nactags.IsNull() && !d.Nactags.IsUnknown() {
		data.Nactags = mistutils.ListOfStringTerraformToSdk(d.Nactags)
	}
	if !d.OsType.IsNull() && !d.OsType.IsUnknown() {
		data.OsType = mistutils.ListOfStringTerraformToSdk(d.OsType)
	}
	if !d.PortTypes.IsNull() && !d.PortTypes.IsUnknown() {
		data.PortTypes = matchingPortTypesTerraformToSdk(d.PortTypes)
	}
	if !d.SiteIds.IsNull() && !d.SiteIds.IsUnknown() {
		data.SiteIds = mistutils.ListOfUuidTerraformToSdk(d.SiteIds)
	}
	if !d.SitegroupIds.IsNull() && !d.SitegroupIds.IsUnknown() {
		data.SitegroupIds = mistutils.ListOfUuidTerraformToSdk(d.SitegroupIds)
	}
	if !d.Vendor.IsNull() && !d.Vendor.IsUnknown() {
		data.Vendor = mistutils.ListOfStringTerraformToSdk(d.Vendor)
	}

	return &data
}
