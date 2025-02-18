package resource_org_networktemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vrfConfigTerraformToSdk(d VrfConfigValue) *models.VrfConfig {
	data := models.VrfConfig{}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = models.ToPointer(d.Enabled.ValueBool())
	}
	return &data
}

func vrfInstanceExtraRouteTerraformToSdk(d basetypes.MapValue) map[string]models.VrfExtraRoute {
	data := make(map[string]models.VrfExtraRoute)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(VrfExtraRoutesValue)

		dataItem := models.VrfExtraRoute{}
		if itemObj.Via.ValueStringPointer() != nil {
			dataItem.Via = models.ToPointer(itemObj.Via.ValueString())
		}
		data[itemName] = dataItem
	}
	return data
}

func vrfInstancesTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchVrfInstance {
	data := make(map[string]models.SwitchVrfInstance)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(VrfInstancesValue)

		dataItem := models.SwitchVrfInstance{}
		if !itemObj.Networks.IsNull() && !itemObj.Networks.IsUnknown() {
			dataItem.Networks = misttransform.ListOfStringTerraformToSdk(itemObj.Networks)
		}
		if !itemObj.VrfExtraRoutes.IsNull() && !itemObj.VrfExtraRoutes.IsUnknown() {
			dataItem.ExtraRoutes = vrfInstanceExtraRouteTerraformToSdk(itemObj.VrfExtraRoutes)
		}

		data[itemName] = dataItem
	}
	return data
}
