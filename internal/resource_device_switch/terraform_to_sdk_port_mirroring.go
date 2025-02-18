package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func portMirroringTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchPortMirroringProperty {
	data := make(map[string]models.SwitchPortMirroringProperty)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(PortMirroringValue)

		dataItem := models.SwitchPortMirroringProperty{}
		if !itemObj.InputNetworksIngress.IsNull() && !itemObj.InputNetworksIngress.IsUnknown() {
			dataItem.InputNetworksIngress = misttransform.ListOfStringTerraformToSdk(itemObj.InputNetworksIngress)
		}
		if dataItem.InputNetworksIngress == nil {
			dataItem.InputNetworksIngress = make([]string, 0)
		}

		if !itemObj.InputPortIdsEgress.IsNull() && !itemObj.InputPortIdsEgress.IsUnknown() {
			dataItem.InputPortIdsEgress = misttransform.ListOfStringTerraformToSdk(itemObj.InputPortIdsEgress)
		}
		if dataItem.InputPortIdsEgress == nil {
			dataItem.InputPortIdsEgress = make([]string, 0)
		}

		if !itemObj.InputPortIdsIngress.IsNull() && !itemObj.InputPortIdsIngress.IsUnknown() {
			dataItem.InputPortIdsIngress = misttransform.ListOfStringTerraformToSdk(itemObj.InputPortIdsIngress)
		}
		if dataItem.InputPortIdsIngress == nil {
			dataItem.InputPortIdsIngress = make([]string, 0)
		}

		if itemObj.OutputNetwork.ValueStringPointer() != nil {
			dataItem.OutputNetwork = models.ToPointer(itemObj.OutputNetwork.ValueString())
		}
		if itemObj.OutputPortId.ValueStringPointer() != nil {
			dataItem.OutputPortId = models.ToPointer(itemObj.OutputPortId.ValueString())
		}

		data[itemName] = dataItem
	}
	return data
}
