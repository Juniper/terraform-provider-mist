package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func evpnConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EvpnConfigValue) *models.EvpnConfig {
	data := models.EvpnConfig{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Role.ValueStringPointer() != nil {
		data.Role = (*models.EvpnConfigRoleEnum)(d.Role.ValueStringPointer())
	}
	return &data
}
