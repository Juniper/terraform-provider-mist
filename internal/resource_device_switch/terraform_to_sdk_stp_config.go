package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func stpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d StpConfigValue) *models.SwitchStpConfig {
	tflog.Debug(ctx, "stpConfigTerraformToSdk")

	data := models.SwitchStpConfig{}

	if d.StpConfigType.ValueStringPointer() != nil {
		data.Type = models.ToPointer(models.SwitchStpConfigTypeEnum(d.StpConfigType.ValueString()))
	}

	return &data
}
