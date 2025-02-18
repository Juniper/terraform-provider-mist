package resource_device_ap

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApIpConfig) IpConfigValue {
	var dns = misttransform.ListOfStringSdkToTerraformEmpty()
	var dnsSuffix = misttransform.ListOfStringSdkToTerraformEmpty()
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var ip basetypes.StringValue
	var ip6 basetypes.StringValue
	var mtu basetypes.Int64Value
	var netmask basetypes.StringValue
	var netmask6 basetypes.StringValue
	var type4 basetypes.StringValue
	var type6 basetypes.StringValue
	var vlanId basetypes.Int64Value

	if d.Dns != nil {
		dns = misttransform.ListOfStringSdkToTerraform(d.Dns)
	}
	if d.DnsSuffix != nil {
		dnsSuffix = misttransform.ListOfStringSdkToTerraform(d.DnsSuffix)
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
	if d.Mtu != nil {
		mtu = types.Int64Value(int64(*d.Mtu))
	}
	if d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d.Netmask6 != nil {
		netmask6 = types.StringValue(*d.Netmask6)
	}
	if d.Type != nil {
		type4 = types.StringValue(string(*d.Type))
	}
	if d.Type6 != nil {
		type6 = types.StringValue(string(*d.Type6))
	}
	if d.VlanId != nil {
		vlanId = types.Int64Value(int64(*d.VlanId))
	}

	dataMapValue := map[string]attr.Value{
		"dns":        dns,
		"dns_suffix": dnsSuffix,
		"gateway":    gateway,
		"gateway6":   gateway6,
		"ip":         ip,
		"ip6":        ip6,
		"mtu":        mtu,
		"netmask":    netmask,
		"netmask6":   netmask6,
		"type":       type4,
		"type6":      type6,
		"vlan_id":    vlanId,
	}
	data, e := NewIpConfigValue(IpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
