package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func oobIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OobIpConfigValue) *models.MxedgeOobIpConfig {
	data := models.MxedgeOobIpConfig{}

	if !d.Autoconf6.IsNull() && !d.Autoconf6.IsUnknown() {
		data.Autoconf6 = d.Autoconf6.ValueBoolPointer()
	}

	if !d.Dhcp6.IsNull() && !d.Dhcp6.IsUnknown() {
		data.Dhcp6 = d.Dhcp6.ValueBoolPointer()
	}

	if !d.Dns.IsNull() && !d.Dns.IsUnknown() {
		data.Dns = mistutils.ListOfStringTerraformToSdk(d.Dns)
	}

	if !d.Gateway.IsNull() && !d.Gateway.IsUnknown() {
		data.Gateway = d.Gateway.ValueStringPointer()
	}

	if !d.Gateway6.IsNull() && !d.Gateway6.IsUnknown() {
		data.Gateway6 = d.Gateway6.ValueStringPointer()
	}

	if !d.Ip.IsNull() && !d.Ip.IsUnknown() {
		data.Ip = d.Ip.ValueStringPointer()
	}

	if !d.Ip6.IsNull() && !d.Ip6.IsUnknown() {
		data.Ip6 = d.Ip6.ValueStringPointer()
	}

	if !d.Netmask.IsNull() && !d.Netmask.IsUnknown() {
		data.Netmask = d.Netmask.ValueStringPointer()
	}

	if !d.Netmask6.IsNull() && !d.Netmask6.IsUnknown() {
		data.Netmask6 = d.Netmask6.ValueStringPointer()
	}

	if !d.OobIpConfigType.IsNull() && !d.OobIpConfigType.IsUnknown() {
		data.Type = (*models.IpTypeEnum)(d.OobIpConfigType.ValueStringPointer())
	}

	if !d.Type6.IsNull() && !d.Type6.IsUnknown() {
		data.Type6 = (*models.IpTypeEnum)(d.Type6.ValueStringPointer())
	}

	return &data
}
