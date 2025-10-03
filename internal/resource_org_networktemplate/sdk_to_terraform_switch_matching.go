package resource_org_networktemplate

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

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func switchMatchingRulesPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.JunosPortConfig) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)
	mapItemType := PortConfigValue{}.Type(ctx)
	for k, d := range m {

		var aeDisableLacp basetypes.BoolValue
		var aeIdx basetypes.Int64Value
		var aeLacpSlow basetypes.BoolValue
		var aggregated basetypes.BoolValue
		var critical basetypes.BoolValue
		var description basetypes.StringValue
		var disableAutoneg basetypes.BoolValue
		var duplex basetypes.StringValue
		var dynamicUsage basetypes.StringValue
		var esilag basetypes.BoolValue
		var mtu basetypes.Int64Value
		var noLocalOverwrite basetypes.BoolValue
		var poeDisabled basetypes.BoolValue
		var portNetwork basetypes.StringValue
		var speed basetypes.StringValue
		var usage = types.StringValue(d.Usage)

		if d.AeDisableLacp != nil {
			aeDisableLacp = types.BoolValue(*d.AeDisableLacp)
		}
		if d.AeIdx != nil {
			aeIdx = types.Int64Value(int64(*d.AeIdx))
		}
		if d.AeLacpSlow != nil {
			aeLacpSlow = types.BoolValue(*d.AeLacpSlow)
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
			disableAutoneg = types.BoolValue(*d.DisableAutoneg)
		}
		if d.Duplex != nil {
			duplex = types.StringValue(string(*d.Duplex))
		}
		if d.DynamicUsage.Value() != nil {
			dynamicUsage = types.StringValue(*d.DynamicUsage.Value())
		}
		if d.Esilag != nil {
			esilag = types.BoolValue(*d.Esilag)
		}
		if d.Mtu != nil {
			mtu = types.Int64Value(int64(*d.Mtu))
		}
		if d.NoLocalOverwrite != nil {
			noLocalOverwrite = types.BoolValue(*d.NoLocalOverwrite)
		}
		if d.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*d.PoeDisabled)
		}
		if d.PortNetwork != nil {
			portNetwork = types.StringValue(*d.PortNetwork)
		}
		if d.Speed != nil {
			speed = types.StringValue(string(*d.Speed))
		}

		dataMapValue := map[string]attr.Value{
			"ae_disable_lacp":    aeDisableLacp,
			"ae_idx":             aeIdx,
			"ae_lacp_slow":       aeLacpSlow,
			"aggregated":         aggregated,
			"critical":           critical,
			"description":        description,
			"disable_autoneg":    disableAutoneg,
			"duplex":             duplex,
			"dynamic_usage":      dynamicUsage,
			"esilag":             esilag,
			"mtu":                mtu,
			"no_local_overwrite": noLocalOverwrite,
			"poe_disabled":       poeDisabled,
			"port_network":       portNetwork,
			"speed":              speed,
			"usage":              usage,
		}
		data, e := NewPortConfigValue(PortConfigValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}
	r, e := types.MapValueFrom(ctx, mapItemType, mapItemValue)
	diags.Append(e...)
	return r
}

func switchMatchingRulesIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMatchingRuleIpConfig) basetypes.ObjectValue {
	var network basetypes.StringValue
	var ipType basetypes.StringValue

	if d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d.Type != nil {
		ipType = types.StringValue(string(*d.Type))
	}

	dataMapValue := map[string]attr.Value{
		"network": network,
		"type":    ipType,
	}
	data, e := NewIpConfigValue(IpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}
func switchMatchingRulesOobIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchMatchingRuleOobIpConfig) basetypes.ObjectValue {
	var oobIpType basetypes.StringValue
	var useMgmtVrf basetypes.BoolValue
	var useMgmtVrfForHostOut basetypes.BoolValue

	if d.Type != nil {
		oobIpType = types.StringValue(string(*d.Type))
	}
	if d.UseMgmtVrf != nil {
		useMgmtVrf = types.BoolValue(*d.UseMgmtVrf)
	}
	if d.UseMgmtVrfForHostOut != nil {
		useMgmtVrfForHostOut = types.BoolValue(*d.UseMgmtVrfForHostOut)
	}

	dataMapValue := map[string]attr.Value{
		"type":                      oobIpType,
		"use_mgmt_vrf":              useMgmtVrf,
		"use_mgmt_vrf_for_host_out": useMgmtVrfForHostOut,
	}
	data, e := NewOobIpConfigValue(OobIpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o
}

func switchMatchingRulesStpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchStpConfig) basetypes.ObjectValue {

	var bridgePriority basetypes.StringValue

	if d.BridgePriority != nil {
		bridgePriority = types.StringValue(*d.BridgePriority)
	}

	dataMapValue := map[string]attr.Value{
		"bridge_priority": bridgePriority,
	}
	data, e := NewStpConfigValue(StpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	o, e := data.ToObjectValue(ctx)
	diags.Append(e...)
	return o

}

func switchMatchingRulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.SwitchMatchingRule) basetypes.ListValue {
	var dataList []MatchingRulesValue

	for _, d := range l {

		var additionalConfigCmds = types.ListNull(types.StringType)
		var matchModel basetypes.StringValue
		var matchName basetypes.StringValue
		var matchNameOffset = types.Int64Value(0)
		var matchRole basetypes.StringValue
		var name basetypes.StringValue
		var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
		var portMirroring = types.MapNull(PortMirroringValue{}.Type(ctx))
		var stpConfig = types.ObjectNull(StpConfigValue{}.AttributeTypes(ctx))
		var ipConfig = types.ObjectNull(IpConfigValue{}.AttributeTypes(ctx))
		var oobIpConfig = types.ObjectNull(OobIpConfigValue{}.AttributeTypes(ctx))

		for key, value := range d.AdditionalProperties {
			if strings.HasPrefix(key, "match_model") {
				matchModel = types.StringValue(value)
			} else if strings.HasPrefix(key, "match_name") {
				matchName = types.StringValue(value)
				if strings.Contains(key, "[") {
					offsetString := strings.Split(strings.Split(key, "[")[1], ":")[0]
					i, e := strconv.Atoi(offsetString)
					if e != nil {
						diags.AddWarning("Unable to extract the switch rule name offset", e.Error())
					} else {
						matchNameOffset = types.Int64Value(int64(i))
					}
				}
			} else if strings.HasPrefix(key, "match_role") {
				matchRole = types.StringValue(value)
			}
		}

		if len(d.AdditionalConfigCmds) > 0 {
			additionalConfigCmds = mistutils.ListOfStringSdkToTerraform(d.AdditionalConfigCmds)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if len(d.PortConfig) > 0 {
			portConfig = switchMatchingRulesPortConfigSdkToTerraform(ctx, diags, d.PortConfig)
		}
		if len(d.PortMirroring) > 0 {
			portMirroring = portMirroringSdkToTerraform(ctx, diags, d.PortMirroring)
		}
		if d.StpConfig != nil {
			stpConfig = switchMatchingRulesStpConfigSdkToTerraform(ctx, diags, d.StpConfig)
		}
		if d.IpConfig != nil {
			ipConfig = switchMatchingRulesIpConfigSdkToTerraform(ctx, diags, d.IpConfig)
		}
		if d.OobIpConfig != nil {
			oobIpConfig = switchMatchingRulesOobIpConfigSdkToTerraform(ctx, diags, d.OobIpConfig)
		}

		dataMapValue := map[string]attr.Value{
			"additional_config_cmds": additionalConfigCmds,
			"match_model":            matchModel,
			"match_name":             matchName,
			"match_name_offset":      matchNameOffset,
			"match_role":             matchRole,
			"name":                   name,
			"port_config":            portConfig,
			"port_mirroring":         portMirroring,
			"stp_config":             stpConfig,
			"ip_config":              ipConfig,
			"oob_ip_config":          oobIpConfig,
		}
		data, e := NewMatchingRulesValue(MatchingRulesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	datalistType := MatchingRulesValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
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
	var rules = types.ListNull(MatchingRulesValue{}.Type(ctx))

	if d != nil && d.Enable != nil {
		enable = types.BoolValue(*d.Enable)
	}
	if d != nil && d.Rules != nil {
		rules = switchMatchingRulesSdkToTerraform(ctx, diags, d.Rules)
	}

	dataMapValue := map[string]attr.Value{
		"enable": enable,
		"rules":  rules,
	}
	data, e := NewSwitchMatchingValue(SwitchMatchingValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
