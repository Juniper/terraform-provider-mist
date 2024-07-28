package resource_site_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func bonjourServicesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanBonjourServiceProperties) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {

		var disable_local basetypes.BoolValue
		var radius_groups basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var scope basetypes.StringValue

		if d.DisableLocal != nil {
			disable_local = types.BoolValue(*d.DisableLocal)
		}
		if d.RadiusGroups != nil {
			radius_groups = mist_transform.ListOfStringSdkToTerraform(ctx, d.RadiusGroups)
		}
		if d.Scope != nil {
			scope = types.StringValue(string(*d.Scope))
		}

		data_map_attr_type := ServicesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disable_local": disable_local,
			"radius_groups": radius_groups,
			"scope":         scope,
		}
		data, e := NewServicesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	r, e := types.MapValueFrom(ctx, ServicesValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return r
}

func bonjourSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanBonjour) BonjourValue {
	var additional_vlan_ids basetypes.ListValue = types.ListNull(types.StringType)
	var enabled basetypes.BoolValue
	var services basetypes.MapValue = types.MapNull(ServicesValue{}.Type(ctx))

	if d != nil && d.AdditionalVlanIds != nil {
		var items []attr.Value
		for _, item := range d.AdditionalVlanIds {
			items = append(items, types.StringValue(item.String()))
		}
		list, _ := types.ListValue(basetypes.StringType{}, items)
		additional_vlan_ids = list
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.Services != nil && len(d.Services) > 0 {
		services = bonjourServicesSdkToTerraform(ctx, diags, d.Services)
	}

	data_map_attr_type := BonjourValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"additional_vlan_ids": additional_vlan_ids,
		"enabled":             enabled,
		"services":            services,
	}
	data, e := NewBonjourValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
