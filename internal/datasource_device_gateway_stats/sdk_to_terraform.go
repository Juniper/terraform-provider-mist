package datasource_device_gateway_stats

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

func SdkToTerraform(ctx context.Context, l []models.StatsDevice) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		gw_js, e := d.MarshalJSON()
		if e != nil {
			diags.AddError("Unable to unMarshal Gateway Stats", e.Error())
		} else {
			gw := models.StatsGateway{}
			e := json.Unmarshal(gw_js, &gw)
			if e != nil {
				diags.AddError("Unable to unMarshal Switch Stats", e.Error())
			}
			elem := deviceGatewayStatSdkToTerraform(ctx, &diags, &gw)
			elements = append(elements, elem)
		}
	}

	dataSet, err := types.SetValue(DeviceGatewayStatsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func deviceGatewayStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsGateway) DeviceGatewayStatsValue {

	var ap_redundancy basetypes.ObjectValue = types.ObjectNull(ApRedundancyValue{}.AttributeTypes(ctx))
	var arp_table_stats basetypes.ObjectValue = types.ObjectNull(ArpTableStatsValue{}.AttributeTypes(ctx))
	var cert_expiry basetypes.Int64Value
	var cluster_config basetypes.ObjectValue = types.ObjectNull(ClusterConfigValue{}.AttributeTypes(ctx))
	var cluster_stat basetypes.ObjectValue = types.ObjectNull(ClusterStatValue{}.AttributeTypes(ctx))
	var conductor_name basetypes.StringValue
	var config_status basetypes.StringValue
	var cpu2_stat basetypes.ObjectValue = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var cpu_stat basetypes.ObjectValue = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var created_time basetypes.Int64Value
	var deviceprofile_id basetypes.StringValue
	var dhcpd2_stat basetypes.MapValue = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var dhcpd_stat basetypes.MapValue = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var ext_ip basetypes.StringValue
	var fwupdate basetypes.ObjectValue = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var has_pcap basetypes.BoolValue
	var hostname basetypes.StringValue
	var id basetypes.StringValue
	var if2_stat basetypes.MapValue = types.MapNull(IfStatValue{}.Type(ctx))
	var if_stat basetypes.MapValue = types.MapNull(IfStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ip2_stat basetypes.ObjectValue = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var ip_stat basetypes.ObjectValue = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var is_ha basetypes.BoolValue
	var last_seen basetypes.NumberValue
	var mac basetypes.StringValue
	var map_id basetypes.StringValue
	var memory2_stat basetypes.ObjectValue = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var memory_stat basetypes.ObjectValue = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modified_time basetypes.Int64Value
	var module2_stat basetypes.ListValue = types.ListNull(ModuleStatValue{}.Type(ctx))
	var module_stat basetypes.ListValue = types.ListNull(ModuleStatValue{}.Type(ctx))
	var name basetypes.StringValue
	var node_name basetypes.StringValue
	var org_id basetypes.StringValue
	var route_summary_stats basetypes.ObjectValue = types.ObjectNull(RouteSummaryStatsValue{}.AttributeTypes(ctx))
	var router_name basetypes.StringValue
	var serial basetypes.StringValue
	var service2_stat basetypes.MapValue = types.MapNull(ServiceStatValue{}.Type(ctx))
	var service_stat basetypes.MapValue = types.MapNull(ServiceStatValue{}.Type(ctx))
	var service_status basetypes.ObjectValue = types.ObjectNull(ServiceStatusValue{}.AttributeTypes(ctx))
	var site_id basetypes.StringValue
	var spu2_stat basetypes.ListValue = types.ListNull(SpuStatValue{}.Type(ctx))
	var spu_stat basetypes.ListValue = types.ListNull(SpuStatValue{}.Type(ctx))
	var status basetypes.StringValue
	var uptime basetypes.NumberValue
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
	if d.ClusterConfig != nil {
		cluster_config = clusterConfigSdkToTerraform(ctx, diags, d.ClusterConfig)
	}
	if d.ClusterStat != nil {
		cluster_stat = clusterStatsSdkToTerraform(ctx, diags, d.ClusterStat)
	}
	if d.ConductorName != nil {
		conductor_name = types.StringValue(*d.ConductorName)
	}
	if d.ConfigStatus != nil {
		config_status = types.StringValue(*d.ConfigStatus)
	}
	if d.Cpu2Stat != nil {
		cpu2_stat = cpuStatsSdkToTerraform(ctx, diags, d.Cpu2Stat)
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
	if d.Dhcpd2Stat != nil && len(d.Dhcpd2Stat) > 0 {
		dhcpd2_stat = dhcpdStatsSdkToTerraform(ctx, diags, d.Dhcpd2Stat)
	}
	if d.DhcpdStat != nil && len(d.DhcpdStat) > 0 {
		dhcpd_stat = dhcpdStatsSdkToTerraform(ctx, diags, d.DhcpdStat)
	}
	if d.ExtIp.Value() != nil {
		ext_ip = types.StringValue(*d.ExtIp.Value())
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HasPcap.Value() != nil {
		has_pcap = types.BoolValue(*d.HasPcap.Value())
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.If2Stat != nil && len(d.If2Stat) > 0 {
		if2_stat = ifStatsSdkToTerraform(ctx, diags, d.If2Stat)
	}
	if d.IfStat != nil && len(d.IfStat) > 0 {
		if_stat = ifStatsSdkToTerraform(ctx, diags, d.IfStat)
	}
	if d.Ip.Value() != nil {
		ip = types.StringValue(*d.Ip.Value())
	}
	if d.Ip2Stat != nil {
		ip2_stat = ipStatsSdkToTerraform(ctx, diags, d.Ip2Stat)
	}
	if d.IpStat != nil {
		ip_stat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.IsHa.Value() != nil {
		is_ha = types.BoolValue(*d.IsHa.Value())
	}
	if d.LastSeen != nil {
		last_seen = types.NumberValue(big.NewFloat(*d.LastSeen))
	}

	mac = types.StringValue(d.Mac)

	if d.MapId.Value() != nil {
		map_id = types.StringValue(d.MapId.Value().String())
	}
	if d.Memory2Stat != nil {
		memory2_stat = memoryStatSdkToTerraform(ctx, diags, d.Memory2Stat)
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
	if d.Module2Stat != nil {
		module2_stat = moduleStatSdkToTerraform(ctx, diags, d.Module2Stat)
	}
	if d.ModuleStat != nil {
		module_stat = moduleStatSdkToTerraform(ctx, diags, d.ModuleStat)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.NodeName != nil {
		node_name = types.StringValue(*d.NodeName)
	}
	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.RouteSummaryStats != nil {
		route_summary_stats = routeSummaryStatsSdkToTerraform(ctx, diags, d.RouteSummaryStats)
	}
	if d.RouterName != nil {
		router_name = types.StringValue(*d.RouterName)
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if d.Service2Stat != nil && len(d.Service2Stat) > 0 {
		service2_stat = serviceStatsSdkToTerraform(ctx, diags, d.Service2Stat)
	}
	if d.ServiceStat != nil && len(d.ServiceStat) > 0 {
		service_stat = serviceStatsSdkToTerraform(ctx, diags, d.ServiceStat)
	}
	if d.ServiceStatus != nil {
		service_status = serviceStatusSdkToTerraform(ctx, diags, d.ServiceStatus)
	}
	if d.SiteId != nil {
		site_id = types.StringValue(d.SiteId.String())
	}
	if d.Spu2Stat != nil {
		spu2_stat = spuStatsSdkToTerraform(ctx, diags, d.Spu2Stat)
	}
	if d.SpuStat != nil {
		spu_stat = spuStatsSdkToTerraform(ctx, diags, d.SpuStat)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.Uptime != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime))
	}
	if d.Version != nil {
		version = types.StringValue(*d.Version)
	}

	data_map_value := map[string]attr.Value{
		"ap_redundancy":       ap_redundancy,
		"arp_table_stats":     arp_table_stats,
		"cert_expiry":         cert_expiry,
		"cluster_config":      cluster_config,
		"cluster_stat":        cluster_stat,
		"conductor_name":      conductor_name,
		"config_status":       config_status,
		"cpu2_stat":           cpu2_stat,
		"cpu_stat":            cpu_stat,
		"created_time":        created_time,
		"deviceprofile_id":    deviceprofile_id,
		"dhcpd2_stat":         dhcpd2_stat,
		"dhcpd_stat":          dhcpd_stat,
		"ext_ip":              ext_ip,
		"fwupdate":            fwupdate,
		"has_pcap":            has_pcap,
		"hostname":            hostname,
		"id":                  id,
		"if2_stat":            if2_stat,
		"if_stat":             if_stat,
		"ip":                  ip,
		"ip2_stat":            ip2_stat,
		"ip_stat":             ip_stat,
		"is_ha":               is_ha,
		"last_seen":           last_seen,
		"mac":                 mac,
		"map_id":              map_id,
		"memory2_stat":        memory2_stat,
		"memory_stat":         memory_stat,
		"model":               model,
		"modified_time":       modified_time,
		"module2_stat":        module2_stat,
		"module_stat":         module_stat,
		"name":                name,
		"node_name":           node_name,
		"org_id":              org_id,
		"route_summary_stats": route_summary_stats,
		"router_name":         router_name,
		"serial":              serial,
		"service2_stat":       service2_stat,
		"service_stat":        service_stat,
		"service_status":      service_status,
		"site_id":             site_id,
		"spu2_stat":           spu2_stat,
		"spu_stat":            spu_stat,
		"status":              status,
		"uptime":              uptime,
		"version":             version,
	}
	data, e := NewDeviceGatewayStatsValue(DeviceGatewayStatsValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}
