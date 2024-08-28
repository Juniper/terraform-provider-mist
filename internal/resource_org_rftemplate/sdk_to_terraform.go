package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data models.RfTemplate) (OrgRftemplateModel, diag.Diagnostics) {
	var state OrgRftemplateModel
	var diags diag.Diagnostics

	var ant_gain_24 types.Int64
	var ant_gain_5 types.Int64
	var ant_gain_6 types.Int64
	var band_24 Band24Value = NewBand24ValueNull()
	var band_24_usage types.String
	var band_5 Band5Value = NewBand5ValueNull()
	var band_5_on_24_radio Band5On24RadioValue
	var band_6 Band6Value = NewBand6ValueNull()
	var country_code types.String
	var id types.String
	var model_specific types.Map = types.MapNull(ModelSpecificValue{}.Type(ctx))
	var name types.String
	var org_id types.String
	var scanning_enabled types.Bool

	if data.AntGain24 != nil {
		ant_gain_24 = types.Int64Value(int64(*data.AntGain24))
	}
	if data.AntGain5 != nil {
		ant_gain_5 = types.Int64Value(int64(*data.AntGain5))
	}
	if data.AntGain6 != nil {
		ant_gain_6 = types.Int64Value(int64(*data.AntGain6))
	}
	if data.Band24 != nil {
		band_24 = band24SdkToTerraform(ctx, &diags, data.Band24)
	}
	if data.Band24Usage != nil {
		band_24_usage = types.StringValue(string(*data.Band24Usage))
	}
	if data.Band5 != nil {
		band_5 = band5SdkToTerraform(ctx, &diags, data.Band5)
	}
	if data.Band5On24Radio != nil {
		band_5_on_24_radio = band5On24RadioSdkToTerraform(ctx, &diags, data.Band5On24Radio)
	}
	if data.Band6 != nil {
		band_6 = band6SdkToTerraform(ctx, &diags, data.Band6)
	}
	if data.CountryCode != nil {
		country_code = types.StringValue(*data.CountryCode)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if data.ModelSpecific != nil && len(data.ModelSpecific) > 0 {
		model_specific = modelSpecificSdkToTerraform(ctx, &diags, data.ModelSpecific)
	}

	name = types.StringValue(data.Name)

	if data.OrgId != nil {
		org_id = types.StringValue(data.OrgId.String())
	}
	if data.ScanningEnabled != nil {
		scanning_enabled = types.BoolValue(*data.ScanningEnabled)
	}

	state.AntGain24 = ant_gain_24
	state.AntGain5 = ant_gain_5
	state.AntGain6 = ant_gain_6
	state.Band24 = band_24
	state.Band24Usage = band_24_usage
	state.Band5 = band_5
	state.Band5On24Radio = band_5_on_24_radio
	state.Band6 = band_6
	state.CountryCode = country_code
	state.Id = id
	state.ModelSpecific = model_specific
	state.Name = name
	state.OrgId = org_id
	state.ScanningEnabled = scanning_enabled

	return state, diags
}
