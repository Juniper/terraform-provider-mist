package resource_org_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func syntheticTestVlansTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ListValue) []models.SynthetictestProperties {
	var data_list []models.SynthetictestProperties
	for _, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VlansValue)
		data := models.SynthetictestProperties{}

		if !plan.CustomTestUrls.IsNull() && !plan.CustomTestUrls.IsUnknown() {
			data.CustomTestUrls = mist_transform.ListOfStringTerraformToSdk(ctx, plan.CustomTestUrls)
		}

		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}

		if !plan.VlanIds.IsNull() && !plan.VlanIds.IsUnknown() {
			var items []models.VlanIdWithVariable
			for _, item := range plan.VlanIds.Elements() {
				var item_interface interface{} = item
				i := item_interface.(basetypes.StringValue)
				v := models.VlanIdWithVariableContainer.FromString(i.ValueString())
				items = append(items, v)
			}
			data.VlanIds = items
		}

		data_list = append(data_list, data)
	}
	return data_list
}
func syntheticTestWanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SynthetictestConfigWanSpeedtest {
	data := models.SynthetictestConfigWanSpeedtest{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewWanSpeedtestValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.Enabled.ValueBoolPointer() != nil {
				data.Enabled = plan.Enabled.ValueBoolPointer()
			}

			if plan.TimeOdFay.ValueStringPointer() != nil {
				data.TimeOdFay = plan.TimeOdFay.ValueStringPointer()
			}
		}
	}
	return &data
}
func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) *models.SynthetictestConfig {
	data := models.SynthetictestConfig{}

	if d.Disabled.ValueBoolPointer() != nil {
		data.Disabled = d.Disabled.ValueBoolPointer()
	}

	if !d.Vlans.IsNull() && !d.Vlans.IsUnknown() {
		data.Vlans = syntheticTestVlansTerraformToSdk(ctx, diags, d.Vlans)
	}

	if !d.WanSpeedtest.IsNull() && !d.WanSpeedtest.IsUnknown() {
		data.WanSpeedtest = syntheticTestWanTerraformToSdk(ctx, diags, d.WanSpeedtest)
	}

	return &data
}
