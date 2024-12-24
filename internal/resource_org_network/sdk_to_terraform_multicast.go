package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func groupMutlicastSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.NetworkMulticastGroup) basetypes.MapValue {
	state_value_map_value := make(map[string]attr.Value)
	for k, v := range d {
		var rp_ip basetypes.StringValue

		if v.RpIp != nil {
			rp_ip = types.StringValue(*v.RpIp)
		}

		data_map_attr_type := GroupsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"rp_ip": rp_ip,
		}
		n, e := NewGroupsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = n
	}
	state_result_map, e := types.MapValueFrom(ctx, GroupsValue{}.Type(ctx), state_value_map_value)
	diags.Append(e...)
	return state_result_map
}

func MutlicastSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkMulticast) MulticastValue {
	var disable_igmp basetypes.BoolValue
	var enabled basetypes.BoolValue
	var groups basetypes.MapValue

	if d.DisableIgmp != nil {
		disable_igmp = types.BoolValue(*d.DisableIgmp)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Groups != nil {
		groups = groupMutlicastSdkToTerraform(ctx, diags, d.Groups)
	}

	data_map_attr_type := MulticastValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"disable_igmp": disable_igmp,
		"enabled":      enabled,
		"groups":       groups,
	}
	data, e := NewMulticastValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
