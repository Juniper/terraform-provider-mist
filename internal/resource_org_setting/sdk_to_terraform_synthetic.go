package resource_org_setting

import (
	"context"

	mist_api "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func syntheticTestVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SynthetictestProperties) basetypes.ListValue {

	var data_list = []VlansValue{}
	for _, d := range l {

		var custom_test_urls basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var disabled basetypes.BoolValue
		var vlan_ids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.CustomTestUrls != nil && len(d.CustomTestUrls) > 0 {
			custom_test_urls = mist_transform.ListOfStringSdkToTerraform(ctx, d.CustomTestUrls)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.VlanIds != nil {
			var items []attr.Value
			for _, item := range d.VlanIds {
				items = append(items, mist_api.VlanAsString(item))
			}
			vlan_ids, _ = types.ListValue(basetypes.StringType{}, items)
		}

		data_map_attr_type := VlansValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"custom_test_urls": custom_test_urls,
			"disabled":         disabled,
			"vlan_ids":         vlan_ids,
		}
		data, e := NewVlansValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}

	r, e := types.ListValueFrom(ctx, VlansValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func syntheticTestWanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfigWanSpeedtest) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	var time_of_day basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.TimeOfDay != nil {
		time_of_day = types.StringValue(*d.TimeOfDay)
	}

	data_map_attr_type := WanSpeedtestValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":     enabled,
		"time_of_day": time_of_day,
	}
	r, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)
	return r
}

func syntheticTestSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfig) SyntheticTestValue {

	var disabled basetypes.BoolValue
	var vlans basetypes.ListValue = types.ListNull(VlansValue{}.Type(ctx))
	var wan_speedtest basetypes.ObjectValue = types.ObjectNull(WanSpeedtestValue{}.AttributeTypes(ctx))

	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Vlans != nil {
		vlans = syntheticTestVlanSdkToTerraform(ctx, diags, d.Vlans)
	}
	if d.WanSpeedtest != nil {
		wan_speedtest = syntheticTestWanSdkToTerraform(ctx, diags, d.WanSpeedtest)
	}

	data_map_attr_type := SyntheticTestValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"disabled":      disabled,
		"vlans":         vlans,
		"wan_speedtest": wan_speedtest,
	}
	data, e := NewSyntheticTestValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
