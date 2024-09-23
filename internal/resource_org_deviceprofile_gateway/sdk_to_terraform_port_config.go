package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func portConfigIpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayPortConfigIpConfig) basetypes.ObjectValue {
	var dns basetypes.ListValue = types.ListNull(types.StringType)
	var dns_suffix basetypes.ListValue = types.ListNull(types.StringType)
	var gateway basetypes.StringValue
	var ip basetypes.StringValue
	var netmask basetypes.StringValue
	var network basetypes.StringValue
	var poser_password basetypes.StringValue
	var pppoe_auth basetypes.StringValue = types.StringValue("none")
	var pppoe_username basetypes.StringValue
	var ip_config_type basetypes.StringValue = types.StringValue("dhcp")

	if g != nil && g.Dns != nil {
		dns = mist_transform.ListOfStringSdkToTerraform(ctx, g.Dns)
	}
	if g != nil && g.DnsSuffix != nil {
		dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, g.DnsSuffix)
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
		poser_password = types.StringValue(*g.PoserPassword)
	}
	if g != nil && g.PppoeAuth != nil {
		pppoe_auth = types.StringValue(string(*g.PppoeAuth))
	}
	if g != nil && g.PppoeUsername != nil {
		pppoe_username = types.StringValue(*g.PppoeUsername)
	}
	if g != nil && g.Type != nil {
		ip_config_type = types.StringValue(string(*g.Type))
	}

	r_attr_type := PortIpConfigValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"dns":            dns,
		"dns_suffix":     dns_suffix,
		"gateway":        gateway,
		"ip":             ip,
		"netmask":        netmask,
		"network":        network,
		"poser_password": poser_password,
		"pppoe_username": pppoe_username,
		"pppoe_auth":     pppoe_auth,
		"type":           ip_config_type,
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigTrafficShappingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayTrafficShaping) basetypes.ObjectValue {
	var class_percentages basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var enabled basetypes.BoolValue = types.BoolValue(false)

	if g != nil && g.ClassPercentages != nil {
		class_percentages = mist_transform.ListOfIntSdkToTerraform(ctx, g.ClassPercentages)
	}
	if g != nil && g.Enabled != nil {
		enabled = types.BoolValue(*g.Enabled)
	}

	r_attr_type := TrafficShapingValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"class_percentages": class_percentages,
		"enabled":           enabled,
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigVpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g map[string]models.GatewayPortVpnPath) basetypes.MapValue {

	port_usage_type := VpnPathsValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range g {

		var bfd_profile basetypes.StringValue = types.StringValue("broadband")
		var bfd_use_tunnel_mode basetypes.BoolValue = types.BoolValue(false)
		var preference basetypes.Int64Value
		var role basetypes.StringValue = types.StringValue("spoke")
		var traffic_shaping basetypes.ObjectValue = types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))

		if v.BfdProfile != nil {
			bfd_profile = types.StringValue(string(*v.BfdProfile))
		}
		if v.BfdUseTunnelMode != nil {
			bfd_use_tunnel_mode = types.BoolValue(*v.BfdUseTunnelMode)
		}
		if v.Preference != nil {
			preference = types.Int64Value(int64(*v.Preference))
		}
		if v.TrafficShaping != nil {
			traffic_shaping = portConfigTrafficShappingSdkToTerraform(ctx, diags, v.TrafficShaping)
		}
		if v.Role != nil {
			role = types.StringValue(string(*v.Role))
		}

		var port_usage_state = map[string]attr.Value{
			"bfd_profile":         bfd_profile,
			"bfd_use_tunnel_mode": bfd_use_tunnel_mode,
			"preference":          preference,
			"role":                role,
			"traffic_shaping":     traffic_shaping,
		}
		port_usage_object, e := NewVpnPathsValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := VpnPathsValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}

func portConfigWanSourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, g *models.GatewayPortWanSourceNat) basetypes.ObjectValue {

	var disabled basetypes.BoolValue = types.BoolValue(false)
	var nat_pool basetypes.StringValue

	if g != nil && g.Disabled != nil {
		disabled = types.BoolValue(*g.Disabled)
	}
	if g != nil && g.NatPool != nil {
		nat_pool = types.StringValue(*g.NatPool)
	}

	r_attr_type := WanSourceNatValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"disabled": disabled,
		"nat_pool": nat_pool,
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)

	return r
}

