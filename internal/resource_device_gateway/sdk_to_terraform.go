package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceGateway) (DeviceGatewayModel, diag.Diagnostics) {
	var state DeviceGatewayModel
	var diags diag.Diagnostics

	var additional_config_cmds types.List = types.ListNull(types.StringType)
	var bgp_config types.Map = types.MapNull(BgpConfigValue{}.Type(ctx))
	var deviceprofile_id types.String
	var dhcpd_config DhcpdConfigValue = NewDhcpdConfigValueNull()
	var dns_servers types.List = types.ListNull(types.StringType)
	var dns_suffix types.List = types.ListNull(types.StringType)
	var extra_routes types.Map = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extra_routes6 types.Map = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var deviceId types.String
	var idp_profiles types.Map = types.MapNull(IdpProfilesValue{}.Type(ctx))
	var image1_url types.String
	var image2_url types.String
	var image3_url types.String
	var ip_configs types.Map = types.MapNull(IpConfigsValue{}.Type(ctx))
	var map_id types.String
	var name types.String
	var networks types.List = types.ListNull(NetworksValue{}.Type(ctx))
	var notes types.String
	var ntp_servers types.List = types.ListNull(types.StringType)
	var oob_ip_config OobIpConfigValue = NewOobIpConfigValueNull()
	var org_id types.String
	var path_preferences types.Map = types.MapNull(PathPreferencesValue{}.Type(ctx))
	var port_config types.Map = types.MapNull(PortConfigValue{}.Type(ctx))
	var router_id types.String
	var routing_policies types.Map = types.MapNull(RoutingPoliciesValue{}.Type(ctx))
	var service_policies types.List = types.ListNull(ServicePoliciesValue{}.Type(ctx))
	var site_id types.String
	var tunnel_configs types.Map = types.MapNull(TunnelConfigsValue{}.Type(ctx))
	var tunnel_provider_options TunnelProviderOptionsValue = NewTunnelProviderOptionsValueNull()
	var vars types.Map = types.MapNull(types.StringType)
	var x types.Float64
	var y types.Float64

	var device_type types.String
	var serial types.String
	var mac types.String
	var model types.String

	if data.AdditionalConfigCmds != nil {
		additional_config_cmds = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalConfigCmds)
	}
	if data.BgpConfig != nil && len(data.BgpConfig) > 0 {
		bgp_config = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)
	}
	if data.DeviceprofileId != nil {
		deviceprofile_id = types.StringValue(data.DeviceprofileId.String())
	}
	if data.DhcpdConfig != nil {
		dhcpd_config = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
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
		deviceId = types.StringValue(data.Id.String())
	}
	if data.IdpProfiles != nil && len(data.IdpProfiles) > 0 {
		idp_profiles = idpProfileSdkToTerraform(ctx, &diags, data.IdpProfiles)
	}
	if data.Image1Url.Value() != nil {
		image1_url = types.StringValue(*data.Image1Url.Value())
	}
	if data.Image2Url.Value() != nil {
		image2_url = types.StringValue(*data.Image2Url.Value())
	}
	if data.Image3Url.Value() != nil {
		image3_url = types.StringValue(*data.Image3Url.Value())
	}
	if data.IpConfigs != nil && len(data.IpConfigs) > 0 {
		ip_configs = ipConfigsSdkToTerraform(ctx, &diags, data.IpConfigs)
	}
	if data.MapId != nil {
		map_id = types.StringValue(data.MapId.String())
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Networks != nil {
		networks = NetworksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
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
	if data.SiteId != nil {
		site_id = types.StringValue(data.SiteId.String())
	}
	if data.ServicePolicies != nil {
		service_policies = servicePoliciesSdkToTerraform(ctx, &diags, data.ServicePolicies)
	}
	if data.TunnelConfigs != nil && len(data.TunnelConfigs) > 0 {
		tunnel_configs = tunnelConfigsSdkToTerraform(ctx, &diags, data.TunnelConfigs)
	}
	if data.TunnelProviderOptions != nil {
		tunnel_provider_options = tunnelProviderSdkToTerraform(ctx, &diags, data.TunnelProviderOptions)
	}
	if data.Vars != nil && len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}
	if data.X != nil {
		x = types.Float64Value(float64(*data.X))
	}
	if data.Y != nil {
		y = types.Float64Value(float64(*data.Y))
	}

	if data.Type != nil {
		device_type = types.StringValue(string(*data.Type))
	}

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}

	state.AdditionalConfigCmds = additional_config_cmds
	state.BgpConfig = bgp_config
	state.DeviceprofileId = deviceprofile_id
	state.DhcpdConfig = dhcpd_config
	state.DnsServers = dns_servers
	state.DnsSuffix = dns_suffix
	state.ExtraRoutes = extra_routes
	state.ExtraRoutes6 = extra_routes6
	state.DeviceId = deviceId
	state.IdpProfiles = idp_profiles
	state.Image1Url = image1_url
	state.Image2Url = image2_url
	state.Image3Url = image3_url
	state.IpConfigs = ip_configs
	state.MapId = map_id
	state.Name = name
	state.Networks = networks
	state.NtpServers = ntp_servers
	state.Notes = notes
	state.OobIpConfig = oob_ip_config
	state.OrgId = org_id
	state.PathPreferences = path_preferences
	state.PortConfig = port_config
	state.RouterId = router_id
	state.RoutingPolicies = routing_policies
	state.ServicePolicies = service_policies
	state.SiteId = site_id
	state.TunnelConfigs = tunnel_configs
	state.TunnelProviderOptions = tunnel_provider_options
	state.Vars = vars
	state.X = x
	state.Y = y
	state.Type = device_type
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
