package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *DeviceGatewayModel) (models.MistDevice, diag.Diagnostics) {
	data := models.DeviceGateway{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if len(plan.MapId.ValueString()) > 0 {
		map_id, e := uuid.Parse(plan.MapId.ValueString())
		if e == nil {
			data.MapId = &map_id
		} else {
			diags.AddError("Bad value for map_id", e.Error())
		}
	} else {
		unset["-map_id"] = ""
	}

	data.Name = plan.Name.ValueStringPointer()
	data.Notes = plan.Notes.ValueStringPointer()

	if plan.AdditionalConfigCmds.IsNull() || plan.AdditionalConfigCmds.IsUnknown() {
		unset["-additional_config_cmds"] = ""
	} else {
		data.AdditionalConfigCmds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.AdditionalConfigCmds)
	}

	if plan.BgpConfig.IsNull() || plan.BgpConfig.IsUnknown() {
		unset["-bgp_config"] = ""
	} else {
		data.BgpConfig = bgpConfigTerraformToSdk(ctx, &diags, plan.BgpConfig)
	}

	if plan.DhcpdConfig.IsNull() || plan.DhcpdConfig.IsUnknown() {
		unset["-dhcpd_config"] = ""
	} else {
		data.DhcpdConfig = dhcpdConfigTerraformToSdk(ctx, &diags, plan.DhcpdConfig)
	}

	data.DnsServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsServers)

	data.DnsSuffix = mist_transform.ListOfStringTerraformToSdk(ctx, plan.DnsSuffix)

	if plan.ExtraRoutes.IsNull() || plan.ExtraRoutes.IsUnknown() {
		unset["-extra_routes"] = ""
	} else {
		data.ExtraRoutes = extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes)
	}

	if plan.ExtraRoutes6.IsNull() || plan.ExtraRoutes6.IsUnknown() {
		unset["-extra_routes6"] = ""
	} else {
		data.ExtraRoutes6 = extraRoutesTerraformToSdk(ctx, &diags, plan.ExtraRoutes6)
	}

	if plan.IdpProfiles.IsNull() || plan.IdpProfiles.IsUnknown() {
		unset["-idp_profiles"] = ""
	} else {
		data.IdpProfiles = idpProfileTerraformToSdk(ctx, &diags, plan.IdpProfiles)
	}

	if plan.IpConfigs.IsNull() || plan.IpConfigs.IsUnknown() {
		unset["-ip_configs"] = ""
	} else {
		data.IpConfigs = ipConfigsTerraformToSdk(ctx, &diags, plan.IpConfigs)
	}

	if plan.Managed.IsNull() || plan.Managed.IsUnknown() {
		unset["-managed"] = ""
	} else {
		data.Managed = plan.Managed.ValueBoolPointer()
	}

	if plan.Networks.IsNull() || plan.Networks.IsUnknown() {
		unset["-networks"] = ""
	} else {
		data.Networks = networksTerraformToSdk(ctx, &diags, plan.Networks)
	}

	data.NtpServers = mist_transform.ListOfStringTerraformToSdk(ctx, plan.NtpServers)

	if plan.OobIpConfig.IsNull() || plan.OobIpConfig.IsUnknown() {
		unset["-oob_ip_config"] = ""
	} else {
		data.OobIpConfig = oobIpConfigTerraformToSdk(ctx, &diags, plan.OobIpConfig)
	}

	if plan.PathPreferences.IsNull() || plan.PathPreferences.IsUnknown() {
		unset["-path_preferences"] = ""
	} else {
		data.PathPreferences = pathPreferencesTerraformToSdk(ctx, &diags, plan.PathPreferences)
	}

	if plan.PortConfig.IsNull() || plan.PortConfig.IsUnknown() {
		unset["-port_config"] = ""
	} else {
		data.PortConfig = portConfigTerraformToSdk(ctx, &diags, plan.PortConfig)
	}

	data.RouterId = plan.RouterId.ValueStringPointer()

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

	if plan.TunnelConfigs.IsNull() || plan.TunnelConfigs.IsUnknown() {
		unset["-tunnel_configs"] = ""
	} else {
		data.TunnelConfigs = tunnelConfigsTerraformToSdk(ctx, &diags, plan.TunnelConfigs)
	}

	if plan.TunnelProviderOptions.IsNull() || plan.TunnelProviderOptions.IsUnknown() {
		unset["-tunnel_provider_options"] = ""
	} else {
		data.TunnelProviderOptions = tunnelProviderOptionsTerraformToSdk(ctx, &diags, plan.TunnelProviderOptions)
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(ctx, &diags, plan.Vars)
	} else {
		unset["-vars"] = ""
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

	if !plan.X.IsNull() && !plan.X.IsUnknown() {
		data.X = plan.X.ValueFloat64Pointer()
	} else {
		unset["-x"] = ""
	}
	if !plan.Y.IsNull() && !plan.Y.IsUnknown() {
		data.Y = plan.Y.ValueFloat64Pointer()
	} else {
		unset["-y"] = ""
	}

	data.AdditionalProperties = unset

	mist_device := models.MistDeviceContainer.FromDeviceGateway(data)
	return mist_device, diags
}
