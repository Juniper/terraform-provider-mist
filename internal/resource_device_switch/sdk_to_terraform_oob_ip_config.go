package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func oobIpConfigsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchOobIpConfig) OobIpConfigValue {
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var network basetypes.StringValue
	var typeOob = types.StringValue("dhcp")
	var useMgmtVrf basetypes.BoolValue
	var useMgmtVrfForHostOut basetypes.BoolValue

	if d != nil && d.Gateway != nil {
		gateway = types.StringValue(*d.Gateway)
	}
	if d != nil && d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d != nil && d.Netmask != nil {
		netmask = types.StringValue(*d.Netmask)
	}
	if d != nil && d.Network != nil {
		network = types.StringValue(*d.Network)
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

	dataMapValue := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"network":                   network,
		"type":                      typeOob,
		"use_mgmt_vrf":              useMgmtVrf,
		"use_mgmt_vrf_for_host_out": useMgmtVrfForHostOut,
	}
	data, e := NewOobIpConfigValue(OobIpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
