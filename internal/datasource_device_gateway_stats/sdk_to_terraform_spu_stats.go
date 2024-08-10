package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func spuStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsGatewaySpuItem) basetypes.ListValue {

	var data_list = []SpuStatValue{}
	for _, d := range l {
		var spu_cpu basetypes.Int64Value
		var spu_current_session basetypes.Int64Value
		var spu_max_session basetypes.Int64Value
		var spu_memory basetypes.Int64Value
		var spu_pending_session basetypes.Int64Value
		var spu_valid_session basetypes.Int64Value

		if d.SpuCpu != nil {
			spu_cpu = types.Int64Value(int64(*d.SpuCpu))
		}
		if d.SpuCurrentSession != nil {
			spu_current_session = types.Int64Value(int64(*d.SpuCurrentSession))
		}
		if d.SpuMaxSession != nil {
			spu_max_session = types.Int64Value(int64(*d.SpuMaxSession))
		}
		if d.SpuMemory != nil {
			spu_memory = types.Int64Value(int64(*d.SpuMemory))
		}
		if d.SpuPendingSession != nil {
			spu_pending_session = types.Int64Value(int64(*d.SpuPendingSession))
		}
		if d.SpuValidSession != nil {
			spu_valid_session = types.Int64Value(int64(*d.SpuValidSession))
		}

		data_map_attr_type := SpuStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"spu_cpu":             spu_cpu,
			"spu_current_session": spu_current_session,
			"spu_max_session":     spu_max_session,
			"spu_memory":          spu_memory,
			"spu_pending_session": spu_pending_session,
			"spu_valid_session":   spu_valid_session,
		}
		data, e := NewSpuStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, SpuStatValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
