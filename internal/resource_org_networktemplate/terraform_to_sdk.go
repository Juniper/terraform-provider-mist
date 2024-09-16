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
		data.AclPolicies = aclPoliciesTerraformToSdk(ctx, &diags, plan.AclPolicies)
	}

	if plan.AclTags.IsNull() || plan.AclTags.IsUnknown() {
		unset["-acl_tags"] = ""
	} else {
		data.AclTags = actTagsTerraformToSdk(ctx, &diags, plan.AclTags)
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
		data.DhcpSnooping = dhcpSnoopingTerraformToSdk(ctx, &diags, plan.DhcpSnooping)
	}

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		data.ExtraRoutes = extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutes6TerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
	}

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		data.MistNac = mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		data.Networks = NetworksTerraformToSdk(ctx, &diags, plan.Networks)
	}

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)
	}

	if plan.OspfAreas.IsNull() || plan.OspfAreas.IsUnknown() {
		unset["-ospf_areas"] = ""
	} else {
		data.OspfAreas = ospfAreasTerraformToSdk(ctx, &diags, plan.OspfAreas)
	}

	if plan.PortMirroring.IsNull() || plan.PortMirroring.IsUnknown() {
		unset["-port_mirroring"] = ""
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
		radius_config := radiusConfigTerraformToSdk(ctx, &diags, plan.RadiusConfig)
		data.RadiusConfig = radius_config
	}

	if plan.RemoteSyslog.IsNull() || plan.RemoteSyslog.IsUnknown() {
		unset["-remote_syslog"] = ""
	} else {
		data.RemoteSyslog = remoteSyslogTerraformToSdk(ctx, &diags, plan.RemoteSyslog)
	}

	if plan.RemoveExistingConfigs.IsNull() || plan.RemoveExistingConfigs.IsUnknown() {
		unset["-remove_existing_configs"] = ""
	} else {
		data.RemoveExistingConfigs = plan.RemoveExistingConfigs.ValueBoolPointer()
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
		data.VrfConfig = vrfConfigTerraformToSdk(ctx, &diags, plan.VrfConfig)
	}

	if plan.VrfInstances.IsNull() || plan.VrfInstances.IsUnknown() {
		unset["-vrf_instances"] = ""
	} else {
		data.VrfInstances = vrfInstancesTerraformToSdk(ctx, &diags, plan.VrfInstances)
	}

	data.AdditionalProperties = unset

	return data, diags
}
