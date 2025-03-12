package resource_org_inventory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func legacyGenDeviceMap(devices *basetypes.ListValue) map[string]*DevicesValue {
	devicesMap := make(map[string]*DevicesValue)
	for _, v := range devices.Elements() {
		var dsi interface{} = v
		var dev = dsi.(DevicesValue)
		var magic = dev.Magic.ValueString()
		var mac = dev.Mac.ValueString()
		devicesMap[mac] = &dev
		if magic != "" {
			// for claimed devices
			devicesMap[magic] = &dev
		}
	}
	return devicesMap
}

func legacyProcessDevice(
	planDeviceInfo string,
	planDeviceSiteId basetypes.StringValue,
	stateMap *map[string]*DevicesValue,
) (string, string, bool) {
	var op, mac string
	var alreadyClaimed = false

	if stateDevice, ok := (*stateMap)[planDeviceInfo]; ok {
		// for already claimed devices
		op = processAction(&planDeviceSiteId, &stateDevice.SiteId)
		mac = stateDevice.Mac.ValueString()
		alreadyClaimed = true
	} else if !planDeviceSiteId.IsNull() && !planDeviceSiteId.IsUnknown() && planDeviceSiteId.ValueString() != "" {
		// for devices not claimed with the site_id set
		op = "assign"
	}

	return op, mac, alreadyClaimed
}

/*
legacyDetectDeviceInfoType detects the type of info (Claim Code or MAC Address)

parameters

	diags : *diag.Diagnostics
	deviceInfo : string
		the string to test

returns:
bool

	true if it's a Claim Code

bool

	true if it's a MAC Address

string

	device Claim Code or MAC Address
*/
func legacyDetectDeviceInfoType(device *DevicesValue) (bool, bool, string) {
	if device.Magic.ValueString() != "" {
		return true, false, *device.Magic.ValueStringPointer()
	} else if device.Mac.ValueString() != "" {
		return false, true, *device.Mac.ValueStringPointer()
	}
	return false, false, ""
}

func legacyProcessPlanedDevices(
	diags *diag.Diagnostics,
	planDevices *basetypes.ListValue,
	stateDevicesMap *map[string]*DevicesValue,
	claim *[]string,
	unassign *[]string,
	assignClaim *map[string]string,
	assign *map[string][]string,
) {
	// process devices in the plan
	// check if they must be claimed/assigned/unassigned
	for i, devPlanAttr := range planDevices.Elements() {
		var op, mac, deviceInfo string
		var alreadyClaimed, isClaimCode, isMac bool

		var dpi interface{} = devPlanAttr
		var planDevice = dpi.(DevicesValue)
		var deviceSiteId = planDevice.SiteId

		isClaimCode, isMac, deviceInfo = legacyDetectDeviceInfoType(&planDevice)
		op, mac, alreadyClaimed = legacyProcessDevice(deviceInfo, deviceSiteId, stateDevicesMap)

		if !alreadyClaimed && isClaimCode {
			*claim = append(*claim, deviceInfo)
			if op == "assign" {
				(*assignClaim)[strings.ToUpper(deviceInfo)] = deviceSiteId.ValueString()
			}
		} else if alreadyClaimed || isMac {
			if isMac {
				mac = deviceInfo
			}
			switch op {
			case "assign":
				(*assign)[deviceSiteId.ValueString()] = append((*assign)[deviceSiteId.ValueString()], mac)
			case "unassign":
				*unassign = append(*unassign, mac)
			}
		} else if !isClaimCode && !isMac {
			diags.AddError(
				"Unable to process a device in \"mist_org_inventory\"",
				fmt.Sprintf("Invalid Claim Code / MAC Address format for the device mist_org_inventory.devices[%d]. Got: \"%s\"", i, deviceInfo),
			)
		}
	}
}

func legacyProcessUnplannedDevices(
	planDevicesMap *map[string]*DevicesValue,
	stateDevices *basetypes.ListValue,
	macsToUnclaim *[]string,
) {
	unclaimedVcMembers := make(map[string]string)
	Vcs := make(map[string]string)
	// process devices in the state
	// check if they must be unclaimed
	for _, devStateAttr := range stateDevices.Elements() {
		var dsi interface{} = devStateAttr
		var device = dsi.(DevicesValue)
		var magic = device.Magic.ValueString()

		isVc := false
		if isVc = !device.VcMac.IsNull() && !device.VcMac.IsUnknown() && device.VcMac.ValueString() != ""; isVc {
			Vcs[device.VcMac.ValueString()] = device.Mac.ValueString()
		}
		// does not unclaim devices not "cloud ready" (without claim code)
		if magic != "" {
			// if we are not able to find the device in the plan based
			// on its claim code, we'll unclaim it
			if _, magicOk := (*planDevicesMap)[magic]; !magicOk { //} && !macOk {
				*macsToUnclaim = append(*macsToUnclaim, device.Mac.ValueString())
				if isVc {
					unclaimedVcMembers[device.VcMac.ValueString()] = Vcs[device.Mac.ValueString()]
				}
			}
		}
	}
	for vcMac, vcMembers := range unclaimedVcMembers {
		if len(vcMembers) == len(Vcs[vcMac]) {
			*macsToUnclaim = append(*macsToUnclaim, vcMac)
		}
	}
}

func legacyTerraformToSdk(
	stateInventory *OrgInventoryModel,
	planInventory *OrgInventoryModel,
) (
	claim []string,
	unclaim []string,
	unassign []string,
	assignClaim map[string]string,
	assign map[string][]string,
	diags diag.Diagnostics,
) {
	assignClaim = make(map[string]string)
	assign = make(map[string][]string)

	// process devices in the plan
	// check if devices must be claimed/assigned/unassigned
	stateDevicesMap := legacyGenDeviceMap(&stateInventory.Devices)
	legacyProcessPlanedDevices(&diags, &planInventory.Devices, &stateDevicesMap, &claim, &unassign, &assignClaim, &assign)

	// process devices in the state
	// check if devices must be unclaimed
	planDevicesMap := legacyGenDeviceMap(&planInventory.Devices)
	legacyProcessUnplannedDevices(&planDevicesMap, &stateInventory.Devices, &unclaim)

	return claim, unclaim, unassign, assignClaim, assign, diags
}
