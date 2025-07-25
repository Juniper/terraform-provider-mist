package resource_device_switch

import (
	"context"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfConfigAreasSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchOspfConfigArea) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var noSummary basetypes.BoolValue

		if d.NoSummary != nil {
			noSummary = types.BoolValue(*d.NoSummary)
		}

		dataMapValue := map[string]attr.Value{
			"no_summary": noSummary,
		}
		data, e := NewAreasValue(AreasValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMap, e := types.MapValueFrom(ctx, AreasValue{}.Type(ctx), stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func ospfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchOspfConfig) OspfConfigValue {
	var areas = types.MapNull(AreasValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var referenceBandwidth basetypes.StringValue

	if d != nil && d.Areas != nil {
		areas = ospfConfigAreasSdkToTerraform(ctx, diags, d.Areas)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if referenceBandwidthInt, ok := d.ReferenceBandwidth.AsNumber(); ok {
		referenceBandwidth = types.StringValue(strconv.FormatInt(int64(*referenceBandwidthInt), 10))
	} else if referenceBandwidthStr, ok := d.ReferenceBandwidth.AsString(); ok {
		referenceBandwidth = types.StringValue(*referenceBandwidthStr)
	}

	dataMapValue := map[string]attr.Value{
		"areas":               areas,
		"enabled":             enabled,
		"reference_bandwidth": referenceBandwidth,
	}
	data, e := NewOspfConfigValue(OspfConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
