package resource_device_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func airistaTerraformToSdk(d AiristaValue) *models.ApAirista {
	data := models.ApAirista{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = models.NewOptional(d.Host.ValueStringPointer())
	}
	if d.Port.ValueInt64Pointer() != nil {
		data.Port = models.NewOptional(models.ToPointer(int(d.Port.ValueInt64())))
	}

	return &data
}
