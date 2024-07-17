package resource_device_switch

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IpConfigValue) *models.JunosIpConfig {
	data := models.JunosIpConfig{}

	if !d.Dns.IsNull() && !d.Dns.IsUnknown() {
		data.Dns = mist_transform.ListOfStringTerraformToSdk(ctx, d.Dns)
	}
	if !d.DnsSuffix.IsNull() && !d.DnsSuffix.IsUnknown() {
		data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, d.DnsSuffix)
	}
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
	if d.IpConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.IpTypeEnum(d.IpConfigType.ValueString()))
	}

	return &data
}
