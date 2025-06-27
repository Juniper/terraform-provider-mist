package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.GatewayTemplate) (OrgGatewaytemplateModel, diag.Diagnostics) {
	var state OrgGatewaytemplateModel
	var diags diag.Diagnostics

	var additionalConfigCmds = types.ListNull(types.StringType)
	var bgpConfig = types.MapNull(BgpConfigValue{}.Type(ctx))
	var dhcpdConfig = NewDhcpdConfigValueNull()
	var dnsOverride = types.BoolValue(false)
	var dnsServers = types.ListNull(types.StringType)
	var dnsSuffix = types.ListNull(types.StringType)
	var extraRoutes = types.MapNull(ExtraRoutesValue{}.Type(ctx))
	var extraRoutes6 = types.MapNull(ExtraRoutes6Value{}.Type(ctx))
	var id types.String
	var idpProfiles = types.MapNull(IdpProfilesValue{}.Type(ctx))
	var ipConfigs = types.MapNull(IpConfigsValue{}.Type(ctx))
	var name = types.StringValue(data.Name)
	var networks = types.ListNull(NetworksValue{}.Type(ctx))
	var ntpOverride = types.BoolValue(false)
	var ntpServers = types.ListNull(types.StringType)
	var oobIpConfig = NewOobIpConfigValueNull()
	var orgId types.String
	var pathPreferences = types.MapNull(PathPreferencesValue{}.Type(ctx))
	var portConfig = types.MapNull(PortConfigValue{}.Type(ctx))
	var routerId types.String
	var routingPolicies = types.MapNull(RoutingPoliciesValue{}.Type(ctx))
	var servicePolicies = types.ListNull(ServicePoliciesValue{}.Type(ctx))
	var tunnelConfigs = types.MapNull(TunnelConfigsValue{}.Type(ctx))
	var tunnelProviderOptions = NewTunnelProviderOptionsValueNull()
	var typeTemplate = types.StringValue("standalone")
	var vrfConfig = NewVrfConfigValueNull()
	var vrfInstances = types.MapNull(VrfInstancesValue{}.Type(ctx))

	if data.AdditionalConfigCmds != nil {
		additionalConfigCmds = mistutils.ListOfStringSdkToTerraform(data.AdditionalConfigCmds)
	}
	if len(data.BgpConfig) > 0 {
		bgpConfig = bgpConfigSdkToTerraform(ctx, &diags, data.BgpConfig)
	}
	if data.DhcpdConfig != nil {
		dhcpdConfig = dhcpdConfigSdkToTerraform(ctx, &diags, data.DhcpdConfig)
	}
	if data.DnsOverride != nil {
		dnsOverride = types.BoolValue(*data.DnsOverride)
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
		id = types.StringValue(data.Id.String())
	}
	if len(data.IdpProfiles) > 0 {
		idpProfiles = idpProfileSdkToTerraform(ctx, &diags, data.IdpProfiles)
	}
	if len(data.IpConfigs) > 0 {
		ipConfigs = ipConfigsSdkToTerraform(ctx, &diags, data.IpConfigs)
	}
	if data.Networks != nil {
		networks = networksSdkToTerraform(ctx, &diags, data.Networks)
	}
	if data.NtpOverride != nil {
		ntpOverride = types.BoolValue(*data.NtpOverride)
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
	if len(data.TunnelConfigs) > 0 {
		tunnelConfigs = tunnelConfigsSdkToTerraform(ctx, &diags, data.TunnelConfigs)
	}
	if data.TunnelProviderOptions != nil {
		if tunnelProviderOptionsTmp, ok := tunnelProviderSdkToTerraform(ctx, &diags, data.TunnelProviderOptions); ok {
			tunnelProviderOptions = tunnelProviderOptionsTmp
		}
	}
	if data.Type != nil {
		typeTemplate = types.StringValue(string(*data.Type))
	}
	if data.VrfConfig != nil {
		vrfConfig = vrfConfigSdkToTerraform(ctx, &diags, data.VrfConfig)
	}
	if len(data.VrfInstances) > 0 {
		vrfInstances = vrfInstancesSdkToTerraform(ctx, &diags, data.VrfInstances)
	}

	state.AdditionalConfigCmds = additionalConfigCmds
	state.BgpConfig = bgpConfig
	state.DhcpdConfig = dhcpdConfig
	state.DnsOverride = dnsOverride
	state.DnsServers = dnsServers
	state.DnsSuffix = dnsSuffix
	state.ExtraRoutes = extraRoutes
	state.ExtraRoutes6 = extraRoutes6
	state.Id = id
	state.IdpProfiles = idpProfiles
	state.IpConfigs = ipConfigs
	state.Name = name
	state.Networks = networks
	state.NtpOverride = ntpOverride
	state.NtpServers = ntpServers
	state.OobIpConfig = oobIpConfig
	state.OrgId = orgId
	state.PathPreferences = pathPreferences
	state.PortConfig = portConfig
	state.RouterId = routerId
	state.RoutingPolicies = routingPolicies
	state.ServicePolicies = servicePolicies
	state.TunnelConfigs = tunnelConfigs
	state.TunnelProviderOptions = tunnelProviderOptions
	state.Type = typeTemplate
	state.VrfConfig = vrfConfig
	state.VrfInstances = vrfInstances

	return state, diags
}
