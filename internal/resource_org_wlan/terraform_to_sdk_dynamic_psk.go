package resource_org_wlan

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func dynamicPskTerraformToSdk(plan DynamicPskValue) *models.WlanDynamicPsk {

	data := models.WlanDynamicPsk{}
	data.DefaultPsk = plan.DefaultPsk.ValueStringPointer()
	data.DefaultVlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.DefaultVlanId.ValueString()))
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.ForceLookup = plan.ForceLookup.ValueBoolPointer()
	if !plan.LocalVlanIds.IsNull() && !plan.LocalVlanIds.IsUnknown() {
		for _, v := range plan.LocalVlanIds.Elements() {
			var vInterface interface{} = v
			s := vInterface.(basetypes.StringValue)
			data.LocalVlanIds = append(data.LocalVlanIds, models.VlanIdWithVariableContainer.FromString(s.ValueString()))
		}
	}
	data.Source = models.ToPointer(models.DynamicPskSourceEnum(plan.Source.ValueString()))

	return &data
}
