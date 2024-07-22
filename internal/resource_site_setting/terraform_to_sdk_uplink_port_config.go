package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func uplinkPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d UplinkPortConfigValue) *models.ApUplinkPortConfig {
	data := models.ApUplinkPortConfig{}

	if d.Dot1x.ValueBoolPointer() != nil {
		data.Dot1x = d.Dot1x.ValueBoolPointer()
	}

	if d.KeepWlansUpIfDown.ValueBoolPointer() != nil {
		data.KeepWlansUpIfDown = d.Dot1x.ValueBoolPointer()
	}

	return &data
}
