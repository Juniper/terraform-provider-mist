package resource_org_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func authKeysTerraformToSdk(plan basetypes.ListValue) []string {
	var items []string
	for _, v := range plan.Elements() {
		if v != nil {
			var vInteface interface{} = v
			vPlan := vInteface.(basetypes.StringValue)
			items = append(items, vPlan.ValueString())
		} else {
			var t string
			items = append(items, t)
		}
	}
	return items
}
func authPairwiseTerraformToSdk(plan basetypes.ListValue) []models.WlanAuthPairwiseItemEnum {
	var items []models.WlanAuthPairwiseItemEnum
	for _, v := range plan.Elements() {
		var vInteface interface{} = v
		vPlan := vInteface.(basetypes.StringValue)
		items = append(items, (models.WlanAuthPairwiseItemEnum)(vPlan.ValueString()))
	}
	return items
}
func authTerraformToSdk(plan AuthValue) *models.WlanAuth {

	data := models.WlanAuth{}
	if plan.AuthType.ValueStringPointer() != nil {
		data.Type = models.WlanAuthTypeEnum(plan.AuthType.ValueString())
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
		data.Keys = authKeysTerraformToSdk(plan.Keys)
	}
	if plan.MultiPskOnly.ValueBoolPointer() != nil {
		data.MultiPskOnly = plan.MultiPskOnly.ValueBoolPointer()
	}
	if plan.Owe.ValueStringPointer() != nil {
		data.Owe = models.ToPointer(models.WlanAuthOweEnum(plan.Owe.ValueString()))
	}
	if !plan.Pairwise.IsNull() && !plan.Pairwise.IsUnknown() {
		data.Pairwise = authPairwiseTerraformToSdk(plan.Pairwise)
	}
	if plan.Psk.ValueStringPointer() != nil {
		data.Psk = models.NewOptional(plan.Psk.ValueStringPointer())
	}
	if plan.WepAsSecondaryAuth.ValueBoolPointer() != nil {
		data.WepAsSecondaryAuth = plan.WepAsSecondaryAuth.ValueBoolPointer()
	}

	return &data
}
