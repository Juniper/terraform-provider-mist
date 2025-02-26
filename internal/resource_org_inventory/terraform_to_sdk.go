package resource_org_inventory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func processAction(planSiteId *basetypes.StringValue, stateSiteId *basetypes.StringValue) (op string) {
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
	planDeviceSiteId *basetypes.StringValue,
	stateDevice *InventoryValue,
) (op string, mac string, alreadyClaimed bool) {
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
	alreadyClaimed = false

	if stateDevice != nil && !stateDevice.IsNull() && !stateDevice.IsUnknown() {
		// for already claimed devices
		op = processAction(planDeviceSiteId, &stateDevice.SiteId)
		mac = stateDevice.Mac.ValueString()
		alreadyClaimed = true
	} else if !planDeviceSiteId.IsNull() && !planDeviceSiteId.IsUnknown() && planDeviceSiteId.ValueString() != "" {
		// for devices not claimed with the site_id set
		op = "assign"
	}

	return op, mac, alreadyClaimed
}

func vcMembersAssignmentSave(
	deviceInfo *string,
	planDevice *InventoryValue,
	stateDevice *InventoryValue,
	vcMacAssignments map[string]map[string][]string,
) (isVc bool, vcMac string) {
	/*
		Function to check if the devices (mostly switches) are part of a Virtual Chassis.
		If true, the rest of the process will be applied to the Virtual Chassis MAC,
		otherwise, the process will continue with the device MAC Address.
		This function is also validating all the VC members are assigned to the same site

		parameters:
			diags: *diag.Diagnostics
			deviceInfo: *string
				the device claim code / MAC address (from the mist_org_inventory resource)
			planDevice : *InventoryValue
				device in the Plan.
			stateDevice : *InventoryValue
				device in the State.
			vcMacToSiteIdMap : *map[string]string
				a map of string to track which device/virtual device is assigned to which site
				(used to validate all the VC members are assigned to the same site)

		returns:
			bool:
				true if the device is a VC member
			string:
				if the device is a VC member, the MAC address of the VC
	*/
	isVc = false
	if stateDevice != nil && !stateDevice.VcMac.IsNull() && !stateDevice.VcMac.IsUnknown() && stateDevice.VcMac.ValueString() != "" {
		isVc = true
		vcMac = stateDevice.VcMac.ValueString()
		planSiteId := planDevice.SiteId.ValueString()

		if planSiteId == "" {
			planSiteId = "00000000-0000-0000-0000-000000000000"
		}

		if _, vcMacExists := vcMacAssignments[vcMac]; !vcMacExists {
			vcMacAssignments[vcMac] = make(map[string][]string)
			vcMacAssignments[vcMac][planSiteId] = []string{*deviceInfo}
		} else if _, siteIdExists := vcMacAssignments[vcMac][planSiteId]; !siteIdExists {
			vcMacAssignments[vcMac][planSiteId] = []string{*deviceInfo}
		} else {
			vcMacAssignments[vcMac][planSiteId] = append(vcMacAssignments[vcMac][planSiteId], *deviceInfo)
		}
	}
	return isVc, vcMac
}

func vcMembersAssignmentCheck(
	diags *diag.Diagnostics,
	vcMacAssignments *map[string]map[string][]string,
) {
	for vcMac, vcSiteMembers := range *vcMacAssignments {
		if len(vcSiteMembers) > 1 {
			errorMessage := ""
			for siteId, members := range vcSiteMembers {
				if siteId == "00000000-0000-0000-0000-000000000000" {
					errorMessage = errorMessage + "\nunassigned:"
				} else {
					errorMessage = errorMessage + fmt.Sprintf("\nsite_id %s:", siteId)
				}
				for _, m := range members {
					errorMessage = errorMessage + fmt.Sprintf("\n\t- mist_org_inventory.inventory[%s]", m)
				}
			}
			diags.AddError(
				"Unable to process a device in \"mist_org_inventory\"",
				fmt.Sprintf(
					"The devices part of the Virtual Chassis %s are currently assigned to different sites:%s"+
						"\nPlease set the same site_id to all the virtual chassis members to it a site, "+
						"or unset it to unassign the virtual chassis from the site.",
					vcMac, errorMessage,
				),
			)
		}
	}
}

