package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternetAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternetAccessValue) *models.NetworkInternetAccess {
	destination_nat := destinationNatTerraformToSdk(ctx, diags, d.DestinationNat)
	static_nat := staticNatTerraformToSdk(ctx, diags, d.StaticNat)
	data := models.NetworkInternetAccess{}
	data.CreateSimpleServicePolicy = d.CreateSimpleServicePolicy.ValueBoolPointer()
	data.DestinationNat = destination_nat
	data.Enabled = d.Enabled.ValueBoolPointer()
	data.Restricted = d.Restricted.ValueBoolPointer()
	data.StaticNat = static_nat

	return &data
}
