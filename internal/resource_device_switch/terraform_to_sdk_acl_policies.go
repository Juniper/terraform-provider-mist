package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func aclPolicyActionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclPolicyAction {

	var data []models.AclPolicyAction
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(ActionsValue)
		data_item := models.AclPolicyAction{}
		if v_plan.Action.ValueStringPointer() != nil {
			data_item.Action = models.ToPointer(models.AllowDenyEnum(v_plan.Action.ValueString()))
		}
		if v_plan.DstTag.ValueStringPointer() != nil {
			data_item.DstTag = models.ToPointer(v_plan.DstTag.ValueString())
		}
		data = append(data, data_item)
	}
	return data
}

func aclPoliciesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.AclPolicy {

	var data []models.AclPolicy
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(AclPoliciesValue)
		data_item := models.AclPolicy{}
		if v_plan.Name.ValueStringPointer() != nil {
			data_item.Name = models.ToPointer(v_plan.Name.ValueString())
		}
		if !v_plan.Actions.IsNull() && !v_plan.Actions.IsUnknown() {
			actions := aclPolicyActionsTerraformToSdk(ctx, diags, v_plan.Actions)
			data_item.Actions = actions
		}
		if !v_plan.SrcTags.IsNull() && !v_plan.SrcTags.IsUnknown() {
			data_item.SrcTags = mist_transform.ListOfStringTerraformToSdk(ctx, v_plan.SrcTags)
		}

		data = append(data, data_item)
	}
	return data
}
