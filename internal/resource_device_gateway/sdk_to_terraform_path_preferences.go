package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func pathPreferencePathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.GatewayPathPreferencesPath) basetypes.ListValue {
	var dataList []PathsValue

	for _, d := range l {
		var cost basetypes.Int64Value
		var disabled basetypes.BoolValue
		var gatewayIp basetypes.StringValue
		var internetAccess basetypes.BoolValue
		var name basetypes.StringValue
		var networks = types.ListNull(types.StringType)
		var targetIps = types.ListNull(types.StringType)
		var typePath basetypes.StringValue
		var wanName basetypes.StringValue

		if d.Cost != nil {
			cost = types.Int64Value(int64(*d.Cost))
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.GatewayIp != nil {
			gatewayIp = types.StringValue(*d.GatewayIp)
		}
		if d.InternetAccess != nil {
			internetAccess = types.BoolValue(*d.InternetAccess)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}
		if d.TargetIps != nil {
			targetIps = mistutils.ListOfStringSdkToTerraform(d.TargetIps)
		}

		typePath = types.StringValue(string(d.Type))

		if d.WanName != nil {
			wanName = types.StringValue(*d.WanName)
		}

		dataMapValue := map[string]attr.Value{
			"cost":            cost,
			"disabled":        disabled,
			"gateway_ip":      gatewayIp,
			"internet_access": internetAccess,
			"name":            name,
			"networks":        networks,
			"target_ips":      targetIps,
			"type":            typePath,
			"wan_name":        wanName,
		}
		data, e := NewPathsValue(PathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := PathsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func pathPreferencesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayPathPreferences) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var paths = types.ListNull(PathsValue{}.Type(ctx))
		var strategy = types.StringValue("ordered")

		if d.Paths != nil {
			paths = pathPreferencePathsSdkToTerraform(ctx, diags, d.Paths)
		}
		if d.Strategy != nil {
			strategy = types.StringValue(string(*d.Strategy))
		}

		dataMapValue := map[string]attr.Value{
			"paths":    paths,
			"strategy": strategy,
		}
		data, e := NewPathPreferencesValue(PathPreferencesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		stateValueMap[k] = data
	}
	stateType := PathPreferencesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
