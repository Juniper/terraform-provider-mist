package resource_org_deviceprofile_assign

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, macs_plan basetypes.SetValue, macs_state basetypes.SetValue) (models.MacAddresses, models.MacAddresses, diag.Diagnostics) {
	var diags diag.Diagnostics
	var macs_to_assign models.MacAddresses
	var macs_to_unssign models.MacAddresses

	macs_to_assign.Macs = diffList(ctx, macs_plan, macs_state)
	macs_to_unssign.Macs = diffList(ctx, macs_state, macs_plan)

	return macs_to_assign, macs_to_unssign, diags
}

func diffList(ctx context.Context, list_one basetypes.SetValue, list_two basetypes.SetValue) []string {
	var diff []string
	for _, elo := range list_one.Elements() {
		var so_interface interface{} = elo
		so := so_interface.(basetypes.StringValue)
		in_both := false
		for _, elt := range list_two.Elements() {
			var st_interface interface{} = elt
			st := st_interface.(basetypes.StringValue)
			if so.Equal(st) {
				in_both = true
			}
		}
		if !in_both {
			diff = append(diff, so.ValueString())
		}
	}
	return diff
}
