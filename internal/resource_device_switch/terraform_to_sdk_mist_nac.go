package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mistNacTerraformToSdk(d MistNacValue) *models.SwitchMistNac {
	data := models.SwitchMistNac{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	return &data
}
