package resource_org_inventory

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, devices_plan *basetypes.ListValue, devices_state *basetypes.ListValue) ([]string, []string, []string, map[string]string, map[string][]string, diag.Diagnostics) {
	var diags diag.Diagnostics
	var claim []string
	var unclaim []string
	var unassign []string
	assign_claim := make(map[string]string)
	assign := make(map[string][]string)

	for _, dev_plan_attr := range devices_plan.Elements() {
		var dpi interface{} = dev_plan_attr
		var dev_plan = dpi.(DevicesValue)
		var op string = ""
		var device_mac = ""
		already_claimed := false
		for _, dev_state_attr := range devices_state.Elements() {
			var dsi interface{} = dev_state_attr
			var dev_state = dsi.(DevicesValue)
			if dev_plan.Magic == dev_state.Magic {
				already_claimed = true
				// Site ID not in state but planed > must be assigned
				if dev_state.SiteId.ValueStringPointer() == nil && dev_state.SiteId.ValueString() == "" &&
					dev_plan.SiteId.ValueStringPointer() != nil && dev_plan.SiteId.ValueString() != "" {
					op = "assign"
					device_mac = dev_state.Mac.ValueString()
					// Site ID in state but not planed > must be unassigned
				} else if dev_state.SiteId.ValueStringPointer() != nil && dev_state.SiteId.ValueString() != "" &&
					dev_plan.SiteId.ValueStringPointer() == nil && dev_plan.SiteId.ValueString() == "" {
					op = "unassign"
					device_mac = dev_state.Mac.ValueString()
					// State Site ID != Plan Site ID > must be reassigned
				} else if dev_state.SiteId.ValueStringPointer() != nil && dev_state.SiteId.ValueString() != "" &&
					dev_plan.SiteId.ValueStringPointer() != nil && dev_plan.SiteId.ValueString() != "" &&
					!dev_state.SiteId.Equal(dev_plan.SiteId) {
					device_mac = dev_state.Mac.ValueString()
					op = "reassign"
				}
			}
		}
		if !already_claimed {
			claim = append(claim, dev_plan.Magic.ValueString())
			if dev_plan.SiteId.ValueStringPointer() != nil && dev_plan.SiteId.ValueString() != "" {
				assign_claim[dev_plan.Magic.ValueString()] = dev_plan.SiteId.ValueString()
			}
		}
		switch op {
		case "assign":
			assign[dev_plan.SiteId.ValueString()] = append(assign[dev_plan.SiteId.ValueString()], device_mac)
		case "reassign":
			assign[dev_plan.SiteId.ValueString()] = append(assign[dev_plan.SiteId.ValueString()], device_mac)
		case "unassign":
			unassign = append(unassign, device_mac)
		}
	}

	for _, dev_state_attr := range devices_state.Elements() {
		var dsi interface{} = dev_state_attr
		var dev_state = dsi.(DevicesValue)
		to_unclaim := true
		for _, dev_plan_attr := range devices_plan.Elements() {
			var dpi interface{} = dev_plan_attr
			var dev_plan = dpi.(DevicesValue)
			if dev_state.Magic.Equal(dev_plan.Magic) {
				to_unclaim = false
			}
		}
		if to_unclaim {
			unclaim = append(unclaim, dev_state.Serial.ValueString())
		}
	}

	return claim, unclaim, unassign, assign_claim, assign, diags
}
