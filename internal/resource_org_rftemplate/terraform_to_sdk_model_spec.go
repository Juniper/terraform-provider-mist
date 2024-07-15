package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func modelSpecificTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.RfTemplateModelSpecificProperty {

	data_map := make(map[string]models.RfTemplateModelSpecificProperty)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(ModelSpecificValue)
		data := models.RfTemplateModelSpecificProperty{}

		if plan.AntGain24.ValueInt64Pointer() != nil {
			data.AntGain24 = models.ToPointer(int(plan.AntGain24.ValueInt64()))
		}
		if plan.AntGain5.ValueInt64Pointer() != nil {
			data.AntGain5 = models.ToPointer(int(plan.AntGain5.ValueInt64()))
		}
		if plan.AntGain6.ValueInt64Pointer() != nil {
			data.AntGain6 = models.ToPointer(int(plan.AntGain6.ValueInt64()))
		}

		if !plan.Band24.IsNull() && !plan.Band24.IsUnknown() {
			plan_band24, _ := NewBand24Value(plan.Band24.AttributeTypes(ctx), plan.Band24.Attributes())
			data.Band24 = band24TerraformToSdk(ctx, diags, plan_band24)

			data.Band24Usage = models.ToPointer(models.RadioBand24UsageEnum(plan.Band24Usage.ValueString()))
		}

		if !plan.Band5.IsNull() && !plan.Band5.IsUnknown() {
			plan_band5, _ := NewBand5Value(plan.Band5.AttributeTypes(ctx), plan.Band5.Attributes())
			data.Band5 = band5TerraformToSdk(ctx, diags, plan_band5)

		}
		if !plan.Band6.IsNull() && !plan.Band6.IsUnknown() {
			plan_band6, _ := NewBand6Value(plan.Band6.AttributeTypes(ctx), plan.Band6.Attributes())
			data.Band6 = band6TerraformToSdk(ctx, diags, plan_band6)

			data_map[k] = data
		}
	}
	return data_map
}
