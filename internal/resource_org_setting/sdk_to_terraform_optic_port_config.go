package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func opticPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OpticPortConfigPort) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {

		var channelized basetypes.BoolValue
		var speed basetypes.StringValue

		if d.Channelized != nil {
			channelized = types.BoolValue(*d.Channelized)
		}
		if d.Speed != nil {
			speed = types.StringValue(*d.Speed)
		}

		data_map_attr_type := OpticPortConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"channelized": channelized,
			"speed":       speed,
		}
		data, e := NewOpticPortConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := OpticPortConfigValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
