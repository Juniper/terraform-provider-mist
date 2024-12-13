package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func dynamicPskTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicPskValue) *models.WlanDynamicPsk {

	data := models.WlanDynamicPsk{}
	data.DefaultPsk = plan.DefaultPsk.ValueStringPointer()
	data.DefaultVlanId = models.ToPointer(models.VlanIdWithVariableContainer.FromString(plan.DefaultVlanId.ValueString()))
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.ForceLookup = plan.ForceLookup.ValueBoolPointer()
	data.Source = models.ToPointer(models.DynamicPskSourceEnum(string(plan.Source.ValueString())))

	return &data
}
