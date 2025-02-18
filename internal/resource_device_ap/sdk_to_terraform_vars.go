package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]string) basetypes.MapValue {
	dataMap := make(map[string]string)
	for k, v := range d {
		dataMap[k] = v
	}
	stateResult, e := types.MapValueFrom(ctx, types.StringType, dataMap)
	diags.Append(e...)
	return stateResult
}
