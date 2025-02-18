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

	var fibRoutes basetypes.Int64Value
	var maxUnicastRoutesSupported basetypes.Int64Value
	var ribRoutes basetypes.Int64Value
	var totalRoutes basetypes.Int64Value

	if d.FibRoutes != nil {
		fibRoutes = types.Int64Value(int64(*d.FibRoutes))
	}
	if d.MaxUnicastRoutesSupported != nil {
		maxUnicastRoutesSupported = types.Int64Value(int64(*d.MaxUnicastRoutesSupported))
	}
	if d.RibRoutes != nil {
		ribRoutes = types.Int64Value(int64(*d.RibRoutes))
	}
	if d.TotalRoutes != nil {
		totalRoutes = types.Int64Value(int64(*d.TotalRoutes))
	}

	dataMapValue := map[string]attr.Value{
		"fib_routes":                   fibRoutes,
		"max_unicast_routes_supported": maxUnicastRoutesSupported,
		"rib_routes":                   ribRoutes,
		"total_routes":                 totalRoutes,
	}
	data, e := basetypes.NewObjectValue(RouteSummaryStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
