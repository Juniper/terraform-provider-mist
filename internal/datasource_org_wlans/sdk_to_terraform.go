package datasource_org_wlans

import (
	"context"
	"strings"

	mist_api "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

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

func wlanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Wlan) OrgWlansValue {

	var acct_immediate_update basetypes.BoolValue = types.BoolValue(false)
	var acct_interim_interval basetypes.Int64Value
	var acct_servers basetypes.ListValue = types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})
	var airwatch basetypes.ObjectValue = types.ObjectNull(AirwatchValue{}.AttributeTypes(ctx))
	var allow_ipv6_ndp basetypes.BoolValue = types.BoolValue(true)
	var allow_mdns basetypes.BoolValue = types.BoolValue(false)
	var allow_ssdp basetypes.BoolValue = types.BoolValue(false)
	var ap_ids basetypes.ListValue = types.ListNull(types.StringType)
	var app_limit basetypes.ObjectValue = types.ObjectNull(AppLimitValue{}.AttributeTypes(ctx))
	var app_qos basetypes.ObjectValue = types.ObjectNull(AppQosValue{}.AttributeTypes(ctx))
	var apply_to basetypes.StringValue
	var arp_filter basetypes.BoolValue
	var auth basetypes.ObjectValue = types.ObjectNull(AuthValue{}.AttributeTypes(ctx))
	var auth_server_selection basetypes.StringValue
	var auth_servers basetypes.ListValue = types.ListValueMust(AuthServersValue{}.Type(ctx), []attr.Value{})
	var auth_servers_nas_id basetypes.StringValue
	var auth_servers_nas_ip basetypes.StringValue
	var auth_servers_retries basetypes.Int64Value
	var auth_servers_timeout basetypes.Int64Value
	var band_steer basetypes.BoolValue
	var band_steer_force_band5 basetypes.BoolValue
	var bands basetypes.ListValue = types.ListNull(types.StringType)
	var block_blacklist_clients basetypes.BoolValue
	var bonjour basetypes.ObjectValue = types.ObjectNull(BonjourValue{}.AttributeTypes(ctx))
	var cisco_cwa basetypes.ObjectValue = types.ObjectNull(CiscoCwaValue{}.AttributeTypes(ctx))
	var client_limit_down basetypes.Int64Value
	var client_limit_down_enabled basetypes.BoolValue
	var client_limit_up basetypes.Int64Value
	var client_limit_up_enabled basetypes.BoolValue
	var coa_servers basetypes.ListValue = types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})
	var created_time basetypes.Float64Value
	var disable_11ax basetypes.BoolValue
	var disable_ht_vht_rates basetypes.BoolValue
	var disable_uapsd basetypes.BoolValue
	var disable_v1_roam_notify basetypes.BoolValue
	var disable_v2_roam_notify basetypes.BoolValue
	var disable_when_gateway_unreachable basetypes.BoolValue
	var disable_when_mxtunnel_down basetypes.BoolValue
	var disable_wmm basetypes.BoolValue
	var dns_server_rewrite basetypes.ObjectValue = types.ObjectNull(DnsServerRewriteValue{}.AttributeTypes(ctx))
	var dtim basetypes.Int64Value
	var dynamic_psk basetypes.ObjectValue = types.ObjectNull(DynamicPskValue{}.AttributeTypes(ctx))
	var dynamic_vlan basetypes.ObjectValue = types.ObjectNull(DynamicVlanValue{}.AttributeTypes(ctx))
	var enable_local_keycaching basetypes.BoolValue
	var enable_wireless_bridging basetypes.BoolValue
	var enable_wireless_bridging_dhcp_tracking basetypes.BoolValue
	var enabled basetypes.BoolValue
	var fast_dot1x_timers basetypes.BoolValue
	var hide_ssid basetypes.BoolValue
	var hostname_ie basetypes.BoolValue
	var hotspot20 basetypes.ObjectValue = types.ObjectNull(Hotspot20Value{}.AttributeTypes(ctx))
	var id basetypes.StringValue
	var inject_dhcp_option_82 basetypes.ObjectValue = types.ObjectNull(InjectDhcpOption82Value{}.AttributeTypes(ctx))
	var interface_wlan basetypes.StringValue
	var isolation basetypes.BoolValue
	var l2_isolation basetypes.BoolValue
	var legacy_overds basetypes.BoolValue
	var limit_bcast basetypes.BoolValue
	var limit_probe_response basetypes.BoolValue
	var max_idletime basetypes.Int64Value
	var max_num_clients basetypes.Int64Value
	var mist_nac basetypes.ObjectValue = types.ObjectNull(MistNacValue{}.AttributeTypes(ctx))
	var modified_time basetypes.Float64Value
	var msp_id basetypes.StringValue = types.StringValue("")
	var mxtunnel_ids basetypes.ListValue = types.ListNull(types.StringType)
	var mxtunnel_name basetypes.ListValue = types.ListNull(types.StringType)
	var no_static_dns basetypes.BoolValue
	var no_static_ip basetypes.BoolValue
	var org_id basetypes.StringValue
	var portal basetypes.ObjectValue = types.ObjectNull(PortalValue{}.AttributeTypes(ctx))
	var portal_allowed_hostnames basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_allowed_subnets basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_api_secret basetypes.StringValue = types.StringValue("")
	var portal_denied_hostnames basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var portal_image basetypes.StringValue = types.StringValue("not_present")
	var portal_sso_url basetypes.StringValue = types.StringValue("")
	var qos basetypes.ObjectValue = types.ObjectNull(QosValue{}.AttributeTypes(ctx))
	var radsec basetypes.ObjectValue = types.ObjectNull(RadsecValue{}.AttributeTypes(ctx))
	var rateset basetypes.MapValue = types.MapNull(RatesetValue{}.Type(ctx))
	var reconnect_clients_when_roaming_mxcluster basetypes.BoolValue
	var roam_mode basetypes.StringValue
	var schedule basetypes.ObjectValue = types.ObjectNull(ScheduleValue{}.AttributeTypes(ctx))
	var sle_excluded basetypes.BoolValue
	var ssid basetypes.StringValue
	var template_id basetypes.StringValue
	var use_eapol_v1 basetypes.BoolValue
	var vlan_enabled basetypes.BoolValue
	var vlan_id basetypes.StringValue
	var vlan_ids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var vlan_pooling basetypes.BoolValue
	var wlan_limit_down basetypes.Int64Value
	var wlan_limit_down_enabled basetypes.BoolValue
	var wlan_limit_up basetypes.Int64Value
	var wlan_limit_up_enabled basetypes.BoolValue
	var wxtag_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var wxtunnel_id basetypes.StringValue = types.StringValue("")
	var wxtunnel_remote_id basetypes.StringValue = types.StringValue("")

	if d.AcctImmediateUpdate != nil {
		acct_immediate_update = types.BoolValue(*d.AcctImmediateUpdate)
	}

	if d.AcctInterimInterval != nil {
		acct_interim_interval = types.Int64Value(int64(*d.AcctInterimInterval))
	}

	if len(d.AcctServers) > 0 {
		acct_servers = radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	} else {
		types.ListValueMust(AcctServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.Airwatch != nil {
		airwatch = airwatchSdkToTerraform(ctx, diags, d.Airwatch)
	}

	if d.AllowIpv6Ndp != nil {
		allow_ipv6_ndp = types.BoolValue(*d.AllowIpv6Ndp)
	}

	if d.AllowMdns != nil {
		allow_mdns = types.BoolValue(*d.AllowMdns)
	}

	if d.AllowSsdp != nil {
		allow_ssdp = types.BoolValue(*d.AllowSsdp)
	}

	if d.ApIds.Value() != nil && len(*d.ApIds.Value()) > 0 {
		ap_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, *d.ApIds.Value())
	}

	if d.AppLimit != nil {
		app_limit = appLimitSdkToTerraform(ctx, diags, d.AppLimit)
	}

	if d.AppQos != nil {
		app_qos = appQosSdkToTerraform(ctx, diags, *d.AppQos)
	}

	if d.ApplyTo != nil {
		apply_to = types.StringValue(string(*d.ApplyTo))
	}

	if d.ArpFilter != nil {
		arp_filter = types.BoolValue(*d.ArpFilter)
	}

	if d.Auth != nil {
		auth = authSdkToTerraform(ctx, diags, d.Auth)
	}

	if d.AuthServerSelection != nil {
		auth_server_selection = types.StringValue(string(*d.AuthServerSelection))
	}

	if len(d.AuthServers) > 0 {
		auth_servers = radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	} else {
		auth_servers = types.ListValueMust(AuthServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.AuthServersNasId.IsValueSet() && d.AuthServersNasId.Value() != nil {
		auth_servers_nas_id = types.StringValue(*d.AuthServersNasId.Value())
	}

	if d.AuthServersNasIp.IsValueSet() && d.AuthServersNasIp.Value() != nil {
		auth_servers_nas_ip = types.StringValue(*d.AuthServersNasIp.Value())
	}

	if d.AuthServersRetries != nil {
		auth_servers_retries = types.Int64Value(int64(*d.AuthServersRetries))
	}

	if d.AuthServersTimeout != nil {
		auth_servers_timeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}

	if d.BandSteer != nil {
		band_steer = types.BoolValue(*d.BandSteer)
	}

	if d.BandSteerForceBand5 != nil {
		band_steer_force_band5 = types.BoolValue(*d.BandSteerForceBand5)
	}

	if d.Bands != nil {
		bands = bandsSdkToTerraform(ctx, diags, d.Bands)
	}

	if d.BlockBlacklistClients != nil {
		block_blacklist_clients = types.BoolValue(*d.BlockBlacklistClients)
	}

	if d.Bonjour != nil {
		bonjour = bonjourSdkToTerraform(ctx, diags, d.Bonjour)
	}

	if d.CiscoCwa != nil {
		cisco_cwa = ciscoCwaSdkToTerraform(ctx, diags, d.CiscoCwa)
	}

	if d.ClientLimitDown != nil {
		client_limit_down = types.Int64Value(int64(*d.ClientLimitDown))
	}

	if d.ClientLimitDownEnabled != nil {
		client_limit_down_enabled = types.BoolValue(*d.ClientLimitDownEnabled)
	}

	if d.ClientLimitUp != nil {
		client_limit_up = types.Int64Value(int64(*d.ClientLimitUp))
	}

	if d.ClientLimitUpEnabled != nil {
		client_limit_up_enabled = types.BoolValue(*d.ClientLimitUpEnabled)
	}

	if len(d.CoaServers) > 0 {
		coa_servers = coaServersSdkToTerraform(ctx, diags, d.CoaServers)
	} else {
		coa_servers = types.ListValueMust(CoaServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if d.CreatedTime != nil {
		created_time = types.Float64Value(*d.CreatedTime)
	}

	if d.Disable11ax != nil {
		disable_11ax = types.BoolValue(*d.Disable11ax)
	}

	if d.DisableHtVhtRates != nil {
		disable_ht_vht_rates = types.BoolValue(*d.DisableHtVhtRates)
	}

	if d.DisableUapsd != nil {
		disable_uapsd = types.BoolValue(*d.DisableUapsd)
	}

	if d.DisableV1RoamNotify != nil {
		disable_v1_roam_notify = types.BoolValue(*d.DisableV1RoamNotify)
	}

	if d.DisableV2RoamNotify != nil {
		disable_v2_roam_notify = types.BoolValue(*d.DisableV2RoamNotify)
	}

	if d.DisableWhenGatewayUnreachable != nil {
		disable_when_gateway_unreachable = types.BoolValue(*d.DisableWhenGatewayUnreachable)
	}

	if d.DisableWhenMxtunnelDown != nil {
		disable_when_mxtunnel_down = types.BoolValue(*d.DisableWhenMxtunnelDown)
	}

	if d.DisableWmm != nil {
		disable_wmm = types.BoolValue(*d.DisableWmm)
	}

	if d.DnsServerRewrite.IsValueSet() && d.DnsServerRewrite.Value() != nil {
		dns_server_rewrite = dnsServerRewriteSdkToTerraform(ctx, diags, d.DnsServerRewrite.Value())
	}

	if d.Dtim != nil {
		dtim = types.Int64Value(int64(*d.Dtim))
	}

	if d.DynamicPsk.IsValueSet() && d.DynamicPsk.Value() != nil {
		dynamic_psk = dynamicPskSdkToTerraform(ctx, diags, d.DynamicPsk.Value())
	}

	if d.DynamicVlan.IsValueSet() && d.DynamicVlan.Value() != nil {
		dynamic_vlan = dynamicVlanSdkToTerraform(ctx, diags, d.DynamicVlan.Value())
	}

	if d.EnableLocalKeycaching != nil {
		enable_local_keycaching = types.BoolValue(*d.EnableLocalKeycaching)
	}

	if d.EnableWirelessBridging != nil {
		enable_wireless_bridging = types.BoolValue(*d.EnableWirelessBridging)
	}

	if d.EnableWirelessBridgingDhcpTracking != nil {
		enable_wireless_bridging_dhcp_tracking = types.BoolValue(*d.EnableWirelessBridgingDhcpTracking)
	}

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	if d.FastDot1xTimers != nil {
		fast_dot1x_timers = types.BoolValue(*d.FastDot1xTimers)
	}

	if d.HideSsid != nil {
		hide_ssid = types.BoolValue(*d.HideSsid)
	}

	if d.HostnameIe != nil {
		hostname_ie = types.BoolValue(*d.HostnameIe)
	}

	if d.Hotspot20 != nil {
		hotspot20 = hotspot20SdkToTerraform(ctx, diags, d.Hotspot20)
	}

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	if d.InjectDhcpOption82 != nil {
		inject_dhcp_option_82 = injectDhcpOption82dkToTerraform(ctx, diags, d.InjectDhcpOption82)
	}

	if d.Interface != nil {
		interface_wlan = types.StringValue(string(*d.Interface))
	}

	if d.Isolation != nil {
		isolation = types.BoolValue(*d.Isolation)
	}

	if d.L2Isolation != nil {
		l2_isolation = types.BoolValue(*d.L2Isolation)
	}

	if d.LegacyOverds != nil {
		legacy_overds = types.BoolValue(*d.LegacyOverds)
	}

	if d.LimitBcast != nil {
		limit_bcast = types.BoolValue(*d.LimitBcast)
	}

	if d.LimitProbeResponse != nil {
		limit_probe_response = types.BoolValue(*d.LimitProbeResponse)
	}

	if d.MaxIdletime != nil {
		max_idletime = types.Int64Value(int64(*d.MaxIdletime))
	}

	if d.MaxNumClients != nil {
		max_num_clients = types.Int64Value(int64(*d.MaxNumClients))
	}

	if d.MistNac != nil {
		mist_nac = mistNacdSkToTerraform(ctx, diags, d.MistNac)
	}

	if d.ModifiedTime != nil {
		modified_time = types.Float64Value(*d.ModifiedTime)
	}

	if d.MspId != nil {
		msp_id = types.StringValue(d.MspId.String())
	}

	if len(d.MxtunnelIds) > 0 {
		mxtunnel_ids = mist_transform.ListOfStringSdkToTerraform(ctx, d.MxtunnelIds)
	}

	if len(d.MxtunnelName) > 0 {
		mxtunnel_name = mist_transform.ListOfStringSdkToTerraform(ctx, d.MxtunnelName)
	}

	if d.NoStaticDns != nil {
		no_static_dns = types.BoolValue(*d.NoStaticDns)
	}

	if d.NoStaticIp != nil {
		no_static_ip = types.BoolValue(*d.NoStaticIp)
	}

	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}

	if d.Portal != nil {
		portal = portalSkToTerraform(ctx, diags, d.Portal)
	}

	if d.PortalAllowedHostnames != nil {
		portal_allowed_hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, d.PortalAllowedHostnames)
	}
	if d.PortalAllowedSubnets != nil {
		portal_allowed_subnets = mist_transform.ListOfStringSdkToTerraform(ctx, d.PortalAllowedSubnets)
	}

	if d.PortalApiSecret.IsValueSet() && d.PortalApiSecret.Value() != nil {
		portal_api_secret = types.StringValue(*d.PortalApiSecret.Value())
	}

	if d.PortalDeniedHostnames != nil {
		portal_denied_hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, d.PortalDeniedHostnames)
	}

	if d.PortalImage.IsValueSet() && d.PortalImage.Value() != nil {
		portal_image = types.StringValue("present")
	}

	if d.PortalSsoUrl.IsValueSet() && d.PortalSsoUrl.Value() != nil {
		portal_sso_url = types.StringValue(*d.PortalSsoUrl.Value())
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
		reconnect_clients_when_roaming_mxcluster = types.BoolValue(*d.ReconnectClientsWhenRoamingMxcluster)
	}

	if d.RoamMode != nil {
		roam_mode = types.StringValue(string(*d.RoamMode))
	}

	if d.Schedule != nil {
		schedule = scheduleSkToTerraform(ctx, diags, d.Schedule)
	}

	if d.SleExcluded != nil {
		sle_excluded = types.BoolValue(*d.SleExcluded)
	}

	ssid = types.StringValue(d.Ssid)

	template_id = types.StringValue(d.TemplateId.Value().String())

	if d.UseEapolV1 != nil {
		use_eapol_v1 = types.BoolValue(*d.UseEapolV1)
	}

	if d.VlanEnabled != nil {
		vlan_enabled = types.BoolValue(*d.VlanEnabled)
	}

	if d.VlanId != nil {
		vlan_id = mist_api.VlanAsString(*d.VlanId)
	}

	if d.VlanIds != nil {
		var list []attr.Value
		if vlan_ids_as_string, ok := d.VlanIds.AsString(); ok {
			for _, vlan := range strings.Split(*vlan_ids_as_string, ",") {
				list = append(list, types.StringValue(vlan))
			}
		} else if vlan_ids_as_list, ok := d.VlanIds.AsArrayOfVlanIdWithVariable2(); ok {
			for _, v := range *vlan_ids_as_list {
				list = append(list, mist_api.VlanAsString(v))
			}
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		vlan_ids = r
	}

	if d.VlanPooling != nil {
		vlan_pooling = types.BoolValue(*d.VlanPooling)
	}

	if d.WlanLimitDown.IsValueSet() && d.WlanLimitDown.Value() != nil {
		wlan_limit_down = types.Int64Value(int64(*d.WlanLimitDown.Value()))
	}

	if d.WlanLimitDownEnabled != nil {
		wlan_limit_down_enabled = types.BoolValue(*d.WlanLimitDownEnabled)
	}

	if d.WlanLimitUp.IsValueSet() && d.WlanLimitUp.Value() != nil {
		wlan_limit_up = types.Int64Value(int64(*d.WlanLimitUp.Value()))
	}

	if d.WlanLimitUpEnabled != nil {
		wlan_limit_up_enabled = types.BoolValue(*d.WlanLimitUpEnabled)
	}

	if d.WxtagIds.IsValueSet() && d.WxtagIds.Value() != nil {
		wxtag_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, *d.WxtagIds.Value())
	}

	if d.WxtunnelId.IsValueSet() && d.WxtunnelId.Value() != nil {
		wxtunnel_id = types.StringValue(*d.WxtunnelId.Value())
	}

	if d.WxtunnelRemoteId.IsValueSet() && d.WxtunnelRemoteId.Value() != nil {
		wxtunnel_remote_id = types.StringValue(*d.WxtunnelRemoteId.Value())
	}

	data_map_value := map[string]attr.Value{
		"acct_immediate_update":                  acct_immediate_update,
		"acct_interim_interval":                  acct_interim_interval,
		"acct_servers":                           acct_servers,
		"airwatch":                               airwatch,
		"allow_ipv6_ndp":                         allow_ipv6_ndp,
		"allow_mdns":                             allow_mdns,
		"allow_ssdp":                             allow_ssdp,
		"ap_ids":                                 ap_ids,
		"app_limit":                              app_limit,
		"app_qos":                                app_qos,
		"apply_to":                               apply_to,
		"arp_filter":                             arp_filter,
		"auth":                                   auth,
		"auth_server_selection":                  auth_server_selection,
		"auth_servers":                           auth_servers,
		"auth_servers_nas_id":                    auth_servers_nas_id,
		"auth_servers_nas_ip":                    auth_servers_nas_ip,
		"auth_servers_retries":                   auth_servers_retries,
		"auth_servers_timeout":                   auth_servers_timeout,
		"band_steer":                             band_steer,
		"band_steer_force_band5":                 band_steer_force_band5,
		"bands":                                  bands,
		"block_blacklist_clients":                block_blacklist_clients,
		"bonjour":                                bonjour,
		"cisco_cwa":                              cisco_cwa,
		"client_limit_down":                      client_limit_down,
		"client_limit_down_enabled":              client_limit_down_enabled,
		"client_limit_up":                        client_limit_up,
		"client_limit_up_enabled":                client_limit_up_enabled,
		"created_time":                           created_time,
		"coa_servers":                            coa_servers,
		"disable_11ax":                           disable_11ax,
		"disable_ht_vht_rates":                   disable_ht_vht_rates,
		"disable_uapsd":                          disable_uapsd,
		"disable_v1_roam_notify":                 disable_v1_roam_notify,
		"disable_v2_roam_notify":                 disable_v2_roam_notify,
		"disable_when_gateway_unreachable":       disable_when_gateway_unreachable,
		"disable_when_mxtunnel_down":             disable_when_mxtunnel_down,
		"disable_wmm":                            disable_wmm,
		"dns_server_rewrite":                     dns_server_rewrite,
		"dtim":                                   dtim,
		"dynamic_psk":                            dynamic_psk,
		"dynamic_vlan":                           dynamic_vlan,
		"enable_local_keycaching":                enable_local_keycaching,
		"enable_wireless_bridging":               enable_wireless_bridging,
		"enable_wireless_bridging_dhcp_tracking": enable_wireless_bridging_dhcp_tracking,
		"enabled":                                enabled,
		"fast_dot1x_timers":                      fast_dot1x_timers,
		"hide_ssid":                              hide_ssid,
		"hostname_ie":                            hostname_ie,
		"hotspot20":                              hotspot20,
		"id":                                     id,
		"inject_dhcp_option_82":                  inject_dhcp_option_82,
		"interface":                              interface_wlan,
		"isolation":                              isolation,
		"l2_isolation":                           l2_isolation,
		"legacy_overds":                          legacy_overds,
		"limit_bcast":                            limit_bcast,
		"limit_probe_response":                   limit_probe_response,
		"max_idletime":                           max_idletime,
		"max_num_clients":                        max_num_clients,
		"mist_nac":                               mist_nac,
		"modified_time":                          modified_time,
		"msp_id":                                 msp_id,
		"mxtunnel_ids":                           mxtunnel_ids,
		"mxtunnel_name":                          mxtunnel_name,
		"no_static_dns":                          no_static_dns,
		"no_static_ip":                           no_static_ip,
		"org_id":                                 org_id,
		"portal":                                 portal,
		"portal_allowed_hostnames":               portal_allowed_hostnames,
		"portal_allowed_subnets":                 portal_allowed_subnets,
		"portal_api_secret":                      portal_api_secret,
		"portal_denied_hostnames":                portal_denied_hostnames,
		"portal_image":                           portal_image,
		"portal_sso_url":                         portal_sso_url,
		"qos":                                    qos,
		"radsec":                                 radsec,
		"rateset":                                rateset,
		"reconnect_clients_when_roaming_mxcluster": reconnect_clients_when_roaming_mxcluster,
		"roam_mode":               roam_mode,
		"schedule":                schedule,
		"sle_excluded":            sle_excluded,
		"ssid":                    ssid,
		"template_id":             template_id,
		"use_eapol_v1":            use_eapol_v1,
		"vlan_enabled":            vlan_enabled,
		"vlan_id":                 vlan_id,
		"vlan_ids":                vlan_ids,
		"vlan_pooling":            vlan_pooling,
		"wlan_limit_down":         wlan_limit_down,
		"wlan_limit_down_enabled": wlan_limit_down_enabled,
		"wlan_limit_up":           wlan_limit_up,
		"wlan_limit_up_enabled":   wlan_limit_up_enabled,
		"wxtag_ids":               wxtag_ids,
		"wxtunnel_id":             wxtunnel_id,
		"wxtunnel_remote_id":      wxtunnel_remote_id,
	}

	data, e := NewOrgWlansValue(OrgWlansValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
