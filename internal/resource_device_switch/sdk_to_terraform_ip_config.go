package resource_device_switch

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.JunosIpConfig) IpConfigValue {
	var dns basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var dns_suffix basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var network basetypes.StringValue
	var type4 basetypes.StringValue

	if d.Dns != nil {
		dns = mist_transform.ListOfStringSdkToTerraform(ctx, d.Dns)
	}
	if d.DnsSuffix != nil {
		dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, d.DnsSuffix)
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

	data_map_attr_type := IpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"dns":        dns,
		"dns_suffix": dns_suffix,
		"gateway":    gateway,
		"ip":         ip,
		"netmask":    netmask,
		"network":    network,
		"type":       type4,
	}
	data, e := NewIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
