package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wiredVnaTerraformToSdk(d WiredVnaValue) *models.SiteSettingWiredVna {
	data := models.SiteSettingWiredVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
