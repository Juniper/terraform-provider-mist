package resource_org_rftemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgRftemplateModel) (*models.RfTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.RfTemplate{}
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueString()

	if plan.AntGain24.IsNull() || plan.AntGain24.IsUnknown() {
		unset["-ant_gain_24"] = ""
	} else {
		data.AntGain24 = models.ToPointer(int(plan.AntGain24.ValueInt64()))
	}

	if plan.AntGain5.IsNull() || plan.AntGain5.IsUnknown() {
		unset["-ant_gain_5"] = ""
	} else {
		data.AntGain5 = models.ToPointer(int(plan.AntGain5.ValueInt64()))
	}

	if plan.AntGain6.IsNull() || plan.AntGain6.IsUnknown() {
		unset["-ant_gain_6ﬂ"] = ""
	} else {
		data.AntGain6 = models.ToPointer(int(plan.AntGain6.ValueInt64()))
	}

	if plan.Band24Usage.IsNull() || plan.Band24Usage.IsUnknown() {
		unset["-band_24_usage"] = ""
	} else {
		data.Band24Usage = models.ToPointer(models.RadioBand24UsageEnum(string(plan.Band24Usage.ValueString())))
	}

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

	if plan.Band5On24Radio.IsNull() || plan.Band5On24Radio.IsUnknown() {
		unset["-band_5_on_24_radio"] = ""
	} else {
		data.Band5On24Radio = band5On24RadioTerraformToSdk(ctx, &diags, plan.Band5On24Radio)
	}

	if plan.Band6.IsNull() || plan.Band6.IsUnknown() {
		unset["-band_6"] = ""
	} else {
		data.Band6 = band6TerraformToSdk(ctx, &diags, plan.Band6)
	}

	if plan.CountryCode.IsNull() || plan.CountryCode.IsUnknown() {
		unset["-country_code"] = ""
	} else {
		data.CountryCode = plan.CountryCode.ValueStringPointer()
	}

	if plan.ModelSpecific.IsNull() || plan.ModelSpecific.IsUnknown() {
		unset["-model_specific"] = ""
	} else {
		data.ModelSpecific = modelSpecificTerraformToSdk(ctx, &diags, plan.ModelSpecific)
	}

	if plan.ScanningEnabled.IsNull() || plan.ScanningEnabled.IsUnknown() {
		unset["-scanning_enabled"] = ""
	} else {
		data.ScanningEnabled = plan.ScanningEnabled.ValueBoolPointer()
	}

	data.AdditionalProperties = unset
	return &data, diags
}
