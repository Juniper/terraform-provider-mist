package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func virtualChassisMembersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwitchVirtualChassisMember) basetypes.ListValue {

	var data_list = []MembersValue{}
	for _, d := range l {
		var mac basetypes.StringValue
		var member_id basetypes.Int64Value
		var vc_role basetypes.StringValue

		mac = types.StringValue(d.Mac)
		member_id = types.Int64Value(int64(d.MemberId))
		vc_role = types.StringValue(string(d.VcRole))

		data_map_attr_type := MembersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"mac":       mac,
			"member_id": member_id,
			"vc_role":   vc_role,
		}
		data, e := NewMembersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, MembersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func virtualChassisSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchVirtualChassis) VirtualChassisValue {
	var members basetypes.ListValue = types.ListNull(MembersValue{}.Type(ctx))
	var preprovisioned basetypes.BoolValue

	if d.Members != nil {
		members = virtualChassisMembersSdkToTerraform(ctx, diags, d.Members)
	}
	if d.Preprovisioned != nil {
		preprovisioned = types.BoolValue(*d.Preprovisioned)
	}

	data_map_attr_type := VirtualChassisValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"members":        members,
		"preprovisioned": preprovisioned,
	}
	data, e := NewVirtualChassisValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
