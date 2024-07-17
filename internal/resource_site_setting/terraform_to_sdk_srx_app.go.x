package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func srxAppTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SrxAppValue) *models.SiteSettingSrxApp {
	data := models.SiteSettingSrxApp{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	return &data
}
