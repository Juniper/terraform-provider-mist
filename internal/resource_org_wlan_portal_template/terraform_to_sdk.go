package resource_org_wlan_portal_template

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgWlanPortalTemplateModel) (models.WlanPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.WlanPortalTemplate{}

	data.PortalTemplate = portalTemplateTerraformToSdk(&diags, &plan.PortalTemplate)

	return data, diags
}
