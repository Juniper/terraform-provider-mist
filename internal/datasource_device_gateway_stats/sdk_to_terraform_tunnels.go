package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsGatewayWanTunnel) basetypes.ListValue {

	var dataList []TunnelsValue
	for _, d := range l {
		var authAlgo basetypes.StringValue
		var encryptAlgo basetypes.StringValue
		var ikeVersion basetypes.StringValue
		var ip basetypes.StringValue
		var lastEvent basetypes.StringValue
		var lastFlapped basetypes.Float64Value
		var node basetypes.StringValue
		var peerHost basetypes.StringValue
		var peerIp basetypes.StringValue
		var priority basetypes.StringValue
		var protocol basetypes.StringValue
		var rxBytes basetypes.Int64Value
		var rxPkts basetypes.Int64Value
		var tunnelName basetypes.StringValue
		var txBytes basetypes.Int64Value
		var txPkts basetypes.Int64Value
		var up basetypes.BoolValue
		var uptime basetypes.Int64Value
		var wanName basetypes.StringValue

		if d.AuthAlgo != nil {
			authAlgo = types.StringValue(*d.AuthAlgo)
		}
		if d.EncryptAlgo != nil {
			encryptAlgo = types.StringValue(*d.EncryptAlgo)
		}
		if d.IkeVersion != nil {
			ikeVersion = types.StringValue(*d.IkeVersion)
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.LastEvent != nil {
			lastEvent = types.StringValue(*d.LastEvent)
		}
		if d.LastFlapped != nil {
			lastFlapped = types.Float64Value(*d.LastFlapped)
		}
		if d.Node != nil {
			node = types.StringValue(*d.Node)
		}
		if d.PeerHost != nil {
			peerHost = types.StringValue(*d.PeerHost)
		}
		if d.PeerIp != nil {
			peerIp = types.StringValue(*d.PeerIp)
		}
		if d.Priority != nil {
			priority = types.StringValue(string(*d.Priority))
		}
		if d.Protocol != nil {
			protocol = types.StringValue(string(*d.Protocol))
		}
		if d.RxBytes.Value() != nil {
			rxBytes = types.Int64Value(int64(*d.RxBytes.Value()))
		}
		if d.RxPkts.Value() != nil {
			rxPkts = types.Int64Value(int64(*d.RxPkts.Value()))
		}
		if d.TunnelName != nil {
			tunnelName = types.StringValue(*d.TunnelName)
		}
		if d.TxBytes.Value() != nil {
			txBytes = types.Int64Value(int64(*d.TxBytes.Value()))
		}
		if d.TxPkts.Value() != nil {
			txPkts = types.Int64Value(int64(*d.TxPkts.Value()))
		}
		if d.Up != nil {
			up = types.BoolValue(*d.Up)
		}
		if d.Uptime != nil {
			uptime = types.Int64Value(int64(*d.Uptime))
		}
		if d.WanName != nil {
			wanName = types.StringValue(*d.WanName)
		}

		dataMapValue := map[string]attr.Value{
			"auth_algo":    authAlgo,
			"encrypt_algo": encryptAlgo,
			"ike_version":  ikeVersion,
			"ip":           ip,
			"last_event":   lastEvent,
			"last_flapped": lastFlapped,
			"node":         node,
			"peer_host":    peerHost,
			"peer_ip":      peerIp,
			"priority":     priority,
			"protocol":     protocol,
			"rx_bytes":     rxBytes,
			"rx_pkts":      rxPkts,
			"tunnel_name":  tunnelName,
			"tx_bytes":     txBytes,
			"tx_pkts":      txPkts,
			"up":           up,
			"uptime":       uptime,
			"wan_name":     wanName,
		}
		data, e := NewTunnelsValue(TunnelsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, TunnelsValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
