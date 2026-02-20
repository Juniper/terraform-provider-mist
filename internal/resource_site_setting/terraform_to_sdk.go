package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func TerraformToSdk(ctx context.Context, plan *SiteSettingModel) (*models.SiteSetting, diag.Diagnostics) {
	data := models.SiteSetting{}
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	if !plan.Analytic.IsNull() && !plan.Analytic.IsUnknown() {
		data.Analytic = analyticTerraformToSdk(plan.Analytic)
	} else {
		unset["-analytic"] = ""
	}

	if plan.ApUpdownThreshold.ValueInt64Pointer() != nil {
		data.ApUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.ApUpdownThreshold.ValueInt64())))
	} else {
		unset["-ap_updown_threshold"] = ""
	}

	if !plan.AutoUpgrade.IsNull() && !plan.AutoUpgrade.IsUnknown() {
		data.AutoUpgrade = siteSettingAutoUpgradeTerraformToSdk(plan.AutoUpgrade)
	} else {
		unset["-auto_upgrade"] = ""
	}

	if !plan.BgpNeighborUpdownThreshold.IsNull() && !plan.BgpNeighborUpdownThreshold.IsUnknown() {
		data.BgpNeighborUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.BgpNeighborUpdownThreshold.ValueInt64())))
	} else {
		unset["-bgp_neighbor_updown_threshold"] = ""
	}

	if !plan.BleConfig.IsNull() && !plan.BleConfig.IsUnknown() {
		data.BleConfig = siteSettingBleConfigTerraformToSdk(plan.BleConfig)
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
		data.CriticalUrlMonitoring = criticalUrlMonitoringTerraformToSdk(plan.CriticalUrlMonitoring)
	} else {
		unset["-critical_url_monitoring"] = ""
	}

	if plan.DeviceUpdownThreshold.ValueInt64Pointer() != nil {
		data.DeviceUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.DeviceUpdownThreshold.ValueInt64())))
	} else {
		unset["-device_updown_threshold"] = ""
	}

	if !plan.EnableUnii4.IsNull() && !plan.EnableUnii4.IsUnknown() {
		data.EnableUnii4 = plan.EnableUnii4.ValueBoolPointer()
	} else {
		unset["-enable_unii_4"] = ""
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

	if !plan.JuniperSrx.IsNull() && !plan.JuniperSrx.IsUnknown() {
		data.JuniperSrx = juniperSrxTerraformToSdk(ctx, &diags, plan.JuniperSrx)
	} else {
		unset["-juniper_srx"] = ""
	}

	if !plan.Led.IsNull() && !plan.Led.IsUnknown() {
		data.Led = ledTerraformToSdk(plan.Led)
	} else {
		unset["-led"] = ""
	}

	if !plan.Marvis.IsNull() && !plan.Marvis.IsUnknown() {
		data.Marvis = marvisTerraformToSdk(ctx, plan.Marvis)
	} else {
		unset["-marvis"] = ""
	}

	if !plan.Occupancy.IsNull() && !plan.Occupancy.IsUnknown() {
		data.Occupancy = occupancyTerraformToSdk(plan.Occupancy)
	} else {
		unset["-occupancy"] = ""
	}

	if plan.PersistConfigOnDevice.ValueBoolPointer() != nil {
		data.PersistConfigOnDevice = plan.PersistConfigOnDevice.ValueBoolPointer()
	} else {
		unset["-persist_config_on_device"] = ""
	}

	if !plan.Proxy.IsNull() && !plan.Proxy.IsUnknown() {
		data.Proxy = proxyTerraformToSdk(plan.Proxy)
	} else {
		unset["-proxy"] = ""

	}

	if plan.RemoveExistingConfigs.ValueBoolPointer() != nil {
		data.RemoveExistingConfigs = plan.RemoveExistingConfigs.ValueBoolPointer()
	} else {
		unset["-remove_existing_configs"] = ""
	}

	if plan.ReportGatt.ValueBoolPointer() != nil {
		data.ReportGatt = plan.ReportGatt.ValueBoolPointer()
	} else {
		unset["-report_gatt"] = ""
	}

	if !plan.Rogue.IsNull() && !plan.Rogue.IsUnknown() {
		data.Rogue = rogueTerraformToSdk(plan.Rogue)
	} else {
		unset["-rogue"] = ""
	}

	if !plan.Rtsa.IsNull() && !plan.Rtsa.IsUnknown() {
		data.Rtsa = rtsaTerraformToSdk(plan.Rtsa)
	} else {
		unset["-rtsa"] = ""
	}

	if !plan.SimpleAlert.IsNull() && !plan.SimpleAlert.IsUnknown() {
		data.SimpleAlert = simpleAlertTerraformToSdk(ctx, plan.SimpleAlert)
	} else {
		unset["-simple_alert"] = ""
	}

	if !plan.Skyatp.IsNull() && !plan.Skyatp.IsUnknown() {
		data.Skyatp = skyAtpTerraformToSdk(plan.Skyatp)
	} else {
		unset["-skyatp"] = ""
	}

	if !plan.SleThresholds.IsNull() && !plan.SleThresholds.IsUnknown() {
		data.SleThresholds = sleThresholdsTerraformToSdk(plan.SleThresholds)
	} else {
		unset["-sle_thresholds"] = ""
	}

	if !plan.SrxApp.IsNull() && !plan.SrxApp.IsUnknown() {
		data.SrxApp = srxAppTerraformToSdk(plan.SrxApp)
	} else {
		unset["-srx_app"] = ""
	}

	if !plan.SshKeys.IsNull() && !plan.SshKeys.IsUnknown() {
		data.SshKeys = mistutils.ListOfStringTerraformToSdk(plan.SshKeys)
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
		data.UplinkPortConfig = uplinkPortConfigTerraformToSdk(plan.UplinkPortConfig)
	} else {
		unset["-uplink_port_config"] = ""
	}

	if plan.UsesDescriptionFromPortUsage.ValueBoolPointer() != nil {
		data.UsesDescriptionFromPortUsage = plan.UsesDescriptionFromPortUsage.ValueBoolPointer()
	} else {
		unset["-uses_description_from_port_usage"] = ""
	}

	if !plan.Vars.IsNull() && !plan.Vars.IsUnknown() {
		data.Vars = varsTerraformToSdk(plan.Vars)
	} else {
		unset["-vars"] = ""
	}

	if !plan.Vna.IsNull() && !plan.Vna.IsUnknown() {
		data.Vna = vnaTerraformToSdk(plan.Vna)
	} else {
		unset["-vna"] = ""
	}

	if !plan.VpnPathUpdownThreshold.IsNull() && !plan.VpnPathUpdownThreshold.IsUnknown() {
		data.VpnPathUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.VpnPathUpdownThreshold.ValueInt64())))
	} else {
		unset["-vpn_path_updown_threshold"] = ""
	}

	if !plan.VpnPeerUpdownThreshold.IsNull() && !plan.VpnPeerUpdownThreshold.IsUnknown() {
		data.VpnPeerUpdownThreshold = models.NewOptional(models.ToPointer(int(plan.VpnPeerUpdownThreshold.ValueInt64())))
	} else {
		unset["-vpn_peer_updown_threshold"] = ""
	}

	if !plan.VsInstance.IsNull() && !plan.VsInstance.IsUnknown() {
		data.VsInstance = vsInstanceTerraformToSdk(plan.VsInstance)
	} else {
		unset["-vs_instance"] = ""
	}

	if !plan.WanVna.IsNull() && !plan.WanVna.IsUnknown() {
		data.WanVna = wanVnaTerraformToSdk(plan.WanVna)
	} else {
		unset["-wan_vna"] = ""
	}

	if !plan.Wids.IsNull() && !plan.Wids.IsUnknown() {
		data.Wids = widsTerraformToSdk(ctx, plan.Wids)
	} else {
		unset["-wids"] = ""
	}

	if !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		data.Wifi = wifiTerraformToSdk(plan.Wifi)
	} else {
		unset["-wifi"] = ""
	}

	if !plan.WiredVna.IsNull() && !plan.WiredVna.IsUnknown() {
		data.WiredVna = wiredVnaTerraformToSdk(plan.WiredVna)
	} else {
		unset["-wired_vna"] = ""
	}

	if !plan.ZoneOccupancyAlert.IsNull() && !plan.ZoneOccupancyAlert.IsUnknown() {
		data.ZoneOccupancyAlert = zoneOccupancyTerraformToSdk(plan.ZoneOccupancyAlert)
	} else {
		unset["-zone_occupancy_alert"] = ""
	}

	data.AdditionalProperties = unset

	return &data, diags
}
