package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func pathPreferencePathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.GatewayPathPreferencesPath) basetypes.ListValue {
	tflog.Debug(ctx, "pathPreferencePathsSdkToTerraform")
	var data_list = []PathsValue{}

	for _, d := range l {
		var cost basetypes.Int64Value
		var disabled basetypes.BoolValue
		var gateway_ip basetypes.StringValue
		var internet_access basetypes.BoolValue = types.BoolValue(false)
		var name basetypes.StringValue
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var target_ips basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var type_path basetypes.StringValue
		var wan_name basetypes.StringValue

		if d.Cost != nil {
			cost = types.Int64Value(int64(*d.Cost))
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.GatewayIp != nil {
			gateway_ip = types.StringValue(*d.GatewayIp)
		}
		if d.InternetAccess != nil {
			internet_access = types.BoolValue(*d.InternetAccess)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}
		if d.TargetIps != nil {
			target_ips = mist_transform.ListOfStringSdkToTerraform(ctx, d.TargetIps)
		}
		if d.Type != nil {
			type_path = types.StringValue(string(*d.Type))
		}
		if d.WanName != nil {
			wan_name = types.StringValue(*d.WanName)
		}

		data_map_attr_type := PathsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"cost":            cost,
			"disabled":        disabled,
			"gateway_ip":      gateway_ip,
			"internet_access": internet_access,
			"name":            name,
			"networks":        networks,
			"target_ips":      target_ips,
			"type":            type_path,
			"wan_name":        wan_name,
		}
		data, e := NewPathsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := PathsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func pathPreferencesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayPathPreferences) basetypes.MapValue {
	tflog.Debug(ctx, "pathPreferencesSdkToTerraform")
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {
		var paths basetypes.ListValue = types.ListNull(PathsValue{}.Type(ctx))
		var strategy basetypes.StringValue = types.StringValue("ordered")

		if d.Paths != nil {
			paths = pathPreferencePathsSdkToTerraform(ctx, diags, d.Paths)
		}
		if d.Strategy != nil {
			strategy = types.StringValue(string(*d.Strategy))
		}

		data_map_attr_type := PathPreferencesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"paths":    paths,
			"strategy": strategy,
		}
		data, e := NewPathPreferencesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		state_value_map[k] = data
	}
	state_type := PathPreferencesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
