package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func centrakTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CentrakValue) *models.ApCentrak {
	data := models.ApCentrak{}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
