package resource_site_networktemplate

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
		itemObj := itemInterface.(VrfExtraRoutesValue)

		var dataItem models.VrfExtraRoute
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
		if !itemObj.MulticastConfig.IsNull() && !itemObj.MulticastConfig.IsUnknown() {
			attrs := itemObj.MulticastConfig.Attributes()
			mc := models.SwitchMulticastConfig{}
			if v, ok := attrs["anycast_rp"].(basetypes.BoolValue); ok && !v.IsNull() && !v.IsUnknown() {
				mc.AnycastRp = v.ValueBoolPointer()
			}
			if v, ok := attrs["rp_ip"].(basetypes.StringValue); ok && !v.IsNull() && !v.IsUnknown() {
				mc.RpIp = v.ValueStringPointer()
			}
			if v, ok := attrs["sbd_subnet"].(basetypes.StringValue); ok && !v.IsNull() && !v.IsUnknown() {
				mc.SbdSubnet = v.ValueStringPointer()
			}
			if v, ok := attrs["sbd_vlan_id"].(basetypes.Int64Value); ok && !v.IsNull() && !v.IsUnknown() {
				mc.SbdVlanId = models.ToPointer(int(v.ValueInt64()))
			}
			dataItem.MulticastConfig = &mc
		}
		data[itemName] = dataItem
	}
	return data
}
