package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wanVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WanVnaValue) *models.SiteSettingWanVna {
	data := models.SiteSettingWanVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}