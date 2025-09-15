package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func modelSpecificTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.RfTemplateModelSpecificProperty {

	dataMap := make(map[string]models.RfTemplateModelSpecificProperty)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(ModelSpecificValue)
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
			planBand24, _ := NewBand24Value(plan.Band24.AttributeTypes(ctx), plan.Band24.Attributes())
			data.Band24 = band24TerraformToSdk(planBand24)

			data.Band24Usage = models.ToPointer(models.RadioBand24UsageEnum(plan.Band24Usage.ValueString()))
		}

		if !plan.Band5.IsNull() && !plan.Band5.IsUnknown() {
			planBand5, _ := NewBand5Value(plan.Band5.AttributeTypes(ctx), plan.Band5.Attributes())
			data.Band5 = band5TerraformToSdk(planBand5)

		}

		if !plan.Band5On24Radio.IsNull() && !plan.Band5On24Radio.IsUnknown() {
			planBand5On24Radio, _ := NewBand5On24RadioValue(plan.Band5On24Radio.AttributeTypes(ctx), plan.Band5On24Radio.Attributes())
			data.Band5On24Radio = band5On24RadioTerraformToSdk(planBand5On24Radio)
		}

		if !plan.Band6.IsNull() && !plan.Band6.IsUnknown() {
			planBand6, _ := NewBand6Value(plan.Band6.AttributeTypes(ctx), plan.Band6.Attributes())
			data.Band6 = band6TerraformToSdk(planBand6)
		}

		dataMap[k] = data
	}
	return dataMap
}
