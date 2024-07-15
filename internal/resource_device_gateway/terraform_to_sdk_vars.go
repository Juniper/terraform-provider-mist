package resource_device_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func varsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, m basetypes.MapValue) map[string]string {
	tflog.Debug(ctx, "varsTerraformToSdk")
	data_map := make(map[string]string)
	for k, v := range m.Elements() {
		data_map[k] = string(v.String())
	}
	return data_map
}
