package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authKeysTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []string {
	var items []string
	for _, v := range plan.Elements() {
		if v != nil {
			var v_inteface interface{} = v
			v_plan := v_inteface.(basetypes.StringValue)
			items = append(items, v_plan.ValueString())
		} else {
			var t string
			items = append(items, t)
		}
	}
	return items
}
func authPairwiseTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan basetypes.ListValue) []models.WlanAuthPairwiseItemEnum {
	var items []models.WlanAuthPairwiseItemEnum
	for _, v := range plan.Elements() {
		var v_inteface interface{} = v
		v_plan := v_inteface.(basetypes.StringValue)
		items = append(items, (models.WlanAuthPairwiseItemEnum)(v_plan.ValueString()))
	}
	return items
}
func authTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AuthValue) *models.WlanAuth {

	data := models.WlanAuth{}
	if plan.AuthType.ValueStringPointer() != nil {
		data.Type = models.WlanAuthTypeEnum(string(plan.AuthType.ValueString()))
	}
	if plan.AnticlogThreshold.ValueInt64Pointer() != nil {
		data.AnticlogThreshold = models.ToPointer(int(plan.AnticlogThreshold.ValueInt64()))
	}
	if plan.EapReauth.ValueBoolPointer() != nil {
		data.EapReauth = plan.EapReauth.ValueBoolPointer()
	}
	if plan.EnableMacAuth.ValueBoolPointer() != nil {
		data.EnableMacAuth = plan.EnableMacAuth.ValueBoolPointer()
	}
	if plan.KeyIdx.ValueInt64Pointer() != nil {
		data.KeyIdx = models.ToPointer(int(plan.KeyIdx.ValueInt64()))
	}
	if !plan.Keys.IsNull() && !plan.Keys.IsUnknown() {
		data.Keys = authKeysTerraformToSdk(ctx, diags, plan.Keys)
	}
	if plan.MultiPskOnly.ValueBoolPointer() != nil {
		data.MultiPskOnly = plan.MultiPskOnly.ValueBoolPointer()
	}
	if plan.Owe.ValueStringPointer() != nil {
		data.Owe = models.ToPointer(models.WlanAuthOweEnum(string(plan.Owe.ValueString())))
	}
	if !plan.Pairwise.IsNull() && !plan.Pairwise.IsUnknown() {
		data.Pairwise = authPairwiseTerraformToSdk(ctx, diags, plan.Pairwise)
	}
	if plan.Psk.ValueStringPointer() != nil {
		data.Psk = models.NewOptional(plan.Psk.ValueStringPointer())
	}
	if plan.WepAsSecondaryAuth.ValueBoolPointer() != nil {
		data.WepAsSecondaryAuth = plan.WepAsSecondaryAuth.ValueBoolPointer()
	}

	return &data
}
