package resource_org_rftemplate

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func band24SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand24) Band24Value {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels = misttransform.ListOfIntSdkToTerraformEmpty()
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
	if d.Channels.Value() != nil {
		channels = misttransform.ListOfIntSdkToTerraform(*d.Channels.Value())
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
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         powerMax,
		"power_min":         powerMin,
		"preamble":          preamble,
	}
	data, e := NewBand24Value(Band24Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func band5SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand5) Band5Value {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels = misttransform.ListOfIntSdkToTerraformEmpty()
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
	if d.Channels.Value() != nil {
		channels = misttransform.ListOfIntSdkToTerraform(*d.Channels.Value())
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
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         powerMax,
		"power_min":         powerMin,
		"preamble":          preamble,
	}
	data, e := NewBand5Value(Band5Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func band5On24RadioSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand5) Band5On24RadioValue {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels = misttransform.ListOfIntSdkToTerraformEmpty()
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
	if d.Channels.Value() != nil {
		channels = misttransform.ListOfIntSdkToTerraform(*d.Channels.Value())
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
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         powerMax,
		"power_min":         powerMin,
		"preamble":          preamble,
	}
	data, e := NewBand5On24RadioValue(Band5On24RadioValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func band6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RftemplateRadioBand6) Band6Value {
	var allowRrmDisable basetypes.BoolValue
	var antGain basetypes.Int64Value
	var antennaMode basetypes.StringValue
	var bandwidth basetypes.Int64Value
	var channels = misttransform.ListOfIntSdkToTerraformEmpty()
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
	if d.Channels.Value() != nil {
		channels = misttransform.ListOfIntSdkToTerraform(*d.Channels.Value())
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
		"allow_rrm_disable": allowRrmDisable,
		"ant_gain":          antGain,
		"antenna_mode":      antennaMode,
		"bandwidth":         bandwidth,
		"channels":          channels,
		"disabled":          disabled,
		"power":             power,
		"power_max":         powerMax,
		"power_min":         powerMin,
		"preamble":          preamble,
		"standard_power":    standardPower,
	}
	data, e := NewBand6Value(Band6Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
