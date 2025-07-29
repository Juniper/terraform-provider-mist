package resource_org_deviceprofile_gateway

import (
	"context"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func idpProfileOverwriteMatchingSeveritiesSdkToTerraform(data []models.IdpProfileMatchingSeverityValueEnum) basetypes.ListValue {
	var items []attr.Value
	var itemsType attr.Type = basetypes.StringType{}
	for _, item := range data {
		value := strings.ReplaceAll(string(item), "\"", "")
		items = append(items, types.StringValue(value))
	}
	list, _ := types.ListValue(itemsType, items)
	return list
}

func idpProfileOverwriteMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpProfileMatching) basetypes.ObjectValue {

	var attackName = types.ListNull(types.StringType)
	var dstSubnet = types.ListNull(types.StringType)
	var severity = types.ListNull(types.StringType)

	if d != nil && d.AttackName != nil {
		attackName = mistutils.ListOfStringSdkToTerraform(d.AttackName)
	}
	if d != nil && d.DstSubnet != nil {
		dstSubnet = mistutils.ListOfStringSdkToTerraform(d.DstSubnet)
	}
	if d != nil && d.Severity != nil {
		severity = idpProfileOverwriteMatchingSeveritiesSdkToTerraform(d.Severity)
	}

	dataMapValue := map[string]attr.Value{
		"attack_name": attackName,
		"dst_subnet":  dstSubnet,
		"severity":    severity,
	}
	data, e := basetypes.NewObjectValue(IpdProfileOverwriteMatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func idpProfileOverwritesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileOverwrite) basetypes.ListValue {
	var dataList []OverwritesValue

	for _, d := range l {
		var action basetypes.StringValue
		var matching = types.ObjectNull(IpdProfileOverwriteMatchingValue{}.AttributeTypes(ctx))

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		if d.Matching != nil {
			matching = idpProfileOverwriteMatchingSdkToTerraform(ctx, diags, d.Matching)
		}

		dataMapValue := map[string]attr.Value{
			"action":   action,
			"matching": matching,
		}
		data, e := NewOverwritesValue(OverwritesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := OverwritesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}

func idpProfileSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.IdpProfile) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var baseProfile basetypes.StringValue
		var name basetypes.StringValue
		var overwrites = types.ListNull(OverwritesValue{}.Type(ctx))

		if d.BaseProfile != nil {
			baseProfile = types.StringValue(string(*d.BaseProfile))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Overwrites != nil {
			overwrites = idpProfileOverwritesSdkToTerraform(ctx, diags, d.Overwrites)
		}

		dataMapValue := map[string]attr.Value{
			"base_profile": baseProfile,
			"name":         name,
			"overwrites":   overwrites,
		}
		data, e := NewIdpProfilesValue(IdpProfilesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := IdpProfilesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
