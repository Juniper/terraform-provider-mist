package resource_org_inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func genDeviceMap(devices *basetypes.ListValue) map[string]DevicesValue {
	deviceMap := make(map[string]DevicesValue)
	for _, v := range devices.Elements() {
		var dsi interface{} = v
		var dev = dsi.(DevicesValue)
		var magic string = strings.ToUpper(dev.Magic.ValueString())
		var mac string = strings.ToUpper(dev.Mac.ValueString())
		if magic != "" {
			// for claimed devices
			deviceMap[magic] = dev
		} else {
			// for adopted devices
			deviceMap[mac] = dev
		}
	}
	return deviceMap
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

func processClaimedDevice(planDevice *DevicesValue, stateMap *map[string]DevicesValue) (string, string, string) {
	var op, siteId, mac string
	var magic string = strings.ToUpper(planDevice.Magic.ValueString())

	if stateDevice, ok := (*stateMap)[magic]; ok {
		// for already claimed devices
		op = processAction(planDevice, &stateDevice)
		mac = stateDevice.Mac.ValueString()
	} else if planDevice.SiteId.ValueString() != "" {
		// for devices not claimed with the site_id set
		op = "assign"
	}
	siteId = planDevice.SiteId.ValueString()

	return op, mac, siteId
}

func processAdoptedDevice(planDevice *DevicesValue, stateMap *map[string]DevicesValue) (string, string, string) {
	var op, siteId string
	var mac string = strings.ToUpper(planDevice.Mac.ValueString())

	if stateDevice, ok := (*stateMap)[mac]; ok {
		// or adopted devices that are already known
		op = processAction(planDevice, &stateDevice)
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

	planMap := genDeviceMap(devices_plan)
	stateMap := genDeviceMap(devices_state)

	for i, dev_plan_attr := range devices_plan.Elements() {
		var dpi interface{} = dev_plan_attr
		var dev_plan = dpi.(DevicesValue)
		var op, mac, siteId string

		if dev_plan.Magic.ValueString() != "" {
			op, mac, siteId = processClaimedDevice(&dev_plan, &stateMap)
		} else if dev_plan.Mac.ValueString() != "" {
			op, mac, siteId = processAdoptedDevice(&dev_plan, &stateMap)
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

	for _, dev_state_attr := range devices_state.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(DevicesValue)
		var magic string = strings.ToUpper(dev_state.Magic.ValueString())
		var mac string = strings.ToUpper(dev_state.Mac.ValueString())
		if _, ok := planMap[magic]; magic != "" && !ok {
			unclaim = append(unclaim, dev_state.Serial.ValueString())
		} else if _, ok = planMap[mac]; magic == "" && !ok {
			unassign = append(unassign, dev_state.Mac.ValueString())
		}
	}

	return claim, unclaim, unassign, assign_claim, assign, diags
}
