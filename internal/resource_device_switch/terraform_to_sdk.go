package resource_device_switch

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *DeviceSwitchModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceSwitch{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if plan.AclPolicies.IsNull() || plan.AclPolicies.IsUnknown() {
		unset["-acl_policies"] = ""
	} else {
		data.AclPolicies = aclPoliciesTerraformToSdk(plan.AclPolicies)
	}

	if plan.AclTags.IsNull() || plan.AclTags.IsUnknown() {
		unset["-acl_tags"] = ""
	} else {
		data.AclTags = actTagsTerraformToSdk(plan.AclTags)
	}

	if plan.AdditionalConfigCmds.IsNull() || plan.AdditionalConfigCmds.IsUnknown() {
		unset["-additional_config_cmds"] = ""
	} else {
		data.AdditionalConfigCmds = mistutils.ListOfStringTerraformToSdk(plan.AdditionalConfigCmds)
	}

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		data.BgpConfig = bgpConfigTerraformToSdk(plan.BgpConfig)
	}

	if plan.DhcpSnooping.IsNull() || plan.DhcpSnooping.IsUnknown() {
		unset["-dhcp_snooping"] = ""
	} else {
		data.DhcpSnooping = dhcpSnoopingTerraformToSdk(plan.DhcpSnooping)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		data.DhcpdConfig = dhcpdConfigTerraformToSdk(plan.DhcpdConfig)
	}

	if plan.DisableAutoConfig.IsNull() || plan.DisableAutoConfig.IsUnknown() {
		unset["-disable_auto_config"] = ""
	} else {
		data.DisableAutoConfig = plan.DisableAutoConfig.ValueBoolPointer()
	}

	if plan.DnsServers.IsNull() || plan.DnsServers.IsUnknown() {
		unset["-dns_servers"] = ""
	} else {
		data.DnsServers = mistutils.ListOfStringTerraformToSdk(plan.DnsServers)
	}

	if plan.DnsSuffix.IsNull() || plan.DnsSuffix.IsUnknown() {
		unset["-dns_suffix"] = ""
	} else {
		data.DnsSuffix = mistutils.ListOfStringTerraformToSdk(plan.DnsSuffix)
	}

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		data.ExtraRoutes = extraRoutesTerraformToSdk(plan.ExtraRoutes)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutes6TerraformToSdk(plan.ExtraRoutes6)
	}

	if plan.IpConfig.IsNull() || plan.IpConfig.IsUnknown() {
		unset["-ip_config"] = ""
	} else {
		data.IpConfig = ipConfigTerraformToSdk(plan.IpConfig)
	}

	if plan.LocalPortConfig.IsNull() || plan.LocalPortConfig.IsUnknown() {
		unset["-local_port_config"] = ""
	} else {
		data.LocalPortConfig = LocalPortConfigTerraformToSdk(ctx, &diags, plan.LocalPortConfig)
	}

	if plan.Managed.IsNull() || plan.Managed.IsUnknown() {
		unset["-managed"] = ""
	} else {
		data.Managed = plan.Managed.ValueBoolPointer()
	}

	if len(plan.MapId.ValueString()) > 0 {
		mapId, e := uuid.Parse(plan.MapId.ValueString())
		if e == nil {
			data.MapId = &mapId
		} else {
			diags.AddError("Bad value for map_id", e.Error())
		}
	} else {
		unset["-map_id"] = ""
	}

	data.Name = models.ToPointer(plan.Name.ValueString())

	if plan.Notes.IsNull() || plan.Notes.IsUnknown() {
		unset["-notes"] = ""
	} else {
		data.Notes = plan.Notes.ValueStringPointer()
	}

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mistNac := mistNacTerraformToSdk(plan.MistNac)
		data.MistNac = mistNac
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := NetworksTerraformToSdk(plan.Networks)
		data.Networks = networks
	}

	if plan.Notes.IsNull() || plan.Notes.IsUnknown() {
		unset["-notes"] = ""
	} else {
		data.Notes = plan.Notes.ValueStringPointer()
	}

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = mistutils.ListOfStringTerraformToSdk(plan.NtpServers)
	}

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		data.OobIpConfig = oobIpConfigTerraformToSdk(plan.OobIpConfig)
	}

	if plan.OspfAreas.IsNull() || plan.OspfAreas.IsUnknown() {
		unset["-ospf_areas"] = ""
	} else {
		data.OspfAreas = ospfAreasTerraformToSdk(plan.OspfAreas)
	}

	if plan.OspfConfig.IsNull() || plan.OspfConfig.IsUnknown() {
		unset["-ospf_config"] = ""
	} else {
		data.OspfConfig = ospfConfigTerraformToSdk(plan.OspfConfig)
	}

	if plan.OtherIpConfigs.IsNull() || plan.OtherIpConfigs.IsUnknown() {
		unset["-other_ip_configs"] = ""
	} else {
		data.OtherIpConfigs = otherIpConfigTerraformToSdk(plan.OtherIpConfigs)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		data.PortConfig = portConfigTerraformToSdk(plan.PortConfig)
	}

	if plan.PortConfigOverwrite.IsNull() || plan.PortConfigOverwrite.IsUnknown() {
		unset["-port_config_overwrite"] = ""
	} else {
		data.PortConfigOverwrite = portConfigOverwriteTerraformToSdk(plan.PortConfigOverwrite)
	}

	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirroring"] = ""
	} else {
		data.PortMirroring = portMirroringTerraformToSdk(plan.PortMirroring)
	}

	if plan.PortUsages.IsNull() || plan.PortUsages.IsUnknown() {
		unset["-port_usages"] = ""
	} else {
		data.PortUsages = portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
	}

	if plan.RadiusConfig.IsNull() || plan.RadiusConfig.IsUnknown() {
		unset["-radius_config"] = ""
	} else {
		data.RadiusConfig = radiusConfigTerraformToSdk(plan.RadiusConfig)
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
		data.StpConfig = stpConfigTerraformToSdk(plan.StpConfig)
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
		data.Vars = varsTerraformToSdk(plan.Vars)
	}

	if plan.VirtualChassis.IsNull() || plan.VirtualChassis.IsUnknown() {
		unset["-virtual_chassis"] = ""
	} else {
		data.VirtualChassis = virtualChassisTerraformToSdk(plan.VirtualChassis)
	}

	if plan.VrfConfig.IsNull() || plan.VrfConfig.IsUnknown() {
		unset["-vrf_config"] = ""
	} else {
		data.VrfConfig = vrfConfigTerraformToSdk(plan.VrfConfig)
	}

	if plan.VrfInstances.IsNull() || plan.VrfInstances.IsUnknown() {
		unset["-vrf_instances"] = ""
	} else {
		data.VrfInstances = vrfInstancesTerraformToSdk(plan.VrfInstances)
	}

	if plan.VrrpConfig.IsNull() || plan.VrrpConfig.IsUnknown() {
		unset["-vrrp_config"] = ""
	} else {
		data.VrrpConfig = vrrpTerraformToSdk(plan.VrrpConfig)
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

	mistDevice := models.MistDeviceContainer.FromDeviceSwitch(data)
	return mistDevice, diags
}
