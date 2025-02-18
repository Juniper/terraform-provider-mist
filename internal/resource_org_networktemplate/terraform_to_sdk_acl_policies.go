package resource_org_networktemplate

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func aclPolicyActionsTerraformToSdk(d basetypes.ListValue) []models.AclPolicyAction {

	var data []models.AclPolicyAction
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(ActionsValue)
		dataItem := models.AclPolicyAction{}
		if vPlan.Action.ValueStringPointer() != nil {
			dataItem.Action = models.ToPointer(models.AllowDenyEnum(vPlan.Action.ValueString()))
		}
		if vPlan.DstTag.ValueStringPointer() != nil {
			dataItem.DstTag = vPlan.DstTag.ValueString()
		}
		data = append(data, dataItem)
	}
	return data
}

func aclPoliciesTerraformToSdk(d basetypes.ListValue) []models.AclPolicy {

	var data []models.AclPolicy
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		vPlan := vInterface.(AclPoliciesValue)
		dataItem := models.AclPolicy{}
		if vPlan.Name.ValueStringPointer() != nil {
			dataItem.Name = models.ToPointer(vPlan.Name.ValueString())
		}
		if !vPlan.Actions.IsNull() && !vPlan.Actions.IsUnknown() {
			actions := aclPolicyActionsTerraformToSdk(vPlan.Actions)
			dataItem.Actions = actions
		}
		if !vPlan.SrcTags.IsNull() && !vPlan.SrcTags.IsUnknown() {
			dataItem.SrcTags = misttransform.ListOfStringTerraformToSdk(vPlan.SrcTags)
		}

		data = append(data, dataItem)
	}
	return data
}
