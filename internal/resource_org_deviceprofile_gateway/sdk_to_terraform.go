package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceprofileGateway) (OrgDeviceprofileGatewayModel, diag.Diagnostics) {
	var state OrgDeviceprofileGatewayModel
	var diags diag.Diagnostics

	var additional_config_cmds types.List = types.ListNull(types.StringType)
	var bgp_config types.Map = types.MapNull(BgpConfigValue{}.Type(ctx))
	var dhcpd_config DhcpdConfigValue = NewDhcpdConfigValueNull()
	var dns_override types.Bool = types.BoolValue(false)
	var dns_servers types.List = types.ListNull(types.StringType)
	var dns_suffix types.List = types.ListNull(types.StringType)
	var extra_routes types.Map = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extra_routes6 types.Map = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var idp_profiles types.Map = types.MapNull(IdpProfilesValue{}.Type(ctx))
	var ip_configs types.Map = types.MapNull(IpConfigsValue{}.Type(ctx))
	var name types.String = types.StringValue(data.Name)
	var networks types.List = types.ListNull(NetworksValue{}.Type(ctx))
	var ntp_override types.Bool = types.BoolValue(false)
	var ntp_servers types.List = types.ListNull(types.StringType)
	var oob_ip_config OobIpConfigValue = NewOobIpConfigValueNull()
	var org_id types.String
	var path_preferences types.Map = types.MapNull(PathPreferencesValue{}.Type(ctx))
	var port_config types.Map = types.MapNull(PortConfigValue{}.Type(ctx))
	var router_id types.String
	var routing_policies types.Map = types.MapNull(RoutingPoliciesValue{}.Type(ctx))
	var service_policies types.List = types.ListNull(ServicePoliciesValue{}.Type(ctx))
	var tunnel_configs types.Map = types.MapNull(TunnelConfigsValue{}.Type(ctx))
	var tunnel_provider_options TunnelProviderOptionsValue = NewTunnelProviderOptionsValueNull()
	var type_template types.String
	var vrf_config VrfConfigValue = NewVrfConfigValueNull()
	var vrf_instances types.Map = types.MapNull(VrfInstancesValue{}.Type(ctx))

	if data.AdditionalConfigCmds != nil {
		additional_config_cmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalConfigCmds)
	}
	if data.BgpConfig != nil && len(data.BgpConfig) > 0 {
		bgp_config = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)
	}
	if data.DhcpdConfig != nil {
		dhcpd_config = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
	}
	if data.DnsOverride != nil {
		dns_override = types.BoolValue(*data.DnsOverride)
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
	if data.IdpProfiles != nil && len(data.IdpProfiles) > 0 {
		idp_profiles = idpProfileSdkToTerraform(ctx, &diags, data.IdpProfiles)
	}
	if data.IpConfigs != nil && len(data.IpConfigs) > 0 {
		ip_configs = ipConfigsSdkToTerraform(ctx, &diags, data.IpConfigs)
	}
	if data.Networks != nil {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpOverride != nil {
		ntp_override = types.BoolValue(*data.NtpOverride)
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
	if data.PathPreferences != nil && len(data.PathPreferences) > 0 {
		path_preferences = pathPreferencesSdkToTerraform(ctx, &diags, data.PathPreferences)
	}
	if data.PortConfig != nil && len(data.PortConfig) > 0 {
		port_config = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
	}
	if data.RouterId != nil {
		router_id = types.StringValue(*data.RouterId)
	}
	if data.RoutingPolicies != nil && len(data.RoutingPolicies) > 0 {
		routing_policies = routingPolociesSdkToTerraform(ctx, &diags, data.RoutingPolicies)
	}
	if data.ServicePolicies != nil && len(data.ServicePolicies) > 0 {
		service_policies = servicePoliciesSdkToTerraform(ctx, &diags, data.ServicePolicies)
	}
	if data.TunnelConfigs != nil && len(data.TunnelConfigs) > 0 {
		tunnel_configs = tunnelConfigsSdkToTerraform(ctx, &diags, data.TunnelConfigs)
	}
	if data.TunnelProviderOptions != nil {
		tunnel_provider_options = tunnelProviderSdkToTerraform(ctx, &diags, data.TunnelProviderOptions)
	}
	if data.Type != nil {
		type_template = types.StringValue(string(*data.Type))
	}
	if data.VrfConfig != nil {
		vrf_config = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if data.VrfInstances != nil && len(data.VrfInstances) > 0 {
		vrf_instances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}

	state.AdditionalConfigCmds = additional_config_cmds
	state.BgpConfig = bgp_config
	state.DhcpdConfig = dhcpd_config
	state.DnsOverride = dns_override
	state.DnsServers = dns_servers
	state.DnsSuffix = dns_suffix
	state.ExtraRoutes = extra_routes
	state.ExtraRoutes6 = extra_routes6
	state.Id = id
	state.IdpProfiles = idp_profiles
	state.IpConfigs = ip_configs
	state.Name = name
	state.Networks = networks
	state.NtpOverride = ntp_override
	state.NtpServers = ntp_servers
	state.OobIpConfig = oob_ip_config
	state.OrgId = org_id
	state.PathPreferences = path_preferences
	state.PortConfig = port_config
	state.RouterId = router_id
	state.RoutingPolicies = routing_policies
	state.ServicePolicies = service_policies
	state.TunnelConfigs = tunnel_configs
	state.TunnelProviderOptions = tunnel_provider_options
	state.Type = type_template
	state.VrfConfig = vrf_config
	state.VrfInstances = vrf_instances

	return state, diags
}
