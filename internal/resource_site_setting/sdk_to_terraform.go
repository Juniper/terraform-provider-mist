package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_list "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteSettingModel, diag.Diagnostics) {
	var state SiteSettingModel
	var diags diag.Diagnostics

	var analytic AnalyticValue = NewAnalyticValueNull()
	var ap_updown_threshold types.Int64
	var auto_upgrade AutoUpgradeValue = NewAutoUpgradeValueNull()
	var blacklist_url types.String
	var ble_config BleConfigValue = NewBleConfigValueNull()
	var config_auto_revert types.Bool
	var config_push_policy ConfigPushPolicyValue = NewConfigPushPolicyValueNull()
	var critical_url_monitoring CriticalUrlMonitoringValue = NewCriticalUrlMonitoringValueNull()
	var device_updown_threshold types.Int64
	var disabled_system_defined_port_usages types.List = types.ListNull(types.StringType)
	var engagement EngagementValue = NewEngagementValueNull()
	var gateway_updown_threshold types.Int64
	var led LedValue = NewLedValueNull()
	var occupancy OccupancyValue = NewOccupancyValueNull()
	// var org_id types.String
	var persist_config_on_device types.Bool
	var proxy ProxyValue = NewProxyValueNull()
	var report_gatt types.Bool
	var rogue RogueValue = NewRogueValueNull()
	var simple_alert SimpleAlertValue = NewSimpleAlertValueNull()
	// var site_id types.String
	var skyatp SkyatpValue = NewSkyatpValueNull()
	var ssh_keys types.List = types.ListNull(types.StringType)
	var ssr SsrValue = NewSsrValueNull()
	var switch_updown_threshold types.Int64
	var synthetic_test SyntheticTestValue = NewSyntheticTestValueNull()
	var track_anonymous_devices types.Bool
	var vars types.Map = basetypes.NewMapNull(types.StringType)
	// var vs_instance VsInstanceValue = NewVsInstanceValueNull()
	var watched_station_url types.String
	var whitelist_url types.String
	var wids WidsValue = NewWidsValueNull()
	var wifi WifiValue = NewWifiValueNull()
	// var wired_vna WiredVnaValue = NewWiredVnaValueNull()
	var zone_occupancy_alert ZoneOccupancyAlertValue = NewZoneOccupancyAlertValueNull()

	state.SiteId = types.StringValue(data.SiteId.String())
	state.OrgId = types.StringValue(data.OrgId.String())

	if data.Analytic != nil {
		analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)
	}

	if data.ApUpdownThreshold.Value() != nil {
		ap_updown_threshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}

	if data.AutoUpgrade != nil {
		auto_upgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)
	}

	if data.BleConfig != nil {
		ble_config = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}

	if data.BlacklistUrl != nil {
		blacklist_url = types.StringValue(*data.BlacklistUrl)
	}

	if data.ConfigAutoRevert != nil {
		config_auto_revert = types.BoolValue(*data.ConfigAutoRevert)
	}

	if data.ConfigPushPolicy != nil {
		config_push_policy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)
	}

	if data.CriticalUrlMonitoring != nil {
		critical_url_monitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)
	}

	if data.DeviceUpdownThreshold != nil {
		device_updown_threshold = types.Int64Value(int64(*data.DeviceUpdownThreshold))
	}

	if data.DisabledSystemDefinedPortUsages != nil {
		disabled_system_defined_port_usages = mist_list.ListOfStringSdkToTerraform(ctx, data.DisabledSystemDefinedPortUsages)
	}

	if data.Engagement != nil {
		engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)
	}

	if data.GatewayUpdownThreshold.Value() != nil {
		gateway_updown_threshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}

	if data.Led != nil {
		led = ledSdkToTerraform(ctx, &diags, data.Led)
	}

	if data.Occupancy != nil {
		occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)
	}

	if data.PersistConfigOnDevice != nil {
		persist_config_on_device = types.BoolValue(*data.PersistConfigOnDevice)
	}

	if data.Proxy != nil {
		proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}

	if data.ReportGatt != nil {
		report_gatt = types.BoolValue(*data.ReportGatt)
	}

	if data.Rogue != nil {
		rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)
	}

	if data.SimpleAlert != nil {
		simple_alert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)
	}

	if data.Skyatp != nil {
		skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)
	}

	// state.SrxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	if data.SshKeys != nil {
		ssh_keys = mist_list.ListOfStringSdkToTerraform(ctx, data.SshKeys)
	}

	if data.Ssr != nil {
		ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}

	if data.SwitchUpdownThreshold.Value() != nil {
		switch_updown_threshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}

	if data.SyntheticTest != nil {
		synthetic_test = synthteticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}

	if data.TrackAnonymousDevices != nil {
		track_anonymous_devices = types.BoolValue(*data.TrackAnonymousDevices)
	}

	if data.Vars != nil {
		vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	}

	// state.Vna = vnaSdkToTerraform(ctx, &diags, data.Vna)

	// state.WanVna = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)

	if data.WatchedStationUrl != nil {
		watched_station_url = types.StringValue(*data.WatchedStationUrl)
	}

	if data.WhitelistUrl != nil {
		whitelist_url = types.StringValue(*data.WhitelistUrl)
	}

	if data.Wids != nil {
		wids = widsSdkToTerraform(ctx, &diags, data.Wids)
	}

	if data.Wifi != nil {
		wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)
	}

	// state.WiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)

	if data.ZoneOccupancyAlert != nil {
		zone_occupancy_alert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)
	}

	state.Analytic = analytic
	state.ApUpdownThreshold = ap_updown_threshold
	state.AutoUpgrade = auto_upgrade
	state.BleConfig = ble_config
	state.BlacklistUrl = blacklist_url
	state.ConfigAutoRevert = config_auto_revert
	state.ConfigPushPolicy = config_push_policy
	state.CriticalUrlMonitoring = critical_url_monitoring
	state.DeviceUpdownThreshold = device_updown_threshold
	state.DisabledSystemDefinedPortUsages = disabled_system_defined_port_usages
	state.Engagement = engagement
	state.GatewayUpdownThreshold = gateway_updown_threshold
	state.Led = led
	state.Occupancy = occupancy
	state.PersistConfigOnDevice = persist_config_on_device
	state.Proxy = proxy
	state.ReportGatt = report_gatt
	state.Rogue = rogue
	state.SimpleAlert = simple_alert
	state.Skyatp = skyatp
	// state.SrxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	state.SshKeys = ssh_keys
	state.Ssr = ssr
	state.SwitchUpdownThreshold = switch_updown_threshold
	state.SyntheticTest = synthetic_test
	state.TrackAnonymousDevices = track_anonymous_devices
	state.Vars = vars
	// state.Vna = vnaSdkToTerraform(ctx, &diags, data.Vna)
	// state.WanVna = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)
	state.WatchedStationUrl = watched_station_url
	state.WhitelistUrl = whitelist_url
	state.Wids = wids
	state.Wifi = wifi
	// state.WiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)
	state.ZoneOccupancyAlert = zone_occupancy_alert

	return state, diags
}
