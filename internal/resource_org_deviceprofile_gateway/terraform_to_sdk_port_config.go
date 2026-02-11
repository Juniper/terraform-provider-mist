package resource_org_deviceprofile_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func wanExtraRoutesPortVpnPathTerraformToSdk(data basetypes.MapValue) map[string]models.WanExtraRoutes {
	dataMap := make(map[string]models.WanExtraRoutes)
	for key, val := range data.Elements() {
		plan := val.(WanExtraRoutesValue)
		if plan.Via.ValueStringPointer() == nil {
			continue
		}

		dataMap[key] = models.WanExtraRoutes{
			Via: plan.Via.ValueStringPointer(),
		}
	}
	return dataMap
}

func wanExtraRoutes6PortVpnPathTerraformToSdk(data basetypes.MapValue) map[string]models.WanExtraRoutes {
	dataMap := make(map[string]models.WanExtraRoutes)
	for key, val := range data.Elements() {
		plan := val.(WanExtraRoutes6Value)
		if plan.Via.ValueStringPointer() == nil {
			continue
		}

		dataMap[key] = models.WanExtraRoutes{
			Via: plan.Via.ValueStringPointer(),
		}
	}
	return dataMap
}

func wanProbeOverridePortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, data basetypes.ObjectValue) *models.GatewayWanProbeOverride {
	result := models.GatewayWanProbeOverride{}
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	plan, err := NewWanProbeOverrideValue(data.AttributeTypes(ctx), data.Attributes())
	if err != nil {
		diags.Append(err...)
		return nil
	}

	if !plan.Ips.IsNull() && !plan.Ips.IsUnknown() {
		result.Ips = mistutils.ListOfStringTerraformToSdk(plan.Ips)
	}

	if !plan.Ip6s.IsNull() && !plan.Ip6s.IsUnknown() {
		result.Ip6s = mistutils.ListOfStringTerraformToSdk(plan.Ip6s)
	}

	if plan.ProbeProfile.ValueStringPointer() != nil {
		result.ProbeProfile = (*models.GatewayWanProbeOverrideProbeProfileEnum)(plan.ProbeProfile.ValueStringPointer())
	}

	return &result
}

func gatewayPortVpnPathTerraformToSdk(ctx context.Context, data basetypes.MapValue) map[string]models.GatewayPortVpnPath {
	dataMap := make(map[string]models.GatewayPortVpnPath)
	for key, val := range data.Elements() {
		var result models.GatewayPortVpnPath
		plan := val.(VpnPathsValue)
		if plan.BfdProfile.ValueStringPointer() != nil {
			result.BfdProfile = models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum(plan.BfdProfile.ValueString()))
		}

		if plan.BfdUseTunnelMode.ValueBoolPointer() != nil {
			result.BfdUseTunnelMode = plan.BfdUseTunnelMode.ValueBoolPointer()
		}

		if plan.Role.ValueStringPointer() != nil {
			result.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))
		}

		if plan.Preference.ValueInt64Pointer() != nil {
			result.Preference = models.ToPointer(int(plan.Preference.ValueInt64()))
		}

		if !plan.TrafficShaping.IsNull() && !plan.TrafficShaping.IsUnknown() {
			result.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, plan.TrafficShaping)
		}

		dataMap[key] = result
	}

	return dataMap
}

func gatewayPortTrafficShapingTerraformToSdk(ctx context.Context, data basetypes.ObjectValue) *models.GatewayTrafficShaping {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var result models.GatewayTrafficShaping
	plan := NewTrafficShapingValueMust(data.AttributeTypes(ctx), data.Attributes())
	if !plan.ClassPercentages.IsNull() && !plan.ClassPercentages.IsUnknown() {
		result.ClassPercentages = mistutils.ListOfIntTerraformToSdk(plan.ClassPercentages)
	}

	if plan.Enabled.ValueBoolPointer() != nil {
		result.Enabled = plan.Enabled.ValueBoolPointer()
	}

	if plan.MaxTxKbps.ValueInt64Pointer() != nil {
		result.MaxTxKbps = models.ToPointer(int(plan.MaxTxKbps.ValueInt64()))
	}

	return &result
}

