package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func destinationNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkDestinationNatProperty) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var internal_ip basetypes.StringValue
		var name basetypes.StringValue
		var port basetypes.Int64Value

		if v.InternalIp != nil {
			internal_ip = types.StringValue(*v.InternalIp)
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Port != nil {
			port = types.Int64Value(int64(*v.Port))
		}

		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": internal_ip,
			"name":        name,
			"port":        port,
		}
		n, e := NewDestinationNatValue(DestinationNatValue{}.AttributeTypes(ctx), state_value_map_attr_value)
		diags.Append(e...)

		state_value_map_value[k] = n
	}
	state_result_map_type := DestinationNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func staticNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkStaticNatProperty) basetypes.MapValue {
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
		n, e := NewStaticNatValue(StaticNatValue{}.AttributeTypes(ctx), state_value_map_attr_value)
		diags.Append(e...)

		state_value_map_value[k] = n
	}
	state_result_map_type := StaticNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkSourceNat) basetypes.ObjectValue {
	var external_ip basetypes.StringValue

	if d != nil && d.ExternalIp != nil {
		external_ip = types.StringValue(*d.ExternalIp)
	}

	r_attr_type := SourceNatValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"external_ip": external_ip,
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
