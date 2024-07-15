package resource_site_wlan

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_hours "terraform-provider-mist/internal/commons/hours"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func scheduleTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ScheduleValue) *models.WlanSchedule {
	data := models.WlanSchedule{}

	data.Enabled = d.Enabled.ValueBoolPointer()

	data.Hours = mist_hours.HoursTerraformToSdk(ctx, diags, d.Hours)

	return &data
}
