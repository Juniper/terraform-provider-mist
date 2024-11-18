package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternetAccessValue) *models.NetworkInternetAccess {
	data := models.NetworkInternetAccess{}

	if !d.CreateSimpleServicePolicy.IsNull() && !d.CreateSimpleServicePolicy.IsUnknown() {
		data.CreateSimpleServicePolicy = d.CreateSimpleServicePolicy.ValueBoolPointer()
	}
	if !d.DestinationNat.IsNull() && !d.DestinationNat.IsUnknown() {
		data.DestinationNat = destinationNatTerraformToSdk(ctx, diags, d.DestinationNat)
	}
	if !d.Enabled.IsNull() && !d.Enabled.IsUnknown() {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.StaticNat.IsNull() && !d.StaticNat.IsUnknown() {
		data.StaticNat = staticNatTerraformToSdk(ctx, diags, d.StaticNat)
	}
	if !d.Restricted.IsNull() && !d.Restricted.IsUnknown() {
		data.Restricted = d.Restricted.ValueBoolPointer()
	}

	return &data
}
