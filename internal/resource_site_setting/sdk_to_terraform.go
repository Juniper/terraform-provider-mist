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
	var state SiteSettingModel
	var diags diag.Diagnostics

	var analytic = NewAnalyticValueNull()
	var apUpdownThreshold types.Int64
	var autoUpgrade = NewAutoUpgradeValueNull()
	var blacklistUrl = types.StringValue("")
	var bleConfig = NewBleConfigValueNull()
	var configAutoRevert types.Bool
	var configPushPolicy = NewConfigPushPolicyValueNull()
	var criticalUrlMonitoring = NewCriticalUrlMonitoringValueNull()
	var deviceUpdownThreshold types.Int64
	var engagement = NewEngagementValueNull()
	var enableUnii4 types.Bool
	var gatewayMgmt = NewGatewayMgmtValueNull()
	var gatewayUpdownThreshold types.Int64
	var juniperSrx = NewJuniperSrxValueNull()
	var led = NewLedValueNull()
	var occupancy = NewOccupancyValueNull()
	var persistConfigOnDevice types.Bool
	var proxy = NewProxyValueNull()
	var removeExistingConfigs types.Bool
	var reportGatt types.Bool
	var rogue = NewRogueValueNull()
	var rtsa = NewRtsaValueNull()
	var simpleAlert = NewSimpleAlertValueNull()
	var skyatp = NewSkyatpValueNull()
	var srxApp = NewSrxAppValueNull()
	var sshKeys = types.ListValueMust(types.StringType, []attr.Value{})
	var ssr = NewSsrValueNull()
	var switchUpdownThreshold types.Int64
	var syntheticTest = NewSyntheticTestValueNull()
	var trackAnonymousDevices types.Bool
	var uplinkPortConfig = NewUplinkPortConfigValueNull()
	var vars = types.MapNull(types.StringType)
	var vsInstance = types.MapNull(VsInstanceValue{}.Type(ctx))
	var vna = NewVnaValueNull()
	var watchedStationUrl = types.StringValue("")
	var whitelistUrl = types.StringValue("")
	var wids = NewWidsValueNull()
	var wifi = NewWifiValueNull()
	var wanVan = NewWanVnaValueNull()
	var wiredVna = NewWiredVnaValueNull()
	var zoneOccupancyAlert = NewZoneOccupancyAlertValueNull()

	state.SiteId = types.StringValue(data.SiteId.String())

	if data.Analytic != nil {
		analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)
	}

	if data.ApUpdownThreshold.Value() != nil {
		apUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}

	if data.AutoUpgrade != nil {
		autoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)
	}

	if data.BleConfig != nil {
		bleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}

	if data.BlacklistUrl != nil {
		blacklistUrl = types.StringValue(*data.BlacklistUrl)
	}

	if data.ConfigAutoRevert != nil {
		configAutoRevert = types.BoolValue(*data.ConfigAutoRevert)
	}

	if data.ConfigPushPolicy != nil {
		configPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)
	}

	if data.CriticalUrlMonitoring != nil {
		criticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)
	}

	if data.DeviceUpdownThreshold.Value() != nil {
		deviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold.Value()))
	}

	if data.Engagement != nil {
		engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)
	}

	if data.EnableUnii4 != nil {
		enableUnii4 = types.BoolValue(*data.EnableUnii4)
	}

	if data.GatewayMgmt != nil {
		gatewayMgmt = gatewayMgmtSdkToTerraform(ctx, &diags, data.GatewayMgmt)
	}

	if data.GatewayUpdownThreshold.Value() != nil {
		gatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}

	if data.JuniperSrx != nil {
		juniperSrx = juniperSrxSdkToTerraform(ctx, &diags, data.JuniperSrx)
	}

	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}

	if data.Occupancy != nil {
		occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)
	}

	if data.PersistConfigOnDevice != nil {
		persistConfigOnDevice = types.BoolValue(*data.PersistConfigOnDevice)
	}

	if data.Proxy != nil {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}

	if data.RemoveExistingConfigs != nil {
		removeExistingConfigs = types.BoolValue(*data.RemoveExistingConfigs)
	}

	if data.ReportGatt != nil {
		reportGatt = types.BoolValue(*data.ReportGatt)
	}

	if data.Rogue != nil {
		rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)
	}

	if data.Rtsa != nil {
		rtsa = rtsaSdkToTerraform(ctx, &diags, data.Rtsa)
	}

	if data.SimpleAlert != nil {
		simpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)
	}

	if data.Skyatp != nil {
		skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)
	}

	if data.SrxApp != nil {
		srxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	}

	if data.SshKeys != nil {
		sshKeys = mistlist.ListOfStringSdkToTerraform(data.SshKeys)
	}

	if data.Ssr != nil {
		ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}

	if data.SwitchUpdownThreshold.Value() != nil {
		switchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}

	if data.SyntheticTest != nil {
		syntheticTest = syntheticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}

	if data.TrackAnonymousDevices != nil {
		trackAnonymousDevices = types.BoolValue(*data.TrackAnonymousDevices)
	}

	if data.UplinkPortConfig != nil {
		uplinkPortConfig = uplinkPortConfigValueSdkToTerraform(ctx, &diags, data.UplinkPortConfig)
	}

	if data.Vars != nil && len(data.Vars) > 0 {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	if data.Vna != nil {
		vna = vnaSdkToTerraform(ctx, &diags, data.Vna)
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

	if data.Wids != nil {
		wids = widsSdkToTerraform(ctx, &diags, data.Wids)
	}

	if data.Wifi != nil {
		wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)
	}

	if data.WiredVna != nil {
		wiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)
	}

	if data.ZoneOccupancyAlert != nil {
		zoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)
	}

	state.Analytic = analytic
	state.ApUpdownThreshold = apUpdownThreshold
	state.AutoUpgrade = autoUpgrade
	state.BleConfig = bleConfig
	state.BlacklistUrl = blacklistUrl
	state.ConfigAutoRevert = configAutoRevert
	state.ConfigPushPolicy = configPushPolicy
	state.CriticalUrlMonitoring = criticalUrlMonitoring
	state.DeviceUpdownThreshold = deviceUpdownThreshold
	state.EnableUnii4 = enableUnii4
	state.Engagement = engagement
	state.GatewayMgmt = gatewayMgmt
	state.GatewayUpdownThreshold = gatewayUpdownThreshold
	state.JuniperSrx = juniperSrx
	state.Led = led
	state.Occupancy = occupancy
	state.PersistConfigOnDevice = persistConfigOnDevice
	state.Proxy = proxy
	state.RemoveExistingConfigs = removeExistingConfigs
	state.ReportGatt = reportGatt
	state.Rogue = rogue
	state.Rtsa = rtsa
	state.SimpleAlert = simpleAlert
	state.Skyatp = skyatp
	state.SrxApp = srxApp
	state.SshKeys = sshKeys
	state.Ssr = ssr
	state.SwitchUpdownThreshold = switchUpdownThreshold
	state.SyntheticTest = syntheticTest
	state.TrackAnonymousDevices = trackAnonymousDevices
	state.UplinkPortConfig = uplinkPortConfig
	state.Vars = vars
	state.Vna = vna
	state.VsInstance = vsInstance
	state.WanVna = wanVan
	state.WatchedStationUrl = watchedStationUrl
	state.WhitelistUrl = whitelistUrl
	state.Wids = wids
	state.Wifi = wifi
	state.WiredVna = wiredVna
	state.ZoneOccupancyAlert = zoneOccupancyAlert

	return state, diags
}
