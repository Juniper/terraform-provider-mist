package datasource_device_switch_stats

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, l *[]models.StatsSwitch, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := deviceSwitchStatSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func deviceSwitchStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitch) DeviceSwitchStatsValue {

	var apRedundancy = types.ObjectNull(ApRedundancyValue{}.AttributeTypes(ctx))
	var arpTableStats = types.ObjectNull(ArpTableStatsValue{}.AttributeTypes(ctx))
	var certExpiry basetypes.Int64Value
	var clients = types.ListNull(ClientsValue{}.Type(ctx))
	var clientsStats = types.ObjectNull(ClientsStatsValue{}.AttributeTypes(ctx))
	var configStatus basetypes.StringValue
	var cpuStat = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var createdTime basetypes.Float64Value
	var deviceprofileId basetypes.StringValue
	var dhcpdStat = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var evpntopoId basetypes.StringValue
	var fwVersionsOutofsync basetypes.BoolValue
	var fwupdate = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var hasPcap basetypes.BoolValue
	var hostname basetypes.StringValue
	var hwRev basetypes.StringValue
	var id basetypes.StringValue
	var ifStat = types.MapNull(IfStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ipStat = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var lastSeen basetypes.Float64Value
	var lastTrouble = types.ObjectNull(LastTroubleValue{}.AttributeTypes(ctx))
	var mac basetypes.StringValue
	var macTableStats = types.ObjectNull(MacTableStatsValue{}.AttributeTypes(ctx))
	var mapId basetypes.StringValue
	var memoryStat = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var moduleStat = types.ListNull(ModuleStatValue{}.Type(ctx))
	var name basetypes.StringValue
	var orgId basetypes.StringValue
	var ports = types.ListNull(PortsValue{}.Type(ctx))
	var routeSummaryStats = types.ObjectNull(RouteSummaryStatsValue{}.AttributeTypes(ctx))
	var serial basetypes.StringValue
	var serviceStat = types.MapNull(ServiceStatValue{}.Type(ctx))
	var siteId basetypes.StringValue
	var status basetypes.StringValue
	var uptime basetypes.NumberValue
	var vcMac basetypes.StringValue
	var vcSetupInfo = types.ObjectNull(VcSetupInfoValue{}.AttributeTypes(ctx))
	var version basetypes.StringValue

	if d.ApRedundancy != nil {
		apRedundancy = apRedundancySdkToTerraform(ctx, diags, d.ApRedundancy)
	}
	if d.ArpTableStats != nil {
		arpTableStats = arpTableStatsSdkToTerraform(ctx, diags, d.ArpTableStats)
	}
	if d.CertExpiry != nil {
		certExpiry = types.Int64Value(*d.CertExpiry)
	}
	if d.Clients != nil {
		clients = clientsSdkToTerraform(ctx, diags, d.Clients)
	}
	if d.ClientsStats != nil {
		clientsStats = clientsStatsSdkToTerraform(ctx, diags, d.ClientsStats)
	}
	if d.ConfigStatus != nil {
		configStatus = types.StringValue(*d.ConfigStatus)
	}
	if d.CpuStat != nil {
		cpuStat = cpuStatsSdkToTerraform(ctx, diags, d.CpuStat)
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofileId = types.StringValue(d.DeviceprofileId.Value().String())
	}
	if d.DhcpdStat != nil && len(d.DhcpdStat) > 0 {
		dhcpdStat = dhcpStatsSdkToTerraform(ctx, diags, d.DhcpdStat)
	}
	if d.EvpntopoId.Value() != nil {
		evpntopoId = types.StringValue(d.EvpntopoId.Value().String())
	}
	if d.FwVersionsOutofsync != nil {
		fwVersionsOutofsync = types.BoolValue(*d.FwVersionsOutofsync)
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HasPcap != nil {
		hasPcap = types.BoolValue(*d.HasPcap)
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.HwRev != nil {
		hwRev = types.StringValue(*d.HwRev)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.IfStat != nil && len(d.IfStat) > 0 {
		ifStat = ifStatsSdkToTerraform(ctx, diags, d.IfStat)
	}
	if d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d.IpStat != nil {
		ipStat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.LastSeen.Value() != nil {
		lastSeen = types.Float64Value(*d.LastSeen.Value())
	}
	if d.LastTrouble != nil {
		lastTrouble = lastTroubleSdkToTerraform(ctx, diags, d.LastTrouble)
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.MacTableStats != nil {
		macTableStats = macTableStatSdkToTerraform(ctx, diags, d.MacTableStats)
	}
	if d.MapId.Value() != nil {
		mapId = types.StringValue(d.MapId.Value().String())
	}
	if d.MemoryStat != nil {
		memoryStat = memoryStatSdkToTerraform(ctx, diags, d.MemoryStat)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
	}
	if d.ModuleStat != nil {
		moduleStat = moduleStatSdkToTerraform(ctx, diags, d.ModuleStat)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Ports != nil {
		ports = portsSdkToTerraform(ctx, diags, d.Ports)
	}
	if d.RouteSummaryStats != nil {
		routeSummaryStats = routeSummaryStatsSdkToTerraform(ctx, diags, d.RouteSummaryStats)
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if d.ServiceStat != nil && len(d.ServiceStat) > 0 {
		serviceStat = serviceStatsSdkToTerraform(ctx, diags, d.ServiceStat)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.Uptime.Value() != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime.Value()))
	}
	if d.VcMac.Value() != nil {
		vcMac = types.StringValue(*d.VcMac.Value())
	}
	if d.VcSetupInfo != nil {
		vcSetupInfo = vcSetupInfoSdkToTerraform(ctx, diags, d.VcSetupInfo)
	}
	if d.Version.Value() != nil {
		version = types.StringValue(*d.Version.Value())
	}

	dataMapValue := map[string]attr.Value{
		"ap_redundancy":         apRedundancy,
		"arp_table_stats":       arpTableStats,
		"cert_expiry":           certExpiry,
		"clients":               clients,
		"clients_stats":         clientsStats,
		"config_status":         configStatus,
		"cpu_stat":              cpuStat,
		"created_time":          createdTime,
		"deviceprofile_id":      deviceprofileId,
		"dhcpd_stat":            dhcpdStat,
		"evpntopo_id":           evpntopoId,
		"fw_versions_outofsync": fwVersionsOutofsync,
		"fwupdate":              fwupdate,
		"has_pcap":              hasPcap,
		"hostname":              hostname,
		"hw_rev":                hwRev,
		"id":                    id,
		"if_stat":               ifStat,
		"ip":                    ip,
		"ip_stat":               ipStat,
		"last_seen":             lastSeen,
		"last_trouble":          lastTrouble,
		"mac":                   mac,
		"mac_table_stats":       macTableStats,
		"map_id":                mapId,
		"memory_stat":           memoryStat,
		"model":                 model,
		"modified_time":         modifiedTime,
		"module_stat":           moduleStat,
		"name":                  name,
		"org_id":                orgId,
		"ports":                 ports,
		"route_summary_stats":   routeSummaryStats,
		"serial":                serial,
		"service_stat":          serviceStat,
		"site_id":               siteId,
		"status":                status,
		"uptime":                uptime,
		"vc_mac":                vcMac,
		"vc_setup_info":         vcSetupInfo,
		"version":               version,
	}
	data, e := NewDeviceSwitchStatsValue(DeviceSwitchStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
