package resource_org_inventory

import (
	"context"
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
		var magic string = strings.ReplaceAll(strings.ToUpper(dev.Magic.ValueString()), "-", "")
		var mac string = strings.ToUpper(dev.Mac.ValueString())
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
	var alreadyClaimed bool = false

	if stateDevice, ok := (*stateMap)[strings.ToUpper(planDeviceInfo)]; ok {
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

func legacyDetectDeviceInfoType(device *DevicesValue) (bool, bool, string) {
	/*
		Function to detect the type of info (Claim Code or MAC Address)

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
	for i, dev_plan_attr := range planDevices.Elements() {
		var op, mac, deviceInfo string
		var alreadyClaimed, isClaimCode, isMac bool

		var dpi interface{} = dev_plan_attr
		var planDevice = dpi.(DevicesValue)
		var deviceSiteId = planDevice.SiteId

		isClaimCode, isMac, deviceInfo = legacyDetectDeviceInfoType(&planDevice)
		op, mac, alreadyClaimed = legacyProcessDevice(deviceInfo, deviceSiteId, stateDevicesMap)

		if !alreadyClaimed && isClaimCode {
			(*claim) = append((*claim), deviceInfo)
			if op == "assign" {
				(*assignClaim)[strings.ToUpper(deviceInfo)] = deviceSiteId.ValueString()
			}
			// } else if !alreadyClaimed && isMac {
			// 	diags.AddError(
			// 		"Unable to process the \"org_inventory\" resource",
			// 		fmt.Sprintf("unable to find a device with the MAC Address %s in the Org Inventory", deviceInfo),
			// 	)
		} else if alreadyClaimed || isMac {
			if isMac {
				mac = deviceInfo
			}
			switch op {
			case "assign":
				(*assign)[deviceSiteId.ValueString()] = append((*assign)[deviceSiteId.ValueString()], mac)
			case "unassign":
				(*unassign) = append((*unassign), mac)
			}
		} else if !isClaimCode && !isMac {
			diags.AddError(
				"Unable to process a device in \"mist_org_inventory\"",
				fmt.Sprintf("Invalid Claim Code / MAC Address format for the device mist_org_inventory.devices[%d]. Got: \"%s\"", i, deviceInfo),
			)
		}
	}
}

func legacyProcessUnplanedDevices(
	stateDevices *basetypes.ListValue,
	planDevicesMap *map[string]*DevicesValue,
	unclaim *[]string,
) {
	// process devices in the state
	// check if they must be unclaimed
	for _, dev_state_attr := range stateDevices.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(DevicesValue)
		var magic string = strings.ToUpper(dev_state.Magic.ValueString())
		var mac string = strings.ToLower(dev_state.Mac.ValueString())
		// does not unclaim devices not "cloud ready" (without claim code)
		if magic != "" {
			_, magic_ok := (*planDevicesMap)[magic]
			_, mac_ok := (*planDevicesMap)[mac]
			// if we are not able to find the device in the plan based
			// on its claim code or its mac, we'll unclaim it
			if !magic_ok && !mac_ok {
				(*unclaim) = append((*unclaim), dev_state.Serial.ValueString())
			}
		}
	}
}

func legacyTerraformToSdk(ctx context.Context, stateDevices *basetypes.ListValue, planDevices *basetypes.ListValue) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var claim []string
	var unclaim []string
	var unassign []string
	assignClaim := make(map[string]string)
	assign := make(map[string][]string)

	// process devices in the plan
	// check if devices must be claimed/assigned/unassigned
	stateDevicesMap := legacyGenDeviceMap(stateDevices)
	legacyProcessPlanedDevices(&diags, planDevices, &stateDevicesMap, &claim, &unassign, &assignClaim, &assign)

	// process devices in the state
	// check if devices must be unclaimed
	planDevicesMap := legacyGenDeviceMap(planDevices)
	legacyProcessUnplanedDevices(stateDevices, &planDevicesMap, &unclaim)

	return claim, unclaim, unassign, assignClaim, assign, diags
}
