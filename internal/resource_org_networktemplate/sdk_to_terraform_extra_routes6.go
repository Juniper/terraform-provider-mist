package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func extraRoute6NextQualifiedSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRoute6NextQualifiedProperties) basetypes.MapValue {

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
func extraRoutes6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ExtraRoute6) basetypes.MapValue {

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
		if len(d.NextQualified) > 0 {
			nextQualified = extraRoute6NextQualifiedSdkToTerraform(ctx, diags, d.NextQualified)
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
