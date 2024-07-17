package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func ledTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d LedValue) *models.ApLed {
	data := models.ApLed{}

	if d.Brightness.ValueInt64Pointer() != nil {
		data.Brightness = models.ToPointer(int(d.Brightness.ValueInt64()))
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
