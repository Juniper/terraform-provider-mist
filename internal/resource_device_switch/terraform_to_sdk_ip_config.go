package resource_device_switch

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigTerraformToSdk(d IpConfigValue) *models.JunosIpConfig {
	data := models.JunosIpConfig{}

	if !d.Dns.IsNull() && !d.Dns.IsUnknown() {
		data.Dns = mistutils.ListOfStringTerraformToSdk(d.Dns)
	}
	if !d.DnsSuffix.IsNull() && !d.DnsSuffix.IsUnknown() {
		data.DnsSuffix = mistutils.ListOfStringTerraformToSdk(d.DnsSuffix)
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
