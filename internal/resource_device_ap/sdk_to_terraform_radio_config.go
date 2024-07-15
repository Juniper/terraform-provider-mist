package resource_device_ap

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand24) basetypes.ObjectValue {
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
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand5) basetypes.ObjectValue {
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
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand6) basetypes.ObjectValue {
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
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func radioConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadio) RadioConfigValue {
	tflog.Debug(ctx, "radioConfigSdkToTerraform")

	var allow_rrm_disable basetypes.BoolValue
	var ant_gain_24 basetypes.Int64Value
	var ant_gain_5 basetypes.Int64Value
	var ant_gain_6 basetypes.Int64Value
	var antenna_mode basetypes.StringValue
	var band_24 basetypes.ObjectValue = types.ObjectNull(Band24Value{}.AttributeTypes(ctx))
	var band_24_usage basetypes.StringValue
	var band_5 basetypes.ObjectValue = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
	var band_5_on_24_radio basetypes.ObjectValue = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
	var band_6 basetypes.ObjectValue = types.ObjectNull(Band6Value{}.AttributeTypes(ctx))
	var indoor_use basetypes.BoolValue
	var scanning_enabled basetypes.BoolValue

	if d.AllowRrmDisable != nil {
		allow_rrm_disable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain24 != nil {
		ant_gain_24 = types.Int64Value(int64(*d.AntGain24))
	}
	if d.AntGain5 != nil {
		ant_gain_5 = types.Int64Value(int64(*d.AntGain5))
	}
	if d.AntGain6 != nil {
		ant_gain_6 = types.Int64Value(int64(*d.AntGain6))
	}
	if d.AntennaMode != nil {
		antenna_mode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Band24 != nil {
		band_24 = band24SdkToTerraform(ctx, diags, d.Band24)
	}
	if d.Band24Usage != nil {
		band_24_usage = types.StringValue(string(*d.Band24Usage))
	}
	if d.Band5 != nil {
		band_5 = band5SdkToTerraform(ctx, diags, d.Band5)
	}
	if d.Band5On24Radio != nil {
		band_5_on_24_radio = band5SdkToTerraform(ctx, diags, d.Band5On24Radio)
	}
	if d.Band6 != nil {
		band_6 = band6SdkToTerraform(ctx, diags, d.Band6)
	}
	if d.IndoorUse != nil {
		indoor_use = types.BoolValue(*d.IndoorUse)
	}
	if d.ScanningEnabled != nil {
		scanning_enabled = types.BoolValue(*d.ScanningEnabled)
	}

	data_map_attr_type := RadioConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allow_rrm_disable":  allow_rrm_disable,
		"ant_gain_24":        ant_gain_24,
		"ant_gain_5":         ant_gain_5,
		"ant_gain_6":         ant_gain_6,
		"antenna_mode":       antenna_mode,
		"band_24":            band_24,
		"band_24_usage":      band_24_usage,
		"band_5":             band_5,
		"band_5_on_24_radio": band_5_on_24_radio,
		"band_6":             band_6,
		"indoor_use":         indoor_use,
		"scanning_enabled":   scanning_enabled,
	}
	data, e := NewRadioConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
