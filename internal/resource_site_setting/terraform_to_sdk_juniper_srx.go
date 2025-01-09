package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func juniperSrxGatewaysTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SiteSettingJuniperSrxGateway {
	var data_list []models.SiteSettingJuniperSrxGateway
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(GatewaysValue)
		data := models.SiteSettingJuniperSrxGateway{}

		if plan.ApiKey.ValueStringPointer() != nil {
			data.ApiKey = plan.ApiKey.ValueStringPointer()
		}

		if plan.ApiUrl.ValueStringPointer() != nil {
			data.ApiUrl = plan.ApiUrl.ValueStringPointer()
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func juniperSrxTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d JuniperSrxValue) *models.SiteSettingJuniperSrx {
	data := models.SiteSettingJuniperSrx{}
	if !d.Gateways.IsNull() && !d.Gateways.IsUnknown() {
		data.Gateways = juniperSrxGatewaysTerraformToSdk(ctx, diags, d.Gateways)
	}
	if d.SendMistNacUserInfo.ValueBoolPointer() != nil {
		data.SendMistNacUserInfo = d.SendMistNacUserInfo.ValueBoolPointer()
	}
	return &data
}
