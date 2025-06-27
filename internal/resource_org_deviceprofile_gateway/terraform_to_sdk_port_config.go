package resource_org_deviceprofile_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func wanExtraRoutesPortVpnPathTerraformToSdk(d basetypes.MapValue) map[string]models.WanExtraRoutes {
	dataMap := make(map[string]models.WanExtraRoutes)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(WanExtraRoutesValue)
		data := models.WanExtraRoutes{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		dataMap[k] = data
	}
	return dataMap
}
func wanProbeOverridePortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayWanProbeOverride {
	data := models.GatewayWanProbeOverride{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan, e := NewWanProbeOverrideValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
			return nil
		}
		if plan.Ips.IsNull() && !plan.Ips.IsUnknown() {
			data.Ips = mistutils.ListOfStringTerraformToSdk(plan.Ips)
		}
		if plan.ProbeProfile.ValueStringPointer() != nil {
			data.ProbeProfile = (*models.GatewayWanProbeOverrideProbeProfileEnum)(plan.ProbeProfile.ValueStringPointer())
		}
		return &data
	}
}
func gatewayPortVpnPathTerraformToSdk(ctx context.Context, d basetypes.MapValue) map[string]models.GatewayPortVpnPath {
	dataMap := make(map[string]models.GatewayPortVpnPath)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(VpnPathsValue)
		data := models.GatewayPortVpnPath{}
		if plan.BfdProfile.ValueStringPointer() != nil {
			data.BfdProfile = models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum(plan.BfdProfile.ValueString()))
		}
		if plan.BfdUseTunnelMode.ValueBoolPointer() != nil {
			data.BfdUseTunnelMode = plan.BfdUseTunnelMode.ValueBoolPointer()
		}
		if plan.Role.ValueStringPointer() != nil {
			data.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))
		}
		if plan.Preference.ValueInt64Pointer() != nil {
			data.Preference = models.ToPointer(int(plan.Preference.ValueInt64()))
		}
		if !plan.TrafficShaping.IsNull() && !plan.TrafficShaping.IsUnknown() {
			data.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, plan.TrafficShaping)
		}

		dataMap[k] = data
	}
	return dataMap
}

func gatewayPortTrafficShapingTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GatewayTrafficShaping {
	data := models.GatewayTrafficShaping{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewTrafficShapingValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.ClassPercentages.IsNull() && !plan.ClassPercentages.IsUnknown() {
			data.ClassPercentages = mistutils.ListOfIntTerraformToSdk(plan.ClassPercentages)
		}
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = plan.Enabled.ValueBoolPointer()
		}
		if plan.MaxTxKbps.ValueInt64Pointer() != nil {
			data.MaxTxKbps = models.ToPointer(int(plan.MaxTxKbps.ValueInt64()))
		}
		return &data
	}
}

func gatewayIpConfigTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GatewayPortConfigIpConfig {
	data := models.GatewayPortConfigIpConfig{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewPortIpConfigValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Dns.IsNull() && !plan.Dns.IsUnknown() {
			data.Dns = mistutils.ListOfStringTerraformToSdk(plan.Dns)
		}
		if plan.DnsSuffix.IsNull() && !plan.DnsSuffix.IsUnknown() {
			data.DnsSuffix = mistutils.ListOfStringTerraformToSdk(plan.DnsSuffix)
		}
		if plan.Gateway.ValueStringPointer() != nil {
			data.Gateway = plan.Gateway.ValueStringPointer()
		}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = plan.Ip.ValueStringPointer()
		}
		if plan.Netmask.ValueStringPointer() != nil {
			data.Netmask = plan.Netmask.ValueStringPointer()
		}
		if plan.Network.ValueStringPointer() != nil {
			data.Network = plan.Network.ValueStringPointer()
		}
		if plan.PoserPassword.ValueStringPointer() != nil {
			data.PoserPassword = plan.PoserPassword.ValueStringPointer()
		}
		if plan.PppoeUsername.ValueStringPointer() != nil {
			data.PppoeUsername = plan.PppoeUsername.ValueStringPointer()
		}
		if plan.PppoeAuth.ValueStringPointer() != nil {
			data.PppoeAuth = models.ToPointer(models.GatewayWanPpoeAuthEnum(plan.PppoeAuth.ValueString()))
		}
		if plan.PortIpConfigType.ValueStringPointer() != nil {
			data.Type = models.ToPointer(models.GatewayWanTypeEnum(plan.PortIpConfigType.ValueString()))
		}
		return &data
	}
}

