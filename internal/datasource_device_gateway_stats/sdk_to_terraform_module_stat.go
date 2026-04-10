package datasource_device_gateway_stats

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func moduleStatFanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemFansItems) basetypes.ListValue {

	var dataList []FansValue
	for _, d := range l {
		var airflow basetypes.StringValue
		var name basetypes.StringValue
		var rpm basetypes.Int64Value
		var status basetypes.StringValue

		if d.Airflow != nil {
			airflow = types.StringValue(*d.Airflow)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Rpm != nil {
			rpm = types.Int64Value(int64(*d.Rpm))
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		dataMapValue := map[string]attr.Value{
			"airflow": airflow,
			"name":    name,
			"rpm":     rpm,
			"status":  status,
		}
		data, e := NewFansValue(FansValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, FansValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}

func moduleStatPoeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ModuleStatItemPoe) basetypes.ObjectValue {

	var maxPower basetypes.NumberValue
	var powerDraw basetypes.NumberValue
	var status basetypes.StringValue

	if d.MaxPower != nil {
		maxPower = types.NumberValue(big.NewFloat(*d.MaxPower))
	}
	if d.PowerDraw != nil {
		powerDraw = types.NumberValue(big.NewFloat(*d.PowerDraw))
	}
	if d.Status != nil {
		status = types.StringValue(*d.Status)
	}

	dataMapValue := map[string]attr.Value{
		"max_power":  maxPower,
		"power_draw": powerDraw,
		"status":     status,
	}
	data, e := basetypes.NewObjectValue(PoeValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
func moduleStatPsusSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemPsusItem) basetypes.ListValue {

	var dataList []PsusValue
	for _, d := range l {
		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		dataMapValue := map[string]attr.Value{
			"name":   name,
			"status": status,
		}
		data, e := NewPsusValue(PsusValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, PsusValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
func moduleStatTemperatureSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemTemperaturesItem) basetypes.ListValue {

	var dataList []TemperaturesValue
	for _, d := range l {
		var celsius basetypes.NumberValue
		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Celsius != nil {
			celsius = types.NumberValue(big.NewFloat(*d.Celsius))
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		dataMapValue := map[string]attr.Value{
			"celsius": celsius,
			"name":    name,
			"status":  status,
		}
		data, e := NewTemperaturesValue(TemperaturesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, TemperaturesValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
func moduleStatVcLinksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemVcLinksItem) basetypes.ListValue {

	var dataList []VcLinksValue
	for _, d := range l {
		var neighborModuleIdx basetypes.Int64Value
		var neighborPortId basetypes.StringValue
		var portId basetypes.StringValue

		if d.NeighborModuleIdx != nil {
			neighborModuleIdx = types.Int64Value(int64(*d.NeighborModuleIdx))
		}
		if d.NeighborPortId != nil {
			neighborPortId = types.StringValue(*d.NeighborPortId)
		}
		if d.PortId != nil {
			portId = types.StringValue(*d.PortId)
		}

		dataMapValue := map[string]attr.Value{
			"neighbor_module_idx": neighborModuleIdx,
			"neighbor_port_id":    neighborPortId,
			"port_id":             portId,
		}
		data, e := NewVcLinksValue(VcLinksValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, VcLinksValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
func moduleStatNetworkResourceSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemNetworkResource) basetypes.ListValue {

	var dataList []NetworkResourcesValue
	for _, d := range l {
		var count basetypes.Int64Value
		var limit basetypes.Int64Value
		var nrType basetypes.StringValue

		if d.Count != nil {
			count = types.Int64Value(int64(*d.Count))
		}
		if d.Limit != nil {
			limit = types.Int64Value(int64(*d.Limit))
		}
		if d.Type != nil {
			nrType = types.StringValue(*d.Type)
		}

		dataMapValue := map[string]attr.Value{
			"count": count,
			"limit": limit,
			"type":  nrType,
		}
		data, e := NewNetworkResourcesValue(NetworkResourcesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, NetworkResourcesValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}

func moduleStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.StatsGatewayModuleStatItem) basetypes.ListValue {

	var dataList []ModuleStatValue
	for _, d := range l {
		var backupVersion basetypes.StringValue
		var biosVersion basetypes.StringValue
		var bootPartition basetypes.StringValue
		var cpldVersion basetypes.StringValue
		var fans = types.ListNull(FansValue{}.Type(ctx))
		var fpgaVersion basetypes.StringValue
		var lastSeen basetypes.Float64Value
		var locating basetypes.BoolValue
		var mac basetypes.StringValue
		var memoryStat = types.ObjectNull(MemoryStatValue{}.AttributeTypes(ctx))
		var model basetypes.StringValue
		var networkResources = types.ListNull(NetworkResourcesValue{}.Type(ctx))
		var opticsCpldVersion basetypes.StringValue
		var pendingVersion basetypes.StringValue
		var poe = types.ObjectNull(PoeValue{}.AttributeTypes(ctx))
		var poeVersion basetypes.StringValue
		var powerCpldVersion basetypes.StringValue
		var psus = types.ListNull(PsusValue{}.Type(ctx))
		var reFpgaVersion basetypes.StringValue
		var recoveryVersion basetypes.StringValue
		var serial basetypes.StringValue
		var status basetypes.StringValue
		var temperatures = types.ListNull(TemperaturesValue{}.Type(ctx))
		var tmcFpgaVersion basetypes.StringValue
		var ubootVersion basetypes.StringValue
		var uptime basetypes.Int64Value
		var vcLinks = types.ListNull(VcLinksValue{}.Type(ctx))
		var vcMode basetypes.StringValue
		var vcRole basetypes.StringValue
		var vcState basetypes.StringValue
		var version basetypes.StringValue

		if d.BackupVersion.Value() != nil {
			backupVersion = types.StringValue(*d.BackupVersion.Value())
		}
		if d.BootPartition != nil {
			bootPartition = types.StringValue(*d.BootPartition)
		}
		if d.BiosVersion.Value() != nil {
			biosVersion = types.StringValue(*d.BiosVersion.Value())
		}
		if d.CpldVersion.Value() != nil {
			cpldVersion = types.StringValue(*d.CpldVersion.Value())
		}
		if len(d.Fans) > 0 {
			fans = moduleStatFanSdkToTerraform(ctx, diags, d.Fans)
		}
		if d.FpgaVersion.Value() != nil {
			fpgaVersion = types.StringValue(*d.FpgaVersion.Value())
		}
		if d.LastSeen.Value() != nil {
			lastSeen = types.Float64Value(*d.LastSeen.Value())
		}
		if d.Locating != nil {
			locating = types.BoolValue(*d.Locating)
		}
		if d.Mac != nil {
			mac = types.StringValue(*d.Mac)
		}
		if d.MemoryStat != nil {
			memoryStat = memoryStatSdkToTerraform(ctx, diags, d.MemoryStat)
		}
		if d.Model.Value() != nil {
			model = types.StringValue(*d.Model.Value())
		}
		if d.NetworkResources != nil {
			networkResources = moduleStatNetworkResourceSdkToTerraform(ctx, diags, d.NetworkResources)
		}
		if d.OpticsCpldVersion.Value() != nil {
			opticsCpldVersion = types.StringValue(*d.OpticsCpldVersion.Value())
		}
		if d.PendingVersion.Value() != nil {
			pendingVersion = types.StringValue(*d.PendingVersion.Value())
		}
		if d.Poe != nil {
			poe = moduleStatPoeSdkToTerraform(ctx, diags, d.Poe)
		}
		if d.PoeVersion.Value() != nil {
			poeVersion = types.StringValue(*d.PoeVersion.Value())
		}
		if d.PowerCpldVersion.Value() != nil {
			powerCpldVersion = types.StringValue(*d.PowerCpldVersion.Value())
		}
		if d.Psus != nil {
			psus = moduleStatPsusSdkToTerraform(ctx, diags, d.Psus)
		}
		if d.ReFpgaVersion.Value() != nil {
			reFpgaVersion = types.StringValue(*d.ReFpgaVersion.Value())
		}
		if d.RecoveryVersion.Value() != nil {
			recoveryVersion = types.StringValue(*d.RecoveryVersion.Value())
		}
		if d.Serial.Value() != nil {
			serial = types.StringValue(*d.Serial.Value())
		}
		if d.Status.Value() != nil {
			status = types.StringValue(*d.Status.Value())
		}
		if d.Temperatures != nil {
			temperatures = moduleStatTemperatureSdkToTerraform(ctx, diags, d.Temperatures)
		}
		if d.TmcFpgaVersion.Value() != nil {
			tmcFpgaVersion = types.StringValue(*d.TmcFpgaVersion.Value())
		}
		if d.UbootVersion.Value() != nil {
			ubootVersion = types.StringValue(*d.UbootVersion.Value())
		}
		if d.Uptime.Value() != nil {
			uptime = types.Int64Value(int64(*d.Uptime.Value()))
		}
		if d.VcLinks != nil {
			vcLinks = moduleStatVcLinksSdkToTerraform(ctx, diags, d.VcLinks)
		}
		if d.VcMode.Value() != nil {
			vcMode = types.StringValue(*d.VcMode.Value())
		}
		if d.VcRole.Value() != nil {
			vcRole = types.StringValue(*d.VcRole.Value())
		}
		if d.VcState.Value() != nil {
			vcState = types.StringValue(*d.VcState.Value())
		}
		if d.Version.Value() != nil {
			version = types.StringValue(*d.Version.Value())
		}

		dataMapValue := map[string]attr.Value{
			"backup_version":      backupVersion,
			"boot_partition":      bootPartition,
			"bios_version":        biosVersion,
			"cpld_version":        cpldVersion,
			"fans":                fans,
			"fpga_version":        fpgaVersion,
			"last_seen":           lastSeen,
			"locating":            locating,
			"mac":                 mac,
			"memory_stat":         memoryStat,
			"model":               model,
			"network_resources":   networkResources,
			"optics_cpld_version": opticsCpldVersion,
			"pending_version":     pendingVersion,
			"poe":                 poe,
			"poe_version":         poeVersion,
			"power_cpld_version":  powerCpldVersion,
			"psus":                psus,
			"re_fpga_version":     reFpgaVersion,
			"recovery_version":    recoveryVersion,
			"serial":              serial,
			"status":              status,
			"temperatures":        temperatures,
			"tmc_fpga_version":    tmcFpgaVersion,
			"uboot_version":       ubootVersion,
			"uptime":              uptime,
			"vc_links":            vcLinks,
			"vc_mode":             vcMode,
			"vc_role":             vcRole,
			"vc_state":            vcState,
			"version":             version,
		}
		data, e := NewModuleStatValue(ModuleStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ModuleStatValue{}.Type(ctx), dataList)
	diags.Append(e...)

	return r
}
