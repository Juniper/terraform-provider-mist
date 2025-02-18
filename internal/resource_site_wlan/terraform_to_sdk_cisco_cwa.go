package resource_site_wlan

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ciscoCwaTerraformToSdk(plan CiscoCwaValue) *models.WlanCiscoCwa {

	data := models.WlanCiscoCwa{}
	if !plan.AllowedHostnames.IsNull() && !plan.AllowedHostnames.IsUnknown() {
		data.AllowedHostnames = misttransform.ListOfStringTerraformToSdk(plan.AllowedHostnames)
	}
	if !plan.AllowedSubnets.IsNull() && !plan.AllowedSubnets.IsUnknown() {
		data.AllowedSubnets = misttransform.ListOfStringTerraformToSdk(plan.AllowedSubnets)
	}
	if !plan.BlockedSubnets.IsNull() && !plan.BlockedSubnets.IsUnknown() {
		data.BlockedSubnets = misttransform.ListOfStringTerraformToSdk(plan.BlockedSubnets)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	return &data
}
