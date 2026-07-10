package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistlist "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteSettingModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Error: SiteSetting model is nil", "The SDK's SiteSetting model is nil.")
		return SiteSettingModel{}, diags
	}

	var allowMist types.Bool
	if data.AllowMist != nil {
		allowMist = types.BoolValue(*data.AllowMist)
	}

	var analytic = NewAnalyticValueNull()
	if data.Analytic != nil {
		analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)
	}

	var apSyntheticTest = NewApSyntheticTestValueNull()
	if data.ApSyntheticTest != nil {
		var additionalVlanIds = types.ListNull(types.StringType)
		if data.ApSyntheticTest.AdditionalVlanIds != nil {
			additionalVlanIds = mistlist.WlanBonjourAdditionalVlanIdsAsArrayOfString(&diags, *data.ApSyntheticTest.AdditionalVlanIds)
		}
		apSyntheticTest = ApSyntheticTestValue{
			AdditionalVlanIds: additionalVlanIds,
			state:             attr.ValueStateKnown,
		}
	}

	var apUpdownThreshold types.Int64
	if data.ApUpdownThreshold.Value() != nil {
		apUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}

	var autoUpgrade = NewAutoUpgradeValueNull()
	if data.AutoUpgrade != nil {
		autoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)
	}

	var blacklistUrl = types.StringValue("")
	if data.BlacklistUrl != nil {
		blacklistUrl = types.StringValue(*data.BlacklistUrl)
	}

	var bgpNeighborUpdownThreshold types.Int64
	if data.BgpNeighborUpdownThreshold.Value() != nil {
		bgpNeighborUpdownThreshold = types.Int64Value(int64(*data.BgpNeighborUpdownThreshold.Value()))
	}

	var bleConfig = NewBleConfigValueNull()
	if data.BleConfig != nil {
		bleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}

	var configAutoRevert types.Bool
	if data.ConfigAutoRevert != nil {
		configAutoRevert = types.BoolValue(*data.ConfigAutoRevert)
	}

	var configPushPolicy = NewConfigPushPolicyValueNull()
	if data.ConfigPushPolicy != nil {
		configPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)
	}

	var criticalUrlMonitoring = NewCriticalUrlMonitoringValueNull()
	if data.CriticalUrlMonitoring != nil {
		criticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)
	}

	var deviceUpdownThreshold types.Int64
	if data.DeviceUpdownThreshold.Value() != nil {
		deviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold.Value()))
	}

	var engagement = NewEngagementValueNull()
	if data.Engagement != nil {
		engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)
	}

	var enableUnii4 types.Bool
	if data.EnableUnii4 != nil {
		enableUnii4 = types.BoolValue(*data.EnableUnii4)
	}

	var gatewayMgmt = NewGatewayMgmtValueNull()
	if data.GatewayMgmt != nil {
		gatewayMgmt = gatewayMgmtSdkToTerraform(ctx, &diags, data.GatewayMgmt)
	}

	var gatewayUpdownThreshold types.Int64
	if data.GatewayUpdownThreshold.Value() != nil {
		gatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}

	var juniperSrx = NewJuniperSrxValueNull()
	if data.JuniperSrx != nil {
		juniperSrx = juniperSrxSdkToTerraform(ctx, &diags, data.JuniperSrx)
	}

	var iotproxy = NewIotproxyValueNull()
	if data.Iotproxy != nil {
		iotproxy = iotproxySdkToTerraform(ctx, &diags, data.Iotproxy)
	}

	var led = NewLedValueNull()
	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}

	var marvis = NewMarvisValueNull()
	if data.Marvis != nil {
		marvis = marvisSdkToTerraform(ctx, &diags, data.Marvis)
	}

	var occupancy = NewOccupancyValueNull()
	if data.Occupancy != nil {
		occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)
	}

	var persistConfigOnDevice types.Bool
	if data.PersistConfigOnDevice != nil {
		persistConfigOnDevice = types.BoolValue(*data.PersistConfigOnDevice)
	}

	var proxy = NewProxyValueNull()
	if data.Proxy != nil {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}

	var removeExistingConfigs types.Bool
	if data.RemoveExistingConfigs != nil {
		removeExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
	}

	var reportGatt types.Bool
	if data.ReportGatt != nil {
		reportGatt = types.BoolValue(*data.ReportGatt)
	}

	var rogue = NewRogueValueNull()
	if data.Rogue != nil {
		rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)
	}

	var rtsa = NewRtsaValueNull()
	if data.Rtsa != nil {
		rtsa = rtsaSdkToTerraform(ctx, &diags, data.Rtsa)
	}

	var simpleAlert = NewSimpleAlertValueNull()
	if data.SimpleAlert != nil {
		simpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)
	}

	var skyatp = NewSkyatpValueNull()
	if data.Skyatp != nil {
		skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)
	}

	var sleThresholds = NewSleThresholdsValueNull()
	if data.SleThresholds != nil {
		sleThresholds = sleThresholdsSdkToTerraform(ctx, &diags, data.SleThresholds)
	}

	var srxApp = NewSrxAppValueNull()
	if data.SrxApp != nil {
		srxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	}

	var sshKeys = types.ListValueMust(types.StringType, []attr.Value{})
	if data.SshKeys != nil {
		sshKeys = mistlist.ListOfStringSdkToTerraform(data.SshKeys)
	}

	var ssr = NewSsrValueNull()
	if data.Ssr != nil {
		ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}

	var switchUpdownThreshold types.Int64
	if data.SwitchUpdownThreshold.Value() != nil {
		switchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}

	var syntheticTest = NewSyntheticTestValueNull()
	if data.SyntheticTest != nil {
		syntheticTest = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}

	var mxedgeMgmt = NewMxedgeMgmtValueNull()
	if data.MxedgeMgmt != nil {
		mxedgeMgmt = mxedgeMgmtSdkToTerraform(ctx, &diags, data.MxedgeMgmt)
	}

	var mxtunnel = NewMxtunnelValueNull()
	if data.Mxtunnel != nil {
		mxtunnel = mxtunnelSdkToTerraform(ctx, &diags, data.Mxtunnel)
	}

	var tuntermMonitoring = types.ListNull(TuntermMonitoringValue{}.Type(ctx))
	if data.TuntermMonitoring != nil {
		var items []attr.Value
		for _, item := range data.TuntermMonitoring {
			var host = types.StringNull()
			var port = types.Int64Null()
			var protocol = types.StringNull()
			var srcVlanId = types.Int64Null()
			var timeout = types.Int64Null()
			if item.Host != nil {
				host = types.StringValue(*item.Host)
			}
			if item.Port != nil {
				port = types.Int64Value(int64(*item.Port))
			}
			if item.Protocol != nil {
				protocol = types.StringValue(string(*item.Protocol))
			}
			if item.SrcVlanId != nil {
				srcVlanId = types.Int64Value(int64(*item.SrcVlanId))
			}
			if item.Timeout != nil {
				timeout = types.Int64Value(int64(*item.Timeout))
			}
			itemObj, e := NewTuntermMonitoringValue(
				TuntermMonitoringValue{}.AttributeTypes(ctx),
				map[string]attr.Value{
					"host":        host,
					"port":        port,
					"protocol":    protocol,
					"src_vlan_id": srcVlanId,
					"timeout":     timeout,
				},
			)
			diags.Append(e...)
			items = append(items, itemObj)
		}
		r, e := types.ListValueFrom(ctx, TuntermMonitoringValue{}.Type(ctx), items)
		diags.Append(e...)
		tuntermMonitoring = r
	}

	var tuntermMonitoringDisabled = types.BoolNull()
	if data.TuntermMonitoringDisabled != nil {
		tuntermMonitoringDisabled = types.BoolValue(*data.TuntermMonitoringDisabled)
	}

	var tuntermMulticastConfig = NewTuntermMulticastConfigValueNull()
	if data.TuntermMulticastConfig != nil {
		tuntermMulticastConfig = tuntermMulticastConfigSdkToTerraform(ctx, &diags, data.TuntermMulticastConfig)
	}

	var trackAnonymousDevices types.Bool
	if data.TrackAnonymousDevices != nil {
		trackAnonymousDevices = types.BoolValue(*data.TrackAnonymousDevices)
	}

	var uplinkPortConfig = NewUplinkPortConfigValueNull()
	if data.UplinkPortConfig != nil {
		uplinkPortConfig = uplinkPortConfigValueSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}

	var vars = types.MapNull(types.StringType)
	if len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	var vna = NewVnaValueNull()
	if data.Vna != nil {
		vna = vnaSdkToTerraform(ctx, &diags, data.Vna)
	}

	var vpnPathUpdownThreshold types.Int64
	if data.VpnPathUpdownThreshold.Value() != nil {
		vpnPathUpdownThreshold = types.Int64Value(int64(*data.VpnPathUpdownThreshold.Value()))
	}

	var vpnPeerUpdownThreshold types.Int64
	if data.VpnPeerUpdownThreshold.Value() != nil {
		vpnPeerUpdownThreshold = types.Int64Value(int64(*data.VpnPeerUpdownThreshold.Value()))
	}

	var vsInstance = types.MapNull(VsInstanceValue{}.Type(ctx))
	if data.VsInstance != nil {
		vsInstance = vsInstanceSdkToTerraform(ctx, &diags, data.VsInstance)
	}

	var varsAnnotations = types.MapNull(VarsAnnotationsValue{}.Type(ctx))
	if data.VarsAnnotations != nil {
		varsAnnotations = varsAnnotationsSdkToTerraform(ctx, &diags, data.VarsAnnotations)
	}

	var wanVan = NewWanVnaValueNull()
	if data.WanVna != nil {
		wanVan = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)
	}

	var watchedStationUrl = types.StringValue("")
	if data.WatchedStationUrl != nil {
		watchedStationUrl = types.StringValue(*data.WatchedStationUrl)
	}

	var whitelistUrl = types.StringValue("")
	if data.WhitelistUrl != nil {
		whitelistUrl = types.StringValue(*data.WhitelistUrl)
	}

	var wids = NewWidsValueNull()
	if data.Wids != nil {
		wids = widsSdkToTerraform(ctx, &diags, data.Wids)
	}

	var wifi = NewWifiValueNull()
	if data.Wifi != nil {
		wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)
	}

	var wiredVna = NewWiredVnaValueNull()
	if data.WiredVna != nil {
		wiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)
	}

	var zoneOccupancyAlert = NewZoneOccupancyAlertValueNull()
	if data.ZoneOccupancyAlert != nil {
		zoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)
	}

	state := SiteSettingModel{
		SiteId:                     types.StringValue(data.SiteId.String()),
		AllowMist:                  allowMist,
		Analytic:                   analytic,
		ApSyntheticTest:            apSyntheticTest,
		ApUpdownThreshold:          apUpdownThreshold,
		AutoUpgrade:                autoUpgrade,
		BgpNeighborUpdownThreshold: bgpNeighborUpdownThreshold,
		BleConfig:                  bleConfig,
		BlacklistUrl:               blacklistUrl,
		ConfigAutoRevert:           configAutoRevert,
		ConfigPushPolicy:           configPushPolicy,
		CriticalUrlMonitoring:      criticalUrlMonitoring,
		DeviceUpdownThreshold:      deviceUpdownThreshold,
		EnableUnii4:                enableUnii4,
		Engagement:                 engagement,
		GatewayMgmt:                gatewayMgmt,
		GatewayUpdownThreshold:     gatewayUpdownThreshold,
		JuniperSrx:                 juniperSrx,
		Iotproxy:                   iotproxy,
		Led:                        led,
		Marvis:                     marvis,
		Occupancy:                  occupancy,
		PersistConfigOnDevice:      persistConfigOnDevice,
		Proxy:                      proxy,
		RemoveExistingConfigs:      removeExistingConfigs,
		ReportGatt:                 reportGatt,
		Rogue:                      rogue,
		Rtsa:                       rtsa,
		SimpleAlert:                simpleAlert,
		Skyatp:                     skyatp,
		SrxApp:                     srxApp,
		SleThresholds:              sleThresholds,
		SshKeys:                    sshKeys,
		Ssr:                        ssr,
		SwitchUpdownThreshold:      switchUpdownThreshold,
		SyntheticTest:              syntheticTest,
		TrackAnonymousDevices:      trackAnonymousDevices,
		UplinkPortConfig:           uplinkPortConfig,
		Vars:                       vars,
		Vna:                        vna,
		VpnPathUpdownThreshold:     vpnPathUpdownThreshold,
		VpnPeerUpdownThreshold:     vpnPeerUpdownThreshold,
		VsInstance:                 vsInstance,
		VarsAnnotations:            varsAnnotations,
		WanVna:                     wanVan,
		WatchedStationUrl:          watchedStationUrl,
		WhitelistUrl:               whitelistUrl,
		Wids:                       wids,
		Wifi:                       wifi,
		WiredVna:                   wiredVna,
		ZoneOccupancyAlert:         zoneOccupancyAlert,
		MxedgeMgmt:                 mxedgeMgmt,
		Mxtunnel:                   mxtunnel,
		TuntermMonitoring:          tuntermMonitoring,
		TuntermMonitoringDisabled:  tuntermMonitoringDisabled,
		TuntermMulticastConfig:     tuntermMulticastConfig,
	}

	return state, diags
}
