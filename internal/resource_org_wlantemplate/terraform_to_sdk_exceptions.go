package resource_org_wlantemplate

import (
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func exceptionsTerraformToSdk(plan ExceptionsValue) *models.TemplateExceptions {

	data := models.TemplateExceptions{}
	if !plan.SiteIds.IsNull() && !plan.SiteIds.IsUnknown() {
		data.SiteIds = misttransform.ListOfUuidTerraformToSdk(plan.SiteIds)
	}
	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		data.SitegroupIds = misttransform.ListOfUuidTerraformToSdk(plan.SitegroupIds)
	}
	return &data
}
