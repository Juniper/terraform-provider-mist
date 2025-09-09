package resource_org_idpprofile

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func overwritesMatchingSeveritySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileMatchingSeverityValueEnum) basetypes.ListValue {
	listAttrTypes := types.StringType
	var listAttrValues []attr.Value
	for _, d := range l {
		listAttrValues = append(listAttrValues, types.StringValue(string(d)))
	}

	r, e := types.ListValueFrom(ctx, listAttrTypes, listAttrValues)
	diags.Append(e...)
	return r
}

func overwritesMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.IdpProfileMatching) basetypes.ObjectValue {
	var attackName = types.ListNull(types.StringType)
	var dstSubnet = types.ListNull(types.StringType)
	var severity = types.ListNull(types.StringType)

	if d.AttackName != nil {
		attackName = mistutils.ListOfStringSdkToTerraform(d.AttackName)
	}
	if d.DstSubnet != nil {
		dstSubnet = mistutils.ListOfStringSdkToTerraform(d.DstSubnet)
	}
	if d.Severity != nil {
		severity = overwritesMatchingSeveritySdkToTerraform(ctx, diags, d.Severity)
	}

	dataMapValue := map[string]attr.Value{
		"attack_name": attackName,
		"dst_subnet":  dstSubnet,
		"severity":    severity,
	}
	data, e := basetypes.NewObjectValue(MatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func overwritesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.IdpProfileOverwrite) basetypes.ListValue {
	var listAttrValues []OverwritesValue
	for _, d := range l {
		var action = types.StringNull()
		var matching = types.ObjectNull(MatchingValue{}.AttributeTypes(ctx))
		var name = types.StringNull()

		if d.Action != nil {
			action = types.StringValue(string(*d.Action))
		}
		if d.Matching != nil {
			matching = overwritesMatchingSdkToTerraform(ctx, diags, d.Matching)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		dataMapValue := map[string]attr.Value{
			"action":   action,
			"matching": matching,
			"name":     name,
		}
		data, e := NewOverwritesValue(OverwritesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		listAttrValues = append(listAttrValues, data)
	}

	r, e := types.ListValueFrom(ctx, OverwritesValue{}.Type(ctx), listAttrValues)
	diags.Append(e...)
	return r
}
