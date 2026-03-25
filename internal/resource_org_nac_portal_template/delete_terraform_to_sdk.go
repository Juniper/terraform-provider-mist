package resource_org_nac_portal_template

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func DeleteTerraformToSdk() (*models.NacPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reset to default values
	alignment := models.PortalTemplateAlignmentEnum("center")
	color := "#1074bc"
	poweredBy := false

	data := models.NacPortalTemplate{
		Alignment: &alignment,
		Color:     &color,
		Logo:      nil,
		PoweredBy: &poweredBy,
	}

	return &data, diags
}
