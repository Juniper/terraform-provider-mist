package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func virtualChassisMemberTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SwitchVirtualChassisMember {
	var data_list []models.SwitchVirtualChassisMember
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(MembersValue)
		data := models.SwitchVirtualChassisMember{}

		if plan.Mac.ValueStringPointer() != nil {
			data.Mac = plan.Mac.ValueString()
		}
		if plan.MemberId.ValueInt64Pointer() != nil {
			data.MemberId = int(plan.MemberId.ValueInt64())
		}
		if plan.VcRole.ValueStringPointer() != nil {
			data.VcRole = models.SwitchVirtualChassisMemberVcRoleEnum(plan.VcRole.ValueString())
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func virtualChassisTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VirtualChassisValue) *models.SwitchVirtualChassis {

	data := models.SwitchVirtualChassis{}

	if !d.Members.IsNull() && !d.Members.IsUnknown() {
		data.Members = virtualChassisMemberTerraformToSdk(ctx, diags, d.Members)
	}
	if d.Preprovisioned.ValueBoolPointer() != nil {
		data.Preprovisioned = d.Preprovisioned.ValueBoolPointer()
	}

	return &data
}
