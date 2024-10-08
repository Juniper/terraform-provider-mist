package resource_site_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func ssrTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SsrValue) *models.SiteSettingSsr {
	data := models.SiteSettingSsr{}

	data.ConductorHosts = mist_transform.ListOfStringTerraformToSdk(ctx, d.ConductorHosts)
	data.DisableStats = d.DisableStats.ValueBoolPointer()

	return &data
}
