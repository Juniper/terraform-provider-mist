package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsTerraformToSdk(plan basetypes.MapValue) map[string]models.WlanAppQosAppsProperties {
	dataMap := make(map[string]models.WlanAppQosAppsProperties)
	for k, v := range plan.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(AppsValue)
		data := models.WlanAppQosAppsProperties{}
		if vPlan.Dscp.ValueInt64Pointer() != nil {
			data.Dscp = models.ToPointer(int(vPlan.Dscp.ValueInt64()))
		}
		if vPlan.DstSubnet.ValueStringPointer() != nil {
			data.DstSubnet = vPlan.DstSubnet.ValueStringPointer()
		}
		if vPlan.SrcSubnet.ValueStringPointer() != nil {
			data.SrcSubnet = vPlan.SrcSubnet.ValueStringPointer()
		}
		dataMap[k] = data
	}
	return dataMap
}
func appQosOthersTerraformToSdk(plan basetypes.ListValue) []models.WlanAppQosOthersItem {
	var dataList []models.WlanAppQosOthersItem
	for _, v := range plan.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(OthersValue)
		data := models.WlanAppQosOthersItem{}
		if vPlan.Dscp.ValueInt64Pointer() != nil {
			data.Dscp = models.ToPointer(int(vPlan.Dscp.ValueInt64()))
		}
		if vPlan.DstSubnet.ValueStringPointer() != nil {
			data.DstSubnet = vPlan.DstSubnet.ValueStringPointer()
		}
		if vPlan.PortRanges.ValueStringPointer() != nil {
			data.PortRanges = vPlan.PortRanges.ValueStringPointer()
		}
		if vPlan.Protocol.ValueStringPointer() != nil {
			data.Protocol = vPlan.Protocol.ValueStringPointer()
		}
		if vPlan.SrcSubnet.ValueStringPointer() != nil {
			data.SrcSubnet = vPlan.SrcSubnet.ValueStringPointer()
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func appQosTerraformToSdk(plan AppQosValue) *models.WlanAppQos {

	data := models.WlanAppQos{}

	apps := appQosAppsTerraformToSdk(plan.Apps)
	data.Apps = apps

	data.Enabled = plan.Enabled.ValueBoolPointer()

	others := appQosOthersTerraformToSdk(plan.Others)
	data.Others = others

	return &data
}
