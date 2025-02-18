package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func iotStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.StatsApIotStatAdditionalProperties) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var value basetypes.Int64Value

		if d.Value.Value() != nil {
			value = types.Int64Value(int64(*d.Value.Value()))
		}

		dataMapValue := map[string]attr.Value{
			"value": value,
		}
		data, e := NewIotStatValue(IotStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, IotStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
