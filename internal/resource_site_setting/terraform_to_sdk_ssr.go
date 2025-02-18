package resource_site_setting

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ssrTerraformToSdk(d SsrValue) *models.SiteSettingSsr {
	data := models.SiteSettingSsr{}

	data.ConductorHosts = misttransform.ListOfStringTerraformToSdk(d.ConductorHosts)
	data.DisableStats = d.DisableStats.ValueBoolPointer()

	return &data
}
