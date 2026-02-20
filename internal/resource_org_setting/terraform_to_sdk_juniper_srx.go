package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

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

func juniperSrxTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d JuniperSrxValue) *models.OrgSettingJuniperSrx {
	data := models.OrgSettingJuniperSrx{}

	if !d.SrxAutoUpgrade.IsNull() && !d.SrxAutoUpgrade.IsUnknown() {
		data.AutoUpgrade = models.ToPointer(juniperSrxAutoUpgradeTerraformToSdk(ctx, diags, d.SrxAutoUpgrade))
	}

	return &data
}
