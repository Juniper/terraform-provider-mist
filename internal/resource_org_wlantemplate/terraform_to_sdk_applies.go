package resource_org_wlantemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func appliesTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan AppliesValue) *models.TemplateApplies {

	data := models.TemplateApplies{}
	if plan.OrgId.ValueStringPointer() != nil {
		data.OrgId = models.ToPointer(uuid.MustParse(plan.OrgId.ValueString()))
	}
	if !plan.SiteIds.IsNull() && !plan.SiteIds.IsUnknown() {
		data.SiteIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.SiteIds)
	}
	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		data.SitegroupIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.SitegroupIds)
	}
	return &data
}
