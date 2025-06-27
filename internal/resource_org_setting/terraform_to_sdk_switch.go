package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchAutoUpgradeTerraformToSdk(ctx context.Context, o basetypes.ObjectValue) *models.SwitchAutoUpgrade {
	data := models.SwitchAutoUpgrade{}
	if o.IsNull() || o.IsUnknown() {
		return &data
	} else {
		d := NewAutoUpgradeValueMust(o.AttributeTypes(ctx), o.Attributes())

		if !d.CustomVersions.IsNull() && !d.CustomVersions.IsUnknown() {
			var customVersions = make(map[string]string)
			for k, v := range d.CustomVersions.Elements() {
				var vInterface interface{} = v
				s := vInterface.(basetypes.StringValue)
				customVersions[k] = s.ValueString()
			}
			data.CustomVersions = customVersions
		}
		if d.Enabled.ValueBoolPointer() != nil {
			data.Enabled = d.Enabled.ValueBoolPointer()
		}
		if d.Snapshot.ValueBoolPointer() != nil {
			data.Snapshot = d.Snapshot.ValueBoolPointer()
		}

		return &data
	}
}

func switchTerraformToSdk(ctx context.Context, d SwitchValue) *models.OrgSettingSwitch {
	data := models.OrgSettingSwitch{}
	if !d.AutoUpgrade.IsNull() && !d.AutoUpgrade.IsUnknown() {
		data.AutoUpgrade = switchAutoUpgradeTerraformToSdk(ctx, d.AutoUpgrade)
	}
	return &data
}
