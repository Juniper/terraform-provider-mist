package datasource_device_gateway_stats

import (
	"context"
	"github.com/tmunzer/mistapi-go/mistapi/models"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vpnPeersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsGatewayVpnPeer) basetypes.ListValue {

	var dataList []VpnPeersValue
	for _, d := range l {
		var isActive basetypes.BoolValue
		var lastSeen basetypes.Float64Value
		var latency basetypes.NumberValue
		var mos basetypes.NumberValue
		var mtu basetypes.Int64Value
		var peerMac basetypes.StringValue
		var peerPortId basetypes.StringValue
		var peerRouterName basetypes.StringValue
		var peerSiteId basetypes.StringValue
		var portId basetypes.StringValue
		var routerName basetypes.StringValue
		var vpType basetypes.StringValue
		var up basetypes.BoolValue
		var uptime basetypes.Int64Value

		if d.IsActive != nil {
			isActive = types.BoolValue(*d.IsActive)
		}
		if d.LastSeen.Value() != nil {
			lastSeen = types.Float64Value(*d.LastSeen.Value())
		}
		if d.Latency != nil {
			latency = types.NumberValue(big.NewFloat(*d.Latency))
		}
		if d.Mos != nil {
			mos = types.NumberValue(big.NewFloat(*d.Mos))
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.PeerMac != nil {
			peerMac = types.StringValue(*d.PeerMac)
		}
		if d.PeerPortId != nil {
			peerPortId = types.StringValue(*d.PeerPortId)
		}
		if d.PeerRouterName != nil {
			peerRouterName = types.StringValue(*d.PeerRouterName)
		}
		if d.PeerSiteId != nil {
			peerSiteId = types.StringValue(d.PeerSiteId.String())
		}
		if d.PortId != nil {
			portId = types.StringValue(*d.PortId)
		}
		if d.RouterName != nil {
			routerName = types.StringValue(*d.RouterName)
		}
		if d.Type != nil {
			vpType = types.StringValue(*d.Type)
		}
		if d.Up != nil {
			up = types.BoolValue(*d.Up)
		}
		if d.Uptime != nil {
			uptime = types.Int64Value(int64(*d.Uptime))
		}

		dataMapValue := map[string]attr.Value{
			"is_active":        isActive,
			"last_seen":        lastSeen,
			"latency":          latency,
			"mos":              mos,
			"mtu":              mtu,
			"peer_mac":         peerMac,
			"peer_port_id":     peerPortId,
			"peer_router_name": peerRouterName,
			"peer_site_id":     peerSiteId,
			"port_id":          portId,
			"router_name":      routerName,
			"type":             vpType,
			"up":               up,
			"uptime":           uptime,
		}
		data, e := NewVpnPeersValue(VpnPeersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, VpnPeersValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
