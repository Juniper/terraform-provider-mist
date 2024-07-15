package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoute6NextQualifiedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRoute6PropertiesNextQualifiedProperties) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var metric basetypes.Int64Value
		var preference basetypes.Int64Value

		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.Preference.Value() != nil {
			preference = types.Int64Value(int64(*d.Preference.Value()))
		}

		data_map_attr_type := NextQualifiedValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"metric":     metric,
			"preference": preference,
		}
		data, e := NewNextQualifiedValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := NextQualifiedValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func extraRoutes6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRoute6Properties) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var discard basetypes.BoolValue
		var metric basetypes.Int64Value
		var next_qualified basetypes.MapValue = types.MapNull(NextQualifiedValue{}.Type(ctx))
		var no_resolve basetypes.BoolValue
		var preference basetypes.Int64Value
		var via basetypes.StringValue

		if d.Discard != nil {
			discard = types.BoolValue(*d.Discard)
		}
		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.NextQualified != nil && len(d.NextQualified) > 0 {
			next_qualified = extraRoute6NextQualifiedSdkToTerraform(ctx, diags, d.NextQualified)
		}
		if d.NoResolve != nil {
			no_resolve = types.BoolValue(*d.NoResolve)
		}
		if d.Preference.Value() != nil {
			preference = types.Int64Value(int64(*d.Preference.Value()))
		}
		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		data_map_attr_type := ExtraRoutesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"discard":        discard,
			"metric":         metric,
			"next_qualified": next_qualified,
			"no_resolve":     no_resolve,
			"preference":     preference,
			"via":            via,
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
