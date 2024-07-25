package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func extraRoute6NextQualifiedTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ExtraRoute6NextQualifiedProperties {
	data := make(map[string]models.ExtraRoute6NextQualifiedProperties)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(NextQualifiedValue)
		v_data := models.ExtraRoute6NextQualifiedProperties{}
		if v_plan.Metric.ValueInt64Pointer() != nil {
			v_data.Metric = models.NewOptional(models.ToPointer(int(v_plan.Metric.ValueInt64())))
		}
		if v_plan.Preference.ValueInt64Pointer() != nil {
			v_data.Preference = models.NewOptional(models.ToPointer(int(v_plan.Preference.ValueInt64())))
		}
		data[k] = v_data
	}
	return data
}
func extraRoutes6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.ExtraRoute6 {
	data := make(map[string]models.ExtraRoute6)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ExtraRoutesValue)

		v_data := models.ExtraRoute6{}
		if v_plan.Discard.ValueBoolPointer() != nil {
			v_data.Discard = models.ToPointer(v_plan.Discard.ValueBool())
		}
		if v_plan.Metric.ValueInt64Pointer() != nil {
			v_data.Metric = models.NewOptional(models.ToPointer(int(v_plan.Metric.ValueInt64())))
		}
		v_data.NextQualified = extraRoute6NextQualifiedTerraformToSdk(ctx, diags, v_plan.NextQualified)
		v_data.NoResolve = models.ToPointer(v_plan.NoResolve.ValueBool())
		if v_plan.Preference.ValueInt64Pointer() != nil {
			v_data.Preference = models.NewOptional(models.ToPointer(int(v_plan.Preference.ValueInt64())))
		}
		if v_plan.Via.ValueStringPointer() != nil {
			v_data.Via = models.ToPointer(v_plan.Via.ValueString())
		}
		data[k] = v_data
	}
	return data
}
