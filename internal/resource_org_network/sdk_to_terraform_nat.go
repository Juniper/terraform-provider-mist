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
	state_value_map_attr_type := DestinationNatValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": types.StringValue(*v.InternalIp),
			"name":        types.StringValue(*v.Name),
			"port":        types.Int64Value(int64(*v.Port)),
		}
		n, e := NewDestinationNatValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := DestinationNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func staticNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkStaticNatProperty) basetypes.MapValue {
	state_value_map_attr_type := StaticNatValue{}.AttributeTypes(ctx)
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		state_value_map_attr_value := map[string]attr.Value{
			"internal_ip": types.StringValue(*v.InternalIp),
			"name":        types.StringValue(*v.Name),
			"wan_name":    types.StringValue(*v.WanName),
		}
		n, e := NewStaticNatValue(state_value_map_attr_type, state_value_map_attr_value)
		diags.Append(e...)
		state_value_map_value[k] = n
	}
	state_result_map_type := StaticNatValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkSourceNat) basetypes.ObjectValue {

	state_value_map_attr_type := SourceNatValue{}.AttributeTypes(ctx)

	var external_ip basetypes.StringValue
	if d.ExteralIp != nil {
		external_ip = types.StringValue(*d.ExteralIp)
	}

	state_value_map_attr_value := map[string]attr.Value{
		"exteral_ip": external_ip,
	}

	n, e := types.ObjectValue(state_value_map_attr_type, state_value_map_attr_value)
	diags.Append(e...)

	return n
}
