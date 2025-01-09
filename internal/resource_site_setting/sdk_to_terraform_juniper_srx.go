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
	var data_list = []GatewaysValue{}
	for _, d := range l {
		var api_key basetypes.StringValue
		var api_url basetypes.StringValue

		if d.ApiKey != nil {
			api_key = types.StringValue(*d.ApiKey)
		}
		if d.ApiUrl != nil {
			api_url = types.StringValue(*d.ApiUrl)
		}

		data_map_value := map[string]attr.Value{
			"api_key": api_key,
			"api_url": api_url,
		}
		data, e := NewGatewaysValue(GatewaysValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, GatewaysValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func juniperSrxSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingJuniperSrx) JuniperSrxValue {

	var gateways basetypes.ListValue = types.ListNull(GatewaysValue{}.Type(ctx))
	var send_mist_nac_user_info basetypes.BoolValue

	if d != nil && d.Gateways != nil {
		gateways = juniperSrxGatewaysSdkToTerraform(ctx, diags, d.Gateways)
	}
	if d != nil && d.SendMistNacUserInfo != nil {
		send_mist_nac_user_info = types.BoolValue(*d.SendMistNacUserInfo)
	}

	data_map_value := map[string]attr.Value{
		"gateways":                gateways,
		"send_mist_nac_user_info": send_mist_nac_user_info,
	}
	data, e := NewJuniperSrxValue(JuniperSrxValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
