package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func stpConfigTerraformToSdk(d StpConfigValue) *models.SwitchStpConfig {

	data := models.SwitchStpConfig{}

	if d.BridgePriority.ValueStringPointer() != nil {
		data.BridgePriority = d.BridgePriority.ValueStringPointer()
	}

	return &data
}
