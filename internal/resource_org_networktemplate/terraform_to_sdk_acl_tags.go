package resource_org_networktemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func actTagSpecsTerraformToSdk(d basetypes.ListValue) []models.AclTagSpec {
	var data []models.AclTagSpec
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		vState := vInterface.(SpecsValue)
		vData := models.AclTagSpec{}
		if vState.PortRange.ValueStringPointer() != nil {
			vData.PortRange = models.ToPointer(vState.PortRange.ValueString())
		}
		if vState.Protocol.ValueStringPointer() != nil {
			vData.Protocol = models.ToPointer(vState.Protocol.ValueString())
		}
		data = append(data, vData)
	}
	return data
}

func actTagsTerraformToSdk(d basetypes.MapValue) map[string]models.AclTag {
	data := make(map[string]models.AclTag)
	for itemName, itemValue := range d.Elements() {
		var itemInterface interface{} = itemValue
		itemObj := itemInterface.(AclTagsValue)

		dataItem := models.AclTag{}
		dataItem.Type = models.AclTagTypeEnum(itemObj.AclTagsType.ValueString())
		if itemObj.GbpTag.ValueInt64Pointer() != nil {
			dataItem.GbpTag = models.ToPointer(int(itemObj.GbpTag.ValueInt64()))
		}
		dataItem.Macs = mistutils.ListOfStringTerraformToSdk(itemObj.Macs)
		if itemObj.Network.ValueStringPointer() != nil {
			dataItem.Network = models.ToPointer(itemObj.Network.ValueString())
		}
		if itemObj.RadiusGroup.ValueStringPointer() != nil {
			dataItem.RadiusGroup = models.ToPointer(itemObj.RadiusGroup.ValueString())
		}
		if !itemObj.Specs.IsNull() && !itemObj.Specs.IsUnknown() {
			dataItem.Specs = actTagSpecsTerraformToSdk(itemObj.Specs)
		}
		if !itemObj.Subnets.IsNull() && !itemObj.Subnets.IsUnknown() {
			dataItem.Subnets = mistutils.ListOfStringTerraformToSdk(itemObj.Subnets)
		}
		data[itemName] = dataItem
	}
	return data
}
