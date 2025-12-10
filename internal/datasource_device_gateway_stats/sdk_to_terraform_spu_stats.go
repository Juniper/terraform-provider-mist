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

	var dataList []SpuStatValue
	for _, d := range l {
		var spuCpu basetypes.Int64Value
		var spuCurrentSession basetypes.Int64Value
		var spuMaxSession basetypes.Int64Value
		var spuMemory basetypes.Int64Value
		var spuPendingSession basetypes.Int64Value
		var spuUptime basetypes.Int64Value
		var spuValidSession basetypes.Int64Value

		if d.SpuCpu != nil {
			spuCpu = types.Int64Value(int64(*d.SpuCpu))
		}
		if d.SpuCurrentSession != nil {
			spuCurrentSession = types.Int64Value(int64(*d.SpuCurrentSession))
		}
		if d.SpuMaxSession != nil {
			spuMaxSession = types.Int64Value(int64(*d.SpuMaxSession))
		}
		if d.SpuMemory != nil {
			spuMemory = types.Int64Value(int64(*d.SpuMemory))
		}
		if d.SpuPendingSession != nil {
			spuPendingSession = types.Int64Value(int64(*d.SpuPendingSession))
		}
		if d.SpuUptime != nil {
			spuUptime = types.Int64Value(int64(*d.SpuUptime))
		}
		if d.SpuValidSession != nil {
			spuValidSession = types.Int64Value(int64(*d.SpuValidSession))
		}

		dataMapValue := map[string]attr.Value{
			"spu_cpu":             spuCpu,
			"spu_current_session": spuCurrentSession,
			"spu_max_session":     spuMaxSession,
			"spu_memory":          spuMemory,
			"spu_pending_session": spuPendingSession,
			"spu_uptime":          spuUptime,
			"spu_valid_session":   spuValidSession,
		}
		data, e := NewSpuStatValue(SpuStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SpuStatValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
