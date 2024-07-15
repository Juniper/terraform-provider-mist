package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func airwatchTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AirwatchValue) *models.WlanAirwatch {

	data := models.WlanAirwatch{}

	data.ApiKey = plan.ApiKey.ValueStringPointer()
	data.ConsoleUrl = plan.ConsoleUrl.ValueStringPointer()
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.Password = plan.Password.ValueStringPointer()
	data.Username = plan.Username.ValueStringPointer()

	return &data
}
