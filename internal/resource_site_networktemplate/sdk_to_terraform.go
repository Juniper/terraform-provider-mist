package resource_site_networktemplate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteNetworktemplateModel, diag.Diagnostics) {
	var state SiteNetworktemplateModel
	var diags diag.Diagnostics

	var acl_policies types.List = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var acl_tags types.Map = types.MapNull(AclTagsValue{}.Type(ctx))
	var additional_config_cmds types.List = types.ListNull(types.StringType)
	var dhcp_snooping DhcpSnoopingValue = NewDhcpSnoopingValueNull()
	var dns_servers types.List = types.ListNull(types.StringType)
	var dns_suffix types.List = types.ListNull(types.StringType)
	var disabled_system_defined_port_usages types.List = types.ListNull(types.StringType)
	var extra_routes types.Map = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extra_routes6 types.Map = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var mist_nac MistNacValue = NewMistNacValueNull()
	var networks types.Map = types.MapNull(NetworksValue{}.Type(ctx))
	var ntp_servers types.List = types.ListNull(types.StringType)
	var ospf_areas types.Map = types.MapNull(OspfAreasValue{}.Type(ctx))
	var port_mirroring types.Map = types.MapNull(PortMirroringValue{}.Type(ctx))
	var port_usages types.Map = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radius_config RadiusConfigValue = NewRadiusConfigValueNull()
	var remote_syslog RemoteSyslogValue = NewRemoteSyslogValueNull()
	var remove_existing_configs types.Bool = types.BoolValue(false)
	var snmp_config SnmpConfigValue = NewSnmpConfigValueNull()
	var siteId types.String = types.StringValue(data.SiteId.String())
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
	if data.DisabledSystemDefinedPortUsages != nil {
		var items []attr.Value
		var items_type attr.Type = basetypes.StringType{}
		for _, item := range data.DisabledSystemDefinedPortUsages {
			items = append(items, types.StringValue(string(item)))
		}
		list, _ := types.ListValue(items_type, items)
		disabled_system_defined_port_usages = list
	}
	if data.ExtraRoutes != nil && len(data.ExtraRoutes) > 0 {
		extra_routes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if data.ExtraRoutes6 != nil && len(data.ExtraRoutes6) > 0 {
		extra_routes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.MistNac != nil {
		mist_nac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Networks != nil && len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntp_servers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)
	}
	if data.OspfAreas != nil {
		ospf_areas = ospfAreasSdkToTerraform(ctx, &diags, data.OspfAreas)
	}
	if data.PortMirroring != nil && len(data.PortMirroring) > 0 {
		port_mirroring = portMirroringSdkToTerraform(ctx, &diags, data.PortMirroring)
	}
	if data.PortUsages != nil && len(data.PortUsages) > 0 {
		port_usages = portUsagesSdkToTerraform(ctx, &diags, data.PortUsages)
	}
	if data.RadiusConfig != nil {
		radius_config = radiusConfigSdkToTerraform(ctx, &diags, data.RadiusConfig)
	}
	if data.RemoveExistingConfigs != nil {
		remove_existing_configs = types.BoolValue(*data.RemoveExistingConfigs)
	}
	if data.RemoteSyslog != nil {
		remote_syslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
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
	if data.VrfInstances != nil {
		vrf_instances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}

	state.AclPolicies = acl_policies
	state.AclTags = acl_tags
	state.AdditionalConfigCmds = additional_config_cmds
	state.DhcpSnooping = dhcp_snooping
	state.DnsServers = dns_servers
	state.DnsSuffix = dns_suffix
	state.DisabledSystemDefinedPortUsages = disabled_system_defined_port_usages
	state.ExtraRoutes = extra_routes
	state.ExtraRoutes6 = extra_routes6
	state.MistNac = mist_nac
	state.NtpServers = ntp_servers
	state.Networks = networks
	state.OspfAreas = ospf_areas
	state.PortMirroring = port_mirroring
	state.PortUsages = port_usages
	state.RadiusConfig = radius_config
	state.RemoteSyslog = remote_syslog
	state.RemoveExistingConfigs = remove_existing_configs
	state.SnmpConfig = snmp_config
	state.SiteId = siteId
	state.SwitchMatching = switch_matching
	state.SwitchMgmt = switch_mgmt
	state.VrfConfig = vrf_config
	state.VrfInstances = vrf_instances

	return state, diags
}
