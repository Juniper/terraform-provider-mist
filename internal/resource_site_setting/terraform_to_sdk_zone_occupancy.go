package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func zoneOccupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ZoneOccupancyAlertValue) *models.SiteZoneOccupancyAlert {
	tflog.Debug(ctx, "zoneOccupancyTerraformToSdk")
	data := models.SiteZoneOccupancyAlert{}

	data.EmailNotifiers = mist_transform.ListOfStringTerraformToSdk(ctx, d.EmailNotifiers)
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.Threshold = models.ToPointer(int(d.Threshold.ValueInt64()))

	return &data
}
