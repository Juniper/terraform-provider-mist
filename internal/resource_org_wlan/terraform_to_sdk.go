package resource_org_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TerraformToSdk(ctx context.Context, plan *OrgWlanModel) (*models.Wlan, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.Wlan{}
	unset := make(map[string]interface{})

	data.Ssid = plan.Ssid.ValueString()

	if len(plan.TemplateId.ValueString()) > 0 {
		data.TemplateId = models.NewOptional(models.ToPointer(uuid.MustParse(plan.TemplateId.ValueString())))
	}

	if plan.AcctImmediateUpdate.IsNull() || plan.AcctImmediateUpdate.IsUnknown() {
		unset["-acct_immediate_update"] = ""
	} else {
		data.AcctImmediateUpdate = plan.AcctImmediateUpdate.ValueBoolPointer()
	}

	if plan.AcctInterimInterval.IsNull() || plan.AcctInterimInterval.IsUnknown() {
		unset["-acct_interim_interval"] = ""
	} else {
		data.AcctInterimInterval = models.ToPointer(int(plan.AcctInterimInterval.ValueInt64()))
	}

	if plan.AcctServers.IsNull() || plan.AcctServers.IsUnknown() {
		unset["-acct_servers"] = ""
	} else {
		acct_servers := radiusAcctServersTerraformToSdk(ctx, &diags, plan.AcctServers)
		data.AcctServers = acct_servers
	}

	if plan.Airwatch.IsNull() || plan.Airwatch.IsUnknown() {
		unset["-airwatch"] = ""
	} else {
		airwatch := airwatchTerraformToSdk(ctx, &diags, plan.Airwatch)
		data.Airwatch = airwatch
	}

	if plan.AllowIpv6Ndp.IsNull() || plan.AllowIpv6Ndp.IsUnknown() {
		unset["-allow_ipv6_ndp"] = ""
	} else {
		data.AllowIpv6Ndp = plan.AllowIpv6Ndp.ValueBoolPointer()
	}

	if plan.AllowMdns.IsNull() || plan.AllowMdns.IsUnknown() {
		unset["-allow_mdns"] = ""
	} else {
		data.AllowMdns = plan.AllowMdns.ValueBoolPointer()
	}

	if plan.AllowSsdp.IsNull() || plan.AllowSsdp.IsUnknown() {
		unset["-allow_ssdp"] = ""
	} else {
		data.AllowSsdp = plan.AllowSsdp.ValueBoolPointer()
	}

	if plan.ApIds.IsNull() || plan.ApIds.IsUnknown() {
		unset["-ap_ids"] = ""
	} else {
		data.ApIds = models.NewOptional(models.ToPointer(mist_transform.ListOfUuidTerraformToSdk(ctx, plan.ApIds)))
	}

	if plan.AppLimit.IsNull() || plan.AppLimit.IsUnknown() {
		unset["-app_limit"] = ""
	} else {
		app_limit := appLimitTerraformToSdk(ctx, &diags, plan.AppLimit)
		data.AppLimit = app_limit
	}

	if plan.AppQos.IsNull() || plan.AppQos.IsUnknown() {
		unset["-app_qos"] = ""
	} else {
		app_qos := appQosTerraformToSdk(ctx, &diags, plan.AppQos)
		data.AppQos = app_qos
	}

	if plan.ApplyTo.IsNull() || plan.ApplyTo.IsUnknown() {
		unset["-apply_to"] = ""
	} else {
		data.ApplyTo = models.ToPointer(models.WlanApplyToEnum(string(plan.ApplyTo.ValueString())))
	}

	if plan.ArpFilter.IsNull() || plan.ArpFilter.IsUnknown() {
		unset["-arp_filter"] = ""
	} else {
		data.ArpFilter = plan.ArpFilter.ValueBoolPointer()
	}

	if plan.Auth.IsNull() || plan.Auth.IsUnknown() {
		unset["-auth"] = ""
	} else {
		auth := authTerraformToSdk(ctx, &diags, plan.Auth)
		data.Auth = auth
	}

	if plan.AuthServerSelection.IsNull() || plan.AuthServerSelection.IsUnknown() {
		unset["-auth_server_selection"] = ""
	} else {
		data.AuthServerSelection = models.ToPointer(models.WlanAuthServerSelectionEnum(plan.AuthServerSelection.ValueString()))
	}

	if plan.AuthServers.IsNull() || plan.AuthServers.IsUnknown() {
		unset["-auth_servers"] = ""
	} else {
		auth_servers := radiusAuthServersTerraformToSdk(ctx, &diags, plan.AuthServers)
		data.AuthServers = auth_servers
	}

	if plan.AuthServersNasId.IsNull() || plan.AuthServersNasId.IsUnknown() {
		unset["-auth_servers_nas_id"] = ""
	} else {
		data.AuthServersNasId = models.NewOptional(plan.AuthServersNasId.ValueStringPointer())
	}

	if plan.AuthServersNasIp.IsNull() || plan.AuthServersNasIp.IsUnknown() {
		unset["-auth_servers_nas_ip"] = ""
	} else {
		data.AuthServersNasIp = models.NewOptional(plan.AuthServersNasIp.ValueStringPointer())
	}

	if plan.AuthServersRetries.IsNull() || plan.AuthServersRetries.IsUnknown() {
		unset["-auth_servers_retries"] = ""
	} else {
		data.AuthServersRetries = models.ToPointer(int(plan.AuthServersRetries.ValueInt64()))
	}

	if plan.AuthServersTimeout.IsNull() || plan.AuthServersTimeout.IsUnknown() {
		unset["-auth_servers_timeout"] = ""
	} else {
		data.AuthServersTimeout = models.ToPointer(int(plan.AuthServersTimeout.ValueInt64()))
	}

	if plan.BandSteer.IsNull() || plan.BandSteer.IsUnknown() {
		unset["-band_steer"] = ""
	} else {
		data.BandSteer = plan.BandSteer.ValueBoolPointer()
	}

	if plan.BandSteerForceBand5.IsNull() || plan.BandSteerForceBand5.IsUnknown() {
		unset["-band_steer_force_band5"] = ""
	} else {
		data.BandSteerForceBand5 = plan.BandSteerForceBand5.ValueBoolPointer()
	}

	if plan.Bands.IsNull() || plan.Bands.IsUnknown() {
		unset["-bands"] = ""
	} else {
		bands := bandsTerraformToSdk(ctx, &diags, plan.Bands)
		data.Bands = bands
	}

	if plan.BlockBlacklistClients.IsNull() || plan.BlockBlacklistClients.IsUnknown() {
		unset["-block_blacklist_clients"] = ""
	} else {
		data.BlockBlacklistClients = plan.BlockBlacklistClients.ValueBoolPointer()
	}

	if plan.Bonjour.IsNull() || plan.Bonjour.IsUnknown() {
		unset["-bonjour"] = ""
	} else {
		bonjour := bonjourTerraformToSdk(ctx, &diags, plan.Bonjour)
		data.Bonjour = bonjour
	}

	if plan.CiscoCwa.IsNull() || plan.CiscoCwa.IsUnknown() {
		unset["-cisco_cwa"] = ""
	} else {
		cisco_cwa := ciscoCwaTerraformToSdk(ctx, &diags, plan.CiscoCwa)
		data.CiscoCwa = cisco_cwa
	}

	if plan.ClientLimitDown.IsNull() || plan.ClientLimitDown.IsUnknown() {
		unset["-client_limit_down"] = ""
	} else {
		data.ClientLimitDown = models.ToPointer(int(plan.ClientLimitDown.ValueInt64()))
	}

	if plan.ClientLimitDownEnabled.IsNull() || plan.ClientLimitDownEnabled.IsUnknown() {
		unset["-client_limit_down_enabled"] = ""
	} else {
		data.ClientLimitDownEnabled = plan.ClientLimitDownEnabled.ValueBoolPointer()
	}

	if plan.ClientLimitUp.IsNull() || plan.ClientLimitUp.IsUnknown() {
		unset["-client_limit_up"] = ""
	} else {
		data.ClientLimitUp = models.ToPointer(int(plan.ClientLimitUp.ValueInt64()))
	}

	if plan.ClientLimitUpEnabled.IsNull() || plan.ClientLimitUpEnabled.IsUnknown() {
		unset["-client_limit_up_enabled"] = ""
	} else {
		data.ClientLimitUpEnabled = plan.ClientLimitUpEnabled.ValueBoolPointer()
	}

	if plan.CoaServers.IsNull() || plan.CoaServers.IsUnknown() {
		unset["-coa_servers"] = ""
	} else {
		coa_servers := coaServerTerraformToSdk(ctx, &diags, plan.CoaServers)
		data.CoaServers = models.NewOptional(models.ToPointer(coa_servers))
	}

	if plan.Disable11ax.IsNull() || plan.Disable11ax.IsUnknown() {
		unset["-disable_11ax"] = ""
	} else {
		data.Disable11ax = plan.Disable11ax.ValueBoolPointer()
	}

	if plan.DisableHtVhtRates.IsNull() || plan.DisableHtVhtRates.IsUnknown() {
		unset["-disable_ht_vht_rates"] = ""
	} else {
		data.DisableHtVhtRates = plan.DisableHtVhtRates.ValueBoolPointer()
	}

	if plan.DisableUapsd.IsNull() || plan.DisableUapsd.IsUnknown() {
		unset["-disable_uapsd"] = ""
	} else {
		data.DisableUapsd = plan.DisableUapsd.ValueBoolPointer()
	}

	if plan.DisableV1RoamNotify.IsNull() || plan.DisableV1RoamNotify.IsUnknown() {
		unset["-disable_v1_roam_notify"] = ""
	} else {
		data.DisableV1RoamNotify = plan.DisableV1RoamNotify.ValueBoolPointer()
	}

	if plan.DisableV2RoamNotify.IsNull() || plan.DisableV2RoamNotify.IsUnknown() {
		unset["-disable_v2_roam_notify"] = ""
	} else {
		data.DisableV2RoamNotify = plan.DisableV2RoamNotify.ValueBoolPointer()
	}

	if plan.DisableWhenGatewayUnreachable.IsNull() || plan.DisableWhenGatewayUnreachable.IsUnknown() {
		unset["-disable_when_gateway_unreachable"] = ""
	} else {
		data.DisableWhenGatewayUnreachable = plan.DisableWhenGatewayUnreachable.ValueBoolPointer()
	}

	if plan.DisableWhenMxtunnelDown.IsNull() || plan.DisableWhenMxtunnelDown.IsUnknown() {
		unset["-disable_when_mxtunnel_down"] = ""
	} else {
		data.DisableWhenMxtunnelDown = plan.DisableWhenMxtunnelDown.ValueBoolPointer()
	}

	if plan.DisableWmm.IsNull() || plan.DisableWmm.IsUnknown() {
		unset["-disable_wmm"] = ""
	} else {
		data.DisableWmm = plan.DisableWmm.ValueBoolPointer()
	}

	if plan.DnsServerRewrite.IsNull() || plan.DnsServerRewrite.IsUnknown() {
		unset["-dns_server_rewrite"] = ""
	} else {
		dns_server_rewrite := dnsServerRewriteTerraformToSdk(ctx, &diags, plan.DnsServerRewrite)
		data.DnsServerRewrite = models.NewOptional(dns_server_rewrite)
	}

	if plan.Dtim.IsNull() || plan.Dtim.IsUnknown() {
		unset["-dtim"] = ""
	} else {
		data.Dtim = models.ToPointer(int(plan.Dtim.ValueInt64()))
	}

	if plan.DynamicPsk.IsNull() || plan.DynamicPsk.IsUnknown() {
		unset["-dynamic_psk"] = ""
	} else {
		dynamic_psk := dynamicPskTerraformToSdk(ctx, &diags, plan.DynamicPsk)
		data.DynamicPsk = models.NewOptional(dynamic_psk)
	}

	if plan.DynamicVlan.IsNull() || plan.DynamicVlan.IsUnknown() {
		unset["-dynamic_vlan"] = ""
	} else {
		dynamic_vlan := dynamicVlanTerraformToSdk(ctx, &diags, plan.DynamicVlan)
		data.DynamicVlan = models.NewOptional(dynamic_vlan)
	}

	if plan.EnableLocalKeycaching.IsNull() || plan.EnableLocalKeycaching.IsUnknown() {
		unset["-enable_local_keycaching"] = ""
	} else {
		data.EnableLocalKeycaching = plan.EnableLocalKeycaching.ValueBoolPointer()
	}

	if plan.EnableWirelessBridging.IsNull() || plan.EnableWirelessBridging.IsUnknown() {
		unset["-enable_wireless_bridging"] = ""
	} else {
		data.EnableWirelessBridging = plan.EnableWirelessBridging.ValueBoolPointer()
	}

	if plan.EnableWirelessBridgingDhcpTracking.IsNull() || plan.EnableWirelessBridgingDhcpTracking.IsUnknown() {
		unset["-enable_wireless_bridging_dhcp_tracking"] = ""
	} else {
		data.EnableWirelessBridgingDhcpTracking = plan.EnableWirelessBridgingDhcpTracking.ValueBoolPointer()
	}

	if plan.Enabled.IsNull() || plan.Enabled.IsUnknown() {
		unset["-enabled"] = ""
	} else {
		data.Enabled = plan.Enabled.ValueBoolPointer()
	}

	if plan.FastDot1xTimers.IsNull() || plan.FastDot1xTimers.IsUnknown() {
		unset["-fast_dot1x_timers"] = ""
	} else {
		data.FastDot1xTimers = plan.FastDot1xTimers.ValueBoolPointer()
	}

	if plan.HideSsid.IsNull() || plan.HideSsid.IsUnknown() {
		unset["-hide_ssid"] = ""
	} else {
		data.HideSsid = plan.HideSsid.ValueBoolPointer()
	}

	if plan.HostnameIe.IsNull() || plan.HostnameIe.IsUnknown() {
		unset["-hostname_ie"] = ""
	} else {
		data.HostnameIe = plan.HostnameIe.ValueBoolPointer()
	}

	if plan.Hotspot20.IsNull() || plan.Hotspot20.IsUnknown() {
		unset["-hotspot20"] = ""
	} else {
		hotspot20 := hotspot20TerraformToSdk(ctx, &diags, plan.Hotspot20)
		data.Hotspot20 = hotspot20
	}

	if plan.InjectDhcpOption82.IsNull() || plan.InjectDhcpOption82.IsUnknown() {
		unset["-inject_dhcp_option_82"] = ""
	} else {
		inject_dhcp_option_82 := injectDhcpOption82TerraformToSdk(ctx, &diags, plan.InjectDhcpOption82)
		data.InjectDhcpOption82 = inject_dhcp_option_82
	}

	if plan.Interface.IsNull() || plan.Interface.IsUnknown() {
		unset["-interface"] = ""
	} else {
		data.Interface = models.ToPointer(models.WlanInterfaceEnum(string(plan.Interface.ValueString())))
	}

	if plan.Isolation.IsNull() || plan.Isolation.IsUnknown() {
		unset["-isolation"] = ""
	} else {
		data.Isolation = plan.Isolation.ValueBoolPointer()
	}

	if plan.L2Isolation.IsNull() || plan.L2Isolation.IsUnknown() {
		unset["-l2_isolation"] = ""
	} else {
		data.L2Isolation = plan.L2Isolation.ValueBoolPointer()
	}

	if plan.LegacyOverds.IsNull() || plan.LegacyOverds.IsUnknown() {
		unset["-legacy_overds"] = ""
	} else {
		data.LegacyOverds = plan.LegacyOverds.ValueBoolPointer()
	}

	if plan.LimitBcast.IsNull() || plan.LimitBcast.IsUnknown() {
		unset["-limit_bcast"] = ""
	} else {
		data.LimitBcast = plan.LimitBcast.ValueBoolPointer()
	}

	if plan.LimitProbeResponse.IsNull() || plan.LimitProbeResponse.IsUnknown() {
		unset["-limit_probe_response"] = ""
	} else {
		data.LimitProbeResponse = plan.LimitProbeResponse.ValueBoolPointer()
	}

	if plan.MaxIdletime.IsNull() || plan.MaxIdletime.IsUnknown() {
		unset["-max_idletime"] = ""
	} else {
		data.MaxIdletime = models.ToPointer(int(plan.MaxIdletime.ValueInt64()))
	}

	if plan.MaxNumClients.IsNull() || plan.MaxNumClients.IsUnknown() {
		unset["-max_num_clients"] = ""
	} else {
		data.MaxNumClients = models.ToPointer(int(plan.MaxNumClients.ValueInt64()))
	}

	if plan.MistNac.IsNull() || plan.MistNac.IsUnknown() {
		unset["-mist_nac"] = ""
	} else {
		mist_nac := mistNacTerraformToSdk(ctx, &diags, plan.MistNac)
		data.MistNac = mist_nac
	}

	if plan.MxtunnelIds.IsNull() || plan.MxtunnelIds.IsUnknown() {
		unset["-mxtunnel_ids"] = ""
	} else {
		data.MxtunnelIds = mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelIds)
	}

	if plan.MxtunnelName.IsNull() || plan.MxtunnelName.IsUnknown() {
		unset["-mxtunnel_name"] = ""
	} else {
		data.MxtunnelName = mist_transform.ListOfStringTerraformToSdk(ctx, plan.MxtunnelName)
	}

	if plan.NoStaticDns.IsNull() || plan.NoStaticDns.IsUnknown() {
		unset["-no_static_dns"] = ""
	} else {
		data.NoStaticDns = plan.NoStaticDns.ValueBoolPointer()
	}

	if plan.NoStaticIp.IsNull() || plan.NoStaticIp.IsUnknown() {
		unset["-no_static_ip"] = ""
	} else {
		data.NoStaticIp = plan.NoStaticIp.ValueBoolPointer()
	}

	if plan.Portal.IsNull() || plan.Portal.IsUnknown() {
		unset["-portal"] = ""
	} else {
		portal := portalTerraformToSdk(ctx, &diags, plan.Portal)
		data.Portal = portal
	}

	if plan.PortalAllowedHostnames.IsNull() || plan.PortalAllowedHostnames.IsUnknown() {
		unset["-portal_allowed_hostnames"] = ""
	} else {
		data.PortalAllowedHostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedHostnames)
	}

	if plan.PortalAllowedSubnets.IsNull() || plan.PortalAllowedSubnets.IsUnknown() {
		unset["-portal_allowed_subnets"] = ""
	} else {
		data.PortalAllowedSubnets = mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalAllowedSubnets)
	}

	if plan.PortalDeniedHostnames.IsNull() || plan.PortalDeniedHostnames.IsUnknown() {
		unset["-portal_denied_hostnames"] = ""
	} else {
		data.PortalDeniedHostnames = mist_transform.ListOfStringTerraformToSdk(ctx, plan.PortalDeniedHostnames)
	}

	if plan.Qos.IsNull() || plan.Qos.IsUnknown() {
		unset["-qos"] = ""
	} else {
		qos := qosTerraformToSdk(ctx, &diags, plan.Qos)
		data.Qos = qos
	}

	if plan.Radsec.IsNull() || plan.Radsec.IsUnknown() {
		unset["-radsec"] = ""
	} else {
		radesc := radsecTerraformToSdk(ctx, &diags, plan.Radsec)
		data.Radsec = radesc
	}

	if plan.Rateset.IsNull() || plan.Rateset.IsUnknown() {
		unset["-rateset"] = ""
	} else {
		data.Rateset = ratesetTerraformToSdk(ctx, &diags, plan.Rateset)
	}

	if plan.ReconnectClientsWhenRoamingMxcluster.IsNull() || plan.ReconnectClientsWhenRoamingMxcluster.IsUnknown() {
		unset["-reconnect_clients_when_roaming_mxcluster"] = ""
	} else {
		data.ReconnectClientsWhenRoamingMxcluster = plan.ReconnectClientsWhenRoamingMxcluster.ValueBoolPointer()
	}

	if plan.RoamMode.IsNull() || plan.RoamMode.IsUnknown() {
		unset["-roam_mode"] = ""
	} else {
		data.RoamMode = models.ToPointer(models.WlanRoamModeEnum(string(plan.RoamMode.ValueString())))
	}

	if plan.Schedule.IsNull() || plan.Schedule.IsUnknown() {
		unset["-schedule"] = ""
	} else {
		schedule := scheduleTerraformToSdk(ctx, &diags, plan.Schedule)
		data.Schedule = schedule
	}

	if plan.SleExcluded.IsNull() || plan.SleExcluded.IsUnknown() {
		unset["-sle_excluded"] = ""
	} else {
		data.SleExcluded = plan.SleExcluded.ValueBoolPointer()
	}

	if plan.UseEapolV1.IsNull() || plan.UseEapolV1.IsUnknown() {
		unset["-use_eapol_v1"] = ""
	} else {
		data.UseEapolV1 = plan.UseEapolV1.ValueBoolPointer()
	}

	if plan.VlanEnabled.IsNull() || plan.VlanEnabled.IsUnknown() {
		unset["-vlan_enabled"] = ""
	} else {
		data.VlanEnabled = plan.VlanEnabled.ValueBoolPointer()
	}

	if plan.VlanId.IsNull() || plan.VlanId.IsUnknown() {
		unset["-vlan_id"] = ""
	} else {
		data.VlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.VlanId.ValueString()))
	}

	if plan.VlanIds.IsNull() || plan.VlanIds.IsUnknown() {
		unset["-vlan_ids"] = ""
	} else {
		var items []models.VlanIdWithVariable
		for _, item := range plan.VlanIds.Elements() {
			var item_interface interface{} = item
			i := item_interface.(basetypes.StringValue)
			v := models.VlanIdWithVariableContainer.FromString(i.ValueString())
			items = append(items, v)
		}
		data.VlanIds = items
	}

	if plan.VlanPooling.IsNull() || plan.VlanPooling.IsUnknown() {
		unset["-vlan_pooling"] = ""
	} else {
		data.VlanPooling = plan.VlanPooling.ValueBoolPointer()
	}

	if plan.WlanLimitDown.IsNull() || plan.WlanLimitDown.IsUnknown() {
		unset["-wlan_limit_down"] = ""
	} else {
		data.WlanLimitDown = models.NewOptional(models.ToPointer(int(plan.WlanLimitDown.ValueInt64())))
	}

	if plan.WlanLimitDownEnabled.IsNull() || plan.WlanLimitDownEnabled.IsUnknown() {
		unset["-wlan_limit_down_enabled"] = ""
	} else {
		data.WlanLimitDownEnabled = plan.WlanLimitDownEnabled.ValueBoolPointer()
	}

	if plan.WlanLimitUp.IsNull() || plan.WlanLimitUp.IsUnknown() {
		unset["-wlan_limit_up"] = ""
	} else {
		data.WlanLimitUp = models.NewOptional(models.ToPointer(int(plan.WlanLimitUp.ValueInt64())))
	}

	if plan.WlanLimitUpEnabled.IsNull() || plan.WlanLimitUpEnabled.IsUnknown() {
		unset["-wlan_limit_up_enabled"] = ""
	} else {
		data.WlanLimitUpEnabled = plan.WlanLimitUpEnabled.ValueBoolPointer()
	}

	if plan.WxtagIds.IsNull() || plan.WxtagIds.IsUnknown() {
		unset["-wxtag_ids"] = ""
	} else {
		data.WxtagIds = models.NewOptional(models.ToPointer(mist_transform.ListOfUuidTerraformToSdk(ctx, plan.WxtagIds)))
	}

	if plan.WxtunnelId.IsNull() || plan.WxtunnelId.IsUnknown() {
		unset["-wxtunnel_id"] = ""
	} else {
		data.WxtunnelId = models.NewOptional(plan.WxtunnelId.ValueStringPointer())
	}

	if plan.WxtunnelRemoteId.IsNull() || plan.WxtunnelRemoteId.IsUnknown() {
		unset["-wxtunnel_remote_id"] = ""
	} else {
		data.WxtunnelRemoteId = models.NewOptional(plan.WxtunnelRemoteId.ValueStringPointer())
	}

	data.AdditionalProperties = unset

	return &data, diags
}
