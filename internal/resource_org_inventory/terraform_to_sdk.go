package resource_org_inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func genDeviceMap(devices *basetypes.ListValue) (map[string]*DevicesValue, map[string]*DevicesValue) {
	devices_map_magic := make(map[string]*DevicesValue)
	devices_map_mac := make(map[string]*DevicesValue)
	for _, v := range devices.Elements() {
		var dsi interface{} = v
		var dev = dsi.(DevicesValue)
		var magic string = strings.ReplaceAll(strings.ToUpper(dev.Magic.ValueString()), "-", "")
		var mac string = strings.ToLower(dev.Mac.ValueString())
		devices_map_mac[mac] = &dev
		if magic != "" {
			// for claimed devices
			devices_map_magic[magic] = &dev
		}
	}
	return devices_map_magic, devices_map_mac
}

func processAction(planDevice *DevicesValue, stateDevice *DevicesValue) string {

	if stateDevice != nil && stateDevice.SiteId.Equal(planDevice.SiteId) {
		return ""
	} else if planDevice.SiteId.ValueStringPointer() == nil {
		// Planned Site ID nil > must be unassigned
		return "unassign"
		//*unassign = append(*unassign, stateDevice.Mac.ValueString())
	} else {
		// Planned Site ID not not nil > must be assigned or reassiogned
		return "assign"
		// (*assign)[planDevice.SiteId.ValueString()] = append((*assign)[planDevice.SiteId.ValueString()], stateDevice.Mac.ValueString())
	}
}

func processClaimedDevice(planDevice *DevicesValue, stateMap *map[string]*DevicesValue) (string, string, string) {
	var op, siteId, mac string
	var magic string = strings.ToUpper(planDevice.Magic.ValueString())

	if stateDevice, ok := (*stateMap)[magic]; ok {
		// for already claimed devices
		op = processAction(planDevice, stateDevice)
		mac = stateDevice.Mac.ValueString()
	} else if planDevice.SiteId.ValueString() != "" {
		// for devices not claimed with the site_id set
		op = "assign"
	}
	siteId = planDevice.SiteId.ValueString()

	return op, mac, siteId
}

func processAdoptedDevice(planDevice *DevicesValue, stateMap *map[string]*DevicesValue) (string, string, string) {
	var op, siteId string
	var mac string = strings.ToLower(planDevice.Mac.ValueString())

	if stateDevice, ok := (*stateMap)[mac]; ok {
		// or adopted devices that are already known
		op = processAction(planDevice, stateDevice)
	} else if planDevice.SiteId.ValueString() != "" {
		// for adopted devices that are unknown and have the site_id set
		op = "assign"
	}
	siteId = planDevice.SiteId.ValueString()

	return op, planDevice.Mac.ValueString(), siteId
}

func TerraformToSdk(ctx context.Context, devices_plan *basetypes.ListValue, devices_state *basetypes.ListValue) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var claim []string
	var unclaim []string
	var unassign []string
	assign_claim := make(map[string]string)
	assign := make(map[string][]string)

	devices_plan_magic, devices_plan_mac := genDeviceMap(devices_plan)
	devices_state_magic, devices_state_mac := genDeviceMap(devices_state)

	// process devices in the plan
	// check if they must be claimed/assigned/unassigned
	for i, dev_plan_attr := range devices_plan.Elements() {
		var dpi interface{} = dev_plan_attr
		var dev_plan = dpi.(DevicesValue)
		var op, mac, siteId string

		if dev_plan.Magic.ValueString() != "" {
			op, mac, siteId = processClaimedDevice(&dev_plan, &devices_state_magic)
		} else if dev_plan.Mac.ValueString() != "" {
			op, mac, siteId = processAdoptedDevice(&dev_plan, &devices_state_mac)
		} else {
			diags.AddError(
				"Unable to process device in \"mist_org_inventory\"",
				fmt.Sprintf("Claim Code and MAC Address not found for the device mist_org_inventory.devices[%s]", string(i)),
			)
		}

		// if the device is not claimed (unknown MAC Address), we will need to claim it to get the MAC address
		if mac == "" && dev_plan.Magic.ValueString() != "" {
			claim = append(claim, dev_plan.Magic.ValueString())
			if siteId != "" {
				assign_claim[dev_plan.Magic.ValueString()] = siteId
			}
		} else {
			switch op {
			case "assign":
				assign[siteId] = append(assign[siteId], mac)
			case "unassign":
				unassign = append(unassign, mac)
			}
		}
	}

	// process devices in the state
	// check if they must be unclaimed
	for _, dev_state_attr := range devices_state.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(DevicesValue)
		var magic string = strings.ToUpper(dev_state.Magic.ValueString())
		var mac string = strings.ToLower(dev_state.Mac.ValueString())
		// does not unclaim devices not "cloud ready" (without claim code)
		if magic != "" {
			_, magic_ok := devices_plan_magic[magic]
			_, mac_ok := devices_plan_mac[mac]
			// if we are not able to find the device in the plan based
			// on its claim code or its mac, we'll unclaim it
			if !magic_ok && !mac_ok {
				unclaim = append(unclaim, dev_state.Serial.ValueString())
			}
		}
	}

	return claim, unclaim, unassign, assign_claim, assign, diags
}
