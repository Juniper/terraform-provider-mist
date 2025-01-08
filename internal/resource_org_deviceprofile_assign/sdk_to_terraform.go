package resource_org_deviceprofile_assign

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(
	ctx context.Context,
	macsState *basetypes.SetValue,
	macsToAssign *models.MacAddresses,
	macsToUnassign *models.MacAddresses,
	macsAssigned []string,
	macsUnassigned []string,
) (*basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var list_value []attr.Value
	var list_value_type attr.Type = basetypes.StringType{}

	// remove unassigned devices from state
	for _, mac := range macsState.Elements() {
		var s_interface interface{} = mac
		s := s_interface.(basetypes.StringValue)
		unassignRequested := macInSlice(s.ValueString(), macsToUnassign.Macs)
		unassignApplied := macInSlice(s.ValueString(), macsUnassigned)
		// if unassignRequested && unassignApplied: that's expected, do not add the device to the state list, do not raise a warning
		// if unassignRequested && !unassignApplied: that's not expected, do not add the device to the state list but raise a warning
		if unassignRequested && !unassignApplied {
			diags.AddWarning(
				"Unable to Unassign the device profile from a device",
				fmt.Sprintf(
					"The device profile has not been unassigned from the MAC Address %s.\n"+
						"This is mostly because the device profile has been manually unassigned from the device. Please check on the Mist UI.\n"+
						"The state has been updated according to the plan",
					mac,
				),
			)
		} else if !unassignRequested {
			list_value = append(list_value, mac)
		}
	}

	// add assigned devices to state
	for _, mac := range macsToAssign.Macs {
		if macInSlice(mac, macsAssigned) {
			list_value = append(list_value, types.StringValue(mac))
		} else {
			diags.AddError(
				"Error Assigning the device profile to the device",
				fmt.Sprintf(
					"The device profile has not been assigned to the MAC Address %s.\n"+
						"Please check the provided MAC Address is valid and belong to your organization.",
					mac,
				),
			)
		}
	}

	new_macs_state, e := types.SetValue(list_value_type, list_value)
	diags.Append(e...)

	return &new_macs_state, diags
}

func macInSlice(mac string, list []string) bool {
	for _, b := range list {
		if b == mac {
			return true
		}
	}
	return false
}

func macInState(mac string, state *basetypes.SetValue) bool {
	for _, b := range state.Elements() {
		var s_interface interface{} = b
		s := s_interface.(basetypes.StringValue)
		if s.ValueString() == mac {
			return true
		}
	}
	return false
}
