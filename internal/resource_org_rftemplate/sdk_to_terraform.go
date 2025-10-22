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

	var antGain24 types.Int64
	var antGain5 types.Int64
	var antGain6 types.Int64
	var band24 = NewBand24ValueNull()
	var band24Usage types.String
	var band5 = NewBand5ValueNull()
	var band5On24Radio = NewBand5On24RadioValueNull()
	var band6 = NewBand6ValueNull()
	var countryCode types.String
	var id types.String
	var modelSpecific = types.MapNull(ModelSpecificValue{}.Type(ctx))
	var name types.String
	var orgId types.String
	var scanningEnabled types.Bool

	if data.AntGain24 != nil {
		antGain24 = types.Int64Value(int64(*data.AntGain24))
	}
	if data.AntGain5 != nil {
		antGain5 = types.Int64Value(int64(*data.AntGain5))
	}
	if data.AntGain6 != nil {
		antGain6 = types.Int64Value(int64(*data.AntGain6))
	}
	if data.Band24 != nil {
		band24 = band24SdkToTerraform(ctx, &diags, data.Band24)
	}
	if data.Band24Usage != nil {
		band24Usage = types.StringValue(string(*data.Band24Usage))
	}
	if data.Band5 != nil {
		band5 = band5SdkToTerraform(ctx, &diags, data.Band5)
	}
	if data.Band5On24Radio != nil {
		band5On24Radio = band5On24RadioSdkToTerraform(ctx, &diags, data.Band5On24Radio)
	}
	if data.Band6 != nil {
		band6 = band6SdkToTerraform(ctx, &diags, data.Band6)
	}
	if data.CountryCode != nil {
		countryCode = types.StringValue(*data.CountryCode)
	}
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}
	if len(data.ModelSpecific) > 0 {
		modelSpecific = modelSpecificSdkToTerraform(ctx, &diags, data.ModelSpecific)
	}

	name = types.StringValue(data.Name)

	if data.OrgId != nil {
		orgId = types.StringValue(data.OrgId.String())
	}
	if data.ScanningEnabled != nil {
		scanningEnabled = types.BoolValue(*data.ScanningEnabled)
	}

	state.AntGain24 = antGain24
	state.AntGain5 = antGain5
	state.AntGain6 = antGain6
	state.Band24 = band24
	state.Band24Usage = band24Usage
	state.Band5 = band5
	state.Band5On24Radio = band5On24Radio
	state.Band6 = band6
	state.CountryCode = countryCode
	state.Id = id
	state.ModelSpecific = modelSpecific
	state.Name = name
	state.OrgId = orgId
	state.ScanningEnabled = scanningEnabled

	return state, diags
}
