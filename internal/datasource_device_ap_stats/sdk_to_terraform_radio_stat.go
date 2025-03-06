package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radioStatBandSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioStat) basetypes.ObjectValue {

	var bandwidth basetypes.Int64Value
	var channel basetypes.Int64Value
	var dynamicChainingEnabled basetypes.BoolValue
	var mac basetypes.StringValue
	var noiseFloor basetypes.Int64Value
	var numClients basetypes.Int64Value
	var numWlans basetypes.Int64Value
	var power basetypes.Int64Value
	var rxBytes basetypes.Int64Value
	var rxPkts basetypes.Int64Value
	var txBytes basetypes.Int64Value
	var txPkts basetypes.Int64Value
	var usage basetypes.StringValue
	var utilAll basetypes.Int64Value
	var utilNonWifi basetypes.Int64Value
	var utilRxInBss basetypes.Int64Value
	var utilRxOtherBss basetypes.Int64Value
	var utilTx basetypes.Int64Value
	var utilUndecodableWifi basetypes.Int64Value
	var utilUnknownWifi basetypes.Int64Value

	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.DynamicChainingEnabled.Value() != nil {
		dynamicChainingEnabled = types.BoolValue(*d.DynamicChainingEnabled.Value())
	}
	if d.Mac.Value() != nil {
		mac = types.StringValue(*d.Mac.Value())
	}
	if d.NoiseFloor.Value() != nil {
		noiseFloor = types.Int64Value(int64(*d.NoiseFloor.Value()))
	}
	if d.NumClients.Value() != nil {
		numClients = types.Int64Value(int64(*d.NumClients.Value()))
	}
	if d.NumWlans != nil {
		numWlans = types.Int64Value(int64(*d.NumWlans))
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.RxBytes.Value() != nil {
		rxBytes = types.Int64Value(int64(*d.RxBytes.Value()))
	}
	if d.RxPkts.Value() != nil {
		rxPkts = types.Int64Value(int64(*d.RxPkts.Value()))
	}
	if d.TxBytes.Value() != nil {
		txBytes = types.Int64Value(int64(*d.TxBytes.Value()))
	}
	if d.TxPkts.Value() != nil {
		txPkts = types.Int64Value(int64(*d.TxPkts.Value()))
	}
	if d.Usage.Value() != nil {
		usage = types.StringValue(*d.Usage.Value())
	}
	if d.UtilAll.Value() != nil {
		utilAll = types.Int64Value(int64(*d.UtilAll.Value()))
	}
	if d.UtilNonWifi.Value() != nil {
		utilNonWifi = types.Int64Value(int64(*d.UtilNonWifi.Value()))
	}
	if d.UtilRxInBss.Value() != nil {
		utilRxInBss = types.Int64Value(int64(*d.UtilRxInBss.Value()))
	}
	if d.UtilRxOtherBss.Value() != nil {
		utilRxOtherBss = types.Int64Value(int64(*d.UtilRxOtherBss.Value()))
	}
	if d.UtilTx.Value() != nil {
		utilTx = types.Int64Value(int64(*d.UtilTx.Value()))
	}
	if d.UtilUndecodableWifi.Value() != nil {
		utilUndecodableWifi = types.Int64Value(int64(*d.UtilUndecodableWifi.Value()))
	}
	if d.UtilUnknownWifi.Value() != nil {
		utilUnknownWifi = types.Int64Value(int64(*d.UtilUnknownWifi.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"bandwidth":                bandwidth,
		"channel":                  channel,
		"dynamic_chaining_enabled": dynamicChainingEnabled,
		"mac":                      mac,
		"noise_floor":              noiseFloor,
		"num_clients":              numClients,
		"num_wlans":                numWlans,
		"power":                    power,
		"rx_bytes":                 rxBytes,
		"rx_pkts":                  rxPkts,
		"tx_bytes":                 txBytes,
		"tx_pkts":                  txPkts,
		"usage":                    usage,
		"util_all":                 utilAll,
		"util_non_wifi":            utilNonWifi,
		"util_rx_in_bss":           utilRxInBss,
		"util_rx_other_bss":        utilRxOtherBss,
		"util_tx":                  utilTx,
		"util_undecodable_wifi":    utilUndecodableWifi,
		"util_unknown_wifi":        utilUnknownWifi,
	}
	data, e := basetypes.NewObjectValue(BandValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func radioStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApRadioStat) basetypes.ObjectValue {

	var band24 = types.ObjectNull(BandValue{}.AttributeTypes(ctx))
	var band5 = types.ObjectNull(BandValue{}.AttributeTypes(ctx))
	var band6 = types.ObjectNull(BandValue{}.AttributeTypes(ctx))

	if d.Band24 != nil {
		band24 = radioStatBandSdkToTerraform(ctx, diags, d.Band24)
	}
	if d.Band5 != nil {
		band5 = radioStatBandSdkToTerraform(ctx, diags, d.Band5)
	}
	if d.Band6 != nil {
		band6 = radioStatBandSdkToTerraform(ctx, diags, d.Band6)
	}

	dataMapValue := map[string]attr.Value{
		"band_24": band24,
		"band_5":  band5,
		"band_6":  band6,
	}
	data, e := basetypes.NewObjectValue(RadioStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
