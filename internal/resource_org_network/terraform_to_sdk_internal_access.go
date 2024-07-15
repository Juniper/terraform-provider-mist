package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func InternalAccessTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d InternalAccessValue) *models.NetworkInternalAccess {
	data := models.NetworkInternalAccess{}
	data.Enabled = d.Enabled.ValueBoolPointer()
	return &data
}
