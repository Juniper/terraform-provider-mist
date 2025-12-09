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

func gpsStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApGpsStat) basetypes.ObjectValue {
	var accuracy basetypes.NumberValue
	var altitude basetypes.NumberValue
	var latitude basetypes.NumberValue
	var longitude basetypes.NumberValue
	var src basetypes.StringValue
	var timestamp basetypes.Float64Value

	if d.Accuracy != nil {
		accuracy = types.NumberValue(big.NewFloat(*d.Accuracy))
	}
	if d.Altitude != nil {
		altitude = types.NumberValue(big.NewFloat(*d.Altitude))
	}
	if d.Latitude != nil {
		latitude = types.NumberValue(big.NewFloat(*d.Latitude))
	}
	if d.Longitude != nil {
		longitude = types.NumberValue(big.NewFloat(*d.Longitude))
	}
	if d.Src != nil {
		src = types.StringValue(string(*d.Src))
	}
	if d.Timestamp != nil {
		timestamp = types.Float64Value(*d.Timestamp)
	}

	dataMapValue := map[string]attr.Value{
		"accuracy":  accuracy,
		"altitude":  altitude,
		"latitude":  latitude,
		"longitude": longitude,
		"src":       src,
		"timestamp": timestamp,
	}
	data, e := types.ObjectValue(GpsStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
