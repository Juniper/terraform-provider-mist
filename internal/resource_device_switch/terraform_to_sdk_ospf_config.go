package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfConfigAreasTerraformToSdk(d basetypes.MapValue) map[string]models.SwitchOspfConfigArea {
	dataMap := make(map[string]models.SwitchOspfConfigArea)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(AreasValue)

		dataItem := models.SwitchOspfConfigArea{}

		if !itemObj.NoSummary.IsNull() && !itemObj.NoSummary.IsUnknown() {
			dataItem.NoSummary = itemObj.NoSummary.ValueBoolPointer()
		}

		dataMap[itemName] = dataItem
	}
	return dataMap
}

func ospfConfigTerraformToSdk(d OspfConfigValue) *models.SwitchOspfConfig {

	data := models.SwitchOspfConfig{}

	if !d.Areas.IsNull() && !d.Areas.IsUnknown() {
		data.Areas = ospfConfigAreasTerraformToSdk(d.Areas)
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.ExportPolicy.ValueStringPointer() != nil {
		data.ExportPolicy = models.ToPointer(d.ExportPolicy.ValueString())
	}
	if d.ImportPolicy.ValueStringPointer() != nil {
		data.ImportPolicy = models.ToPointer(d.ImportPolicy.ValueString())
	}
	if d.ReferenceBandwidth.ValueStringPointer() != nil {
		data.ReferenceBandwidth = models.ToPointer((models.SwitchOspfConfigReferenceBandwidthContainer.FromString(d.ReferenceBandwidth.ValueString())))
	}

	return &data
}
