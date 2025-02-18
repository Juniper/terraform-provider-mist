package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func juniperSrxGatewaysSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SiteSettingJuniperSrxGateway) basetypes.ListValue {
	var dataList []GatewaysValue
	for _, d := range l {
		var apiKey basetypes.StringValue
		var apiUrl basetypes.StringValue

		if d.ApiKey != nil {
			apiKey = types.StringValue(*d.ApiKey)
		}
		if d.ApiUrl != nil {
			apiUrl = types.StringValue(*d.ApiUrl)
		}

		dataMapValue := map[string]attr.Value{
			"api_key": apiKey,
			"api_url": apiUrl,
		}
		data, e := NewGatewaysValue(GatewaysValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, GatewaysValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func juniperSrxSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingJuniperSrx) JuniperSrxValue {

	var gateways = types.ListNull(GatewaysValue{}.Type(ctx))
	var sendMistNacUserInfo basetypes.BoolValue

	if d != nil && d.Gateways != nil {
		gateways = juniperSrxGatewaysSdkToTerraform(ctx, diags, d.Gateways)
	}
	if d != nil && d.SendMistNacUserInfo != nil {
		sendMistNacUserInfo = types.BoolValue(*d.SendMistNacUserInfo)
	}

	dataMapValue := map[string]attr.Value{
		"gateways":                gateways,
		"send_mist_nac_user_info": sendMistNacUserInfo,
	}
	data, e := NewJuniperSrxValue(JuniperSrxValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