func gatewayIpConfigTerraformToSdk(ctx context.Context, data basetypes.ObjectValue) *models.GatewayPortConfigIpConfig {
	var result models.GatewayPortConfigIpConfig
	if data.IsNull() || data.IsUnknown() {
		return &result
	}

	plan := NewPortIpConfigValueMust(data.AttributeTypes(ctx), data.Attributes())
	if !plan.Dns.IsNull() && !plan.Dns.IsUnknown() {
		result.Dns = mistutils.ListOfStringTerraformToSdk(plan.Dns)
	}

	if !plan.DnsSuffix.IsNull() && !plan.DnsSuffix.IsUnknown() {
		result.DnsSuffix = mistutils.ListOfStringTerraformToSdk(plan.DnsSuffix)
	}

	if plan.Gateway.ValueStringPointer() != nil {
		result.Gateway = plan.Gateway.ValueStringPointer()
	}

	if plan.Gateway6.ValueStringPointer() != nil {
		result.Gateway6 = plan.Gateway6.ValueStringPointer()
	}

	if plan.Ip.ValueStringPointer() != nil {
		result.Ip = plan.Ip.ValueStringPointer()
	}

	if plan.Ip6.ValueStringPointer() != nil {
		result.Ip6 = plan.Ip6.ValueStringPointer()
	}

	if plan.Netmask.ValueStringPointer() != nil {
		result.Netmask = plan.Netmask.ValueStringPointer()
	}

	if plan.Netmask6.ValueStringPointer() != nil {
		result.Netmask6 = plan.Netmask6.ValueStringPointer()
	}

	if plan.Network.ValueStringPointer() != nil {
		result.Network = plan.Network.ValueStringPointer()
	}

	if plan.PoserPassword.ValueStringPointer() != nil {
		result.PoserPassword = plan.PoserPassword.ValueStringPointer()
	}

	if plan.PppoeUsername.ValueStringPointer() != nil {
		result.PppoeUsername = plan.PppoeUsername.ValueStringPointer()
	}

	if plan.PppoeAuth.ValueStringPointer() != nil {
		result.PppoeAuth = models.ToPointer(models.GatewayWanPpoeAuthEnum(plan.PppoeAuth.ValueString()))
	}

	if plan.PortIpConfigType.ValueStringPointer() != nil {
		result.Type = models.ToPointer(models.GatewayWanTypeEnum(plan.PortIpConfigType.ValueString()))
	}

	if plan.Type6.ValueStringPointer() != nil {
		result.Type6 = models.ToPointer(models.GatewayWanType6Enum(plan.Type6.ValueString()))
	}

	return &result
}

func portConfigWanSourceNatTerraformToSdk(ctx context.Context, data basetypes.ObjectValue) *models.GatewayPortWanSourceNat {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var result models.GatewayPortWanSourceNat
	plan := NewWanSourceNatValueMust(data.AttributeTypes(ctx), data.Attributes())
	if !plan.Disabled.IsNull() && !plan.Disabled.IsUnknown() {
		result.Disabled = plan.Disabled.ValueBoolPointer()
	}

	if plan.NatPool.ValueStringPointer() != nil {
		result.NatPool = plan.NatPool.ValueStringPointer()
	}

	if plan.Nat6Pool.ValueStringPointer() != nil {
		result.Nat6Pool = plan.Nat6Pool.ValueStringPointer()
	}

	return &result
}

