package resource_org_wlantemplate

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func exceptionsTerraformToSdk(plan ExceptionsValue) *models.TemplateExceptions {

	data := models.TemplateExceptions{}
	if !plan.SiteIds.IsNull() && !plan.SiteIds.IsUnknown() {
		data.SiteIds = mistutils.ListOfUuidTerraformToSdk(plan.SiteIds)
	}
	if !plan.SitegroupIds.IsNull() && !plan.SitegroupIds.IsUnknown() {
		data.SitegroupIds = mistutils.ListOfUuidTerraformToSdk(plan.SitegroupIds)
	}
	return &data
}
