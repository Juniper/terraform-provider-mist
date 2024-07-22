package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func rtsaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteSettingRtsa) RtsaValue {
	var app_waking basetypes.BoolValue
	var disable_dead_reckoning basetypes.BoolValue
	var disable_pressure_sensor basetypes.BoolValue
	var enabled basetypes.BoolValue
	var track_asset basetypes.BoolValue

	if d.AppWaking != nil {
		app_waking = types.BoolValue(*d.AppWaking)
	}
	if d.DisableDeadReckoning != nil {
		disable_dead_reckoning = types.BoolValue(*d.DisableDeadReckoning)
	}
	if d.DisablePressureSensor != nil {
		disable_pressure_sensor = types.BoolValue(*d.DisablePressureSensor)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TrackAsset != nil {
		track_asset = types.BoolValue(*d.TrackAsset)
	}

	data_map_attr_type := RtsaValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"app_waking":              app_waking,
		"disable_dead_reckoning":  disable_dead_reckoning,
		"disable_pressure_sensor": disable_pressure_sensor,
		"enabled":                 enabled,
		"track_asset":             track_asset,
	}
	data, e := NewRtsaValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
