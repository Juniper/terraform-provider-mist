package resource_device_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func centrakTerraformToSdk(d CentrakValue) *models.ApCentrak {
	data := models.ApCentrak{}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
