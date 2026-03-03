package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func versionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d VersionsValue) *models.MxedgeVersions {
	data := models.MxedgeVersions{}

	if !d.Mxagent.IsNull() && !d.Mxagent.IsUnknown() {
		data.Mxagent = d.Mxagent.ValueStringPointer()
	}

	if !d.Tunterm.IsNull() && !d.Tunterm.IsUnknown() {
		data.Tunterm = d.Tunterm.ValueStringPointer()
	}

	return &data
}
