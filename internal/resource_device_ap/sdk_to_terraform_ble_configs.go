package resource_device_ap

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bleConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.BleConfig) BleConfigValue {

	var beacon_enabled basetypes.BoolValue
	var beacon_rate basetypes.Int64Value
	var beacon_rate_mode basetypes.StringValue
	var beam_disabled basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var custom_ble_packet_enabled basetypes.BoolValue
	var custom_ble_packet_frame basetypes.StringValue
	var custom_ble_packet_freq_msec basetypes.Int64Value
	var eddystone_uid_adv_power basetypes.Int64Value
	var eddystone_uid_beams basetypes.StringValue
	var eddystone_uid_enabled basetypes.BoolValue
	var eddystone_uid_freq_msec basetypes.Int64Value
	var eddystone_uid_instance basetypes.StringValue
	var eddystone_uid_namespace basetypes.StringValue
	var eddystone_url_adv_power basetypes.Int64Value
	var eddystone_url_beams basetypes.StringValue
	var eddystone_url_enabled basetypes.BoolValue
	var eddystone_url_freq_msec basetypes.Int64Value
	var eddystone_url_url basetypes.StringValue
	var ibeacon_adv_power basetypes.Int64Value
	var ibeacon_beams basetypes.StringValue
	var ibeacon_enabled basetypes.BoolValue
	var ibeacon_freq_msec basetypes.Int64Value
	var ibeacon_major basetypes.Int64Value
	var ibeacon_minor basetypes.Int64Value
	var ibeacon_uuid basetypes.StringValue
	var power basetypes.Int64Value
	var power_mode basetypes.StringValue

	if d.BeaconEnabled != nil {
		beacon_enabled = types.BoolValue(*d.BeaconEnabled)
	}
	if d.BeaconRate != nil {
		beacon_rate = types.Int64Value(int64(*d.BeaconRate))
	}
	if d.BeaconRateMode != nil {
		beacon_rate_mode = types.StringValue(string(*d.BeaconRateMode))
	}
	if d.BeamDisabled != nil {
		beam_disabled = mist_transform.ListOfIntSdkToTerraform(ctx, d.BeamDisabled)
	}
	if d.CustomBlePacketEnabled != nil {
		custom_ble_packet_enabled = types.BoolValue(*d.CustomBlePacketEnabled)
	}
	if d.CustomBlePacketFrame != nil {
		custom_ble_packet_frame = types.StringValue(*d.CustomBlePacketFrame)
	}
	if d.CustomBlePacketFreqMsec != nil {
		custom_ble_packet_freq_msec = types.Int64Value(int64(*d.CustomBlePacketFreqMsec))
	}
	if d.EddystoneUidAdvPower != nil {
		eddystone_uid_adv_power = types.Int64Value(int64(*d.EddystoneUidAdvPower))
	}
	if d.EddystoneUidBeams != nil {
		eddystone_uid_beams = types.StringValue(*d.EddystoneUidBeams)
	}
	if d.EddystoneUidEnabled != nil {
		eddystone_uid_enabled = types.BoolValue(*d.EddystoneUidEnabled)
	}
	if d.EddystoneUidFreqMsec != nil {
		eddystone_uid_freq_msec = types.Int64Value(int64(*d.EddystoneUidFreqMsec))
	}
	if d.EddystoneUidInstance != nil {
		eddystone_uid_instance = types.StringValue(*d.EddystoneUidInstance)
	}
	if d.EddystoneUidNamespace != nil {
		eddystone_uid_namespace = types.StringValue(*d.EddystoneUidNamespace)
	}
	if d.EddystoneUrlAdvPower != nil {
		eddystone_url_adv_power = types.Int64Value(int64(*d.EddystoneUrlAdvPower))
	}
	if d.EddystoneUrlBeams != nil {
		eddystone_url_beams = types.StringValue(*d.EddystoneUrlBeams)
	}
	if d.EddystoneUrlEnabled != nil {
		eddystone_url_enabled = types.BoolValue(*d.EddystoneUrlEnabled)
	}
	if d.EddystoneUrlFreqMsec != nil {
		eddystone_url_freq_msec = types.Int64Value(int64(*d.EddystoneUrlFreqMsec))
	}
	if d.EddystoneUrlUrl != nil {
		eddystone_url_url = types.StringValue(*d.EddystoneUrlUrl)
	}
	if d.IbeaconAdvPower != nil {
		ibeacon_adv_power = types.Int64Value(int64(*d.IbeaconAdvPower))
	}
	if d.IbeaconBeams != nil {
		ibeacon_beams = types.StringValue(*d.IbeaconBeams)
	}
	if d.IbeaconEnabled != nil {
		ibeacon_enabled = types.BoolValue(*d.IbeaconEnabled)
	}
	if d.IbeaconFreqMsec != nil {
		ibeacon_freq_msec = types.Int64Value(int64(*d.IbeaconFreqMsec))
	}
	if d.IbeaconMajor != nil {
		ibeacon_major = types.Int64Value(int64(*d.IbeaconMajor))
	}
	if d.IbeaconMinor != nil {
		ibeacon_minor = types.Int64Value(int64(*d.IbeaconMinor))
	}
	if d.IbeaconUuid != nil {
		ibeacon_uuid = types.StringValue(d.IbeaconUuid.String())
	}
	if d.Power != nil {
		power = types.Int64Value(int64(*d.Power))
	}
	if d.PowerMode != nil {
		power_mode = types.StringValue(string(*d.PowerMode))
	}

	data_map_attr_type := BleConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"beacon_enabled":              beacon_enabled,
		"beacon_rate":                 beacon_rate,
		"beacon_rate_mode":            beacon_rate_mode,
		"beam_disabled":               beam_disabled,
		"custom_ble_packet_enabled":   custom_ble_packet_enabled,
		"custom_ble_packet_frame":     custom_ble_packet_frame,
		"custom_ble_packet_freq_msec": custom_ble_packet_freq_msec,
		"eddystone_uid_adv_power":     eddystone_uid_adv_power,
		"eddystone_uid_beams":         eddystone_uid_beams,
		"eddystone_uid_enabled":       eddystone_uid_enabled,
		"eddystone_uid_freq_msec":     eddystone_uid_freq_msec,
		"eddystone_uid_instance":      eddystone_uid_instance,
		"eddystone_uid_namespace":     eddystone_uid_namespace,
		"eddystone_url_adv_power":     eddystone_url_adv_power,
		"eddystone_url_beams":         eddystone_url_beams,
		"eddystone_url_enabled":       eddystone_url_enabled,
		"eddystone_url_freq_msec":     eddystone_url_freq_msec,
		"eddystone_url_url":           eddystone_url_url,
		"ibeacon_adv_power":           ibeacon_adv_power,
		"ibeacon_beams":               ibeacon_beams,
		"ibeacon_enabled":             ibeacon_enabled,
		"ibeacon_freq_msec":           ibeacon_freq_msec,
		"ibeacon_major":               ibeacon_major,
		"ibeacon_minor":               ibeacon_minor,
		"ibeacon_uuid":                ibeacon_uuid,
		"power":                       power,
		"power_mode":                  power_mode,
	}
	data, e := NewBleConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
