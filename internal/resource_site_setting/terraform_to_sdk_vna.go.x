package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) *models.SiteSettingVna {
	tflog.Debug(ctx, "vnaTerraformToSdk")
	data := models.SiteSettingVna{}

	data.Enabled = d.Enabled.ValueBool()

	return &data
}
