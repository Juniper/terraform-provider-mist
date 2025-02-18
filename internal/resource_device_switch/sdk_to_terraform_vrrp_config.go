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
	dataMapValue := make(map[string]attr.Value)
	for k, d := range m {

		var priority basetypes.Int64Value

		if d.Priority != nil {
			priority = types.Int64Value(int64(*d.Priority))
		}

		itemMapAttrType := GroupsValue{}.AttributeTypes(ctx)
		itemMapValue := map[string]attr.Value{
			"priority": priority,
		}
		data, e := NewGroupsValue(itemMapAttrType, itemMapValue)
		diags.Append(e...)

		dataMapValue[k] = data
	}
	stateType := GroupsValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, dataMapValue)
	diags.Append(e...)
	return stateResult
}

func vrrpConfigInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VrrpConfig) VrrpConfigValue {
	var enabled basetypes.BoolValue
	var groups = types.MapNull(GroupsValue{}.Type(ctx))

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Groups != nil && len(d.Groups) > 0 {
		groups = vrrpGroupsSdkToTerraform(ctx, diags, d.Groups)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"groups":  groups,
	}
	data, e := NewVrrpConfigValue(VrrpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
