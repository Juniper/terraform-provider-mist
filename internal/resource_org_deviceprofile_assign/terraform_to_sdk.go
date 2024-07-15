package resource_org_deviceprofile_assign

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, macs_plan basetypes.ListValue, macs_state basetypes.ListValue) (models.MacAddresses, models.MacAddresses, diag.Diagnostics) {
	var diags diag.Diagnostics
	var macs_to_assign models.MacAddresses
	var macs_to_unssign models.MacAddresses

	macs_to_assign.Macs = diffList(ctx, macs_plan, macs_state)
	macs_to_unssign.Macs = diffList(ctx, macs_state, macs_plan)

	return macs_to_assign, macs_to_unssign, diags
}

func diffList(ctx context.Context, list_one basetypes.ListValue, list_two basetypes.ListValue) []string {
	var diff []string
	for _, elo := range list_one.Elements() {
		in_both := false
		for _, elt := range list_two.Elements() {
			if elo.String() == elt.String() {
				in_both = true
				break
			}
		}
		if !in_both {
			var s_interface interface{} = elo
			s := s_interface.(basetypes.StringValue)
			diff = append(diff, s.ValueString())
		}
	}
	return diff
}
