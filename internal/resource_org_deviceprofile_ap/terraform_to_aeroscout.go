package resource_org_deviceprofile_ap

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func aeroscoutTerraformToSdk(d AeroscoutValue) *models.ApAeroscout {
	data := models.ApAeroscout{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = models.NewOptional(d.Host.ValueStringPointer())
	}
	if d.LocateConnected.ValueBoolPointer() != nil {
		data.LocateConnected = d.LocateConnected.ValueBoolPointer()
	}

	return &data
}
