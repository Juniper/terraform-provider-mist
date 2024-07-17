package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func occupancyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d OccupancyValue) *models.SiteOccupancyAnalytics {
	data := models.SiteOccupancyAnalytics{}

	if d.AssetsEnabled.ValueBoolPointer() != nil {
		data.AssetsEnabled = d.AssetsEnabled.ValueBoolPointer()
	}
	if d.ClientsEnabled.ValueBoolPointer() != nil {
		data.ClientsEnabled = d.ClientsEnabled.ValueBoolPointer()
	}
	if d.MinDuration.ValueInt64Pointer() != nil {
		data.MinDuration = models.ToPointer(int(d.MinDuration.ValueInt64()))
	}
	if d.SdkclientsEnabled.ValueBoolPointer() != nil {
		data.SdkclientsEnabled = d.SdkclientsEnabled.ValueBoolPointer()
	}
	if d.UnconnectedClientsEnabled.ValueBoolPointer() != nil {
		data.UnconnectedClientsEnabled = d.UnconnectedClientsEnabled.ValueBoolPointer()
	}

	return &data
}
