package resource_org_wlantemplate

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
)

func appliesTerraformToSdk(plan AppliesValue) *models.TemplateApplies {

	data := models.TemplateApplies{}
	if plan.OrgId.ValueStringPointer() != nil {
		data.OrgId = models.ToPointer(uuid.MustParse(plan.OrgId.ValueString()))
	}
	if !plan.SiteIds.IsNull() && !plan.SiteIds.IsUnknown() {
		data.SiteIds = misttransform.ListOfUuidTerraformToSdk(plan.SiteIds)
	}
	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(plan.SitegroupIds)
	}
	return &data
}
