package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func proxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ProxyValue) *models.Proxy {
	data := models.Proxy{}

	if !d.Url.IsNull() && !d.Url.IsUnknown() {
		data.Url = d.Url.ValueStringPointer()
	}

	return &data
}
