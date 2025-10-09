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

func juniperSrxAutoUpgradeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.JuniperSrxAutoUpgrade) basetypes.ObjectValue {

	var customVersions = types.MapNull(types.StringType)
	var enabled basetypes.BoolValue
	var snapshot basetypes.BoolValue

	if d.CustomVersions != nil {
		rMapValue := make(map[string]attr.Value)
		for k, v := range d.CustomVersions {
			rMapValue[k] = types.StringValue(v)
		}
		m, e := types.MapValueFrom(ctx, types.StringType, rMapValue)
		diags.Append(e...)
		customVersions = m
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Snapshot != nil {
		snapshot = types.BoolValue(*d.Snapshot)
	}

	dataMapValue := map[string]attr.Value{
		"custom_versions": customVersions,
		"enabled":         enabled,
		"snapshot":        snapshot,
	}
	data, e := NewSrxAutoUpgradeValue(SrxAutoUpgradeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o

}

func juniperSrxSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingJuniperSrx) JuniperSrxValue {
	var autoUpgrade = types.ObjectNull(SrxAutoUpgradeValue{}.AttributeTypes(ctx))
	var gateways = types.ListNull(GatewaysValue{}.Type(ctx))
	var sendMistNacUserInfo basetypes.BoolValue

	if d != nil && d.AutoUpgrade != nil {
		autoUpgrade = juniperSrxAutoUpgradeSdkToTerraform(ctx, diags, *d.AutoUpgrade)
	}
	if d != nil && d.Gateways != nil {
		gateways = juniperSrxGatewaysSdkToTerraform(ctx, diags, d.Gateways)
	}
	if d != nil && d.SendMistNacUserInfo != nil {
		sendMistNacUserInfo = types.BoolValue(*d.SendMistNacUserInfo)
	}

	dataMapValue := map[string]attr.Value{
		"auto_upgrade":            autoUpgrade,
		"gateways":                gateways,
		"send_mist_nac_user_info": sendMistNacUserInfo,
	}
	data, e := NewJuniperSrxValue(JuniperSrxValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
