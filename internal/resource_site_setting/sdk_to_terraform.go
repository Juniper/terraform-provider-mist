package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mist_list "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func SdkToTerraform(ctx context.Context, data *models.SiteSetting) (SiteSettingModel, diag.Diagnostics) {
	var state SiteSettingModel
	var diags diag.Diagnostics

	state.SiteId = types.StringValue(data.SiteId.String())
	state.OrgId = types.StringValue(data.OrgId.String())

	// state.Analytic = analyticSdkToTerraform(ctx, &diags, data.Analytic)

	if data.ApUpdownThreshold.Value() != nil {
		state.ApUpdownThreshold = types.Int64Value(int64(*data.ApUpdownThreshold.Value()))
	}

	if data.AutoUpgrade != nil {
		state.AutoUpgrade = autoUpgradeSdkToTerraform(ctx, &diags, *data.AutoUpgrade)
	}

	if data.BleConfig != nil {
		state.BleConfig = bleConfigsSdkToTerraform(ctx, &diags, data.BleConfig)
	}

	if data.BlacklistUrl != nil {
		state.BlacklistUrl = types.StringValue(*data.BlacklistUrl)
	}

	if data.ConfigAutoRevert != nil {
		state.ConfigAutoRevert = types.BoolValue(*data.ConfigAutoRevert)
	}

	if data.ConfigPushPolicy != nil {
		state.ConfigPushPolicy = configPushPolicySdkToTerraform(ctx, &diags, data.ConfigPushPolicy)
	}

	if data.CriticalUrlMonitoring != nil {
		state.CriticalUrlMonitoring = criticalUrlMonitoringSdkToTerraform(ctx, &diags, data.CriticalUrlMonitoring)
	}

	if data.DeviceUpdownThreshold != nil {
		state.DeviceUpdownThreshold = types.Int64Value(int64(*data.DeviceUpdownThreshold))
	}

	if data.DisabledSystemDefinedPortUsages != nil {
		state.DisabledSystemDefinedPortUsages = mist_list.ListOfStringSdkToTerraform(ctx, data.DisabledSystemDefinedPortUsages)
	} else {
		state.DisabledSystemDefinedPortUsages = types.ListNull(types.StringType)
	}

	if data.Engagement != nil {
		state.Engagement = engagementSdkToTerraform(ctx, &diags, data.Engagement)
	}

	if data.GatewayUpdownThreshold.Value() != nil {
		state.GatewayUpdownThreshold = types.Int64Value(int64(*data.GatewayUpdownThreshold.Value()))
	}

	if data.Led != nil {
		state.Led = ledSdkToTerraform(ctx, &diags, data.Led)
	}

	if data.Occupancy != nil {
		state.Occupancy = occupancySdkToTerraform(ctx, &diags, data.Occupancy)
	}

	if data.PersistConfigOnDevice != nil {
		state.PersistConfigOnDevice = types.BoolValue(*data.PersistConfigOnDevice)
	}

	if data.Proxy != nil {
		state.Proxy = proxySdkToTerraform(ctx, &diags, data.Proxy)
	}

	if data.ReportGatt != nil {
		state.ReportGatt = types.BoolValue(*data.ReportGatt)
	}

	if data.Rogue != nil {
		state.Rogue = rogueSdkToTerraform(ctx, &diags, data.Rogue)
	}

	if data.SimpleAlert != nil {
		state.SimpleAlert = simpleAlertSdkToTerraform(ctx, &diags, data.SimpleAlert)
	}

	if data.Skyatp != nil {
		state.Skyatp = skyAtpSdkToTerraform(ctx, &diags, data.Skyatp)
	}

	// state.SrxApp = srxAppSdkToTerraform(ctx, &diags, data.SrxApp)
	if data.SshKeys != nil {
		state.SshKeys = mist_list.ListOfStringSdkToTerraform(ctx, data.SshKeys)
	} else {
		state.SshKeys = types.ListNull(types.StringType)
	}

	if data.Ssr != nil {
		state.Ssr = ssrSdkToTerraform(ctx, &diags, data.Ssr)
	}

	if data.SwitchUpdownThreshold.Value() != nil {
		state.SwitchUpdownThreshold = types.Int64Value(int64(*data.SwitchUpdownThreshold.Value()))
	}

	if data.SyntheticTest != nil {
		state.SyntheticTest = synthteticTestSdkToTerraform(ctx, &diags, data.SyntheticTest)
	}

	if data.TrackAnonymousDevices != nil {
		state.TrackAnonymousDevices = types.BoolValue(*data.TrackAnonymousDevices)
	}

	if data.Vars != nil {
		state.Vars = varsSdkToTerraform(ctx, &diags, data.Vars)
	} else {
		state.Vars = types.MapNull(types.StringType)
	}

	// state.Vna = vnaSdkToTerraform(ctx, &diags, data.Vna)

	// state.WanVna = wanVnaSdkToTerraform(ctx, &diags, data.WanVna)

	if data.WatchedStationUrl != nil {
		state.WatchedStationUrl = types.StringValue(*data.WatchedStationUrl)
	}

	if data.WhitelistUrl != nil {
		state.WhitelistUrl = types.StringValue(*data.WhitelistUrl)
	}

	if data.Wids != nil {
		state.Wids = widsSdkToTerraform(ctx, &diags, data.Wids)
	}

	if data.Wifi != nil {
		state.Wifi = wifiSdkToTerraform(ctx, &diags, data.Wifi)
	}

	// state.WiredVna = wiredVnaSdkToTerraform(ctx, &diags, data.WiredVna)

	if data.ZoneOccupancyAlert != nil {
		state.ZoneOccupancyAlert = zoneOccupancySdkToTerraform(ctx, &diags, *data.ZoneOccupancyAlert)
	}

	return state, diags
}
