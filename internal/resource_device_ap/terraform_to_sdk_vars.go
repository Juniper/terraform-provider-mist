package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, m basetypes.MapValue) map[string]string {
	data_map := make(map[string]string)
	for k, v := range m.Elements() {
		var vi interface{} = v
		vd := vi.(basetypes.StringValue)
		data_map[k] = string(vd.ValueString())
	}
	return data_map
}
