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

	state.Id = types.StringValue(data.Id.String())
	state.OrgId = types.StringValue(data.OrgId.String())
	state.Name = types.StringValue(data.Name)

	state.AntGain24 = types.Int64Value(int64(*data.AntGain24))

	state.AntGain5 = types.Int64Value(int64(*data.AntGain5))

	state.AntGain6 = types.Int64Value(int64(*data.AntGain6))

	if data.Band24 != nil {
		state.Band24 = band24SdkToTerraform(ctx, &diags, data.Band24)
	} else {
		state.Band24 = NewBand24ValueNull()
	}

	state.Band24Usage = types.StringValue(string(*data.Band24Usage))

	if data.Band5 != nil {
		state.Band5 = band5SdkToTerraform(ctx, &diags, data.Band5)
	} else {
		state.Band5 = NewBand5ValueNull()
	}

	if data.Band6 != nil {
		state.Band6 = band6SdkToTerraform(ctx, &diags, data.Band6)
	} else {
		state.Band6 = NewBand6ValueNull()
	}

	state.CountryCode = types.StringValue(*data.CountryCode)

	if data.ModelSpecific != nil {
		state.ModelSpecific = modelSpecificSdkToTerraform(ctx, &diags, data.ModelSpecific)
	} else {
		state.ModelSpecific = types.MapNull(ModelSpecificValue{}.Type(ctx))
	}

	return state, diags
}
