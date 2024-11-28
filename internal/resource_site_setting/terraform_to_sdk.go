package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *SiteSettingModel) (*models.SiteSetting, diag.Diagnostics) {
	data := models.SiteSetting{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if !plan.Analytic.IsNull() && !plan.Analytic.IsUnknown() {
		data.Analytic = analyticTerraformToSdk(ctx, &diags, plan.Analytic)
	} else {
		unset["-analytic"] = ""
	}

	if plan.ApUpdownThreshold.ValueInt64Pointer() != nil {
		data.ApUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.ApUpdownThreshold.ValueInt64())))
	} else {
		unset["-ap_updown_threshold"] = ""
	}

	if !plan.AutoUpgrade.IsNull() && !plan.AutoUpgrade.IsUnknown() {
		data.AutoUpgrade = siteSettingAutoUpgradeTerraformToSdk(ctx, &diags, plan.AutoUpgrade)
	} else {
		unset["-auto_upgrade"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		data.BleConfig = siteSettingBleConfigTerraformToSdk(ctx, &diags, plan.BleConfig)
	} else {
		unset["-ble_config"] = ""
	}

	if plan.ConfigAutoRevert.ValueBoolPointer() != nil {
		data.ConfigAutoRevert = plan.ConfigAutoRevert.ValueBoolPointer()
	} else {
		unset["-config_auto_revert"] = ""
	}

	if !plan.ConfigPushPolicy.IsNull() && !plan.ConfigPushPolicy.IsUnknown() {
		data.ConfigPushPolicy = pushPolicyConfigTerraformToSdk(ctx, &diags, plan.ConfigPushPolicy)
	} else {
		unset["-config_push_policy"] = ""
	}

	if !plan.CriticalUrlMonitoring.IsNull() && !plan.CriticalUrlMonitoring.IsUnknown() {
		data.CriticalUrlMonitoring = criticalUrlMonitoringTerraformToSdk(ctx, &diags, plan.CriticalUrlMonitoring)
	} else {
		unset["-critical_url_monitoring"] = ""
	}

	if plan.DeviceUpdownThreshold.ValueInt64Pointer() != nil {
		data.DeviceUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.DeviceUpdownThreshold.ValueInt64())))
	} else {
		unset["-device_updown_threshold"] = ""
	}

	if !plan.Engagement.IsNull() && !plan.Engagement.IsUnknown() {
		data.Engagement = engagementTerraformToSdk(ctx, &diags, plan.Engagement)
	} else {
		unset["-engagement"] = ""
	}

	if !plan.GatewayMgmt.IsNull() && !plan.GatewayMgmt.IsUnknown() {
		data.GatewayMgmt = gatewayMgmtTerraformToSdk(ctx, &diags, plan.GatewayMgmt)
	} else {
		unset["-gateway_mgmt"] = ""
	}

	if plan.GatewayUpdownThreshold.ValueInt64Pointer() != nil {
		data.GatewayUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.GatewayUpdownThreshold.ValueInt64())))
	} else {
		unset["-gateway_updown_threshold"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		led := ledTerraformToSdk(ctx, &diags, plan.Led)
		data.Led = led
	} else {
		unset["-led"] = ""
	}

	if !plan.Occupancy.IsNull() && !plan.Occupancy.IsUnknown() {
		data.Occupancy = occupancyTerraformToSdk(ctx, &diags, plan.Occupancy)
	} else {
		unset["-occupancy"] = ""
	}

	if plan.PersistConfigOnDevice.ValueBoolPointer() != nil {
		data.PersistConfigOnDevice = plan.PersistConfigOnDevice.ValueBoolPointer()
	} else {
		unset["-persist_config_on_device"] = ""
	}

	if !plan.Proxy.IsNull() && !plan.Proxy.IsUnknown() {
		data.Proxy = proxyTerraformToSdk(ctx, &diags, plan.Proxy)
	} else {
		unset["-proxy"] = ""

	}

	if plan.RemoveExistingConfigs.ValueBoolPointer() != nil {
		data.RemoveExistingConfigs = plan.RemoveExistingConfigs.ValueBoolPointer()
	} else {
		unset["remove_existing_configs"] = ""
	}

	if plan.ReportGatt.ValueBoolPointer() != nil {
		data.ReportGatt = plan.ReportGatt.ValueBoolPointer()
	} else {
		unset["-report_gatt"] = ""
	}

	if !plan.Rogue.IsNull() && !plan.Rogue.IsUnknown() {
		data.Rogue = rogueTerraformToSdk(ctx, &diags, plan.Rogue)
	} else {
		unset["-rogue"] = ""
	}

	if !plan.Rtsa.IsNull() && !plan.Rtsa.IsUnknown() {
		data.Rtsa = rtsaTerraformToSdk(ctx, &diags, plan.Rtsa)
	} else {
		unset["-rtsa"] = ""
	}

	if !plan.SimpleAlert.IsNull() && !plan.SimpleAlert.IsUnknown() {
		data.SimpleAlert = simpleAlertTerraformToSdk(ctx, &diags, plan.SimpleAlert)
	} else {
		unset["-simple_alert"] = ""
	}

	if !plan.Skyatp.IsNull() && !plan.Skyatp.IsUnknown() {
		data.Skyatp = skyAtpTerraformToSdk(ctx, &diags, plan.Skyatp)
	} else {
		unset["-skyatp"] = ""
	}

	if !plan.SrxApp.IsNull() && !plan.SrxApp.IsUnknown() {
		data.SrxApp = srxAppTerraformToSdk(ctx, &diags, plan.SrxApp)
	} else {
		unset["-srx_app"] = ""
	}

	if !plan.SshKeys.IsNull() && !plan.SshKeys.IsUnknown() {
		data.SshKeys = mist_transform.ListOfStringTerraformToSdk(ctx, plan.SshKeys)
	} else {
		unset["-ssh_keys"] = ""
	}

	if !plan.Ssr.IsNull() && !plan.Ssr.IsUnknown() {
		data.Ssr = ssrTerraformToSdk(ctx, &diags, plan.Ssr)
	} else {
		unset["-ssr"] = ""
	}

	if plan.SwitchUpdownThreshold.ValueInt64Pointer() != nil {
		data.SwitchUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.SwitchUpdownThreshold.ValueInt64())))
	} else {
		unset["-switch_updown_threshold"] = ""
	}

	if !plan.SyntheticTest.IsNull() && !plan.SyntheticTest.IsUnknown() {
		data.SyntheticTest = syntheticTestTerraformToSdk(ctx, &diags, plan.SyntheticTest)
	} else {
		unset["-synthetic_test"] = ""
	}

	if plan.TrackAnonymousDevices.ValueBoolPointer() != nil {
		data.TrackAnonymousDevices = plan.TrackAnonymousDevices.ValueBoolPointer()
	} else {
		unset["-track_anonymous_devices"] = ""
	}

	if !plan.UplinkPortConfig.IsNull() && !plan.UplinkPortConfig.IsUnknown() {
		data.UplinkPortConfig = uplinkPortConfigTerraformToSdk(ctx, &diags, plan.UplinkPortConfig)
	} else {
		unset["-uplink_port_config"] = ""
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(ctx, &diags, plan.Vars)
	} else {
		unset["-var"] = ""
	}

	if !plan.Vna.IsNull() && !plan.Vna.IsUnknown() {
		data.Vna = vnaTerraformToSdk(ctx, &diags, plan.Vna)
	} else {
		unset["-vna"] = ""
	}

	if !plan.VsInstance.IsNull() && !plan.VsInstance.IsUnknown() {
		data.VsInstance = vsInstanceTerraformToSdk(ctx, &diags, plan.VsInstance)
	} else {
		unset["-vs_instance"] = ""
	}

	if !plan.WanVna.IsNull() && !plan.WanVna.IsUnknown() {
		data.WanVna = wanVnaTerraformToSdk(ctx, &diags, plan.WanVna)
	} else {
		unset["-wan_vna"] = ""
	}

	if !plan.Wids.IsNull() && !plan.Wids.IsUnknown() {
		data.Wids = widsTerraformToSdk(ctx, &diags, plan.Wids)
	} else {
		unset["-wids"] = ""
	}

	if !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		data.Wifi = wifiTerraformToSdk(ctx, &diags, plan.Wifi)
	} else {
		unset["-wifi"] = ""
	}

	if !plan.WiredVna.IsNull() && !plan.WiredVna.IsUnknown() {
		data.WiredVna = wiredVnaTerraformToSdk(ctx, &diags, plan.WiredVna)
	} else {
		unset["-wired_vna"] = ""
	}

	if !plan.ZoneOccupancyAlert.IsNull() && !plan.ZoneOccupancyAlert.IsUnknown() {
		data.ZoneOccupancyAlert = zoneOccupancyTerraformToSdk(ctx, &diags, plan.ZoneOccupancyAlert)
	} else {
		unset["-zone_occupancy_alert"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
