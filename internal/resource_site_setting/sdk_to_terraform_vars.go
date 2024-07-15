package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func varsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]string) basetypes.MapValue {
	tflog.Debug(ctx, "varsSdkToTerraform")
	data_map := make(map[string]string)
	for k, v := range d {
		data_map[k] = v
	}
	state_result, e := types.MapValueFrom(ctx, types.StringType, data_map)
	diags.Append(e...)
	return state_result
}
