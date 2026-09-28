package resource_device_gateway

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mnhaConfigTerraformToSdk(d MnhaConfigValue) *models.GatewayMnhaConfig {
	data := models.GatewayMnhaConfig{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}

	return &data
}
