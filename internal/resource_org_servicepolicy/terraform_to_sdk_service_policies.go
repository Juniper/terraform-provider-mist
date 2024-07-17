package resource_org_servicepolicy

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appqoeTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d AppqoeValue) *models.ServicePolicyAppqoe {
	data := models.ServicePolicyAppqoe{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		if d.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(d.Enabled.ValueBool())
		}
		return &data
	}
}

func ewfRuleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.ServicePolicyEwfRule {
	var data_list []models.ServicePolicyEwfRule
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(EwfValue)
		data := models.ServicePolicyEwfRule{}
		if plan.AlertOnly.ValueBoolPointer() != nil {
			data.AlertOnly = models.ToPointer(plan.AlertOnly.ValueBool())
		}
		if plan.BlockMessage.ValueStringPointer() != nil {
			data.BlockMessage = models.ToPointer(plan.BlockMessage.ValueString())
		}
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
		if plan.Profile.ValueStringPointer() != nil {
			data.Profile = models.ToPointer(models.ServicePolicyEwfRuleProfileEnum(plan.Profile.ValueString()))
		}

		data_list = append(data_list, data)
	}
	return data_list
}
