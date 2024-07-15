package datasource_device_ap_stats

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ipStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IpStat) basetypes.ObjectValue {

	var dhcp_server basetypes.StringValue
	var dns basetypes.ListValue = types.ListNull(types.StringType)
	var dns_suffix basetypes.ListValue = types.ListNull(types.StringType)
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var ip basetypes.StringValue
	var ip6 basetypes.StringValue
	var ips basetypes.MapValue = types.MapNull(types.StringType)
	var netmask basetypes.StringValue
	var netmask6 basetypes.StringValue

	if d.DhcpServer.Value() != nil {
		dhcp_server = types.StringValue(*d.DhcpServer.Value())
	}
	if d.Dns != nil {
		dns = mist_transform.ListOfStringSdkToTerraform(ctx, d.Dns)
	}
	if d.DnsSuffix != nil {
		dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, d.DnsSuffix)
	}
	if d.Gateway.Value() != nil {
		gateway = types.StringValue(*d.Gateway.Value())
	}
	if d.Gateway6.Value() != nil {
		gateway6 = types.StringValue(*d.Gateway6.Value())
	}
	if d.Ip.Value() != nil {
		ip = types.StringValue(*d.Ip.Value())
	}
	if d.Ip6.Value() != nil {
		ip6 = types.StringValue(*d.Ip6.Value())
	}
	if d.Ips != nil {
		map_attr_values := make(map[string]attr.Value)
		for k, v := range d.Ips {
			map_attr_values[k] = types.StringValue(v)
		}
		map_attr, e := types.MapValueFrom(ctx, types.StringType, map_attr_values)
		diags.Append(e...)
		ips = map_attr
	}
	if d.Netmask.Value() != nil {
		netmask = types.StringValue(*d.Netmask.Value())
	}
	if d.Netmask6.Value() != nil {
		netmask6 = types.StringValue(*d.Netmask6.Value())
	}

	data_map_attr_type := IpStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"dhcp_server": dhcp_server,
		"dns":         dns,
		"dns_suffix":  dns_suffix,
		"gateway":     gateway,
		"gateway6":    gateway6,
		"ip":          ip,
		"ip6":         ip6,
		"ips":         ips,
		"netmask":     netmask,
		"netmask6":    netmask6,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
