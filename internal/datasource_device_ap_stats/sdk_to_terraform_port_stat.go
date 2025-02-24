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
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var fullDuplex basetypes.BoolValue
		var rxBytes basetypes.NumberValue
		var rxErrors basetypes.NumberValue
		var rxPkts basetypes.NumberValue
		var speed basetypes.Int64Value
		var txBytes basetypes.NumberValue
		var txPkts basetypes.NumberValue
		var up basetypes.BoolValue

		if d.FullDuplex.Value() != nil {
			fullDuplex = types.BoolValue(*d.FullDuplex.Value())
		}
		if d.RxBytes.Value() != nil {
			rxBytes = types.NumberValue(big.NewFloat(*d.RxBytes.Value()))
		}
		if d.RxErrors.Value() != nil {
			rxErrors = types.NumberValue(big.NewFloat(*d.RxErrors.Value()))
		}
		if d.RxPkts.Value() != nil {
			rxPkts = types.NumberValue(big.NewFloat(*d.RxPkts.Value()))
		}
		if d.Speed.Value() != nil {
			speed = types.Int64Value(int64(*d.Speed.Value()))
		}
		if d.TxBytes.Value() != nil {
			txBytes = types.NumberValue(big.NewFloat(*d.TxBytes.Value()))
		}
		if d.TxPkts.Value() != nil {
			txPkts = types.NumberValue(big.NewFloat(*d.TxPkts.Value()))
		}
		if d.Up.Value() != nil {
			up = types.BoolValue(*d.Up.Value())
		}

		dataMapValue := map[string]attr.Value{
			"full_duplex": fullDuplex,
			"rx_bytes":    rxBytes,
			"rx_errors":   rxErrors,
			"rx_pkts":     rxPkts,
			"speed":       speed,
			"tx_bytes":    txBytes,
			"tx_pkts":     txPkts,
			"up":          up,
		}
		data, e := NewPortStatValue(PortStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PortStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
