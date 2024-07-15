package resource_org_deviceprofile_assign

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, macs_state *basetypes.ListValue, macs_assigned []string, macs_unassigned []string) (*basetypes.ListValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var list_value []attr.Value
	var list_value_type attr.Type = basetypes.StringType{}
	for _, mac := range macs_state.Elements() {
		var s_interface interface{} = mac
		s := s_interface.(basetypes.StringValue)
		if !stringInSlice(s.ValueString(), macs_unassigned) {
			list_value = append(list_value, mac)
		}
	}
	for _, mac := range macs_assigned {
		list_value = append(list_value, types.StringValue(mac))
	}

	new_macs_state, e := types.ListValue(list_value_type, list_value)
	diags.Append(e...)

	return &new_macs_state, diags
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
