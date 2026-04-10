package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteSettingModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Error: SiteSetting model is nil", "The SDK's SiteSetting model is nil.")
		return SiteSettingModel{}, diags
	}

	var analytic = NewAnalyticValueNull()
	var apUpdownThreshold types.Int64
	var autoUpgrade = NewAutoUpgradeValueNull()
	var autoUpgradeEsl = NewAutoUpgradeEslValueNull()
	var blacklistUrl = types.StringValue("")
	var bgpNeighborUpdownThreshold types.Int64
	var bleConfig = NewBleConfigValueNull()
	var configAutoRevert = types.BoolValue(false)
	var configPushPolicy = NewConfigPushPolicyValueNull()
	var criticalUrlMonitoring = NewCriticalUrlMonitoringValueNull()
	var deviceUpdownThreshold types.Int64
	var engagement = NewEngagementValueNull()
	var enableUnii4 = types.BoolValue(false)
	var gatewayMgmt = NewGatewayMgmtValueNull()
	var gatewayUpdownThreshold types.Int64
	var juniperSrx = NewJuniperSrxValueNull()
	var marvis = NewMarvisValueNull()
	var led = NewLedValueNull()
	var occupancy = NewOccupancyValueNull()
	var persistConfigOnDevice = types.BoolValue(false)
	var proxy = NewProxyValueNull()
	var removeExistingConfigs types.Bool
	var reportGatt types.Bool
	var rogue = NewRogueValueNull()
	var rtsa = NewRtsaValueNull()
	var simpleAlert = NewSimpleAlertValueNull()
	var skyatp = NewSkyatpValueNull()
	var sleThresholds = NewSleThresholdsValueNull()
	var srxApp = NewSrxAppValueNull()
	var sshKeys = types.ListNull(types.StringType)
	var ssr = NewSsrValueNull()
	var switchUpdownThreshold types.Int64
	var syntheticTest = NewSyntheticTestValueNull()
	var trackAnonymousDevices types.Bool
	var uplinkPortConfig = NewUplinkPortConfigValueNull()
	var usesDescriptionFromPortUsage types.Bool
	var vars = types.MapNull(types.StringType)
	var vna = NewVnaValueNull()
	var vpnPathUpdownThreshold types.Int64
	var vpnPeerUpdownThreshold types.Int64
	var vsInstance = types.MapNull(VsInstanceValue{}.Type(ctx))
	var wanVan = NewWanVnaValueNull()
	var watchedStationUrl = types.StringValue("")
	var whitelistUrl = types.StringValue("")
	var wids = NewWidsValueNull()
	var wifi = NewWifiValueNull()
	var wiredVna = NewWiredVnaValueNull()
	var zoneOccupancyAlert = NewZoneOccupancyAlertValueNull()

	if !mistutils.IsSdkDataEmpty(data.Analytic) {
		analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)
	}

	if data.ApUpdownThreshold.Value() != nil {
		apUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}

	if !mistutils.IsSdkDataEmpty(data.AutoUpgrade) {
		autoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)
	}

	if !mistutils.IsSdkDataEmpty(data.AutoUpgradeEsl) {
		autoUpgradeEsl = autoUpgradeEslSdkToTerraform(ctx, &diags, *data.AutoUpgradeEsl)
	}

	if data.BlacklistUrl != nil {
		blacklistUrl = types.StringValue(*data.BlacklistUrl)
	}

	if data.BgpNeighborUpdownThreshold.Value() != nil {
		bgpNeighborUpdownThreshold = types.Int64Value(int64(*data.BgpNeighborUpdownThreshold.Value()))
	}

	if !mistutils.IsSdkDataEmpty(data.BleConfig) {
		bleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}

	if data.ConfigAutoRevert != nil {
		configAutoRevert = types.BoolValue(*data.ConfigAutoRevert)
	}

	if !mistutils.IsSdkDataEmpty(data.ConfigPushPolicy) {
		configPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)
	}

	if !mistutils.IsSdkDataEmpty(data.CriticalUrlMonitoring) {
		criticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)
	}

	if data.DeviceUpdownThreshold.Value() != nil {
		deviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold.Value()))
	}

	if !mistutils.IsSdkDataEmpty(data.Engagement) {
		engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)
	}

	if data.EnableUnii4 != nil {
		enableUnii4 = types.BoolValue(*data.EnableUnii4)
	}

	if !mistutils.IsSdkDataEmpty(data.GatewayMgmt) {
		gatewayMgmt = gatewayMgmtSdkToTerraform(ctx, &diags, data.GatewayMgmt)
	}

	if data.GatewayUpdownThreshold.Value() != nil {
		gatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}

	if !mistutils.IsSdkDataEmpty(data.JuniperSrx) {
		juniperSrx = juniperSrxSdkToTerraform(ctx, &diags, data.JuniperSrx)
	}

	if !mistutils.IsSdkDataEmpty(data.Led) {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}

	if !mistutils.IsSdkDataEmpty(data.Marvis) {
		marvis = marvisSdkToTerraform(ctx, &diags, data.Marvis)
	}

	if !mistutils.IsSdkDataEmpty(data.Occupancy) {
		occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)
	}

	if data.PersistConfigOnDevice != nil {
		persistConfigOnDevice = types.BoolValue(*data.PersistConfigOnDevice)
	}

	if !mistutils.IsSdkDataEmpty(data.Proxy) {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}

	if data.RemoveExistingConfigs != nil {
		removeExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
	}

	if data.ReportGatt != nil {
		reportGatt = types.BoolValue(*data.ReportGatt)
	}

	if !mistutils.IsSdkDataEmpty(data.Rogue) {
		rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)
	}

	if !mistutils.IsSdkDataEmpty(data.Rtsa) {
		rtsa = rtsaSdkToTerraform(ctx, &diags, data.Rtsa)
	}

	if !mistutils.IsSdkDataEmpty(data.SimpleAlert) {
		simpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)
	}

	if !mistutils.IsSdkDataEmpty(data.Skyatp) {
		skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)
	}

	if !mistutils.IsSdkDataEmpty(data.SleThresholds) {
		sleThresholds = sleThresholdsSdkToTerraform(ctx, &diags, data.SleThresholds)
	}

	if !mistutils.IsSdkDataEmpty(data.SrxApp) {
		srxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	}

	if len(data.SshKeys) > 0 {
		sshKeys = mistutils.ListOfStringSdkToTerraform(data.SshKeys)
	}

	if !mistutils.IsSdkDataEmpty(data.Ssr) {
		ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}

	if data.SwitchUpdownThreshold.Value() != nil {
		switchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}

	if !mistutils.IsSdkDataEmpty(data.SyntheticTest) {
		syntheticTest = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}

	if data.TrackAnonymousDevices != nil {
		trackAnonymousDevices = types.BoolValue(*data.TrackAnonymousDevices)
	}

	if !mistutils.IsSdkDataEmpty(data.UplinkPortConfig) {
		uplinkPortConfig = uplinkPortConfigValueSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}

	if data.UsesDescriptionFromPortUsage != nil {
		usesDescriptionFromPortUsage = types.BoolValue(*data.UsesDescriptionFromPortUsage)
	}

	if len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	if !mistutils.IsSdkDataEmpty(data.Vna) {
		vna = vnaSdkToTerraform(ctx, &diags, data.Vna)
	}

	if data.VpnPathUpdownThreshold.Value() != nil {
		vpnPathUpdownThreshold = types.Int64Value(int64(*data.VpnPathUpdownThreshold.Value()))
	}

	if data.VpnPeerUpdownThreshold.Value() != nil {
		vpnPeerUpdownThreshold = types.Int64Value(int64(*data.VpnPeerUpdownThreshold.Value()))
	}

	if data.VsInstance != nil {
		vsInstance = vsInstanceSdkToTerraform(ctx, &diags, data.VsInstance)
	}

	if data.WanVna != nil {
		wanVan = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)
	}

	if data.WatchedStationUrl != nil {
		watchedStationUrl = types.StringValue(*data.WatchedStationUrl)
	}

	if data.WhitelistUrl != nil {
		whitelistUrl = types.StringValue(*data.WhitelistUrl)
	}

	if !mistutils.IsSdkDataEmpty(data.Wids) {
		wids = widsSdkToTerraform(ctx, &diags, data.Wids)
	}

	if !mistutils.IsSdkDataEmpty(data.Wifi) {
		wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)
	}

	if !mistutils.IsSdkDataEmpty(data.WiredVna) {
		wiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)
	}

	if !mistutils.IsSdkDataEmpty(data.ZoneOccupancyAlert) {
		zoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)
	}

	state := SiteSettingModel{
		SiteId:                       types.StringValue(data.SiteId.String()),
		Analytic:                     analytic,
		ApUpdownThreshold:            apUpdownThreshold,
		AutoUpgrade:                  autoUpgrade,
		AutoUpgradeEsl:               autoUpgradeEsl,
		BgpNeighborUpdownThreshold:   bgpNeighborUpdownThreshold,
		BleConfig:                    bleConfig,
		BlacklistUrl:                 blacklistUrl,
		ConfigAutoRevert:             configAutoRevert,
		ConfigPushPolicy:             configPushPolicy,
		CriticalUrlMonitoring:        criticalUrlMonitoring,
		DeviceUpdownThreshold:        deviceUpdownThreshold,
		EnableUnii4:                  enableUnii4,
		Engagement:                   engagement,
		GatewayMgmt:                  gatewayMgmt,
		GatewayUpdownThreshold:       gatewayUpdownThreshold,
		JuniperSrx:                   juniperSrx,
		Led:                          led,
		Marvis:                       marvis,
		Occupancy:                    occupancy,
		PersistConfigOnDevice:        persistConfigOnDevice,
		Proxy:                        proxy,
		RemoveExistingConfigs:        removeExistingConfigs,
		ReportGatt:                   reportGatt,
		Rogue:                        rogue,
		Rtsa:                         rtsa,
		SimpleAlert:                  simpleAlert,
		Skyatp:                       skyatp,
		SrxApp:                       srxApp,
		SleThresholds:                sleThresholds,
		SshKeys:                      sshKeys,
		Ssr:                          ssr,
		SwitchUpdownThreshold:        switchUpdownThreshold,
		SyntheticTest:                syntheticTest,
		TrackAnonymousDevices:        trackAnonymousDevices,
		UplinkPortConfig:             uplinkPortConfig,
		UsesDescriptionFromPortUsage: usesDescriptionFromPortUsage,
		Vars:                         vars,
		Vna:                          vna,
		VpnPathUpdownThreshold:       vpnPathUpdownThreshold,
		VpnPeerUpdownThreshold:       vpnPeerUpdownThreshold,
		VsInstance:                   vsInstance,
		WanVna:                       wanVan,
		WatchedStationUrl:            watchedStationUrl,
		WhitelistUrl:                 whitelistUrl,
		Wids:                         wids,
		Wifi:                         wifi,
		WiredVna:                     wiredVna,
		ZoneOccupancyAlert:           zoneOccupancyAlert,
	}

	return state, diags
}
