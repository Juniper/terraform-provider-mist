package resource_org_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ssrTerraformToSdk(d SsrValue) *models.SettingSsr {
	data := models.SettingSsr{}

	if !d.ConductorHosts.IsNull() && !d.ConductorHosts.IsUnknown() {
		data.ConductorHosts = mistutils.ListOfStringTerraformToSdk(d.ConductorHosts)
	}
	if d.ConductorToken.ValueStringPointer() != nil {
		data.ConductorToken = d.ConductorToken.ValueStringPointer()
	}
	if d.DisableStats.ValueBoolPointer() != nil {
		data.DisableStats = d.DisableStats.ValueBoolPointer()
	}

	return &data
}
