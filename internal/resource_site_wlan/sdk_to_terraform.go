package resource_site_wlan

import (
	"context"
	"strings"

	mistapi "github.com/Juniper/terraform-provider-mist/internal/commons/api_response"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data *models.Wlan) (SiteWlanModel, diag.Diagnostics) {
	var state SiteWlanModel
	var diags diag.Diagnostics

	var acctImmediateUpdate = types.BoolValue(false)
	var acctInterimInterval types.Int64
	var acctServers = types.ListValueMust(AcctServersValue{}.Type(ctx), []attr.Value{})
	var airwatch = NewAirwatchValueNull()
	var allowIpv6Ndp = types.BoolValue(true)
	var allowMdns = types.BoolValue(false)
	var allowSsdp = types.BoolValue(false)
	var apIds = types.ListNull(types.StringType)
	var appLimit = NewAppLimitValueNull()
	var appQos = NewAppQosValueNull()
	var applyTo types.String
	var arpFilter types.Bool
	var auth = NewAuthValueNull()
	var authServerSelection types.String
	var authServers = types.ListValueMust(AuthServersValue{}.Type(ctx), []attr.Value{})
	var authServersNasId types.String
	var authServersNasIp types.String
	var authServersRetries types.Int64
	var authServersTimeout types.Int64
	var bandSteer types.Bool
	var bandSteerForceBand5 types.Bool
	var bands = types.ListNull(types.StringType)
	var blockBlacklistClients types.Bool
	var bonjour = NewBonjourValueNull()
	var ciscoCwa = NewCiscoCwaValueNull()
	var clientLimitDown types.Int64
	var clientLimitDownEnabled types.Bool
	var clientLimitUp types.Int64
	var clientLimitUpEnabled types.Bool
	var coaServers = types.ListValueMust(CoaServersValue{}.Type(ctx), []attr.Value{})
	var disable11ax types.Bool
	var disable11be types.Bool
	var disableHtVhtRates types.Bool
	var disableUapsd types.Bool
	var disableV1RoamNotify types.Bool
	var disableV2RoamNotify types.Bool
	var disableWhenGatewayUnreachable types.Bool
	var disableWhenMxtunnelDown types.Bool
	var disableWmm types.Bool
	var dnsServerRewrite = NewDnsServerRewriteValueNull()
	var dtim types.Int64
	var dynamicPsk = NewDynamicPskValueNull()
	var dynamicVlan = NewDynamicVlanValueNull()
	var enableLocalKeycaching types.Bool
	var enableWirelessBridging types.Bool
	var enableWirelessBridgingDhcpTracking types.Bool
	var enabled types.Bool
	var fastDot1xTimers types.Bool
	var hideSsid types.Bool
	var hostnameIe types.Bool
	var hotspot20 = NewHotspot20ValueNull()
	var id types.String
	var injectDhcpOption82 = NewInjectDhcpOption82ValueNull()
	var interfaceWlan types.String
	var isolation types.Bool
	var l2Isolation types.Bool
	var legacyOverds types.Bool
	var limitBcast types.Bool
	var limitProbeResponse types.Bool
	var maxIdletime types.Int64
	var maxNumClients types.Int64
	var mistNac = NewMistNacValueNull()
	var mspId = types.StringValue("")
	var mxtunnelIds = types.ListNull(types.StringType)
	var mxtunnelName = types.ListNull(types.StringType)
	var noStaticDns types.Bool
	var noStaticIp types.Bool
	var orgId types.String
	var portal = NewPortalValueNull()
	var portalAllowedHostnames = misttransform.ListOfStringSdkToTerraformEmpty()
	var portalAllowedSubnets = misttransform.ListOfStringSdkToTerraformEmpty()
	var portalApiSecret = types.StringValue("")
	var portalDeniedHostnames = misttransform.ListOfStringSdkToTerraformEmpty()
	var portalImage = types.StringValue("not_present")
	var portalSsoUrl = types.StringValue("")
	var qos = NewQosValueNull()
	var radsec = NewRadsecValueNull()
	var rateset = types.MapNull(RatesetValue{}.Type(ctx))
	var reconnectClientsWhenRoamingMxcluster types.Bool
	var roamMode types.String
	var schedule = NewScheduleValueNull()
	var siteId = types.StringValue("")
	var sleExcluded types.Bool
	var ssid types.String
	var useEapolV1 types.Bool
	var vlanEnabled types.Bool
	var vlanId types.String
	var vlanIds = misttransform.ListOfStringSdkToTerraformEmpty()
	var vlanPooling types.Bool
	var wlanLimitDown types.Int64
	var wlanLimitDownEnabled types.Bool
	var wlanLimitUp types.Int64
	var wlanLimitUpEnabled types.Bool
	var wxtagIds = misttransform.ListOfUuidSdkToTerraformEmpty()
	var wxtunnelId = types.StringValue("")
	var wxtunnelRemoteId = types.StringValue("")

	if data.AcctImmediateUpdate != nil {
		acctImmediateUpdate = types.BoolValue(*data.AcctImmediateUpdate)
	}

	if data.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*data.AcctInterimInterval))
	}

	if len(data.AcctServers) > 0 {
		acctServers = radiusServersAcctSdkToTerraform(ctx, &diags, data.AcctServers)
	} else {
		types.ListValueMust(AcctServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.Airwatch != nil {
		airwatch = airwatchSdkToTerraform(ctx, &diags, data.Airwatch)
	}

	if data.AllowIpv6Ndp != nil {
		allowIpv6Ndp = types.BoolValue(*data.AllowIpv6Ndp)
	}

	if data.AllowMdns != nil {
		allowMdns = types.BoolValue(*data.AllowMdns)
	}

	if data.AllowSsdp != nil {
		allowSsdp = types.BoolValue(*data.AllowSsdp)
	}

	if data.ApIds.Value() != nil && len(*data.ApIds.Value()) > 0 {
		apIds = misttransform.ListOfUuidSdkToTerraform(*data.ApIds.Value())
	}

	if data.AppLimit != nil {
		appLimit = appLimitSdkToTerraform(ctx, &diags, data.AppLimit)
	}

	if data.AppQos != nil {
		appQos = appQosSdkToTerraform(ctx, &diags, *data.AppQos)
	}

	if data.ApplyTo != nil {
		applyTo = types.StringValue(string(*data.ApplyTo))
	}

	if data.ArpFilter != nil {
		arpFilter = types.BoolValue(*data.ArpFilter)
	}

	if data.Auth != nil {
		auth = authSdkToTerraform(ctx, &diags, data.Auth)
	}

	if data.AuthServerSelection != nil {
		authServerSelection = types.StringValue(string(*data.AuthServerSelection))
	}

	if len(data.AuthServers) > 0 {
		authServers = radiusServersAuthSdkToTerraform(ctx, &diags, data.AuthServers)
	} else {
		authServers = types.ListValueMust(AuthServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.AuthServersNasId.IsValueSet() && data.AuthServersNasId.Value() != nil {
		authServersNasId = types.StringValue(*data.AuthServersNasId.Value())
	}

	if data.AuthServersNasIp.IsValueSet() && data.AuthServersNasIp.Value() != nil {
		authServersNasIp = types.StringValue(*data.AuthServersNasIp.Value())
	}

	if data.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*data.AuthServersRetries))
	}

	if data.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*data.AuthServersTimeout))
	}

	if data.BandSteer != nil {
		bandSteer = types.BoolValue(*data.BandSteer)
	}

	if data.BandSteerForceBand5 != nil {
		bandSteerForceBand5 = types.BoolValue(*data.BandSteerForceBand5)
	}

	if data.Bands != nil {
		bands = bandsSdkToTerraform(ctx, &diags, data.Bands)
	}

	if data.BlockBlacklistClients != nil {
		blockBlacklistClients = types.BoolValue(*data.BlockBlacklistClients)
	}

	if data.Bonjour != nil {
		bonjour = bonjourSdkToTerraform(ctx, &diags, data.Bonjour)
	}

	if data.CiscoCwa != nil {
		ciscoCwa = ciscoCwaSdkToTerraform(ctx, &diags, data.CiscoCwa)
	}

	if data.ClientLimitDown != nil {
		clientLimitDown = types.Int64Value(int64(*data.ClientLimitDown))
	}

	if data.ClientLimitDownEnabled != nil {
		clientLimitDownEnabled = types.BoolValue(*data.ClientLimitDownEnabled)
	}

	if data.ClientLimitUp != nil {
		clientLimitUp = types.Int64Value(int64(*data.ClientLimitUp))
	}

	if data.ClientLimitUpEnabled != nil {
		clientLimitUpEnabled = types.BoolValue(*data.ClientLimitUpEnabled)
	}

	if len(data.CoaServers) > 0 {
		coaServers = coaServersSdkToTerraform(ctx, &diags, data.CoaServers)
	} else {
		coaServers = types.ListValueMust(CoaServersValue{}.Type(ctx), make([]attr.Value, 0))
	}

	if data.Disable11ax != nil {
		disable11ax = types.BoolValue(*data.Disable11ax)
	}

	if data.Disable11be != nil {
		disable11be = types.BoolValue(*data.Disable11be)
	}

	if data.DisableHtVhtRates != nil {
		disableHtVhtRates = types.BoolValue(*data.DisableHtVhtRates)
	}

	if data.DisableUapsd != nil {
		disableUapsd = types.BoolValue(*data.DisableUapsd)
	}

	if data.DisableV1RoamNotify != nil {
		disableV1RoamNotify = types.BoolValue(*data.DisableV1RoamNotify)
	}

	if data.DisableV2RoamNotify != nil {
		disableV2RoamNotify = types.BoolValue(*data.DisableV2RoamNotify)
	}

	if data.DisableWhenGatewayUnreachable != nil {
		disableWhenGatewayUnreachable = types.BoolValue(*data.DisableWhenGatewayUnreachable)
	}

	if data.DisableWhenMxtunnelDown != nil {
		disableWhenMxtunnelDown = types.BoolValue(*data.DisableWhenMxtunnelDown)
	}

	if data.DisableWmm != nil {
		disableWmm = types.BoolValue(*data.DisableWmm)
	}

	if data.DnsServerRewrite.IsValueSet() && data.DnsServerRewrite.Value() != nil {
		dnsServerRewrite = dnsServerRewriteSdkToTerraform(ctx, &diags, data.DnsServerRewrite.Value())
	}

	if data.Dtim != nil {
		dtim = types.Int64Value(int64(*data.Dtim))
	}

	if data.DynamicPsk.IsValueSet() && data.DynamicPsk.Value() != nil {
		dynamicPsk = dynamicPskSdkToTerraform(ctx, &diags, data.DynamicPsk.Value())
	}

	if data.DynamicVlan.IsValueSet() && data.DynamicVlan.Value() != nil {
		dynamicVlan = dynamicVlanSdkToTerraform(ctx, &diags, data.DynamicVlan.Value())
	}

	if data.EnableLocalKeycaching != nil {
		enableLocalKeycaching = types.BoolValue(*data.EnableLocalKeycaching)
	}

	if data.EnableWirelessBridging != nil {
		enableWirelessBridging = types.BoolValue(*data.EnableWirelessBridging)
	}

	if data.EnableWirelessBridgingDhcpTracking != nil {
		enableWirelessBridgingDhcpTracking = types.BoolValue(*data.EnableWirelessBridgingDhcpTracking)
	}

	if data.Enabled != nil {
		enabled = types.BoolValue(*data.Enabled)
	}

	if data.FastDot1xTimers != nil {
		fastDot1xTimers = types.BoolValue(*data.FastDot1xTimers)
	}

	if data.HideSsid != nil {
		hideSsid = types.BoolValue(*data.HideSsid)
	}

	if data.HostnameIe != nil {
		hostnameIe = types.BoolValue(*data.HostnameIe)
	}

	if data.Hotspot20 != nil {
		hotspot20 = hotspot20SdkToTerraform(ctx, &diags, data.Hotspot20)
	}

	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}

	if data.InjectDhcpOption82 != nil {
		injectDhcpOption82 = injectDhcpOption82dkToTerraform(ctx, &diags, data.InjectDhcpOption82)
	}

	if data.Interface != nil {
		interfaceWlan = types.StringValue(string(*data.Interface))
	}

	if data.Isolation != nil {
		isolation = types.BoolValue(*data.Isolation)
	}

	if data.L2Isolation != nil {
		l2Isolation = types.BoolValue(*data.L2Isolation)
	}

	if data.LegacyOverds != nil {
		legacyOverds = types.BoolValue(*data.LegacyOverds)
	}

	if data.LimitBcast != nil {
		limitBcast = types.BoolValue(*data.LimitBcast)
	}

	if data.LimitProbeResponse != nil {
		limitProbeResponse = types.BoolValue(*data.LimitProbeResponse)
	}

	if data.MaxIdletime != nil {
		maxIdletime = types.Int64Value(int64(*data.MaxIdletime))
	}

	if data.MaxNumClients != nil {
		maxNumClients = types.Int64Value(int64(*data.MaxNumClients))
	}

	if data.MistNac != nil {
		mistNac = mistNacdSkToTerraform(ctx, &diags, data.MistNac)
	}

	if data.MspId != nil {
		mspId = types.StringValue(data.MspId.String())
	}

	if len(data.MxtunnelIds) > 0 {
		mxtunnelIds = misttransform.ListOfStringSdkToTerraform(data.MxtunnelIds)
	}

	if len(data.MxtunnelName) > 0 {
		mxtunnelName = misttransform.ListOfStringSdkToTerraform(data.MxtunnelName)
	}

	if data.NoStaticDns != nil {
		noStaticDns = types.BoolValue(*data.NoStaticDns)
	}

	if data.NoStaticIp != nil {
		noStaticIp = types.BoolValue(*data.NoStaticIp)
	}

	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}

	if data.Portal != nil {
		portal = portalSkToTerraform(ctx, &diags, data.Portal)
	}

	if data.PortalAllowedHostnames != nil {
		portalAllowedHostnames = misttransform.ListOfStringSdkToTerraform(data.PortalAllowedHostnames)
	}
	if data.PortalAllowedSubnets != nil {
		portalAllowedSubnets = misttransform.ListOfStringSdkToTerraform(data.PortalAllowedSubnets)
	}

	if data.PortalApiSecret.IsValueSet() && data.PortalApiSecret.Value() != nil {
		portalApiSecret = types.StringValue(*data.PortalApiSecret.Value())
	}

	if data.PortalDeniedHostnames != nil {
		portalDeniedHostnames = misttransform.ListOfStringSdkToTerraform(data.PortalDeniedHostnames)
	}

	if data.PortalImage.IsValueSet() && data.PortalImage.Value() != nil {
		portalImage = types.StringValue("present")
	}

	if data.PortalSsoUrl.IsValueSet() && data.PortalSsoUrl.Value() != nil {
		portalSsoUrl = types.StringValue(*data.PortalSsoUrl.Value())
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
		reconnectClientsWhenRoamingMxcluster = types.BoolValue(*data.ReconnectClientsWhenRoamingMxcluster)
	}

	if data.RoamMode != nil {
		roamMode = types.StringValue(string(*data.RoamMode))
	}

	if data.Schedule != nil {
		schedule = scheduleSkToTerraform(ctx, &diags, data.Schedule)
	}

	if data.SiteId != nil {
		siteId = types.StringValue(data.SiteId.String())
	}

	if data.SleExcluded != nil {
		sleExcluded = types.BoolValue(*data.SleExcluded)
	}

	ssid = types.StringValue(data.Ssid)

	if data.UseEapolV1 != nil {
		useEapolV1 = types.BoolValue(*data.UseEapolV1)
	}

	if data.VlanEnabled != nil {
		vlanEnabled = types.BoolValue(*data.VlanEnabled)
	}

	if data.VlanId != nil {
		vlanId = mistapi.VlanAsString(*data.VlanId)
	}

	if data.VlanIds != nil {
		var list []attr.Value
		if vlanIdsAsString, ok := data.VlanIds.AsString(); ok {
			for _, vlan := range strings.Split(*vlanIdsAsString, ",") {
				list = append(list, types.StringValue(vlan))
			}
		} else if vlanIdsAsList, ok := data.VlanIds.AsArrayOfVlanIdWithVariable2(); ok {
			for _, v := range *vlanIdsAsList {
				list = append(list, mistapi.VlanAsString(v))
			}
		}
		r, e := types.ListValue(basetypes.StringType{}, list)
		diags.Append(e...)
		vlanIds = r
	}

	if data.VlanPooling != nil {
		vlanPooling = types.BoolValue(*data.VlanPooling)
	}

	if data.WlanLimitDown.IsValueSet() && data.WlanLimitDown.Value() != nil {
		wlanLimitDown = types.Int64Value(int64(*data.WlanLimitDown.Value()))
	}

	if data.WlanLimitDownEnabled != nil {
		wlanLimitDownEnabled = types.BoolValue(*data.WlanLimitDownEnabled)
	}

	if data.WlanLimitUp.IsValueSet() && data.WlanLimitUp.Value() != nil {
		wlanLimitUp = types.Int64Value(int64(*data.WlanLimitUp.Value()))
	}

	if data.WlanLimitUpEnabled != nil {
		wlanLimitUpEnabled = types.BoolValue(*data.WlanLimitUpEnabled)
	}

	if data.WxtagIds.IsValueSet() && data.WxtagIds.Value() != nil {
		wxtagIds = misttransform.ListOfUuidSdkToTerraform(*data.WxtagIds.Value())
	}

	if data.WxtunnelId.IsValueSet() && data.WxtunnelId.Value() != nil {
		wxtunnelId = types.StringValue(*data.WxtunnelId.Value())
	}

	if data.WxtunnelRemoteId.IsValueSet() && data.WxtunnelRemoteId.Value() != nil {
		wxtunnelRemoteId = types.StringValue(*data.WxtunnelRemoteId.Value())
	}

	state.AcctImmediateUpdate = acctImmediateUpdate
	state.AcctInterimInterval = acctInterimInterval
	state.AcctServers = acctServers
	state.Airwatch = airwatch
	state.AllowIpv6Ndp = allowIpv6Ndp
	state.AllowMdns = allowMdns
	state.AllowSsdp = allowSsdp
	state.ApIds = apIds
	state.AppLimit = appLimit
	state.AppQos = appQos
	state.ApplyTo = applyTo
	state.ArpFilter = arpFilter
	state.Auth = auth
	state.AuthServerSelection = authServerSelection
	state.AuthServers = authServers
	state.AuthServersNasId = authServersNasId
	state.AuthServersNasIp = authServersNasIp
	state.AuthServersRetries = authServersRetries
	state.AuthServersTimeout = authServersTimeout
	state.BandSteer = bandSteer
	state.BandSteerForceBand5 = bandSteerForceBand5
	state.Bands = bands
	state.BlockBlacklistClients = blockBlacklistClients
	state.Bonjour = bonjour
	state.CiscoCwa = ciscoCwa
	state.ClientLimitDown = clientLimitDown
	state.ClientLimitDownEnabled = clientLimitDownEnabled
	state.ClientLimitUp = clientLimitUp
	state.ClientLimitUpEnabled = clientLimitUpEnabled
	state.CoaServers = coaServers
	state.Disable11ax = disable11ax
	state.Disable11be = disable11be
	state.DisableHtVhtRates = disableHtVhtRates
	state.DisableUapsd = disableUapsd
	state.DisableV1RoamNotify = disableV1RoamNotify
	state.DisableV2RoamNotify = disableV2RoamNotify
	state.DisableWhenGatewayUnreachable = disableWhenGatewayUnreachable
	state.DisableWhenMxtunnelDown = disableWhenMxtunnelDown
	state.DisableWmm = disableWmm
	state.DnsServerRewrite = dnsServerRewrite
	state.Dtim = dtim
	state.DynamicPsk = dynamicPsk
	state.DynamicVlan = dynamicVlan
	state.EnableLocalKeycaching = enableLocalKeycaching
	state.EnableWirelessBridging = enableWirelessBridging
	state.EnableWirelessBridgingDhcpTracking = enableWirelessBridgingDhcpTracking
	state.Enabled = enabled
	state.FastDot1xTimers = fastDot1xTimers
	state.HideSsid = hideSsid
	state.HostnameIe = hostnameIe
	state.Hotspot20 = hotspot20
	state.Id = id
	state.InjectDhcpOption82 = injectDhcpOption82
	state.Interface = interfaceWlan
	state.Isolation = isolation
	state.L2Isolation = l2Isolation
	state.LegacyOverds = legacyOverds
	state.LimitBcast = limitBcast
	state.LimitProbeResponse = limitProbeResponse
	state.MaxIdletime = maxIdletime
	state.MaxNumClients = maxNumClients
	state.MistNac = mistNac
	state.MspId = mspId
	state.MxtunnelIds = mxtunnelIds
	state.MxtunnelName = mxtunnelName
	state.NoStaticDns = noStaticDns
	state.NoStaticIp = noStaticIp
	state.OrgId = orgId
	state.Portal = portal
	state.PortalAllowedHostnames = portalAllowedHostnames
	state.PortalAllowedSubnets = portalAllowedSubnets
	state.PortalApiSecret = portalApiSecret
	state.PortalDeniedHostnames = portalDeniedHostnames
	state.PortalImage = portalImage
	state.PortalSsoUrl = portalSsoUrl
	state.Qos = qos
	state.Radsec = radsec
	state.Rateset = rateset
	state.ReconnectClientsWhenRoamingMxcluster = reconnectClientsWhenRoamingMxcluster
	state.RoamMode = roamMode
	state.Schedule = schedule
	state.SiteId = siteId
	state.SleExcluded = sleExcluded
	state.Ssid = ssid
	state.UseEapolV1 = useEapolV1
	state.VlanEnabled = vlanEnabled
	state.VlanId = vlanId
	state.VlanIds = vlanIds
	state.VlanPooling = vlanPooling
	state.WlanLimitDown = wlanLimitDown
	state.WlanLimitDownEnabled = wlanLimitDownEnabled
	state.WlanLimitUp = wlanLimitUp
	state.WlanLimitUpEnabled = wlanLimitUpEnabled
	state.WxtagIds = wxtagIds
	state.WxtunnelId = wxtunnelId
	state.WxtunnelRemoteId = wxtunnelRemoteId

	return state, diags
}
