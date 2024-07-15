package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func injectDhcpOption82TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan InjectDhcpOption82Value) *models.WlanInjectDhcpOption82 {

	data := models.WlanInjectDhcpOption82{}
	data.CircuitId = plan.CircuitId.ValueStringPointer()
	data.Enabled = plan.Enabled.ValueBoolPointer()

	return &data
}
