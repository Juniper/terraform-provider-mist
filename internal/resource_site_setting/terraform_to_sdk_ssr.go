package resource_site_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ssrTerraformToSdk(d SsrValue) *models.SiteSettingSsr {
	data := models.SiteSettingSsr{}

	data.ConductorHosts = mistutils.ListOfStringTerraformToSdk(d.ConductorHosts)
	data.DisableStats = d.DisableStats.ValueBoolPointer()

	return &data
}
