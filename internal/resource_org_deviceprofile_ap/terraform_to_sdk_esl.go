package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func eslTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EslConfigValue) *models.ApEslConfig {
	data := models.ApEslConfig{}

	if d.Cacert.ValueStringPointer() != nil {
		data.Cacert = d.Cacert.ValueStringPointer()
	}
	if d.Channel.ValueInt64Pointer() != nil {
		data.Channel = models.ToPointer(int(d.Channel.ValueInt64()))
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Host.ValueStringPointer() != nil {
		data.Host = d.Host.ValueStringPointer()
	}
	if d.Port.ValueInt64Pointer() != nil {
		data.Port = models.ToPointer(int(d.Port.ValueInt64()))
	}
	if d.EslConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.ApEslTypeEnum(d.EslConfigType.ValueString()))
	}
	if d.VerifyCert.ValueBoolPointer() != nil {
		data.VerifyCert = d.VerifyCert.ValueBoolPointer()
	}
	if d.VlanId.ValueInt64Pointer() != nil {
		data.VlanId = models.ToPointer(int(d.VlanId.ValueInt64()))
	}

	return &data
}
