package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mistNacTerraformToSdk(d MistNacValue) *models.WlanMistNac {
	data := models.WlanMistNac{}
	data.Enabled = d.Enabled.ValueBoolPointer()
	return &data
}
