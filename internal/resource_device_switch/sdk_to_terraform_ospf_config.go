package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfAreasConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OspfConfigArea) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var no_summary basetypes.BoolValue

		if d.NoSummary != nil {
			no_summary = types.BoolValue(*d.NoSummary)
		}

		data_map_attr_type := AreasValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"no_summary": no_summary,
		}
		data, e := NewAreasValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := AreasValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func ospfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.OspfConfig) OspfConfigValue {

	var areas basetypes.MapValue = types.MapNull(AreasValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var reference_bandwidth basetypes.StringValue

	if d.Areas != nil {
		areas = ospfAreasConfigSdkToTerraform(ctx, diags, d.Areas)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.ReferenceBandwidth != nil {
		reference_bandwidth = types.StringValue(*d.ReferenceBandwidth)
	}

	data_map_attr_type := OspfConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"areas":               areas,
		"enabled":             enabled,
		"reference_bandwidth": reference_bandwidth,
	}
	data, e := NewOspfConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
