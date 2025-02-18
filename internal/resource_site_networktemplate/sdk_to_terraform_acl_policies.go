package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func aclPolicyActionsToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclPolicyAction) basetypes.ListValue {
	var dataList []attr.Value
	for _, d := range l {

		var action basetypes.StringValue
		var dstTag basetypes.StringValue

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		dstTag = types.StringValue(d.DstTag)

		dataMapValue := map[string]attr.Value{
			"action":  action,
			"dst_tag": dstTag,
		}
		data, e := NewActionsValue(ActionsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	stateListType := ActionsValue{}.Type(ctx)
	stateList, e := types.ListValueFrom(ctx, stateListType, dataList)
	diags.Append(e...)
	return stateList
}

func aclPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclPolicy) basetypes.ListValue {
	var dataList []attr.Value
	for _, d := range l {

		var actions = types.ListNull(ActionsValue{}.Type(ctx))
		var name basetypes.StringValue
		var srcTags = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.Actions != nil {
			actions = aclPolicyActionsToTerraform(ctx, diags, d.Actions)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.SrcTags != nil {
			srcTags = misttransform.ListOfStringSdkToTerraform(d.SrcTags)
		}

		dataMapValue := map[string]attr.Value{
			"actions":  actions,
			"name":     name,
			"src_tags": srcTags,
		}
		data, e := NewAclPoliciesValue(AclPoliciesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	stateListType := AclPoliciesValue{}.Type(ctx)
	stateList, e := types.ListValueFrom(ctx, stateListType, dataList)
	diags.Append(e...)

	return stateList
}
