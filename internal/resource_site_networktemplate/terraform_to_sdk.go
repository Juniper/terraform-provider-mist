package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *SiteNetworktemplateModel) (*models.SiteSetting, diag.Diagnostics) {
	var diags diag.Diagnostics

	unset := make(map[string]interface{})

	data := models.SiteSetting{}

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

	if plan.AutoUpgradeLinecard.IsNull() || plan.AutoUpgradeLinecard.IsUnknown() {
		unset["-auto_upgrade_linecard"] = ""
	} else {
		data.AutoUpgradeLinecard = plan.AutoUpgradeLinecard.ValueBoolPointer()
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

	if plan.DhcpSnooping.IsNull() || plan.DhcpSnooping.IsUnknown() {
		unset["-dhcp_snooping"] = ""
	} else {
		data.DhcpSnooping = dhcpSnoopingTerraformToSdk(plan.DhcpSnooping)
	}

	if plan.DisabledSystemDefinedPortUsages.IsNull() || plan.DisabledSystemDefinedPortUsages.IsUnknown() {
		unset["-disabled_system_defined_port_usages"] = ""
	} else {
		var items []models.SystemDefinedPortUsagesEnum
		for _, item := range plan.DisabledSystemDefinedPortUsages.Elements() {
			var sInterface interface{} = item
			s := sInterface.(basetypes.StringValue)
			items = append(items, models.SystemDefinedPortUsagesEnum(s.ValueString()))
		}
		data.DisabledSystemDefinedPortUsages = items
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

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		data.MistNac = mistNacTerraformToSdk(plan.MistNac)
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		data.Networks = NetworksTerraformToSdk(plan.Networks)
	}

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = mistutils.ListOfStringTerraformToSdk(plan.NtpServers)
	}

	if plan.OspfAreas.IsNull() || plan.OspfAreas.IsUnknown() {
		unset["-ospf_areas"] = ""
	} else {
		data.OspfAreas = ospfAreasTerraformToSdk(plan.OspfAreas)
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

	if plan.RoutingPolicies.IsNull() || plan.RoutingPolicies.IsUnknown() {
		unset["-routing_policies"] = ""
	} else {
		data.RoutingPolicies = routingPoliciesTerraformToSdk(ctx, plan.RoutingPolicies)
	}

	if plan.SnmpConfig.IsNull() || plan.SnmpConfig.IsUnknown() {
		unset["-snmp_config"] = ""
	} else {
		data.SnmpConfig = snmpConfigTerraformToSdk(ctx, &diags, plan.SnmpConfig)
	}

	if plan.SwitchMatching.IsNull() || plan.SwitchMatching.IsUnknown() {
		unset["-switch_matching"] = ""
	} else {
		data.SwitchMatching = switchMatchingTerraformToSdk(ctx, &diags, plan.SwitchMatching)
	}

	if plan.SwitchMgmt.IsNull() || plan.SwitchMgmt.IsUnknown() {
		unset["-switch_mgmt"] = ""
	} else {
		data.SwitchMgmt = switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
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

	data.AdditionalProperties = unset

	return &data, diags
}
