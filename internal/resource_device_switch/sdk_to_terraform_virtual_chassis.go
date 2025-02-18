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

	var dataList []MembersValue
	for _, d := range l {
		var mac basetypes.StringValue
		var memberId basetypes.Int64Value
		var vcRole basetypes.StringValue

		mac = types.StringValue(d.Mac)
		memberId = types.Int64Value(int64(d.MemberId))
		vcRole = types.StringValue(string(d.VcRole))

		dataMapValue := map[string]attr.Value{
			"mac":       mac,
			"member_id": memberId,
			"vc_role":   vcRole,
		}
		data, e := NewMembersValue(MembersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, MembersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func virtualChassisSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchVirtualChassis) VirtualChassisValue {
	var members = types.ListNull(MembersValue{}.Type(ctx))
	var preprovisioned basetypes.BoolValue

	if d.Members != nil {
		members = virtualChassisMembersSdkToTerraform(ctx, diags, d.Members)
	}
	if d.Preprovisioned != nil {
		preprovisioned = types.BoolValue(*d.Preprovisioned)
	}

	dataMapValue := map[string]attr.Value{
		"members":        members,
		"preprovisioned": preprovisioned,
	}
	data, e := NewVirtualChassisValue(VirtualChassisValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
