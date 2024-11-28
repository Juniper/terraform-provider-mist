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
	var type_oob basetypes.StringValue = types.StringValue("dhcp")
	var use_mgmt_vrf basetypes.BoolValue
	var use_mgmt_vrf_for_host_out basetypes.BoolValue

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
		type_oob = types.StringValue(string(*d.Type))
	}
	if d != nil && d.UseMgmtVrf != nil {
		use_mgmt_vrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d != nil && d.UseMgmtVrfForHostOut != nil {
		use_mgmt_vrf_for_host_out = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}

	data_map_attr_type := OobIpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"gateway":                   gateway,
		"ip":                        ip,
		"netmask":                   netmask,
		"network":                   network,
		"type":                      type_oob,
		"use_mgmt_vrf":              use_mgmt_vrf,
		"use_mgmt_vrf_for_host_out": use_mgmt_vrf_for_host_out,
	}
	data, e := NewOobIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
