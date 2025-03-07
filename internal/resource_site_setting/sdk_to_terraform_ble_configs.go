package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func bleConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.BleConfig) BleConfigValue {

	var beaconEnabled basetypes.BoolValue
	var beaconRate basetypes.Int64Value
	var beaconRateMode basetypes.StringValue
	var beamDisabled = mistutils.ListOfIntSdkToTerraformEmpty()
	var customBlePacketEnabled basetypes.BoolValue
	var customBlePacketFrame basetypes.StringValue
	var customBlePacketFreqMsec basetypes.Int64Value
	var eddystoneUidAdvPower basetypes.Int64Value
	var eddystoneUidBeams basetypes.StringValue
	var eddystoneUidEnabled basetypes.BoolValue
	var eddystoneUidFreqMsec basetypes.Int64Value
	var eddystoneUidInstance basetypes.StringValue
	var eddystoneUidNamespace basetypes.StringValue
	var eddystoneUrlAdvPower basetypes.Int64Value
	var eddystoneUrlBeams basetypes.StringValue
	var eddystoneUrlEnabled basetypes.BoolValue
	var eddystoneUrlFreqMsec basetypes.Int64Value
	var eddystoneUrlUrl basetypes.StringValue
	var ibeaconAdvPower basetypes.Int64Value
	var ibeaconBeams basetypes.StringValue
	var ibeaconEnabled basetypes.BoolValue
	var ibeaconFreqMsec basetypes.Int64Value
	var ibeaconMajor basetypes.Int64Value
	var ibeaconMinor basetypes.Int64Value
	var ibeaconUuid basetypes.StringValue
	var power basetypes.Int64Value
	var powerMode basetypes.StringValue

	if d != nil && d.BeaconEnabled != nil {
		beaconEnabled = types.BoolValue(*d.BeaconEnabled)
	}
	if d != nil && d.BeaconRate != nil {
		beaconRate = types.Int64Value(int64(*d.BeaconRate))
	}
	if d != nil && d.BeaconRateMode != nil {
		beaconRateMode = types.StringValue(string(*d.BeaconRateMode))
	}
	if d != nil && d.BeamDisabled != nil {
		beamDisabled = mistutils.ListOfIntSdkToTerraform(d.BeamDisabled)
	}
	if d != nil && d.CustomBlePacketEnabled != nil {
		customBlePacketEnabled = types.BoolValue(*d.CustomBlePacketEnabled)
	}
	if d != nil && d.CustomBlePacketFrame != nil {
		customBlePacketFrame = types.StringValue(*d.CustomBlePacketFrame)
	}
	if d != nil && d.CustomBlePacketFreqMsec != nil {
		customBlePacketFreqMsec = types.Int64Value(int64(*d.CustomBlePacketFreqMsec))
	}
	if d != nil && d.EddystoneUidAdvPower != nil {
		eddystoneUidAdvPower = types.Int64Value(int64(*d.EddystoneUidAdvPower))
	}
	if d != nil && d.EddystoneUidBeams != nil {
		eddystoneUidBeams = types.StringValue(*d.EddystoneUidBeams)
	}
	if d != nil && d.EddystoneUidEnabled != nil {
		eddystoneUidEnabled = types.BoolValue(*d.EddystoneUidEnabled)
	}
	if d != nil && d.EddystoneUidFreqMsec != nil {
		eddystoneUidFreqMsec = types.Int64Value(int64(*d.EddystoneUidFreqMsec))
	}
	if d != nil && d.EddystoneUidInstance != nil {
		eddystoneUidInstance = types.StringValue(*d.EddystoneUidInstance)
	}
	if d != nil && d.EddystoneUidNamespace != nil {
		eddystoneUidNamespace = types.StringValue(*d.EddystoneUidNamespace)
	}
	if d != nil && d.EddystoneUrlAdvPower != nil {
		eddystoneUrlAdvPower = types.Int64Value(int64(*d.EddystoneUrlAdvPower))
	}
	if d != nil && d.EddystoneUrlBeams != nil {
		eddystoneUrlBeams = types.StringValue(*d.EddystoneUrlBeams)
	}
	if d != nil && d.EddystoneUrlEnabled != nil {
		eddystoneUrlEnabled = types.BoolValue(*d.EddystoneUrlEnabled)
	}
	if d != nil && d.EddystoneUrlFreqMsec != nil {
		eddystoneUrlFreqMsec = types.Int64Value(int64(*d.EddystoneUrlFreqMsec))
	}
	if d != nil && d.EddystoneUrlUrl != nil {
		eddystoneUrlUrl = types.StringValue(*d.EddystoneUrlUrl)
	}
	if d != nil && d.IbeaconAdvPower != nil {
		ibeaconAdvPower = types.Int64Value(int64(*d.IbeaconAdvPower))
	}
	if d != nil && d.IbeaconBeams != nil {
		ibeaconBeams = types.StringValue(*d.IbeaconBeams)
	}
	if d != nil && d.IbeaconEnabled != nil {
		ibeaconEnabled = types.BoolValue(*d.IbeaconEnabled)
	}
	if d != nil && d.IbeaconFreqMsec != nil {
		ibeaconFreqMsec = types.Int64Value(int64(*d.IbeaconFreqMsec))
	}
	if d != nil && d.IbeaconMajor != nil {
		ibeaconMajor = types.Int64Value(int64(*d.IbeaconMajor))
	}
	if d != nil && d.IbeaconMinor != nil {
		ibeaconMinor = types.Int64Value(int64(*d.IbeaconMinor))
	}
	if d != nil && d.IbeaconUuid != nil {
		ibeaconUuid = types.StringValue(d.IbeaconUuid.String())
	}
	if d != nil && d.Power != nil {
		power = types.Int64Value(int64(*d.Power))
	}
	if d != nil && d.PowerMode != nil {
		powerMode = types.StringValue(string(*d.PowerMode))
	}

	dataMapValue := map[string]attr.Value{
		"beacon_enabled":              beaconEnabled,
		"beacon_rate":                 beaconRate,
		"beacon_rate_mode":            beaconRateMode,
		"beam_disabled":               beamDisabled,
		"custom_ble_packet_enabled":   customBlePacketEnabled,
		"custom_ble_packet_frame":     customBlePacketFrame,
		"custom_ble_packet_freq_msec": customBlePacketFreqMsec,
		"eddystone_uid_adv_power":     eddystoneUidAdvPower,
		"eddystone_uid_beams":         eddystoneUidBeams,
		"eddystone_uid_enabled":       eddystoneUidEnabled,
		"eddystone_uid_freq_msec":     eddystoneUidFreqMsec,
		"eddystone_uid_instance":      eddystoneUidInstance,
		"eddystone_uid_namespace":     eddystoneUidNamespace,
		"eddystone_url_adv_power":     eddystoneUrlAdvPower,
		"eddystone_url_beams":         eddystoneUrlBeams,
		"eddystone_url_enabled":       eddystoneUrlEnabled,
		"eddystone_url_freq_msec":     eddystoneUrlFreqMsec,
		"eddystone_url_url":           eddystoneUrlUrl,
		"ibeacon_adv_power":           ibeaconAdvPower,
		"ibeacon_beams":               ibeaconBeams,
		"ibeacon_enabled":             ibeaconEnabled,
		"ibeacon_freq_msec":           ibeaconFreqMsec,
		"ibeacon_major":               ibeaconMajor,
		"ibeacon_minor":               ibeaconMinor,
		"ibeacon_uuid":                ibeaconUuid,
		"power":                       power,
		"power_mode":                  powerMode,
	}
	data, e := NewBleConfigValue(BleConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
