package resource_org_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func cacertsConfigsTerraformToSdk(d basetypes.ListValue) []models.OrgSettingCacertsConfig {
	var dataList []models.OrgSettingCacertsConfig
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(CacertsConfigsValue)
		data := models.OrgSettingCacertsConfig{}

		data.Cert = plan.Cert.ValueString()
		if plan.CrlEnabled.ValueBoolPointer() != nil {
			data.CrlEnabled = plan.CrlEnabled.ValueBoolPointer()
		}
		if plan.CrlUrl.ValueStringPointer() != nil {
			data.CrlUrl = plan.CrlUrl.ValueStringPointer()
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if plan.OcspEnabled.ValueBoolPointer() != nil {
			data.OcspEnabled = plan.OcspEnabled.ValueBoolPointer()
		}
		if plan.OcspUrl.ValueStringPointer() != nil {
			data.OcspUrl = plan.OcspUrl.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}
	return dataList
}
