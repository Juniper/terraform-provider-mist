package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func synthteticTestVlansSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SynthetictestProperties) basetypes.ListValue {
	tflog.Debug(ctx, "synthteticTestVlansSdkToTerraform")
	var data_list = []VlansValue{}
	for _, d := range l {
		var custom_test_urls basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var disabled basetypes.BoolValue
		var vlan_ids basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)

		if d.CustomTestUrls != nil {
			custom_test_urls = mist_transform.ListOfStringSdkToTerraform(ctx, d.CustomTestUrls)
		}
		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.VlanIds != nil {
			vlan_ids = mist_transform.ListOfIntSdkToTerraform(ctx, d.VlanIds)
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
	data_list_type := VlansValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}

func synthteticTestSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SynthetictestConfig) SyntheticTestValue {
	tflog.Debug(ctx, "synthteticTestSdkToTerraform")

	var disabled basetypes.BoolValue
	var vlans basetypes.ListValue = types.ListNull(VlansValue{}.Type(ctx))

	if d != nil && d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d != nil && d.Vlans != nil {
		vlans = synthteticTestVlansSdkToTerraform(ctx, diags, d.Vlans)
	}

	data_map_attr_type := SyntheticTestValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"disabled": disabled,
		"vlans":    vlans,
	}
	data, e := NewSyntheticTestValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
