package resource_org_deviceprofile_ap

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand24) basetypes.ObjectValue {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channel basetypes.Int64Value
	var channels = mistutils.ListOfIntSdkToTerraformEmpty()
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var powerMax basetypes.Int64Value
	var powerMin basetypes.Int64Value
	var preamble basetypes.StringValue

	if d.AllowRrmDisable != nil {
		allowRrmDisable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		antGain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antennaMode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Channels.Value() != nil {
		channels = mistutils.ListOfIntSdkToTerraform(*d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		powerMax = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		powerMin = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}

	dataMapValue := map[string]attr.Value{
		"allow_rrm_disable": allowRrmDisable,
		"ant_gain":          antGain,
		"antenna_mode":      antennaMode,
		"bandwidth":         bandwidth,
		"channel":           channel,
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         powerMax,
		"power_min":         powerMin,
		"preamble":          preamble,
	}
	data, e := basetypes.NewObjectValue(Band24Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand5) basetypes.ObjectValue {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaBeamPattern basetypes.StringValue
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channel basetypes.Int64Value
	var channels = mistutils.ListOfIntSdkToTerraformEmpty()
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var powerMax basetypes.Int64Value
	var powerMin basetypes.Int64Value
	var preamble basetypes.StringValue

	if d.AllowRrmDisable != nil {
		allowRrmDisable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		antGain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antennaMode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Channels.Value() != nil {
		channels = mistutils.ListOfIntSdkToTerraform(*d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		powerMax = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		powerMin = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}

	dataMapValue := map[string]attr.Value{
		"allow_rrm_disable":    allowRrmDisable,
		"ant_gain":             antGain,
		"antenna_beam_pattern": antennaBeamPattern,
		"antenna_mode":         antennaMode,
		"bandwidth":            bandwidth,
		"channel":              channel,
		"channels":             channels,
		"disabled":             disabled,
		"power":                power,
		"power_max":            powerMax,
		"power_min":            powerMin,
		"preamble":             preamble,
	}
	data, e := basetypes.NewObjectValue(Band5Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadioBand6) basetypes.ObjectValue {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaBeamPattern basetypes.StringValue
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channel basetypes.Int64Value
	var channels = mistutils.ListOfIntSdkToTerraformEmpty()
	var disabled basetypes.BoolValue
	var power basetypes.Int64Value
	var powerMax basetypes.Int64Value
	var powerMin basetypes.Int64Value
	var preamble basetypes.StringValue
	var standardPower basetypes.BoolValue

	if d.AllowRrmDisable != nil {
		allowRrmDisable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain.Value() != nil {
		antGain = types.Int64Value(int64(*d.AntGain.Value()))
	}
	if d.AntennaMode != nil {
		antennaMode = types.StringValue(string(*d.AntennaMode))
	}
	if d.Bandwidth != nil {
		bandwidth = types.Int64Value(int64(*d.Bandwidth))
	}
	if d.Channel.Value() != nil {
		channel = types.Int64Value(int64(*d.Channel.Value()))
	}
	if d.Channels.Value() != nil {
		channels = mistutils.ListOfIntSdkToTerraform(*d.Channels.Value())
	}
	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Power.Value() != nil {
		power = types.Int64Value(int64(*d.Power.Value()))
	}
	if d.PowerMax.Value() != nil {
		powerMax = types.Int64Value(int64(*d.PowerMax.Value()))
	}
	if d.PowerMin.Value() != nil {
		powerMin = types.Int64Value(int64(*d.PowerMin.Value()))
	}
	if d.Preamble != nil {
		preamble = types.StringValue(string(*d.Preamble))
	}
	if d.StandardPower != nil {
		standardPower = types.BoolValue(*d.StandardPower)
	}

	dataMapValue := map[string]attr.Value{
		"allow_rrm_disable":    allowRrmDisable,
		"ant_gain":             antGain,
		"antenna_beam_pattern": antennaBeamPattern,
		"antenna_mode":         antennaMode,
		"bandwidth":            bandwidth,
		"channel":              channel,
		"channels":             channels,
		"disabled":             disabled,
		"power":                power,
		"power_max":            powerMax,
		"power_min":            powerMin,
		"preamble":             preamble,
		"standard_power":       standardPower,
	}
	data, e := basetypes.NewObjectValue(Band6Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func radioConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApRadio) RadioConfigValue {

	var allowRrmDisable basetypes.BoolValue
	var antGain24 basetypes.Int64Value
	var antGain5 basetypes.Int64Value
	var antGain6 basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var antennaSelect basetypes.StringValue
	var band24 = types.ObjectNull(Band24Value{}.AttributeTypes(ctx))
	var band24Usage basetypes.StringValue
	var band5 = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
	var band5On24Radio = types.ObjectNull(Band5Value{}.AttributeTypes(ctx))
	var band6 = types.ObjectNull(Band6Value{}.AttributeTypes(ctx))
	var fullAutomaticRrm basetypes.BoolValue
	var indoorUse basetypes.BoolValue
	var rrmManaged basetypes.BoolValue
	var scanningEnabled basetypes.BoolValue

	if d.AllowRrmDisable != nil {
		allowRrmDisable = types.BoolValue(*d.AllowRrmDisable)
	}
	if d.AntGain24 != nil {
		antGain24 = types.Int64Value(int64(*d.AntGain24))
	}
	if d.AntGain5 != nil {
		antGain5 = types.Int64Value(int64(*d.AntGain5))
	}
	if d.AntGain6 != nil {
		antGain6 = types.Int64Value(int64(*d.AntGain6))
	}
	if d.AntennaMode != nil {
		antennaMode = types.StringValue(string(*d.AntennaMode))
	}
	if d.AntennaSelect != nil {
		antennaSelect = types.StringValue(string(*d.AntennaSelect))
	}
	if d.Band24 != nil {
		band24 = band24SdkToTerraform(ctx, diags, d.Band24)
	}
	if d.Band24Usage != nil {
		band24Usage = types.StringValue(string(*d.Band24Usage))
	}
	if d.Band5 != nil {
		band5 = band5SdkToTerraform(ctx, diags, d.Band5)
	}
	if d.Band5On24Radio != nil {
		band5On24Radio = band5SdkToTerraform(ctx, diags, d.Band5On24Radio)
	}
	if d.Band6 != nil {
		band6 = band6SdkToTerraform(ctx, diags, d.Band6)
	}
	if d.FullAutomaticRrm != nil {
		fullAutomaticRrm = types.BoolValue(*d.FullAutomaticRrm)
	}
	if d.IndoorUse != nil {
		indoorUse = types.BoolValue(*d.IndoorUse)
	}
	if d.RrmManaged != nil {
		rrmManaged = types.BoolValue(*d.RrmManaged)
	}
	if d.ScanningEnabled != nil {
		scanningEnabled = types.BoolValue(*d.ScanningEnabled)
	}

	dataMapValue := map[string]attr.Value{
		"allow_rrm_disable":  allowRrmDisable,
		"ant_gain_24":        antGain24,
		"ant_gain_5":         antGain5,
		"ant_gain_6":         antGain6,
		"antenna_mode":       antennaMode,
		"antenna_select":     antennaSelect,
		"band_24":            band24,
		"band_24_usage":      band24Usage,
		"band_5":             band5,
		"band_5_on_24_radio": band5On24Radio,
		"band_6":             band6,
		"full_automatic_rrm": fullAutomaticRrm,
		"indoor_use":         indoorUse,
		"rrm_managed":        rrmManaged,
		"scanning_enabled":   scanningEnabled,
	}
	data, e := NewRadioConfigValue(RadioConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
