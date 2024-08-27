package resource_org_inventory

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func processAction(planDevice *DevicesValue, stateDevice *DevicesValue, assign map[string][]string, unassign []string) (map[string][]string, []string) {

	// Site ID not in state but planed > must be assigned
	if !stateDevice.SiteId.Equal(planDevice.SiteId) && planDevice.SiteId.ValueString() != "" {
		assign[planDevice.SiteId.ValueString()] = append(assign[planDevice.SiteId.ValueString()], stateDevice.Mac.ValueString())
		// Site ID in state but not planed > must be unassigned

	} else if !stateDevice.SiteId.Equal(planDevice.SiteId) && planDevice.SiteId.ValueString() == "" {
		unassign = append(unassign, stateDevice.Mac.ValueString())
		// State Site ID != Plan Site ID > must be reassigned

	}
	return assign, unassign
}

func TerraformToSdk(ctx context.Context, devices_plan *basetypes.ListValue, devices_state *basetypes.ListValue) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var claim []string
	var unclaim []string
	var unassign []string
	assign_claim := make(map[string]string)
	assign := make(map[string][]string)

	stateMap := genDeviceMap(devices_state)
	planMap := genDeviceMap(devices_plan)

	for _, dev_plan_attr := range devices_plan.Elements() {
		var dpi interface{} = dev_plan_attr
		var dev_plan = dpi.(DevicesValue)

		var op string
		var already_claimed bool = false
		var magic string = strings.ToUpper(dev_plan.Magic.ValueString())
		var mac string = strings.ToUpper(dev_plan.Mac.ValueString())

		if dev_state, ok := stateMap[magic]; ok {
			// for already claimed devices
			already_claimed = true
			assign, unassign = processAction(&dev_plan, &dev_state, assign, unassign)
			tflog.Warn(ctx, "FOUND MAGIC "+magic+"with op "+op)
		} else if dev_state, ok = stateMap[mac]; ok {
			// for already adopted devices
			already_claimed = true
			assign, unassign = processAction(&dev_plan, &dev_state, assign, unassign)
			tflog.Warn(ctx, "FOUND MAC "+mac+"with op "+op)
		} else {
			tflog.Warn(ctx, "NOT FOUND MAC "+mac+" OR MAGIC "+magic)

		}

		if !already_claimed && !dev_plan.Magic.IsNull() && !dev_plan.Magic.IsUnknown() {
			claim = append(claim, dev_plan.Magic.ValueString())
			if dev_plan.SiteId.ValueStringPointer() != nil && dev_plan.SiteId.ValueString() != "" {
				assign_claim[dev_plan.Magic.ValueString()] = dev_plan.SiteId.ValueString()
			}
		}

	}

	for _, dev_state_attr := range devices_state.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(DevicesValue)
		var magic string = strings.ToUpper(dev_state.Magic.ValueString())
		if _, ok := planMap[magic]; magic != "" && !ok {
			unclaim = append(unclaim, dev_state.Serial.ValueString())
		}
	}

	return claim, unclaim, unassign, assign_claim, assign, diags
}
