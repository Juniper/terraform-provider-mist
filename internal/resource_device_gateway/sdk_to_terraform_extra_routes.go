package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayExtraRoute) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		data_map_attr_type := ExtraRoutesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"via": via,
		}
		data, e := NewExtraRoutesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := ExtraRoutesValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
