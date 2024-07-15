package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vlanIdsSkToTerraform(ctx context.Context, diags *diag.Diagnostics, data []int) basetypes.ListValue {

	var list []attr.Value
	for _, v := range data {
		list = append(list, types.Int64Value(int64(v)))
	}
	r, e := types.ListValue(basetypes.Int64Type{}, list)
	diags.Append(e...)

	return r
}
