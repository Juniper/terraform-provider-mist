package resource_org_wlan

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanDynamicVlan) DynamicVlanValue {

	var defaultVlanIds = mistutils.ListOfStringSdkToTerraformEmpty()
	var enabled basetypes.BoolValue
	var localVlanIds = mistutils.ListOfStringSdkToTerraformEmpty()
	var typeDynamicVlan basetypes.StringValue
	var vlans = types.MapNull(types.StringType)

	if d != nil && d.DefaultVlanIds != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		for _, item := range d.DefaultVlanIds {
			vlanId := mistutils.WlanDynamicVlanAsString(item)
			items = append(items, vlanId)
		}
		r, e := types.ListValue(itemsType, items)
		diags.Append(e...)
		defaultVlanIds = r
	}

	if d != nil && d.DefaultVlanIds == nil && d.DefaultVlanId != nil {
		var items []attr.Value
		var itemsType attr.Type = basetypes.StringType{}
		items = append(items, types.StringValue(d.DefaultVlanId.String()))
		r, e := types.ListValue(itemsType, items)
		diags.Append(e...)
		defaultVlanIds = r
	}

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.LocalVlanIds != nil {
		var list []attr.Value
		for _, v := range d.LocalVlanIds {
			list = append(list, types.StringValue(v.String()))
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		localVlanIds = r
	}
	if d != nil && d.Type != nil {
		typeDynamicVlan = types.StringValue(string(*d.Type))
	}
	if d != nil && d.Vlans != nil {
		vlansAttr := make(map[string]attr.Value)
		for k, v := range d.Vlans {
			vlansAttr[k] = types.StringValue(v)
		}
		vlans = types.MapValueMust(basetypes.StringType{}, vlansAttr)
	}

	dataMapValue := map[string]attr.Value{
		"default_vlan_ids": defaultVlanIds,
		"enabled":          enabled,
		"local_vlan_ids":   localVlanIds,
		"type":             typeDynamicVlan,
		"vlans":            vlans,
	}
	data, e := NewDynamicVlanValue(DynamicVlanValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
