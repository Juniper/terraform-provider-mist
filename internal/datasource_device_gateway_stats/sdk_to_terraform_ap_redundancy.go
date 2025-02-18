package datasource_device_gateway_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func moduleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ApRedundancyModule) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {

		var numAps basetypes.Int64Value
		var numApsWithSwitchRedundancy basetypes.Int64Value

		if d.NumAps != nil {
			numAps = types.Int64Value(int64(*d.NumAps))
		}
		if d.NumApsWithSwitchRedundancy != nil {
			numApsWithSwitchRedundancy = types.Int64Value(int64(*d.NumApsWithSwitchRedundancy))
		}

		dataMapValue := map[string]attr.Value{
			"num_aps":                        numAps,
			"num_aps_with_switch_redundancy": numApsWithSwitchRedundancy,
		}
		data, e := NewModulesValue(ModulesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, ModulesValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}

func apRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRedundancy) basetypes.ObjectValue {
	var modules = types.MapNull(ModuleStatValue{}.Type(ctx))
	var numAps basetypes.Int64Value
	var numApsWithSwitchRedundancy basetypes.Int64Value

	if d.Modules != nil && len(d.Modules) > 0 {
		modules = moduleSdkToTerraform(ctx, diags, d.Modules)
	}
	if d.NumAps != nil {
		numAps = types.Int64Value(int64(*d.NumAps))
	}
	if d.NumApsWithSwitchRedundancy != nil {
		numApsWithSwitchRedundancy = types.Int64Value(int64(*d.NumApsWithSwitchRedundancy))
	}

	dataMapValue := map[string]attr.Value{
		"modules":                        modules,
		"num_aps":                        numAps,
		"num_aps_with_switch_redundancy": numApsWithSwitchRedundancy,
	}
	data, e := types.ObjectValue(ApRedundancyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
