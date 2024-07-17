package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func rogueTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d RogueValue) *models.SiteRogue {
	data := models.SiteRogue{}

	data.Enabled = d.Enabled.ValueBoolPointer()
	data.HoneypotEnabled = d.HoneypotEnabled.ValueBoolPointer()
	data.MinDuration = models.ToPointer(int(d.MinDuration.ValueInt64()))
	data.MinRssi = models.ToPointer(int(d.MinRssi.ValueInt64()))
	data.WhitelistedBssids = mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedBssids)
	data.WhitelistedSsids = mist_transform.ListOfStringTerraformToSdk(ctx, d.WhitelistedSsids)

	return &data
}
