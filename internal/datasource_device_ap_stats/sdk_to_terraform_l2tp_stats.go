package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func l2tpStatsSessionSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsApL2tpStatSession) basetypes.ListValue {
	var dataList []SessionsValue
	for _, d := range l {
		var localSid basetypes.Int64Value
		var remoteId basetypes.StringValue
		var remoteSid basetypes.Int64Value
		var state basetypes.StringValue

		if d.LocalSid.Value() != nil {
			localSid = types.Int64Value(int64(*d.LocalSid.Value()))
		}
		if d.RemoteId.Value() != nil {
			remoteId = types.StringValue(*d.RemoteId.Value())
		}
		if d.RemoteSid.Value() != nil {
			remoteSid = types.Int64Value(int64(*d.RemoteSid.Value()))
		}
		if d.State != nil {
			state = types.StringValue(string(*d.State))
		}

		dataMapValue := map[string]attr.Value{
			"local_sid":  localSid,
			"remote_id":  remoteId,
			"remote_sid": remoteSid,
			"state":      state,
		}
		data, e := NewSessionsValue(SessionsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, SessionsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func l2tpStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.StatsApL2tpStat) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var sessions = types.ListUnknown(SessionsValue{}.Type(ctx))
		var state basetypes.StringValue
		var uptime basetypes.Int64Value
		var wxtunnelId basetypes.StringValue

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
			wxtunnelId = types.StringValue(d.WxtunnelId.Value().String())
		}

		dataMapValue := map[string]attr.Value{
			"sessions":    sessions,
			"state":       state,
			"uptime":      uptime,
			"wxtunnel_id": wxtunnelId,
		}
		data, e := NewL2tpStatValue(L2tpStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, L2tpStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
