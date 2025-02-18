package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vnaTerraformToSdk(d VnaValue) *models.SiteSettingVna {
	data := models.SiteSettingVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
