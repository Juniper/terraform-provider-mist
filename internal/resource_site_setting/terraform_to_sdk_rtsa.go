package resource_site_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func rtsaTerraformToSdk(d RtsaValue) *models.SiteSettingRtsa {
	data := models.SiteSettingRtsa{}

	if d.AppWaking.ValueBoolPointer() != nil {
		data.AppWaking = d.AppWaking.ValueBoolPointer()
	}
	if d.DisableDeadReckoning.ValueBoolPointer() != nil {
		data.DisableDeadReckoning = d.DisableDeadReckoning.ValueBoolPointer()
	}
	if d.DisablePressureSensor.ValueBoolPointer() != nil {
		data.DisablePressureSensor = d.DisablePressureSensor.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.TrackAsset.ValueBoolPointer() != nil {
		data.TrackAsset = d.TrackAsset.ValueBoolPointer()
	}

	return &data
}
