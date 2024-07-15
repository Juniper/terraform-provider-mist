package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func l2tpStatsSessionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ApStatsL2TpStatSession) basetypes.ListValue {
	var data_list = []SessionsValue{}
	for _, d := range l {
		var local_sid basetypes.Int64Value
		var remote_id basetypes.StringValue
		var remote_sid basetypes.Int64Value
		var state basetypes.StringValue

		if d.LocalSid.Value() != nil {
			local_sid = types.Int64Value(int64(*d.LocalSid.Value()))
		}
		if d.RemoteId.Value() != nil {
			remote_id = types.StringValue(*d.RemoteId.Value())
		}
		if d.RemoteSid.Value() != nil {
			remote_sid = types.Int64Value(int64(*d.RemoteSid.Value()))
		}
		if d.State != nil {
			state = types.StringValue(string(*d.State))
		}

		data_map_attr_type := SessionsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"local_sid":  local_sid,
			"remote_id":  remote_id,
			"remote_sid": remote_sid,
			"state":      state,
		}
		data, e := NewSessionsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, SessionsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func l2tpStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ApStatsL2TpStat) basetypes.MapValue {
	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var sessions basetypes.ListValue = types.ListUnknown(SessionsValue{}.Type(ctx))
		var state basetypes.StringValue
		var uptime basetypes.Int64Value
		var wxtunnel_id basetypes.StringValue

		if d.Sessions != nil {
			sessions = l2tpStatsSessionSdkToTerraform(ctx, diags, d.Sessions)
		}
		if d.State != nil {
			state = types.StringValue(string(*d.State))
		}
		if d.Uptime.Value() != nil {
			uptime = types.Int64Value(int64(*d.Uptime.Value()))
		}
		if d.WxtunnelId.Value() != nil {
			wxtunnel_id = types.StringValue(d.WxtunnelId.Value().String())
		}

		data_map_attr_type := L2tpStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"sessions":    sessions,
			"state":       state,
			"uptime":      uptime,
			"wxtunnel_id": wxtunnel_id,
		}
		data, e := NewL2tpStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, L2tpStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
