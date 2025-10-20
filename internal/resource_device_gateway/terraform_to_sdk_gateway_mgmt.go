package resource_device_gateway

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func gatewayMgmtTerraformToSdk(d GatewayMgmtValue) *models.GatewayMgmt {
	data := models.GatewayMgmt{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		if d.ConfigRevertTimer.ValueInt64Pointer() != nil {
			data.ConfigRevertTimer = models.ToPointer(int(d.ConfigRevertTimer.ValueInt64()))
		}
		return &data
	}
}
