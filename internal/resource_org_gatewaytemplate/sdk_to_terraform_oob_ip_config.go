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
	var typeOob = types.StringValue("dhcp")
	var useMgmtVrf basetypes.BoolValue
	var useMgmtVrfForHostOut basetypes.BoolValue
	var vlanId basetypes.StringValue

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
		typeOob = types.StringValue(string(*d.Type))
	}
	if d != nil && d.UseMgmtVrf != nil {
		useMgmtVrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d != nil && d.UseMgmtVrfForHostOut != nil {
		useMgmtVrfForHostOut = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}
	if d != nil && d.VlanId != nil {
		vlanId = types.StringValue(*d.VlanId)
	}

	dataMapValue := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"type":                      typeOob,
		"use_mgmt_vrf":              useMgmtVrf,
		"use_mgmt_vrf_for_host_out": useMgmtVrfForHostOut,
		"vlan_id":                   vlanId,
	}
	data, e := basetypes.NewObjectValue(Node1Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func oobIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayOobIpConfig) OobIpConfigValue {
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var node1 = types.ObjectNull(Node1Value{}.AttributeTypes(ctx))
	var typeOob = types.StringValue("dhcp")
	var useMgmtVrf basetypes.BoolValue
	var useMgmtVrfForHostOut basetypes.BoolValue
	var vlanId basetypes.StringValue

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
		typeOob = types.StringValue(string(*d.Type))
	}
	if d != nil && d.UseMgmtVrf != nil {
		useMgmtVrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d != nil && d.UseMgmtVrfForHostOut != nil {
		useMgmtVrfForHostOut = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}
	if d != nil && d.VlanId != nil {
		vlanId = types.StringValue(*d.VlanId)
	}

	dataMapValue := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"node1":                     node1,
		"type":                      typeOob,
		"use_mgmt_vrf":              useMgmtVrf,
		"use_mgmt_vrf_for_host_out": useMgmtVrfForHostOut,
		"vlan_id":                   vlanId,
	}
	data, e := NewOobIpConfigValue(OobIpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
