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

	var assetsEnabled basetypes.BoolValue
	var clientsEnabled basetypes.BoolValue
	var minDuration basetypes.Int64Value
	var sdkclientsEnabled basetypes.BoolValue
	var unconnectedClientsEnabled basetypes.BoolValue

	if d != nil && d.AssetsEnabled != nil {
		assetsEnabled = types.BoolValue(*d.AssetsEnabled)
	}
	if d != nil && d.ClientsEnabled != nil {
		clientsEnabled = types.BoolValue(*d.ClientsEnabled)
	}
	if d != nil && d.MinDuration != nil {
		minDuration = types.Int64Value(int64(*d.MinDuration))
	}
	if d != nil && d.SdkclientsEnabled != nil {
		sdkclientsEnabled = types.BoolValue(*d.SdkclientsEnabled)
	}
	if d != nil && d.UnconnectedClientsEnabled != nil {
		unconnectedClientsEnabled = types.BoolValue(*d.UnconnectedClientsEnabled)
	}

	dataMapValue := map[string]attr.Value{
		"assets_enabled":              assetsEnabled,
		"clients_enabled":             clientsEnabled,
		"min_duration":                minDuration,
		"sdkclients_enabled":          sdkclientsEnabled,
		"unconnected_clients_enabled": unconnectedClientsEnabled,
	}
	data, e := NewOccupancyValue(OccupancyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
