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
	var appWaking basetypes.BoolValue
	var disableDeadReckoning basetypes.BoolValue
	var disablePressureSensor basetypes.BoolValue
	var enabled basetypes.BoolValue
	var trackAsset basetypes.BoolValue

	if d.AppWaking != nil {
		appWaking = types.BoolValue(*d.AppWaking)
	}
	if d.DisableDeadReckoning != nil {
		disableDeadReckoning = types.BoolValue(*d.DisableDeadReckoning)
	}
	if d.DisablePressureSensor != nil {
		disablePressureSensor = types.BoolValue(*d.DisablePressureSensor)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TrackAsset != nil {
		trackAsset = types.BoolValue(*d.TrackAsset)
	}

	dataMapValue := map[string]attr.Value{
		"app_waking":              appWaking,
		"disable_dead_reckoning":  disableDeadReckoning,
		"disable_pressure_sensor": disablePressureSensor,
		"enabled":                 enabled,
		"track_asset":             trackAsset,
	}
	data, e := NewRtsaValue(RtsaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
