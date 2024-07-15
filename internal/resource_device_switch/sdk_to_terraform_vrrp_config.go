package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func vrrpGroupsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VrrpConfigGroup) basetypes.MapValue {
	data_map_value := make(map[string]attr.Value)
	for k, d := range m {

		var priority basetypes.Int64Value

		if d.Priority != nil {
			priority = types.Int64Value(int64(*d.Priority))
		}

		data_map_attr_type := GroupsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"priority": priority,
		}
		data, e := NewGroupsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_map_value[k] = data
	}
	state_type := GroupsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}

func vrrpConfigInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VrrpConfig) VrrpConfigValue {
	var enabled basetypes.BoolValue
	var groups basetypes.MapValue = types.MapNull(GroupsValue{}.Type(ctx))

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Groups != nil {
		groups = vrrpGroupsSdkToTerraform(ctx, diags, d.Groups)
	}

	data_map_attr_type := VrrpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
		"groups":  groups,
	}
	data, e := NewVrrpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
