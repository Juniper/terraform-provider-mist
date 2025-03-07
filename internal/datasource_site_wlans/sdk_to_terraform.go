package datasource_site_wlans

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

func SdkToTerraform(ctx context.Context, l *[]models.Wlan, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := wlanSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func wlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Wlan) SiteWlansValue {

	var acctImmediateUpdate = types.BoolValue(false)
	var acctInterimInterval basetypes.Int64Value
	var acctServers = types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})
	var airwatch = types.ObjectNull(AirwatchValue{}.AttributeTypes(ctx))
	var allowIpv6Ndp = types.BoolValue(true)
	var allowMdns = types.BoolValue(false)
	var allowSsdp = types.BoolValue(false)
	var apIds = types.ListNull(types.StringType)
	var appLimit = types.ObjectNull(AppLimitValue{}.AttributeTypes(ctx))
	var appQos = types.ObjectNull(AppQosValue{}.AttributeTypes(ctx))
	var applyTo basetypes.StringValue
	var arpFilter basetypes.BoolValue
	var auth = types.ObjectNull(AuthValue{}.AttributeTypes(ctx))
	var authServerSelection basetypes.StringValue
	var authServers = types.ListValueMust(AuthServersValue{}.Type(ctx), []attr.Value{})
	var authServersNasId basetypes.StringValue
	var authServersNasIp basetypes.StringValue
	var authServersRetries basetypes.Int64Value
	var authServersTimeout basetypes.Int64Value
	var bandSteer basetypes.BoolValue
	var bandSteerForceBand5 basetypes.BoolValue
	var bands = types.ListNull(types.StringType)
	var blockBlacklistClients basetypes.BoolValue
	var bonjour = types.ObjectNull(BonjourValue{}.AttributeTypes(ctx))
	var ciscoCwa = types.ObjectNull(CiscoCwaValue{}.AttributeTypes(ctx))
	var clientLimitDown basetypes.Int64Value
	var clientLimitDownEnabled basetypes.BoolValue
	var clientLimitUp basetypes.Int64Value
	var clientLimitUpEnabled basetypes.BoolValue
	var createdTime basetypes.Float64Value
	var coaServers = types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})
	var disable11ax basetypes.BoolValue
	var disable11be basetypes.BoolValue
	var disableHtVhtRates basetypes.BoolValue
	var disableUapsd basetypes.BoolValue
	var disableV1RoamNotify basetypes.BoolValue
	var disableV2RoamNotify basetypes.BoolValue
	var disableWhenGatewayUnreachable basetypes.BoolValue
	var disableWhenMxtunnelDown basetypes.BoolValue
	var disableWmm basetypes.BoolValue
	var dnsServerRewrite = types.ObjectNull(DnsServerRewriteValue{}.AttributeTypes(ctx))
	var dtim basetypes.Int64Value
	var dynamicPsk = types.ObjectNull(DynamicPskValue{}.AttributeTypes(ctx))
	var dynamicVlan = types.ObjectNull(DynamicVlanValue{}.AttributeTypes(ctx))
	var enableLocalKeycaching basetypes.BoolValue
	var enableWirelessBridging basetypes.BoolValue
	var enableWirelessBridgingDhcpTracking basetypes.BoolValue
	var enabled basetypes.BoolValue
	var fastDot1xTimers basetypes.BoolValue
	var hideSsid basetypes.BoolValue
	var hostnameIe basetypes.BoolValue
	var hotspot20 = types.ObjectNull(Hotspot20Value{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var injectDhcpOption82 = types.ObjectNull(InjectDhcpOption82Value{}.AttributeTypes(ctx))
	var interfaceWlan basetypes.StringValue
	var isolation basetypes.BoolValue
	var l2Isolation basetypes.BoolValue
	var legacyOverds basetypes.BoolValue
	var limitBcast basetypes.BoolValue
	var limitProbeResponse basetypes.BoolValue
	var maxIdletime basetypes.Int64Value
	var maxNumClients basetypes.Int64Value
	var mistNac = types.ObjectNull(MistNacValue{}.AttributeTypes(ctx))
	var modifiedTime basetypes.Float64Value
	var mspId = types.StringValue("")
	var mxtunnelIds = types.ListNull(types.StringType)
	var mxtunnelName = types.ListNull(types.StringType)
	var noStaticDns basetypes.BoolValue
	var noStaticIp basetypes.BoolValue
	var orgId basetypes.StringValue
	var portal = types.ObjectNull(PortalValue{}.AttributeTypes(ctx))
	var portalAllowedHostnames = mistutils.ListOfStringSdkToTerraformEmpty()
	var portalAllowedSubnets = mistutils.ListOfStringSdkToTerraformEmpty()
	var portalApiSecret = types.StringValue("")
	var portalDeniedHostnames = mistutils.ListOfStringSdkToTerraformEmpty()
	var portalImage = types.StringValue("not_present")
	var portalSsoUrl = types.StringValue("")
	var qos = types.ObjectNull(QosValue{}.AttributeTypes(ctx))
	var radsec = types.ObjectNull(RadsecValue{}.AttributeTypes(ctx))
	var rateset = types.MapNull(RatesetValue{}.Type(ctx))
	var reconnectClientsWhenRoamingMxcluster basetypes.BoolValue
	var roamMode basetypes.StringValue
	var schedule = types.ObjectNull(ScheduleValue{}.AttributeTypes(ctx))
	var siteId = types.StringValue("")
	var sleExcluded basetypes.BoolValue
	var ssid basetypes.StringValue
	var useEapolV1 basetypes.BoolValue
	var vlanEnabled basetypes.BoolValue
	var vlanId basetypes.StringValue
	var vlanIds = mistutils.ListOfStringSdkToTerraformEmpty()
	var vlanPooling basetypes.BoolValue
	var wlanLimitDown basetypes.Int64Value
	var wlanLimitDownEnabled basetypes.BoolValue
	var wlanLimitUp basetypes.Int64Value
	var wlanLimitUpEnabled basetypes.BoolValue
	var wxtagIds = mistutils.ListOfUuidSdkToTerraformEmpty()
	var wxtunnelId = types.StringValue("")
	var wxtunnelRemoteId = types.StringValue("")

	if d.AcctImmediateUpdate != nil {
		acctImmediateUpdate = types.BoolValue(*d.AcctImmediateUpdate)
	}

	if d.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*d.AcctInterimInterval))
	}

	if len(d.AcctServers) > 0 {
		acctServers = radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	} else {
		types.ListValueMust(AcctServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.Airwatch != nil {
		airwatch = airwatchSdkToTerraform(ctx, diags, d.Airwatch)
	}

	if d.AllowIpv6Ndp != nil {
		allowIpv6Ndp = types.BoolValue(*d.AllowIpv6Ndp)
	}

	if d.AllowMdns != nil {
		allowMdns = types.BoolValue(*d.AllowMdns)
	}

	if d.AllowSsdp != nil {
		allowSsdp = types.BoolValue(*d.AllowSsdp)
	}

	if d.ApIds.Value() != nil && len(*d.ApIds.Value()) > 0 {
		apIds = mistutils.ListOfUuidSdkToTerraform(*d.ApIds.Value())
	}

	if d.AppLimit != nil {
		appLimit = appLimitSdkToTerraform(ctx, diags, d.AppLimit)
	}

	if d.AppQos != nil {
		appQos = appQosSdkToTerraform(ctx, diags, *d.AppQos)
	}

	if d.ApplyTo != nil {
		applyTo = types.StringValue(string(*d.ApplyTo))
	}

	if d.ArpFilter != nil {
		arpFilter = types.BoolValue(*d.ArpFilter)
	}

	if d.Auth != nil {
		auth = authSdkToTerraform(ctx, diags, d.Auth)
	}

	if d.AuthServerSelection != nil {
		authServerSelection = types.StringValue(string(*d.AuthServerSelection))
	}

	if len(d.AuthServers) > 0 {
		authServers = radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	} else {
		authServers = types.ListValueMust(AuthServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.AuthServersNasId.IsValueSet() && d.AuthServersNasId.Value() != nil {
		authServersNasId = types.StringValue(*d.AuthServersNasId.Value())
	}

	if d.AuthServersNasIp.IsValueSet() && d.AuthServersNasIp.Value() != nil {
		authServersNasIp = types.StringValue(*d.AuthServersNasIp.Value())
	}

	if d.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*d.AuthServersRetries))
	}

	if d.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}

	if d.BandSteer != nil {
		bandSteer = types.BoolValue(*d.BandSteer)
	}

	if d.BandSteerForceBand5 != nil {
		bandSteerForceBand5 = types.BoolValue(*d.BandSteerForceBand5)
	}

	if d.Bands != nil {
		bands = bandsSdkToTerraform(ctx, diags, d.Bands)
	}

	if d.BlockBlacklistClients != nil {
		blockBlacklistClients = types.BoolValue(*d.BlockBlacklistClients)
	}

	if d.Bonjour != nil {
		bonjour = bonjourSdkToTerraform(ctx, diags, d.Bonjour)
	}

	if d.CiscoCwa != nil {
		ciscoCwa = ciscoCwaSdkToTerraform(ctx, diags, d.CiscoCwa)
	}

	if d.ClientLimitDown != nil {
		clientLimitDown = types.Int64Value(int64(*d.ClientLimitDown))
	}

	if d.ClientLimitDownEnabled != nil {
		clientLimitDownEnabled = types.BoolValue(*d.ClientLimitDownEnabled)
	}

	if d.ClientLimitUp != nil {
		clientLimitUp = types.Int64Value(int64(*d.ClientLimitUp))
	}

	if d.ClientLimitUpEnabled != nil {
		clientLimitUpEnabled = types.BoolValue(*d.ClientLimitUpEnabled)
	}

	if len(d.CoaServers) > 0 {
		coaServers = coaServersSdkToTerraform(ctx, diags, d.CoaServers)
	} else {
		coaServers = types.ListValueMust(CoaServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}

	if d.Disable11ax != nil {
		disable11ax = types.BoolValue(*d.Disable11ax)
	}

	if d.Disable11be != nil {
		disable11be = types.BoolValue(*d.Disable11be)
	}

	if d.DisableHtVhtRates != nil {
		disableHtVhtRates = types.BoolValue(*d.DisableHtVhtRates)
	}

	if d.DisableUapsd != nil {
		disableUapsd = types.BoolValue(*d.DisableUapsd)
	}

	if d.DisableV1RoamNotify != nil {
		disableV1RoamNotify = types.BoolValue(*d.DisableV1RoamNotify)
	}

	if d.DisableV2RoamNotify != nil {
		disableV2RoamNotify = types.BoolValue(*d.DisableV2RoamNotify)
	}

	if d.DisableWhenGatewayUnreachable != nil {
		disableWhenGatewayUnreachable = types.BoolValue(*d.DisableWhenGatewayUnreachable)
	}

	if d.DisableWhenMxtunnelDown != nil {
		disableWhenMxtunnelDown = types.BoolValue(*d.DisableWhenMxtunnelDown)
	}

	if d.DisableWmm != nil {
		disableWmm = types.BoolValue(*d.DisableWmm)
	}

	if d.DnsServerRewrite.IsValueSet() && d.DnsServerRewrite.Value() != nil {
		dnsServerRewrite = dnsServerRewriteSdkToTerraform(ctx, diags, d.DnsServerRewrite.Value())
	}

	if d.Dtim != nil {
		dtim = types.Int64Value(int64(*d.Dtim))
	}

	if d.DynamicPsk.IsValueSet() && d.DynamicPsk.Value() != nil {
		dynamicPsk = dynamicPskSdkToTerraform(ctx, diags, d.DynamicPsk.Value())
	}

	if d.DynamicVlan.IsValueSet() && d.DynamicVlan.Value() != nil {
		dynamicVlan = dynamicVlanSdkToTerraform(ctx, diags, d.DynamicVlan.Value())
	}

	if d.EnableLocalKeycaching != nil {
		enableLocalKeycaching = types.BoolValue(*d.EnableLocalKeycaching)
	}

	if d.EnableWirelessBridging != nil {
		enableWirelessBridging = types.BoolValue(*d.EnableWirelessBridging)
	}

	if d.EnableWirelessBridgingDhcpTracking != nil {
		enableWirelessBridgingDhcpTracking = types.BoolValue(*d.EnableWirelessBridgingDhcpTracking)
	}

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	if d.FastDot1xTimers != nil {
		fastDot1xTimers = types.BoolValue(*d.FastDot1xTimers)
	}

	if d.HideSsid != nil {
		hideSsid = types.BoolValue(*d.HideSsid)
	}

	if d.HostnameIe != nil {
		hostnameIe = types.BoolValue(*d.HostnameIe)
	}

	if d.Hotspot20 != nil {
		hotspot20 = hotspot20SdkToTerraform(ctx, diags, d.Hotspot20)
	}

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	if d.InjectDhcpOption82 != nil {
		injectDhcpOption82 = injectDhcpOption82dkToTerraform(ctx, diags, d.InjectDhcpOption82)
	}

	if d.Interface != nil {
		interfaceWlan = types.StringValue(string(*d.Interface))
	}

	if d.Isolation != nil {
		isolation = types.BoolValue(*d.Isolation)
	}

	if d.L2Isolation != nil {
		l2Isolation = types.BoolValue(*d.L2Isolation)
	}

	if d.LegacyOverds != nil {
		legacyOverds = types.BoolValue(*d.LegacyOverds)
	}

	if d.LimitBcast != nil {
		limitBcast = types.BoolValue(*d.LimitBcast)
	}

	if d.LimitProbeResponse != nil {
		limitProbeResponse = types.BoolValue(*d.LimitProbeResponse)
	}

	if d.MaxIdletime != nil {
		maxIdletime = types.Int64Value(int64(*d.MaxIdletime))
	}

	if d.MaxNumClients != nil {
		maxNumClients = types.Int64Value(int64(*d.MaxNumClients))
	}

	if d.MistNac != nil {
		mistNac = mistNacdSkToTerraform(ctx, diags, d.MistNac)
	}

	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}

	if d.MspId != nil {
		mspId = types.StringValue(d.MspId.String())
	}

	if len(d.MxtunnelIds) > 0 {
		mxtunnelIds = mistutils.ListOfStringSdkToTerraform(d.MxtunnelIds)
	}

	if len(d.MxtunnelName) > 0 {
		mxtunnelName = mistutils.ListOfStringSdkToTerraform(d.MxtunnelName)
	}

	if d.NoStaticDns != nil {
		noStaticDns = types.BoolValue(*d.NoStaticDns)
	}

	if d.NoStaticIp != nil {
		noStaticIp = types.BoolValue(*d.NoStaticIp)
	}

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}

	if d.Portal != nil {
		portal = portalSkToTerraform(ctx, diags, d.Portal)
	}

	if d.PortalAllowedHostnames != nil {
		portalAllowedHostnames = mistutils.ListOfStringSdkToTerraform(d.PortalAllowedHostnames)
	}
	if d.PortalAllowedSubnets != nil {
		portalAllowedSubnets = mistutils.ListOfStringSdkToTerraform(d.PortalAllowedSubnets)
	}

	if d.PortalApiSecret.IsValueSet() && d.PortalApiSecret.Value() != nil {
		portalApiSecret = types.StringValue(*d.PortalApiSecret.Value())
	}

	if d.PortalDeniedHostnames != nil {
		portalDeniedHostnames = mistutils.ListOfStringSdkToTerraform(d.PortalDeniedHostnames)
	}

	if d.PortalImage.IsValueSet() && d.PortalImage.Value() != nil {
		portalImage = types.StringValue("present")
	}

	if d.PortalSsoUrl.IsValueSet() && d.PortalSsoUrl.Value() != nil {
		portalSsoUrl = types.StringValue(*d.PortalSsoUrl.Value())
	}

	if d.Qos != nil {
		qos = qosSkToTerraform(ctx, diags, d.Qos)
	}

	if d.Radsec != nil {
		radsec = radsecSkToTerraform(ctx, diags, d.Radsec)
	}

	if d.Rateset != nil {
		rateset = ratesetSkToTerraform(ctx, diags, d.Rateset)
	}

	if d.ReconnectClientsWhenRoamingMxcluster != nil {
		reconnectClientsWhenRoamingMxcluster = types.BoolValue(*d.ReconnectClientsWhenRoamingMxcluster)
	}

	if d.RoamMode != nil {
		roamMode = types.StringValue(string(*d.RoamMode))
	}

	if d.Schedule != nil {
		schedule = scheduleSkToTerraform(ctx, diags, d.Schedule)
	}

	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}

	if d.SleExcluded != nil {
		sleExcluded = types.BoolValue(*d.SleExcluded)
	}

	ssid = types.StringValue(d.Ssid)

	if d.UseEapolV1 != nil {
		useEapolV1 = types.BoolValue(*d.UseEapolV1)
	}

	if d.VlanEnabled != nil {
		vlanEnabled = types.BoolValue(*d.VlanEnabled)
	}

	if d.VlanId != nil {
		vlanId = mistutils.VlanAsString(*d.VlanId)
	}

	if d.VlanIds != nil {
		var list []attr.Value
		if vlanIdsAsString, ok := d.VlanIds.AsString(); ok {
			for _, vlan := range strings.Split(*vlanIdsAsString, ",") {
				list = append(list, types.StringValue(vlan))
			}
		} else if vlanIdsAsList, ok := d.VlanIds.AsArrayOfVlanIdWithVariable2(); ok {
			for _, v := range *vlanIdsAsList {
				list = append(list, mistutils.VlanAsString(v))
			}
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		vlanIds = r
	}

	if d.VlanPooling != nil {
		vlanPooling = types.BoolValue(*d.VlanPooling)
	}

	if d.WlanLimitDown.IsValueSet() && d.WlanLimitDown.Value() != nil {
		wlanLimitDown = types.Int64Value(int64(*d.WlanLimitDown.Value()))
	}

	if d.WlanLimitDownEnabled != nil {
		wlanLimitDownEnabled = types.BoolValue(*d.WlanLimitDownEnabled)
	}

	if d.WlanLimitUp.IsValueSet() && d.WlanLimitUp.Value() != nil {
		wlanLimitUp = types.Int64Value(int64(*d.WlanLimitUp.Value()))
	}

	if d.WlanLimitUpEnabled != nil {
		wlanLimitUpEnabled = types.BoolValue(*d.WlanLimitUpEnabled)
	}

	if d.WxtagIds.IsValueSet() && d.WxtagIds.Value() != nil {
		wxtagIds = mistutils.ListOfUuidSdkToTerraform(*d.WxtagIds.Value())
	}

	if d.WxtunnelId.IsValueSet() && d.WxtunnelId.Value() != nil {
		wxtunnelId = types.StringValue(*d.WxtunnelId.Value())
	}

	if d.WxtunnelRemoteId.IsValueSet() && d.WxtunnelRemoteId.Value() != nil {
		wxtunnelRemoteId = types.StringValue(*d.WxtunnelRemoteId.Value())
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
		"coa_servers":                            coaServers,
		"created_time":                           createdTime,
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
		"site_id":                 siteId,
		"sle_excluded":            sleExcluded,
		"ssid":                    ssid,
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

	data, e := NewSiteWlansValue(SiteWlansValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
