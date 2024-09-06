package resource_org_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data models.NetworkTemplate) (OrgNetworktemplateModel, diag.Diagnostics) {
	var state OrgNetworktemplateModel
	var diags diag.Diagnostics

	var acl_policies types.List = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var acl_tags types.Map = types.MapNull(AclTagsValue{}.Type(ctx))
	var additional_config_cmds types.List = types.ListNull(types.StringType)
	var dhcp_snooping DhcpSnoopingValue = NewDhcpSnoopingValueNull()
	var dns_servers types.List = types.ListNull(types.StringType)
	var dns_suffix types.List = types.ListNull(types.StringType)
	var extra_routes types.Map = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extra_routes6 types.Map = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var mist_nac MistNacValue = NewMistNacValueNull()
	var name types.String
	var networks types.Map = types.MapNull(NetworksValue{}.Type(ctx))
	var ntp_servers types.List = types.ListNull(types.StringType)
	var org_id types.String
	var port_mirroring types.Map = types.MapNull(PortMirroringValue{}.Type(ctx))
	var port_usages types.Map = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radius_config RadiusConfigValue = NewRadiusConfigValueNull()
	var remote_syslog RemoteSyslogValue = NewRemoteSyslogValueNull()
	var remove_existing_configs types.Bool = types.BoolValue(false)
	var snmp_config SnmpConfigValue = NewSnmpConfigValueNull()
	var switch_matching SwitchMatchingValue = NewSwitchMatchingValueNull()
	var switch_mgmt SwitchMgmtValue = NewSwitchMgmtValueNull()
	var vrf_config VrfConfigValue = NewVrfConfigValueNull()
	var vrf_instances types.Map = types.MapNull(VrfInstancesValue{}.Type(ctx))

	if data.AclPolicies != nil {
		acl_policies = aclPoliciesSdkToTerraform(ctx, &diags, data.AclPolicies)
	}
	if data.AclTags != nil && len(data.AclTags) > 0 {
		acl_tags = aclTagsSdkToTerraform(ctx, &diags, data.AclTags)
	}
	if data.AdditionalConfigCmds != nil {
		additional_config_cmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalConfigCmds)
	}
	if data.DhcpSnooping != nil {
		dhcp_snooping = dhcpSnoopingSdkToTerraform(ctx, &diags, data.DhcpSnooping)
	}
	if data.DnsServers != nil {
		dns_servers = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsServers)
	}
	if data.DnsSuffix != nil {
		dns_suffix = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsSuffix)
	}
	if data.ExtraRoutes != nil && len(data.ExtraRoutes) > 0 {
		extra_routes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if data.ExtraRoutes6 != nil && len(data.ExtraRoutes6) > 0 {
		extra_routes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.MistNac != nil {
		mist_nac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Networks != nil && len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntp_servers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.PortMirroring != nil {
		port_mirroring = portMirroringSdkToTerraform(ctx, &diags, data.PortMirroring)
	}
	if data.PortUsages != nil && len(data.PortUsages) > 0 {
		port_usages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)
	}
	if data.RadiusConfig != nil {
		radius_config = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)
	}
	if data.RemoteSyslog != nil {
		remote_syslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
	}
	if data.RemoveExistingConfigs != nil {
		state.RemoveExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
	}
	if data.SnmpConfig != nil {
		snmp_config = snmpConfigSdkToTerraform(ctx, &diags, data.SnmpConfig)
	}
	if data.SwitchMatching != nil {
		switch_matching = switchMatchingSdkToTerraform(ctx, &diags, data.SwitchMatching)
	}
	if data.SwitchMgmt != nil {
		switch_mgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.VrfConfig != nil {
		vrf_config = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if data.VrfInstances != nil && len(data.VrfInstances) > 0 {
		vrf_instances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}

	state.Id = id
	state.OrgId = org_id
	state.Name = name
	state.AclPolicies = acl_policies
	state.AclTags = acl_tags
	state.AdditionalConfigCmds = additional_config_cmds
	state.DhcpSnooping = dhcp_snooping
	state.DnsServers = dns_servers
	state.DnsSuffix = dns_suffix
	state.ExtraRoutes = extra_routes
	state.ExtraRoutes6 = extra_routes6
	state.MistNac = mist_nac
	state.NtpServers = ntp_servers
	state.Networks = networks
	state.PortMirroring = port_mirroring
	state.PortUsages = port_usages
	state.RadiusConfig = radius_config
	state.RemoteSyslog = remote_syslog
	state.RemoveExistingConfigs = remove_existing_configs
	state.SnmpConfig = snmp_config
	state.SwitchMatching = switch_matching
	state.SwitchMgmt = switch_mgmt
	state.VrfConfig = vrf_config
	state.VrfInstances = vrf_instances

	return state, diags
}
