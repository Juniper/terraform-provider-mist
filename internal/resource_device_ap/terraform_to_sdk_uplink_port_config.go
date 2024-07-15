package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func uplinkPortConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d UplinkPortConfigValue) *models.ApUplinkPortConfig {
	tflog.Debug(ctx, "uplinkPortConfigTerraformToSdk")
	data := models.ApUplinkPortConfig{}

	if d.Dot1x.ValueBoolPointer() != nil {
		data.Dot1x = d.Dot1x.ValueBoolPointer()
	}
	if d.KeepWlansUpIfDown.ValueBoolPointer() != nil {
		data.KeepWlansUpIfDown = d.KeepWlansUpIfDown.ValueBoolPointer()
	}

	return &data
}
