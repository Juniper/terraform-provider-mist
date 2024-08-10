package datasource_device_ap_stats

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portStatdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.StatsApPortStat) basetypes.MapValue {
	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var full_duplex basetypes.BoolValue
		var rx_bytes basetypes.NumberValue
		var rx_errors basetypes.NumberValue
		var rx_pkts basetypes.NumberValue
		var speed basetypes.Int64Value
		var tx_bytes basetypes.NumberValue
		var tx_pkts basetypes.NumberValue
		var up basetypes.BoolValue

		if d.FullDuplex.Value() != nil {
			full_duplex = types.BoolValue(*d.FullDuplex.Value())
		}
		if d.RxBytes.Value() != nil {
			rx_bytes = types.NumberValue(big.NewFloat(*d.RxBytes.Value()))
		}
		if d.RxErrors.Value() != nil {
			rx_errors = types.NumberValue(big.NewFloat(*d.RxErrors.Value()))
		}
		if d.RxPkts.Value() != nil {
			rx_pkts = types.NumberValue(big.NewFloat(*d.RxPkts.Value()))
		}
		if d.Speed.Value() != nil {
			speed = types.Int64Value(int64(*d.Speed.Value()))
		}
		if d.TxBytes.Value() != nil {
			tx_bytes = types.NumberValue(big.NewFloat(*d.TxBytes.Value()))
		}
		if d.TxPkts.Value() != nil {
			tx_pkts = types.NumberValue(big.NewFloat(*d.TxPkts.Value()))
		}
		if d.Up.Value() != nil {
			up = types.BoolValue(*d.Up.Value())
		}

		data_map_attr_type := PortStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"full_duplex": full_duplex,
			"rx_bytes":    rx_bytes,
			"rx_errors":   rx_errors,
			"rx_pkts":     rx_pkts,
			"speed":       speed,
			"tx_bytes":    tx_bytes,
			"tx_pkts":     tx_pkts,
			"up":          up,
		}
		data, e := NewPortStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, PortStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
