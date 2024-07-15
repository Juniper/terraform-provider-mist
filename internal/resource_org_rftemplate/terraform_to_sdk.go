package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgRftemplateModel) (*models.RfTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	unset := make(map[string]interface{})

	data := models.RfTemplate{}

	data.Name = plan.Name.ValueString()

	data.AntGain24 = models.ToPointer(int(plan.AntGain24.ValueInt64()))
	data.AntGain5 = models.ToPointer(int(plan.AntGain5.ValueInt64()))
	data.AntGain6 = models.ToPointer(int(plan.AntGain6.ValueInt64()))

	data.Band24Usage = models.ToPointer(models.RadioBand24UsageEnum(string(plan.Band24Usage.ValueString())))

	if plan.Band24.IsNull() || plan.Band24.IsUnknown() {
		unset["-band_24"] = ""
	} else {
		data.Band24 = band24TerraformToSdk(ctx, &diags, plan.Band24)
	}

	if plan.Band5.IsNull() || plan.Band5.IsUnknown() {
		unset["-band_5"] = ""
	} else {
		data.Band5 = band5TerraformToSdk(ctx, &diags, plan.Band5)
	}

	if plan.Band6.IsNull() || plan.Band6.IsUnknown() {
		unset["-band_6"] = ""
	} else {
		data.Band6 = band6TerraformToSdk(ctx, &diags, plan.Band6)
	}

	data.CountryCode = plan.CountryCode.ValueStringPointer()

	if plan.ModelSpecific.IsNull() || plan.ModelSpecific.IsUnknown() {
		unset["-model_specific"] = ""
	} else {
		data.ModelSpecific = modelSpecificTerraformToSdk(ctx, &diags, plan.ModelSpecific)
	}

	data.ScanningEnabled = plan.ScanningEnabled.ValueBoolPointer()

	data.AdditionalProperties = unset

	return &data, diags
}
