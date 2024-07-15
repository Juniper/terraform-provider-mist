package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func proxyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ProxyValue) *models.Proxy {
	tflog.Debug(ctx, "proxyTerraformToSdk")
	data := models.Proxy{}

	data.Url = d.Url.ValueStringPointer()

	return &data
}
