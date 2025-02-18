package resource_org_deviceprofile_assign

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(macsState *basetypes.SetValue, macsToAssign *models.MacAddresses, macsToUnassign *models.MacAddresses, macsAssigned []string, macsUnassigned []string) (*basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var listValue []attr.Value
	var listValueType attr.Type = basetypes.StringType{}

	// remove unassigned devices from state
	for _, mac := range macsState.Elements() {
		var sInterface interface{} = mac
		s := sInterface.(basetypes.StringValue)
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
			listValue = append(listValue, mac)
		}
	}

	// add assigned devices to state
	for _, mac := range macsToAssign.Macs {
		if macInSlice(mac, macsAssigned) {
			listValue = append(listValue, types.StringValue(mac))
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

	newMacsState, e := types.SetValue(listValueType, listValue)
	diags.Append(e...)

	return &newMacsState, diags
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
		var sInterface interface{} = b
		s := sInterface.(basetypes.StringValue)
		if s.ValueString() == mac {
			return true
		}
	}
	return false
}