func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.GatewayPortConfig) basetypes.MapValue {
	port_usage_type := PortConfigValue{}.AttributeTypes(ctx)
	state_value_map := make(map[string]attr.Value)
	for k, v := range d {
		var ae_disable_lacp basetypes.BoolValue = types.BoolValue(false)
		var ae_idx basetypes.StringValue
		var ae_lacp_force_up basetypes.BoolValue = types.BoolValue(false)
		var aggregated basetypes.BoolValue = types.BoolValue(false)
		var critical basetypes.BoolValue = types.BoolValue(false)
		var description basetypes.StringValue
		var disable_autoneg basetypes.BoolValue = types.BoolValue(false)
		var disabled basetypes.BoolValue = types.BoolValue(false)
		var dsl_type basetypes.StringValue = types.StringValue("vdsl")
		var dsl_vci basetypes.Int64Value = types.Int64Value(35)
		var dsl_vpi basetypes.Int64Value = types.Int64Value(0)
		var duplex basetypes.StringValue = types.StringValue("auto")
		var ip_config basetypes.ObjectValue = types.ObjectNull(PortIpConfigValue{}.AttributeTypes(ctx))
		var lte_apn basetypes.StringValue
		var lte_auth basetypes.StringValue = types.StringValue("none")
		var lte_backup basetypes.BoolValue
		var lte_password basetypes.StringValue
		var lte_username basetypes.StringValue
		var mtu basetypes.Int64Value
		var name basetypes.StringValue
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraform(ctx, v.Networks)
		var outer_vlan_id basetypes.Int64Value
		var poe_disabled basetypes.BoolValue = types.BoolValue(false)
		var port_network basetypes.StringValue
		var preserve_dscp basetypes.BoolValue = types.BoolValue(true)
		var redundant basetypes.BoolValue
		var reth_idx basetypes.Int64Value
		var reth_node basetypes.StringValue
		var reth_nodes basetypes.ListValue = mist_transform.ListOfStringSdkToTerraform(ctx, v.RethNodes)
		var speed basetypes.StringValue = types.StringValue("auto")
		var ssr_no_virtual_mac basetypes.BoolValue = types.BoolValue(false)
		var svr_port_range basetypes.StringValue = types.StringValue("none")
		var traffic_shaping basetypes.ObjectValue = types.ObjectNull(TrafficShapingValue{}.AttributeTypes(ctx))
		var usage basetypes.StringValue = types.StringValue(string(v.Usage))
		var vlan_id basetypes.Int64Value
		var vpn_paths basetypes.MapValue = types.MapNull(VpnPathsValue{}.Type(ctx))
		var wan_arp_policer basetypes.StringValue = types.StringValue("default")
		var wan_ext_ip basetypes.StringValue
		var wan_source_nat basetypes.ObjectValue = types.ObjectNull(WanSourceNatValue{}.AttributeTypes(ctx))
		var wan_type basetypes.StringValue = types.StringValue("broadband")

		if v.AeDisableLacp != nil {
			ae_disable_lacp = types.BoolValue(*v.AeDisableLacp)
		}
		if v.AeIdx.Value() != nil {
			ae_idx = types.StringValue(*v.AeIdx.Value())
		}
		if v.AeLacpForceUp != nil {
			ae_lacp_force_up = types.BoolValue(*v.AeLacpForceUp)
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
			disable_autoneg = types.BoolValue(*v.DisableAutoneg)
		}
		if v.Disabled != nil {
			disabled = types.BoolValue(*v.Disabled)
		}
		if v.DslType != nil {
			dsl_type = types.StringValue(string(*v.DslType))
		}
		if v.DslVci != nil {
			dsl_vci = types.Int64Value(int64(*v.DslVci))
		}
		if v.DslVpi != nil {
			dsl_vpi = types.Int64Value(int64(*v.DslVpi))
		}
		if v.Duplex != nil {
			duplex = types.StringValue(string(*v.Duplex))
		}
		if v.IpConfig != nil {
			ip_config = portConfigIpConfigSdkToTerraform(ctx, diags, v.IpConfig)
		}
		if v.LteApn != nil {
			lte_apn = types.StringValue(*v.LteApn)
		}
		if v.LteAuth != nil {
			lte_auth = types.StringValue(string(*v.LteAuth))
		}
		if v.LteBackup != nil {
			lte_backup = types.BoolValue(*v.LteBackup)
		}
		if v.LtePassword != nil {
			lte_password = types.StringValue(*v.LtePassword)
		}
		if v.LteUsername != nil {
			lte_username = types.StringValue(*v.LteUsername)
		}
		if v.Mtu != nil {
			mtu = types.Int64Value(int64(*v.Mtu))
		}
		if v.Name != nil {
			name = types.StringValue(*v.Name)
		}
		if v.OuterVlanId != nil {
			outer_vlan_id = types.Int64Value(int64(*v.OuterVlanId))
		}
		if v.PoeDisabled != nil {
			poe_disabled = types.BoolValue(*v.PoeDisabled)
		}
		if v.PortNetwork != nil {
			port_network = types.StringValue(*v.PortNetwork)
		}
		if v.PreserveDscp != nil {
			preserve_dscp = types.BoolValue(*v.PreserveDscp)
		}
		if v.Redundant != nil {
			redundant = types.BoolValue(*v.Redundant)
		}
		if v.RethIdx != nil {
			reth_idx = types.Int64Value(int64(*v.RethIdx))
		}
		if v.RethNode != nil {
			reth_node = types.StringValue(*v.RethNode)
		}
		if v.Speed != nil {
			speed = types.StringValue(*v.Speed)
		}
		if v.SsrNoVirtualMac != nil {
			ssr_no_virtual_mac = types.BoolValue(*v.SsrNoVirtualMac)
		}
		if v.SvrPortRange != nil {
			svr_port_range = types.StringValue(*v.SvrPortRange)
		}
		if v.TrafficShaping != nil {
			traffic_shaping = portConfigTrafficShappingSdkToTerraform(ctx, diags, v.TrafficShaping)
		}
		if v.VlanId != nil {
			vlan_id = types.Int64Value(int64(*v.VlanId))
		}
		if v.VpnPaths != nil && len(v.VpnPaths) > 0 {
			vpn_paths = portConfigVpnPathsSdkToTerraform(ctx, diags, v.VpnPaths)
		}
		if v.WanArpPolicer != nil {
			wan_arp_policer = types.StringValue(string(*v.WanArpPolicer))
		}
		if v.WanExtIp != nil {
			wan_ext_ip = types.StringValue(*v.WanExtIp)
		}
		if v.WanSourceNat != nil {
			wan_source_nat = portConfigWanSourceNatSdkToTerraform(ctx, diags, v.WanSourceNat)
		}
		if v.WanType != nil {
			wan_type = types.StringValue(string(*v.WanType))
		}

		var port_usage_state = map[string]attr.Value{
			"ae_disable_lacp":    ae_disable_lacp,
			"ae_idx":             ae_idx,
			"ae_lacp_force_up":   ae_lacp_force_up,
			"aggregated":         aggregated,
			"critical":           critical,
			"description":        description,
			"disable_autoneg":    disable_autoneg,
			"disabled":           disabled,
			"dsl_type":           dsl_type,
			"dsl_vci":            dsl_vci,
			"dsl_vpi":            dsl_vpi,
			"duplex":             duplex,
			"ip_config":          ip_config,
			"lte_apn":            lte_apn,
			"lte_auth":           lte_auth,
			"lte_backup":         lte_backup,
			"lte_password":       lte_password,
			"lte_username":       lte_username,
			"mtu":                mtu,
			"name":               name,
			"networks":           networks,
			"outer_vlan_id":      outer_vlan_id,
			"poe_disabled":       poe_disabled,
			"port_network":       port_network,
			"preserve_dscp":      preserve_dscp,
			"redundant":          redundant,
			"reth_idx":           reth_idx,
			"reth_node":          reth_node,
			"reth_nodes":         reth_nodes,
			"speed":              speed,
			"ssr_no_virtual_mac": ssr_no_virtual_mac,
			"svr_port_range":     svr_port_range,
			"traffic_shaping":    traffic_shaping,
			"usage":              usage,
			"vlan_id":            vlan_id,
			"vpn_paths":          vpn_paths,
			"wan_arp_policer":    wan_arp_policer,
			"wan_ext_ip":         wan_ext_ip,
			"wan_source_nat":     wan_source_nat,
			"wan_type":           wan_type,
		}
		port_usage_object, e := NewPortConfigValue(port_usage_type, port_usage_state)
		diags.Append(e...)
		state_value_map[k] = port_usage_object
	}
	state_type := PortConfigValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
