package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func siteSettingAutoUpgradeEslTerraformToSdk(d AutoUpgradeEslValue) *models.SiteSettingAutoUpgradeEsl {
	data := models.SiteSettingAutoUpgradeEsl{}

	data.AllowDowngrade = d.AllowDowngrade.ValueBoolPointer()

	customVersions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		customVersions[k] = vd.ValueString()
	}
	if len(customVersions) > 0 {
		data.CustomVersions = customVersions
	}

	if d.DayOfWeek.ValueStringPointer() != nil {
		data.DayOfWeek = models.ToPointer(models.DayOfWeekEnum(d.DayOfWeek.ValueString()))
	}
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	data.Version = d.Version.ValueStringPointer()

	return &data
}
