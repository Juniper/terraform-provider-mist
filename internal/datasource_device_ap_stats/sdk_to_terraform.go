package datasource_device_ap_stats

import (
	"context"
	"math/big"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, l *[]models.StatsAp, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := deviceApStatSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func deviceApStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsAp) DeviceApStatsValue {

	var autoPlacement = types.ObjectNull(AutoPlacementValue{}.AttributeTypes(ctx))
	var autoUpgradeStat = types.ObjectNull(AutoUpgradeStatValue{}.AttributeTypes(ctx))
	var bleStat = types.ObjectNull(BleStatValue{}.AttributeTypes(ctx))
	var certExpiry basetypes.NumberValue
	var configReverted basetypes.BoolValue
	var cpuSystem basetypes.Int64Value
	var cpuUtil basetypes.Int64Value
	var createdTime basetypes.Int64Value
	var deviceprofileId basetypes.StringValue
	var envStat = types.ObjectNull(EnvStatValue{}.AttributeTypes(ctx))
	var eslStat = types.ObjectNull(EslStatValue{}.AttributeTypes(ctx))
	var extIp basetypes.StringValue
	var fwupdate = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var hwRev basetypes.StringValue
	var id basetypes.StringValue
	var inactiveWiredVlans = types.ListNull(types.Int64Type)
	var iotStat = types.MapNull(IotStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ipConfig = types.ObjectNull(IpConfigValue{}.AttributeTypes(ctx))
	var ipStat = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var l2tpStat = types.MapNull(L2tpStatValue{}.Type(ctx))
	var lastSeen basetypes.NumberValue
	var lastTrouble = types.ObjectNull(LastTroubleValue{}.AttributeTypes(ctx))
	var led = types.ObjectNull(LedValue{}.AttributeTypes(ctx))
	var lldpStat = types.ObjectNull(LldpStatValue{}.AttributeTypes(ctx))
	var locating basetypes.BoolValue
	var locked basetypes.BoolValue
	var mac basetypes.StringValue
	var mapId basetypes.StringValue
	var memUsedKb basetypes.Int64Value
	var meshDownlinks = types.MapNull(MeshDownlinksValue{}.Type(ctx))
	var meshUplink = types.ObjectNull(MeshUplinkValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modifiedTime basetypes.Int64Value
	var mount basetypes.StringValue
	var name basetypes.StringValue
	var notes basetypes.StringValue
	var numClients basetypes.Int64Value
	var orgId basetypes.StringValue
	var portStat = types.MapNull(PortStatValue{}.Type(ctx))
	var powerBudget basetypes.Int64Value
	var powerConstrained basetypes.BoolValue
	var powerOpmode basetypes.StringValue
	var powerSrc basetypes.StringValue
	var radioStat = types.ObjectNull(RadioStatValue{}.AttributeTypes(ctx))
	var rxBps basetypes.NumberValue
	var rxBytes basetypes.Int64Value
	var rxPkts basetypes.Int64Value
	var serial basetypes.StringValue
	var siteId basetypes.StringValue
	var status basetypes.StringValue
	var switchRedundancy = types.ObjectNull(SwitchRedundancyValue{}.AttributeTypes(ctx))
	var txBps basetypes.NumberValue
	var txBytes basetypes.NumberValue
	var txPkts basetypes.NumberValue
	var uptime basetypes.NumberValue
	var usbStat = types.ObjectNull(UsbStatValue{}.AttributeTypes(ctx))
	var version basetypes.StringValue
	var x basetypes.Float64Value
	var y basetypes.Float64Value

	if d.AutoPlacement != nil {
		autoPlacement = autoPlacementSdkToTerraform(ctx, diags, d.AutoPlacement)
	}
	if d.AutoUpgradeStat != nil {
		autoUpgradeStat = autoUpgradeStatsSdkToTerraform(ctx, diags, d.AutoUpgradeStat)
	}
	if d.BleStat != nil {
		bleStat = bleStatsSdkToTerraform(ctx, diags, d.BleStat)
	}
	if d.CertExpiry.Value() != nil {
		certExpiry = types.NumberValue(big.NewFloat(*d.CertExpiry.Value()))
	}
	if d.ConfigReverted.Value() != nil {
		configReverted = types.BoolValue(*d.ConfigReverted.Value())
	}
	if d.CpuSystem.Value() != nil {
		cpuSystem = types.Int64Value(*d.CpuSystem.Value())
	}
	if d.CpuUtil.Value() != nil {
		cpuUtil = types.Int64Value(int64(*d.CpuUtil.Value()))
	}
	if d.CreatedTime != nil {
		createdTime = types.Int64Value(int64(*d.CreatedTime))
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofileId = types.StringValue(d.DeviceprofileId.Value().String())
	}
	if d.EnvStat != nil {
		envStat = envStatsSdkToTerraform(ctx, diags, d.EnvStat)
	}
	if d.EslStat.Value() != nil {
		eslStat = eslStatsSdkToTerraform(ctx, diags, d.EslStat.Value())
	}
	if d.ExtIp.Value() != nil {
		extIp = types.StringValue(*d.ExtIp.Value())
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HwRev.Value() != nil {
		hwRev = types.StringValue(*d.HwRev.Value())
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.InactiveWiredVlans != nil {
		inactiveWiredVlans = misttransform.ListOfIntSdkToTerraform(d.InactiveWiredVlans)
	}
	if d.IotStat != nil && len(d.IotStat) > 0 {
		iotStat = iotStatsSdkToTerraform(ctx, diags, d.IotStat)
	}
	if d.Ip.Value() != nil {
		ip = types.StringValue(*d.Ip.Value())
	}
	if d.IpConfig != nil {
		ipConfig = ipConfigSdkToTerraform(ctx, diags, d.IpConfig)
	}
	if d.IpStat != nil {
		ipStat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.L2tpStat != nil && len(d.L2tpStat) > 0 {
		l2tpStat = l2tpStatsSdkToTerraform(ctx, diags, d.L2tpStat)
	}
	if d.LastSeen.Value() != nil {
		lastSeen = types.NumberValue(big.NewFloat(*d.LastSeen.Value()))
	}
	if d.LastTrouble != nil {
		lastTrouble = lastTroubleSdkToTerraform(ctx, diags, d.LastTrouble)
	}
	if d.Led != nil {
		led = ledSdkToTerraform(ctx, diags, d.Led)
	}
	if d.LldpStat != nil {
		lldpStat = lldpSdkToTerraform(ctx, diags, d.LldpStat)
	}
	if d.Locating.Value() != nil {
		locating = types.BoolValue(*d.Locating.Value())
	}
	if d.Locked.Value() != nil {
		locked = types.BoolValue(*d.Locked.Value())
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.MapId.Value() != nil {
		mapId = types.StringValue(d.MapId.Value().String())
	}
	if d.MemUsedKb.Value() != nil {
		memUsedKb = types.Int64Value(*d.MemUsedKb.Value())
	}
	if d.MeshDownlinks != nil && len(d.MeshDownlinks) > 0 {
		meshDownlinks = meshDownlinksSdkToTerraform(ctx, diags, d.MeshDownlinks)
	}
	if d.MeshUplink != nil {
		meshUplink = meshUplinkSdkToTerraform(ctx, diags, d.MeshUplink)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Int64Value(int64(*d.ModifiedTime))
	}
	if d.Mount.Value() != nil {
		mount = types.StringValue(*d.Mount.Value())
	}
	if d.Name.Value() != nil {
		name = types.StringValue(*d.Name.Value())
	}
	if d.Notes.Value() != nil {
		notes = types.StringValue(*d.Notes.Value())
	}
	if d.NumClients.Value() != nil {
		numClients = types.Int64Value(int64(*d.NumClients.Value()))
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.PortStat.Value() != nil && len(*d.PortStat.Value()) > 0 {
		portStat = portStatdkToTerraform(ctx, diags, *d.PortStat.Value())
	}
	if d.PowerBudget.Value() != nil {
		powerBudget = types.Int64Value(int64(*d.PowerBudget.Value()))
	}
	if d.PowerConstrained.Value() != nil {
		powerConstrained = types.BoolValue(*d.PowerConstrained.Value())
	}
	if d.PowerOpmode.Value() != nil {
		powerOpmode = types.StringValue(*d.PowerOpmode.Value())
	}
	if d.PowerSrc.Value() != nil {
		powerSrc = types.StringValue(*d.PowerSrc.Value())
	}
	if d.RadioStat != nil {
		radioStat = radioStatSdkToTerraform(ctx, diags, d.RadioStat)
	}
	if d.RxBps.Value() != nil {
		rxBps = types.NumberValue(big.NewFloat(*d.RxBps.Value()))
	}
	if d.RxBytes.Value() != nil {
		rxBytes = types.Int64Value(*d.RxBytes.Value())
	}
	if d.RxPkts.Value() != nil {
		rxPkts = types.Int64Value(int64(*d.RxPkts.Value()))
	}
	if d.Serial.Value() != nil {
		serial = types.StringValue(*d.Serial.Value())
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.Status.Value() != nil {
		status = types.StringValue(*d.Status.Value())
	}
	if d.SwitchRedundancy != nil {
		switchRedundancy = SwitchRedundancySdkToTerraform(ctx, diags, d.SwitchRedundancy)
	}
	if d.TxBps.Value() != nil {
		txBps = types.NumberValue(big.NewFloat(*d.TxBps.Value()))
	}
	if d.TxBytes.Value() != nil {
		txBytes = types.NumberValue(big.NewFloat(*d.TxBytes.Value()))
	}
	if d.TxPkts.Value() != nil {
		txPkts = types.NumberValue(big.NewFloat(*d.TxPkts.Value()))
	}
	if d.Uptime.Value() != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime.Value()))
	}
	if d.UsbStat != nil {
		usbStat = usbStatsSdkToTerraform(ctx, diags, d.UsbStat)
	}
	if d.Version.Value() != nil {
		version = types.StringValue(*d.Version.Value())
	}
	if d.X.Value() != nil {
		x = types.Float64Value(*d.X.Value())
	}
	if d.Y.Value() != nil {
		y = types.Float64Value(*d.Y.Value())
	}

	dataMapValue := map[string]attr.Value{
		"auto_placement":       autoPlacement,
		"auto_upgrade_stat":    autoUpgradeStat,
		"ble_stat":             bleStat,
		"cert_expiry":          certExpiry,
		"config_reverted":      configReverted,
		"cpu_system":           cpuSystem,
		"cpu_util":             cpuUtil,
		"created_time":         createdTime,
		"deviceprofile_id":     deviceprofileId,
		"env_stat":             envStat,
		"esl_stat":             eslStat,
		"ext_ip":               extIp,
		"fwupdate":             fwupdate,
		"hw_rev":               hwRev,
		"id":                   id,
		"inactive_wired_vlans": inactiveWiredVlans,
		"iot_stat":             iotStat,
		"ip":                   ip,
		"ip_config":            ipConfig,
		"ip_stat":              ipStat,
		"l2tp_stat":            l2tpStat,
		"last_seen":            lastSeen,
		"last_trouble":         lastTrouble,
		"led":                  led,
		"lldp_stat":            lldpStat,
		"locating":             locating,
		"locked":               locked,
		"mac":                  mac,
		"map_id":               mapId,
		"mem_used_kb":          memUsedKb,
		"mesh_downlinks":       meshDownlinks,
		"mesh_uplink":          meshUplink,
		"model":                model,
		"modified_time":        modifiedTime,
		"mount":                mount,
		"name":                 name,
		"notes":                notes,
		"num_clients":          numClients,
		"org_id":               orgId,
		"port_stat":            portStat,
		"power_budget":         powerBudget,
		"power_constrained":    powerConstrained,
		"power_opmode":         powerOpmode,
		"power_src":            powerSrc,
		"radio_stat":           radioStat,
		"rx_bps":               rxBps,
		"rx_bytes":             rxBytes,
		"rx_pkts":              rxPkts,
		"serial":               serial,
		"site_id":              siteId,
		"status":               status,
		"switch_redundancy":    switchRedundancy,
		"tx_bps":               txBps,
		"tx_bytes":             txBytes,
		"tx_pkts":              txPkts,
		"uptime":               uptime,
		"usb_stat":             usbStat,
		"version":              version,
		"x":                    x,
		"y":                    y,
	}
	data, e := NewDeviceApStatsValue(DeviceApStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
