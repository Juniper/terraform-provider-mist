package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func uplinkPortConfigTerraformToSdk(d UplinkPortConfigValue) *models.ApUplinkPortConfig {
	data := models.ApUplinkPortConfig{}

	if d.Dot1x.ValueBoolPointer() != nil {
		data.Dot1x = d.Dot1x.ValueBoolPointer()
	}

	if d.KeepWlansUpIfDown.ValueBoolPointer() != nil {
		data.KeepWlansUpIfDown = d.Dot1x.ValueBoolPointer()
	}

	return &data
}
