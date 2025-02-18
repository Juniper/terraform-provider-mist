package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wanVnaTerraformToSdk(d WanVnaValue) *models.SiteSettingWanVna {
	data := models.SiteSettingWanVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