func portConfigWanSourceNatTerraformToSdk(ctx context.Context, d basetypes.ObjectValue) *models.GatewayPortWanSourceNat {
	data := models.GatewayPortWanSourceNat{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewWanSourceNatValueMust(d.AttributeTypes(ctx), d.Attributes())
		if !plan.Disabled.IsNull() && !plan.Disabled.IsUnknown() {
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}
		if plan.NatPool.ValueStringPointer() != nil {
			data.NatPool = plan.NatPool.ValueStringPointer()
		}
		return &data
	}
}
func portConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayPortConfig {
	dataMap := make(map[string]models.GatewayPortConfig)
	for k, v := range d.Elements() {
		var vInterface interface{} = v
		plan := vInterface.(PortConfigValue)
		data := models.GatewayPortConfig{}

		if plan.AeDisableLacp.ValueBoolPointer() != nil {
			data.AeDisableLacp = plan.AeDisableLacp.ValueBoolPointer()
		}
		if plan.AeIdx.ValueStringPointer() != nil {
			data.AeIdx = models.NewOptional(plan.AeIdx.ValueStringPointer())
		}
		if plan.AeLacpForceUp.ValueBoolPointer() != nil {
			data.AeLacpForceUp = plan.AeLacpForceUp.ValueBoolPointer()
		}
		if plan.Aggregated.ValueBoolPointer() != nil {
			data.Aggregated = plan.Aggregated.ValueBoolPointer()
		}
		if plan.Critical.ValueBoolPointer() != nil {
			data.Critical = plan.Critical.ValueBoolPointer()
		}
		if plan.Usage.ValueStringPointer() != nil {
			data.Usage = models.GatewayPortUsageEnum(plan.Usage.ValueString())
		}
		if plan.Description.ValueStringPointer() != nil {
			data.Description = plan.Description.ValueStringPointer()
		}
		if plan.DisableAutoneg.ValueBoolPointer() != nil {
			data.DisableAutoneg = plan.DisableAutoneg.ValueBoolPointer()
		}
		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = plan.Disabled.ValueBoolPointer()
		}
		if plan.DslType.ValueStringPointer() != nil {
			data.DslType = models.ToPointer(models.GatewayPortDslTypeEnum(plan.DslType.ValueString()))
		}
		if plan.DslVci.ValueInt64Pointer() != nil {
			data.DslVci = models.ToPointer(int(plan.DslVci.ValueInt64()))
		}
		if plan.DslVpi.ValueInt64Pointer() != nil {
			data.DslVpi = models.ToPointer(int(plan.DslVpi.ValueInt64()))
		}
		if plan.Duplex.ValueStringPointer() != nil {
			data.Duplex = models.ToPointer(models.GatewayPortDuplexEnum(plan.Duplex.ValueString()))
		}

		t, _ := plan.PortIpConfig.ToObjectValue(ctx)
		data.IpConfig = gatewayIpConfigTerraformToSdk(ctx, t)

		if plan.LteApn.ValueStringPointer() != nil {
			data.LteApn = plan.LteApn.ValueStringPointer()
		}
		if plan.LteAuth.ValueStringPointer() != nil {
			data.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(plan.LteAuth.ValueString()))
		}
		if plan.LteBackup.ValueBoolPointer() != nil {
			data.LteBackup = plan.LteBackup.ValueBoolPointer()
		}
		if plan.LtePassword.ValueStringPointer() != nil {
			data.LtePassword = plan.LtePassword.ValueStringPointer()
		}
		if plan.LteUsername.ValueStringPointer() != nil {
			data.LteUsername = plan.LteUsername.ValueStringPointer()
		}
		if plan.Mtu.ValueInt64Pointer() != nil {
			data.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = plan.Name.ValueStringPointer()
		}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mistutils.ListOfStringTerraformToSdk(plan.Networks)
		}
		if plan.OuterVlanId.ValueInt64Pointer() != nil {
			data.OuterVlanId = models.ToPointer(int(plan.OuterVlanId.ValueInt64()))
		}
		if plan.PoeDisabled.ValueBoolPointer() != nil {
			data.PoeDisabled = plan.PoeDisabled.ValueBoolPointer()
		}
		if plan.PortNetwork.ValueStringPointer() != nil {
			data.PortNetwork = plan.PortNetwork.ValueStringPointer()
		}
		if plan.PreserveDscp.ValueBoolPointer() != nil {
			data.PreserveDscp = plan.PreserveDscp.ValueBoolPointer()
		}
		if plan.Redundant.ValueBoolPointer() != nil {
			data.Redundant = plan.Redundant.ValueBoolPointer()
		}
		if plan.RedundantGroup.ValueInt64Pointer() != nil {
			data.RedundantGroup = models.ToPointer(int(plan.RedundantGroup.ValueInt64()))
		}
		if plan.RethIdx.ValueStringPointer() != nil {
			data.RethIdx = models.ToPointer(models.GatewayPortConfigRethIdxContainer.FromString(plan.RethIdx.ValueString()))
		}
		if plan.RethNode.ValueStringPointer() != nil {
			data.RethNode = plan.RethNode.ValueStringPointer()
		}
		if plan.Speed.ValueStringPointer() != nil {
			data.Speed = plan.Speed.ValueStringPointer()
		}
		if plan.SsrNoVirtualMac.ValueBoolPointer() != nil {
			data.SsrNoVirtualMac = plan.SsrNoVirtualMac.ValueBoolPointer()
		}
		if plan.SvrPortRange.ValueStringPointer() != nil {
			data.SvrPortRange = plan.SvrPortRange.ValueStringPointer()
		}

		data.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, plan.TrafficShaping)

		if plan.VlanId.ValueStringPointer() != nil {
			data.VlanId = models.ToPointer(models.GatewayPortVlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
		}

		data.VpnPaths = gatewayPortVpnPathTerraformToSdk(ctx, plan.VpnPaths)

		data.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(plan.WanArpPolicer.ValueString()))

		if plan.WanDisableSpeedtest.ValueBoolPointer() != nil {
			data.WanDisableSpeedtest = plan.WanDisableSpeedtest.ValueBoolPointer()
		}
		if plan.WanExtIp.ValueStringPointer() != nil {
			data.WanExtIp = plan.WanExtIp.ValueStringPointer()
		}

		data.WanExtraRoutes = wanExtraRoutesPortVpnPathTerraformToSdk(plan.WanExtraRoutes)

		if !plan.WanNetworks.IsNull() && !plan.WanNetworks.IsUnknown() {
			data.WanNetworks = mistutils.ListOfStringTerraformToSdk(plan.WanNetworks)
		}

		data.WanProbeOverride = wanProbeOverridePortVpnPathTerraformToSdk(ctx, diags, plan.WanProbeOverride)
		data.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, plan.WanSourceNat)

		if plan.WanType.ValueStringPointer() != nil {
			data.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(plan.WanType.ValueString()))
		}
		dataMap[k] = data
	}
	return dataMap
}
