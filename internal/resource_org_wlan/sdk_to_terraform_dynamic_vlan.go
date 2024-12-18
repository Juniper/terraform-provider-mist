package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicVlan) DynamicVlanValue {

	var default_vlan_ids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var enabled basetypes.BoolValue
	var local_vlan_ids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var type_dynamic_vlan basetypes.StringValue
	var vlans basetypes.MapValue = types.MapNull(types.StringType)

	if d != nil && d.DefaultVlanIds != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range d.DefaultVlanIds {
			items = append(items, types.StringValue(item.String()))
		}
		r, e := types.ListValue(items_type, items)
		diags.Append(e...)
		default_vlan_ids = r
	}

	if d != nil && d.DefaultVlanIds == nil && d.DefaultVlanId != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		items = append(items, types.StringValue(d.DefaultVlanId.String()))
		r, e := types.ListValue(items_type, items)
		diags.Append(e...)
		default_vlan_ids = r
	}

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.LocalVlanIds != nil {
		var list []attr.Value
		for _, v := range d.LocalVlanIds {
			list = append(list, types.StringValue(string(v.String())))
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		local_vlan_ids = r
	}
	if d != nil && d.Type != nil {
		type_dynamic_vlan = types.StringValue(string(*d.Type))
	}
	if d != nil && d.Vlans != nil {
		vlans_attr := make(map[string]attr.Value)
		for k, v := range d.Vlans {
			vlans_attr[k] = types.StringValue(string(v))
		}
		vlans = types.MapValueMust(basetypes.StringType{}, vlans_attr)
	}

	data_map_attr_type := DynamicVlanValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"default_vlan_ids": default_vlan_ids,
		"enabled":          enabled,
		"local_vlan_ids":   local_vlan_ids,
		"type":             type_dynamic_vlan,
		"vlans":            vlans,
	}
	data, e := NewDynamicVlanValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
