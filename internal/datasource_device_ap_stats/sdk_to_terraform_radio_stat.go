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
	var dynamic_chaining_enalbed basetypes.BoolValue
	var mac basetypes.StringValue
	var noise_floor basetypes.Int64Value
	var num_clients basetypes.Int64Value
	var power basetypes.Int64Value
	var rx_bytes basetypes.Int64Value
	var rx_pkts basetypes.Int64Value
	var tx_bytes basetypes.Int64Value
	var tx_pkts basetypes.Int64Value
	var usage basetypes.StringValue
	var util_all basetypes.Int64Value
	var util_non_wifi basetypes.Int64Value
	var util_rx_in_bss basetypes.Int64Value
	var util_rx_other_bss basetypes.Int64Value
	var util_tx basetypes.Int64Value
	var util_undecodable_wifi basetypes.Int64Value
	var util_unknown_wifi basetypes.Int64Value

	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.DynamicChainingEnalbed.Value() != nil {
		dynamic_chaining_enalbed = types.BoolValue(*d.DynamicChainingEnalbed.Value())
	}
	if d.Mac.Value() != nil {
		mac = types.StringValue(*d.Mac.Value())
	}
	if d.NoiseFloor.Value() != nil {
		noise_floor = types.Int64Value(int64(*d.NoiseFloor.Value()))
	}
	if d.NumClients.Value() != nil {
		num_clients = types.Int64Value(int64(*d.NumClients.Value()))
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.RxBytes.Value() != nil {
		rx_bytes = types.Int64Value(int64(*d.RxBytes.Value()))
	}
	if d.RxPkts.Value() != nil {
		rx_pkts = types.Int64Value(int64(*d.RxPkts.Value()))
	}
	if d.TxBytes.Value() != nil {
		tx_bytes = types.Int64Value(int64(*d.TxBytes.Value()))
	}
	if d.TxPkts.Value() != nil {
		tx_pkts = types.Int64Value(int64(*d.TxPkts.Value()))
	}
	if d.Usage.Value() != nil {
		usage = types.StringValue(*d.Usage.Value())
	}
	if d.UtilAll.Value() != nil {
		util_all = types.Int64Value(int64(*d.UtilAll.Value()))
	}
	if d.UtilNonWifi.Value() != nil {
		util_non_wifi = types.Int64Value(int64(*d.UtilNonWifi.Value()))
	}
	if d.UtilRxInBss.Value() != nil {
		util_rx_in_bss = types.Int64Value(int64(*d.UtilRxInBss.Value()))
	}
	if d.UtilRxOtherBss.Value() != nil {
		util_rx_other_bss = types.Int64Value(int64(*d.UtilRxOtherBss.Value()))
	}
	if d.UtilTx.Value() != nil {
		util_tx = types.Int64Value(int64(*d.UtilTx.Value()))
	}
	if d.UtilUndecodableWifi.Value() != nil {
		util_undecodable_wifi = types.Int64Value(int64(*d.UtilUndecodableWifi.Value()))
	}
	if d.UtilUnknownWifi.Value() != nil {
		util_unknown_wifi = types.Int64Value(int64(*d.UtilUnknownWifi.Value()))
	}

	data_map_attr_type := BandValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"bandwidth":                bandwidth,
		"channel":                  channel,
		"dynamic_chaining_enalbed": dynamic_chaining_enalbed,
		"mac":                      mac,
		"noise_floor":              noise_floor,
		"num_clients":              num_clients,
		"power":                    power,
		"rx_bytes":                 rx_bytes,
		"rx_pkts":                  rx_pkts,
		"tx_bytes":                 tx_bytes,
		"tx_pkts":                  tx_pkts,
		"usage":                    usage,
		"util_all":                 util_all,
		"util_non_wifi":            util_non_wifi,
		"util_rx_in_bss":           util_rx_in_bss,
		"util_rx_other_bss":        util_rx_other_bss,
		"util_tx":                  util_tx,
		"util_undecodable_wifi":    util_undecodable_wifi,
		"util_unknown_wifi":        util_unknown_wifi,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func radioStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApRadioStat) basetypes.ObjectValue {

	var band_24 basetypes.ObjectValue = types.ObjectNull(BandValue{}.AttributeTypes(ctx))
	var band_5 basetypes.ObjectValue = types.ObjectNull(BandValue{}.AttributeTypes(ctx))
	var band_6 basetypes.ObjectValue = types.ObjectNull(BandValue{}.AttributeTypes(ctx))

	if d.Band24 != nil {
		band_24 = radioStatBandSdkToTerraform(ctx, diags, d.Band24)
	}
	if d.Band5 != nil {
		band_5 = radioStatBandSdkToTerraform(ctx, diags, d.Band5)
	}
	if d.Band6 != nil {
		band_6 = radioStatBandSdkToTerraform(ctx, diags, d.Band6)
	}

	data_map_attr_type := RadioStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"band_24": band_24,
		"band_5":  band_5,
		"band_6":  band_6,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
