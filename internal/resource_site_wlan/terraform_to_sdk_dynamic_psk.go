package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func dynamicPskTerraformToSdk(plan DynamicPskValue) *models.WlanDynamicPsk {

	data := models.WlanDynamicPsk{}
	data.DefaultPsk = plan.DefaultPsk.ValueStringPointer()
	data.DefaultVlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.DefaultVlanId.ValueString()))
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.ForceLookup = plan.ForceLookup.ValueBoolPointer()
	data.Source = models.ToPointer(models.DynamicPskSourceEnum(plan.Source.ValueString()))

	return &data
}
