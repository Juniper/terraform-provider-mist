package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func actionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclPolicyAction) basetypes.ListValue {

	var list_attr_values []attr.Value
	for _, d := range l {
		var action basetypes.StringValue
		var dst_tag basetypes.StringValue

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		dst_tag = types.StringValue(d.DstTag)

		data_map_attr_type := ActionsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":  action,
			"dst_tag": dst_tag,
		}
		data, e := NewActionsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		list_attr_values = append(list_attr_values, data)
	}

	list_attr_types := ActionsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, list_attr_types, list_attr_values)
	diags.Append(e...)
	return r
}

func aclPoliciesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclPolicy) basetypes.ListValue {
	var data_list []attr.Value
	for _, d := range l {

		var actions basetypes.ListValue = types.ListNull(ActionsValue{}.Type(ctx))
		var name basetypes.StringValue
		var src_tags basetypes.ListValue = types.ListNull(types.StringType)

		if d.Actions != nil {
			actions = actionsSdkToTerraform(ctx, diags, d.Actions)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.SrcTags != nil {
			src_tags = mist_transform.ListOfStringSdkToTerraform(ctx, d.SrcTags)
		}

		data_map_attr_type := AclPoliciesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"actions":  actions,
			"name":     name,
			"src_tags": src_tags,
		}
		data, e := NewAclPoliciesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}

	state_list_type := AclPoliciesValue{}.Type(ctx)
	state_list, e := types.ListValueFrom(ctx, state_list_type, data_list)
	diags.Append(e...)

	return state_list
}
