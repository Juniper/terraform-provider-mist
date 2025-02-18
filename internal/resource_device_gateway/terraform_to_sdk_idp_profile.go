package resource_device_gateway

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func idpProfileMatchingSeverityTerraformToSdk(list basetypes.ListValue) []models.IdpProfileMatchingSeverityValueEnum {
	var items []models.IdpProfileMatchingSeverityValueEnum
	for _, item := range list.Elements() {
		var iface interface{} = item
		val := iface.(basetypes.StringValue)
		s := models.IdpProfileMatchingSeverityValueEnum(val.ValueString())
		items = append(items, s)
	}
	return items
}

func idpProfileMatchingTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.IdpProfileMatching {
	data := models.IdpProfileMatching{}
	if d.IsNull() || d.IsUnknown() {
		return &data
	} else {
		plan := NewIpdProfileOverwriteMatchingValueMust(d.AttributeTypes(ctx), d.Attributes())

		if !plan.AttackName.IsNull() && !plan.AttackName.IsUnknown() {
			data.AttackName = misttransform.ListOfStringTerraformToSdk(plan.AttackName)
		}
		if !plan.DstSubnet.IsNull() && !plan.DstSubnet.IsUnknown() {
			data.DstSubnet = misttransform.ListOfStringTerraformToSdk(plan.DstSubnet)
		}
		if !plan.Severity.IsNull() && !plan.Severity.IsUnknown() {
			data.Severity = idpProfileMatchingSeverityTerraformToSdk(plan.Severity)
		}

		return &data
	}
}

func idpProfileOverwritesTerraformToSdk(ctx context.Context, d basetypes.ListValue) []models.IdpProfileOverwrite {
	var dataList []models.IdpProfileOverwrite
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(OverwritesValue)
		data := models.IdpProfileOverwrite{}

		if plan.Action.ValueStringPointer() != nil {
			data.Action = models.ToPointer(models.IdpProfileActionEnum(plan.Action.ValueString()))
		}
		if !plan.IpdProfileOverwriteMatching.IsNull() && !plan.IpdProfileOverwriteMatching.IsUnknown() {
			data.Matching = idpProfileMatchingTerraformToSdk(ctx, plan.IpdProfileOverwriteMatching)
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func idpProfileTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.IdpProfile {
	dataMap := make(map[string]models.IdpProfile)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(IdpProfilesValue)

		data := models.IdpProfile{}
		if plan.BaseProfile.ValueStringPointer() != nil {
			data.BaseProfile = models.ToPointer(models.IdpProfileBaseProfileEnum(plan.BaseProfile.ValueString()))
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if !plan.Overwrites.IsNull() && !plan.Overwrites.IsUnknown() {
			data.Overwrites = idpProfileOverwritesTerraformToSdk(ctx, plan.Overwrites)
		}

		dataMap[k] = data
	}
	return dataMap
}
