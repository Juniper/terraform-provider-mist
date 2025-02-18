package resource_site_evpn_topology

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchTerraformToSdk(d basetypes.MapValue) []models.EvpnTopologySwitch {

	var data []models.EvpnTopologySwitch
	for mac, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(SwitchesValue)
		dataItem := models.EvpnTopologySwitch{}
		dataItem.Mac = mac

		if !plan.Pod.IsNull() && !plan.Pod.IsUnknown() {
			dataItem.Pod = models.ToPointer(int(plan.Pod.ValueInt64()))
		}
		if !plan.Pods.IsNull() && !plan.Pods.IsUnknown() {
			dataItem.Pods = misttransform.ListOfIntTerraformToSdk(plan.Pods)
		}
		if !plan.Role.IsNull() && !plan.Role.IsUnknown() {
			dataItem.Role = models.EvpnTopologySwitchRoleEnum(plan.Role.ValueString())
		}

		data = append(data, dataItem)
	}
	return data
}
