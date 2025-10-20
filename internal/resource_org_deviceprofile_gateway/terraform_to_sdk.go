package resource_org_deviceprofile_gateway

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgDeviceprofileGatewayModel) (models.Deviceprofile, diag.Diagnostics) {
	data := models.DeviceprofileGateway{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueString()

	additionalConfigCmds := mistutils.ListOfStringTerraformToSdk(plan.AdditionalConfigCmds)
	data.AdditionalConfigCmds = additionalConfigCmds

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		data.BgpConfig = bgpConfigTerraformToSdk(plan.BgpConfig)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		data.DhcpdConfig = dhcpdConfigTerraformToSdk(plan.DhcpdConfig)
	}

	if plan.DnsOverride.IsNull() || plan.DnsOverride.IsUnknown() {
		unset["-dnsOverride"] = ""
	} else {
		data.DnsOverride = plan.DnsOverride.ValueBoolPointer()
	}

	if plan.DnsServers.IsNull() || plan.DnsServers.IsUnknown() {
		unset["-dns_servers"] = ""
	} else {
		data.DnsServers = mistutils.ListOfStringTerraformToSdk(plan.DnsServers)
	}

	if plan.DnsSuffix.IsNull() || plan.DnsSuffix.IsUnknown() {
		unset["-dns_suffix"] = ""
	} else {
		data.DnsSuffix = mistutils.ListOfStringTerraformToSdk(plan.DnsSuffix)
	}

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		data.ExtraRoutes = extraRoutesTerraformToSdk(plan.ExtraRoutes)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutes6TerraformToSdk(plan.ExtraRoutes6)
	}

	if plan.IdpProfiles.IsNull() || plan.IdpProfiles.IsUnknown() {
		unset["-idp_profiles"] = ""
	} else {
		data.IdpProfiles = idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
	}

	if plan.IpConfigs.IsNull() || plan.IpConfigs.IsUnknown() {
		unset["-ip_configs"] = ""
	} else {
		data.IpConfigs = ipConfigsTerraformToSdk(plan.IpConfigs)
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		data.Networks = networksTerraformToSdk(ctx, &diags, plan.Networks)
	}

	if plan.NtpOverride.IsNull() || plan.NtpOverride.IsUnknown() {
		unset["-ntpOverride"] = ""
	} else {
		data.NtpOverride = plan.NtpOverride.ValueBoolPointer()
	}

	if plan.NtpServers.IsNull() || plan.NtpServers.IsUnknown() {
		unset["-ntp_servers"] = ""
	} else {
		data.NtpServers = mistutils.ListOfStringTerraformToSdk(plan.NtpServers)
	}

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		data.OobIpConfig = oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
	}

	if plan.PathPreferences.IsNull() || plan.PathPreferences.IsUnknown() {
		unset["-path_preferences"] = ""
	} else {
		data.PathPreferences = pathPreferencesTerraformToSdk(plan.PathPreferences)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		data.PortConfig = portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
	}

	if plan.RouterId.IsNull() || plan.RouterId.IsUnknown() {
		unset["-router_id"] = ""
	} else {
		data.RouterId = plan.RouterId.ValueStringPointer()
	}

	if plan.RoutingPolicies.IsNull() || plan.RoutingPolicies.IsUnknown() {
		unset["-routing_policies"] = ""
	} else {
		data.RoutingPolicies = routingPoliciesTerraformToSdk(ctx, &diags, plan.RoutingPolicies)
	}

	if plan.ServicePolicies.IsNull() || plan.ServicePolicies.IsUnknown() {
		unset["-service_policies"] = ""
	} else {
		data.ServicePolicies = servicePoliciesTerraformToSdk(ctx, &diags, plan.ServicePolicies)
	}

	if plan.SsrAdditionalConfigCmds.IsNull() || plan.SsrAdditionalConfigCmds.IsUnknown() {
		unset["-ssr_additional_config_cmds"] = ""
	} else {
		data.SsrAdditionalConfigCmds = mistutils.ListOfStringTerraformToSdk(plan.SsrAdditionalConfigCmds)
	}

	if plan.TunnelConfigs.IsNull() || plan.TunnelConfigs.IsUnknown() {
		unset["-tunnel_configs"] = ""
	} else {
		data.TunnelConfigs = tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
	}

	if plan.TunnelProviderOptions.IsNull() || plan.TunnelProviderOptions.IsUnknown() {
		unset["-tunnel_provider_options"] = ""
	} else {
		data.TunnelProviderOptions = tunnelProviderOptionsTerraformToSdk(ctx, plan.TunnelProviderOptions)
	}

	if data.UrlFilteringDenyMsg == nil && (plan.UrlFilteringDenyMsg.IsNull() || plan.UrlFilteringDenyMsg.IsUnknown()) {
		unset["-url_filtering_deny_msg"] = ""
	} else if !plan.UrlFilteringDenyMsg.IsNull() && !plan.UrlFilteringDenyMsg.IsUnknown() {
		data.UrlFilteringDenyMsg = plan.UrlFilteringDenyMsg.ValueStringPointer()
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

	data.Type = string(models.ConstDeviceTypeGatewayEnum_GATEWAY)
	data.AdditionalProperties = unset

	deviceprofile := models.DeviceprofileContainer.FromDeviceprofileGateway(data)

	return deviceprofile, diags
}
