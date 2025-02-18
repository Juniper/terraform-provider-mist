package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func skyAtpTerraformToSdk(d SkyatpValue) *models.SiteSettingSkyatp {
	data := models.SiteSettingSkyatp{}

	data.Enabled = d.Enabled.ValueBoolPointer()
	data.SendIpMacMapping = d.SendIpMacMapping.ValueBoolPointer()

	return &data
}
