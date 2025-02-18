package resource_org_rftemplate

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

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

func band24TerraformToSdk(plan Band24Value) *models.RftemplateRadioBand24 {

	data := models.RftemplateRadioBand24{}

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

	return &data
}

func band5TerraformToSdk(plan Band5Value) *models.RftemplateRadioBand5 {

	data := models.RftemplateRadioBand5{}

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
		data.Bandwidth = models.ToPointer(models.Dot11Bandwidth5Enum(plan.Bandwidth.ValueInt64()))
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

	return &data
}

func band5On24RadioTerraformToSdk(plan Band5On24RadioValue) *models.RftemplateRadioBand5 {

	data := models.RftemplateRadioBand5{}

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
		data.Bandwidth = models.ToPointer(models.Dot11Bandwidth5Enum(plan.Bandwidth.ValueInt64()))
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

	return &data
}

func band6TerraformToSdk(plan Band6Value) *models.RftemplateRadioBand6 {

	data := models.RftemplateRadioBand6{}

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
		data.Bandwidth = models.ToPointer(models.Dot11Bandwidth6Enum(plan.Bandwidth.ValueInt64()))
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

	return &data
}
