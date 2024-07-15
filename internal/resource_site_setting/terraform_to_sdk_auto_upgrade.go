package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func siteSettingAutoUpgradeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AutoUpgradeValue) *models.SiteSettingAutoUpgrade {
	tflog.Debug(ctx, "siteSettingAutoUpgradeTerraformToSdk")
	data := models.SiteSettingAutoUpgrade{}

	custom_versions := make(map[string]string)
	for k, v := range d.CustomVersions.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		custom_versions[k] = vd.ValueString()
	}
	data.CustomVersions = custom_versions
	data.DayOfWeek = models.ToPointer(models.DayOfWeekEnum(d.DayOfWeek.ValueString()))
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	data.Version = models.ToPointer(models.SiteAutoUpgradeVersionEnum(d.Version.ValueString()))

	return &data
}