func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, data basetypes.MapValue) map[string]models.GatewayPortConfig {
	dataMap := make(map[string]models.GatewayPortConfig)
	for key, val := range data.Elements() {
		var result models.GatewayPortConfig
		portConfig := val.(PortConfigValue)
		if portConfig.AeDisableLacp.ValueBoolPointer() != nil {
			result.AeDisableLacp = portConfig.AeDisableLacp.ValueBoolPointer()
		}

		if portConfig.AeIdx.ValueStringPointer() != nil {
			result.AeIdx = models.NewOptional(portConfig.AeIdx.ValueStringPointer())
		}

		if portConfig.AeLacpForceUp.ValueBoolPointer() != nil {
			result.AeLacpForceUp = portConfig.AeLacpForceUp.ValueBoolPointer()
		}

		if portConfig.Aggregated.ValueBoolPointer() != nil {
			result.Aggregated = portConfig.Aggregated.ValueBoolPointer()
		}

		if portConfig.Critical.ValueBoolPointer() != nil {
			result.Critical = portConfig.Critical.ValueBoolPointer()
		}

		if portConfig.Usage.ValueStringPointer() != nil {
			result.Usage = models.GatewayPortUsageEnum(portConfig.Usage.ValueString())
		}

		if portConfig.Description.ValueStringPointer() != nil {
			result.Description = portConfig.Description.ValueStringPointer()
		}

		if portConfig.DisableAutoneg.ValueBoolPointer() != nil {
			result.DisableAutoneg = portConfig.DisableAutoneg.ValueBoolPointer()
		}

		if portConfig.Disabled.ValueBoolPointer() != nil {
			result.Disabled = portConfig.Disabled.ValueBoolPointer()
		}

		if portConfig.DslType.ValueStringPointer() != nil {
			result.DslType = models.ToPointer(models.GatewayPortDslTypeEnum(portConfig.DslType.ValueString()))
		}

		if portConfig.DslVci.ValueInt64Pointer() != nil {
			result.DslVci = models.ToPointer(int(portConfig.DslVci.ValueInt64()))
		}

		if portConfig.DslVpi.ValueInt64Pointer() != nil {
			result.DslVpi = models.ToPointer(int(portConfig.DslVpi.ValueInt64()))
		}

		if portConfig.Duplex.ValueStringPointer() != nil {
			result.Duplex = models.ToPointer(models.GatewayPortDuplexEnum(portConfig.Duplex.ValueString()))
		}

		if !portConfig.PortIpConfig.IsNull() && !portConfig.PortIpConfig.IsUnknown() {
			t, _ := portConfig.PortIpConfig.ToObjectValue(ctx)
			result.IpConfig = gatewayIpConfigTerraformToSdk(ctx, t)
		}

		if portConfig.LteApn.ValueStringPointer() != nil {
			result.LteApn = portConfig.LteApn.ValueStringPointer()
		}

		if portConfig.LteAuth.ValueStringPointer() != nil {
			result.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(portConfig.LteAuth.ValueString()))
		}

		if portConfig.LteBackup.ValueBoolPointer() != nil {
			result.LteBackup = portConfig.LteBackup.ValueBoolPointer()
		}

		if portConfig.LtePassword.ValueStringPointer() != nil {
			result.LtePassword = portConfig.LtePassword.ValueStringPointer()
		}

		if portConfig.LteUsername.ValueStringPointer() != nil {
			result.LteUsername = portConfig.LteUsername.ValueStringPointer()
		}

		if portConfig.Mtu.ValueInt64Pointer() != nil {
			result.Mtu = models.ToPointer(int(portConfig.Mtu.ValueInt64()))
		}

		if portConfig.Name.ValueStringPointer() != nil {
			result.Name = portConfig.Name.ValueStringPointer()
		}

		if !portConfig.Networks.IsNull() && !portConfig.Networks.IsUnknown() {
			result.Networks = mistutils.ListOfStringTerraformToSdk(portConfig.Networks)
		}

		if portConfig.OuterVlanId.ValueInt64Pointer() != nil {
			result.OuterVlanId = models.ToPointer(int(portConfig.OuterVlanId.ValueInt64()))
		}

		if portConfig.PoeDisabled.ValueBoolPointer() != nil {
			result.PoeDisabled = portConfig.PoeDisabled.ValueBoolPointer()
		}

		if portConfig.PortNetwork.ValueStringPointer() != nil {
			result.PortNetwork = portConfig.PortNetwork.ValueStringPointer()
		}

		if portConfig.PreserveDscp.ValueBoolPointer() != nil {
			result.PreserveDscp = portConfig.PreserveDscp.ValueBoolPointer()
		}

		if portConfig.Redundant.ValueBoolPointer() != nil {
			result.Redundant = portConfig.Redundant.ValueBoolPointer()
		}

		if portConfig.RedundantGroup.ValueInt64Pointer() != nil {
			result.RedundantGroup = models.ToPointer(int(portConfig.RedundantGroup.ValueInt64()))
		}

		if portConfig.RethIdx.ValueStringPointer() != nil {
			result.RethIdx = models.ToPointer(models.GatewayPortConfigRethIdxContainer.FromString(portConfig.RethIdx.ValueString()))
		}

		if portConfig.RethNode.ValueStringPointer() != nil {
			result.RethNode = portConfig.RethNode.ValueStringPointer()
		}

		if portConfig.Speed.ValueStringPointer() != nil {
			result.Speed = portConfig.Speed.ValueStringPointer()
		}

		if portConfig.SsrNoVirtualMac.ValueBoolPointer() != nil {
			result.SsrNoVirtualMac = portConfig.SsrNoVirtualMac.ValueBoolPointer()
		}

		if portConfig.SvrPortRange.ValueStringPointer() != nil {
			result.SvrPortRange = portConfig.SvrPortRange.ValueStringPointer()
		}

		if !portConfig.TrafficShaping.IsNull() && !portConfig.TrafficShaping.IsUnknown() {
			result.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, portConfig.TrafficShaping)
		}

		if portConfig.VlanId.ValueStringPointer() != nil {
			result.VlanId = models.ToPointer(models.GatewayPortVlanIdWithVariableContainer.FromString(portConfig.VlanId.ValueString()))
		}

		if !portConfig.VpnPaths.IsNull() && !portConfig.VpnPaths.IsUnknown() {
			result.VpnPaths = gatewayPortVpnPathTerraformToSdk(ctx, portConfig.VpnPaths)
		}

		if portConfig.WanArpPolicer.ValueStringPointer() != nil {
			result.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(portConfig.WanArpPolicer.ValueString()))
		}

		if portConfig.WanDisableSpeedtest.ValueBoolPointer() != nil {
			result.WanDisableSpeedtest = portConfig.WanDisableSpeedtest.ValueBoolPointer()
		}

		if portConfig.WanExtIp.ValueStringPointer() != nil {
			result.WanExtIp = portConfig.WanExtIp.ValueStringPointer()
		}

		if !portConfig.WanExtIp6.IsNull() && !portConfig.WanExtIp6.IsUnknown() {
			result.WanExtIp6 = portConfig.WanExtIp6.ValueStringPointer()
		}

		if !portConfig.WanExtraRoutes.IsNull() && !portConfig.WanExtraRoutes.IsUnknown() {
			result.WanExtraRoutes = wanExtraRoutesPortVpnPathTerraformToSdk(portConfig.WanExtraRoutes)
		}

		if !portConfig.WanExtraRoutes6.IsNull() && !portConfig.WanExtraRoutes6.IsUnknown() {
			result.WanExtraRoutes6 = wanExtraRoutes6PortVpnPathTerraformToSdk(portConfig.WanExtraRoutes6)
		}

		if !portConfig.WanNetworks.IsNull() && !portConfig.WanNetworks.IsUnknown() {
			result.WanNetworks = mistutils.ListOfStringTerraformToSdk(portConfig.WanNetworks)
		}

		if !portConfig.WanProbeOverride.IsNull() && !portConfig.WanProbeOverride.IsUnknown() {
			result.WanProbeOverride = wanProbeOverridePortVpnPathTerraformToSdk(ctx, diags, portConfig.WanProbeOverride)
		}

		if !portConfig.WanSourceNat.IsNull() && !portConfig.WanSourceNat.IsUnknown() {
			result.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, portConfig.WanSourceNat)
		}

		if portConfig.WanType.ValueStringPointer() != nil {
			result.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(portConfig.WanType.ValueString()))
		}

		dataMap[key] = result
	}

	return dataMap
}
