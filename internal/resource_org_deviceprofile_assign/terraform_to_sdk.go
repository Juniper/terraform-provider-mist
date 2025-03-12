package resource_org_deviceprofile_assign

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(macsPlan basetypes.SetValue, macsState basetypes.SetValue) (models.MacAddresses, models.MacAddresses, diag.Diagnostics) {
	var diags diag.Diagnostics
	var macsToAssign models.MacAddresses
	var macsToUnassign models.MacAddresses

	macsToAssign.Macs = diffList(macsPlan, macsState)
	macsToUnassign.Macs = diffList(macsState, macsPlan)

	return macsToAssign, macsToUnassign, diags
}

func diffList(listOne basetypes.SetValue, listTwo basetypes.SetValue) []string {
	var diff []string
	for _, elo := range listOne.Elements() {
		var soInterface interface{} = elo
		so := soInterface.(basetypes.StringValue)
		inBoth := false
		for _, elt := range listTwo.Elements() {
			var stInterface interface{} = elt
			st := stInterface.(basetypes.StringValue)
			if so.Equal(st) {
				inBoth = true
			}
		}
		if !inBoth {
			diff = append(diff, so.ValueString())
		}
	}
	return diff
}
