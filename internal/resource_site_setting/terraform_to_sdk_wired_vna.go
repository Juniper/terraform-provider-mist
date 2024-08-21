package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wiredVnaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d WiredVnaValue) *models.SiteSettingWiredVna {
	data := models.SiteSettingWiredVna{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
