package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func envStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApEnvStat) basetypes.ObjectValue {
	var accel_x basetypes.Float64Value
	var accel_y basetypes.Float64Value
	var accel_z basetypes.Float64Value
	var ambient_temp basetypes.Int64Value
	var attitude basetypes.Int64Value
	var cpu_temp basetypes.Int64Value
	var humidity basetypes.Int64Value
	var magne_x basetypes.Float64Value
	var magne_y basetypes.Float64Value
	var magne_z basetypes.Float64Value
	var pressure basetypes.Float64Value
	var vcore_voltage basetypes.Int64Value

	if d.AccelX.Value() != nil {
		accel_x = types.Float64Value(*d.AccelX.Value())
	}
	if d.AccelY.Value() != nil {
		accel_y = types.Float64Value(*d.AccelY.Value())
	}
	if d.AccelZ.Value() != nil {
		accel_z = types.Float64Value(*d.AccelZ.Value())
	}
	if d.AmbientTemp.Value() != nil {
		ambient_temp = types.Int64Value(int64(*d.AmbientTemp.Value()))
	}
	if d.Attitude.Value() != nil {
		attitude = types.Int64Value(int64(*d.Attitude.Value()))
	}
	if d.CpuTemp.Value() != nil {
		cpu_temp = types.Int64Value(int64(*d.CpuTemp.Value()))
	}
	if d.Humidity.Value() != nil {
		humidity = types.Int64Value(int64(*d.Humidity.Value()))
	}
	if d.MagneX.Value() != nil {
		magne_x = types.Float64Value(*d.MagneX.Value())
	}
	if d.MagneY.Value() != nil {
		magne_y = types.Float64Value(*d.MagneY.Value())
	}
	if d.MagneZ.Value() != nil {
		magne_z = types.Float64Value(*d.MagneZ.Value())
	}
	if d.Pressure.Value() != nil {
		pressure = types.Float64Value(*d.Pressure.Value())
	}
	if d.VcoreVoltage.Value() != nil {
		vcore_voltage = types.Int64Value(int64(*d.VcoreVoltage.Value()))
	}

	data_map_attr_type := EnvStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"accel_x":       accel_x,
		"accel_y":       accel_y,
		"accel_z":       accel_z,
		"ambient_temp":  ambient_temp,
		"attitude":      attitude,
		"cpu_temp":      cpu_temp,
		"humidity":      humidity,
		"magne_x":       magne_x,
		"magne_y":       magne_y,
		"magne_z":       magne_z,
		"pressure":      pressure,
		"vcore_voltage": vcore_voltage,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
