package resource_org_mxtunnel

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ipsecSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxtunnelIpsec) IpsecValue {
	var dnsServers = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var enabled = types.BoolNull()
	var extraRoutes = types.ListNull(ExtraRoutesValue{}.Type(ctx))
	var splitTunnel = types.BoolNull()
	var useMxedge = types.BoolNull()

	if d.DnsServers.Value() != nil {
		var items []attr.Value
		for _, v := range *d.DnsServers.Value() {
			items = append(items, types.StringValue(v))
		}
		r, e := types.ListValue(types.StringType, items)
		diags.Append(e...)
		dnsServers = r
	}

	if d.DnsSuffix != nil {
		var items []attr.Value
		for _, v := range d.DnsSuffix {
			items = append(items, types.StringValue(v))
		}
		r, e := types.ListValue(types.StringType, items)
		diags.Append(e...)
		dnsSuffix = r
	}

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	if d.ExtraRoutes != nil {
		var dataList []attr.Value
		for _, item := range d.ExtraRoutes {
			var dest = types.StringNull()
			var nextHop = types.StringNull()

			if item.Dest != nil {
				dest = types.StringValue(*item.Dest)
			}
			if item.NextHop != nil {
				nextHop = types.StringValue(*item.NextHop)
			}

			itemAttrType := ExtraRoutesValue{}.AttributeTypes(ctx)
			itemValue := map[string]attr.Value{
				"dest":     dest,
				"next_hop": nextHop,
			}
			itemObj, e := NewExtraRoutesValue(itemAttrType, itemValue)
			diags.Append(e...)
			dataList = append(dataList, itemObj)
		}
		r, e := types.ListValueFrom(ctx, ExtraRoutesValue{}.Type(ctx), dataList)
		diags.Append(e...)
		extraRoutes = r
	}

	if d.SplitTunnel != nil {
		splitTunnel = types.BoolValue(*d.SplitTunnel)
	}

	if d.UseMxedge != nil {
		useMxedge = types.BoolValue(*d.UseMxedge)
	}

	dataMapAttrType := IpsecValue{}.AttributeTypes(ctx)
	dataMapValue := map[string]attr.Value{
		"dns_servers":  dnsServers,
		"dns_suffix":   dnsSuffix,
		"enabled":      enabled,
		"extra_routes": extraRoutes,
		"split_tunnel": splitTunnel,
		"use_mxedge":   useMxedge,
	}
	data, e := NewIpsecValue(dataMapAttrType, dataMapValue)
	diags.Append(e...)

	return data
}
