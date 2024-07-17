package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func analyticTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AnalyticValue) *models.SiteSettingAnalytic {
	data := models.NewSiteSettingAnalytic()

	data.SetEnabled(d.Enabled.ValueBool())

	return data
}