func processPlanedDevices(
	diags *diag.Diagnostics,
	planDevices *basetypes.MapValue,
	stateDevicesMap *map[string]*InventoryValue,
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
	var vcMacAssignments = make(map[string]map[string][]string)
	for deviceInfo, d := range planDevices.Elements() {
		var op, mac string
		var alreadyClaimed bool

		var di interface{} = d
		var planDevice = di.(InventoryValue)
		var deviceSiteId = planDevice.SiteId
		stateDevice := (*stateDevicesMap)[strings.ToUpper(deviceInfo)]

		// mac will be empty if the device is not already in the state
		op, mac, alreadyClaimed = findDeviceInState(&deviceSiteId, stateDevice)
		isClaimCode, isMac := DetectDeviceInfoType(diags, deviceInfo)
		isVc, vcMac := vcMembersAssignmentSave(&deviceInfo, &planDevice, stateDevice, vcMacAssignments)
		if !alreadyClaimed && isClaimCode {
			*claim = append(*claim, deviceInfo)
			if op == "assign" {
				(*assignClaim)[strings.ToUpper(deviceInfo)] = deviceSiteId.ValueString()
			}
		} else if alreadyClaimed || isMac {
			if isVc {
				mac = vcMac
			} else if isMac {
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
	vcMembersAssignmentCheck(diags, &vcMacAssignments)
}

func processUnplanedDevices(
	planDevicesMap *map[string]*InventoryValue,
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
	unclaimedVcMembers := make(map[string]string)
	Vcs := make(map[string]string)

	for deviceInfo, d := range stateDevices.Elements() {
		var di interface{} = d
		var device = di.(InventoryValue)
		var unclaimWhenDestroyed = device.UnclaimWhenDestroyed.ValueBool()
		isVc := false
		if isVc = !device.VcMac.IsNull() && !device.VcMac.IsUnknown() && device.VcMac.ValueString() != ""; isVc {
			Vcs[device.VcMac.ValueString()] = device.Mac.ValueString()
		}

		if _, ok := (*planDevicesMap)[deviceInfo]; !ok && unclaimWhenDestroyed {
			*unclaim = append(*unclaim, device.Serial.ValueString())
			if isVc {
				unclaimedVcMembers[device.VcMac.ValueString()] = Vcs[device.Mac.ValueString()]
			}
		}
	}
	for vcMac, vcMembers := range unclaimedVcMembers {
		if len(vcMembers) == len(Vcs[vcMac]) {
			*unclaim = append(*unclaim, vcMac)
		}
	}
}

func mapTerraformToSdk(
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
	stateDevicesMap := GenDeviceMap(&stateInventory.Inventory)
	processPlanedDevices(&diags, &planInventory.Inventory, &stateDevicesMap, &claim, &unassign, &assignClaim, &assign)

	// process devices in the state
	// check if devices must be unclaimed
	planDevicesMap := GenDeviceMap(&planInventory.Inventory)
	processUnplanedDevices(&planDevicesMap, &stateInventory.Inventory, &unclaim)

	return claim, unclaim, unassign, assignClaim, assign, diags
}

func TerraformToSdk(
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

	if !planInventory.Devices.IsNull() && !planInventory.Devices.IsUnknown() {
		return legacyTerraformToSdk(&stateInventory.Devices, &planInventory.Devices)
	} else {
		return mapTerraformToSdk(stateInventory, planInventory)
	}
}

func DeleteOrgInventory(stateInventory *OrgInventoryModel) (unclaim []string, diags diag.Diagnostics) {
	if !stateInventory.Devices.IsNull() && !stateInventory.Devices.IsUnknown() {
		planDevicesMap := make(map[string]*DevicesValue)
		legacyProcessUnplanedDevices(&planDevicesMap, &stateInventory.Devices, &unclaim)
	} else {
		planDevicesMap := make(map[string]*InventoryValue)
		processUnplanedDevices(&planDevicesMap, &stateInventory.Inventory, &unclaim)
	}

	return unclaim, diags
}
