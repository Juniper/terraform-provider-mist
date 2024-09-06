package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *OrgNetworktemplateModel) (models.NetworkTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	unset := make(map[string]interface{})

	data := models.NetworkTemplate{}

	data.Name = models.ToPointer(plan.Name.ValueString())

	if plan.AclPolicies.IsNull() || plan.AclPolicies.IsUnknown() {
		unset["-acl_policies"] = ""
	} else {
		acl_policies := aclPoliciesTerraformToSdk(ctx, &diags, plan.AclPolicies)
		data.AclPolicies = acl_policies
	}

	if plan.AclTags.IsNull() || plan.AclTags.IsUnknown() {
		unset["-acl_tags"] = ""
	} else {
		acl_tags := actTagsTerraformToSdk(ctx, &diags, plan.AclTags)
		data.AclTags = acl_tags
	}

	if plan.AdditionalConfigCmds.IsNull() || plan.AdditionalConfigCmds.IsUnknown() {
		unset["-additional_config_cmds"] = ""
	} else {
		data.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	}

	if plan.DnsServers.IsNull() || plan.DnsServers.IsUnknown() {
		unset["-dns_servers"] = ""
	} else {
		data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)
	}

	if plan.DnsSuffix.IsNull() || plan.DnsSuffix.IsUnknown() {
		unset["-dns_suffix"] = ""
	} else {
		data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)
	}

	if plan.DhcpSnooping.IsNull() || plan.DhcpSnooping.IsUnknown() {
		unset["-dhcp_snooping"] = ""
	} else {
		dhcp_snooping := dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
		data.DhcpSnooping = dhcp_snooping
	}

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extra_routes := extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
		data.ExtraRoutes = extra_routes
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		extra_routes6 := extraRoutes6TerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
		data.ExtraRoutes6 = extra_routes6
	}

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

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	}

	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirroring"] = ""
	} else {
		port_mirroring := portMirroringTerraformToSdk(ctx, &diags, plan.PortMirroring)
		data.PortMirroring = port_mirroring
	}

	if plan.PortUsages.IsNull() || plan.PortUsages.IsUnknown() {
		unset["-port_usages"] = ""
	} else {
		port_usages := portUsageTerraformToSdk(ctx, &diags, plan.PortUsages)
		data.PortUsages = port_usages
	}

	if plan.RadiusConfig.IsNull() || plan.RadiusConfig.IsUnknown() {
		unset["-radius_config"] = ""
	} else {
		radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
		data.RadiusConfig = radius_config
	}

	if plan.RemoteSyslog.IsNull() || plan.RemoteSyslog.IsUnknown() {
		unset["-remote_syslog"] = ""
	} else {
		remote_syslog := remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
		data.RemoteSyslog = remote_syslog
	}

	if plan.SnmpConfig.IsNull() || plan.SnmpConfig.IsUnknown() {
		unset["-snmp_config"] = ""
	} else {
		snmp_config := snmpConfigTerraformToSdk(ctx, &diags, plan.SnmpConfig)
		data.SnmpConfig = snmp_config
	}

	if plan.SwitchMatching.IsNull() || plan.SwitchMatching.IsUnknown() {
		unset["-switch_matching"] = ""
	} else {
		switch_matching := switchMatchingTerraformToSdk(ctx, &diags, plan.SwitchMatching)
		data.SwitchMatching = switch_matching
	}

	if plan.SwitchMgmt.IsNull() || plan.SwitchMgmt.IsUnknown() {
		unset["-switch_mgmt"] = ""
	} else {
		switch_mgmt := switchMgmtTerraformToSdk(ctx, &diags, plan.SwitchMgmt)
		data.SwitchMgmt = switch_mgmt
	}

	if plan.VrfConfig.IsNull() || plan.VrfConfig.IsUnknown() {
		unset["-vrf_config"] = ""
	} else {
		vrf_config := vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
		data.VrfConfig = vrf_config
	}

	if plan.VrfInstances.IsNull() || plan.VrfInstances.IsUnknown() {
		unset["-vrf_instances"] = ""
	} else {
		vrf_instances := vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
		data.VrfInstances = vrf_instances
	}

	data.AdditionalProperties = unset

	return data, diags
}
