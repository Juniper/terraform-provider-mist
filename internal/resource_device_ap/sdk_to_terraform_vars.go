package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]string) basetypes.MapValue {
	data_map := make(map[string]string)
	for k, v := range d {
		data_map[k] = v
	}
	state_result, e := types.MapValueFrom(ctx, types.StringType, data_map)
	diags.Append(e...)
	return state_result
}
