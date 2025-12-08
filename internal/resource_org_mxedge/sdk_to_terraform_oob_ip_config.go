package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func oobIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeOobIpConfig) OobIpConfigValue {

	var autoconf6 types.Bool
	var dhcp6 types.Bool
	var dns = types.ListNull(types.StringType)
	var gateway types.String
	var gateway6 types.String
	var ip types.String
	var ip6 types.String
	var netmask types.String
	var netmask6 types.String
	var oobIpConfigType types.String
	var type6 types.String

	if d.Autoconf6 != nil {
		autoconf6 = types.BoolValue(*d.Autoconf6)
	}
	if d.Dhcp6 != nil {
		dhcp6 = types.BoolValue(*d.Dhcp6)
	}
	if d.Dns != nil {
		dns = mistutils.ListOfStringSdkToTerraform(d.Dns)
	}
	if d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d.Gateway6 != nil {
		gateway6 = types.StringValue(*d.Gateway6)
	}
	if d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d.Ip6 != nil {
		ip6 = types.StringValue(*d.Ip6)
	}
	if d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d.Netmask6 != nil {
		netmask6 = types.StringValue(*d.Netmask6)
	}
	if d.Type != nil {
		oobIpConfigType = types.StringValue(string(*d.Type))
	}
	if d.Type6 != nil {
		type6 = types.StringValue(string(*d.Type6))
	}

	data_map_attr_type := OobIpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"autoconf6": autoconf6,
		"dhcp6":     dhcp6,
		"dns":       dns,
		"gateway":   gateway,
		"gateway6":  gateway6,
		"ip":        ip,
		"ip6":       ip6,
		"netmask":   netmask,
		"netmask6":  netmask6,
		"type":      oobIpConfigType,
		"type6":     type6,
	}
	data, e := NewOobIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
