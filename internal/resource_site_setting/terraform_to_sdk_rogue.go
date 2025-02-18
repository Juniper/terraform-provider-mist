package resource_site_setting

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func rogueTerraformToSdk(d RogueValue) *models.SiteRogue {
	data := models.SiteRogue{}

	data.Enabled = d.Enabled.ValueBoolPointer()
	data.HoneypotEnabled = d.HoneypotEnabled.ValueBoolPointer()
	data.MinDuration = models.ToPointer(int(d.MinDuration.ValueInt64()))
	data.MinRssi = models.ToPointer(int(d.MinRssi.ValueInt64()))
	data.WhitelistedBssids = misttransform.ListOfStringTerraformToSdk(d.WhitelistedBssids)
	data.WhitelistedSsids = misttransform.ListOfStringTerraformToSdk(d.WhitelistedSsids)

	return &data
}
