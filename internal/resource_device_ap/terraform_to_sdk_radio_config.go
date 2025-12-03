package resource_device_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func channelsTerraformToSdk(d basetypes.ListValue) []int {
	var data []int
	for _, v := range d.Elements() {
		var i interface{} = v
		d := i.(basetypes.Int64Value)
		data = append(data, int(d.ValueInt64()))
	}
	return data
}

func band24TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApRadioBand24 {
	data := models.ApRadioBand24{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewBand24Value(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.AllowRrmDisable.ValueBoolPointer() != nil {
				data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
			}
			if plan.AntGain.ValueInt64Pointer() != nil {
				data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
			}
			if plan.AntennaMode.ValueStringPointer() != nil {
				data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
			}
			if plan.Bandwidth.ValueInt64Pointer() != nil {
				data.Bandwidth = models.ToPointer(models.Dot11Bandwidth24Enum(plan.Bandwidth.ValueInt64()))
			}
			if plan.Channel.ValueInt64Pointer() != nil {
				data.Channel = models.NewOptional(models.ToPointer(int(plan.Channel.ValueInt64())))
			}
			if !plan.Channels.IsNull() && !plan.Channels.IsUnknown() {
				data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(plan.Channels)))
			}
			if plan.Disabled.ValueBoolPointer() != nil {
				data.Disabled = plan.Disabled.ValueBoolPointer()
			}
			if plan.Power.ValueInt64Pointer() != nil {
				data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
			}
			if plan.PowerMax.ValueInt64Pointer() != nil {
				data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
			}
			if plan.PowerMin.ValueInt64Pointer() != nil {
				data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
			}
			if plan.Preamble.ValueStringPointer() != nil {
				data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))
			}
		}
	}
	return &data
}

func band5TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApRadioBand5 {
	data := models.ApRadioBand5{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewBand5Value(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.AllowRrmDisable.ValueBoolPointer() != nil {
				data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
			}
			if plan.AntGain.ValueInt64Pointer() != nil {
				data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
			}
			if plan.AntennaBeamPattern.ValueStringPointer() != nil {
				data.AntennaBeamPattern = (*models.RadioBandAntennaBeamPatternEnum)(plan.AntennaBeamPattern.ValueStringPointer())
			}
			if plan.AntennaMode.ValueStringPointer() != nil {
				data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
			}
			if plan.Bandwidth.ValueInt64Pointer() != nil {
				data.Bandwidth = models.ToPointer(models.Dot11Bandwidth5Enum(plan.Bandwidth.ValueInt64()))
			}
			if plan.Channel.ValueInt64Pointer() != nil {
				data.Channel = models.NewOptional(models.ToPointer(int(plan.Channel.ValueInt64())))
			}
			if !plan.Channels.IsNull() && !plan.Channels.IsUnknown() {
				data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(plan.Channels)))
			}
			if plan.Disabled.ValueBoolPointer() != nil {
				data.Disabled = plan.Disabled.ValueBoolPointer()
			}
			if plan.Power.ValueInt64Pointer() != nil {
				data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
			}
			if plan.PowerMax.ValueInt64Pointer() != nil {
				data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
			}
			if plan.PowerMin.ValueInt64Pointer() != nil {
				data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
			}
			if plan.Preamble.ValueStringPointer() != nil {
				data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))
			}
		}
	}
	return &data
}

