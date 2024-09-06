package resource_site_wlan_portal_template

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func DeleteTerraformToSdk(ctx context.Context) (*models.WlanPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.WlanPortalTemplate{}
	tmp := NewPortalTemplateValueNull()
	data.PortalTemplate = portalTemplateTerraformToSdk(ctx, &diags, &tmp)
	return &data, diags
}
