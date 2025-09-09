package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func siteSettingAutoUpgradeEslTerraformToSdk(d AutoUpgradeEslValue) *models.SiteSettingAutoUpgradeEsl {
	data := models.SiteSettingAutoUpgradeEsl{}

	if d.AllowDowngrade.ValueBoolPointer() != nil {
		data.AllowDowngrade = d.AllowDowngrade.ValueBoolPointer()
	}

	if !d.CustomVersions.IsNull() || !d.CustomVersions.IsUnknown() {
		customVersions := make(map[string]string)
		for k, v := range d.CustomVersions.Elements() {
			var vi interface{} = v
			vd := vi.(basetypes.StringValue)
			customVersions[k] = vd.ValueString()
		}
		data.CustomVersions = customVersions
	}

	if d.DayOfWeek.ValueStringPointer() != nil {
		data.DayOfWeek = models.ToPointer(models.DayOfWeekEnum(d.DayOfWeek.ValueString()))
	}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if d.TimeOfDay.ValueStringPointer() != nil {
		data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	}

	if d.Version.ValueStringPointer() != nil {
		data.Version = d.Version.ValueStringPointer()
	}

	return &data
}