func band6TerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.ApRadioBand6 {
	data := models.ApRadioBand6{}
	if !d.IsNull() && !d.IsUnknown() {
		plan, e := NewBand6Value(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if plan.AllowRrmDisable.ValueBoolPointer() != nil {
				data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
			}
			if plan.AntGain.ValueInt64Pointer() != nil {
				data.AntGain = models.NewOptional(models.ToPointer(int(plan.AntGain.ValueInt64())))
			}
			if plan.AntennaBeamPattern.ValueStringPointer() != nil {
				data.AntennaBeamPattern = (*models.RadioBandAntennaBeamPatternEnum)(plan.AntennaBeamPattern.ValueStringPointer())
			}
			if plan.AntennaMode.ValueStringPointer() != nil {
				data.AntennaMode = models.ToPointer(models.RadioBandAntennaModeEnum(plan.AntennaMode.ValueString()))
			}
			if plan.Bandwidth.ValueInt64Pointer() != nil {
				data.Bandwidth = models.ToPointer(models.Dot11Bandwidth6Enum(plan.Bandwidth.ValueInt64()))
			}
			if plan.Channel.ValueInt64Pointer() != nil {
				data.Channel = models.NewOptional(models.ToPointer(int(plan.Channel.ValueInt64())))
			}
			if !plan.Channels.IsNull() && !plan.Channels.IsUnknown() {
				data.Channels = models.NewOptional(models.ToPointer(channelsTerraformToSdk(plan.Channels)))
			}
			if plan.Disabled.ValueBoolPointer() != nil {
				data.Disabled = plan.Disabled.ValueBoolPointer()
			}
			if plan.Power.ValueInt64Pointer() != nil {
				data.Power = models.NewOptional(models.ToPointer(int(plan.Power.ValueInt64())))
			}
			if plan.PowerMax.ValueInt64Pointer() != nil {
				data.PowerMax = models.NewOptional(models.ToPointer(int(plan.PowerMax.ValueInt64())))
			}
			if plan.PowerMin.ValueInt64Pointer() != nil {
				data.PowerMin = models.NewOptional(models.ToPointer(int(plan.PowerMin.ValueInt64())))
			}
			if plan.Preamble.ValueStringPointer() != nil {
				data.Preamble = models.ToPointer(models.RadioBandPreambleEnum(plan.Preamble.ValueString()))
			}
			if plan.StandardPower.ValueBoolPointer() != nil {
				data.StandardPower = plan.StandardPower.ValueBoolPointer()
			}
		}
	}
	return &data
}

func radioConfigTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan RadioConfigValue) *models.ApRadio {

	data := models.ApRadio{}

	if plan.AllowRrmDisable.ValueBoolPointer() != nil {
		data.AllowRrmDisable = plan.AllowRrmDisable.ValueBoolPointer()
	}
	if plan.AntGain24.ValueInt64Pointer() != nil {
		data.AntGain24 = models.ToPointer(int(plan.AntGain24.ValueInt64()))
	}
	if plan.AntGain5.ValueInt64Pointer() != nil {
		data.AntGain5 = models.ToPointer(int(plan.AntGain5.ValueInt64()))
	}
	if plan.AntGain6.ValueInt64Pointer() != nil {
		data.AntGain6 = models.ToPointer(int(plan.AntGain6.ValueInt64()))
	}
	if !plan.AntennaMode.IsNull() && !plan.AntennaMode.IsUnknown() {
		data.AntennaMode = models.ToPointer(models.ApRadioAntennaModeEnum(plan.AntennaMode.ValueString()))
	}
	if !plan.AntennaSelect.IsNull() && !plan.AntennaSelect.IsUnknown() {
		data.AntennaSelect = models.ToPointer(models.AntennaSelectEnum(plan.AntennaSelect.ValueString()))
	}
	if !plan.Band24.IsNull() && !plan.Band24.IsUnknown() {
		data.Band24 = band24TerraformToSdk(ctx, diags, plan.Band24)
	}
	data.Band24Usage = models.ToPointer(models.RadioBand24UsageEnum(plan.Band24Usage.ValueString()))

	if !plan.Band5On24Radio.IsNull() && !plan.Band5On24Radio.IsUnknown() {
		data.Band5On24Radio = band5TerraformToSdk(ctx, diags, plan.Band5On24Radio)
	}
	if !plan.Band5.IsNull() && !plan.Band5.IsUnknown() {
		data.Band5 = band5TerraformToSdk(ctx, diags, plan.Band5)
	}
	if !plan.Band6.IsNull() && !plan.Band6.IsUnknown() {
		data.Band6 = band6TerraformToSdk(ctx, diags, plan.Band6)
	}
	if plan.FullAutomaticRrm.ValueBoolPointer() != nil {
		data.FullAutomaticRrm = plan.FullAutomaticRrm.ValueBoolPointer()
	}
	if plan.IndoorUse.ValueBoolPointer() != nil {
		data.IndoorUse = plan.IndoorUse.ValueBoolPointer()
	}
	if plan.RrmManaged.ValueBoolPointer() != nil {
		data.RrmManaged = plan.RrmManaged.ValueBoolPointer()
	}
	if plan.ScanningEnabled.ValueBoolPointer() != nil {
		data.ScanningEnabled = plan.ScanningEnabled.ValueBoolPointer()
	}

	return &data
}
