package resource_org_gatewaytemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func wanExtraRoutesPortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.WanExtraRoutes {
	data_map := make(map[string]models.WanExtraRoutes)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(WanExtraRoutesValue)
		data := models.WanExtraRoutes{}
		if plan.Via.ValueStringPointer() != nil {
			data.Via = plan.Via.ValueStringPointer()
		}

		data_map[k] = data
	}
	return data_map
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
			data.Ips = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Ips)
		}
		if plan.ProbeProfile.ValueStringPointer() != nil {
			data.ProbeProfile = (*models.GatewayWanProbeOverrideProbeProfileEnum)(plan.ProbeProfile.ValueStringPointer())
		}
		return &data
	}
}
func gatewayPortVpnPathTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.GatewayPortVpnPath {
	data_map := make(map[string]models.GatewayPortVpnPath)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VpnPathsValue)
		data := models.GatewayPortVpnPath{}
		if plan.BfdProfile.ValueStringPointer() != nil {
			data.BfdProfile = models.ToPointer(models.GatewayPortVpnPathBfdProfileEnum(plan.BfdProfile.ValueString()))
		}
		if plan.BfdUseTunnelMode.ValueBoolPointer() != nil {
			data.BfdUseTunnelMode = models.ToPointer(plan.BfdUseTunnelMode.ValueBool())
		}
		if plan.Role.ValueStringPointer() != nil {
			data.Role = models.ToPointer(models.GatewayPortVpnPathRoleEnum(plan.Role.ValueString()))
		}
		if plan.LinkName.ValueStringPointer() != nil {
			data.LinkName = plan.LinkName.ValueStringPointer()
		}
		if plan.Preference.ValueInt64Pointer() != nil {
			data.Preference = models.ToPointer(int(plan.Preference.ValueInt64()))
		}
		if !plan.TrafficShaping.IsNull() && !plan.TrafficShaping.IsUnknown() {
			data.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)
		}

		data_map[k] = data
	}
	return data_map
}

func gatewayPortTrafficShapingTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayTrafficShaping {
	data := models.GatewayTrafficShaping{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewTrafficShapingValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.ClassPercentages.IsNull() && !plan.ClassPercentages.IsUnknown() {
			data.ClassPercentages = mist_transform.ListOfIntTerraformToSdk(ctx, plan.ClassPercentages)
		}
		if plan.Enabled.ValueBoolPointer() != nil {
			data.Enabled = models.ToPointer(plan.Enabled.ValueBool())
		}
		return &data
	}
}

func gatewayIpConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayPortConfigIpConfig {
	data := models.GatewayPortConfigIpConfig{}
	if d.IsNull() || d.IsUnknown() {
		return nil
	} else {
		plan := NewPortIpConfigValueMust(d.AttributeTypes(ctx), d.Attributes())
		if plan.Dns.IsNull() && !plan.Dns.IsUnknown() {
			data.Dns = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Dns)
		}
		if plan.DnsSuffix.IsNull() && !plan.DnsSuffix.IsUnknown() {
			data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
		}
		if plan.Gateway.ValueStringPointer() != nil {
			data.Gateway = models.ToPointer(plan.Gateway.ValueString())
		}
		if plan.Ip.ValueStringPointer() != nil {
			data.Ip = models.ToPointer(plan.Ip.ValueString())
		}
		if plan.Netmask.ValueStringPointer() != nil {
			data.Netmask = models.ToPointer(plan.Netmask.ValueString())
		}
		if plan.Network.ValueStringPointer() != nil {
			data.Network = models.ToPointer(plan.Network.ValueString())
		}
		if plan.PoserPassword.ValueStringPointer() != nil {
			data.PoserPassword = models.ToPointer(plan.PoserPassword.ValueString())
		}
		if plan.PppoeUsername.ValueStringPointer() != nil {
			data.PppoeUsername = models.ToPointer(plan.PppoeUsername.ValueString())
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

func portConfigWanSourceNatTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.GatewayPortWanSourceNat {
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
	data_map := make(map[string]models.GatewayPortConfig)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(PortConfigValue)
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
			data.Description = models.ToPointer(plan.Description.ValueString())
		}
		if plan.DisableAutoneg.ValueBoolPointer() != nil {
			data.DisableAutoneg = models.ToPointer(plan.DisableAutoneg.ValueBool())
		}
		if plan.Disabled.ValueBoolPointer() != nil {
			data.Disabled = models.ToPointer(plan.Disabled.ValueBool())
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
		data.IpConfig = gatewayIpConfigTerraformToSdk(ctx, diags, t)

		if plan.LteApn.ValueStringPointer() != nil {
			data.LteApn = models.ToPointer(plan.LteApn.ValueString())
		}
		if plan.LteAuth.ValueStringPointer() != nil {
			data.LteAuth = models.ToPointer(models.GatewayPortLteAuthEnum(plan.LteAuth.ValueString()))
		}
		if plan.LteBackup.ValueBoolPointer() != nil {
			data.LteBackup = models.ToPointer(plan.LteBackup.ValueBool())
		}
		if plan.LtePassword.ValueStringPointer() != nil {
			data.LtePassword = models.ToPointer(plan.LtePassword.ValueString())
		}
		if plan.LteUsername.ValueStringPointer() != nil {
			data.LteUsername = models.ToPointer(plan.LteUsername.ValueString())
		}
		if plan.Mtu.ValueInt64Pointer() != nil {
			data.Mtu = models.ToPointer(int(plan.Mtu.ValueInt64()))
		}
		if plan.Name.ValueStringPointer() != nil {
			data.Name = models.ToPointer(plan.Name.ValueString())
		}
		if plan.Name.IsNull() && !plan.Name.IsUnknown() {
			data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		}
		if plan.OuterVlanId.ValueInt64Pointer() != nil {
			data.OuterVlanId = models.ToPointer(int(plan.OuterVlanId.ValueInt64()))
		}
		if plan.PoeDisabled.ValueBoolPointer() != nil {
			data.PoeDisabled = models.ToPointer(plan.PoeDisabled.ValueBool())
		}
		if plan.PortNetwork.ValueStringPointer() != nil {
			data.PortNetwork = models.ToPointer(plan.PortNetwork.ValueString())
		}
		if plan.PreserveDscp.ValueBoolPointer() != nil {
			data.PreserveDscp = models.ToPointer(plan.PreserveDscp.ValueBool())
		}
		if plan.Redundant.ValueBoolPointer() != nil {
			data.Redundant = models.ToPointer(plan.Redundant.ValueBool())
		}
		if plan.RethIdx.ValueInt64Pointer() != nil {
			data.RethIdx = models.ToPointer(int(plan.RethIdx.ValueInt64()))
		}
		if plan.RethNode.ValueStringPointer() != nil {
			data.RethNode = models.ToPointer(plan.RethNode.ValueString())
		}
		if plan.Speed.ValueStringPointer() != nil {
			data.Speed = models.ToPointer(plan.Speed.ValueString())
		}
		if plan.SsrNoVirtualMac.ValueBoolPointer() != nil {
			data.SsrNoVirtualMac = models.ToPointer(plan.SsrNoVirtualMac.ValueBool())
		}
		if plan.SvrPortRange.ValueStringPointer() != nil {
			data.SvrPortRange = models.ToPointer(plan.SvrPortRange.ValueString())
		}

		data.TrafficShaping = gatewayPortTrafficShapingTerraformToSdk(ctx, diags, plan.TrafficShaping)

		if plan.VlanId.ValueInt64Pointer() != nil {
			data.VlanId = models.ToPointer(int(plan.VlanId.ValueInt64()))
		}

		data.VpnPaths = gatewayPortVpnPathTerraformToSdk(ctx, diags, plan.VpnPaths)

		data.WanArpPolicer = models.ToPointer(models.GatewayPortWanArpPolicerEnum(plan.WanArpPolicer.ValueString()))

		if plan.WanExtIp.ValueStringPointer() != nil {
			data.WanExtIp = models.ToPointer(plan.WanExtIp.ValueString())
		}

		data.WanExtraRoutes = wanExtraRoutesPortVpnPathTerraformToSdk(ctx, diags, plan.WanExtraRoutes)

		if !plan.WanNetworks.IsNull() && !plan.WanNetworks.IsUnknown() {
			data.WanNetworks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.WanNetworks)
		}

		data.WanProbeOverride = wanProbeOverridePortVpnPathTerraformToSdk(ctx, diags, plan.WanProbeOverride)
		data.WanSourceNat = portConfigWanSourceNatTerraformToSdk(ctx, diags, plan.WanSourceNat)

		if plan.WanType.ValueStringPointer() != nil {
			data.WanType = models.ToPointer(models.GatewayPortWanTypeEnum(plan.WanType.ValueString()))
		}
		data_map[k] = data
	}
	return data_map
}
