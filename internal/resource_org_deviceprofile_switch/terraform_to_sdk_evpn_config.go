package resource_org_deviceprofile_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func evpnConfigTerraformToSdk(d EvpnConfigValue) *models.EvpnConfig {
	data := models.EvpnConfig{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Role.ValueStringPointer() != nil {
		data.Role = models.ToPointer(models.EvpnConfigRoleEnum(d.Role.ValueString()))
	}

	return &data
}
