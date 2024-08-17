package resource_device_switch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *DeviceSwitchModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceSwitch{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if plan.AclPolicies.IsNull() || plan.AclPolicies.IsUnknown() {
		unset["-acl_policies"] = ""
	} else {
		data.AclPolicies = aclPoliciesTerraformToSdk(ctx, &diags, plan.AclPolicies)
	}

	if plan.AclTags.IsNull() || plan.AclTags.IsUnknown() {
		unset["-acl_tags"] = ""
	} else {
		data.AclTags = actTagsTerraformToSdk(ctx, &diags, plan.AclTags)
	}

	data.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)

	if plan.DhcpSnooping.IsNull() || plan.DhcpSnooping.IsUnknown() {
		unset["-dhcp_snooping"] = ""
	} else {
		data.DhcpSnooping = dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		data.DhcpdConfig = dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
	}

	if plan.DisableAutoConfig.IsNull() || plan.DisableAutoConfig.IsUnknown() {
		unset["-disable_auto_config"] = ""
	} else {
		data.DisableAutoConfig = plan.DisableAutoConfig.ValueBoolPointer()
	}

	data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)

	data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		data.ExtraRoutes = extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	}

	if plan.EvpnConfig.IsNull() || plan.EvpnConfig.IsUnknown() {
		unset["-evpn_config"] = ""
	} else {
		data.EvpnConfig = evpnConfigTerraformToSdk(ctx, &diags, plan.EvpnConfig)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutes6TerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
	}

	if plan.IpConfig.IsNull() || plan.IpConfig.IsUnknown() {
		unset["-ip_config"] = ""
	} else {
		data.IpConfig = ipConfigTerraformToSdk(ctx, &diags, plan.IpConfig)
	}

	if plan.Managed.IsNull() || plan.Managed.IsUnknown() {
		unset["-managed"] = ""
	} else {
		data.Managed = plan.Managed.ValueBoolPointer()
	}

	if len(plan.MapId.ValueString()) > 0 {
		map_id, e := uuid.Parse(plan.MapId.ValueString())
		if e == nil {
			data.MapId = &map_id
		} else {
			diags.AddError("Bad value for map_id", e.Error())
		}
	} else {
		unset["-map_id"] = nil
	}

	data.Name = models.ToPointer(plan.Name.ValueString())

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mist_nac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
		data.MistNac = mist_nac
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := NetworksTerraformToSdk(ctx, &diags, plan.Networks)
		data.Networks = networks
	}

	data.Notes = plan.Notes.ValueStringPointer()

	data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		data.OobIpConfig = oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
	}

	if plan.OspfConfig.IsNull() || plan.OspfConfig.IsUnknown() {
		unset["-ospf_config"] = ""
	} else {
		data.OspfConfig = ospfConfigTerraformToSdk(ctx, &diags, plan.OspfConfig)
	}

	if plan.OtherIpConfigs.IsNull() || plan.OtherIpConfigs.IsUnknown() {
		unset["-other_ip_configs"] = ""
	} else {
		data.OtherIpConfigs = otherIpConfigTerraformToSdk(ctx, &diags, plan.OtherIpConfigs)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		data.PortConfig = portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
	}

	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirrorings"] = ""
	} else {
		data.PortMirroring = portMirroringTerraformToSdk(ctx, &diags, plan.PortMirroring)
	}

	if plan.PortUsages.IsNull() || plan.PortUsages.IsUnknown() {
		unset["-port_usages"] = ""
	} else {
		data.PortUsages = portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	}

	if plan.RadiusConfig.IsNull() || plan.RadiusConfig.IsUnknown() {
		unset["-radius_config"] = ""
	} else {
		data.RadiusConfig = radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
	}

	if plan.RemoteSyslog.IsNull() || plan.RemoteSyslog.IsUnknown() {
		unset["-remote_syslog"] = ""
	} else {
		data.RemoteSyslog = remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
	}

	if plan.RouterId.ValueStringPointer() == nil {
		unset["-router_id"] = ""
	} else {
		data.RouterId = plan.RouterId.ValueStringPointer()
	}

	if plan.Role.ValueStringPointer() == nil {
		unset["-role"] = ""
	} else {
		data.Role = plan.Role.ValueStringPointer()
	}

	if plan.SnmpConfig.IsNull() || plan.SnmpConfig.IsUnknown() {
		unset["-snmp_config"] = ""
	} else {
		data.SnmpConfig = snmpConfigTerraformToSdk(ctx, &diags, plan.SnmpConfig)
	}

	if plan.StpConfig.IsNull() || plan.StpConfig.IsUnknown() {
		unset["-stp_config"] = ""
	} else {
		data.StpConfig = stpConfigTerraformToSdk(ctx, &diags, plan.StpConfig)
	}

	if plan.SwitchMgmt.IsNull() || plan.SwitchMgmt.IsUnknown() {
		unset["-switch_mgmt"] = ""
	} else {
		data.SwitchMgmt = switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
	}

	if plan.UseRouterIdAsSourceIp.ValueBoolPointer() == nil {
		unset["-use_router_id_as_source_ip"] = ""
	} else {
		data.UseRouterIdAsSourceIp = plan.UseRouterIdAsSourceIp.ValueBoolPointer()
	}

	if plan.Vars.IsNull() || plan.Vars.IsUnknown() {
		unset["-vars"] = ""
	} else {
		data.Vars = varsTerraformToSdk(ctx, &diags, plan.Vars)
	}

	if plan.VirtualChassis.IsNull() || plan.VirtualChassis.IsUnknown() {
		unset["-virtual_chassis"] = ""
	} else {
		data.VirtualChassis = virtualChassisTerraformToSdk(ctx, &diags, plan.VirtualChassis)
	}

	if plan.VrfConfig.IsNull() || plan.VrfConfig.IsUnknown() {
		unset["-vrf_config"] = ""
	} else {
		data.VrfConfig = vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
	}

	if plan.VrfInstances.IsNull() || plan.VrfInstances.IsUnknown() {
		unset["-vrf_instances"] = ""
	} else {
		data.VrfInstances = vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
	}

	if plan.VrrpConfig.IsNull() || plan.VrrpConfig.IsUnknown() {
		unset["-vrrp_config"] = ""
	} else {
		data.VrrpConfig = vrrpTerraformToSdk(ctx, &diags, plan.VrrpConfig)
	}

	if !plan.X.IsNull() && !plan.X.IsUnknown() {
		data.X = plan.X.ValueFloat64Pointer()
	} else {
		unset["-x"] = ""
	}
	if !plan.Y.IsNull() && !plan.Y.IsUnknown() {
		data.Y = plan.Y.ValueFloat64Pointer()
	} else {
		unset["-y"] = ""
	}

	data.AdditionalProperties = unset

	mist_device := models.MistDeviceContainer.FromDeviceSwitch(data)
	return mist_device, diags
}
