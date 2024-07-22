package datasource_device_switch_stats

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, l []models.ListOrgDevicesStatsResponse) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		sw_js, e := d.MarshalJSON()
		if e != nil {
			diags.AddError("Unable to Marshal Switch Stats", e.Error())
		} else {
			sw := models.SwitchStats{}
			e := json.Unmarshal(sw_js, &sw)
			if e != nil {
				diags.AddError("Unable to unMarshal Switch Stats", e.Error())
			}
			elem := deviceSwitchStatSdkToTerraform(ctx, &diags, &sw)
			elements = append(elements, elem)
		}
	}

	dataSet, err := types.SetValue(DeviceSwitchStatsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func deviceSwitchStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchStats) DeviceSwitchStatsValue {

	var ap_redundancy basetypes.ObjectValue = types.ObjectNull(ApRedundancyValue{}.AttributeTypes(ctx))
	var arp_table_stats basetypes.ObjectValue = types.ObjectNull(ArpTableStatsValue{}.AttributeTypes(ctx))
	var cert_expiry basetypes.Int64Value
	var clients basetypes.ListValue = types.ListNull(ClientsValue{}.Type(ctx))
	var clients_stats basetypes.ObjectValue = types.ObjectNull(ClientsStatsValue{}.AttributeTypes(ctx))
	var config_status basetypes.StringValue
	var cpu_stat basetypes.ObjectValue = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var created_time basetypes.Int64Value
	var deviceprofile_id basetypes.StringValue
	var dhcpd_stat basetypes.MapValue = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var evpntopo_id basetypes.StringValue
	var fw_versions_outofsync basetypes.BoolValue
	var fwupdate basetypes.ObjectValue = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var has_pcap basetypes.BoolValue
	var hostname basetypes.StringValue
	var hw_rev basetypes.StringValue
	var id basetypes.StringValue
	var if_stat basetypes.MapValue = types.MapNull(IfStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ip_stat basetypes.ObjectValue = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var last_seen basetypes.NumberValue
	var last_trouble basetypes.ObjectValue = types.ObjectNull(LastTroubleValue{}.AttributeTypes(ctx))
	var mac basetypes.StringValue
	var mac_table_stats basetypes.ObjectValue = types.ObjectNull(MacTableStatsValue{}.AttributeTypes(ctx))
	var map_id basetypes.StringValue
	var memory_stat basetypes.ObjectValue = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modified_time basetypes.Int64Value
	var module_stat basetypes.ListValue = types.ListNull(ModuleStatValue{}.Type(ctx))
	var name basetypes.StringValue
	var org_id basetypes.StringValue
	var route_summary_stats basetypes.ObjectValue = types.ObjectNull(RouteSummaryStatsValue{}.AttributeTypes(ctx))
	var serial basetypes.StringValue
	var service_stat basetypes.MapValue = types.MapNull(ServiceStatValue{}.Type(ctx))
	var site_id basetypes.StringValue
	var status basetypes.StringValue
	var uptime basetypes.NumberValue
	var vc_mac basetypes.StringValue
	var vc_setup_info basetypes.ObjectValue = types.ObjectNull(VcSetupInfoValue{}.AttributeTypes(ctx))
	var version basetypes.StringValue

	if d.ApRedundancy != nil {
		ap_redundancy = apRedundancySdkToTerraform(ctx, diags, d.ApRedundancy)
	}
	if d.ArpTableStats != nil {
		arp_table_stats = arpTableStatsSdkToTerraform(ctx, diags, d.ArpTableStats)
	}
	if d.CertExpiry != nil {
		cert_expiry = types.Int64Value(*d.CertExpiry)
	}
	if d.Clients != nil {
		clients = clientsSdkToTerraform(ctx, diags, d.Clients)
	}
	if d.ClientsStats != nil {
		clients_stats = clientsStatsSdkToTerraform(ctx, diags, d.ClientsStats)
	}
	if d.ConfigStatus != nil {
		config_status = types.StringValue(*d.ConfigStatus)
	}
	if d.CpuStat != nil {
		cpu_stat = cpuStatsSdkToTerraform(ctx, diags, d.CpuStat)
	}
	if d.CreatedTime != nil {
		created_time = types.Int64Value(int64(*d.CreatedTime))
	}
	if d.DeviceprofileId.Value() != nil {
		deviceprofile_id = types.StringValue(d.DeviceprofileId.Value().String())
	}
	if d.DhcpdStat != nil && len(d.DhcpdStat) > 0 {
		dhcpd_stat = dhcpStatsSdkToTerraform(ctx, diags, d.DhcpdStat)
	}
	if d.EvpntopoId.Value() != nil {
		evpntopo_id = types.StringValue(d.EvpntopoId.Value().String())
	}
	if d.FwVersionsOutofsync != nil {
		fw_versions_outofsync = types.BoolValue(*d.FwVersionsOutofsync)
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HasPcap != nil {
		has_pcap = types.BoolValue(*d.HasPcap)
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.HwRev != nil {
		hw_rev = types.StringValue(*d.HwRev)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.IfStat != nil && len(d.IfStat) > 0 {
		if_stat = ifStatsSdkToTerraform(ctx, diags, d.IfStat)
	}
	if d.Ip != nil {
		ip = types.StringValue(*d.Ip)
	}
	if d.IpStat != nil {
		ip_stat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.LastSeen != nil {
		last_seen = types.NumberValue(big.NewFloat(*d.LastSeen))
	}
	if d.LastTrouble != nil {
		last_trouble = lastTroubleSdkToTerraform(ctx, diags, d.LastTrouble)
	}
	if d.Mac != nil {
		mac = types.StringValue(*d.Mac)
	}
	if d.MacTableStats != nil {
		mac_table_stats = macTableStatSdkToTerraform(ctx, diags, d.MacTableStats)
	}
	if d.MapId.Value() != nil {
		map_id = types.StringValue(d.MapId.Value().String())
	}
	if d.MemoryStat != nil {
		memory_stat = memoryStatSdkToTerraform(ctx, diags, d.MemoryStat)
	}
	if d.Model != nil {
		model = types.StringValue(*d.Model)
	}
	if d.ModifiedTime != nil {
		modified_time = types.Int64Value(int64(*d.ModifiedTime))
	}
	if d.ModuleStat != nil {
		module_stat = moduleStatSdkToTerraform(ctx, diags, d.ModuleStat)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.RouteSummaryStats != nil {
		route_summary_stats = routeSummaryStatsSdkToTerraform(ctx, diags, d.RouteSummaryStats)
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if d.ServiceStat != nil && len(d.ServiceStat) > 0 {
		service_stat = serviceStatsSdkToTerraform(ctx, diags, d.ServiceStat)
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.Uptime != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime))
	}
	if d.VcMac.Value() != nil {
		vc_mac = types.StringValue(*d.VcMac.Value())
	}
	if d.VcSetupInfo != nil {
		vc_setup_info = vcSetupInfoSdkToTerraform(ctx, diags, d.VcSetupInfo)
	}
	if d.Version != nil {
		version = types.StringValue(*d.Version)
	}

	data_map_value := map[string]attr.Value{
		"ap_redundancy":         ap_redundancy,
		"arp_table_stats":       arp_table_stats,
		"cert_expiry":           cert_expiry,
		"clients":               clients,
		"clients_stats":         clients_stats,
		"config_status":         config_status,
		"cpu_stat":              cpu_stat,
		"created_time":          created_time,
		"deviceprofile_id":      deviceprofile_id,
		"dhcpd_stat":            dhcpd_stat,
		"evpntopo_id":           evpntopo_id,
		"fw_versions_outofsync": fw_versions_outofsync,
		"fwupdate":              fwupdate,
		"has_pcap":              has_pcap,
		"hostname":              hostname,
		"hw_rev":                hw_rev,
		"id":                    id,
		"if_stat":               if_stat,
		"ip":                    ip,
		"ip_stat":               ip_stat,
		"last_seen":             last_seen,
		"last_trouble":          last_trouble,
		"mac":                   mac,
		"mac_table_stats":       mac_table_stats,
		"map_id":                map_id,
		"memory_stat":           memory_stat,
		"model":                 model,
		"modified_time":         modified_time,
		"module_stat":           module_stat,
		"name":                  name,
		"org_id":                org_id,
		"route_summary_stats":   route_summary_stats,
		"serial":                serial,
		"service_stat":          service_stat,
		"site_id":               site_id,
		"status":                status,
		"uptime":                uptime,
		"vc_mac":                vc_mac,
		"vc_setup_info":         vc_setup_info,
		"version":               version,
	}
	data, e := NewDeviceSwitchStatsValue(DeviceSwitchStatsValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
