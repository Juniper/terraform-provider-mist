package resource_site_wlan

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func airwatchTerraformToSdk(plan AirwatchValue) *models.WlanAirwatch {

	data := models.WlanAirwatch{}

	data.ApiKey = plan.ApiKey.ValueStringPointer()
	data.ConsoleUrl = plan.ConsoleUrl.ValueStringPointer()
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.Password = plan.Password.ValueStringPointer()
	data.Username = plan.Username.ValueStringPointer()

	return &data
}
