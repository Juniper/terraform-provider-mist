package resource_org_mxtunnel

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func autoPreemptionTerraformToSdk(_ context.Context, _ *diag.Diagnostics, d AutoPreemptionValue) *models.AutoPreemption {
	data := models.AutoPreemption{}

	if !d.DayOfWeek.IsNull() && !d.DayOfWeek.IsUnknown() && d.DayOfWeek.ValueString() != "" {
		data.DayOfWeek = (*models.DayOfWeekEnum)(d.DayOfWeek.ValueStringPointer())
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.TimeOfDay.IsNull() && !d.TimeOfDay.IsUnknown() {
		data.TimeOfDay = d.TimeOfDay.ValueStringPointer()
	}

	return &data
}
