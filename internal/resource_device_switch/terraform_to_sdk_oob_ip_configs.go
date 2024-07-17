package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) *models.SwitchOobIpConfig {

	data := models.SwitchOobIpConfig{}

	if d.Gateway.ValueStringPointer() != nil {
		data.Gateway = d.Gateway.ValueStringPointer()
	}
	if d.Ip.ValueStringPointer() != nil {
		data.Ip = d.Ip.ValueStringPointer()
	}
	if d.Netmask.ValueStringPointer() != nil {
		data.Netmask = d.Netmask.ValueStringPointer()
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = d.Network.ValueStringPointer()
	}
	if d.OobIpConfigType.ValueStringPointer() != nil {
		data.Type = (*models.IpTypeEnum)(d.OobIpConfigType.ValueStringPointer())
	}
	if d.UseMgmtVrf.ValueBoolPointer() != nil {
		data.UseMgmtVrf = d.UseMgmtVrf.ValueBoolPointer()
	}
	if d.UseMgmtVrfForHostOut.ValueBoolPointer() != nil {
		data.UseMgmtVrfForHostOut = d.UseMgmtVrfForHostOut.ValueBoolPointer()
	}

	return &data
}
