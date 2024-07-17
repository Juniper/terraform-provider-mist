package resource_device_gateway

import (
	"context"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func idpProfileOverwriteMatchingSeveritiesSdkToTerraform(ctx context.Context, data []models.IdpProfileMatchingSeverityValueEnum) basetypes.ListValue {
	var items []attr.Value
	var items_type attr.Type = basetypes.StringType{}
	for _, item := range data {
		value := strings.ReplaceAll(string(item), "\"", "")
		items = append(items, types.StringValue(value))
	}
	list, _ := types.ListValue(items_type, items)
	return list
}

func idpProfileOverwriteMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpProfileMatching) basetypes.ObjectValue {

	var attack_name basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var dst_subnet basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var severity basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

	if d != nil && d.AttackName != nil {
		attack_name = mist_transform.ListOfStringSdkToTerraform(ctx, d.AttackName)
	}
	if d != nil && d.DstSubnet != nil {
		dst_subnet = mist_transform.ListOfStringSdkToTerraform(ctx, d.DstSubnet)
	}
	if d != nil && d.Severity != nil {
		severity = idpProfileOverwriteMatchingSeveritiesSdkToTerraform(ctx, d.Severity)
	}

	data_map_attr_type := IpdProfileOverwriteMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"attack_name": attack_name,
		"dst_subnet":  dst_subnet,
		"severity":    severity,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func idpProfileOverwritesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileOverwrite) basetypes.ListValue {
	var data_list = []OverwritesValue{}

	for _, d := range l {
		var action basetypes.StringValue = types.StringValue("alert")
		var matching basetypes.ObjectValue = types.ObjectNull(IpdProfileOverwriteMatchingValue{}.AttributeTypes(ctx))

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		if d.Matching != nil {
			matching = idpProfileOverwriteMatchingSdkToTerraform(ctx, diags, d.Matching)
		}

		data_map_attr_type := OverwritesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"action":   action,
			"matching": matching,
		}
		data, e := NewOverwritesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := OverwritesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func idpProfileSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.IdpProfile) basetypes.MapValue {

	state_value_map := make(map[string]attr.Value)
	for k, d := range m {

		var base_profile basetypes.StringValue
		var name basetypes.StringValue
		var overwrites basetypes.ListValue = types.ListNull(OverwritesValue{}.Type(ctx))

		if d.BaseProfile != nil {
			base_profile = types.StringValue(string(*d.BaseProfile))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Overwrites != nil {
			overwrites = idpProfileOverwritesSdkToTerraform(ctx, diags, d.Overwrites)
		}

		data_map_attr_type := IdpProfilesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"base_profile": base_profile,
			"name":         name,
			"overwrites":   overwrites,
		}
		data, e := NewIdpProfilesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := IdpProfilesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
