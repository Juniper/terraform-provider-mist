package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func juniperSrxGatewaysTerraformToSdk(d basetypes.ListValue) []models.SiteSettingJuniperSrxGateway {
	var dataList []models.SiteSettingJuniperSrxGateway
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(GatewaysValue)
		data := models.SiteSettingJuniperSrxGateway{}

		if plan.ApiKey.ValueStringPointer() != nil {
			data.ApiKey = plan.ApiKey.ValueStringPointer()
		}

		if plan.ApiUrl.ValueStringPointer() != nil {
			data.ApiUrl = plan.ApiUrl.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func juniperSrxTerraformToSdk(d JuniperSrxValue) *models.SiteSettingJuniperSrx {
	data := models.SiteSettingJuniperSrx{}
	if !d.Gateways.IsNull() && !d.Gateways.IsUnknown() {
		data.Gateways = juniperSrxGatewaysTerraformToSdk(d.Gateways)
	}
	if d.SendMistNacUserInfo.ValueBoolPointer() != nil {
		data.SendMistNacUserInfo = d.SendMistNacUserInfo.ValueBoolPointer()
	}
	return &data
}
