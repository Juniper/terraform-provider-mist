package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ospfAreasNetworksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OspfAreasNetwork) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var auth_keys basetypes.MapValue = types.MapNull(types.StringType)
		var auth_password basetypes.StringValue
		var auth_type basetypes.StringValue
		var bfd_minimum_interval basetypes.Int64Value
		var dead_interval basetypes.Int64Value
		var export_policy basetypes.StringValue
		var hello_interval basetypes.Int64Value
		var import_policy basetypes.StringValue
		var interface_type basetypes.StringValue
		var metric basetypes.Int64Value
		var no_readvertise_to_overlay basetypes.BoolValue
		var passive basetypes.BoolValue

		if d.AuthKeys != nil {
			auth_keys_vm := make(map[string]string)
			for k, v := range d.AuthKeys {
				auth_keys_vm[k] = v
			}
			auth_keys, _ = types.MapValueFrom(ctx, types.StringType, auth_keys_vm)
		}
		if d.AuthPassword != nil {
			auth_password = types.StringValue(*d.AuthPassword)
		}
		if d.AuthType != nil {
			auth_type = types.StringValue(string(*d.AuthType))
		}
		if d.BfdMinimumInterval != nil {
			bfd_minimum_interval = types.Int64Value(int64(*d.BfdMinimumInterval))
		}
		if d.DeadInterval != nil {
			dead_interval = types.Int64Value(int64(*d.DeadInterval))
		}
		if d.ExportPolicy != nil {
			export_policy = types.StringValue(*d.ExportPolicy)
		}
		if d.HelloInterval != nil {
			hello_interval = types.Int64Value(int64(*d.HelloInterval))
		}
		if d.ImportPolicy != nil {
			import_policy = types.StringValue(*d.ImportPolicy)
		}
		if d.InterfaceType != nil {
			interface_type = types.StringValue(string(*d.InterfaceType))
		}
		if d.Metric.Value() != nil {
			metric = types.Int64Value(int64(*d.Metric.Value()))
		}
		if d.NoReadvertiseToOverlay != nil {
			no_readvertise_to_overlay = types.BoolValue(*d.NoReadvertiseToOverlay)
		}
		if d.Passive != nil {
			passive = types.BoolValue(*d.Passive)
		}

		data_map_attr_type := OspfNetworksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"auth_keys":                 auth_keys,
			"auth_password":             auth_password,
			"auth_type":                 auth_type,
			"bfd_minimum_interval":      bfd_minimum_interval,
			"dead_interval":             dead_interval,
			"export_policy":             export_policy,
			"hello_interval":            hello_interval,
			"import_policy":             import_policy,
			"interface_type":            interface_type,
			"metric":                    metric,
			"no_readvertise_to_overlay": no_readvertise_to_overlay,
			"passive":                   passive,
		}
		data, e := NewOspfNetworksValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := OspfNetworksValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func ospfAreasSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.OspfArea) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {

		var include_loopback basetypes.BoolValue
		var networks basetypes.MapValue = types.MapNull(OspfNetworksValue{}.Type(ctx))
		var area_type basetypes.StringValue

		if d.IncludeLoopback != nil {
			include_loopback = types.BoolValue(*d.IncludeLoopback)
		}
		if d.Networks != nil {
			networks = ospfAreasNetworksSdkToTerraform(ctx, diags, d.Networks)
		}
		if d.Type != nil {
			area_type = types.StringValue(string(*d.Type))
		}

		data_map_attr_type := OspfAreasValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"include_loopback": include_loopback,
			"networks":         networks,
			"type":             area_type,
		}
		data, e := NewOspfAreasValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data

	}
	state_result_map_type := OspfAreasValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
