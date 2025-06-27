package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func wanExtraRoutesPortConfigIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WanExtraRoutes) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var via basetypes.StringValue
		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}
		dataMapValue := map[string]attr.Value{
			"via": via,
		}
		data, e := NewWanExtraRoutesValue(WanExtraRoutesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, WanExtraRoutesValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}
func wanProbeOverridePortConfigIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayWanProbeOverride) basetypes.ObjectValue {
	var ips = types.ListNull(types.StringType)
	var probeProfile basetypes.StringValue

	if g != nil && g.Ips != nil {
		ips = mistutils.ListOfStringSdkToTerraform(g.Ips)
	}
	if g != nil && g.ProbeProfile != nil {
		probeProfile = types.StringValue(string(*g.ProbeProfile))
	}

	rAttrValue := map[string]attr.Value{
		"ips":           ips,
		"probe_profile": probeProfile,
	}

	r, e := basetypes.NewObjectValue(WanProbeOverrideValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}
func portConfigIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayPortConfigIpConfig) basetypes.ObjectValue {
	var dns = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var network basetypes.StringValue
	var poserPassword basetypes.StringValue
	var pppoeAuth = types.StringValue("none")
	var pppoeUsername basetypes.StringValue
	var ipConfigType = types.StringValue("dhcp")

	if g != nil && g.Dns != nil {
		dns = mistutils.ListOfStringSdkToTerraform(g.Dns)
	}
	if g != nil && g.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(g.DnsSuffix)
	}
	if g != nil && g.Gateway != nil {
		gateway = types.StringValue(*g.Gateway)
	}
	if g != nil && g.Ip != nil {
		ip = types.StringValue(*g.Ip)
	}
	if g != nil && g.Netmask != nil {
		netmask = types.StringValue(*g.Netmask)
	}
	if g != nil && g.Network != nil {
		network = types.StringValue(*g.Network)
	}
	if g != nil && g.PoserPassword != nil {
		poserPassword = types.StringValue(*g.PoserPassword)
	}
	if g != nil && g.PppoeAuth != nil {
		pppoeAuth = types.StringValue(string(*g.PppoeAuth))
	}
	if g != nil && g.PppoeUsername != nil {
		pppoeUsername = types.StringValue(*g.PppoeUsername)
	}
	if g != nil && g.Type != nil {
		ipConfigType = types.StringValue(string(*g.Type))
	}

	rAttrValue := map[string]attr.Value{
		"dns":            dns,
		"dns_suffix":     dnsSuffix,
		"gateway":        gateway,
		"ip":             ip,
		"netmask":        netmask,
		"network":        network,
		"poser_password": poserPassword,
		"pppoe_username": pppoeUsername,
		"pppoe_auth":     pppoeAuth,
		"type":           ipConfigType,
	}

	r, e := basetypes.NewObjectValue(PortIpConfigValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func portConfigTrafficShapingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayTrafficShaping) basetypes.ObjectValue {
	var classPercentages = mistutils.ListOfIntSdkToTerraformEmpty()
	var enabled = types.BoolValue(false)
	var maxTxKbps basetypes.Int64Value

	if g != nil && g.ClassPercentages != nil {
		classPercentages = mistutils.ListOfIntSdkToTerraform(g.ClassPercentages)
	}
	if g != nil && g.Enabled != nil {
		enabled = types.BoolValue(*g.Enabled)
	}
	if g != nil && g.MaxTxKbps != nil {
		maxTxKbps = types.Int64Value(int64(*g.MaxTxKbps))
	}

	rAttrValue := map[string]attr.Value{
		"class_percentages": classPercentages,
		"enabled":           enabled,
		"max_tx_kbps":       maxTxKbps,
	}

	r, e := basetypes.NewObjectValue(TrafficShapingValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func portConfigVpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g map[string]models.GatewayPortVpnPath) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, v := range g {

		var bfdProfile = types.StringValue("broadband")
		var bfdUseTunnelMode = types.BoolValue(false)
		var preference basetypes.Int64Value
		var role = types.StringValue("spoke")
		var trafficShaping = types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))

		if v.BfdProfile != nil {
			bfdProfile = types.StringValue(string(*v.BfdProfile))
		}
		if v.BfdUseTunnelMode != nil {
			bfdUseTunnelMode = types.BoolValue(*v.BfdUseTunnelMode)
		}
		if v.Preference != nil {
			preference = types.Int64Value(int64(*v.Preference))
		}
		if v.TrafficShaping != nil {
			trafficShaping = portConfigTrafficShapingSdkToTerraform(ctx, diags, v.TrafficShaping)
		}
		if v.Role != nil {
			role = types.StringValue(string(*v.Role))
		}

		var portUsageState = map[string]attr.Value{
			"bfd_profile":         bfdProfile,
			"bfd_use_tunnel_mode": bfdUseTunnelMode,
			"preference":          preference,
			"role":                role,
			"traffic_shaping":     trafficShaping,
		}
		portUsageObject, e := NewVpnPathsValue(VpnPathsValue{}.AttributeTypes(ctx), portUsageState)
		diags.Append(e...)
		stateValueMap[k] = portUsageObject
	}

	stateResult, e := types.MapValueFrom(ctx, VpnPathsValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}

func portConfigWanSourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayPortWanSourceNat) basetypes.ObjectValue {

	var disabled = types.BoolValue(false)
	var natPool basetypes.StringValue

	if g != nil && g.Disabled != nil {
		disabled = types.BoolValue(*g.Disabled)
	}
	if g != nil && g.NatPool != nil {
		natPool = types.StringValue(*g.NatPool)
	}

	rAttrValue := map[string]attr.Value{
		"disabled": disabled,
		"nat_pool": natPool,
	}

	r, e := basetypes.NewObjectValue(WanSourceNatValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)

	return r
}

func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.GatewayPortConfig) basetypes.MapValue {

	stateValueMap := make(map[string]attr.Value)
	for k, v := range d {
		var aeDisableLacp = types.BoolValue(false)
		var aeIdx basetypes.StringValue
		var aeLacpForceUp = types.BoolValue(false)
		var aggregated = types.BoolValue(false)
		var critical = types.BoolValue(false)
		var description basetypes.StringValue
		var disableAutoneg = types.BoolValue(false)
		var disabled = types.BoolValue(false)
		var dslType = types.StringValue("vdsl")
		var dslVci = types.Int64Value(35)
		var dslVpi = types.Int64Value(0)
		var duplex = types.StringValue("auto")
		var ipConfig = types.ObjectNull(PortIpConfigValue{}.AttributeTypes(ctx))
		var lteApn basetypes.StringValue
		var lteAuth = types.StringValue("none")
		var lteBackup basetypes.BoolValue
		var ltePassword basetypes.StringValue
		var lteUsername basetypes.StringValue
		var mtu basetypes.Int64Value
		var name basetypes.StringValue
		var networks = types.ListNull(types.StringType)
		var outerVlanId basetypes.Int64Value
		var poeDisabled = types.BoolValue(false)
		var portNetwork basetypes.StringValue
		var preserveDscp = types.BoolValue(true)
		var redundant basetypes.BoolValue
		var redundantGroup basetypes.Int64Value
		var rethIdx basetypes.StringValue
		var rethNode basetypes.StringValue
		var rethNodes = types.ListNull(types.StringType)
		var speed = types.StringValue("auto")
		var ssrNoVirtualMac = types.BoolValue(false)
		var svrPortRange = types.StringValue("none")
		var trafficShaping = types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))
		var usage = types.StringValue(string(v.Usage))
		var vlanId basetypes.StringValue
		var vpnPaths = types.MapNull(VpnPathsValue{}.Type(ctx))
		var wanArpPolicer = types.StringValue("default")
		var wanDisableSpeedtest types.Bool
		var wanExtIp basetypes.StringValue
		var wanExtraRoutes = types.MapValueMust(WanExtraRoutesValue{}.Type(ctx), map[string]attr.Value{})
		var wanNetworks = mistutils.ListOfStringSdkToTerraform(v.WanNetworks)
		var wanProbeOverride = types.ObjectNull(WanProbeOverrideValue{}.AttributeTypes(ctx))
		var wanSourceNat = types.ObjectNull(WanSourceNatValue{}.AttributeTypes(ctx))
		var wanType = types.StringValue("broadband")

		if v.AeDisableLacp != nil {
			aeDisableLacp = types.BoolValue(*v.AeDisableLacp)
		}
		if v.AeIdx.Value() != nil {
			aeIdx = types.StringValue(*v.AeIdx.Value())
		}
		if v.AeLacpForceUp != nil {
			aeLacpForceUp = types.BoolValue(*v.AeLacpForceUp)
		}
		if v.Aggregated != nil {
			aggregated = types.BoolValue(*v.Aggregated)
		}
		if v.Critical != nil {
			critical = types.BoolValue(*v.Critical)
		}
		if v.Description != nil {
			description = types.StringValue(*v.Description)
		}
		if v.DisableAutoneg != nil {
			disableAutoneg = types.BoolValue(*v.DisableAutoneg)
		}
		if v.Disabled != nil {
			disabled = types.BoolValue(*v.Disabled)
		}
		if v.DslType != nil {
			dslType = types.StringValue(string(*v.DslType))
		}
		if v.DslVci != nil {
			dslVci = types.Int64Value(int64(*v.DslVci))
		}
		if v.DslVpi != nil {
			dslVpi = types.Int64Value(int64(*v.DslVpi))
		}
		if v.Duplex != nil {
			duplex = types.StringValue(string(*v.Duplex))
		}
		if v.IpConfig != nil {
			ipConfig = portConfigIpConfigSdkToTerraform(ctx, diags, v.IpConfig)
		}
		if v.LteApn != nil {
			lteApn = types.StringValue(*v.LteApn)
		}
		if v.LteAuth != nil {
			lteAuth = types.StringValue(string(*v.LteAuth))
		}
		if v.LteBackup != nil {
			lteBackup = types.BoolValue(*v.LteBackup)
		}
		if v.LtePassword != nil {
			ltePassword = types.StringValue(*v.LtePassword)
		}
		if v.LteUsername != nil {
			lteUsername = types.StringValue(*v.LteUsername)
		}
		if v.Mtu != nil {
			mtu = types.Int64Value(int64(*v.Mtu))
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(v.Networks)
		}
		if v.OuterVlanId != nil {
			outerVlanId = types.Int64Value(int64(*v.OuterVlanId))
		}
		if v.PoeDisabled != nil {
			poeDisabled = types.BoolValue(*v.PoeDisabled)
		}
		if v.PortNetwork != nil {
			portNetwork = types.StringValue(*v.PortNetwork)
		}
		if v.PreserveDscp != nil {
			preserveDscp = types.BoolValue(*v.PreserveDscp)
		}
		if v.Redundant != nil {
			redundant = types.BoolValue(*v.Redundant)
		}
		if v.RedundantGroup != nil {
			redundantGroup = types.Int64Value(int64(*v.RedundantGroup))
		}
		if v.RethIdx != nil {
			rethIdx = mistutils.GatewayPortConfigRethIdxAsString(v.RethIdx)
		}
		if v.RethNode != nil {
			rethNode = types.StringValue(*v.RethNode)
		}
		if v.Speed != nil {
			speed = types.StringValue(*v.Speed)
		}
		if v.SsrNoVirtualMac != nil {
			ssrNoVirtualMac = types.BoolValue(*v.SsrNoVirtualMac)
		}
		if v.SvrPortRange != nil {
			svrPortRange = types.StringValue(*v.SvrPortRange)
		}
		if v.TrafficShaping != nil {
			trafficShaping = portConfigTrafficShapingSdkToTerraform(ctx, diags, v.TrafficShaping)
		}
		if v.VlanId != nil {
			vlanId = mistutils.GatewayVlanAsString(*v.VlanId)
		}
		if len(v.VpnPaths) > 0 {
			vpnPaths = portConfigVpnPathsSdkToTerraform(ctx, diags, v.VpnPaths)
		}
		if v.WanArpPolicer != nil {
			wanArpPolicer = types.StringValue(string(*v.WanArpPolicer))
		}
		if v.WanExtIp != nil {
			wanExtIp = types.StringValue(*v.WanExtIp)
		}
		if v.WanDisableSpeedtest != nil {
			wanDisableSpeedtest = types.BoolValue(*v.WanDisableSpeedtest)
		}
		if len(v.WanExtraRoutes) > 0 {
			wanExtraRoutes = wanExtraRoutesPortConfigIpConfigSdkToTerraform(ctx, diags, v.WanExtraRoutes)
		}
		if v.WanProbeOverride != nil {
			wanProbeOverride = wanProbeOverridePortConfigIpConfigSdkToTerraform(ctx, diags, v.WanProbeOverride)
		}
		if v.WanSourceNat != nil {
			wanSourceNat = portConfigWanSourceNatSdkToTerraform(ctx, diags, v.WanSourceNat)
		}
		if v.WanType != nil {
			wanType = types.StringValue(string(*v.WanType))
		}

		var portUsageState = map[string]attr.Value{
			"ae_disable_lacp":       aeDisableLacp,
			"ae_idx":                aeIdx,
			"ae_lacp_force_up":      aeLacpForceUp,
			"aggregated":            aggregated,
			"critical":              critical,
			"description":           description,
			"disable_autoneg":       disableAutoneg,
			"disabled":              disabled,
			"dsl_type":              dslType,
			"dsl_vci":               dslVci,
			"dsl_vpi":               dslVpi,
			"duplex":                duplex,
			"ip_config":             ipConfig,
			"lte_apn":               lteApn,
			"lte_auth":              lteAuth,
			"lte_backup":            lteBackup,
			"lte_password":          ltePassword,
			"lte_username":          lteUsername,
			"mtu":                   mtu,
			"name":                  name,
			"networks":              networks,
			"outer_vlan_id":         outerVlanId,
			"poe_disabled":          poeDisabled,
			"port_network":          portNetwork,
			"preserve_dscp":         preserveDscp,
			"redundant":             redundant,
			"redundant_group":       redundantGroup,
			"reth_idx":              rethIdx,
			"reth_node":             rethNode,
			"reth_nodes":            rethNodes,
			"speed":                 speed,
			"ssr_no_virtual_mac":    ssrNoVirtualMac,
			"svr_port_range":        svrPortRange,
			"traffic_shaping":       trafficShaping,
			"usage":                 usage,
			"vlan_id":               vlanId,
			"vpn_paths":             vpnPaths,
			"wan_arp_policer":       wanArpPolicer,
			"wan_disable_speedtest": wanDisableSpeedtest,
			"wan_ext_ip":            wanExtIp,
			"wan_extra_routes":      wanExtraRoutes,
			"wan_networks":          wanNetworks,
			"wan_probe_override":    wanProbeOverride,
			"wan_source_nat":        wanSourceNat,
			"wan_type":              wanType,
		}
		portUsageObject, e := NewPortConfigValue(PortConfigValue{}.AttributeTypes(ctx), portUsageState)
		diags.Append(e...)
		stateValueMap[k] = portUsageObject
	}

	stateResult, e := types.MapValueFrom(ctx, PortConfigValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}
