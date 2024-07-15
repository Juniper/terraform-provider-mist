package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func modelSpecificSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.RfTemplateModelSpecificProperty) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var ant_gain_24 basetypes.Int64Value
		var ant_gain_5 basetypes.Int64Value
		var ant_gain_6 basetypes.Int64Value
		var band_24 basetypes.ObjectValue = types.ObjectNull(Band24Value{}.AttributeTypes(ctx))
		var band_24_usage basetypes.StringValue
		var band_5 basetypes.ObjectValue = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
		var band_5_on_24_radio basetypes.ObjectValue = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
		var band_6 basetypes.ObjectValue = types.ObjectNull(Band6Value{}.AttributeTypes(ctx))

		if d.AntGain24 != nil {
			ant_gain_24 = types.Int64Value(int64(*d.AntGain24))
		}
		if d.AntGain5 != nil {
			ant_gain_5 = types.Int64Value(int64(*d.AntGain5))
		}
		if d.AntGain6 != nil {
			ant_gain_6 = types.Int64Value(int64(*d.AntGain6))
		}
		if d.Band24 != nil {
			band_24, _ = band24SdkToTerraform(ctx, diags, d.Band24).ToObjectValue(ctx)
		}
		if d.Band24Usage != nil {
			band_24_usage = types.StringValue(string(*d.Band24Usage))
		}
		if d.Band5 != nil {
			band_5, _ = band5SdkToTerraform(ctx, diags, d.Band5).ToObjectValue(ctx)
		}
		if d.Band5On24Radio != nil {
			band_5_on_24_radio, _ = band5SdkToTerraform(ctx, diags, d.Band5On24Radio).ToObjectValue(ctx)
		}
		if d.Band6 != nil {
			band_6, _ = band6SdkToTerraform(ctx, diags, d.Band6).ToObjectValue(ctx)
		}

		data_map_attr_type := ModelSpecificValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ant_gain_24":        ant_gain_24,
			"ant_gain_5":         ant_gain_5,
			"ant_gain_6":         ant_gain_6,
			"band_24":            band_24,
			"band_24_usage":      band_24_usage,
			"band_5":             band_5,
			"band_5_on_24_radio": band_5_on_24_radio,
			"band_6":             band_6,
		}
		data, e := NewModelSpecificValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := ModelSpecificValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
