package resource_org_setting

import (
	"context"

	mistapi "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func syntheticTestVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SynthetictestProperties) basetypes.ListValue {

	var dataList []VlansValue
	for _, d := range l {

		var customTestUrls = misttransform.ListOfStringSdkToTerraformEmpty()
		var disabled basetypes.BoolValue
		var vlanIds = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.CustomTestUrls != nil && len(d.CustomTestUrls) > 0 {
			customTestUrls = misttransform.ListOfStringSdkToTerraform(d.CustomTestUrls)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.VlanIds != nil {
			var items []attr.Value
			for _, item := range d.VlanIds {
				items = append(items, mistapi.VlanAsString(item))
			}
			vlanIds, _ = types.ListValue(basetypes.StringType{}, items)
		}

		dataMapValue := map[string]attr.Value{
			"custom_test_urls": customTestUrls,
			"disabled":         disabled,
			"vlan_ids":         vlanIds,
		}
		data, e := NewVlansValue(VlansValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, VlansValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func syntheticTestWanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfigWanSpeedtest) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	var timeOfDay basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		timeOfDay = types.StringValue(*d.TimeOfDay)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":     enabled,
		"time_of_day": timeOfDay,
	}
	r, e := basetypes.NewObjectValue(WanSpeedtestValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return r
}

func syntheticTestSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfig) SyntheticTestValue {

	var disabled basetypes.BoolValue
	var vlans = types.ListNull(VlansValue{}.Type(ctx))
	var wanSpeedtest = types.ObjectNull(WanSpeedtestValue{}.AttributeTypes(ctx))

	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Vlans != nil {
		vlans = syntheticTestVlanSdkToTerraform(ctx, diags, d.Vlans)
	}
	if d.WanSpeedtest != nil {
		wanSpeedtest = syntheticTestWanSdkToTerraform(ctx, diags, d.WanSpeedtest)
	}

	dataMapValue := map[string]attr.Value{
		"disabled":      disabled,
		"vlans":         vlans,
		"wan_speedtest": wanSpeedtest,
	}
	data, e := NewSyntheticTestValue(SyntheticTestValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
