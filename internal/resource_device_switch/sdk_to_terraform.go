package resource_device_switch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceSwitch) (DeviceSwitchModel, diag.Diagnostics) {
	var state DeviceSwitchModel
	var diags diag.Diagnostics

	var acl_policies types.List = types.ListNull(AclPoliciesValue{}.Type(ctx))
	var acl_tags types.Map = types.MapNull(AclTagsValue{}.Type(ctx))
	var additional_config_cmds types.List = types.ListNull(types.StringType)
	var dhcp_snooping DhcpSnoopingValue = NewDhcpSnoopingValueNull()
	var dhcpd_config DhcpdConfigValue = NewDhcpdConfigValueNull()
	var device_id types.String
	var disable_auto_config types.Bool
	var dns_servers types.List = types.ListNull(types.StringType)
	var dns_suffix types.List = types.ListNull(types.StringType)
	var extra_routes types.Map = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extra_routes6 types.Map = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var image1_url types.String = types.StringValue("not_present")
	var image2_url types.String = types.StringValue("not_present")
	var image3_url types.String = types.StringValue("not_present")
	var ip_config IpConfigValue = NewIpConfigValueNull()
	var local_port_config types.Map = types.MapNull(LocalPortConfigValue{}.Type(ctx))
	var managed types.Bool
	var map_id types.String
	var mist_nac MistNacValue = NewMistNacValueNull()
	var name types.String
	var notes types.String
	var networks types.Map = types.MapNull(NetworksValue{}.Type(ctx))
	var ntp_servers types.List = types.ListNull(types.StringType)
	var oob_ip_config OobIpConfigValue = NewOobIpConfigValueNull()
	var ospf_areas types.Map = types.MapNull(OspfAreasValue{}.Type(ctx))
	var other_ip_configs types.Map = types.MapNull(OtherIpConfigsValue{}.Type(ctx))
	var org_id types.String
	var port_config types.Map = types.MapNull(PortConfigValue{}.Type(ctx))
	var port_mirroring types.Map = types.MapNull(PortMirroringValue{}.Type(ctx))
	var port_usages types.Map = types.MapNull(PortUsagesValue{}.Type(ctx))
	var radius_config RadiusConfigValue = NewRadiusConfigValueNull()
	var remote_syslog RemoteSyslogValue = NewRemoteSyslogValueNull()
	var role types.String
	var router_id types.String
	var site_id types.String
	var snmp_config SnmpConfigValue = NewSnmpConfigValueNull()
	var stp_config StpConfigValue = NewStpConfigValueNull()
	var switch_mgmt SwitchMgmtValue = NewSwitchMgmtValueNull()
	var use_router_id_as_source_ip types.Bool
	var vars types.Map = types.MapNull(types.StringType)
	var virtual_chassis VirtualChassisValue = NewVirtualChassisValueNull()
	var vrf_config VrfConfigValue = NewVrfConfigValueNull()
	var vrf_instances types.Map = types.MapNull(VrfInstancesValue{}.Type(ctx))
	var vrrp_config VrrpConfigValue = NewVrrpConfigValueNull()
	var x types.Float64
	var y types.Float64

	var device_type types.String
	var serial types.String
	var mac types.String
	var model types.String

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
	if data.DhcpdConfig != nil {
		dhcpd_config = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
	}
	if data.DnsServers != nil {
		dns_servers = mist_transform.ListOfStringSdkToTerraform(ctx, data.DnsServers)
	}
	if data.DisableAutoConfig != nil {
		disable_auto_config = types.BoolValue(*data.DisableAutoConfig)
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
		device_id = types.StringValue(data.Id.String())
	}
	if data.Image1Url.Value() != nil {
		image1_url = types.StringValue("present")
	}
	if data.Image2Url.Value() != nil {
		image2_url = types.StringValue("present")
	}
	if data.Image3Url.Value() != nil {
		image3_url = types.StringValue("present")
	}
	if data.IpConfig != nil {
		ip_config = ipConfigSdkToTerraform(ctx, &diags, data.IpConfig)
	}
	if data.LocalPortConfig != nil {
		local_port_config = localPortConfigSdkToTerraform(ctx, &diags, data.LocalPortConfig)
	}
	if data.Managed != nil {
		managed = types.BoolValue(*data.Managed)
	}
	if data.MapId != nil {
		map_id = types.StringValue(data.MapId.String())
	}
	if data.MistNac != nil {
		mist_nac = mistNacSdkToTerraform(ctx, &diags, data.MistNac)
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}
	if data.Networks != nil && len(data.Networks) > 0 {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpServers != nil {
		ntp_servers = mist_transform.ListOfStringSdkToTerraform(ctx, data.NtpServers)
	}
	if data.OobIpConfig != nil {
		oob_ip_config = oobIpConfigsSdkToTerraform(ctx, &diags, data.OobIpConfig)
	}
	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.OspfAreas != nil {
		ospf_areas = ospfAreasSdkToTerraform(ctx, &diags, data.OspfAreas)
	}
	if data.OtherIpConfigs != nil && len(data.OtherIpConfigs) > 0 {
		other_ip_configs = otherIpConfigsSdkToTerraform(ctx, &diags, data.OtherIpConfigs)
	}
	if data.PortConfig != nil && len(data.PortConfig) > 0 {
		port_config = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
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
	if data.RemoteSyslog != nil {
		remote_syslog = remoteSyslogSdkToTerraform(ctx, &diags, data.RemoteSyslog)
	}
	if data.Role != nil {
		role = types.StringValue(*data.Role)
	}
	if data.RouterId != nil {
		router_id = types.StringValue(*data.RouterId)
	}
	if data.SiteId != nil {
		site_id = types.StringValue(data.SiteId.String())
	}
	if data.SnmpConfig != nil {
		snmp_config = snmpConfigSdkToTerraform(ctx, &diags, data.SnmpConfig)
	}
	if data.StpConfig != nil {
		stp_config = stpConfigSdkToTerraform(ctx, &diags, *data.StpConfig)
	}
	if data.SwitchMgmt != nil {
		switch_mgmt = switchMgmtSdkToTerraform(ctx, &diags, data.SwitchMgmt)
	}
	if data.UseRouterIdAsSourceIp != nil {
		use_router_id_as_source_ip = types.BoolValue(*data.UseRouterIdAsSourceIp)
	}
	if data.VirtualChassis != nil {
		virtual_chassis = virtualChassisSdkToTerraform(ctx, &diags, data.VirtualChassis)
	}
	if data.Vars != nil && len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}
	if data.VrfConfig != nil {
		vrf_config = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if data.VrfInstances != nil && len(data.VrfInstances) > 0 {
		vrf_instances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}
	if data.VrrpConfig != nil {
		vrrp_config = vrrpConfigInstancesSdkToTerraform(ctx, &diags, data.VrrpConfig)
	}
	if data.X != nil {
		x = types.Float64Value(float64(*data.X))
	}
	if data.Y != nil {
		y = types.Float64Value(float64(*data.Y))
	}

	device_type = types.StringValue(string(data.Type))

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}

	state.AclPolicies = acl_policies
	state.AclTags = acl_tags
	state.AdditionalConfigCmds = additional_config_cmds
	state.DeviceId = device_id
	state.DhcpSnooping = dhcp_snooping
	state.DhcpdConfig = dhcpd_config
	state.DisableAutoConfig = disable_auto_config
	state.DnsServers = dns_servers
	state.DnsSuffix = dns_suffix
	state.ExtraRoutes = extra_routes
	state.ExtraRoutes6 = extra_routes6
	state.Image1Url = image1_url
	state.Image2Url = image2_url
	state.Image3Url = image3_url
	state.IpConfig = ip_config
	state.LocalPortConfig = local_port_config
	state.Managed = managed
	state.MapId = map_id
	state.MistNac = mist_nac
	state.Name = name
	state.Notes = notes
	state.NtpServers = ntp_servers
	state.Networks = networks
	state.OobIpConfig = oob_ip_config
	state.OrgId = org_id
	state.OspfAreas = ospf_areas
	state.OtherIpConfigs = other_ip_configs
	state.PortConfig = port_config
	state.PortMirroring = port_mirroring
	state.PortUsages = port_usages
	state.RadiusConfig = radius_config
	state.RemoteSyslog = remote_syslog
	state.Role = role
	state.RouterId = router_id
	state.SiteId = site_id
	state.SnmpConfig = snmp_config
	state.StpConfig = stp_config
	state.SwitchMgmt = switch_mgmt
	state.UseRouterIdAsSourceIp = use_router_id_as_source_ip
	state.Vars = vars
	state.VirtualChassis = virtual_chassis
	state.VrfConfig = vrf_config
	state.VrfInstances = vrf_instances
	state.VrrpConfig = vrrp_config
	state.X = x
	state.Y = y
	state.Type = device_type
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
