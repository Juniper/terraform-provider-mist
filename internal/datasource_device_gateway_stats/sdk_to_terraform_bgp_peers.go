package datasource_device_gateway_stats

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bgpPeersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.BgpPeer) basetypes.ListValue {

	var dataList []BgpPeersValue
	for _, d := range l {
		var evpnOverlay basetypes.BoolValue
		var forOverlay basetypes.BoolValue
		var localAs basetypes.StringValue
		var neighbor basetypes.StringValue
		var neighborAs basetypes.StringValue
		var neighborMac basetypes.StringValue
		var node basetypes.StringValue
		var rxPkts basetypes.Int64Value
		var rxRoutes basetypes.Int64Value
		var state basetypes.StringValue
		var timestamp basetypes.Float64Value
		var txPkts basetypes.Int64Value
		var txRoutes basetypes.Int64Value
		var up basetypes.BoolValue
		var uptime basetypes.Int64Value
		var vrfName basetypes.StringValue

		if d.EvpnOverlay != nil {
			evpnOverlay = types.BoolValue(*d.EvpnOverlay)
		}
		if d.ForOverlay != nil {
			forOverlay = types.BoolValue(*d.ForOverlay)
		}
		if d.LocalAs != nil {
			localAs = mistutils.BgpAsAsString(d.LocalAs)
		}
		if d.Neighbor != nil {
			neighbor = types.StringValue(*d.Neighbor)
		}
		if d.NeighborAs != nil {
			neighborAs = mistutils.BgpAsAsString(d.NeighborAs)
		}
		if d.NeighborMac != nil {
			neighborMac = types.StringValue(*d.NeighborMac)
		}
		if d.Node != nil {
			node = types.StringValue(*d.Node)
		}
		if d.RxPkts.Value() != nil {
			rxPkts = types.Int64Value(int64(*d.RxPkts.Value()))
		}
		if d.RxRoutes != nil {
			rxRoutes = types.Int64Value(int64(*d.RxRoutes))
		}
		if d.State != nil {
			state = types.StringValue(string(*d.State))
		}
		if d.Timestamp != nil {
			timestamp = types.Float64Value(*d.Timestamp)
		}
		if d.TxPkts.Value() != nil {
			txPkts = types.Int64Value(int64(*d.TxPkts.Value()))
		}
		if d.TxRoutes != nil {
			txRoutes = types.Int64Value(int64(*d.TxRoutes))
		}
		if d.Up != nil {
			up = types.BoolValue(*d.Up)
		}
		if d.Uptime != nil {
			uptime = types.Int64Value(int64(*d.Uptime))
		}
		if d.VrfName != nil {
			vrfName = types.StringValue(*d.VrfName)
		}

		dataMapValue := map[string]attr.Value{
			"evpn_overlay": evpnOverlay,
			"for_overlay":  forOverlay,
			"local_as":     localAs,
			"neighbor":     neighbor,
			"neighbor_as":  neighborAs,
			"neighbor_mac": neighborMac,
			"node":         node,
			"rx_pkts":      rxPkts,
			"rx_routes":    rxRoutes,
			"state":        state,
			"timestamp":    timestamp,
			"tx_pkts":      txPkts,
			"tx_routes":    txRoutes,
			"up":           up,
			"uptime":       uptime,
			"vrf_name":     vrfName,
		}
		data, e := NewBgpPeersValue(BgpPeersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, BgpPeersValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
