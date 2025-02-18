package resource_org_network

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func InternalAccessTerraformToSdk(d InternalAccessValue) *models.NetworkInternalAccess {
	data := models.NetworkInternalAccess{}
	data.Enabled = d.Enabled.ValueBoolPointer()
	return &data
}
