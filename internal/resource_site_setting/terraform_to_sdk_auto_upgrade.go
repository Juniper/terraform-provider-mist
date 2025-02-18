package resource_site_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func siteSettingAutoUpgradeTerraformToSdk(d AutoUpgradeValue) *models.SiteSettingAutoUpgrade {
	data := models.SiteSettingAutoUpgrade{}

	customVersions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		customVersions[k] = vd.ValueString()
	}
	data.CustomVersions = customVersions
	data.DayOfWeek = models.ToPointer(models.DayOfWeekEnum(d.DayOfWeek.ValueString()))
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	data.Version = models.ToPointer(models.SiteAutoUpgradeVersionEnum(d.Version.ValueString()))

	return &data
}
