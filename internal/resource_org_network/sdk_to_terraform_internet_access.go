package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessDestinationNatProperty) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var internal_ip basetypes.StringValue
		var name basetypes.StringValue
		var port basetypes.StringValue
		var wan_name basetypes.StringValue

		if v.InternalIp != nil {
			internal_ip = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Port != nil {
			port = types.StringValue(*v.Port)
		}
		if v.WanName != nil {
			wan_name = types.StringValue(*v.WanName)
		}

		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": internal_ip,
			"name":        name,
			"port":        port,
			"wan_name":    wan_name,
		}
		n, e := NewInternetAccessDestinationNatValue(InternetAccessDestinationNatValue{}.AttributeTypes(ctx), state_value_map_attr_value)
		diags.Append(e...)

		state_value_map_value[k] = n
	}
	state_result_map, e := types.MapValueFrom(ctx, InternetAccessDestinationNatValue{}.Type(ctx), state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func staticNatInternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkInternetAccessStaticNatProperty) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var internal_ip basetypes.StringValue
		var name basetypes.StringValue
		var wan_name basetypes.StringValue

		if v.InternalIp != nil {
			internal_ip = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.WanName != nil {
			wan_name = types.StringValue(*v.WanName)
		}

		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": internal_ip,
			"name":        name,
			"wan_name":    wan_name,
		}
		n, e := NewInternetAccessStaticNatValue(InternetAccessStaticNatValue{}.AttributeTypes(ctx), state_value_map_attr_value)
		diags.Append(e...)

		state_value_map_value[k] = n
	}
	state_result_map, e := types.MapValueFrom(ctx, InternetAccessStaticNatValue{}.Type(ctx), state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func InternetAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternetAccess) InternetAccessValue {
	var create_simple_service_policy basetypes.BoolValue = types.BoolValue(false)
	var destination_nat basetypes.MapValue = types.MapNull(InternetAccessDestinationNatValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var restricted basetypes.BoolValue = types.BoolValue(false)
	var static_nac basetypes.MapValue = types.MapNull(InternetAccessStaticNatValue{}.Type(ctx))

	if d.CreateSimpleServicePolicy != nil {
		create_simple_service_policy = types.BoolValue(*d.CreateSimpleServicePolicy)
	}
	if d.DestinationNat != nil && len(d.DestinationNat) > 0 {
		destination_nat = destinationNatInternetAccessSdkToTerraform(ctx, diags, d.DestinationNat)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Restricted != nil {
		restricted = types.BoolValue(*d.Restricted)
	}
	if d.StaticNat != nil && len(d.StaticNat) > 0 {
		static_nac = staticNatInternetAccessSdkToTerraform(ctx, diags, d.StaticNat)
	}

	data_map_attr_type := InternetAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"create_simple_service_policy": create_simple_service_policy,
		"destination_nat":              destination_nat,
		"enabled":                      enabled,
		"restricted":                   restricted,
		"static_nat":                   static_nac,
	}
	data, e := NewInternetAccessValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
