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
	var radiusM basetypes.NumberValue
	var x basetypes.Float64Value

	if d.Radius != nil {
		radius = types.NumberValue(big.NewFloat(*d.Radius))
	}
	if d.RadiusM != nil {
		radiusM = types.NumberValue(big.NewFloat(*d.RadiusM))
	}
	if d.X != nil {
		x = types.Float64Value(*d.X)
	}

	dataMapValue := map[string]attr.Value{
		"radius":   radius,
		"radius_m": radiusM,
		"x":        x,
	}
	data, e := types.ObjectValue(ProbabilitySurfaceValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func autoPlacementInfoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoPlacementInfo) basetypes.ObjectValue {
	var clusterNumber basetypes.Int64Value
	var orientationStats basetypes.Int64Value
	var probabilitySurface = types.ObjectNull(ProbabilitySurfaceValue{}.AttributeTypes(ctx))

	if d.ClusterNumber != nil {
		clusterNumber = types.Int64Value(int64(*d.ClusterNumber))
	}
	if d.OrientationStats != nil {
		orientationStats = types.Int64Value(int64(*d.OrientationStats))
	}
	if d.ProbabilitySurface != nil {
		probabilitySurface = autoPlacementInfoProbaSdkToTerraform(ctx, diags, d.ProbabilitySurface)
	}

	dataMapValue := map[string]attr.Value{
		"cluster_number":      clusterNumber,
		"orientation_stats":   orientationStats,
		"probability_surface": probabilitySurface,
	}
	data, e := types.ObjectValue(InfoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func autoPlacementSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoPlacement) basetypes.ObjectValue {
	var info = types.ObjectNull(InfoValue{}.AttributeTypes(ctx))
	var recommendedAnchor basetypes.BoolValue
	var status basetypes.StringValue
	var statusDetail basetypes.StringValue
	var useAutoPlacement basetypes.BoolValue
	var x basetypes.Float64Value
	var xM basetypes.Float64Value
	var y basetypes.Float64Value
	var yM basetypes.Float64Value

	if d.Info != nil {
		info = autoPlacementInfoSdkToTerraform(ctx, diags, d.Info)
	}
	if d.RecommendedAnchor != nil {
		recommendedAnchor = types.BoolValue(*d.RecommendedAnchor)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.StatusDetail != nil {
		statusDetail = types.StringValue(*d.StatusDetail)
	}
	if d.UseAutoPlacement != nil {
		useAutoPlacement = types.BoolValue(*d.UseAutoPlacement)
	}
	if d.X != nil {
		x = types.Float64Value(*d.X)
	}
	if d.XM != nil {
		xM = types.Float64Value(*d.XM)
	}
	if d.Y != nil {
		y = types.Float64Value(*d.Y)
	}
	if d.YM != nil {
		yM = types.Float64Value(*d.YM)
	}

	dataMapValue := map[string]attr.Value{
		"info":               info,
		"recommended_anchor": recommendedAnchor,
		"status":             status,
		"status_detail":      statusDetail,
		"use_auto_placement": useAutoPlacement,
		"x":                  x,
		"x_m":                xM,
		"y":                  y,
		"y_m":                yM,
	}
	data, e := types.ObjectValue(AutoPlacementValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
