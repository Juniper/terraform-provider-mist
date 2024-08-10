package datasource_device_ap_stats

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func autoPlacementInfoProbaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoPlacementInfoProbabilitySurface) basetypes.ObjectValue {
	var radius basetypes.NumberValue
	var radius_m basetypes.NumberValue
	var x basetypes.Float64Value

	if d.Radius != nil {
		radius = types.NumberValue(big.NewFloat(*d.Radius))
	}
	if d.RadiusM != nil {
		radius_m = types.NumberValue(big.NewFloat(*d.RadiusM))
	}
	if d.X != nil {
		x = types.Float64Value(*d.X)
	}

	data_map_attr_type := ProbabilitySurfaceValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"radius":   radius,
		"radius_m": radius_m,
		"x":        x,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func autoPlacementInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoPlacementInfo) basetypes.ObjectValue {
	var cluster_number basetypes.Int64Value
	var orientation_stats basetypes.Int64Value
	var probability_surface basetypes.ObjectValue = types.ObjectNull(ProbabilitySurfaceValue{}.AttributeTypes(ctx))

	if d.ClusterNumber != nil {
		cluster_number = types.Int64Value(int64(*d.ClusterNumber))
	}
	if d.OrientationStats != nil {
		orientation_stats = types.Int64Value(int64(*d.OrientationStats))
	}
	if d.ProbabilitySurface != nil {
		probability_surface = autoPlacementInfoProbaSdkToTerraform(ctx, diags, d.ProbabilitySurface)
	}

	data_map_attr_type := InfoValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"cluster_number":      cluster_number,
		"orientation_stats":   orientation_stats,
		"probability_surface": probability_surface,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func autoPlacementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoPlacement) basetypes.ObjectValue {
	var info basetypes.ObjectValue = types.ObjectNull(InfoValue{}.AttributeTypes(ctx))
	var recommended_anchor basetypes.BoolValue
	var status basetypes.StringValue
	var status_detail basetypes.StringValue
	var use_auto_placement basetypes.BoolValue
	var x basetypes.Float64Value
	var x_m basetypes.Float64Value
	var y basetypes.Float64Value
	var y_m basetypes.Float64Value

	if d.Info != nil {
		info = autoPlacementInfoSdkToTerraform(ctx, diags, d.Info)
	}
	if d.RecommendedAnchor != nil {
		recommended_anchor = types.BoolValue(*d.RecommendedAnchor)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.StatusDetail != nil {
		status_detail = types.StringValue(*d.StatusDetail)
	}
	if d.UseAutoPlacement != nil {
		use_auto_placement = types.BoolValue(*d.UseAutoPlacement)
	}
	if d.X != nil {
		x = types.Float64Value(*d.X)
	}
	if d.XM != nil {
		x_m = types.Float64Value(*d.XM)
	}
	if d.Y != nil {
		y = types.Float64Value(*d.Y)
	}
	if d.YM != nil {
		y_m = types.Float64Value(*d.YM)
	}

	data_map_attr_type := AutoPlacementValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"info":               info,
		"recommended_anchor": recommended_anchor,
		"status":             status,
		"status_detail":      status_detail,
		"use_auto_placement": use_auto_placement,
		"x":                  x,
		"x_m":                x_m,
		"y":                  y,
		"y_m":                y_m,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
