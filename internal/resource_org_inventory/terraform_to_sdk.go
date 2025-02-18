package resource_org_inventory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func processAction(planSiteId *basetypes.StringValue, stateSiteId *basetypes.StringValue) string {
	/*
		Function to define the required action for a specific device (assign/unassign/nothing)

		parameters:
			planSiteId : *basetypes.StringValue
				planed siteId for the device
			stateSiteId : *basetypes.StringValue
				planed siteId for the device

		returns:
			string
				the op to apply to the device (assign/unassign/nothing)
	*/
	if stateSiteId.ValueString() == planSiteId.ValueString() {
		return ""
	} else if planSiteId.IsNull() || planSiteId.IsUnknown() || planSiteId.ValueString() == "" {
		// Planned Site ID is not set > must be unassigned
		return "unassign"
	} else {
		// Planned Site ID is set > must be assigned or reassiogned
		return "assign"
	}
}

func findDeviceInState(
	planDeviceInfo string,
	planDeviceSiteId basetypes.StringValue,
	stateMap *map[string]InventoryValue,
) (string, string, bool) {
	/*
		Function to find a device in the list coming from the Mist Inventory based on the Claim Code
		or the MAC Address

		parameters:
			planDeviceInfo : string
				the planed device Claim Code or MAC Address
			planDeviceSiteId : basetypes.StringValue
				the planed device Site ID
			stateMap : *map[string]InventoryValue
				map of the devices in the Mist inventory. The key may be the device Claim Code or MAC address
				(depeending on the value type in planDeviceInfo) and the value is DeviceValue

		returns:
			string
				the op to apply to the device (assign/unassign/nothing)
			string
				the device MAC Address (required for assign/unassign ops)
			bool
				if the device is already claimed (only used when planDeviceInfo is a claim code)
	*/
	var op, mac string
	var alreadyClaimed = false

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

func processPlanedDevices(
	diags *diag.Diagnostics,
	planDevices *basetypes.MapValue,
	stateDevicesMap *map[string]InventoryValue,
	claim *[]string,
	unassign *[]string,
	assignClaim *map[string]string,
	assign *map[string][]string,
) {
	/*
		Function to process the planed devices and detect which type of action should be applied. Depending
		on the required action, the device will be added to one of the required list

		parameters:
			diags: *diag.Diagnostics
			planDevices : *basetypes.MapValue
				map of devices in the plan. Key is the device Claim Code or MAC Address, Value is a DeviceValue Nested
				Object with the SiteId, the UnclaimWhenDestroyed bit and the information retrieved from the Mist Inventory
			stateDevicesMap : *map[string]InventoryValue
				map of devices in the plan. Key is the device Claim Code or MAC Address, Value is a DeviceValue Nested
				Object with the SiteId, the UnclaimWhenDestroyed bit and the information retrieved from the Mist Inventory
			claim : *[]string
				list of claim codes (string) that must be claimed to the Mist Org
			unassign : *[]string
				list of MAC Address (string) that must be unassigned from Mist Sites
			assignClaim : *map[string]string
				map of  ClaimCodes / SiteId of the devices that must be claimed then assigned to a site. This is required
				because we don't have the device MAC address at this time (we only have the claim code, the MAC Addresss
				which is required for the "assign" op will be known after the claim)
				the key is the device Claim Code
				the value is the site id where the device must be assigned to after the claim
			assign : *map[string][]string
				map of siteId / list of MAC address (string) that must be assigned to a site
				the key is the siteId where the device(s) must be claimed to
				the value is a list of MAC Address that must be assigned to the site
	*/
	for deviceInfo, d := range planDevices.Elements() {
		var op, mac string
		var alreadyClaimed bool

		var di interface{} = d
		var planDevice = di.(InventoryValue)
		var deviceSiteId = planDevice.SiteId

		op, mac, alreadyClaimed = findDeviceInState(deviceInfo, deviceSiteId, stateDevicesMap)
		isClaimCode, isMac := DetectDeviceInfoType(diags, deviceInfo)
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
				fmt.Sprintf("Invalid Claim Code / MAC Address format. Got: \"%s\"", deviceInfo),
			)
		}
	}
}

func processUnplanedDevices(
	planDevicesMap *map[string]InventoryValue,
	stateDevices *basetypes.MapValue,
	unclaim *[]string,
) {
	/*
		Function to process the planed devices and detect which devices must be unclaimed

		parameters:
			planDevicesMap : *map[string]DeviceValue
				map of the devices in the Plan. The key may be the device Claim Code or MAC address
				(depending on the value type in planDeviceInfo) and the value is DeviceValue
			stateDevices : *basetypes.MapValue
				map of devices in the state (claimed / managed by the provider). Key is the device Claim Code
				or MAC Address, Value is a Nested Object with the SiteId and the UnclaimWhenDestroyed bit
			unclaim : *[]string
				list of serial numbers (serial) that must be unclaim from the Mist Inventory
	*/

	for deviceInfo, d := range stateDevices.Elements() {
		var di interface{} = d
		var device = di.(InventoryValue)
		var unclaimWhenDestroyed = device.UnclaimWhenDestroyed.ValueBool()

		if _, ok := (*planDevicesMap)[strings.ToUpper(deviceInfo)]; !ok && unclaimWhenDestroyed {
			*unclaim = append(*unclaim, device.Serial.ValueString())
		}
	}
}

func mapTerraformToSdk(stateInventory *OrgInventoryModel, planInventory *OrgInventoryModel) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	// var knownDevice
	var claim []string
	var unclaim []string
	var unassign []string
	assignClaim := make(map[string]string)
	assign := make(map[string][]string)

	// process devices in the plan
	// check if devices must be claimed/assigned/unassigned
	stateDevicesMap := GenDeviceMap(&stateInventory.Inventory)
	processPlanedDevices(&diags, &planInventory.Inventory, &stateDevicesMap, &claim, &unassign, &assignClaim, &assign)

	// process devices in the state
	// check if devices must be unclaimed
	planDevicesMap := GenDeviceMap(&planInventory.Inventory)
	processUnplanedDevices(&planDevicesMap, &stateInventory.Inventory, &unclaim)

	return claim, unclaim, unassign, assignClaim, assign, diags
}

func TerraformToSdk(stateInventory *OrgInventoryModel, planInventory *OrgInventoryModel) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {

	if !planInventory.Devices.IsNull() && !planInventory.Devices.IsUnknown() {
		return legacyTerraformToSdk(&stateInventory.Devices, &planInventory.Devices)
	} else {
		return mapTerraformToSdk(stateInventory, planInventory)
	}
}

func DeleteOrgInventory(stateInventory *OrgInventoryModel) ([]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var unclaim []string

	if !stateInventory.Devices.IsNull() && !stateInventory.Devices.IsUnknown() {
		for _, d := range stateInventory.Devices.Elements() {
			var di interface{} = d
			var device = di.(DevicesValue)
			var unclaimWhenDestroyed = device.UnclaimWhenDestroyed.ValueBool()

			if unclaimWhenDestroyed {
				unclaim = append(unclaim, device.Serial.ValueString())
			}
		}
	} else {
		for _, d := range stateInventory.Inventory.Elements() {
			var di interface{} = d
			var device = di.(InventoryValue)
			var unclaimWhenDestroyed = device.UnclaimWhenDestroyed.ValueBool()

			if unclaimWhenDestroyed {
				unclaim = append(unclaim, device.Serial.ValueString())
			}
		}
	}

	return unclaim, diags
}
