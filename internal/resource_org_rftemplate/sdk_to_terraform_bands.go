package resource_org_rftemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand24) Band24Value {
	var allow_rrm_disable basetypes.BoolValue
	var ant_gain basetypes.Int64Value
	var antenna_mode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var power_max basetypes.Int64Value
	var power_min basetypes.Int64Value
	var preamble basetypes.StringValue

	if d.AllowRrmDisable != nil {
		allow_rrm_disable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		ant_gain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antenna_mode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channels.Value() != nil {
		channels = mist_transform.ListOfIntSdkToTerraform(ctx, *d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		power_max = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		power_min = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}

	data_map_attr_type := Band24Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": allow_rrm_disable,
		"ant_gain":          ant_gain,
		"antenna_mode":      antenna_mode,
		"bandwidth":         bandwidth,
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         power_max,
		"power_min":         power_min,
		"preamble":          preamble,
	}
	data, e := NewBand24Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand5) Band5Value {
	var allow_rrm_disable basetypes.BoolValue
	var ant_gain basetypes.Int64Value
	var antenna_mode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var power_max basetypes.Int64Value
	var power_min basetypes.Int64Value
	var preamble basetypes.StringValue

	if d.AllowRrmDisable != nil {
		allow_rrm_disable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		ant_gain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antenna_mode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channels.Value() != nil {
		channels = mist_transform.ListOfIntSdkToTerraform(ctx, *d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		power_max = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		power_min = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}

	data_map_attr_type := Band5Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": allow_rrm_disable,
		"ant_gain":          ant_gain,
		"antenna_mode":      antenna_mode,
		"bandwidth":         bandwidth,
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         power_max,
		"power_min":         power_min,
		"preamble":          preamble,
	}
	data, e := NewBand5Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand6) Band6Value {
	var allow_rrm_disable basetypes.BoolValue
	var ant_gain basetypes.Int64Value
	var antenna_mode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var power_max basetypes.Int64Value
	var power_min basetypes.Int64Value
	var preamble basetypes.StringValue
	var standard_power basetypes.BoolValue

	if d.AllowRrmDisable != nil {
		allow_rrm_disable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		ant_gain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antenna_mode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channels.Value() != nil {
		channels = mist_transform.ListOfIntSdkToTerraform(ctx, *d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		power_max = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		power_min = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}
	if d.StandardPower != nil {
		standard_power = types.BoolValue(*d.StandardPower)
	}

	data_map_attr_type := Band6Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable": allow_rrm_disable,
		"ant_gain":          ant_gain,
		"antenna_mode":      antenna_mode,
		"bandwidth":         bandwidth,
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         power_max,
		"power_min":         power_min,
		"preamble":          preamble,
		"standard_power":    standard_power,
	}
	data, e := NewBand6Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
