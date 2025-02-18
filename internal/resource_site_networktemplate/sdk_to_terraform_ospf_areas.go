package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfAreasNetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OspfAreasNetwork) basetypes.MapValue {
	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var authKeys = types.MapNull(types.StringType)
		var authPassword basetypes.StringValue
		var authType basetypes.StringValue
		var bfdMinimumInterval basetypes.Int64Value
		var deadInterval basetypes.Int64Value
		var exportPolicy basetypes.StringValue
		var helloInterval basetypes.Int64Value
		var importPolicy basetypes.StringValue
		var interfaceType basetypes.StringValue
		var metric basetypes.Int64Value
		var noReadvertiseToOverlay basetypes.BoolValue
		var passive basetypes.BoolValue

		if d.AuthKeys != nil {
			authKeysVm := make(map[string]string)
			for k, v := range d.AuthKeys {
				authKeysVm[k] = v
			}
			authKeys, _ = types.MapValueFrom(ctx, types.StringType, authKeysVm)
		}
		if d.AuthPassword != nil {
			authPassword = types.StringValue(*d.AuthPassword)
		}
		if d.AuthType != nil {
			authType = types.StringValue(string(*d.AuthType))
		}
		if d.BfdMinimumInterval != nil {
			bfdMinimumInterval = types.Int64Value(int64(*d.BfdMinimumInterval))
		}
		if d.DeadInterval != nil {
			deadInterval = types.Int64Value(int64(*d.DeadInterval))
		}
		if d.ExportPolicy != nil {
			exportPolicy = types.StringValue(*d.ExportPolicy)
		}
		if d.HelloInterval != nil {
			helloInterval = types.Int64Value(int64(*d.HelloInterval))
		}
		if d.ImportPolicy != nil {
			importPolicy = types.StringValue(*d.ImportPolicy)
		}
		if d.InterfaceType != nil {
			interfaceType = types.StringValue(string(*d.InterfaceType))
		}
		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.NoReadvertiseToOverlay != nil {
			noReadvertiseToOverlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.Passive != nil {
			passive = types.BoolValue(*d.Passive)
		}

		dataMapValue := map[string]attr.Value{
			"auth_keys":                 authKeys,
			"auth_password":             authPassword,
			"auth_type":                 authType,
			"bfd_minimum_interval":      bfdMinimumInterval,
			"dead_interval":             deadInterval,
			"export_policy":             exportPolicy,
			"hello_interval":            helloInterval,
			"import_policy":             importPolicy,
			"interface_type":            interfaceType,
			"metric":                    metric,
			"no_readvertise_to_overlay": noReadvertiseToOverlay,
			"passive":                   passive,
		}
		data, e := NewOspfNetworksValue(OspfNetworksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data
	}
	stateResultMapType := OspfNetworksValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}

func ospfAreasSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OspfArea) basetypes.MapValue {

	stateValueMapValue := make(map[string]attr.Value)
	for k, d := range m {

		var includeLoopback basetypes.BoolValue
		var networks = types.MapNull(OspfNetworksValue{}.Type(ctx))
		var areaType basetypes.StringValue

		if d.IncludeLoopback != nil {
			includeLoopback = types.BoolValue(*d.IncludeLoopback)
		}
		if d.Networks != nil {
			networks = ospfAreasNetworksSdkToTerraform(ctx, diags, d.Networks)
		}
		if d.Type != nil {
			areaType = types.StringValue(string(*d.Type))
		}

		dataMapValue := map[string]attr.Value{
			"include_loopback": includeLoopback,
			"networks":         networks,
			"type":             areaType,
		}
		data, e := NewOspfAreasValue(OspfAreasValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMapValue[k] = data

	}
	stateResultMapType := OspfAreasValue{}.Type(ctx)
	stateResultMap, e := types.MapValueFrom(ctx, stateResultMapType, stateValueMapValue)
	diags.Append(e...)
	return stateResultMap
}
