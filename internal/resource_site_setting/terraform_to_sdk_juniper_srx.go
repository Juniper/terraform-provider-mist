package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
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

func juniperSrxAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) models.JuniperSrxAutoUpgrade {
	data := models.JuniperSrxAutoUpgrade{}
	if !d.IsNull() && !d.IsUnknown() {
		item, e := NewSrxAutoUpgradeValue(SrxAutoUpgradeValue{}.AttributeTypes(ctx), d.Attributes())
		diags.Append(e...)
		if e != nil {
			diags.Append(e...)
			return data
		} else {
			if !item.CustomVersions.IsNull() && !item.CustomVersions.IsUnknown() {
				rMap := make(map[string]string)
				for k, v := range item.CustomVersions.Elements() {
					var vInterface interface{} = v
					vString := vInterface.(basetypes.StringValue)
					if vString.ValueStringPointer() != nil {
						rMap[k] = vString.ValueString()
					}
				}
				data.CustomVersions = rMap
			}
			if item.Enabled.ValueBoolPointer() != nil {
				data.Enabled = models.ToPointer(item.Enabled.ValueBool())
			}
			if item.Snapshot.ValueBoolPointer() != nil {
				data.Snapshot = models.ToPointer(item.Snapshot.ValueBool())
			}
			if item.Version.ValueStringPointer() != nil {
				data.Version = models.ToPointer(item.Version.ValueString())
			}
		}
	}
	return data
}

func juniperSrxTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d JuniperSrxValue) *models.SiteSettingJuniperSrx {
	data := models.SiteSettingJuniperSrx{}
	if !d.SrxAutoUpgrade.IsNull() && !d.SrxAutoUpgrade.IsUnknown() {
		data.AutoUpgrade = models.ToPointer(juniperSrxAutoUpgradeTerraformToSdk(ctx, diags, d.SrxAutoUpgrade))
	}
	if !d.Gateways.IsNull() && !d.Gateways.IsUnknown() {
		data.Gateways = juniperSrxGatewaysTerraformToSdk(d.Gateways)
	}
	if d.SendMistNacUserInfo.ValueBoolPointer() != nil {
		data.SendMistNacUserInfo = d.SendMistNacUserInfo.ValueBoolPointer()
	}
	return &data
}
