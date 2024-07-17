package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, m basetypes.MapValue) map[string]string {
	data_map := make(map[string]string)
	for k, v := range m.Elements() {
		data_map[k] = string(v.String())
	}
	return data_map
}
