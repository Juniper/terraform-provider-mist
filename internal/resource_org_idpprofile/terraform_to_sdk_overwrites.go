package resource_org_idpprofile

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func overwritesMatchingSeverityTerraformToSdk(l basetypes.ListValue) []models.IdpProfileMatchingSeverityValueEnum {
	var dataList []models.IdpProfileMatchingSeverityValueEnum
	for _, v := range l.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(basetypes.StringValue)
		severity := (models.IdpProfileMatchingSeverityValueEnum)(plan.ValueString())
		dataList = append(dataList, severity)
	}
	return dataList
}

func overwritesMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpProfileMatching {
	data := models.IdpProfileMatching{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewMatchingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.AttackName.IsNull() && !plan.AttackName.IsUnknown() {
				data.AttackName = misttransform.ListOfStringTerraformToSdk(plan.AttackName)
			}
			if !plan.DstSubnet.IsNull() && !plan.DstSubnet.IsUnknown() {
				data.DstSubnet = misttransform.ListOfStringTerraformToSdk(plan.DstSubnet)
			}
			if !plan.Severity.IsNull() && !plan.Severity.IsUnknown() {
				data.Severity = overwritesMatchingSeverityTerraformToSdk(plan.Severity)
			}
		}
	}
	return &data
}

func overwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []models.IdpProfileOverwrite {
	var dataList []models.IdpProfileOverwrite
	for _, v := range l.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(OverwritesValue)
		data := models.IdpProfileOverwrite{}

		if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
			data.Action = (*models.IdpProfileActionEnum)(plan.Action.ValueStringPointer())
		}
		if !plan.Matching.IsNull() && !plan.Matching.IsUnknown() {
			data.Matching = overwritesMatchingTerraformToSdk(ctx, diags, plan.Matching)
		}
		data.Name = plan.Name.ValueStringPointer()

		dataList = append(dataList, data)
	}
	return dataList
}
