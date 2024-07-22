package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func rtsaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RtsaValue) *models.SiteSettingRtsa {
	data := models.SiteSettingRtsa{}

	if d.AppWaking.ValueBoolPointer() != nil {
		data.AppWaking = d.AppWaking.ValueBoolPointer()
	}
	if d.DisableDeadReckoning.ValueBoolPointer() != nil {
		data.DisableDeadReckoning = d.DisableDeadReckoning.ValueBoolPointer()
	}
	if d.DisablePressureSensor.ValueBoolPointer() != nil {
		data.DisablePressureSensor = d.DisablePressureSensor.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.TrackAsset.ValueBoolPointer() != nil {
		data.TrackAsset = d.TrackAsset.ValueBoolPointer()
	}

	return &data
}
