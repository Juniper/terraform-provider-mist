package resource_site_setting

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func zoneOccupancyTerraformToSdk(d ZoneOccupancyAlertValue) *models.SiteZoneOccupancyAlert {
	data := models.SiteZoneOccupancyAlert{}

	data.EmailNotifiers = misttransform.ListOfStringTerraformToSdk(d.EmailNotifiers)
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.Threshold = models.ToPointer(int(d.Threshold.ValueInt64()))

	return &data
}
