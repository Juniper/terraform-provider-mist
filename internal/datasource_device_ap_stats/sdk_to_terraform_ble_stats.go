package datasource_device_ap_stats

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bleStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApBle) basetypes.ObjectValue {
	var beacon_enabled basetypes.BoolValue
	var beacon_rate basetypes.Int64Value
	var eddystone_uid_enabled basetypes.BoolValue
	var eddystone_uid_freq_msec basetypes.Int64Value
	var eddystone_uid_instance basetypes.StringValue
	var eddystone_uid_namespace basetypes.StringValue
	var eddystone_url_enabled basetypes.BoolValue
	var eddystone_url_freq_msec basetypes.Int64Value
	var eddystone_url_url basetypes.StringValue
	var ibeacon_enabled basetypes.BoolValue
	var ibeacon_freq_msec basetypes.Int64Value
	var ibeacon_major basetypes.Int64Value
	var ibeacon_minor basetypes.Int64Value
	var ibeacon_uuid basetypes.StringValue
	var major basetypes.Int64Value
	var minors basetypes.ListValue = types.ListNull(types.StringType)
	var power basetypes.Int64Value
	var rx_bytes basetypes.Int64Value
	var rx_pkts basetypes.Int64Value
	var tx_bytes basetypes.Int64Value
	var tx_pkts basetypes.Int64Value
	var tx_resets basetypes.Int64Value
	var uuid basetypes.StringValue

	if d.BeaconEnabled.Value() != nil {
		beacon_enabled = types.BoolValue(*d.BeaconEnabled.Value())
	}
	if d.BeaconRate.Value() != nil {
		beacon_rate = types.Int64Value(int64(*d.BeaconRate.Value()))
	}
	if d.EddystoneUidEnabled.Value() != nil {
		eddystone_uid_enabled = types.BoolValue(*d.EddystoneUidEnabled.Value())
	}
	if d.EddystoneUidFreqMsec.Value() != nil {
		eddystone_uid_freq_msec = types.Int64Value(int64(*d.EddystoneUidFreqMsec.Value()))
	}
	if d.EddystoneUidInstance.Value() != nil {
		eddystone_uid_instance = types.StringValue(*d.EddystoneUidInstance.Value())
	}
	if d.EddystoneUidNamespace.Value() != nil {
		eddystone_uid_namespace = types.StringValue(*d.EddystoneUidNamespace.Value())
	}
	if d.EddystoneUrlEnabled.Value() != nil {
		eddystone_url_enabled = types.BoolValue(*d.EddystoneUrlEnabled.Value())
	}
	if d.EddystoneUrlFreqMsec.Value() != nil {
		eddystone_url_freq_msec = types.Int64Value(int64(*d.EddystoneUrlFreqMsec.Value()))
	}
	if d.EddystoneUrlUrl.Value() != nil {
		eddystone_url_url = types.StringValue(*d.EddystoneUrlUrl.Value())
	}
	if d.IbeaconEnabled.Value() != nil {
		ibeacon_enabled = types.BoolValue(*d.IbeaconEnabled.Value())
	}
	if d.IbeaconFreqMsec.Value() != nil {
		ibeacon_freq_msec = types.Int64Value(int64(*d.IbeaconFreqMsec.Value()))
	}
	if d.IbeaconMajor.Value() != nil {
		ibeacon_major = types.Int64Value(int64(*d.IbeaconMajor.Value()))
	}
	if d.IbeaconMinor.Value() != nil {
		ibeacon_minor = types.Int64Value(int64(*d.IbeaconMinor.Value()))
	}
	if d.IbeaconUuid.Value() != nil {
		ibeacon_uuid = types.StringValue(d.IbeaconUuid.Value().String())
	}
	if d.Major.Value() != nil {
		major = types.Int64Value(int64(*d.Major.Value()))
	}
	if d.Minors != nil {
		minors = mist_transform.ListOfIntSdkToTerraform(ctx, d.Minors)
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
	if d.TxResets.Value() != nil {
		tx_resets = types.Int64Value(int64(*d.TxResets.Value()))
	}
	if d.Uuid.Value() != nil {
		uuid = types.StringValue(d.Uuid.Value().String())
	}

	data_map_attr_type := BleStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"beacon_enabled":          beacon_enabled,
		"beacon_rate":             beacon_rate,
		"eddystone_uid_enabled":   eddystone_uid_enabled,
		"eddystone_uid_freq_msec": eddystone_uid_freq_msec,
		"eddystone_uid_instance":  eddystone_uid_instance,
		"eddystone_uid_namespace": eddystone_uid_namespace,
		"eddystone_url_enabled":   eddystone_url_enabled,
		"eddystone_url_freq_msec": eddystone_url_freq_msec,
		"eddystone_url_url":       eddystone_url_url,
		"ibeacon_enabled":         ibeacon_enabled,
		"ibeacon_freq_msec":       ibeacon_freq_msec,
		"ibeacon_major":           ibeacon_major,
		"ibeacon_minor":           ibeacon_minor,
		"ibeacon_uuid":            ibeacon_uuid,
		"major":                   major,
		"minors":                  minors,
		"power":                   power,
		"rx_bytes":                rx_bytes,
		"rx_pkts":                 rx_pkts,
		"tx_bytes":                tx_bytes,
		"tx_pkts":                 tx_pkts,
		"tx_resets":               tx_resets,
		"uuid":                    uuid,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
