package datasource_device_ap_stats

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ipConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApIpConfig) basetypes.ObjectValue {
	tflog.Debug(ctx, "ipConfigSdkToTerraform")
	var dns basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var dns_suffix basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var ip basetypes.StringValue
	var ip6 basetypes.StringValue
	var mtu basetypes.Int64Value
	var netmask basetypes.StringValue
	var netmask6 basetypes.StringValue
	var type4 basetypes.StringValue
	var type6 basetypes.StringValue
	var vlan_id basetypes.Int64Value

	if d.Dns != nil {
		dns = mist_transform.ListOfStringSdkToTerraform(ctx, d.Dns)
	}
	if d.DnsSuffix != nil {
		dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, d.DnsSuffix)
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
		vlan_id = types.Int64Value(int64(*d.VlanId))
	}

	data_map_attr_type := IpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"dns":        dns,
		"dns_suffix": dns_suffix,
		"gateway":    gateway,
		"gateway6":   gateway6,
		"ip":         ip,
		"ip6":        ip6,
		"mtu":        mtu,
		"netmask":    netmask,
		"netmask6":   netmask6,
		"type":       type4,
		"type6":      type6,
		"vlan_id":    vlan_id,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
