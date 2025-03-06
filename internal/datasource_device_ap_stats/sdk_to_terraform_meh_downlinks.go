package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func meshDownlinksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ApStatMeshDownlink) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
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
		if d.RxBps != nil {
			rxBps = types.Int64Value(int64(*d.RxBps))
		}
		if d.RxBytes != nil {
			rxBytes = types.Int64Value(int64(*d.RxBytes))
		}
		if d.RxPackets != nil {
			rxPackets = types.Int64Value(int64(*d.RxPackets))
		}
		if d.RxRate != nil {
			rxRate = types.Int64Value(int64(*d.RxRate))
		}
		if d.RxRetries != nil {
			rxRetries = types.Int64Value(int64(*d.RxRetries))
		}
		if d.SiteId != nil {
			siteId = types.StringValue(d.SiteId.String())
		}
		if d.Snr != nil {
			snr = types.Int64Value(int64(*d.Snr))
		}
		if d.TxBps != nil {
			txBps = types.Int64Value(int64(*d.TxBps))
		}
		if d.TxBytes != nil {
			txBytes = types.Int64Value(int64(*d.TxBytes))
		}
		if d.TxPackets != nil {
			txPackets = types.Int64Value(int64(*d.TxPackets))
		}
		if d.TxRate != nil {
			txRate = types.Int64Value(int64(*d.TxRate))
		}
		if d.TxRetries != nil {
			txRetries = types.Int64Value(int64(*d.TxRetries))
		}

		dataMapValue := map[string]attr.Value{
			"band":       band,
			"channel":    channel,
			"idle_time":  idleTime,
			"last_seen":  lastSeen,
			"proto":      proto,
			"rssi":       rssi,
			"rx_bps":     rxBps,
			"rx_bytes":   rxBytes,
			"rx_packets": rxPackets,
			"rx_rate":    rxRate,
			"rx_retries": rxRetries,
			"site_id":    siteId,
			"snr":        snr,
			"tx_bps":     txBps,
			"tx_bytes":   txBytes,
			"tx_packets": txPackets,
			"tx_rate":    txRate,
			"tx_retries": txRetries,
		}
		data, e := NewMeshDownlinksValue(MeshDownlinksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, IotStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
