package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func stpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d StpConfigValue) *models.SwitchStpConfig {

	data := models.SwitchStpConfig{}

	if d.VstpEnabled.ValueBoolPointer() != nil {
		data.VstpEnabled = d.VstpEnabled.ValueBoolPointer()
	}

	return &data
}
