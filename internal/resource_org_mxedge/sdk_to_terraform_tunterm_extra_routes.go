package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermExtraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.MxedgeTuntermExtraRoute) basetypes.MapValue {

	state_value_map_type := TuntermExtraRoutesValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d {
		var via types.String

		if v.Via != nil {
			via = types.StringValue(*v.Via)
		}

		data_map_attr_type := TuntermExtraRoutesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"via": via,
		}
		data, e := NewTuntermExtraRoutesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}

	state_result, e := types.MapValueFrom(ctx, state_value_map_type, state_value_map)
	diags.Append(e...)
	return state_result
}
