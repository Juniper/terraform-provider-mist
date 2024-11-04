package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func opticPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.OpticPortConfigPort {
	data_map := make(map[string]models.OpticPortConfigPort)
	for k, data_attr := range d.Elements() {
		var data_interface interface{} = data_attr
		plan := data_interface.(OpticPortConfigValue)

		data := models.OpticPortConfigPort{}
		if plan.Channelized.ValueBoolPointer() != nil {
			data.Channelized = plan.Channelized.ValueBoolPointer()
		}
		if plan.Speed.ValueStringPointer() != nil {
			data.Speed = models.ToPointer(plan.Speed.ValueString())
		}
		data_map[k] = data
	}
	return data_map
}
