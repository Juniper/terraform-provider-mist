package datasource_device_gateway_stats

import (
	"context"
	"math/big"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ifStatsServpInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IfStatPropertyServpInfo) basetypes.ObjectValue {

	var asn basetypes.StringValue
	var city basetypes.StringValue
	var country_code basetypes.StringValue
	var latitude basetypes.NumberValue
	var longitude basetypes.NumberValue
	var org basetypes.StringValue
	var region_code basetypes.StringValue

	if d.Asn != nil {
		asn = types.StringValue(*d.Asn)
	}
	if d.City != nil {
		city = types.StringValue(*d.City)
	}
	if d.CountryCode != nil {
		country_code = types.StringValue(*d.CountryCode)
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
		region_code = types.StringValue(*d.RegionCode)
	}

	data_map_attr_type := ServpInfoValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"asn":          asn,
		"city":         city,
		"country_code": country_code,
		"latitude":     latitude,
		"longitude":    longitude,
		"org":          org,
		"region_code":  region_code,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func ifStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.IfStatProperty) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var address_mode basetypes.StringValue
		var ips basetypes.ListValue = types.ListNull(types.StringType)
		var nat_addresses basetypes.ListValue = types.ListNull(types.StringType)
		var network_name basetypes.StringValue
		var port_id basetypes.StringValue
		var port_usage basetypes.StringValue
		var redundancy_state basetypes.StringValue
		var rx_bytes basetypes.Int64Value
		var rx_pkts basetypes.Int64Value
		var servp_info basetypes.ObjectValue = types.ObjectNull(ServpInfoValue{}.AttributeTypes(ctx))
		var tx_bytes basetypes.Int64Value
		var tx_pkts basetypes.Int64Value
		var up basetypes.BoolValue
		var vlan basetypes.Int64Value
		var wan_name basetypes.StringValue
		var wan_type basetypes.StringValue

		if d.AddressMode != nil {
			address_mode = types.StringValue(*d.AddressMode)
		}
		if d.Ips != nil {
			ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.Ips)
		}
		if d.NatAddresses != nil {
			nat_addresses = mist_transform.ListOfStringSdkToTerraform(ctx, d.NatAddresses)
		}
		if d.NetworkName != nil {
			network_name = types.StringValue(*d.NetworkName)
		}
		if d.PortId != nil {
			port_id = types.StringValue(*d.PortId)
		}
		if d.PortUsage != nil {
			port_usage = types.StringValue(*d.PortUsage)
		}
		if d.RedundancyState != nil {
			redundancy_state = types.StringValue(*d.RedundancyState)
		}
		if d.RxBytes != nil {
			rx_bytes = types.Int64Value(int64(*d.RxBytes))
		}
		if d.RxPkts != nil {
			rx_pkts = types.Int64Value(int64(*d.RxPkts))
		}
		if d.ServpInfo != nil {
			servp_info = ifStatsServpInfoSdkToTerraform(ctx, diags, d.ServpInfo)
		}
		if d.TxBytes != nil {
			tx_bytes = types.Int64Value(int64(*d.TxBytes))
		}
		if d.TxPkts != nil {
			tx_pkts = types.Int64Value(int64(*d.TxPkts))
		}
		if d.Up != nil {
			up = types.BoolValue(*d.Up)
		}
		if d.Vlan != nil {
			vlan = types.Int64Value(int64(*d.Vlan))
		}
		if d.WanName != nil {
			wan_name = types.StringValue(*d.WanName)
		}
		if d.WanType != nil {
			wan_type = types.StringValue(*d.WanType)
		}

		data_map_attr_type := IfStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"address_mode":     address_mode,
			"ips":              ips,
			"nat_addresses":    nat_addresses,
			"network_name":     network_name,
			"port_id":          port_id,
			"port_usage":       port_usage,
			"redundancy_state": redundancy_state,
			"rx_bytes":         rx_bytes,
			"rx_pkts":          rx_pkts,
			"servp_info":       servp_info,
			"tx_bytes":         tx_bytes,
			"tx_pkts":          tx_pkts,
			"up":               up,
			"vlan":             vlan,
			"wan_name":         wan_name,
			"wan_type":         wan_type,
		}
		data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, IfStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
