package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRouteSdkNextQualifiedToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRouteNextQualifiedProperties) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var metric basetypes.Int64Value
		var preference basetypes.Int64Value

		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.Preference.Value() != nil {
			preference = types.Int64Value(int64(*d.Preference.Value()))
		}

		dataMapValue := map[string]attr.Value{
			"metric":     metric,
			"preference": preference,
		}
		data, e := NewNextQualifiedValue(NextQualifiedValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := NextQualifiedValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
func extraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRoute) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var discard basetypes.BoolValue
		var metric basetypes.Int64Value
		var nextQualified = types.MapNull(NextQualifiedValue{}.Type(ctx))
		var noResolve basetypes.BoolValue
		var preference basetypes.Int64Value
		var via basetypes.StringValue

		if d.Discard != nil {
			discard = types.BoolValue(*d.Discard)
		}
		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.NextQualified != nil && len(d.NextQualified) > 0 {
			nextQualified = extraRouteSdkNextQualifiedToTerraform(ctx, diags, d.NextQualified)
		}
		if d.NoResolve != nil {
			noResolve = types.BoolValue(*d.NoResolve)
		}
		if d.Preference.Value() != nil {
			preference = types.Int64Value(int64(*d.Preference.Value()))
		}
		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		dataMapValue := map[string]attr.Value{
			"discard":        discard,
			"metric":         metric,
			"next_qualified": nextQualified,
			"no_resolve":     noResolve,
			"preference":     preference,
			"via":            via,
		}
		data, e := NewExtraRoutesValue(ExtraRoutesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := ExtraRoutesValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
