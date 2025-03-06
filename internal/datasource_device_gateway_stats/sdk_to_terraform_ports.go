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

func portsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsGatewayPort) basetypes.ListValue {

	var dataList []PortsValue
	for _, d := range l {
		var active basetypes.BoolValue
		var authState basetypes.StringValue
		var disabled basetypes.BoolValue
		var forSite basetypes.BoolValue
		var fullDuplex basetypes.BoolValue
		var jitter basetypes.NumberValue
		var latency basetypes.NumberValue
		var loss basetypes.NumberValue
		var lteIccid basetypes.StringValue
		var lteImei basetypes.StringValue
		var lteImsi basetypes.StringValue
		var macCount basetypes.Int64Value
		var macLimit basetypes.Int64Value
		var neighborMac basetypes.StringValue
		var neighborPortDesc basetypes.StringValue
		var neighborSystemName basetypes.StringValue
		var poeDisabled basetypes.BoolValue
		var poeMode basetypes.StringValue
		var poeOn basetypes.BoolValue
		var portId basetypes.StringValue
		var portMac basetypes.StringValue
		var portUsage basetypes.StringValue
		var powerDraw basetypes.NumberValue
		var rxBcastPkts basetypes.Int64Value
		var rxBps basetypes.Int64Value
		var rxBytes basetypes.Int64Value
		var rxErrors basetypes.Int64Value
		var rxMcastPkts basetypes.Int64Value
		var rxPkts basetypes.Int64Value
		var speed basetypes.Int64Value
		var stpRole basetypes.StringValue
		var stpState basetypes.StringValue
		var txBcastPkts basetypes.Int64Value
		var txBps basetypes.Int64Value
		var txBytes basetypes.Int64Value
		var txErrors basetypes.Int64Value
		var txMcastPkts basetypes.Int64Value
		var txPkts basetypes.Int64Value
		var iType basetypes.StringValue
		var unconfigured basetypes.BoolValue
		var up basetypes.BoolValue
		var xcvrModel basetypes.StringValue
		var xcvrPartNumber basetypes.StringValue
		var xcvrSerial basetypes.StringValue

		if d.Active != nil {
			active = types.BoolValue(*d.Active)
		}
		if d.AuthState != nil {
			authState = types.StringValue(string(*d.AuthState))
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.ForSite != nil {
			forSite = types.BoolValue(*d.ForSite)
		}
		if d.FullDuplex != nil {
			fullDuplex = types.BoolValue(*d.FullDuplex)
		}
		if d.Jitter != nil {
			jitter = types.NumberValue(big.NewFloat(*d.Jitter))
		}
		if d.Latency != nil {
			latency = types.NumberValue(big.NewFloat(*d.Latency))
		}
		if d.Loss != nil {
			loss = types.NumberValue(big.NewFloat(*d.Loss))
		}
		if d.LteIccid.Value() != nil {
			lteIccid = types.StringValue(*d.LteIccid.Value())
		}
		if d.LteImei.Value() != nil {
			lteImei = types.StringValue(*d.LteImei.Value())
		}
		if d.LteImsi.Value() != nil {
			lteImsi = types.StringValue(*d.LteImsi.Value())
		}
		if d.MacCount != nil {
			macCount = types.Int64Value(int64(*d.MacCount))
		}
		if d.MacLimit != nil {
			macLimit = types.Int64Value(int64(*d.MacLimit))
		}

		neighborMac = types.StringValue(d.NeighborMac)

		if d.NeighborPortDesc != nil {
			neighborPortDesc = types.StringValue(*d.NeighborPortDesc)
		}
		if d.NeighborSystemName != nil {
			neighborSystemName = types.StringValue(*d.NeighborSystemName)
		}
		if d.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PoeMode != nil {
			poeMode = types.StringValue(string(*d.PoeMode))
		}
		if d.PoeOn != nil {
			poeOn = types.BoolValue(*d.PoeOn)
		}

		portId = types.StringValue(d.PortId)

		portMac = types.StringValue(d.PortMac)

		if d.PortUsage != nil {
			portUsage = types.StringValue(string(*d.PortUsage))
		}
		if d.PowerDraw != nil {
			powerDraw = types.NumberValue(big.NewFloat(*d.PowerDraw))
		}
		if d.RxBcastPkts != nil {
			rxBcastPkts = types.Int64Value(int64(*d.RxBcastPkts))
		}
		if d.RxBps != nil {
			rxBps = types.Int64Value(int64(*d.RxBps))
		}
		if d.RxBytes != nil {
			rxBytes = types.Int64Value(int64(*d.RxBytes))
		}
		if d.RxErrors != nil {
			rxErrors = types.Int64Value(int64(*d.RxErrors))
		}
		if d.RxMcastPkts != nil {
			rxMcastPkts = types.Int64Value(int64(*d.RxMcastPkts))
		}
		if d.Speed != nil {
			speed = types.Int64Value(int64(*d.Speed))
		}
		if d.StpRole != nil {
			stpRole = types.StringValue(string(*d.StpRole))
		}
		if d.StpState != nil {
			stpState = types.StringValue(string(*d.StpState))
		}
		if d.TxBcastPkts != nil {
			txBcastPkts = types.Int64Value(int64(*d.TxBcastPkts))
		}
		if d.TxBps != nil {
			txBps = types.Int64Value(int64(*d.TxBps))
		}
		if d.TxBytes != nil {
			txBytes = types.Int64Value(int64(*d.TxBytes))
		}
		if d.TxErrors != nil {
			txErrors = types.Int64Value(int64(*d.TxErrors))
		}
		if d.TxMcastPkts != nil {
			txMcastPkts = types.Int64Value(int64(*d.TxMcastPkts))
		}
		if d.Type != nil {
			iType = types.StringValue(string(*d.Type))
		}
		if d.Unconfigured != nil {
			unconfigured = types.BoolValue(*d.Unconfigured)
		}
		if d.Up != nil {
			up = types.BoolValue(*d.Up)
		}
		if d.XcvrModel != nil {
			xcvrModel = types.StringValue(*d.XcvrModel)
		}
		if d.XcvrPartNumber != nil {
			xcvrPartNumber = types.StringValue(*d.XcvrPartNumber)
		}
		if d.XcvrSerial != nil {
			xcvrSerial = types.StringValue(*d.XcvrSerial)
		}

		dataMapValue := map[string]attr.Value{
			"active":               active,
			"auth_state":           authState,
			"disabled":             disabled,
			"for_site":             forSite,
			"full_duplex":          fullDuplex,
			"jitter":               jitter,
			"latency":              latency,
			"loss":                 loss,
			"lte_iccid":            lteIccid,
			"lte_imei":             lteImei,
			"lte_imsi":             lteImsi,
			"mac_count":            macCount,
			"mac_limit":            macLimit,
			"neighbor_mac":         neighborMac,
			"neighbor_port_desc":   neighborPortDesc,
			"neighbor_system_name": neighborSystemName,
			"poe_disabled":         poeDisabled,
			"poe_mode":             poeMode,
			"poe_on":               poeOn,
			"port_id":              portId,
			"port_mac":             portMac,
			"port_usage":           portUsage,
			"power_draw":           powerDraw,
			"rx_bcast_pkts":        rxBcastPkts,
			"rx_bps":               rxBps,
			"rx_bytes":             rxBytes,
			"rx_errors":            rxErrors,
			"rx_mcast_pkts":        rxMcastPkts,
			"rx_pkts":              rxPkts,
			"speed":                speed,
			"stp_role":             stpRole,
			"stp_state":            stpState,
			"tx_bcast_pkts":        txBcastPkts,
			"tx_bps":               txBps,
			"tx_bytes":             txBytes,
			"tx_errors":            txErrors,
			"tx_mcast_pkts":        txMcastPkts,
			"tx_pkts":              txPkts,
			"type":                 iType,
			"unconfigured":         unconfigured,
			"up":                   up,
			"xcvr_model":           xcvrModel,
			"xcvr_part_number":     xcvrPartNumber,
			"xcvr_serial":          xcvrSerial,
		}
		data, e := NewPortsValue(PortsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, PortsValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
