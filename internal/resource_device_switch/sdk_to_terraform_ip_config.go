package resource_device_switch

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.JunosIpConfig) IpConfigValue {
	var dns = mistutils.ListOfStringSdkToTerraformEmpty()
	var dnsSuffix = mistutils.ListOfStringSdkToTerraformEmpty()
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var network basetypes.StringValue
	var type4 basetypes.StringValue

	if d.Dns != nil {
		dns = mistutils.ListOfStringSdkToTerraform(d.Dns)
	}
	if d.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(d.DnsSuffix)
	}
	if d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d.Type != nil {
		type4 = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"dns":        dns,
		"dns_suffix": dnsSuffix,
		"gateway":    gateway,
		"ip":         ip,
		"netmask":    netmask,
		"network":    network,
		"type":       type4,
	}
	data, e := NewIpConfigValue(IpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
