package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func skyAtpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SkyatpValue) *models.SiteSettingSkyatp {
	tflog.Debug(ctx, "skyAtpTerraformToSdk")
	data := models.SiteSettingSkyatp{}

	data.Enabled = d.Enabled.ValueBoolPointer()
	data.SendIpMacMapping = d.SendIpMacMapping.ValueBoolPointer()

	return &data
}
