package resource_site_networktemplate

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfAreasNetworksTerraformToSdk(d basetypes.MapValue) map[string]models.OspfAreasNetwork {
	dataMap := make(map[string]models.OspfAreasNetwork)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(OspfNetworksValue)

		dataItem := models.OspfAreasNetwork{}
		if !itemObj.AuthKeys.IsNull() && !itemObj.AuthKeys.IsUnknown() {
			authKeysVm := make(map[string]string)
			for k, v := range itemObj.AuthKeys.Elements() {
				authKeysVm[k] = v.(basetypes.StringValue).ValueString()
			}
			dataItem.AuthKeys = authKeysVm
		}
		if !itemObj.AuthPassword.IsNull() && !itemObj.AuthPassword.IsUnknown() {
			dataItem.AuthPassword = itemObj.AuthPassword.ValueStringPointer()
		}
		if !itemObj.AuthType.IsNull() && !itemObj.AuthType.IsUnknown() {
			dataItem.AuthType = (*models.OspfAreaNetworkAuthTypeEnum)(itemObj.AuthType.ValueStringPointer())
		}
		if !itemObj.BfdMinimumInterval.IsNull() && !itemObj.BfdMinimumInterval.IsUnknown() {
			dataItem.BfdMinimumInterval = models.ToPointer(int(itemObj.BfdMinimumInterval.ValueInt64()))
		}
		if !itemObj.DeadInterval.IsNull() && !itemObj.DeadInterval.IsUnknown() {
			dataItem.DeadInterval = models.ToPointer(int(itemObj.DeadInterval.ValueInt64()))
		}
		if !itemObj.ExportPolicy.IsNull() && !itemObj.ExportPolicy.IsUnknown() {
			dataItem.ExportPolicy = itemObj.ExportPolicy.ValueStringPointer()
		}
		if !itemObj.HelloInterval.IsNull() && !itemObj.HelloInterval.IsUnknown() {
			dataItem.HelloInterval = models.ToPointer(int(itemObj.HelloInterval.ValueInt64()))
		}
		if !itemObj.ImportPolicy.IsNull() && !itemObj.ImportPolicy.IsUnknown() {
			dataItem.ImportPolicy = itemObj.ImportPolicy.ValueStringPointer()
		}
		if !itemObj.InterfaceType.IsNull() && !itemObj.InterfaceType.IsUnknown() {
			dataItem.InterfaceType = (*models.OspfAreaNetworkInterfaceTypeEnum)(itemObj.InterfaceType.ValueStringPointer())
		}
		if !itemObj.Metric.IsNull() && !itemObj.Metric.IsUnknown() {
			dataItem.Metric = models.NewOptional(models.ToPointer(int(itemObj.Metric.ValueInt64())))
		}
		if !itemObj.NoReadvertiseToOverlay.IsNull() && !itemObj.NoReadvertiseToOverlay.IsUnknown() {
			dataItem.NoReadvertiseToOverlay = itemObj.NoReadvertiseToOverlay.ValueBoolPointer()
		}
		if !itemObj.Passive.IsNull() && !itemObj.Passive.IsUnknown() {
			dataItem.Passive = itemObj.Passive.ValueBoolPointer()
		}

		dataMap[itemName] = dataItem
	}
	return dataMap
}

func ospfAreasTerraformToSdk(d basetypes.MapValue) map[string]models.OspfArea {
	data := make(map[string]models.OspfArea)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(OspfAreasValue)

		dataItem := models.OspfArea{}
		if !itemObj.IncludeLoopback.IsNull() && !itemObj.IncludeLoopback.IsUnknown() {
			dataItem.IncludeLoopback = itemObj.IncludeLoopback.ValueBoolPointer()
		}
		if !itemObj.OspfNetworks.IsNull() && !itemObj.OspfNetworks.IsUnknown() {
			dataItem.Networks = ospfAreasNetworksTerraformToSdk(itemObj.OspfNetworks)
		}
		if !itemObj.OspfAreasType.IsNull() && !itemObj.OspfAreasType.IsUnknown() {
			dataItem.Type = (*models.OspfAreaTypeEnum)(itemObj.OspfAreasType.ValueStringPointer())
		}
		data[itemName] = dataItem
	}
	return data
}
