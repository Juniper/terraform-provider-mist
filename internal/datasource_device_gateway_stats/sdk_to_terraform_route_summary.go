package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func routeSummaryStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RouteSummaryStats) basetypes.ObjectValue {

	var fib_routes basetypes.Int64Value
	var max_unicast_routes_supported basetypes.Int64Value
	var rib_routes basetypes.Int64Value
	var total_routes basetypes.Int64Value

	if d.FibRoutes != nil {
		fib_routes = types.Int64Value(int64(*d.FibRoutes))
	}
	if d.MaxUnicastRoutesSupported != nil {
		max_unicast_routes_supported = types.Int64Value(int64(*d.MaxUnicastRoutesSupported))
	}
	if d.RibRoutes != nil {
		rib_routes = types.Int64Value(int64(*d.RibRoutes))
	}
	if d.TotalRoutes != nil {
		total_routes = types.Int64Value(int64(*d.TotalRoutes))
	}

	data_map_attr_type := RouteSummaryStatsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"fib_routes":                   fib_routes,
		"max_unicast_routes_supported": max_unicast_routes_supported,
		"rib_routes":                   rib_routes,
		"total_routes":                 total_routes,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
