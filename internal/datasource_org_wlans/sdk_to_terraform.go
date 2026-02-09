package datasource_org_wlans

import (
	"context"
	"strings"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *[]models.Wlan, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics
	for _, item := range *data {
		elem := wlanSdkToTerraform(ctx, &diags, &item)
		*elements = append(*elements, elem)
	}

	return diags
}

func wlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Wlan) OrgWlansValue {
	if data == nil {
		return OrgWlansValue{}
	}

	var acctImmediateUpdate = types.BoolValue(false)
	if data.AcctImmediateUpdate != nil {
		acctImmediateUpdate = types.BoolValue(*data.AcctImmediateUpdate)
	}

	var acctInterimInterval basetypes.Int64Value
	if data.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*data.AcctInterimInterval))
	}

	var acctServers = types.ListValueMust(AcctServersValue{}.Type(ctx), make([]attr.Value, 0))
	if len(data.AcctServers) > 0 {
		acctServers = radiusServersAcctSdkToTerraform(ctx, diags, data.AcctServers)
	}

	var airwatch = types.ObjectNull(AirwatchValue{}.AttributeTypes(ctx))
	if data.Airwatch != nil {
		airwatch = airwatchSdkToTerraform(ctx, diags, data.Airwatch)
	}

	var allowIpv6Ndp = types.BoolValue(true)
	if data.AllowIpv6Ndp != nil {
		allowIpv6Ndp = types.BoolValue(*data.AllowIpv6Ndp)
	}

	var allowMdns = types.BoolValue(false)
	if data.AllowMdns != nil {
		allowMdns = types.BoolValue(*data.AllowMdns)
	}

	var allowSsdp = types.BoolValue(false)
	if data.AllowSsdp != nil {
		allowSsdp = types.BoolValue(*data.AllowSsdp)
	}

	var apIds = types.ListNull(types.StringType)
	if data.ApIds.IsValueSet() && data.ApIds.Value() != nil {
		apIds = mistutils.ListOfUuidSdkToTerraform(*data.ApIds.Value())
	}

	var appLimit = types.ObjectNull(AppLimitValue{}.AttributeTypes(ctx))
	if data.AppLimit != nil {
		appLimit = appLimitSdkToTerraform(ctx, diags, data.AppLimit)
	}

	var appQos = types.ObjectNull(AppQosValue{}.AttributeTypes(ctx))
	if data.AppQos != nil {
		appQos = appQosSdkToTerraform(ctx, diags, *data.AppQos)
	}

	var applyTo basetypes.StringValue
	if data.ApplyTo != nil {
		applyTo = types.StringValue(string(*data.ApplyTo))
	}

	var arpFilter basetypes.BoolValue
	if data.ArpFilter != nil {
		arpFilter = types.BoolValue(*data.ArpFilter)
	}

	var auth = types.ObjectNull(AuthValue{}.AttributeTypes(ctx))
	if data.Auth != nil {
		auth = authSdkToTerraform(ctx, diags, data.Auth)
	}

	var authServerSelection basetypes.StringValue
	if data.AuthServerSelection != nil {
		authServerSelection = types.StringValue(string(*data.AuthServerSelection))
	}

	var authServers = types.ListValueMust(AuthServersValue{}.Type(ctx), make([]attr.Value, 0))
	if len(data.AuthServers) > 0 {
		authServers = radiusServersAuthSdkToTerraform(ctx, diags, data.AuthServers)
	}

	var authServersNasId basetypes.StringValue
	if data.AuthServersNasId.IsValueSet() && data.AuthServersNasId.Value() != nil {
		authServersNasId = types.StringValue(*data.AuthServersNasId.Value())
	}

	var authServersNasIp basetypes.StringValue
	if data.AuthServersNasIp.IsValueSet() && data.AuthServersNasIp.Value() != nil {
		authServersNasIp = types.StringValue(*data.AuthServersNasIp.Value())
	}

	var authServersRetries basetypes.Int64Value
	if data.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*data.AuthServersRetries))
	}

	var authServersTimeout basetypes.Int64Value
	if data.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*data.AuthServersTimeout))
	}

	var bandSteer basetypes.BoolValue
	if data.BandSteer != nil {
		bandSteer = types.BoolValue(*data.BandSteer)
	}

	var bandSteerForceBand5 basetypes.BoolValue
	if data.BandSteerForceBand5 != nil {
		bandSteerForceBand5 = types.BoolValue(*data.BandSteerForceBand5)
	}

	var bands = types.ListNull(types.StringType)
	if data.Bands != nil {
		bands = bandsSdkToTerraform(ctx, diags, data.Bands)
	}

	var blockBlacklistClients basetypes.BoolValue
	if data.BlockBlacklistClients != nil {
		blockBlacklistClients = types.BoolValue(*data.BlockBlacklistClients)
	}

	var bonjour = types.ObjectNull(BonjourValue{}.AttributeTypes(ctx))
	if data.Bonjour != nil {
		bonjour = bonjourSdkToTerraform(ctx, diags, data.Bonjour)
	}

	var ciscoCwa = types.ObjectNull(CiscoCwaValue{}.AttributeTypes(ctx))
	if data.CiscoCwa != nil {
		ciscoCwa = ciscoCwaSdkToTerraform(ctx, diags, data.CiscoCwa)
	}

	var clientLimitDown basetypes.StringValue
	if data.ClientLimitDown != nil {
		clientLimitDown = mistutils.WlanLimitAsString(data.ClientLimitDown)
	}

	var clientLimitDownEnabled basetypes.BoolValue
	if data.ClientLimitDownEnabled != nil {
		clientLimitDownEnabled = types.BoolValue(*data.ClientLimitDownEnabled)
	}

	var clientLimitUp basetypes.StringValue
	if data.ClientLimitUp != nil {
		clientLimitUp = mistutils.WlanLimitAsString(data.ClientLimitUp)
	}

	var clientLimitUpEnabled basetypes.BoolValue
	if data.ClientLimitUpEnabled != nil {
		clientLimitUpEnabled = types.BoolValue(*data.ClientLimitUpEnabled)
	}

	var coaServers = types.ListValueMust(CoaServersValue{}.Type(ctx), make([]attr.Value, 0))
	if len(data.CoaServers) > 0 {
		coaServers = coaServersSdkToTerraform(ctx, diags, data.CoaServers)
	}

	var createdTime basetypes.Float64Value
	if data.CreatedTime != nil {
		createdTime = types.Float64Value(*data.CreatedTime)
	}

	var disable11ax basetypes.BoolValue
	if data.Disable11ax != nil {
		disable11ax = types.BoolValue(*data.Disable11ax)
	}

	var disable11be basetypes.BoolValue
	if data.Disable11be != nil {
		disable11be = types.BoolValue(*data.Disable11be)
	}

	var disableHtVhtRates basetypes.BoolValue
	if data.DisableHtVhtRates != nil {
		disableHtVhtRates = types.BoolValue(*data.DisableHtVhtRates)
	}

	var disableUapsd basetypes.BoolValue
	if data.DisableUapsd != nil {
		disableUapsd = types.BoolValue(*data.DisableUapsd)
	}

	var disableV1RoamNotify basetypes.BoolValue
	if data.DisableV1RoamNotify != nil {
		disableV1RoamNotify = types.BoolValue(*data.DisableV1RoamNotify)
	}

	var disableV2RoamNotify basetypes.BoolValue
	if data.DisableV2RoamNotify != nil {
		disableV2RoamNotify = types.BoolValue(*data.DisableV2RoamNotify)
	}

	var disableWhenGatewayUnreachable basetypes.BoolValue
	if data.DisableWhenGatewayUnreachable != nil {
		disableWhenGatewayUnreachable = types.BoolValue(*data.DisableWhenGatewayUnreachable)
	}

	var disableWhenMxtunnelDown basetypes.BoolValue
	if data.DisableWhenMxtunnelDown != nil {
		disableWhenMxtunnelDown = types.BoolValue(*data.DisableWhenMxtunnelDown)
	}

	var disableWmm basetypes.BoolValue
	if data.DisableWmm != nil {
		disableWmm = types.BoolValue(*data.DisableWmm)
	}

	var dnsServerRewrite = types.ObjectNull(DnsServerRewriteValue{}.AttributeTypes(ctx))
	if data.DnsServerRewrite.IsValueSet() && data.DnsServerRewrite.Value() != nil {
		dnsServerRewrite = dnsServerRewriteSdkToTerraform(ctx, diags, data.DnsServerRewrite.Value())
	}

	var dtim basetypes.Int64Value
	if data.Dtim != nil {
		dtim = types.Int64Value(int64(*data.Dtim))
	}

	var dynamicPsk = types.ObjectNull(DynamicPskValue{}.AttributeTypes(ctx))
	if data.DynamicPsk.IsValueSet() && data.DynamicPsk.Value() != nil {
		dynamicPsk = dynamicPskSdkToTerraform(ctx, diags, data.DynamicPsk.Value())
	}

	var dynamicVlan = types.ObjectNull(DynamicVlanValue{}.AttributeTypes(ctx))
	if data.DynamicVlan.IsValueSet() && data.DynamicVlan.Value() != nil {
		dynamicVlan = dynamicVlanSdkToTerraform(ctx, diags, data.DynamicVlan.Value())
	}

	var enableLocalKeycaching basetypes.BoolValue
	if data.EnableLocalKeycaching != nil {
		enableLocalKeycaching = types.BoolValue(*data.EnableLocalKeycaching)
	}

	var enableWirelessBridging basetypes.BoolValue
	if data.EnableWirelessBridging != nil {
		enableWirelessBridging = types.BoolValue(*data.EnableWirelessBridging)
	}

	var enableWirelessBridgingDhcpTracking basetypes.BoolValue
	if data.EnableWirelessBridgingDhcpTracking != nil {
		enableWirelessBridgingDhcpTracking = types.BoolValue(*data.EnableWirelessBridgingDhcpTracking)
	}

	var enabled basetypes.BoolValue
	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	var fastDot1xTimers basetypes.BoolValue
	if data.FastDot1xTimers != nil {
		fastDot1xTimers = types.BoolValue(*data.FastDot1xTimers)
	}

	var hideSsid basetypes.BoolValue
	if data.HideSsid != nil {
		hideSsid = types.BoolValue(*data.HideSsid)
	}

	var hostnameIe basetypes.BoolValue
	if data.HostnameIe != nil {
		hostnameIe = types.BoolValue(*data.HostnameIe)
	}

	var hotspot20 = types.ObjectNull(Hotspot20Value{}.AttributeTypes(ctx))
	if data.Hotspot20 != nil {
		hotspot20 = hotspot20SdkToTerraform(ctx, diags, data.Hotspot20)
	}

	var id basetypes.StringValue
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}

	var injectDhcpOption82 = types.ObjectNull(InjectDhcpOption82Value{}.AttributeTypes(ctx))
	if data.InjectDhcpOption82 != nil {
		injectDhcpOption82 = injectDhcpOption82dkToTerraform(ctx, diags, data.InjectDhcpOption82)
	}

	var interfaceWlan basetypes.StringValue
	if data.Interface != nil {
		interfaceWlan = types.StringValue(string(*data.Interface))
	}

	var isolation basetypes.BoolValue
	if data.Isolation != nil {
		isolation = types.BoolValue(*data.Isolation)
	}

	var l2Isolation basetypes.BoolValue
	if data.L2Isolation != nil {
		l2Isolation = types.BoolValue(*data.L2Isolation)
	}

	var legacyOverds basetypes.BoolValue
	if data.LegacyOverds != nil {
		legacyOverds = types.BoolValue(*data.LegacyOverds)
	}

	var limitBcast basetypes.BoolValue
	if data.LimitBcast != nil {
		limitBcast = types.BoolValue(*data.LimitBcast)
	}

	var limitProbeResponse basetypes.BoolValue
	if data.LimitProbeResponse != nil {
		limitProbeResponse = types.BoolValue(*data.LimitProbeResponse)
	}

	var maxIdletime basetypes.Int64Value
	if data.MaxIdletime != nil {
		maxIdletime = types.Int64Value(int64(*data.MaxIdletime))
	}

	var maxNumClients basetypes.Int64Value
	if data.MaxNumClients != nil {
		maxNumClients = types.Int64Value(int64(*data.MaxNumClients))
	}

	var mistNac = types.ObjectNull(MistNacValue{}.AttributeTypes(ctx))
	if data.MistNac != nil {
		mistNac = mistNacSkToTerraform(ctx, diags, data.MistNac)
	}

	var modifiedTime basetypes.Float64Value
	if data.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*data.ModifiedTime)
	}

	var mspId = types.StringValue("")
	if data.MspId != nil {
		mspId = types.StringValue(data.MspId.String())
	}

	var mxtunnelIds = types.ListNull(types.StringType)
	if len(data.MxtunnelIds) > 0 {
		mxtunnelIds = mistutils.ListOfStringSdkToTerraform(data.MxtunnelIds)
	}

	var mxtunnelName = types.ListNull(types.StringType)
	if len(data.MxtunnelName) > 0 {
		mxtunnelName = mistutils.ListOfStringSdkToTerraform(data.MxtunnelName)
	}

	var noStaticDns basetypes.BoolValue
	if data.NoStaticDns != nil {
		noStaticDns = types.BoolValue(*data.NoStaticDns)
	}

	var noStaticIp basetypes.BoolValue
	if data.NoStaticIp != nil {
		noStaticIp = types.BoolValue(*data.NoStaticIp)
	}

	var orgId basetypes.StringValue
	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}

	var portal = types.ObjectNull(PortalValue{}.AttributeTypes(ctx))
	if data.Portal != nil {
		portal = portalSkToTerraform(ctx, diags, data.Portal)
	}

	var portalAllowedHostnames = mistutils.ListOfStringSdkToTerraformEmpty()
	if data.PortalAllowedHostnames != nil {
		portalAllowedHostnames = mistutils.ListOfStringSdkToTerraform(data.PortalAllowedHostnames)
	}

	var portalAllowedSubnets = mistutils.ListOfStringSdkToTerraformEmpty()
	if data.PortalAllowedSubnets != nil {
		portalAllowedSubnets = mistutils.ListOfStringSdkToTerraform(data.PortalAllowedSubnets)
	}

	var portalApiSecret = types.StringValue("")
	if data.PortalApiSecret.IsValueSet() && data.PortalApiSecret.Value() != nil {
		portalApiSecret = types.StringValue(*data.PortalApiSecret.Value())
	}

	var portalDeniedHostnames = mistutils.ListOfStringSdkToTerraformEmpty()
	if data.PortalDeniedHostnames != nil {
		portalDeniedHostnames = mistutils.ListOfStringSdkToTerraform(data.PortalDeniedHostnames)
	}

	var portalImage = types.StringValue("not_present")
	if data.PortalImage.IsValueSet() && data.PortalImage.Value() != nil {
		portalImage = types.StringValue("present")
	}

	var portalSsoUrl = types.StringValue("")
	if data.PortalSsoUrl.IsValueSet() && data.PortalSsoUrl.Value() != nil {
		portalSsoUrl = types.StringValue(*data.PortalSsoUrl.Value())
	}

	var qos = types.ObjectNull(QosValue{}.AttributeTypes(ctx))
	if data.Qos != nil {
		qos = qosSkToTerraform(ctx, diags, data.Qos)
	}

	var radsec = types.ObjectNull(RadsecValue{}.AttributeTypes(ctx))
	if data.Radsec != nil {
		radsec = radsecSkToTerraform(ctx, diags, data.Radsec)
	}

	var rateset = types.MapNull(RatesetValue{}.Type(ctx))
	if data.Rateset != nil {
		rateset = ratesetSkToTerraform(ctx, diags, data.Rateset)
	}

	var reconnectClientsWhenRoamingMxcluster basetypes.BoolValue
	if data.ReconnectClientsWhenRoamingMxcluster != nil {
		reconnectClientsWhenRoamingMxcluster = types.BoolValue(*data.ReconnectClientsWhenRoamingMxcluster)
	}

	var roamMode basetypes.StringValue
	if data.RoamMode != nil {
		switch strings.ToLower(string(*data.RoamMode)) {
		case "none", "okc":
			roamMode = types.StringValue(strings.ToUpper(string(*data.RoamMode)))
		case "11r":
			roamMode = types.StringValue("11r")
		default:
			roamMode = types.StringValue(string(*data.RoamMode))
		}
	}

	var schedule = types.ObjectNull(ScheduleValue{}.AttributeTypes(ctx))
	if data.Schedule != nil {
		schedule = scheduleSkToTerraform(ctx, diags, data.Schedule)
	}

	var sleExcluded basetypes.BoolValue
	if data.SleExcluded != nil {
		sleExcluded = types.BoolValue(*data.SleExcluded)
	}

	ssid := types.StringValue(data.Ssid)

	templateId := types.StringValue(data.TemplateId.Value().String())

	var useEapolV1 basetypes.BoolValue
	if data.UseEapolV1 != nil {
		useEapolV1 = types.BoolValue(*data.UseEapolV1)
	}

	var vlanEnabled basetypes.BoolValue
	if data.VlanEnabled != nil {
		vlanEnabled = types.BoolValue(*data.VlanEnabled)
	}

	var vlanId basetypes.StringValue
	if data.VlanId.IsValueSet() && data.VlanId.Value() != nil {
		vlanId = mistutils.WlanVlanAsString(*data.VlanId.Value())
	}

	var vlanIds = mistutils.ListOfStringSdkToTerraformEmpty()
	if data.VlanIds != nil {
		vlanIds = mistutils.WlanVlanIdsAsArrayOfString(diags, data.VlanIds)
	}

	var vlanPooling basetypes.BoolValue
	if data.VlanPooling != nil {
		vlanPooling = types.BoolValue(*data.VlanPooling)
	}

	var wlanLimitDown basetypes.StringValue
	if data.WlanLimitDown != nil {
		wlanLimitDown = mistutils.WlanLimitAsString(data.WlanLimitDown)
	}

	var wlanLimitDownEnabled basetypes.BoolValue
	if data.WlanLimitDownEnabled != nil {
		wlanLimitDownEnabled = types.BoolValue(*data.WlanLimitDownEnabled)
	}

	var wlanLimitUp basetypes.StringValue
	if data.WlanLimitUp != nil {
		wlanLimitUp = mistutils.WlanLimitAsString(data.WlanLimitUp)
	}

	var wlanLimitUpEnabled basetypes.BoolValue
	if data.WlanLimitUpEnabled != nil {
		wlanLimitUpEnabled = types.BoolValue(*data.WlanLimitUpEnabled)
	}

	var wxtagIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	if data.WxtagIds.IsValueSet() && data.WxtagIds.Value() != nil {
		wxtagIds = mistutils.ListOfUuidSdkToTerraform(*data.WxtagIds.Value())
	}

	var wxtunnelId = types.StringValue("")
	if data.WxtunnelId.IsValueSet() && data.WxtunnelId.Value() != nil {
		wxtunnelId = types.StringValue(*data.WxtunnelId.Value())
	}

	var wxtunnelRemoteId = types.StringValue("")
	if data.WxtunnelRemoteId.IsValueSet() && data.WxtunnelRemoteId.Value() != nil {
		wxtunnelRemoteId = types.StringValue(*data.WxtunnelRemoteId.Value())
	}

	dataMapValue := map[string]attr.Value{
		"acct_immediate_update":                  acctImmediateUpdate,
		"acct_interim_interval":                  acctInterimInterval,
		"acct_servers":                           acctServers,
		"airwatch":                               airwatch,
		"allow_ipv6_ndp":                         allowIpv6Ndp,
		"allow_mdns":                             allowMdns,
		"allow_ssdp":                             allowSsdp,
		"ap_ids":                                 apIds,
		"app_limit":                              appLimit,
		"app_qos":                                appQos,
		"apply_to":                               applyTo,
		"arp_filter":                             arpFilter,
		"auth":                                   auth,
		"auth_server_selection":                  authServerSelection,
		"auth_servers":                           authServers,
		"auth_servers_nas_id":                    authServersNasId,
		"auth_servers_nas_ip":                    authServersNasIp,
		"auth_servers_retries":                   authServersRetries,
		"auth_servers_timeout":                   authServersTimeout,
		"band_steer":                             bandSteer,
		"band_steer_force_band5":                 bandSteerForceBand5,
		"bands":                                  bands,
		"block_blacklist_clients":                blockBlacklistClients,
		"bonjour":                                bonjour,
		"cisco_cwa":                              ciscoCwa,
		"client_limit_down":                      clientLimitDown,
		"client_limit_down_enabled":              clientLimitDownEnabled,
		"client_limit_up":                        clientLimitUp,
		"client_limit_up_enabled":                clientLimitUpEnabled,
		"created_time":                           createdTime,
		"coa_servers":                            coaServers,
		"disable_11ax":                           disable11ax,
		"disable_11be":                           disable11be,
		"disable_ht_vht_rates":                   disableHtVhtRates,
		"disable_uapsd":                          disableUapsd,
		"disable_v1_roam_notify":                 disableV1RoamNotify,
		"disable_v2_roam_notify":                 disableV2RoamNotify,
		"disable_when_gateway_unreachable":       disableWhenGatewayUnreachable,
		"disable_when_mxtunnel_down":             disableWhenMxtunnelDown,
		"disable_wmm":                            disableWmm,
		"dns_server_rewrite":                     dnsServerRewrite,
		"dtim":                                   dtim,
		"dynamic_psk":                            dynamicPsk,
		"dynamic_vlan":                           dynamicVlan,
		"enable_local_keycaching":                enableLocalKeycaching,
		"enable_wireless_bridging":               enableWirelessBridging,
		"enable_wireless_bridging_dhcp_tracking": enableWirelessBridgingDhcpTracking,
		"enabled":                                enabled,
		"fast_dot1x_timers":                      fastDot1xTimers,
		"hide_ssid":                              hideSsid,
		"hostname_ie":                            hostnameIe,
		"hotspot20":                              hotspot20,
		"id":                                     id,
		"inject_dhcp_option_82":                  injectDhcpOption82,
		"interface":                              interfaceWlan,
		"isolation":                              isolation,
		"l2_isolation":                           l2Isolation,
		"legacy_overds":                          legacyOverds,
		"limit_bcast":                            limitBcast,
		"limit_probe_response":                   limitProbeResponse,
		"max_idletime":                           maxIdletime,
		"max_num_clients":                        maxNumClients,
		"mist_nac":                               mistNac,
		"modified_time":                          modifiedTime,
		"msp_id":                                 mspId,
		"mxtunnel_ids":                           mxtunnelIds,
		"mxtunnel_name":                          mxtunnelName,
		"no_static_dns":                          noStaticDns,
		"no_static_ip":                           noStaticIp,
		"org_id":                                 orgId,
		"portal":                                 portal,
		"portal_allowed_hostnames":               portalAllowedHostnames,
		"portal_allowed_subnets":                 portalAllowedSubnets,
		"portal_api_secret":                      portalApiSecret,
		"portal_denied_hostnames":                portalDeniedHostnames,
		"portal_image":                           portalImage,
		"portal_sso_url":                         portalSsoUrl,
		"qos":                                    qos,
		"radsec":                                 radsec,
		"rateset":                                rateset,
		"reconnect_clients_when_roaming_mxcluster": reconnectClientsWhenRoamingMxcluster,
		"roam_mode":               roamMode,
		"schedule":                schedule,
		"sle_excluded":            sleExcluded,
		"ssid":                    ssid,
		"template_id":             templateId,
		"use_eapol_v1":            useEapolV1,
		"vlan_enabled":            vlanEnabled,
		"vlan_id":                 vlanId,
		"vlan_ids":                vlanIds,
		"vlan_pooling":            vlanPooling,
		"wlan_limit_down":         wlanLimitDown,
		"wlan_limit_down_enabled": wlanLimitDownEnabled,
		"wlan_limit_up":           wlanLimitUp,
		"wlan_limit_up_enabled":   wlanLimitUpEnabled,
		"wxtag_ids":               wxtagIds,
		"wxtunnel_id":             wxtunnelId,
		"wxtunnel_remote_id":      wxtunnelRemoteId,
	}
	result, err := NewOrgWlansValue(OrgWlansValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(err...)

	return result
}
