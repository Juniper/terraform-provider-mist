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
	var accelX basetypes.Float64Value
	var accelY basetypes.Float64Value
	var accelZ basetypes.Float64Value
	var ambientTemp basetypes.Int64Value
	var attitude basetypes.Int64Value
	var cpuTemp basetypes.Int64Value
	var humidity basetypes.Int64Value
	var magneX basetypes.Float64Value
	var magneY basetypes.Float64Value
	var magneZ basetypes.Float64Value
	var pressure basetypes.Float64Value
	var vcoreVoltage basetypes.Int64Value

	if d.AccelX.Value() != nil {
		accelX = types.Float64Value(*d.AccelX.Value())
	}
	if d.AccelY.Value() != nil {
		accelY = types.Float64Value(*d.AccelY.Value())
	}
	if d.AccelZ.Value() != nil {
		accelZ = types.Float64Value(*d.AccelZ.Value())
	}
	if d.AmbientTemp.Value() != nil {
		ambientTemp = types.Int64Value(int64(*d.AmbientTemp.Value()))
	}
	if d.Attitude.Value() != nil {
		attitude = types.Int64Value(int64(*d.Attitude.Value()))
	}
	if d.CpuTemp.Value() != nil {
		cpuTemp = types.Int64Value(int64(*d.CpuTemp.Value()))
	}
	if d.Humidity.Value() != nil {
		humidity = types.Int64Value(int64(*d.Humidity.Value()))
	}
	if d.MagneX.Value() != nil {
		magneX = types.Float64Value(*d.MagneX.Value())
	}
	if d.MagneY.Value() != nil {
		magneY = types.Float64Value(*d.MagneY.Value())
	}
	if d.MagneZ.Value() != nil {
		magneZ = types.Float64Value(*d.MagneZ.Value())
	}
	if d.Pressure.Value() != nil {
		pressure = types.Float64Value(*d.Pressure.Value())
	}
	if d.VcoreVoltage.Value() != nil {
		vcoreVoltage = types.Int64Value(int64(*d.VcoreVoltage.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"accel_x":       accelX,
		"accel_y":       accelY,
		"accel_z":       accelZ,
		"ambient_temp":  ambientTemp,
		"attitude":      attitude,
		"cpu_temp":      cpuTemp,
		"humidity":      humidity,
		"magne_x":       magneX,
		"magne_y":       magneY,
		"magne_z":       magneZ,
		"pressure":      pressure,
		"vcore_voltage": vcoreVoltage,
	}
	data, e := types.ObjectValue(EnvStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
