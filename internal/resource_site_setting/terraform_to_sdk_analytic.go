package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func analyticTerraformToSdk(d AnalyticValue) *models.SiteSettingAnalytic {
	data := models.SiteSettingAnalytic{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
