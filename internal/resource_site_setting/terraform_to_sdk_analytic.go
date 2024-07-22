package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func analyticTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AnalyticValue) *models.SiteSettingAnalytic {
	data := models.SiteSettingAnalytic{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}

	return &data
}
