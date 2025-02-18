package resource_device_switch

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func extraRouteNextQualifiedTerraformToSdk(d basetypes.MapValue) map[string]models.ExtraRouteNextQualifiedProperties {
	data := make(map[string]models.ExtraRouteNextQualifiedProperties)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(NextQualifiedValue)
		vData := models.ExtraRouteNextQualifiedProperties{}
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
func extraRoutesTerraformToSdk(d basetypes.MapValue) map[string]models.ExtraRoute {
	data := make(map[string]models.ExtraRoute)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(ExtraRoutesValue)

		vData := models.ExtraRoute{}
		if vPlan.Discard.ValueBoolPointer() != nil {
			vData.Discard = models.ToPointer(vPlan.Discard.ValueBool())
		}
		if vPlan.Metric.ValueInt64Pointer() != nil {
			vData.Metric = models.NewOptional(models.ToPointer(int(vPlan.Metric.ValueInt64())))
		}
		vData.NextQualified = extraRouteNextQualifiedTerraformToSdk(vPlan.NextQualified)
		if vPlan.NoResolve.ValueBoolPointer() != nil {
			vData.NoResolve = models.ToPointer(vPlan.NoResolve.ValueBool())
		}
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
