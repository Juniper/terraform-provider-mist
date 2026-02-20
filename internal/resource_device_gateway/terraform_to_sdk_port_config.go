package resource_device_gateway

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
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	plan, err := NewWanProbeOverrideValue(data.AttributeTypes(ctx), data.Attributes())
	if err != nil {
		diags.Append(err...)
		return nil
	}

	var result models.GatewayWanProbeOverride
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

		if plan.Preference.ValueInt64Pointer() != nil {
			result.Preference = models.ToPointer(int(plan.Preference.ValueInt64()))
		}

		if plan.Role.ValueStringPointer() != nil {
			result.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))
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

func gatewayIpConfigTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GatewayPortConfigIpConfig {
	var result models.GatewayPortConfigIpConfig
	if d.IsNull() || d.IsUnknown() {
		return &result
	}

	plan := NewPortIpConfigValueMust(d.AttributeTypes(ctx), d.Attributes())
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

func portConfigWanSourceNatTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GatewayPortWanSourceNat {
	if d.IsNull() || d.IsUnknown() {
		return nil
	}

	var result models.GatewayPortWanSourceNat
	plan := NewWanSourceNatValueMust(d.AttributeTypes(ctx), d.Attributes())
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
		plan := val.(PortConfigValue)
		if plan.AeDisableLacp.ValueBoolPointer() != nil {
			result.AeDisableLacp = plan.AeDisableLacp.ValueBoolPointer()
		}

		if plan.AeIdx.ValueStringPointer() != nil {
			result.AeIdx = models.NewOptional(plan.AeIdx.ValueStringPointer())
		}

		if plan.AeLacpForceUp.ValueBoolPointer() != nil {
			result.AeLacpForceUp = plan.AeLacpForceUp.ValueBoolPointer()
		}

		if plan.Aggregated.ValueBoolPointer() != nil {
			result.Aggregated = plan.Aggregated.ValueBoolPointer()
		}

		if plan.Critical.ValueBoolPointer() != nil {
			result.Critical = plan.Critical.ValueBoolPointer()
		}

		if plan.Usage.ValueStringPointer() != nil {
			result.Usage = models.GatewayPortUsageEnum(plan.Usage.ValueString())
		}

		if plan.Description.ValueStringPointer() != nil {
			result.Description = plan.Description.ValueStringPointer()
		}

		if plan.DisableAutoneg.ValueBoolPointer() != nil {
			result.DisableAutoneg = plan.DisableAutoneg.ValueBoolPointer()
		}

		if plan.Disabled.ValueBoolPointer() != nil {
			result.Disabled = plan.Disabled.ValueBoolPointer()
		}

		if plan.DslType.ValueStringPointer() != nil {
			result.DslType = models.ToPointer(models.GatewayPortDslTypeEnum(plan.DslType.ValueString()))
		}

		if plan.DslVci.ValueInt64Pointer() != nil {
			result.DslVci = models.ToPointer(int(plan.DslVci.ValueInt64()))
		}

		if plan.DslVpi.ValueInt64Pointer() != nil {
			result.DslVpi = models.ToPointer(int(plan.DslVpi.ValueInt64()))
		}

		if plan.Duplex.ValueStringPointer() != nil {
			result.Duplex = models.ToPointer(models.GatewayPortDuplexEnum(plan.Duplex.ValueString()))
		}

		if !plan.PortIpConfig.IsNull() && !plan.PortIpConfig.IsUnknown() {
			val, err := plan.PortIpConfig.ToObjectValue(ctx)
			if err != nil {
				diags.Append(err...)
			}
			result.IpConfig = gatewayIpConfigTerraformToSdk(ctx, val)
		}

		if plan.LteApn.ValueStringPointer() != nil {
			result.LteApn = plan.LteApn.ValueStringPointer()
		}

		if plan.LteAuth.ValueStringPointer() != nil {
			result.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(plan.LteAuth.ValueString()))
		}

		if plan.LteBackup.ValueBoolPointer() != nil {
			result.LteBackup = plan.LteBackup.ValueBoolPointer()
		}

		if plan.LtePassword.ValueStringPointer() != nil {
			result.LtePassword = plan.LtePassword.ValueStringPointer()
		}

		if plan.LteUsername.ValueStringPointer() != nil {
			result.LteUsername = plan.LteUsername.ValueStringPointer()
		}

		if plan.Mtu.ValueInt64Pointer() != nil {
			result.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
		}

		if plan.Name.ValueStringPointer() != nil {
			result.Name = plan.Name.ValueStringPointer()
		}

		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			result.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}

		if plan.OuterVlanId.ValueInt64Pointer() != nil {
			result.OuterVlanId = models.ToPointer(int(plan.OuterVlanId.ValueInt64()))
		}

		if plan.PoeDisabled.ValueBoolPointer() != nil {
			result.PoeDisabled = plan.PoeDisabled.ValueBoolPointer()
		}

		if plan.PortNetwork.ValueStringPointer() != nil {
			result.PortNetwork = plan.PortNetwork.ValueStringPointer()
		}

		if plan.PreserveDscp.ValueBoolPointer() != nil {
			result.PreserveDscp = plan.PreserveDscp.ValueBoolPointer()
		}

		if plan.Redundant.ValueBoolPointer() != nil {
			result.Redundant = plan.Redundant.ValueBoolPointer()
		}

		if plan.RedundantGroup.ValueInt64Pointer() != nil {
			result.RedundantGroup = models.ToPointer(int(plan.RedundantGroup.ValueInt64()))
		}

		if plan.RethIdx.ValueStringPointer() != nil {
			result.RethIdx = models.ToPointer(models.GatewayPortConfigRethIdxContainer.FromString(plan.RethIdx.ValueString()))
		}

		if plan.RethNode.ValueStringPointer() != nil {
			result.RethNode = plan.RethNode.ValueStringPointer()
		}

		if plan.Speed.ValueStringPointer() != nil {
			result.Speed = plan.Speed.ValueStringPointer()
		}

		if plan.SsrNoVirtualMac.ValueBoolPointer() != nil {
			result.SsrNoVirtualMac = plan.SsrNoVirtualMac.ValueBoolPointer()
		}

		if plan.SvrPortRange.ValueStringPointer() != nil {
			result.SvrPortRange = plan.SvrPortRange.ValueStringPointer()
		}

		if !plan.TrafficShaping.IsNull() && !plan.TrafficShaping.IsUnknown() {
			result.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, plan.TrafficShaping)
		}

		if plan.VlanId.ValueStringPointer() != nil {
			result.VlanId = models.ToPointer(models.GatewayPortVlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		if !plan.VpnPaths.IsNull() && !plan.VpnPaths.IsUnknown() {
			result.VpnPaths = gatewayPortVpnPathTerraformToSdk(ctx, plan.VpnPaths)
		}

		if plan.WanArpPolicer.ValueStringPointer() != nil {
			result.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(plan.WanArpPolicer.ValueString()))
		}

		if plan.WanSpeedtestMode.ValueStringPointer() != nil {
			result.WanSpeedtestMode = models.ToPointer(models.GatewayPortConfigWanSpeedtestModeEnum(plan.WanSpeedtestMode.ValueString()))
		}

		if plan.WanExtIp.ValueStringPointer() != nil {
			result.WanExtIp = plan.WanExtIp.ValueStringPointer()
		}

		if !plan.WanExtIp6.IsNull() && !plan.WanExtIp6.IsUnknown() {
			result.WanExtIp6 = plan.WanExtIp6.ValueStringPointer()
		}

		if !plan.WanExtraRoutes.IsNull() && !plan.WanExtraRoutes.IsUnknown() {
			result.WanExtraRoutes = wanExtraRoutesPortVpnPathTerraformToSdk(plan.WanExtraRoutes)
		}

		if !plan.WanExtraRoutes6.IsNull() && !plan.WanExtraRoutes6.IsUnknown() {
			result.WanExtraRoutes6 = wanExtraRoutes6PortVpnPathTerraformToSdk(plan.WanExtraRoutes6)
		}

		if !plan.WanNetworks.IsNull() && !plan.WanNetworks.IsUnknown() {
			result.WanNetworks = mistutils.ListOfStringTerraformToSdk(plan.WanNetworks)
		}

		if !plan.WanProbeOverride.IsNull() && !plan.WanProbeOverride.IsUnknown() {
			result.WanProbeOverride = wanProbeOverridePortVpnPathTerraformToSdk(ctx, diags, plan.WanProbeOverride)
		}

		if !plan.WanSourceNat.IsNull() && !plan.WanSourceNat.IsUnknown() {
			result.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, plan.WanSourceNat)
		}

		if plan.WanType.ValueStringPointer() != nil {
			result.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(plan.WanType.ValueString()))
		}

		dataMap[key] = result
	}
	return dataMap
}
