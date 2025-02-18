package resource_org_gatewaytemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgGatewaytemplateModel) (*models.GatewayTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.GatewayTemplate{}

	data.Name = plan.Name.ValueString()

	additionalConfigCmds := misttransform.ListOfStringTerraformToSdk(plan.AdditionalConfigCmds)
	data.AdditionalConfigCmds = additionalConfigCmds

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		bgpConfig := bgpConfigTerraformToSdk(plan.BgpConfig)
		data.BgpConfig = bgpConfig
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		dhcpdConfig := dhcpdConfigTerraformToSdk(plan.DhcpdConfig)
		data.DhcpdConfig = &dhcpdConfig
	}

	if plan.DnsOverride.IsNull() || plan.DnsOverride.IsUnknown() {
		unset["-dns_override"] = ""
	} else {
		data.DnsOverride = plan.DnsOverride.ValueBoolPointer()
	}

	if plan.DnsServers.IsNull() || plan.DnsServers.IsUnknown() {
		unset["-dns_servers"] = ""
	} else {
		data.DnsServers = misttransform.ListOfStringTerraformToSdk(plan.DnsServers)
	}

	if plan.DnsSuffix.IsNull() || plan.DnsSuffix.IsUnknown() {
		unset["-dns_suffix"] = ""
	} else {
		data.DnsSuffix = misttransform.ListOfStringTerraformToSdk(plan.DnsSuffix)
	}

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		extraRoutes := extraRoutesTerraformToSdk(plan.ExtraRoutes)
		data.ExtraRoutes = extraRoutes
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutesTerraformToSdk(plan.ExtraRoutes6)
	}

	if plan.IdpProfiles.IsNull() || plan.IdpProfiles.IsUnknown() {
		unset["-idp_profiles"] = ""
	} else {
		idpProfiles := idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
		data.IdpProfiles = idpProfiles
	}

	if plan.IpConfigs.IsNull() || plan.IpConfigs.IsUnknown() {
		unset["-ip_configs"] = ""
	} else {
		ipConfigs := ipConfigsTerraformToSdk(plan.IpConfigs)
		data.IpConfigs = ipConfigs
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		networks := networksTerraformToSdk(ctx, &diags, plan.Networks)
		data.Networks = networks
	}

	if plan.NtpOverride.IsNull() || plan.NtpOverride.IsUnknown() {
		unset["-ntp_override"] = ""
	} else {
		data.NtpOverride = plan.NtpOverride.ValueBoolPointer()
	}

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = misttransform.ListOfStringTerraformToSdk(plan.NtpServers)
	}

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		oobIpConfig := oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
		data.OobIpConfig = oobIpConfig
	}

	if plan.PathPreferences.IsNull() || plan.PathPreferences.IsUnknown() {
		unset["-path_preferences"] = ""
	} else {
		pathPreferences := pathPreferencesTerraformToSdk(plan.PathPreferences)
		data.PathPreferences = pathPreferences
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		portConfig := portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
		data.PortConfig = portConfig
	}

	if plan.RouterId.IsNull() || plan.RouterId.IsUnknown() {
		unset["-router_id"] = ""
	} else {
		data.RouterId = plan.RouterId.ValueStringPointer()
	}

	if plan.RoutingPolicies.IsNull() || plan.RoutingPolicies.IsUnknown() {
		unset["-routing_policies"] = ""
	} else {
		routingPolicies := routingPoliciesTerraformToSdk(ctx, plan.RoutingPolicies)
		data.RoutingPolicies = routingPolicies
	}

	if plan.ServicePolicies.IsNull() || plan.ServicePolicies.IsUnknown() {
		unset["-service_policies"] = ""
	} else {
		servicePolicies := servicePoliciesTerraformToSdk(ctx, &diags, plan.ServicePolicies)
		data.ServicePolicies = servicePolicies
	}

	if plan.TunnelConfigs.IsNull() || plan.TunnelConfigs.IsUnknown() {
		unset["-tunnel_configs"] = ""
	} else {
		tunnelConfigs := tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
		data.TunnelConfigs = tunnelConfigs
	}

	if plan.TunnelProviderOptions.IsNull() || plan.TunnelProviderOptions.IsUnknown() {
		unset["-tunnel_provider_options"] = ""
	} else {
		data.TunnelProviderOptions = tunnelProviderOptionsTerraformToSdk(ctx, plan.TunnelProviderOptions)
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

	data.Type = models.ToPointer(models.GatewayTemplateTypeEnum(plan.Type.ValueString()))

	data.AdditionalProperties = unset
	return &data, diags
}
