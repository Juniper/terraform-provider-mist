package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func injectDhcpOption82TerraformToSdk(plan InjectDhcpOption82Value) *models.WlanInjectDhcpOption82 {

	data := models.WlanInjectDhcpOption82{}
	data.CircuitId = plan.CircuitId.ValueStringPointer()
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
