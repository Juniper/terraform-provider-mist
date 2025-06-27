package resource_device_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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
		itemObj := itemInterface.(ExtraRoutesValue)

		dataItem := models.VrfExtraRoute{}
		if itemObj.Via.ValueStringPointer() != nil {
			dataItem.Via = models.ToPointer(itemObj.Via.ValueString())
		}
		data[itemName] = dataItem
	}
	return data
}

func vrfInstanceExtraRoute6TerraformToSdk(d basetypes.MapValue) map[string]models.VrfExtraRoute {
	data := make(map[string]models.VrfExtraRoute)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(VrfExtraRoutes6Value)

		var dataItem models.VrfExtraRoute
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

		var dataItem models.SwitchVrfInstance
		if !itemObj.EvpnAutoLoopbackSubnet.IsNull() && !itemObj.EvpnAutoLoopbackSubnet.IsUnknown() {
			dataItem.EvpnAutoLoopbackSubnet = itemObj.EvpnAutoLoopbackSubnet.ValueStringPointer()
		}
		if !itemObj.EvpnAutoLoopbackSubnet6.IsNull() && !itemObj.EvpnAutoLoopbackSubnet6.IsUnknown() {
			dataItem.EvpnAutoLoopbackSubnet6 = itemObj.EvpnAutoLoopbackSubnet6.ValueStringPointer()
		}
		if !itemObj.Networks.IsNull() && !itemObj.Networks.IsUnknown() {
			dataItem.Networks = mistutils.ListOfStringTerraformToSdk(itemObj.Networks)
		}
		if !itemObj.VrfExtraRoutes.IsNull() && !itemObj.VrfExtraRoutes.IsUnknown() {
			dataItem.ExtraRoutes = vrfInstanceExtraRouteTerraformToSdk(itemObj.VrfExtraRoutes)
		}
		if !itemObj.VrfExtraRoutes6.IsNull() && !itemObj.VrfExtraRoutes6.IsUnknown() {
			dataItem.ExtraRoutes = vrfInstanceExtraRoute6TerraformToSdk(itemObj.VrfExtraRoutes6)
		}

		data[itemName] = dataItem
	}
	return data
}
