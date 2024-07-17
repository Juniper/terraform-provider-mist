package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func zoneOccupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ZoneOccupancyAlertValue) *models.SiteZoneOccupancyAlert {
	data := models.SiteZoneOccupancyAlert{}

	data.EmailNotifiers = mist_transform.ListOfStringTerraformToSdk(ctx, d.EmailNotifiers)
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.Threshold = models.ToPointer(int(d.Threshold.ValueInt64()))

	return &data
}
