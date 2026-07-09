package resource_org_mxtunnel

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ipsecTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IpsecValue) *models.MxtunnelIpsec {
	data := models.MxtunnelIpsec{}

	if !d.DnsServers.IsNull() && !d.DnsServers.IsUnknown() {
		var items []string
		for _, v := range d.DnsServers.Elements() {
			s, ok := v.(interface{ ValueString() string })
			if ok {
				items = append(items, s.ValueString())
			}
		}
		data.DnsServers = models.NewOptional(&items)
	} else {
		data.DnsServers = models.NewOptional[[]string](nil)
	}

	if !d.DnsSuffix.IsNull() && !d.DnsSuffix.IsUnknown() {
		var items []string
		for _, v := range d.DnsSuffix.Elements() {
			s, ok := v.(interface{ ValueString() string })
			if ok {
				items = append(items, s.ValueString())
			}
		}
		data.DnsSuffix = items
	}

	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	if !d.ExtraRoutes.IsNull() && !d.ExtraRoutes.IsUnknown() {
		var routes []models.MxtunnelIpsecExtraRoute
		for _, item := range d.ExtraRoutes.Elements() {
			itemValue, ok := item.(ExtraRoutesValue)
			if !ok {
				continue
			}
			route := models.MxtunnelIpsecExtraRoute{}
			if !itemValue.Dest.IsNull() && !itemValue.Dest.IsUnknown() {
				route.Dest = itemValue.Dest.ValueStringPointer()
			}
			if !itemValue.NextHop.IsNull() && !itemValue.NextHop.IsUnknown() {
				route.NextHop = itemValue.NextHop.ValueStringPointer()
			}
			routes = append(routes, route)
		}
		data.ExtraRoutes = routes
	}

	if !d.SplitTunnel.IsNull() && !d.SplitTunnel.IsUnknown() {
		data.SplitTunnel = d.SplitTunnel.ValueBoolPointer()
	}

	if !d.UseMxedge.IsNull() && !d.UseMxedge.IsUnknown() {
		data.UseMxedge = d.UseMxedge.ValueBoolPointer()
	}

	_ = ctx
	_ = diags
	return &data
}
