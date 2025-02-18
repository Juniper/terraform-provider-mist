package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanTerraformToSdk(plan DynamicVlanValue) *models.WlanDynamicVlan {

	var localVlanIds []models.VlanIdWithVariable
	for _, item := range plan.LocalVlanIds.Elements() {
		var itemInterface interface{} = item
		i := itemInterface.(basetypes.StringValue)
		j := models.VlanIdWithVariableContainer.FromString(i.ValueString())
		localVlanIds = append(localVlanIds, j)
	}

	vlans := make(map[string]string)
	for k, v := range plan.Vlans.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(basetypes.StringValue)
		vlans[k] = vPlan.ValueString()
	}

	var defaultVlanIds []models.WlanDynamicVlanDefaultVlanId
	for _, item := range plan.DefaultVlanIds.Elements() {
		var itemInterface interface{} = item
		i := itemInterface.(basetypes.StringValue)
		j := models.WlanDynamicVlanDefaultVlanIdContainer.FromString(i.ValueString())
		defaultVlanIds = append(defaultVlanIds, j)
	}

	data := models.WlanDynamicVlan{}
	data.DefaultVlanIds = defaultVlanIds
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.LocalVlanIds = localVlanIds
	data.Type = models.ToPointer(models.WlanDynamicVlanTypeEnum(plan.DynamicVlanType.ValueString()))
	data.Vlans = vlans

	return &data
}
