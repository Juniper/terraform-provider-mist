package resource_org_idpprofile

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func overwritesMatchingSeverityTerraformToSdk(ctx context.Context, l basetypes.ListValue) []models.IdpProfileMatchingSeverityValueEnum {
	var data_list []models.IdpProfileMatchingSeverityValueEnum
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(basetypes.StringValue)
		severity := (models.IdpProfileMatchingSeverityValueEnum)(plan.ValueString())
		data_list = append(data_list, severity)
	}
	return data_list
}

func overwritesMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpProfileMatching {
	data := models.IdpProfileMatching{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewMatchingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.AttackName.IsNull() && !plan.AttackName.IsUnknown() {
				data.AttackName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AttackName)
			}
			if !plan.DstSubnet.IsNull() && !plan.DstSubnet.IsUnknown() {
				data.DstSubnet = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstSubnet)
			}
			if !plan.Severity.IsNull() && !plan.Severity.IsUnknown() {
				data.Severity = overwritesMatchingSeverityTerraformToSdk(ctx, plan.Severity)
			}
		}
	}
	return &data
}

func overwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []models.IdpProfileOverwrite {
	var data_list []models.IdpProfileOverwrite
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OverwritesValue)
		data := models.IdpProfileOverwrite{}

		if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
			data.Action = (*models.IdpProfileActionEnum)(plan.Action.ValueStringPointer())
		}
		if !plan.Matching.IsNull() && !plan.Matching.IsUnknown() {
			data.Matching = overwritesMatchingTerraformToSdk(ctx, diags, plan.Matching)
		}
		data.Name = plan.Name.ValueStringPointer()

		data_list = append(data_list, data)
	}
	return data_list
}
