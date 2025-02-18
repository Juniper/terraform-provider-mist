package resource_device_switch

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func virtualChassisMemberTerraformToSdk(d basetypes.ListValue) []models.SwitchVirtualChassisMember {
	var dataList []models.SwitchVirtualChassisMember
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(MembersValue)
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

		dataList = append(dataList, data)
	}
	return dataList
}

func virtualChassisTerraformToSdk(d VirtualChassisValue) *models.SwitchVirtualChassis {

	data := models.SwitchVirtualChassis{}

	if !d.Members.IsNull() && !d.Members.IsUnknown() {
		data.Members = virtualChassisMemberTerraformToSdk(d.Members)
	}
	if d.Preprovisioned.ValueBoolPointer() != nil {
		data.Preprovisioned = d.Preprovisioned.ValueBoolPointer()
	}

	return &data
}
