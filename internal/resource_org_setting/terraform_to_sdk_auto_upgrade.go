package resource_org_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func autoUpgradeTerraformToSdk(d AutoUpgradeValue) *models.OrgSettingAutoUpgrade {
	data := models.OrgSettingAutoUpgrade{}

	if !d.CustomVersions.IsNull() && !d.CustomVersions.IsUnknown() {
		rMap := make(map[string]string)
		for k, v := range d.CustomVersions.Elements() {
			var vInterface interface{} = v
			vString := vInterface.(basetypes.StringValue)
			if vString.ValueStringPointer() != nil {
				rMap[k] = vString.ValueString()
			}
		}
		data.CustomVersions = rMap
	}
	if d.DayOfWeek.ValueStringPointer() != nil {
		dayOfWeek := models.DayOfWeekEnum(d.DayOfWeek.ValueString())
		data.DayOfWeek = &dayOfWeek
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if d.TimeOfDay.ValueStringPointer() != nil {
		data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	}
	if d.Version.ValueStringPointer() != nil {
		version := models.SiteAutoUpgradeVersionEnum(d.Version.ValueString())
		data.Version = &version
	}

	return &data
}
