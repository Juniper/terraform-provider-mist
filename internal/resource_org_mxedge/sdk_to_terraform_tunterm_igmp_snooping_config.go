package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func tuntermIgmpSnoopingConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermIgmpSnoopingConfig) TuntermIgmpSnoopingConfigValue {

	var enabled types.Bool
	var querier = types.ObjectNull(QuerierValue{}.AttributeTypes(ctx))
	var vlanIds = types.ListNull(types.Int64Type)

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Querier != nil {
		querierValue := querierSdkToTerraform(ctx, diags, d.Querier)
		querierObj, e := querierValue.ToObjectValue(ctx)
		diags.Append(e...)
		querier = querierObj
	}
	if d.VlanIds != nil {
		vlanIds_list := make([]attr.Value, len(d.VlanIds))
		for i, v := range d.VlanIds {
			vlanIds_list[i] = types.Int64Value(int64(v))
		}
		vlanIds_result, e := types.ListValue(types.Int64Type, vlanIds_list)
		diags.Append(e...)
		vlanIds = vlanIds_result
	}

	data_map_attr_type := TuntermIgmpSnoopingConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":  enabled,
		"querier":  querier,
		"vlan_ids": vlanIds,
	}
	data, e := NewTuntermIgmpSnoopingConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func querierSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermIgmpSnoopingQuerier) QuerierValue {

	var maxResponseTime types.Int64
	var mtu types.Int64
	var queryInterval types.Int64
	var robustness types.Int64
	var version types.Int64

	if d.MaxResponseTime != nil {
		maxResponseTime = types.Int64Value(int64(*d.MaxResponseTime))
	}
	if d.Mtu != nil {
		mtu = types.Int64Value(int64(*d.Mtu))
	}
	if d.QueryInterval != nil {
		queryInterval = types.Int64Value(int64(*d.QueryInterval))
	}
	if d.Robustness != nil {
		robustness = types.Int64Value(int64(*d.Robustness))
	}
	if d.Version != nil {
		version = types.Int64Value(int64(*d.Version))
	}

	data_map_attr_type := QuerierValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"max_response_time": maxResponseTime,
		"mtu":               mtu,
		"query_interval":    queryInterval,
		"robustness":        robustness,
		"version":           version,
	}
	data, e := NewQuerierValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
