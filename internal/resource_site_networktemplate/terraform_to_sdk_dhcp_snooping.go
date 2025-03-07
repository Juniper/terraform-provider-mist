package resource_site_networktemplate

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func dhcpSnoopingTerraformToSdk(d DhcpSnoopingValue) *models.DhcpSnooping {
	data := models.DhcpSnooping{}
	if d.AllNetworks.ValueBoolPointer() != nil {
		data.AllNetworks = models.ToPointer(d.AllNetworks.ValueBool())
	}
	if d.EnableArpSpoofCheck.ValueBoolPointer() != nil {
		data.EnableArpSpoofCheck = models.ToPointer(d.EnableArpSpoofCheck.ValueBool())
	}
	if d.EnableIpSourceGuard.ValueBoolPointer() != nil {
		data.EnableIpSourceGuard = models.ToPointer(d.EnableIpSourceGuard.ValueBool())
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	if !d.Networks.IsNull() && !d.Networks.IsUnknown() {
		data.Networks = mistutils.ListOfStringTerraformToSdk(d.Networks)
	}
	return &data
}
