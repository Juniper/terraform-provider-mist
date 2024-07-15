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
	var idle_time basetypes.Int64Value
	var last_seen basetypes.Int64Value
	var proto basetypes.StringValue
	var rssi basetypes.Int64Value
	var rx_bps basetypes.Int64Value
	var rx_bytes basetypes.Int64Value
	var rx_packets basetypes.Int64Value
	var rx_rate basetypes.Int64Value
	var rx_retries basetypes.Int64Value
	var site_id basetypes.StringValue
	var snr basetypes.Int64Value
	var tx_bps basetypes.Int64Value
	var tx_bytes basetypes.Int64Value
	var tx_packets basetypes.Int64Value
	var tx_rate basetypes.Int64Value
	var tx_retries basetypes.Int64Value
	var uplink_ap_id basetypes.StringValue

	if d.Band != nil {
		band = types.StringValue(*d.Band)
	}
	if d.Channel != nil {
		channel = types.Int64Value(int64(*d.Channel))
	}
	if d.IdleTime != nil {
		idle_time = types.Int64Value(int64(*d.IdleTime))
	}
	if d.LastSeen != nil {
		last_seen = types.Int64Value(int64(*d.LastSeen))
	}
	if d.Proto != nil {
		proto = types.StringValue(*d.Proto)
	}
	if d.Rssi != nil {
		rssi = types.Int64Value(int64(*d.Rssi))
	}
	if d.RxBps != nil {
		rx_bps = types.Int64Value(int64(*d.RxBps))
	}
	if d.RxBytes != nil {
		rx_bytes = types.Int64Value(int64(*d.RxBytes))
	}
	if d.RxPackets != nil {
		rx_packets = types.Int64Value(int64(*d.RxPackets))
	}
	if d.RxRate != nil {
		rx_rate = types.Int64Value(int64(*d.RxRate))
	}
	if d.RxRetries != nil {
		rx_retries = types.Int64Value(int64(*d.RxRetries))
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.Snr != nil {
		snr = types.Int64Value(int64(*d.Snr))
	}
	if d.TxBps != nil {
		tx_bps = types.Int64Value(int64(*d.TxBps))
	}
	if d.TxBytes != nil {
		tx_bytes = types.Int64Value(int64(*d.TxBytes))
	}
	if d.TxPackets != nil {
		tx_packets = types.Int64Value(int64(*d.TxPackets))
	}
	if d.TxRate != nil {
		tx_rate = types.Int64Value(int64(*d.TxRate))
	}
	if d.TxRetries != nil {
		tx_retries = types.Int64Value(int64(*d.TxRetries))
	}
	if d.UplinkApId != nil {
		uplink_ap_id = types.StringValue(d.UplinkApId.String())
	}

	data_map_attr_type := MeshUplinkValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"band":         band,
		"channel":      channel,
		"idle_time":    idle_time,
		"last_seen":    last_seen,
		"proto":        proto,
		"rssi":         rssi,
		"rx_bps":       rx_bps,
		"rx_bytes":     rx_bytes,
		"rx_packets":   rx_packets,
		"rx_rate":      rx_rate,
		"rx_retries":   rx_retries,
		"site_id":      site_id,
		"snr":          snr,
		"tx_bps":       tx_bps,
		"tx_bytes":     tx_bytes,
		"tx_packets":   tx_packets,
		"tx_rate":      tx_rate,
		"tx_retries":   tx_retries,
		"uplink_ap_id": uplink_ap_id,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
