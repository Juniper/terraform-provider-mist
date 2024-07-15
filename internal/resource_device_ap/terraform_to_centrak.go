package resource_device_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func centrakTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CentrakValue) *models.ApCentrak {
	tflog.Debug(ctx, "centrakTerraformToSdk")
	data := models.ApCentrak{}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
