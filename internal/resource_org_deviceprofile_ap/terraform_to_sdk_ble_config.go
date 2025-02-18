package resource_org_deviceprofile_ap

import (
	"github.com/google/uuid"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func bleConfigTerraformToSdk(d BleConfigValue) *models.BleConfig {

	data := models.BleConfig{}

	if d.BeaconEnabled.ValueBoolPointer() != nil {
		data.BeaconEnabled = d.BeaconEnabled.ValueBoolPointer()
	}

	if d.BeaconRate.ValueInt64Pointer() != nil {
		data.BeaconRate = models.ToPointer(int(d.BeaconRate.ValueInt64()))
	}

	if d.BeaconRateMode.ValueStringPointer() != nil {
		data.BeaconRateMode = models.ToPointer(models.BleConfigBeaconRateModeEnum(d.BeaconRateMode.ValueString()))
	}

	if !d.BeamDisabled.IsNull() && !d.BeamDisabled.IsUnknown() {
		data.BeamDisabled = misttransform.ListOfIntTerraformToSdk(d.BeamDisabled)
	}

	if d.CustomBlePacketEnabled.ValueBoolPointer() != nil {
		data.CustomBlePacketEnabled = d.CustomBlePacketEnabled.ValueBoolPointer()
	}

	if d.CustomBlePacketFrame.ValueStringPointer() != nil {
		data.CustomBlePacketFrame = d.CustomBlePacketFrame.ValueStringPointer()
	}

	if d.CustomBlePacketFreqMsec.ValueInt64Pointer() != nil {
		data.CustomBlePacketFreqMsec = models.ToPointer(int(d.CustomBlePacketFreqMsec.ValueInt64()))
	}

	if d.EddystoneUidAdvPower.ValueInt64Pointer() != nil {
		data.EddystoneUidAdvPower = models.ToPointer(int(d.EddystoneUidAdvPower.ValueInt64()))
	}

	if d.EddystoneUidBeams.ValueStringPointer() != nil {
		data.EddystoneUidBeams = d.EddystoneUidBeams.ValueStringPointer()
	}

	if d.EddystoneUidEnabled.ValueBoolPointer() != nil {
		data.EddystoneUidEnabled = d.EddystoneUidEnabled.ValueBoolPointer()
	}

	if d.EddystoneUidFreqMsec.ValueInt64Pointer() != nil {
		data.EddystoneUidFreqMsec = models.ToPointer(int(d.EddystoneUidFreqMsec.ValueInt64()))
	}

	if d.EddystoneUidInstance.ValueStringPointer() != nil {
		data.EddystoneUidInstance = d.EddystoneUidInstance.ValueStringPointer()
	}

	if d.EddystoneUidNamespace.ValueStringPointer() != nil {
		data.EddystoneUidNamespace = d.EddystoneUidNamespace.ValueStringPointer()
	}

	if d.EddystoneUrlAdvPower.ValueInt64Pointer() != nil {
		data.EddystoneUrlAdvPower = models.ToPointer(int(d.EddystoneUrlAdvPower.ValueInt64()))
	}

	if d.EddystoneUrlBeams.ValueStringPointer() != nil {
		data.EddystoneUrlBeams = d.EddystoneUrlBeams.ValueStringPointer()
	}

	if d.EddystoneUrlEnabled.ValueBoolPointer() != nil {
		data.EddystoneUrlEnabled = d.EddystoneUrlEnabled.ValueBoolPointer()
	}

	if d.EddystoneUrlFreqMsec.ValueInt64Pointer() != nil {
		data.EddystoneUrlFreqMsec = models.ToPointer(int(d.EddystoneUrlFreqMsec.ValueInt64()))
	}

	if d.EddystoneUrlUrl.ValueStringPointer() != nil {
		data.EddystoneUrlUrl = d.EddystoneUrlUrl.ValueStringPointer()
	}

	if d.IbeaconAdvPower.ValueInt64Pointer() != nil {
		data.IbeaconAdvPower = models.ToPointer(int(d.IbeaconAdvPower.ValueInt64()))
	}

	if d.IbeaconBeams.ValueStringPointer() != nil {
		data.IbeaconBeams = d.IbeaconBeams.ValueStringPointer()
	}

	if d.IbeaconEnabled.ValueBoolPointer() != nil {
		data.IbeaconEnabled = d.IbeaconEnabled.ValueBoolPointer()
	}

	if d.IbeaconFreqMsec.ValueInt64Pointer() != nil {
		data.IbeaconFreqMsec = models.ToPointer(int(d.IbeaconFreqMsec.ValueInt64()))
	}

	if d.IbeaconMajor.ValueInt64Pointer() != nil {
		data.IbeaconMajor = models.ToPointer(int(d.IbeaconMajor.ValueInt64()))
	}

	if d.IbeaconMinor.ValueInt64Pointer() != nil {
		data.IbeaconMinor = models.ToPointer(int(d.IbeaconMinor.ValueInt64()))
	}

	if d.IbeaconUuid.ValueStringPointer() != nil {
		data.IbeaconUuid = models.ToPointer(uuid.MustParse(d.IbeaconUuid.ValueString()))
	}

	if d.Power.ValueInt64Pointer() != nil {
		data.Power = models.ToPointer(int(d.Power.ValueInt64()))
	}

	if d.PowerMode.ValueStringPointer() != nil {
		data.PowerMode = models.ToPointer(models.BleConfigPowerModeEnum(d.PowerMode.ValueString()))
	}

	return &data
}
