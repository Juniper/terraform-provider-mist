package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func syntheticTestWanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.SynthetictestConfigWanSpeedtest {
	data := models.SynthetictestConfigWanSpeedtest{}
	if !d.IsNull() && !d.IsUnknown() {
		vd, e := NewWanSpeedtestValue(WanSpeedtestValue{}.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			data.Enabled = vd.Enabled.ValueBoolPointer()
			data.TimeOfDay = vd.TimeOfDay.ValueStringPointer()
		}
	}
	return &data
}

func syntheticTestVlansTerraformToSdk(d basetypes.ListValue) []models.SynthetictestProperties {
	var dataList []models.SynthetictestProperties
	for _, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VlansValue)
		data := models.SynthetictestProperties{}

		if !plan.CustomTestUrls.IsNull() && !plan.CustomTestUrls.IsUnknown() {
			data.CustomTestUrls = mistutils.ListOfStringTerraformToSdk(plan.CustomTestUrls)
		}

		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}

		if !plan.VlanIds.IsNull() && !plan.VlanIds.IsUnknown() {
			var items []models.VlanIdWithVariable
			for _, item := range plan.VlanIds.Elements() {
				var itemInterface interface{} = item
				i := itemInterface.(basetypes.StringValue)
				v := models.VlanIdWithVariableContainer.FromString(i.ValueString())
				items = append(items, v)
			}
			data.VlanIds = items
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func syntheticTestTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SyntheticTestValue) *models.SynthetictestConfig {
	data := models.SynthetictestConfig{}

	if d.Disabled.ValueBoolPointer() != nil {
		data.Disabled = d.Disabled.ValueBoolPointer()
	}

	if !d.Vlans.IsNull() && !d.Vlans.IsUnknown() {
		data.Vlans = syntheticTestVlansTerraformToSdk(d.Vlans)
	}

	if !d.WanSpeedtest.IsNull() && !d.WanSpeedtest.IsUnknown() {
		data.WanSpeedtest = syntheticTestWanTerraformToSdk(ctx, diags, d.WanSpeedtest)
	}

	return &data
}
