package resource_org_network

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func groupMulticastTerraformToSdk(d basetypes.MapValue) map[string]models.NetworkMulticastGroup {
	dataMap := make(map[string]models.NetworkMulticastGroup)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(GroupsValue)
		data := models.NetworkMulticastGroup{}
		data.RpIp = vPlan.RpIp.ValueStringPointer()
		dataMap[k] = data
	}
	return dataMap
}

func MulticastTerraformToSdk(d MulticastValue) *models.NetworkMulticast {
	data := models.NetworkMulticast{}

	if d.DisableIgmp.ValueBoolPointer() != nil {
		data.DisableIgmp = d.DisableIgmp.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Groups.IsNull() && !d.Groups.IsUnknown() {
		data.Groups = groupMulticastTerraformToSdk(d.Groups)
	}

	return &data
}
