package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mistNacTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MistNacValue) *models.SwitchMistNac {
	data := models.SwitchMistNac{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if d.Network.ValueStringPointer() != nil {
		data.Network = models.ToPointer(d.Network.ValueString())
	}
	return &data
}
