package datasource_device_ap_stats

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ipStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IpStat) basetypes.ObjectValue {

	var dhcpServer basetypes.StringValue
	var dns = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var gateway basetypes.StringValue
	var gateway6 basetypes.StringValue
	var ip basetypes.StringValue
	var ip6 basetypes.StringValue
	var ips = types.MapNull(types.StringType)
	var netmask basetypes.StringValue
	var netmask6 basetypes.StringValue

	if d.DhcpServer.Value() != nil {
		dhcpServer = types.StringValue(*d.DhcpServer.Value())
	}
	if d.Dns != nil {
		dns = mistutils.ListOfStringSdkToTerraform(d.Dns)
	}
	if d.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(d.DnsSuffix)
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
	if len(d.Ips) > 0 {
		mapAttrValues := make(map[string]attr.Value)
		for k, v := range d.Ips {
			mapAttrValues[k] = types.StringValue(v)
		}
		mapAttr, e := types.MapValueFrom(ctx, types.StringType, mapAttrValues)
		diags.Append(e...)
		ips = mapAttr
	}
	if d.Netmask.Value() != nil {
		netmask = types.StringValue(*d.Netmask.Value())
	}
	if d.Netmask6.Value() != nil {
		netmask6 = types.StringValue(*d.Netmask6.Value())
	}

	dataMapValue := map[string]attr.Value{
		"dhcp_server": dhcpServer,
		"dns":         dns,
		"dns_suffix":  dnsSuffix,
		"gateway":     gateway,
		"gateway6":    gateway6,
		"ip":          ip,
		"ip6":         ip6,
		"ips":         ips,
		"netmask":     netmask,
		"netmask6":    netmask6,
	}
	data, e := basetypes.NewObjectValue(IpStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
