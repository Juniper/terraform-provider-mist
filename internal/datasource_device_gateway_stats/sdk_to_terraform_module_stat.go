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

func moduleStatErrorSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemErrorsItems) basetypes.ListValue {

	var data_list = []ErrorsValue{}
	for _, d := range l {
		var feature basetypes.StringValue
		var minimum_version basetypes.StringValue
		var reason basetypes.StringValue
		var since basetypes.Int64Value
		var type_error basetypes.StringValue

		if d.Feature != nil {
			feature = types.StringValue(*d.Feature)
		}
		if d.MinimumVersion != nil {
			minimum_version = types.StringValue(*d.MinimumVersion)
		}
		if d.Reason != nil {
			reason = types.StringValue(*d.Reason)
		}
		// if d.Since != nil {
		since = types.Int64Value(int64(d.Since))
		// }
		// if d.Type != nil {
		type_error = types.StringValue(d.Type)
		//}

		data_map_attr_type := ErrorsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"feature":         feature,
			"minimum_version": minimum_version,
			"reason":          reason,
			"since":           since,
			"type":            type_error,
		}
		data, e := NewErrorsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ErrorsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatFanSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemFansItems) basetypes.ListValue {

	var data_list = []FansValue{}
	for _, d := range l {
		var airflow basetypes.StringValue
		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Airflow != nil {
			airflow = types.StringValue(*d.Airflow)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		data_map_attr_type := FansValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"airflow": airflow,
			"name":    name,
			"status":  status,
		}
		data, e := NewFansValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, FansValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatPicPortGroupSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemPicsItemPortGroupsItem) basetypes.ListValue {

	var data_list = []PortGroupsValue{}
	for _, d := range l {
		var count basetypes.Int64Value
		var type_pg basetypes.StringValue

		if d.Count != nil {
			count = types.Int64Value(int64(*d.Count))
		}
		if d.Type != nil {
			type_pg = types.StringValue(*d.Type)
		}

		data_map_attr_type := PortGroupsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"count": count,
			"type":  type_pg,
		}
		data, e := NewPortGroupsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, PortGroupsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatPicSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemPicsItem) basetypes.ListValue {

	var data_list = []PicsValue{}
	for _, d := range l {
		var index basetypes.Int64Value
		var model_number basetypes.StringValue
		var port_groups basetypes.ListValue = types.ListNull(PortGroupsValue{}.Type(ctx))

		if d.Index != nil {
			index = types.Int64Value(int64(*d.Index))
		}
		if d.ModelNumber != nil {
			model_number = types.StringValue(*d.ModelNumber)
		}
		if d.PortGroups != nil {
			port_groups = moduleStatPicPortGroupSdkToTerraform(ctx, diags, d.PortGroups)
		}

		data_map_attr_type := PicsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"index":        index,
			"model_number": model_number,
			"port_groups":  port_groups,
		}
		data, e := NewPicsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, PicsValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}

func moduleStatPoeSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ModuleStatItemPoe) basetypes.ObjectValue {

	var max_power basetypes.NumberValue
	var power_draw basetypes.NumberValue

	if d.MaxPower != nil {
		max_power = types.NumberValue(big.NewFloat(*d.MaxPower))
	}
	if d.PowerDraw != nil {
		power_draw = types.NumberValue(big.NewFloat(*d.PowerDraw))
	}

	data_map_attr_type := PoeValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"max_power":  max_power,
		"power_draw": power_draw,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
func moduleStatPsusSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemPsusItem) basetypes.ListValue {

	var data_list = []PsusValue{}
	for _, d := range l {
		var name basetypes.StringValue
		var status basetypes.StringValue

		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}
		if d.Status != nil {
			status = types.StringValue(*d.Status)
		}

		data_map_attr_type := PsusValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"name":   name,
			"status": status,
		}
		data, e := NewPsusValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, PsusValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatTemperatureSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemTemperaturesItem) basetypes.ListValue {

	var data_list = []TemperaturesValue{}
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

		data_map_attr_type := TemperaturesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"celsius": celsius,
			"name":    name,
			"status":  status,
		}
		data, e := NewTemperaturesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, TemperaturesValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatVcLinksSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItemVcLinksItem) basetypes.ListValue {

	var data_list = []VcLinksValue{}
	for _, d := range l {
		var neighbor_module_idx basetypes.Int64Value
		var neighbor_port_id basetypes.StringValue
		var port_id basetypes.StringValue

		if d.NeighborModuleIdx != nil {
			neighbor_module_idx = types.Int64Value(int64(*d.NeighborModuleIdx))
		}
		if d.NeighborPortId != nil {
			neighbor_port_id = types.StringValue(*d.NeighborPortId)
		}
		if d.PortId != nil {
			port_id = types.StringValue(*d.PortId)
		}

		data_map_attr_type := VcLinksValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"neighbor_module_idx": neighbor_module_idx,
			"neighbor_port_id":    neighbor_port_id,
			"port_id":             port_id,
		}
		data, e := NewVcLinksValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, VcLinksValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
func moduleStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.ModuleStatItem) basetypes.ListValue {

	var data_list = []ModuleStatValue{}
	for _, d := range l {
		var backup_version basetypes.StringValue
		var bios_version basetypes.StringValue
		var cpld_version basetypes.StringValue
		var errors basetypes.ListValue = types.ListNull(ErrorsValue{}.Type(ctx))
		var fans basetypes.ListValue = types.ListNull(FansValue{}.Type(ctx))
		var fpga_version basetypes.StringValue
		var last_seen basetypes.NumberValue
		var model basetypes.StringValue
		var optics_cpld_version basetypes.StringValue
		var pending_version basetypes.StringValue
		var pics basetypes.ListValue = types.ListNull(PicsValue{}.Type(ctx))
		var poe basetypes.ObjectValue = types.ObjectNull(PoeValue{}.AttributeTypes(ctx))
		var poe_version basetypes.StringValue
		var power_cpld_version basetypes.StringValue
		var psus basetypes.ListValue = types.ListNull(PsusValue{}.Type(ctx))
		var re_fpga_version basetypes.StringValue
		var recovery_version basetypes.StringValue
		var serial basetypes.StringValue
		var status basetypes.StringValue
		var temperatures basetypes.ListValue = types.ListNull(TemperaturesValue{}.Type(ctx))
		var tmc_fpga_version basetypes.StringValue
		var uboot_version basetypes.StringValue
		var uptime basetypes.Int64Value
		var vc_links basetypes.ListValue = types.ListNull(VcLinksValue{}.Type(ctx))
		var vc_mode basetypes.StringValue
		var vc_role basetypes.StringValue
		var vc_state basetypes.StringValue
		var version basetypes.StringValue

		if d.BackupVersion.Value() != nil {
			backup_version = types.StringValue(*d.BackupVersion.Value())
		}
		if d.BiosVersion.Value() != nil {
			bios_version = types.StringValue(*d.BiosVersion.Value())
		}
		if d.CpldVersion.Value() != nil {
			cpld_version = types.StringValue(*d.CpldVersion.Value())
		}
		if d.Errors != nil {
			errors = moduleStatErrorSdkToTerraform(ctx, diags, d.Errors)
		}
		if d.Fans != nil {
			fans = moduleStatFanSdkToTerraform(ctx, diags, d.Fans)
		}
		if d.FpgaVersion.Value() != nil {
			fpga_version = types.StringValue(*d.FpgaVersion.Value())
		}
		if d.LastSeen.Value() != nil {
			last_seen = types.NumberValue(big.NewFloat(*d.LastSeen.Value()))
		}
		if d.Model.Value() != nil {
			model = types.StringValue(*d.Model.Value())
		}
		if d.OpticsCpldVersion.Value() != nil {
			optics_cpld_version = types.StringValue(*d.OpticsCpldVersion.Value())
		}
		if d.PendingVersion.Value() != nil {
			pending_version = types.StringValue(*d.PendingVersion.Value())
		}
		if d.Pics != nil {
			pics = moduleStatPicSdkToTerraform(ctx, diags, d.Pics)
		}
		if d.Poe != nil {
			poe = moduleStatPoeSdkToTerraform(ctx, diags, d.Poe)
		}
		if d.PoeVersion.Value() != nil {
			poe_version = types.StringValue(*d.PoeVersion.Value())
		}
		if d.PowerCpldVersion.Value() != nil {
			power_cpld_version = types.StringValue(*d.PowerCpldVersion.Value())
		}
		if d.Psus != nil {
			psus = moduleStatPsusSdkToTerraform(ctx, diags, d.Psus)
		}
		if d.ReFpgaVersion.Value() != nil {
			re_fpga_version = types.StringValue(*d.ReFpgaVersion.Value())
		}
		if d.RecoveryVersion.Value() != nil {
			recovery_version = types.StringValue(*d.RecoveryVersion.Value())
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
			tmc_fpga_version = types.StringValue(*d.TmcFpgaVersion.Value())
		}
		if d.UbootVersion.Value() != nil {
			uboot_version = types.StringValue(*d.UbootVersion.Value())
		}
		if d.Uptime.Value() != nil {
			uptime = types.Int64Value(int64(*d.Uptime.Value()))
		}
		if d.VcLinks != nil {
			vc_links = moduleStatVcLinksSdkToTerraform(ctx, diags, d.VcLinks)
		}
		if d.VcMode.Value() != nil {
			vc_mode = types.StringValue(*d.VcMode.Value())
		}
		if d.VcRole.Value() != nil {
			vc_role = types.StringValue(*d.VcRole.Value())
		}
		if d.VcState.Value() != nil {
			vc_state = types.StringValue(*d.VcState.Value())
		}
		if d.Version.Value() != nil {
			version = types.StringValue(*d.Version.Value())
		}

		data_map_attr_type := ModuleStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"backup_version":      backup_version,
			"bios_version":        bios_version,
			"cpld_version":        cpld_version,
			"errors":              errors,
			"fans":                fans,
			"fpga_version":        fpga_version,
			"last_seen":           last_seen,
			"model":               model,
			"optics_cpld_version": optics_cpld_version,
			"pending_version":     pending_version,
			"pics":                pics,
			"poe":                 poe,
			"poe_version":         poe_version,
			"power_cpld_version":  power_cpld_version,
			"psus":                psus,
			"re_fpga_version":     re_fpga_version,
			"recovery_version":    recovery_version,
			"serial":              serial,
			"status":              status,
			"temperatures":        temperatures,
			"tmc_fpga_version":    tmc_fpga_version,
			"uboot_version":       uboot_version,
			"uptime":              uptime,
			"vc_links":            vc_links,
			"vc_mode":             vc_mode,
			"vc_role":             vc_role,
			"vc_state":            vc_state,
			"version":             version,
		}
		data, e := NewModuleStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ModuleStatValue{}.Type(ctx), data_list)
	diags.Append(e...)

	return r
}
