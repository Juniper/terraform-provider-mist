package resource_site_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.Wlan) (SiteWlanModel, diag.Diagnostics) {
	var state SiteWlanModel
	var diags diag.Diagnostics

	if data.AcctImmediateUpdate != nil {
		state.AcctImmediateUpdate = types.BoolValue(*data.AcctImmediateUpdate)
	} else {
		state.AcctImmediateUpdate = types.BoolNull()
	}

	if data.AcctInterimInterval != nil {
		state.AcctInterimInterval = types.Int64Value(int64(*data.AcctInterimInterval))
	} else {
		state.AcctInterimInterval = types.Int64Null()
	}

	//if data.AcctServers != nil {
	state.AcctServers = radiusServersAcctSdkToTerraform(ctx, &diags, data.AcctServers)
	//} else {
	//state.AcctServers = types.ListNull(AcctServersValue{}.Type(ctx))
	//}

	if data.Airwatch != nil {
		state.Airwatch = airwatchSdkToTerraform(ctx, &diags, data.Airwatch)
	} else {
		state.Airwatch = NewAirwatchValueNull()
	}

	if data.AllowIpv6Ndp != nil {
		state.AllowIpv6Ndp = types.BoolValue(*data.AllowIpv6Ndp)
	} else {
		state.AllowIpv6Ndp = types.BoolNull()
	}

	if data.AllowMdns != nil {
		state.AllowMdns = types.BoolValue(*data.AllowMdns)
	} else {
		state.AllowMdns = types.BoolValue(false)
	}

	if data.AllowSsdp != nil {
		state.AllowSsdp = types.BoolValue(*data.AllowSsdp)
	} else {
		state.AllowSsdp = types.BoolValue(false)
	}

	if data.ApIds.IsValueSet() && data.ApIds.Value() != nil {
		state.ApIds = mist_transform.ListOfUuidSdkToTerraform(ctx, *data.ApIds.Value())
	} else {
		state.ApIds = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	}

	if data.AppLimit != nil {
		state.AppLimit = appLimitSdkToTerraform(ctx, &diags, data.AppLimit)
	} else {
		state.AppLimit = NewAppLimitValueNull()
	}

	if data.AppQos != nil {
		state.AppQos = appQosSdkToTerraform(ctx, &diags, *data.AppQos)
	} else {
		state.AppQos = NewAppQosValueNull()
	}

	if data.ApplyTo != nil {
		state.ApplyTo = types.StringValue(string(*data.ApplyTo))
	} else {
		state.ApplyTo = types.StringNull()
	}

	if data.ArpFilter != nil {
		state.ArpFilter = types.BoolValue(*data.ArpFilter)
	} else {
		state.ArpFilter = types.BoolNull()
	}

	if data.Auth != nil {
		state.Auth = authSdkToTerraform(ctx, &diags, data.Auth)
	} else {
		state.Auth = NewAuthValueNull()
	}

	if data.AuthServerSelection != nil {
		state.AuthServerSelection = types.StringValue(string(*data.AuthServerSelection))
	} else {
		state.AuthServerSelection = types.StringNull()
	}

	if data.AuthServers != nil {
		state.AuthServers = radiusServersAuthSdkToTerraform(ctx, &diags, data.AuthServers)
	} else {
		state.AuthServers = types.ListNull(AuthServersValue{}.Type(ctx))
	}

	if data.AuthServersNasId.IsValueSet() && data.AuthServersNasId.Value() != nil {
		state.AuthServersNasId = types.StringValue(*data.AuthServersNasId.Value())
	} else {
		state.AuthServersNasId = types.StringNull()
	}

	if data.AuthServersNasIp.IsValueSet() && data.AuthServersNasIp.Value() != nil {
		state.AuthServersNasIp = types.StringValue(*data.AuthServersNasIp.Value())
	} else {
		state.AuthServersNasIp = types.StringNull()
	}

	if data.AuthServersRetries != nil {
		state.AuthServersRetries = types.Int64Value(int64(*data.AuthServersRetries))
	} else {
		state.AuthServersRetries = types.Int64Null()
	}

	if data.AuthServersTimeout != nil {
		state.AuthServersTimeout = types.Int64Value(int64(*data.AuthServersTimeout))
	} else {
		state.AuthServersTimeout = types.Int64Null()
	}

	if data.BandSteer != nil {
		state.BandSteer = types.BoolValue(*data.BandSteer)
	} else {
		state.BandSteer = types.BoolNull()
	}

	if data.BandSteerForceBand5 != nil {
		state.BandSteerForceBand5 = types.BoolValue(*data.BandSteerForceBand5)
	} else {
		state.BandSteerForceBand5 = types.BoolNull()
	}

	if data.Bands != nil {
		state.Bands = bandsSdkToTerraform(ctx, &diags, data.Bands)
	} else {
		state.Bands = types.ListNull(types.StringType)
	}

	if data.BlockBlacklistClients != nil {
		state.BlockBlacklistClients = types.BoolValue(*data.BlockBlacklistClients)
	} else {
		state.BlockBlacklistClients = types.BoolNull()
	}

	if data.Bonjour != nil {
		state.Bonjour = bonjourSdkToTerraform(ctx, &diags, data.Bonjour)
	} else {
		state.Bonjour = NewBonjourValueNull()
	}

	if data.CiscoCwa != nil {
		state.CiscoCwa = ciscoCwaSdkToTerraform(ctx, &diags, data.CiscoCwa)
	} else {
		state.CiscoCwa = NewCiscoCwaValueNull()
	}

	if data.ClientLimitDown != nil {
		state.ClientLimitDown = types.Int64Value(int64(*data.ClientLimitDown))
	} else {
		state.ClientLimitDown = types.Int64Null()
	}

	if data.ClientLimitDownEnabled != nil {
		state.ClientLimitDownEnabled = types.BoolValue(*data.ClientLimitDownEnabled)
	} else {
		state.ClientLimitDownEnabled = types.BoolNull()
	}

	if data.ClientLimitUp != nil {
		state.ClientLimitUp = types.Int64Value(int64(*data.ClientLimitUp))
	} else {
		state.ClientLimitUp = types.Int64Null()
	}

	if data.ClientLimitUpEnabled != nil {
		state.ClientLimitUpEnabled = types.BoolValue(*data.ClientLimitUpEnabled)
	} else {
		state.ClientLimitUpEnabled = types.BoolNull()
	}

	if data.CoaServers.IsValueSet() && data.CoaServers.Value() != nil {
		state.CoaServers = coaServersSdkToTerraform(ctx, &diags, *data.CoaServers.Value())
	} else {
		state.CoaServers = types.ListNull(CoaServersValue{}.Type(ctx))
	}

	if data.Disable11ax != nil {
		state.Disable11ax = types.BoolValue(*data.Disable11ax)
	} else {
		state.Disable11ax = types.BoolNull()
	}

	if data.DisableHtVhtRates != nil {
		state.DisableHtVhtRates = types.BoolValue(*data.DisableHtVhtRates)
	} else {
		state.DisableHtVhtRates = types.BoolNull()
	}

	if data.DisableUapsd != nil {
		state.DisableUapsd = types.BoolValue(*data.DisableUapsd)
	} else {
		state.DisableUapsd = types.BoolNull()
	}

	if data.DisableV1RoamNotify != nil {
		state.DisableV1RoamNotify = types.BoolValue(*data.DisableV1RoamNotify)
	} else {
		state.DisableV1RoamNotify = types.BoolNull()
	}

	if data.DisableV2RoamNotify != nil {
		state.DisableV2RoamNotify = types.BoolValue(*data.DisableV2RoamNotify)
	} else {
		state.DisableV2RoamNotify = types.BoolNull()
	}

	if data.DisableWmm != nil {
		state.DisableWmm = types.BoolValue(*data.DisableWmm)
	} else {
		state.DisableWmm = types.BoolNull()
	}

	if data.DnsServerRewrite.IsValueSet() && data.DnsServerRewrite.Value() != nil {
		state.DnsServerRewrite = dnsServerRewriteSdkToTerraform(ctx, &diags, data.DnsServerRewrite.Value())
	} else {
		state.DnsServerRewrite = NewDnsServerRewriteValueNull()
	}

	if data.Dtim != nil {
		state.Dtim = types.Int64Value(int64(*data.Dtim))
	} else {
		state.Dtim = types.Int64Null()
	}

	if data.DynamicPsk.IsValueSet() && data.DynamicPsk.Value() != nil {
		state.DynamicPsk = dynamicPskSdkToTerraform(ctx, &diags, data.DynamicPsk.Value())
	} else {
		state.DynamicPsk = NewDynamicPskValueNull()
	}

	if data.DynamicVlan.IsValueSet() && data.DynamicVlan.Value() != nil {
		state.DynamicVlan = dynamicVlanSdkToTerraform(ctx, &diags, data.DynamicVlan.Value())
	} else {
		state.DynamicVlan = NewDynamicVlanValueNull()
	}

	if data.EnableLocalKeycaching != nil {
		state.EnableLocalKeycaching = types.BoolValue(*data.EnableLocalKeycaching)
	} else {
		state.EnableLocalKeycaching = types.BoolNull()
	}

	if data.EnableWirelessBridging != nil {
		state.EnableWirelessBridging = types.BoolValue(*data.EnableWirelessBridging)
	} else {
		state.EnableWirelessBridging = types.BoolNull()
	}

	if data.EnableWirelessBridgingDhcpTracking != nil {
		state.EnableWirelessBridgingDhcpTracking = types.BoolValue(*data.EnableWirelessBridgingDhcpTracking)
	} else {
		state.EnableWirelessBridgingDhcpTracking = types.BoolNull()
	}

	if data.Enabled != nil {
		state.Enabled = types.BoolValue(*data.Enabled)
	} else {
		state.Enabled = types.BoolNull()
	}

	if data.FastDot1xTimers != nil {
		state.FastDot1xTimers = types.BoolValue(*data.FastDot1xTimers)
	} else {
		state.FastDot1xTimers = types.BoolNull()
	}

	if data.HideSsid != nil {
		state.HideSsid = types.BoolValue(*data.HideSsid)
	} else {
		state.HideSsid = types.BoolNull()
	}

	if data.HostnameIe != nil {
		state.HostnameIe = types.BoolValue(*data.HostnameIe)
	} else {
		state.HostnameIe = types.BoolNull()
	}

	if data.Hotspot20 != nil {
		state.Hotspot20 = hotspot20SdkToTerraform(ctx, &diags, data.Hotspot20)
	} else {
		state.Hotspot20 = NewHotspot20ValueNull()
	}
	if data.Id != nil {
		state.Id = types.StringValue(data.Id.String())
	} else {
		state.Id = types.StringNull()
	}

	if data.InjectDhcpOption82 != nil {
		state.InjectDhcpOption82 = injectDhcpOption82dkToTerraform(ctx, &diags, data.InjectDhcpOption82)
	} else {
		state.InjectDhcpOption82 = NewInjectDhcpOption82ValueNull()
	}

	if data.Interface != nil {
		state.Interface = types.StringValue(string(*data.Interface))
	} else {
		state.Interface = types.StringNull()
	}
	if data.Isolation != nil {
		state.Isolation = types.BoolValue(*data.Isolation)
	} else {
		state.Isolation = types.BoolNull()
	}

	if data.Isolation != nil {
		state.L2Isolation = types.BoolValue(*data.Isolation)
	} else {
		state.Isolation = types.BoolNull()
	}

	if data.LegacyOverds != nil {
		state.LegacyOverds = types.BoolValue(*data.LegacyOverds)
	} else {
		state.LegacyOverds = types.BoolNull()
	}

	if data.LimitBcast != nil {
		state.LimitBcast = types.BoolValue(*data.LimitBcast)
	} else {
		state.LimitBcast = types.BoolNull()
	}

	if data.LimitProbeResponse != nil {
		state.LimitProbeResponse = types.BoolValue(*data.LimitProbeResponse)
	} else {
		state.LimitProbeResponse = types.BoolNull()
	}

	if data.MaxIdletime != nil {
		state.MaxIdletime = types.Int64Value(int64(*data.MaxIdletime))
	} else {
		state.MaxIdletime = types.Int64Null()
	}

	if data.MistNac != nil {
		state.MistNac = mistNacdSkToTerraform(ctx, &diags, data.MistNac)
	} else {
		state.MistNac = NewMistNacValueNull()
	}

	if data.MspId != nil {
		state.MspId = types.StringValue(data.MspId.String())
	} else {
		state.MspId = types.StringValue("")
	}

	if data.MxtunnelIds != nil {
		state.MxtunnelIds = mist_transform.ListOfStringSdkToTerraform(ctx, data.MxtunnelIds)
	} else {
		state.MxtunnelIds = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	}

	if data.MxtunnelName != nil {
		state.MxtunnelName = mist_transform.ListOfStringSdkToTerraform(ctx, data.MxtunnelName)
	} else {
		state.MxtunnelName = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	}

	if data.NoStaticDns != nil {
		state.NoStaticDns = types.BoolValue(*data.NoStaticDns)
	} else {
		state.NoStaticDns = types.BoolNull()
	}

	if data.NoStaticIp != nil {
		state.NoStaticIp = types.BoolValue(*data.NoStaticIp)
	} else {
		state.NoStaticIp = types.BoolNull()
	}

	if data.OrgId != nil {
		state.OrgId = types.StringValue(data.OrgId.String())
	} else {
		state.OrgId = types.StringNull()
	}

	if data.Portal != nil {
		state.Portal = portalSkToTerraform(ctx, &diags, data.Portal)
	} else {
		state.Portal = NewPortalValueNull()
	}

	if data.PortalAllowedHostnames != nil {
		state.PortalAllowedHostnames = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalAllowedHostnames)
	} else {
		state.PortalAllowedHostnames = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	}

	if data.PortalAllowedSubnets != nil {
		state.PortalAllowedSubnets = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalAllowedSubnets)
	} else {
		state.PortalAllowedSubnets = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	}

	if data.PortalApiSecret.IsValueSet() && data.PortalApiSecret.Value() != nil {
		state.PortalApiSecret = types.StringValue(*data.PortalApiSecret.Value())
	} else {
		state.PortalApiSecret = types.StringValue("")
	}

	if data.PortalDeniedHostnames != nil {
		state.PortalDeniedHostnames = mist_transform.ListOfStringSdkToTerraform(ctx, data.PortalDeniedHostnames)
	} else {
		state.PortalDeniedHostnames = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	}

	if data.PortalImage.IsValueSet() && data.PortalImage.Value() != nil {
		state.PortalImage = types.StringValue(*data.PortalImage.Value())
	} else {
		state.PortalImage = types.StringValue("")
	}

	if data.PortalSsoUrl.IsValueSet() && data.PortalSsoUrl.Value() != nil {
		state.PortalSsoUrl = types.StringValue(*data.PortalSsoUrl.Value())
	} else {
		state.PortalSsoUrl = types.StringValue("")
	}

	if data.PortalTemplateUrl.IsValueSet() && data.PortalTemplateUrl.Value() != nil {
		state.PortalTemplateUrl = types.StringValue(*data.PortalTemplateUrl.Value())
	} else {
		state.PortalTemplateUrl = types.StringValue("")
	}

	if data.Qos != nil {
		state.Qos = qosSkToTerraform(ctx, &diags, data.Qos)
	} else {
		state.Qos = NewQosValueNull()
	}

	if data.Radsec != nil {
		state.Radsec = radsecSkToTerraform(ctx, &diags, data.Radsec)
	} else {
		state.Radsec = NewRadsecValueNull()
	}

	if data.RoamMode != nil {
		state.RoamMode = types.StringValue(string(*data.RoamMode))
	} else {
		state.RoamMode = types.StringNull()
	}

	if data.Schedule != nil {
		state.Schedule = scheduleSkToTerraform(ctx, &diags, data.Schedule)
	} else {
		state.Schedule = NewScheduleValueNull()
	}

	if data.SiteId != nil {
		state.SiteId = types.StringValue(data.SiteId.String())
	} else {
		state.SiteId = types.StringValue("")
	}

	if data.SleExcluded != nil {
		state.SleExcluded = types.BoolValue(*data.SleExcluded)
	} else {
		state.SleExcluded = types.BoolNull()
	}

	state.Ssid = types.StringValue(data.Ssid)

	if data.Thumbnail.IsValueSet() && data.Thumbnail.Value() != nil {
		state.Thumbnail = types.StringValue(*data.Thumbnail.Value())
	} else {
		state.Thumbnail = types.StringValue("")
	}

	if data.UseEapolV1 != nil {
		state.UseEapolV1 = types.BoolValue(*data.UseEapolV1)
	} else {
		state.UseEapolV1 = types.BoolNull()
	}

	if data.VlanEnabled != nil {
		state.VlanEnabled = types.BoolValue(*data.VlanEnabled)
	} else {
		state.VlanEnabled = types.BoolNull()
	}

	if data.VlanId.IsValueSet() && data.VlanId.Value() != nil {
		state.VlanId = types.Int64Value(int64(*data.VlanId.Value()))
	} else {
		state.VlanId = types.Int64Null()
	}

	if data.VlanIds != nil {
		state.VlanIds = vlanIdsSkToTerraform(ctx, &diags, data.VlanIds)
	} else {
		state.VlanIds = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	}
	if data.VlanPooling != nil {
		state.VlanPooling = types.BoolValue(*data.VlanPooling)
	} else {
		state.VlanPooling = types.BoolNull()
	}

	if data.WlanLimitDown.IsValueSet() && data.WlanLimitDown.Value() != nil {
		state.WlanLimitDown = types.Int64Value(int64(*data.WlanLimitDown.Value()))
	} else {
		state.WlanLimitDown = types.Int64Null()
	}

	if data.WlanLimitDownEnabled != nil {
		state.WlanLimitDownEnabled = types.BoolValue(*data.WlanLimitDownEnabled)
	} else {
		state.WlanLimitDownEnabled = types.BoolNull()
	}

	if data.WlanLimitUp.IsValueSet() && data.WlanLimitUp.Value() != nil {
		state.WlanLimitUp = types.Int64Value(int64(*data.WlanLimitUp.Value()))
	} else {
		state.WlanLimitUp = types.Int64Null()
	}

	if data.WlanLimitUpEnabled != nil {
		state.WlanLimitUpEnabled = types.BoolValue(*data.WlanLimitUpEnabled)
	} else {
		state.WlanLimitUpEnabled = types.BoolNull()
	}

	if data.WxtagIds.IsValueSet() && data.WxtagIds.Value() != nil {
		state.WxtagIds = mist_transform.ListOfUuidSdkToTerraform(ctx, *data.WxtagIds.Value())
	} else {
		state.WxtagIds = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	}

	if data.WxtunnelId.IsValueSet() && data.WxtunnelId.Value() != nil {
		state.WxtunnelId = types.StringValue(*data.WxtunnelId.Value())
	} else {
		state.WxtunnelId = types.StringValue("")
	}
	if data.WxtunnelRemoteId.IsValueSet() && data.WxtunnelRemoteId.Value() != nil {
		state.WxtunnelRemoteId = types.StringValue(*data.WxtunnelRemoteId.Value())
	} else {
		state.WxtunnelRemoteId = types.StringValue("")
	}

	return state, diags
}
