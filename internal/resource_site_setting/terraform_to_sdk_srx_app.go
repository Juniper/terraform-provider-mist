package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func srxAppTerraformToSdk(d SrxAppValue) *models.SiteSettingSrxApp {
	data := models.SiteSettingSrxApp{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
