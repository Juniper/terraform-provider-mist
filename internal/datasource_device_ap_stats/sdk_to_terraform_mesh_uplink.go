package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func meshUplinkSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStatMeshUplink) basetypes.ObjectValue {

	var band basetypes.StringValue
	var channel basetypes.Int64Value
	var idleTime basetypes.Int64Value
	var lastSeen basetypes.Int64Value
	var proto basetypes.StringValue
	var rssi basetypes.Int64Value
	var rxBps basetypes.Int64Value
	var rxBytes basetypes.Int64Value
	var rxPackets basetypes.Int64Value
	var rxRate basetypes.Int64Value
	var rxRetries basetypes.Int64Value
	var siteId basetypes.StringValue
	var snr basetypes.Int64Value
	var txBps basetypes.Int64Value
	var txBytes basetypes.Int64Value
	var txPackets basetypes.Int64Value
	var txRate basetypes.Int64Value
	var txRetries basetypes.Int64Value
	var uplinkApId basetypes.StringValue

	if d.Band != nil {
		band = types.StringValue(*d.Band)
	}
	if d.Channel != nil {
		channel = types.Int64Value(int64(*d.Channel))
	}
	if d.IdleTime != nil {
		idleTime = types.Int64Value(int64(*d.IdleTime))
	}
	if d.LastSeen.Value() != nil {
		lastSeen = types.Int64Value(int64(*d.LastSeen.Value()))
	}
	if d.Proto != nil {
		proto = types.StringValue(*d.Proto)
	}
	if d.Rssi != nil {
		rssi = types.Int64Value(int64(*d.Rssi))
	}
	if d.RxBps.Value() != nil {
		rxBps = types.Int64Value(int64(*d.RxBps.Value()))
	}
	if d.RxBytes.Value() != nil {
		rxBytes = types.Int64Value(int64(*d.RxBytes.Value()))
	}
	if d.RxPackets.Value() != nil {
		rxPackets = types.Int64Value(int64(*d.RxPackets.Value()))
	}
	if d.RxRate.Value() != nil {
		rxRate = types.Int64Value(int64(*d.RxRate.Value()))
	}
	if d.RxRetries.Value() != nil {
		rxRetries = types.Int64Value(int64(*d.RxRetries.Value()))
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.Snr != nil {
		snr = types.Int64Value(int64(*d.Snr))
	}
	if d.TxBps.Value() != nil {
		txBps = types.Int64Value(int64(*d.TxBps.Value()))
	}
	if d.TxBytes.Value() != nil {
		txBytes = types.Int64Value(int64(*d.TxBytes.Value()))
	}
	if d.TxPackets.Value() != nil {
		txPackets = types.Int64Value(int64(*d.TxPackets.Value()))
	}
	if d.TxRate.Value() != nil {
		txRate = types.Int64Value(int64(*d.TxRate.Value()))
	}
	if d.TxRetries.Value() != nil {
		txRetries = types.Int64Value(int64(*d.TxRetries.Value()))
	}
	if d.UplinkApId != nil {
		uplinkApId = types.StringValue(d.UplinkApId.String())
	}

	dataMapValue := map[string]attr.Value{
		"band":         band,
		"channel":      channel,
		"idle_time":    idleTime,
		"last_seen":    lastSeen,
		"proto":        proto,
		"rssi":         rssi,
		"rx_bps":       rxBps,
		"rx_bytes":     rxBytes,
		"rx_packets":   rxPackets,
		"rx_rate":      rxRate,
		"rx_retries":   rxRetries,
		"site_id":      siteId,
		"snr":          snr,
		"tx_bps":       txBps,
		"tx_bytes":     txBytes,
		"tx_packets":   txPackets,
		"tx_rate":      txRate,
		"tx_retries":   txRetries,
		"uplink_ap_id": uplinkApId,
	}
	data, e := basetypes.NewObjectValue(MeshUplinkValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
