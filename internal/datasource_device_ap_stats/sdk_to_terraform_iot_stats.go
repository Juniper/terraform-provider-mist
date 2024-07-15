package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func iotStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ApStatsIotStatAdditionalProperties) basetypes.MapValue {
	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var value basetypes.Int64Value

		if d.Value.Value() != nil {
			value = types.Int64Value(int64(*d.Value.Value()))
		}

		data_map_attr_type := IotStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"value": value,
		}
		data, e := NewIotStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, IotStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
