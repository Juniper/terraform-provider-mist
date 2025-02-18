package resource_org_wlan_portal_template

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func DeleteTerraformToSdk() (*models.WlanPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.WlanPortalTemplate{}
	tmp := NewPortalTemplateValueNull()
	data.PortalTemplate = portalTemplateTerraformToSdk(&diags, &tmp)
	return &data, diags
}
