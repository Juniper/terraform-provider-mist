package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func occupancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteOccupancyAnalytics) OccupancyValue {

	var assets_enabled basetypes.BoolValue
	var clients_enabled basetypes.BoolValue
	var min_duration basetypes.Int64Value
	var sdkclients_enabled basetypes.BoolValue
	var unconnected_clients_enabled basetypes.BoolValue

	if d != nil && d.AssetsEnabled != nil {
		assets_enabled = types.BoolValue(*d.AssetsEnabled)
	}
	if d != nil && d.ClientsEnabled != nil {
		clients_enabled = types.BoolValue(*d.ClientsEnabled)
	}
	if d != nil && d.MinDuration != nil {
		min_duration = types.Int64Value(int64(*d.MinDuration))
	}
	if d != nil && d.SdkclientsEnabled != nil {
		sdkclients_enabled = types.BoolValue(*d.SdkclientsEnabled)
	}
	if d != nil && d.UnconnectedClientsEnabled != nil {
		unconnected_clients_enabled = types.BoolValue(*d.UnconnectedClientsEnabled)
	}

	data_map_attr_type := OccupancyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"assets_enabled":              assets_enabled,
		"clients_enabled":             clients_enabled,
		"min_duration":                min_duration,
		"sdkclients_enabled":          sdkclients_enabled,
		"unconnected_clients_enabled": unconnected_clients_enabled,
	}
	data, e := NewOccupancyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
