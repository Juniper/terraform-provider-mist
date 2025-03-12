package datasource_device_gateway_stats

import (
	"context"
	"math/big"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ifStatsServpInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IfStatPropertyServpInfo) basetypes.ObjectValue {

	var asn basetypes.StringValue
	var city basetypes.StringValue
	var countryCode basetypes.StringValue
	var latitude basetypes.NumberValue
	var longitude basetypes.NumberValue
	var org basetypes.StringValue
	var regionCode basetypes.StringValue

	if d.Asn != nil {
		asn = types.StringValue(*d.Asn)
	}
	if d.City != nil {
		city = types.StringValue(*d.City)
	}
	if d.CountryCode != nil {
		countryCode = types.StringValue(*d.CountryCode)
	}
	if d.Latitude != nil {
		latitude = types.NumberValue(big.NewFloat(*d.Latitude))
	}
	if d.Longitude != nil {
		longitude = types.NumberValue(big.NewFloat(*d.Longitude))
	}
	if d.Org != nil {
		org = types.StringValue(*d.Org)
	}
	if d.RegionCode != nil {
		regionCode = types.StringValue(*d.RegionCode)
	}

	dataMapValue := map[string]attr.Value{
		"asn":          asn,
		"city":         city,
		"country_code": countryCode,
		"latitude":     latitude,
		"longitude":    longitude,
		"org":          org,
		"region_code":  regionCode,
	}
	data, e := basetypes.NewObjectValue(ServpInfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func ifStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.IfStatProperty) basetypes.MapValue {

	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var addressMode basetypes.StringValue
		var ips = types.ListNull(types.StringType)
		var natAddresses = types.ListNull(types.StringType)
		var networkName basetypes.StringValue
		var portId basetypes.StringValue
		var portUsage basetypes.StringValue
		var redundancyState basetypes.StringValue
		var rxBytes basetypes.Int64Value
		var rxPkts basetypes.Int64Value
		var servpInfo = types.ObjectNull(ServpInfoValue{}.AttributeTypes(ctx))
		var txBytes basetypes.Int64Value
		var txPkts basetypes.Int64Value
		var up basetypes.BoolValue
		var vlan basetypes.Int64Value
		var wanName basetypes.StringValue
		var wanType basetypes.StringValue

		if d.AddressMode != nil {
			addressMode = types.StringValue(*d.AddressMode)
		}
		if d.Ips != nil {
			ips = mistutils.ListOfStringSdkToTerraform(d.Ips)
		}
		if d.NatAddresses != nil {
			natAddresses = mistutils.ListOfStringSdkToTerraform(d.NatAddresses)
		}
		if d.NetworkName != nil {
			networkName = types.StringValue(*d.NetworkName)
		}
		if d.PortId != nil {
			portId = types.StringValue(*d.PortId)
		}
		if d.PortUsage != nil {
			portUsage = types.StringValue(*d.PortUsage)
		}
		if d.RedundancyState != nil {
			redundancyState = types.StringValue(*d.RedundancyState)
		}
		if d.RxBytes.Value() != nil {
			rxBytes = types.Int64Value(int64(*d.RxBytes.Value()))
		}
		if d.RxPkts.Value() != nil {
			rxPkts = types.Int64Value(int64(*d.RxPkts.Value()))
		}
		if d.ServpInfo != nil {
			servpInfo = ifStatsServpInfoSdkToTerraform(ctx, diags, d.ServpInfo)
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
		if d.Vlan != nil {
			vlan = types.Int64Value(int64(*d.Vlan))
		}
		if d.WanName != nil {
			wanName = types.StringValue(*d.WanName)
		}
		if d.WanType != nil {
			wanType = types.StringValue(*d.WanType)
		}

		dataMapValue := map[string]attr.Value{
			"address_mode":     addressMode,
			"ips":              ips,
			"nat_addresses":    natAddresses,
			"network_name":     networkName,
			"port_id":          portId,
			"port_usage":       portUsage,
			"redundancy_state": redundancyState,
			"rx_bytes":         rxBytes,
			"rx_pkts":          rxPkts,
			"servp_info":       servpInfo,
			"tx_bytes":         txBytes,
			"tx_pkts":          txPkts,
			"up":               up,
			"vlan":             vlan,
			"wan_name":         wanName,
			"wan_type":         wanType,
		}
		data, e := basetypes.NewObjectValue(IfStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, IfStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
