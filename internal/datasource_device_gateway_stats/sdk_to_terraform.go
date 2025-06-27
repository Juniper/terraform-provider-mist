package datasource_device_gateway_stats

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func SdkToTerraform(ctx context.Context, l *[]models.StatsGateway, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := deviceGatewayStatSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func deviceGatewayStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsGateway) DeviceGatewayStatsValue {

	var apRedundancy = types.ObjectNull(ApRedundancyValue{}.AttributeTypes(ctx))
	var arpTableStats = types.ObjectNull(ArpTableStatsValue{}.AttributeTypes(ctx))
	var bgpPeers = types.ListNull(BgpPeersValue{}.Type(ctx))
	var certExpiry basetypes.Int64Value
	var clusterConfig = types.ObjectNull(ClusterConfigValue{}.AttributeTypes(ctx))
	var clusterStat = types.ObjectNull(ClusterStatValue{}.AttributeTypes(ctx))
	var conductorName basetypes.StringValue
	var configStatus basetypes.StringValue
	var cpu2Stat = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var cpuStat = types.ObjectNull(CpuStatValue{}.AttributeTypes(ctx))
	var createdTime basetypes.Float64Value
	var deviceprofileId basetypes.StringValue
	var dhcpd2Stat = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var dhcpdStat = types.MapNull(DhcpdStatValue{}.Type(ctx))
	var extIp basetypes.StringValue
	var fwupdate = types.ObjectNull(FwupdateValue{}.AttributeTypes(ctx))
	var hasPcap basetypes.BoolValue
	var hostname basetypes.StringValue
	var id basetypes.StringValue
	var if2Stat = types.MapNull(IfStatValue{}.Type(ctx))
	var ifStat = types.MapNull(IfStatValue{}.Type(ctx))
	var ip basetypes.StringValue
	var ip2Stat = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var ipStat = types.ObjectNull(IpStatValue{}.AttributeTypes(ctx))
	var isHa basetypes.BoolValue
	var lastSeen basetypes.Float64Value
	var mac basetypes.StringValue
	var mapId basetypes.StringValue
	var memory2Stat = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var memoryStat = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
	var model basetypes.StringValue
	var modifiedTime basetypes.Float64Value
	var module2Stat = types.ListNull(ModuleStatValue{}.Type(ctx))
	var moduleStat = types.ListNull(ModuleStatValue{}.Type(ctx))
	var name basetypes.StringValue
	var nodeName basetypes.StringValue
	var orgId basetypes.StringValue
	var ports = types.ListNull(PortsValue{}.Type(ctx))
	var routeSummaryStats = types.ObjectNull(RouteSummaryStatsValue{}.AttributeTypes(ctx))
	var routerName basetypes.StringValue
	var serial basetypes.StringValue
	var service2Stat = types.MapNull(ServiceStatValue{}.Type(ctx))
	var serviceStat = types.MapNull(ServiceStatValue{}.Type(ctx))
	var serviceStatus = types.ObjectNull(ServiceStatusValue{}.AttributeTypes(ctx))
	var siteId basetypes.StringValue
	var spu2Stat = types.ListNull(SpuStatValue{}.Type(ctx))
	var spuStat = types.ListNull(SpuStatValue{}.Type(ctx))
	var status basetypes.StringValue
	var tunnels = types.ListNull(TunnelsValue{}.Type(ctx))
	var uptime basetypes.NumberValue
	var version basetypes.StringValue
	var vpnPeers = types.ListNull(VpnPeersValue{}.Type(ctx))

	if d.ApRedundancy != nil {
		apRedundancy = apRedundancySdkToTerraform(ctx, diags, d.ApRedundancy)
	}
	if d.ArpTableStats != nil {
		arpTableStats = arpTableStatsSdkToTerraform(ctx, diags, d.ArpTableStats)
	}
	if d.BgpPeers != nil {
		bgpPeers = bgpPeersSdkToTerraform(ctx, diags, d.BgpPeers)
	}
	if d.CertExpiry != nil {
		certExpiry = types.Int64Value(*d.CertExpiry)
	}
	if d.ClusterConfig != nil {
		clusterConfig = clusterConfigSdkToTerraform(ctx, diags, d.ClusterConfig)
	}
	if d.ClusterStat != nil {
		clusterStat = clusterStatsSdkToTerraform(ctx, diags, d.ClusterStat)
	}
	if d.ConductorName != nil {
		conductorName = types.StringValue(*d.ConductorName)
	}
	if d.ConfigStatus != nil {
		configStatus = types.StringValue(*d.ConfigStatus)
	}
	if d.Cpu2Stat != nil {
		cpu2Stat = cpuStatsSdkToTerraform(ctx, diags, d.Cpu2Stat)
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
	if len(d.Dhcpd2Stat) > 0 {
		dhcpd2Stat = dhcpdStatsSdkToTerraform(ctx, diags, d.Dhcpd2Stat)
	}
	if len(d.DhcpdStat) > 0 {
		dhcpdStat = dhcpdStatsSdkToTerraform(ctx, diags, d.DhcpdStat)
	}
	if d.ExtIp.Value() != nil {
		extIp = types.StringValue(*d.ExtIp.Value())
	}
	if d.Fwupdate != nil {
		fwupdate = fwupdateSdkToTerraform(ctx, diags, d.Fwupdate)
	}
	if d.HasPcap.Value() != nil {
		hasPcap = types.BoolValue(*d.HasPcap.Value())
	}
	if d.Hostname != nil {
		hostname = types.StringValue(*d.Hostname)
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if len(d.If2Stat) > 0 {
		if2Stat = ifStatsSdkToTerraform(ctx, diags, d.If2Stat)
	}
	if len(d.IfStat) > 0 {
		ifStat = ifStatsSdkToTerraform(ctx, diags, d.IfStat)
	}
	if d.Ip.Value() != nil {
		ip = types.StringValue(*d.Ip.Value())
	}
	if d.Ip2Stat != nil {
		ip2Stat = ipStatsSdkToTerraform(ctx, diags, d.Ip2Stat)
	}
	if d.IpStat != nil {
		ipStat = ipStatsSdkToTerraform(ctx, diags, d.IpStat)
	}
	if d.IsHa.Value() != nil {
		isHa = types.BoolValue(*d.IsHa.Value())
	}
	if d.LastSeen.Value() != nil {
		lastSeen = types.Float64Value(*d.LastSeen.Value())
	}

	mac = types.StringValue(d.Mac)

	if d.MapId.Value() != nil {
		mapId = types.StringValue(d.MapId.Value().String())
	}
	if d.Memory2Stat != nil {
		memory2Stat = memoryStatSdkToTerraform(ctx, diags, d.Memory2Stat)
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
	if d.Module2Stat != nil {
		module2Stat = moduleStatSdkToTerraform(ctx, diags, d.Module2Stat)
	}
	if d.ModuleStat != nil {
		moduleStat = moduleStatSdkToTerraform(ctx, diags, d.ModuleStat)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.NodeName != nil {
		nodeName = types.StringValue(*d.NodeName)
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
	if d.RouterName != nil {
		routerName = types.StringValue(*d.RouterName)
	}
	if d.Serial != nil {
		serial = types.StringValue(*d.Serial)
	}
	if len(d.Service2Stat) > 0 {
		service2Stat = serviceStatsSdkToTerraform(ctx, diags, d.Service2Stat)
	}
	if len(d.ServiceStat) > 0 {
		serviceStat = serviceStatsSdkToTerraform(ctx, diags, d.ServiceStat)
	}
	if d.ServiceStatus != nil {
		serviceStatus = serviceStatusSdkToTerraform(ctx, diags, d.ServiceStatus)
	}
	if d.SiteId != nil {
		siteId = types.StringValue(d.SiteId.String())
	}
	if d.Spu2Stat != nil {
		spu2Stat = spuStatsSdkToTerraform(ctx, diags, d.Spu2Stat)
	}
	if d.SpuStat != nil {
		spuStat = spuStatsSdkToTerraform(ctx, diags, d.SpuStat)
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}
	if d.Tunnels != nil {
		tunnels = tunnelsSdkToTerraform(ctx, diags, d.Tunnels)
	}
	if d.Uptime.Value() != nil {
		uptime = types.NumberValue(big.NewFloat(*d.Uptime.Value()))
	}
	if d.Version.Value() != nil {
		version = types.StringValue(*d.Version.Value())
	}
	if d.VpnPeers != nil {
		vpnPeers = vpnPeersSdkToTerraform(ctx, diags, d.VpnPeers)
	}

	dataMapValue := map[string]attr.Value{
		"ap_redundancy":       apRedundancy,
		"arp_table_stats":     arpTableStats,
		"bgp_peers":           bgpPeers,
		"cert_expiry":         certExpiry,
		"cluster_config":      clusterConfig,
		"cluster_stat":        clusterStat,
		"conductor_name":      conductorName,
		"config_status":       configStatus,
		"cpu2_stat":           cpu2Stat,
		"cpu_stat":            cpuStat,
		"created_time":        createdTime,
		"deviceprofile_id":    deviceprofileId,
		"dhcpd2_stat":         dhcpd2Stat,
		"dhcpd_stat":          dhcpdStat,
		"ext_ip":              extIp,
		"fwupdate":            fwupdate,
		"has_pcap":            hasPcap,
		"hostname":            hostname,
		"id":                  id,
		"if2_stat":            if2Stat,
		"if_stat":             ifStat,
		"ip":                  ip,
		"ip2_stat":            ip2Stat,
		"ip_stat":             ipStat,
		"is_ha":               isHa,
		"last_seen":           lastSeen,
		"mac":                 mac,
		"map_id":              mapId,
		"memory2_stat":        memory2Stat,
		"memory_stat":         memoryStat,
		"model":               model,
		"modified_time":       modifiedTime,
		"module2_stat":        module2Stat,
		"module_stat":         moduleStat,
		"name":                name,
		"node_name":           nodeName,
		"org_id":              orgId,
		"ports":               ports,
		"route_summary_stats": routeSummaryStats,
		"router_name":         routerName,
		"serial":              serial,
		"service2_stat":       service2Stat,
		"service_stat":        serviceStat,
		"service_status":      serviceStatus,
		"site_id":             siteId,
		"spu2_stat":           spu2Stat,
		"spu_stat":            spuStat,
		"status":              status,
		"tunnels":             tunnels,
		"uptime":              uptime,
		"version":             version,
		"vpn_peers":           vpnPeers,
	}
	data, e := NewDeviceGatewayStatsValue(DeviceGatewayStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
