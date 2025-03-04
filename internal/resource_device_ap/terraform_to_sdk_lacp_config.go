package resource_device_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func lacpConfigTerraformToSdk(d LacpConfigValue) *models.DeviceApLacpConfig {

	var data models.DeviceApLacpConfig

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
