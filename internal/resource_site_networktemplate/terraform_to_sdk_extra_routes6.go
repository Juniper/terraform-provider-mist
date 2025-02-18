package resource_site_networktemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func extraRoute6NextQualifiedTerraformToSdk(d basetypes.MapValue) map[string]models.ExtraRoute6NextQualifiedProperties {
	data := make(map[string]models.ExtraRoute6NextQualifiedProperties)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(NextQualifiedValue)
		vData := models.ExtraRoute6NextQualifiedProperties{}
		if vPlan.Metric.ValueInt64Pointer() != nil {
			vData.Metric = models.NewOptional(models.ToPointer(int(vPlan.Metric.ValueInt64())))
		}
		if vPlan.Preference.ValueInt64Pointer() != nil {
			vData.Preference = models.NewOptional(models.ToPointer(int(vPlan.Preference.ValueInt64())))
		}
		data[k] = vData
	}
	return data
}
func extraRoutes6TerraformToSdk(d basetypes.MapValue) map[string]models.ExtraRoute6 {
	data := make(map[string]models.ExtraRoute6)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(ExtraRoutesValue)

		vData := models.ExtraRoute6{}
		if vPlan.Discard.ValueBoolPointer() != nil {
			vData.Discard = models.ToPointer(vPlan.Discard.ValueBool())
		}
		if vPlan.Metric.ValueInt64Pointer() != nil {
			vData.Metric = models.NewOptional(models.ToPointer(int(vPlan.Metric.ValueInt64())))
		}
		vData.NextQualified = extraRoute6NextQualifiedTerraformToSdk(vPlan.NextQualified)
		vData.NoResolve = models.ToPointer(vPlan.NoResolve.ValueBool())
		if vPlan.Preference.ValueInt64Pointer() != nil {
			vData.Preference = models.NewOptional(models.ToPointer(int(vPlan.Preference.ValueInt64())))
		}
		if vPlan.Via.ValueStringPointer() != nil {
			vData.Via = models.ToPointer(vPlan.Via.ValueString())
		}
		data[k] = vData
	}
	return data
}
