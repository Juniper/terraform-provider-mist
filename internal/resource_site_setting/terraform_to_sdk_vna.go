package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VnaValue) *models.SiteSettingVna {
	data := models.SiteSettingVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
