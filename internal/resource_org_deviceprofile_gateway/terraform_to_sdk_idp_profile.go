package resource_org_deviceprofile_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func idpProfileMatchingSeverityTerraformToSdk(ctx context.Context, list basetypes.ListValue) []models.IdpProfileMatchingSeverityValueEnum {
	var items []models.IdpProfileMatchingSeverityValueEnum
	for _, item := range list.Elements() {
		s := models.IdpProfileMatchingSeverityValueEnum(item.String())
		items = append(items, s)
	}
	return items
}

func idpProfileMatchingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.IdpProfileMatching {
	data := models.IdpProfileMatching{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewIpdProfileOverwriteMatchingValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {

			if !plan.AttackName.IsNull() && !plan.AttackName.IsUnknown() {
				data.AttackName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AttackName)
			}
			if !plan.DstSubnet.IsNull() && !plan.DstSubnet.IsUnknown() {
				data.DstSubnet = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DstSubnet)
			}
			if !plan.Severity.IsNull() && !plan.Severity.IsUnknown() {
				data.Severity = idpProfileMatchingSeverityTerraformToSdk(ctx, plan.Severity)
			}
		}
	}
	return &data
}

func idpProfileOverwritesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.IdpProfileOverwrite {
	var data_list []models.IdpProfileOverwrite
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(OverwritesValue)
		data := models.IdpProfileOverwrite{}

		if plan.Action.ValueStringPointer() != nil {
			data.Action = models.ToPointer(models.IdpProfileActionEnum(plan.Action.ValueString()))
		}
		if !plan.IpdProfileOverwriteMatching.IsNull() && !plan.IpdProfileOverwriteMatching.IsUnknown() {
			data.Matching = idpProfileMatchingTerraformToSdk(ctx, diags, plan.IpdProfileOverwriteMatching)
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func idpProfileTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.IdpProfile {
	data_map := make(map[string]models.IdpProfile)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(IdpProfilesValue)

		data := models.IdpProfile{}
		if plan.BaseProfile.ValueStringPointer() != nil {
			data.BaseProfile = models.ToPointer(models.IdpProfileBaseProfileEnum(plan.BaseProfile.ValueString()))
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if !plan.Overwrites.IsNull() && !plan.Overwrites.IsUnknown() {
			data.Overwrites = idpProfileOverwritesTerraformToSdk(ctx, diags, plan.Overwrites)
		}

		data_map[k] = data
	}
	return data_map
}
