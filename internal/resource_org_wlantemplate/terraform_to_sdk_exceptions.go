package resource_org_wlantemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func exceptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan ExceptionsValue) *models.TemplateExceptions {

	data := models.TemplateExceptions{}
	if !plan.SiteIds.IsNull() && !plan.SiteIds.IsUnknown() {
		data.SiteIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.SiteIds)
	}
	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		data.SitegroupIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.SitegroupIds)
	}
	return &data
}
