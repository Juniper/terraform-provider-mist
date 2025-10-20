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

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var antGain24 basetypes.Int64Value
		var antGain5 basetypes.Int64Value
		var antGain6 basetypes.Int64Value
		var band24 = types.ObjectNull(Band24Value{}.AttributeTypes(ctx))
		var band24Usage basetypes.StringValue
		var band5 = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
		var band5On24Radio = types.ObjectNull(Band5On24RadioValue{}.AttributeTypes(ctx))
		var band6 = types.ObjectNull(Band6Value{}.AttributeTypes(ctx))

		if d.AntGain24 != nil {
			antGain24 = types.Int64Value(int64(*d.AntGain24))
		}
		if d.AntGain5 != nil {
			antGain5 = types.Int64Value(int64(*d.AntGain5))
		}
		if d.AntGain6 != nil {
			antGain6 = types.Int64Value(int64(*d.AntGain6))
		}
		if d.Band24 != nil {
			band24, _ = band24SdkToTerraform(ctx, diags, d.Band24).ToObjectValue(ctx)
		}
		if d.Band24Usage != nil {
			band24Usage = types.StringValue(string(*d.Band24Usage))
		}
		if d.Band5 != nil {
			band5, _ = band5SdkToTerraform(ctx, diags, d.Band5).ToObjectValue(ctx)
		}
		if d.Band5On24Radio != nil {
			band5On24Radio, _ = band5On24RadioSdkToTerraform(ctx, diags, d.Band5On24Radio).ToObjectValue(ctx)
		}
		if d.Band6 != nil {
			band6, _ = band6SdkToTerraform(ctx, diags, d.Band6).ToObjectValue(ctx)
		}

		dataMapValue := map[string]attr.Value{
			"ant_gain_24":        antGain24,
			"ant_gain_5":         antGain5,
			"ant_gain_6":         antGain6,
			"band_24":            band24,
			"band_24_usage":      band24Usage,
			"band_5":             band5,
			"band_5_on_24_radio": band5On24Radio,
			"band_6":             band6,
		}
		data, e := NewModelSpecificValue(ModelSpecificValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := ModelSpecificValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
