package datasource_device_ap_stats

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bleStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApBle) basetypes.ObjectValue {
	var beaconEnabled basetypes.BoolValue
	var beaconRate basetypes.Int64Value
	var eddystoneUidEnabled basetypes.BoolValue
	var eddystoneUidFreqMsec basetypes.Int64Value
	var eddystoneUidInstance basetypes.StringValue
	var eddystoneUidNamespace basetypes.StringValue
	var eddystoneUrlEnabled basetypes.BoolValue
	var eddystoneUrlFreqMsec basetypes.Int64Value
	var eddystoneUrlUrl basetypes.StringValue
	var ibeaconEnabled basetypes.BoolValue
	var ibeaconFreqMsec basetypes.Int64Value
	var ibeaconMajor basetypes.Int64Value
	var ibeaconMinor basetypes.Int64Value
	var ibeaconUuid basetypes.StringValue
	var major basetypes.Int64Value
	var minors = types.ListNull(types.StringType)
	var power basetypes.Int64Value
	var rxBytes basetypes.Int64Value
	var rxPkts basetypes.Int64Value
	var txBytes basetypes.Int64Value
	var txPkts basetypes.Int64Value
	var txResets basetypes.Int64Value
	var uuid basetypes.StringValue

	if d.BeaconEnabled.Value() != nil {
		beaconEnabled = types.BoolValue(*d.BeaconEnabled.Value())
	}
	if d.BeaconRate.Value() != nil {
		beaconRate = types.Int64Value(int64(*d.BeaconRate.Value()))
	}
	if d.EddystoneUidEnabled.Value() != nil {
		eddystoneUidEnabled = types.BoolValue(*d.EddystoneUidEnabled.Value())
	}
	if d.EddystoneUidFreqMsec.Value() != nil {
		eddystoneUidFreqMsec = types.Int64Value(int64(*d.EddystoneUidFreqMsec.Value()))
	}
	if d.EddystoneUidInstance.Value() != nil {
		eddystoneUidInstance = types.StringValue(*d.EddystoneUidInstance.Value())
	}
	if d.EddystoneUidNamespace.Value() != nil {
		eddystoneUidNamespace = types.StringValue(*d.EddystoneUidNamespace.Value())
	}
	if d.EddystoneUrlEnabled.Value() != nil {
		eddystoneUrlEnabled = types.BoolValue(*d.EddystoneUrlEnabled.Value())
	}
	if d.EddystoneUrlFreqMsec.Value() != nil {
		eddystoneUrlFreqMsec = types.Int64Value(int64(*d.EddystoneUrlFreqMsec.Value()))
	}
	if d.EddystoneUrlUrl.Value() != nil {
		eddystoneUrlUrl = types.StringValue(*d.EddystoneUrlUrl.Value())
	}
	if d.IbeaconEnabled.Value() != nil {
		ibeaconEnabled = types.BoolValue(*d.IbeaconEnabled.Value())
	}
	if d.IbeaconFreqMsec.Value() != nil {
		ibeaconFreqMsec = types.Int64Value(int64(*d.IbeaconFreqMsec.Value()))
	}
	if d.IbeaconMajor.Value() != nil {
		ibeaconMajor = types.Int64Value(int64(*d.IbeaconMajor.Value()))
	}
	if d.IbeaconMinor.Value() != nil {
		ibeaconMinor = types.Int64Value(int64(*d.IbeaconMinor.Value()))
	}
	if d.IbeaconUuid.Value() != nil {
		ibeaconUuid = types.StringValue(d.IbeaconUuid.Value().String())
	}
	if d.Major.Value() != nil {
		major = types.Int64Value(int64(*d.Major.Value()))
	}
	if d.Minors != nil {
		minors = mistutils.ListOfIntSdkToTerraform(d.Minors)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.RxBytes.Value() != nil {
		rxBytes = types.Int64Value(*d.RxBytes.Value())
	}
	if d.RxPkts.Value() != nil {
		rxPkts = types.Int64Value(*d.RxPkts.Value())
	}
	if d.TxBytes.Value() != nil {
		txBytes = types.Int64Value(*d.TxBytes.Value())
	}
	if d.TxPkts.Value() != nil {
		txPkts = types.Int64Value(*d.TxPkts.Value())
	}
	if d.TxResets.Value() != nil {
		txResets = types.Int64Value(int64(*d.TxResets.Value()))
	}
	if d.Uuid.Value() != nil {
		uuid = types.StringValue(d.Uuid.Value().String())
	}

	dataMapValue := map[string]attr.Value{
		"beacon_enabled":          beaconEnabled,
		"beacon_rate":             beaconRate,
		"eddystone_uid_enabled":   eddystoneUidEnabled,
		"eddystone_uid_freq_msec": eddystoneUidFreqMsec,
		"eddystone_uid_instance":  eddystoneUidInstance,
		"eddystone_uid_namespace": eddystoneUidNamespace,
		"eddystone_url_enabled":   eddystoneUrlEnabled,
		"eddystone_url_freq_msec": eddystoneUrlFreqMsec,
		"eddystone_url_url":       eddystoneUrlUrl,
		"ibeacon_enabled":         ibeaconEnabled,
		"ibeacon_freq_msec":       ibeaconFreqMsec,
		"ibeacon_major":           ibeaconMajor,
		"ibeacon_minor":           ibeaconMinor,
		"ibeacon_uuid":            ibeaconUuid,
		"major":                   major,
		"minors":                  minors,
		"power":                   power,
		"rx_bytes":                rxBytes,
		"rx_pkts":                 rxPkts,
		"tx_bytes":                txBytes,
		"tx_pkts":                 txPkts,
		"tx_resets":               txResets,
		"uuid":                    uuid,
	}
	data, e := types.ObjectValue(BleStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
