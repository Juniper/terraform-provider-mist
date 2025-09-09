package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.DeviceGateway) (DeviceGatewayModel, diag.Diagnostics) {
	var state DeviceGatewayModel
	var diags diag.Diagnostics

	var additionalConfigCmds = types.ListNull(types.StringType)
	var bgpConfig = types.MapNull(BgpConfigValue{}.Type(ctx))
	var dhcpdConfig = NewDhcpdConfigValueNull()
	var dnsServers = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var deviceId types.String
	var idpProfiles = types.MapNull(IdpProfilesValue{}.Type(ctx))
	var image1Url = types.StringValue("not_present")
	var image2Url = types.StringValue("not_present")
	var image3Url = types.StringValue("not_present")
	var ipConfigs = types.MapNull(IpConfigsValue{}.Type(ctx))
	var managed types.Bool
	var mapId types.String
	var name types.String
	var networks = types.ListNull(NetworksValue{}.Type(ctx))
	var notes types.String
	var ntpServers = types.ListNull(types.StringType)
	var oobIpConfig = NewOobIpConfigValueNull()
	var orgId types.String
	var pathPreferences = types.MapNull(PathPreferencesValue{}.Type(ctx))
	var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
	var routerId types.String
	var routingPolicies = types.MapNull(RoutingPoliciesValue{}.Type(ctx))
	var servicePolicies = types.ListNull(ServicePoliciesValue{}.Type(ctx))
	var siteId types.String
	var ssrAdditionalConfigCmds = types.ListNull(types.StringType)
	var tunnelConfigs = types.MapNull(TunnelConfigsValue{}.Type(ctx))
	var tunnelProviderOptions = NewTunnelProviderOptionsValueNull()
	var vars = types.MapNull(types.StringType)
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))
	var x types.Float64
	var y types.Float64

	var deviceType types.String
	var serial types.String
	var mac types.String
	var model types.String

	if data.AdditionalConfigCmds != nil {
		additionalConfigCmds = mistutils.ListOfStringSdkToTerraform(data.AdditionalConfigCmds)
	}
	if len(data.BgpConfig) > 0 {
		bgpConfig = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)
	}
	if data.DhcpdConfig != nil {
		dhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
	}
	if data.DnsServers != nil {
		dnsServers = mistutils.ListOfStringSdkToTerraform(data.DnsServers)
	}
	if data.DnsSuffix != nil {
		dnsSuffix = mistutils.ListOfStringSdkToTerraform(data.DnsSuffix)
	}
	if len(data.ExtraRoutes) > 0 {
		extraRoutes = extraRoutesSdkToTerraform(ctx, &diags, data.ExtraRoutes)
	}
	if len(data.ExtraRoutes6) > 0 {
		extraRoutes6 = extraRoutes6SdkToTerraform(ctx, &diags, data.ExtraRoutes6)
	}
	if data.Id != nil {
		deviceId = types.StringValue(data.Id.String())
	}
	if len(data.IdpProfiles) > 0 {
		idpProfiles = idpProfileSdkToTerraform(ctx, &diags, data.IdpProfiles)
	}
	if data.Image1Url.Value() != nil {
		image1Url = types.StringValue("present")
	}
	if data.Image2Url.Value() != nil {
		image2Url = types.StringValue("present")
	}
	if data.Image3Url.Value() != nil {
		image3Url = types.StringValue("present")
	}
	if len(data.IpConfigs) > 0 {
		ipConfigs = ipConfigsSdkToTerraform(ctx, &diags, data.IpConfigs)
	}
	if data.Managed != nil {
		managed = types.BoolValue(*data.Managed)
	}
	if data.MapId != nil {
		mapId = types.StringValue(data.MapId.String())
	}
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	if data.Networks != nil {
		networks = networksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.Notes != nil {
		notes = types.StringValue(*data.Notes)
	}
	if data.NtpServers != nil {
		ntpServers = mistutils.ListOfStringSdkToTerraform(data.NtpServers)
	}
	if data.OobIpConfig != nil {
		oobIpConfig = oobIpConfigsSdkToTerraform(ctx, &diags, data.OobIpConfig)
	}
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if len(data.PathPreferences) > 0 {
		pathPreferences = pathPreferencesSdkToTerraform(ctx, &diags, data.PathPreferences)
	}
	if len(data.PortConfig) > 0 {
		portConfig = portConfigSdkToTerraform(ctx, &diags, data.PortConfig)
	}
	if data.RouterId != nil {
		routerId = types.StringValue(*data.RouterId)
	}
	if len(data.RoutingPolicies) > 0 {
		routingPolicies = routingPoliciesSdkToTerraform(ctx, &diags, data.RoutingPolicies)
	}
	if data.ServicePolicies != nil {
		servicePolicies = servicePoliciesSdkToTerraform(ctx, &diags, data.ServicePolicies)
	}
	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}
	if data.SsrAdditionalConfigCmds != nil {
		ssrAdditionalConfigCmds = mistutils.ListOfStringSdkToTerraform(data.SsrAdditionalConfigCmds)
	}
	if len(data.TunnelConfigs) > 0 {
		tunnelConfigs = tunnelConfigsSdkToTerraform(ctx, &diags, data.TunnelConfigs)
	}
	if data.TunnelProviderOptions != nil {
		if tunnelProviderOptionsTmp, ok := tunnelProviderSdkToTerraform(ctx, &diags, data.TunnelProviderOptions); ok {
			tunnelProviderOptions = tunnelProviderOptionsTmp
		}
	}
	if len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}
	if data.VrfConfig != nil {
		vrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if len(data.VrfInstances) > 0 {
		vrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}
	if data.X != nil {
		x = types.Float64Value(*data.X)
	}
	if data.Y != nil {
		y = types.Float64Value(*data.Y)
	}

	deviceType = types.StringValue(data.Type)

	if data.Serial != nil {
		serial = types.StringValue(*data.Serial)
	}

	if data.Mac != nil {
		mac = types.StringValue(*data.Mac)
	}

	if data.Model != nil {
		model = types.StringValue(*data.Model)
	}

	state.AdditionalConfigCmds = additionalConfigCmds
	state.BgpConfig = bgpConfig
	state.DhcpdConfig = dhcpdConfig
	state.DnsServers = dnsServers
	state.DnsSuffix = dnsSuffix
	state.ExtraRoutes = extraRoutes
	state.ExtraRoutes6 = extraRoutes6
	state.DeviceId = deviceId
	state.IdpProfiles = idpProfiles
	state.Image1Url = image1Url
	state.Image2Url = image2Url
	state.Image3Url = image3Url
	state.IpConfigs = ipConfigs
	state.Managed = managed
	state.MapId = mapId
	state.Name = name
	state.Networks = networks
	state.NtpServers = ntpServers
	state.Notes = notes
	state.OobIpConfig = oobIpConfig
	state.OrgId = orgId
	state.PathPreferences = pathPreferences
	state.PortConfig = portConfig
	state.RouterId = routerId
	state.RoutingPolicies = routingPolicies
	state.ServicePolicies = servicePolicies
	state.SiteId = siteId
	state.SsrAdditionalConfigCmds = ssrAdditionalConfigCmds
	state.TunnelConfigs = tunnelConfigs
	state.TunnelProviderOptions = tunnelProviderOptions
	state.Vars = vars
	state.VrfConfig = vrfConfig
	state.VrfInstances = vrfInstances
	state.X = x
	state.Y = y
	state.Type = deviceType
	state.Serial = serial
	state.Mac = mac
	state.Model = model

	return state, diags
}
