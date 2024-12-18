package resource_org_wlan

import (
	"context"
	"strings"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *models.Wlan) (OrgWlanModel, diag.Diagnostics) {
	var state OrgWlanModel
	var diags diag.Diagnostics

	var acct_immediate_update types.Bool = types.BoolValue(false)
	var acct_interim_interval types.Int64
	var acct_servers types.List = types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})
	var airwatch AirwatchValue = NewAirwatchValueNull()
	var allow_ipv6_ndp types.Bool = types.BoolValue(true)
	var allow_mdns types.Bool = types.BoolValue(false)
	var allow_ssdp types.Bool = types.BoolValue(false)
	var ap_ids types.List = types.ListNull(types.StringType)
	var app_limit AppLimitValue = NewAppLimitValueNull()
	var app_qos AppQosValue = NewAppQosValueNull()
	var apply_to types.String
	var arp_filter types.Bool
	var auth AuthValue = NewAuthValueNull()
	var auth_server_selection types.String
	var auth_servers types.List = types.ListValueMust(AuthServersValue{}.Type(ctx), []attr.Value{})
	var auth_servers_nas_id types.String
	var auth_servers_nas_ip types.String
	var auth_servers_retries types.Int64
	var auth_servers_timeout types.Int64
	var band_steer types.Bool
	var band_steer_force_band5 types.Bool
	var bands types.List = types.ListNull(types.StringType)
	var block_blacklist_clients types.Bool
	var bonjour BonjourValue = NewBonjourValueNull()
	var cisco_cwa CiscoCwaValue = NewCiscoCwaValueNull()
	var client_limit_down types.Int64
	var client_limit_down_enabled types.Bool
	var client_limit_up types.Int64
	var client_limit_up_enabled types.Bool
	var coa_servers types.List = types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})
	var disable_11ax types.Bool
	var disable_ht_vht_rates types.Bool
	var disable_uapsd types.Bool
	var disable_v1_roam_notify types.Bool
	var disable_v2_roam_notify types.Bool
	var disable_when_gateway_unreachable types.Bool
	var disable_when_mxtunnel_down types.Bool
	var disable_wmm types.Bool
	var dns_server_rewrite DnsServerRewriteValue = NewDnsServerRewriteValueNull()
	var dtim types.Int64
	var dynamic_psk DynamicPskValue = NewDynamicPskValueNull()
	var dynamic_vlan DynamicVlanValue = NewDynamicVlanValueNull()
	var enable_local_keycaching types.Bool
	var enable_wireless_bridging types.Bool
	var enable_wireless_bridging_dhcp_tracking types.Bool
	var enabled types.Bool
	var fast_dot1x_timers types.Bool
	var hide_ssid types.Bool
	var hostname_ie types.Bool
	var hotspot20 Hotspot20Value = NewHotspot20ValueNull()
	var id types.String
	var inject_dhcp_option_82 InjectDhcpOption82Value = NewInjectDhcpOption82ValueNull()
	var interface_wlan types.String
	var isolation types.Bool
	var l2_isolation types.Bool
	var legacy_overds types.Bool
	var limit_bcast types.Bool
	var limit_probe_response types.Bool
	var max_idletime types.Int64
	var max_num_clients types.Int64
	var mist_nac MistNacValue = NewMistNacValueNull()
	var msp_id types.String = types.StringValue("")
	var mxtunnel_ids types.List = types.ListNull(types.StringType)
	var mxtunnel_name types.List = types.ListNull(types.StringType)
	var no_static_dns types.Bool
	var no_static_ip types.Bool
	var org_id types.String
	var portal PortalValue = NewPortalValueNull()
	var portal_allowed_hostnames types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_allowed_subnets types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_api_secret types.String = types.StringValue("")
	var portal_denied_hostnames types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_image types.String = types.StringValue("not_present")
	var portal_sso_url types.String = types.StringValue("")
	var qos QosValue
	var radsec RadsecValue = NewRadsecValueNull()
	var rateset types.Map = types.MapNull(RatesetValue{}.Type(ctx))
	var reconnect_clients_when_roaming_mxcluster types.Bool
	var roam_mode types.String
	var schedule ScheduleValue = NewScheduleValueNull()
	var sle_excluded types.Bool
	var ssid types.String
	var template_id types.String
	var use_eapol_v1 types.Bool
	var vlan_enabled types.Bool
	var vlan_id types.String
	var vlan_ids types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var vlan_pooling types.Bool
	var wlan_limit_down types.Int64
	var wlan_limit_down_enabled types.Bool
	var wlan_limit_up types.Int64
	var wlan_limit_up_enabled types.Bool
	var wxtag_ids types.List = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var wxtunnel_id types.String = types.StringValue("")
	var wxtunnel_remote_id types.String = types.StringValue("")

	if data.AcctImmediateUpdate != nil {
		acct_immediate_update = types.BoolValue(*data.AcctImmediateUpdate)
	}

	if data.AcctInterimInterval != nil {
		acct_interim_interval = types.Int64Value(int64(*data.AcctInterimInterval))
	}

	if len(data.AcctServers) > 0 {
		acct_servers = radiusServersAcctSdkToTerraform(ctx, &diags, data.AcctServers)
	} else {
		types.ListValueMust(AcctServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.Airwatch != nil {
		airwatch = airwatchSdkToTerraform(ctx, &diags, data.Airwatch)
	}

	if data.AllowIpv6Ndp != nil {
		allow_ipv6_ndp = types.BoolValue(*data.AllowIpv6Ndp)
	}

	if data.AllowMdns != nil {
		allow_mdns = types.BoolValue(*data.AllowMdns)
	}

	if data.AllowSsdp != nil {
		allow_ssdp = types.BoolValue(*data.AllowSsdp)
	}

	if data.ApIds.IsValueSet() && len(*data.ApIds.Value()) > 0 {
		ap_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, *data.ApIds.Value())
	}

	if data.AppLimit != nil {
		app_limit = appLimitSdkToTerraform(ctx, &diags, data.AppLimit)
	}

	if data.AppQos != nil {
		app_qos = appQosSdkToTerraform(ctx, &diags, *data.AppQos)
	}

	if data.ApplyTo != nil {
		apply_to = types.StringValue(string(*data.ApplyTo))
	}

	if data.ArpFilter != nil {
		arp_filter = types.BoolValue(*data.ArpFilter)
	}

	if data.Auth != nil {
		auth = authSdkToTerraform(ctx, &diags, data.Auth)
	}

	if data.AuthServerSelection != nil {
		auth_server_selection = types.StringValue(string(*data.AuthServerSelection))
	}

	if len(data.AuthServers) > 0 {
		auth_servers = radiusServersAuthSdkToTerraform(ctx, &diags, data.AuthServers)
	} else {
		auth_servers = types.ListValueMust(AuthServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.AuthServersNasId.IsValueSet() && data.AuthServersNasId.Value() != nil {
		auth_servers_nas_id = types.StringValue(*data.AuthServersNasId.Value())
	}

	if data.AuthServersNasIp.IsValueSet() && data.AuthServersNasIp.Value() != nil {
		auth_servers_nas_ip = types.StringValue(*data.AuthServersNasIp.Value())
	}

	if data.AuthServersRetries != nil {
		auth_servers_retries = types.Int64Value(int64(*data.AuthServersRetries))
	}

	if data.AuthServersTimeout != nil {
		auth_servers_timeout = types.Int64Value(int64(*data.AuthServersTimeout))
	}

	if data.BandSteer != nil {
		band_steer = types.BoolValue(*data.BandSteer)
	}

	if data.BandSteerForceBand5 != nil {
		band_steer_force_band5 = types.BoolValue(*data.BandSteerForceBand5)
	}

	if data.Bands != nil {
		bands = bandsSdkToTerraform(ctx, &diags, data.Bands)
	}

	if data.BlockBlacklistClients != nil {
		block_blacklist_clients = types.BoolValue(*data.BlockBlacklistClients)
	}

	if data.Bonjour != nil {
		bonjour = bonjourSdkToTerraform(ctx, &diags, data.Bonjour)
	}

	if data.CiscoCwa != nil {
		cisco_cwa = ciscoCwaSdkToTerraform(ctx, &diags, data.CiscoCwa)
	}

	if data.ClientLimitDown != nil {
		client_limit_down = types.Int64Value(int64(*data.ClientLimitDown))
	}

	if data.ClientLimitDownEnabled != nil {
		client_limit_down_enabled = types.BoolValue(*data.ClientLimitDownEnabled)
	}

	if data.ClientLimitUp != nil {
		client_limit_up = types.Int64Value(int64(*data.ClientLimitUp))
	}

	if data.ClientLimitUpEnabled != nil {
		client_limit_up_enabled = types.BoolValue(*data.ClientLimitUpEnabled)
	}

	if len(data.CoaServers) > 0 {
		coa_servers = coaServersSdkToTerraform(ctx, &diags, data.CoaServers)
	} else {
		coa_servers = types.ListValueMust(CoaServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.Disable11ax != nil {
		disable_11ax = types.BoolValue(*data.Disable11ax)
	}

	if data.DisableHtVhtRates != nil {
		disable_ht_vht_rates = types.BoolValue(*data.DisableHtVhtRates)
	}

	if data.DisableUapsd != nil {
		disable_uapsd = types.BoolValue(*data.DisableUapsd)
	}

	if data.DisableV1RoamNotify != nil {
		disable_v1_roam_notify = types.BoolValue(*data.DisableV1RoamNotify)
	}

	if data.DisableV2RoamNotify != nil {
		disable_v2_roam_notify = types.BoolValue(*data.DisableV2RoamNotify)
	}

	if data.DisableWhenGatewayUnreachable != nil {
		disable_when_gateway_unreachable = types.BoolValue(*data.DisableWhenGatewayUnreachable)
	}

	if data.DisableWhenMxtunnelDown != nil {
		disable_when_mxtunnel_down = types.BoolValue(*data.DisableWhenMxtunnelDown)
	}

	if data.DisableWmm != nil {
		disable_wmm = types.BoolValue(*data.DisableWmm)
	}

	if data.DnsServerRewrite.IsValueSet() && data.DnsServerRewrite.Value() != nil {
		dns_server_rewrite = dnsServerRewriteSdkToTerraform(ctx, &diags, data.DnsServerRewrite.Value())
	}

	if data.Dtim != nil {
		dtim = types.Int64Value(int64(*data.Dtim))
	}

	if data.DynamicPsk.IsValueSet() && data.DynamicPsk.Value() != nil {
		dynamic_psk = dynamicPskSdkToTerraform(ctx, &diags, data.DynamicPsk.Value())
	}

	if data.DynamicVlan.IsValueSet() && data.DynamicVlan.Value() != nil {
		dynamic_vlan = dynamicVlanSdkToTerraform(ctx, &diags, data.DynamicVlan.Value())
	}

	if data.EnableLocalKeycaching != nil {
		enable_local_keycaching = types.BoolValue(*data.EnableLocalKeycaching)
	}

	if data.EnableWirelessBridging != nil {
		enable_wireless_bridging = types.BoolValue(*data.EnableWirelessBridging)
	}

	if data.EnableWirelessBridgingDhcpTracking != nil {
		enable_wireless_bridging_dhcp_tracking = types.BoolValue(*data.EnableWirelessBridgingDhcpTracking)
	}

	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	if data.FastDot1xTimers != nil {
		fast_dot1x_timers = types.BoolValue(*data.FastDot1xTimers)
	}

	if data.HideSsid != nil {
		hide_ssid = types.BoolValue(*data.HideSsid)
	}

	if data.HostnameIe != nil {
		hostname_ie = types.BoolValue(*data.HostnameIe)
	}

	if data.Hotspot20 != nil {
		hotspot20 = hotspot20SdkToTerraform(ctx, &diags, data.Hotspot20)
	}

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}

	if data.InjectDhcpOption82 != nil {
		inject_dhcp_option_82 = injectDhcpOption82dkToTerraform(ctx, &diags, data.InjectDhcpOption82)
	}

	if data.Interface != nil {
		interface_wlan = types.StringValue(string(*data.Interface))
	}

	if data.Isolation != nil {
		isolation = types.BoolValue(*data.Isolation)
	}

	if data.L2Isolation != nil {
		l2_isolation = types.BoolValue(*data.L2Isolation)
	}

	if data.LegacyOverds != nil {
		legacy_overds = types.BoolValue(*data.LegacyOverds)
	}

	if data.LimitBcast != nil {
		limit_bcast = types.BoolValue(*data.LimitBcast)
	}

	if data.LimitProbeResponse != nil {
		limit_probe_response = types.BoolValue(*data.LimitProbeResponse)
	}

	if data.MaxIdletime != nil {
		max_idletime = types.Int64Value(int64(*data.MaxIdletime))
	}

	if data.MaxNumClients != nil {
		max_num_clients = types.Int64Value(int64(*data.MaxNumClients))
	}

	if data.MistNac != nil {
		mist_nac = mistNacdSkToTerraform(ctx, &diags, data.MistNac)
	}

	if data.MspId != nil {
		msp_id = types.StringValue(data.MspId.String())
	}

	if len(data.MxtunnelIds) > 0 {
		mxtunnel_ids = mist_transform.ListOfStringSdkToTerraform(ctx, data.MxtunnelIds)
	}

	if len(data.MxtunnelName) > 0 {
		mxtunnel_name = mist_transform.ListOfStringSdkToTerraform(ctx, data.MxtunnelName)
	}

	if data.NoStaticDns != nil {
		no_static_dns = types.BoolValue(*data.NoStaticDns)
	}

	if data.NoStaticIp != nil {
		no_static_ip = types.BoolValue(*data.NoStaticIp)
	}

	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}

	if data.Portal != nil {
		portal = portalSkToTerraform(ctx, &diags, data.Portal)
	}

	if data.PortalAllowedHostnames != nil {
		portal_allowed_hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalAllowedHostnames)
	}
	if data.PortalAllowedSubnets != nil {
		portal_allowed_subnets = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalAllowedSubnets)
	}

	if data.PortalApiSecret.IsValueSet() && data.PortalApiSecret.Value() != nil {
		portal_api_secret = types.StringValue(*data.PortalApiSecret.Value())
	}

	if data.PortalDeniedHostnames != nil {
		portal_denied_hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalDeniedHostnames)
	}

	if data.PortalImage.IsValueSet() && data.PortalImage.Value() != nil {
		portal_image = types.StringValue("present")
	}

	if data.PortalSsoUrl.IsValueSet() && data.PortalSsoUrl.Value() != nil {
		portal_sso_url = types.StringValue(*data.PortalSsoUrl.Value())
	}

	if data.Qos != nil {
		qos = qosSkToTerraform(ctx, &diags, data.Qos)
	}

	if data.Radsec != nil {
		radsec = radsecSkToTerraform(ctx, &diags, data.Radsec)
	}

	if data.Rateset != nil {
		rateset = ratesetSkToTerraform(ctx, &diags, data.Rateset)
	}

	if data.ReconnectClientsWhenRoamingMxcluster != nil {
		reconnect_clients_when_roaming_mxcluster = types.BoolValue(*data.ReconnectClientsWhenRoamingMxcluster)
	}

	if data.RoamMode != nil {
		roam_mode = types.StringValue(string(*data.RoamMode))
	}

	if data.Schedule != nil {
		schedule = scheduleSkToTerraform(ctx, &diags, data.Schedule)
	}

	if data.SleExcluded != nil {
		sle_excluded = types.BoolValue(*data.SleExcluded)
	}

	ssid = types.StringValue(data.Ssid)

	if data.TemplateId.IsValueSet() && data.TemplateId.Value() != nil {
		template_id = types.StringValue(data.TemplateId.Value().String())
	}

	if data.UseEapolV1 != nil {
		use_eapol_v1 = types.BoolValue(*data.UseEapolV1)
	}

	if data.VlanEnabled != nil {
		vlan_enabled = types.BoolValue(*data.VlanEnabled)
	}

	if data.VlanId != nil {
		vlan_id = types.StringValue(string(data.VlanId.String()))
	}

	if data.VlanIds != nil {
		var list []attr.Value
		if vlan_ids_as_string, ok := data.VlanIds.AsString(); ok {
			for _, vlan := range strings.Split(*vlan_ids_as_string, ",") {
				list = append(list, types.StringValue(vlan))
			}
		} else if vlan_ids_as_list, ok := data.VlanIds.AsArrayOfVlanIdWithVariable2(); ok {
			for _, v := range *vlan_ids_as_list {
				list = append(list, types.StringValue(v.String()))
			}
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		vlan_ids = r
	}

	if data.VlanPooling != nil {
		vlan_pooling = types.BoolValue(*data.VlanPooling)
	}

	if data.WlanLimitDown.IsValueSet() && data.WlanLimitDown.Value() != nil {
		wlan_limit_down = types.Int64Value(int64(*data.WlanLimitDown.Value()))
	}

	if data.WlanLimitDownEnabled != nil {
		wlan_limit_down_enabled = types.BoolValue(*data.WlanLimitDownEnabled)
	}

	if data.WlanLimitUp.IsValueSet() && data.WlanLimitUp.Value() != nil {
		wlan_limit_up = types.Int64Value(int64(*data.WlanLimitUp.Value()))
	}

	if data.WlanLimitUpEnabled != nil {
		wlan_limit_up_enabled = types.BoolValue(*data.WlanLimitUpEnabled)
	}

	if data.WxtagIds.IsValueSet() && data.WxtagIds.Value() != nil {
		wxtag_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, *data.WxtagIds.Value())
	}

	if data.WxtunnelId.IsValueSet() && data.WxtunnelId.Value() != nil {
		wxtunnel_id = types.StringValue(*data.WxtunnelId.Value())
	}

	if data.WxtunnelRemoteId.IsValueSet() && data.WxtunnelRemoteId.Value() != nil {
		wxtunnel_remote_id = types.StringValue(*data.WxtunnelRemoteId.Value())
	}

	state.AcctImmediateUpdate = acct_immediate_update
	state.AcctInterimInterval = acct_interim_interval
	state.AcctServers = acct_servers
	state.Airwatch = airwatch
	state.AllowIpv6Ndp = allow_ipv6_ndp
	state.AllowMdns = allow_mdns
	state.AllowSsdp = allow_ssdp
	state.ApIds = ap_ids
	state.AppLimit = app_limit
	state.AppQos = app_qos
	state.ApplyTo = apply_to
	state.ArpFilter = arp_filter
	state.Auth = auth
	state.AuthServerSelection = auth_server_selection
	state.AuthServers = auth_servers
	state.AuthServersNasId = auth_servers_nas_id
	state.AuthServersNasIp = auth_servers_nas_ip
	state.AuthServersRetries = auth_servers_retries
	state.AuthServersTimeout = auth_servers_timeout
	state.BandSteer = band_steer
	state.BandSteerForceBand5 = band_steer_force_band5
	state.Bands = bands
	state.BlockBlacklistClients = block_blacklist_clients
	state.Bonjour = bonjour
	state.CiscoCwa = cisco_cwa
	state.ClientLimitDown = client_limit_down
	state.ClientLimitDownEnabled = client_limit_down_enabled
	state.ClientLimitUp = client_limit_up
	state.ClientLimitUpEnabled = client_limit_up_enabled
	state.CoaServers = coa_servers
	state.Disable11ax = disable_11ax
	state.DisableHtVhtRates = disable_ht_vht_rates
	state.DisableUapsd = disable_uapsd
	state.DisableV1RoamNotify = disable_v1_roam_notify
	state.DisableV2RoamNotify = disable_v2_roam_notify
	state.DisableWhenGatewayUnreachable = disable_when_gateway_unreachable
	state.DisableWhenMxtunnelDown = disable_when_mxtunnel_down
	state.DisableWmm = disable_wmm
	state.DnsServerRewrite = dns_server_rewrite
	state.Dtim = dtim
	state.DynamicPsk = dynamic_psk
	state.DynamicVlan = dynamic_vlan
	state.EnableLocalKeycaching = enable_local_keycaching
	state.EnableWirelessBridging = enable_wireless_bridging
	state.EnableWirelessBridgingDhcpTracking = enable_wireless_bridging_dhcp_tracking
	state.Enabled = enabled
	state.FastDot1xTimers = fast_dot1x_timers
	state.HideSsid = hide_ssid
	state.HostnameIe = hostname_ie
	state.Hotspot20 = hotspot20
	state.Id = id
	state.InjectDhcpOption82 = inject_dhcp_option_82
	state.Interface = interface_wlan
	state.Isolation = isolation
	state.L2Isolation = l2_isolation
	state.LegacyOverds = legacy_overds
	state.LimitBcast = limit_bcast
	state.LimitProbeResponse = limit_probe_response
	state.MaxIdletime = max_idletime
	state.MaxNumClients = max_num_clients
	state.MistNac = mist_nac
	state.MspId = msp_id
	state.MxtunnelIds = mxtunnel_ids
	state.MxtunnelName = mxtunnel_name
	state.NoStaticDns = no_static_dns
	state.NoStaticIp = no_static_ip
	state.OrgId = org_id
	state.Portal = portal
	state.PortalAllowedHostnames = portal_allowed_hostnames
	state.PortalAllowedSubnets = portal_allowed_subnets
	state.PortalApiSecret = portal_api_secret
	state.PortalDeniedHostnames = portal_denied_hostnames
	state.PortalImage = portal_image
	state.PortalSsoUrl = portal_sso_url
	state.Qos = qos
	state.Radsec = radsec
	state.Rateset = rateset
	state.ReconnectClientsWhenRoamingMxcluster = reconnect_clients_when_roaming_mxcluster
	state.RoamMode = roam_mode
	state.Schedule = schedule
	state.SleExcluded = sle_excluded
	state.Ssid = ssid
	state.TemplateId = template_id
	state.UseEapolV1 = use_eapol_v1
	state.VlanEnabled = vlan_enabled
	state.VlanId = vlan_id
	state.VlanIds = vlan_ids
	state.VlanPooling = vlan_pooling
	state.WlanLimitDown = wlan_limit_down
	state.WlanLimitDownEnabled = wlan_limit_down_enabled
	state.WlanLimitUp = wlan_limit_up
	state.WlanLimitUpEnabled = wlan_limit_up_enabled
	state.WxtagIds = wxtag_ids
	state.WxtunnelId = wxtunnel_id
	state.WxtunnelRemoteId = wxtunnel_remote_id

	return state, diags
}
