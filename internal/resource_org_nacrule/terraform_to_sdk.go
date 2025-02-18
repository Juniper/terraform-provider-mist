package resource_org_nacrule

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgNacruleModel) (models.NacRule, diag.Diagnostics) {
	var diags diag.Diagnostics

	unset := make(map[string]interface{})
	data := models.NacRule{}

	if !plan.Action.IsNull() && !plan.Action.IsUnknown() {
		data.Action = *models.ToPointer(models.NacRuleActionEnum(plan.Action.ValueString()))
	} else {
		unset["-action"] = ""
	}
	if !plan.ApplyTags.IsNull() && !plan.ApplyTags.IsUnknown() {
		data.ApplyTags = misttransform.ListOfStringTerraformToSdk(plan.ApplyTags)
	} else {
		unset["-apply_tags"] = ""
	}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
	} else {
		unset["-enabled"] = ""
	}
	if !plan.Matching.IsNull() && !plan.Matching.IsUnknown() {
		data.Matching = matchingTerraformToSdk(plan.Matching)
	} else {
		unset["-matching"] = ""
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = plan.Name.ValueString()
	} else {
		unset["-name"] = ""
	}
	if !plan.NotMatching.IsNull() && !plan.NotMatching.IsUnknown() {
		data.NotMatching = notMatchingTerraformToSdk(plan.NotMatching)
	} else {
		unset["-not_matching"] = ""
	}
	if !plan.Order.IsNull() && !plan.Order.IsUnknown() {
		data.Order = models.ToPointer(int(plan.Order.ValueInt64()))
	} else {
		unset["-order"] = ""
	}

	return data, diags
}
