package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func radsecTlsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RadsecTlsValue) *models.MxclusterRadsecTls {
	data := models.MxclusterRadsecTls{}

	if !d.Keypair.IsNull() && !d.Keypair.IsUnknown() {
		data.Keypair = d.Keypair.ValueStringPointer()
	}

	return &data
}
