package resource_org_idpprofile

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func overwritesMatchingSeveritySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileMatchingSeverityValueEnum) basetypes.ListValue {
	list_attr_types := types.StringType
	var list_attr_values []attr.Value
	for _, d := range l {
		list_attr_values = append(list_attr_values, types.StringValue(string(d)))
	}

	r, e := types.ListValueFrom(ctx, list_attr_types, list_attr_values)
	diags.Append(e...)
	return r
}

func overwritesMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpProfileMatching) basetypes.ObjectValue {
	var attack_name basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var dst_subnet basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var severity basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

	if d.AttackName != nil {
		attack_name = mist_transform.ListOfStringSdkToTerraform(ctx, d.AttackName)
	}
	if d.DstSubnet != nil {
		dst_subnet = mist_transform.ListOfStringSdkToTerraform(ctx, d.DstSubnet)
	}
	if d.Severity != nil {
		severity = overwritesMatchingSeveritySdkToTerraform(ctx, diags, d.Severity)
	}

	data_map_attr_type := MatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"attack_name": attack_name,
		"dst_subnet":  dst_subnet,
		"severity":    severity,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func overwritesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileOverwrite) basetypes.ListValue {
	var list_attr_values []OverwritesValue
	for _, d := range l {
		var action basetypes.StringValue
		var matching basetypes.ObjectValue = types.ObjectNull(MatchingValue{}.AttributeTypes(ctx))
		var name basetypes.StringValue

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		if d.Matching != nil {
			matching = overwritesMatchingSdkToTerraform(ctx, diags, d.Matching)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		data_map_attr_type := OverwritesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":   action,
			"matching": matching,
			"name":     name,
		}
		data, e := NewOverwritesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		list_attr_values = append(list_attr_values, data)
	}

	r, e := types.ListValueFrom(ctx, OverwritesValue{}.Type(ctx), list_attr_values)
	diags.Append(e...)
	return r
}
