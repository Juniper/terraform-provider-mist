package resource_org_wlan

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ciscoCwaTerraformToSdk(plan CiscoCwaValue) *models.WlanCiscoCwa {

	data := models.WlanCiscoCwa{}
	if !plan.AllowedHostnames.IsNull() && !plan.AllowedHostnames.IsUnknown() {
		data.AllowedHostnames = mistutils.ListOfStringTerraformToSdk(plan.AllowedHostnames)
	}
	if !plan.AllowedSubnets.IsNull() && !plan.AllowedSubnets.IsUnknown() {
		data.AllowedSubnets = mistutils.ListOfStringTerraformToSdk(plan.AllowedSubnets)
	}
	if !plan.BlockedSubnets.IsNull() && !plan.BlockedSubnets.IsUnknown() {
		data.BlockedSubnets = mistutils.ListOfStringTerraformToSdk(plan.BlockedSubnets)
	}
	if plan.Enabled.ValueBoolPointer() != nil {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}
	return &data
}
