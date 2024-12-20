package resource_site_networktemplate

import (
	"context"
	"strconv"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchMatchingRulesPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosPortConfig) basetypes.MapValue {
	map_item_value := make(map[string]attr.Value)
	map_item_type := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		var ae_disable_lacp basetypes.BoolValue
		var ae_idx basetypes.Int64Value
		var ae_lacp_slow basetypes.BoolValue
		var aggregated basetypes.BoolValue
		var critical basetypes.BoolValue
		var description basetypes.StringValue
		var disable_autoneg basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamic_usage basetypes.StringValue
		var esilag basetypes.BoolValue
		var mtu basetypes.Int64Value
		var no_local_overwrite basetypes.BoolValue
		var poe_disabled basetypes.BoolValue
		var speed basetypes.StringValue
		var usage basetypes.StringValue = types.StringValue(d.Usage)

		if d.AeDisableLacp != nil {
			ae_disable_lacp = types.BoolValue(*d.AeDisableLacp)
		}
		if d.AeIdx != nil {
			ae_idx = types.Int64Value(int64(*d.AeIdx))
		}
		if d.AeLacpSlow != nil {
			ae_lacp_slow = types.BoolValue(*d.AeLacpSlow)
		}
		if d.Aggregated != nil {
			aggregated = types.BoolValue(*d.Aggregated)
		}
		if d.Critical != nil {
			critical = types.BoolValue(*d.Critical)
		}
		if d.Description != nil {
			description = types.StringValue(*d.Description)
		}
		if d.DisableAutoneg != nil {
			disable_autoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicUsage.Value() != nil {
			dynamic_usage = types.StringValue(*d.DynamicUsage.Value())
		}
		if d.Esilag != nil {
			esilag = types.BoolValue(*d.Esilag)
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.NoLocalOverwrite != nil {
			no_local_overwrite = types.BoolValue(*d.NoLocalOverwrite)
		}
		if d.PoeDisabled != nil {
			poe_disabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		data_map_attr_type := PortConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ae_disable_lacp":    ae_disable_lacp,
			"ae_idx":             ae_idx,
			"ae_lacp_slow":       ae_lacp_slow,
			"aggregated":         aggregated,
			"critical":           critical,
			"description":        description,
			"disable_autoneg":    disable_autoneg,
			"duplex":             duplex,
			"dynamic_usage":      dynamic_usage,
			"esilag":             esilag,
			"mtu":                mtu,
			"no_local_overwrite": no_local_overwrite,
			"poe_disabled":       poe_disabled,
			"speed":              speed,
			"usage":              usage,
		}
		data, e := NewPortConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_item_value[k] = data
	}
	r, e := types.MapValueFrom(ctx, map_item_type, map_item_value)
	diags.Append(e...)
	return r
}
func switchMatchingRulesIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMatchingRuleIpConfig) basetypes.ObjectValue {
	var network basetypes.StringValue
	var ip_type basetypes.StringValue

	if d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d.Type != nil {
		ip_type = types.StringValue(string(*d.Type))
	}

	data_map_attr_type := IpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"network": network,
		"type":    ip_type,
	}
	data, e := NewIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func switchMatchingRulesOobIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMatchingRuleOobIpConfig) basetypes.ObjectValue {
	var oob_ip_type basetypes.StringValue
	var use_mgmt_vrf basetypes.BoolValue
	var use_mgmt_vrf_for_host_out basetypes.BoolValue

	if d.Type != nil {
		oob_ip_type = types.StringValue(string(*d.Type))
	}
	if d.UseMgmtVrf != nil {
		use_mgmt_vrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d.UseMgmtVrfForHostOut != nil {
		use_mgmt_vrf_for_host_out = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}

	data_map_attr_type := OobIpConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"type":                      oob_ip_type,
		"use_mgmt_vrf":              use_mgmt_vrf,
		"use_mgmt_vrf_for_host_out": use_mgmt_vrf_for_host_out,
	}
	data, e := NewOobIpConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func switchMatchingRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwitchMatchingRule) basetypes.ListValue {
	var data_list = []MatchingRulesValue{}

	for _, d := range l {

		var additional_config_cmds basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
		var match_model basetypes.StringValue = types.StringValue("")
		var match_name basetypes.StringValue = types.StringValue("")
		var match_name_offset basetypes.Int64Value = types.Int64Value(0)
		var match_role basetypes.StringValue
		var match_type basetypes.StringValue = types.StringValue("")
		var match_value basetypes.StringValue = types.StringValue("")
		var name basetypes.StringValue
		var port_config basetypes.MapValue = types.MapNull(PortConfigValue{}.Type(ctx))
		var port_mirroring basetypes.MapValue = types.MapNull(PortMirroringValue{}.Type(ctx))
		var ip_config basetypes.ObjectValue = types.ObjectNull(IpConfigValue{}.AttributeTypes(ctx))
		var oob_ip_config basetypes.ObjectValue = types.ObjectNull(OobIpConfigValue{}.AttributeTypes(ctx))

		for key, value := range d.AdditionalProperties {
			if strings.HasPrefix(key, "match_model") {
				match_model = types.StringValue(value)
				// backward compatibility
				match_type = types.StringValue(key)
				match_value = types.StringValue(value)
			} else if strings.HasPrefix(key, "match_name") {
				match_name = types.StringValue(value)
				if strings.Contains(key, "[") {
					offset_string := strings.Split(strings.Split(key, "[")[1], ":")[0]
					i, e := strconv.Atoi(offset_string)
					if e != nil {
						diags.AddWarning("Unable to extract the switch rule name offset", e.Error())
					} else {
						match_name_offset = types.Int64Value(int64(i))
					}
				}
				// backward compatibility
				match_type = types.StringValue(key)
				match_value = types.StringValue(value)
			} else if strings.HasPrefix(key, "match_role") {
				match_role = types.StringValue(value)
			} else if strings.HasPrefix(key, "match_") {
				match_type = types.StringValue(key)
				match_value = types.StringValue(value)
			}
		}

		if d.AdditionalConfigCmds != nil {
			additional_config_cmds = mist_transform.ListOfStringSdkToTerraform(ctx, d.AdditionalConfigCmds)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.PortConfig != nil && len(d.PortConfig) > 0 {
			port_config = switchMatchingRulesPortConfigSdkToTerraform(ctx, diags, d.PortConfig)
		}
		if d.PortMirroring != nil && len(d.PortMirroring) > 0 {
			port_mirroring = portMirroringSdkToTerraform(ctx, diags, d.PortMirroring)
		}
		if d.IpConfig != nil {
			ip_config = switchMatchingRulesIpConfigSdkToTerraform(ctx, diags, d.IpConfig)
		}
		if d.OobIpConfig != nil {
			oob_ip_config = switchMatchingRulesOobIpConfigSdkToTerraform(ctx, diags, d.OobIpConfig)
		}

		data_map_attr_type := MatchingRulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"additional_config_cmds": additional_config_cmds,
			"match_model":            match_model,
			"match_name":             match_name,
			"match_name_offset":      match_name_offset,
			"match_role":             match_role,
			"match_type":             match_type,
			"match_value":            match_value,
			"name":                   name,
			"port_config":            port_config,
			"port_mirroring":         port_mirroring,
			"ip_config":              ip_config,
			"oob_ip_config":          oob_ip_config,
		}
		data, e := NewMatchingRulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := MatchingRulesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	if e.HasError() {
		for _, f := range e.Errors() {
			tflog.Error(ctx, "switchMatchingRulesSdkToTerraform", map[string]interface{}{
				"summary": f.Summary(),
				"error":   f.Detail()})

		}
	}
	diags.Append(e...)
	return r
}

func switchMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMatching) SwitchMatchingValue {

	var enable basetypes.BoolValue
	var rules basetypes.ListValue = types.ListNull(MatchingRulesValue{}.Type(ctx))

	if d != nil && d.Enable != nil {
		enable = types.BoolValue(*d.Enable)
	}
	if d != nil && d.Rules != nil {
		rules = switchMatchingRulesSdkToTerraform(ctx, diags, d.Rules)
	}

	data_map_attr_type := SwitchMatchingValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enable": enable,
		"rules":  rules,
	}
	data, e := NewSwitchMatchingValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
