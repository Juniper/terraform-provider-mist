package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func oobIpConfigsNode1SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayOobIpConfigNode1) basetypes.ObjectValue {
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var type_oob basetypes.StringValue = types.StringValue("dhcp")
	var use_mgmt_vrf basetypes.BoolValue
	var use_mgmt_vrf_for_host_out basetypes.BoolValue
	var vlan_id basetypes.StringValue

	if d != nil && d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d != nil && d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d != nil && d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d != nil && d.Type != nil {
		type_oob = types.StringValue(string(*d.Type))
	}
	if d != nil && d.UseMgmtVrf != nil {
		use_mgmt_vrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d != nil && d.UseMgmtVrfForHostOut != nil {
		use_mgmt_vrf_for_host_out = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}
	if d != nil && d.VlanId != nil {
		vlan_id = types.StringValue(*d.VlanId)
	}

	data_map_attr_type := Node1Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"type":                      type_oob,
		"use_mgmt_vrf":              use_mgmt_vrf,
		"use_mgmt_vrf_for_host_out": use_mgmt_vrf_for_host_out,
		"vlan_id":                   vlan_id,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func oobIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayOobIpConfig) OobIpConfigValue {
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var node1 basetypes.ObjectValue = types.ObjectNull(Node1Value{}.AttributeTypes(ctx))
	var type_oob basetypes.StringValue = types.StringValue("dhcp")
	var use_mgmt_vrf basetypes.BoolValue
	var use_mgmt_vrf_for_host_out basetypes.BoolValue
	var vlan_id basetypes.StringValue

	if d != nil && d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d != nil && d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d != nil && d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d != nil && d.Node1 != nil {
		node1 = oobIpConfigsNode1SdkToTerraform(ctx, diags, d.Node1)
	}
	if d != nil && d.Type != nil {
		type_oob = types.StringValue(string(*d.Type))
	}
	if d != nil && d.UseMgmtVrf != nil {
		use_mgmt_vrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d != nil && d.UseMgmtVrfForHostOut != nil {
		use_mgmt_vrf_for_host_out = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}
	if d != nil && d.VlanId != nil {
		vlan_id = types.StringValue(*d.VlanId)
	}

	data_map_attr_type := OobIpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"node1":                     node1,
		"type":                      type_oob,
		"use_mgmt_vrf":              use_mgmt_vrf,
		"use_mgmt_vrf_for_host_out": use_mgmt_vrf_for_host_out,
		"vlan_id":                   vlan_id,
	}
	data, e := NewOobIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
