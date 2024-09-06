package resource_org_wlan_portal_template

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWlanPortalTemplateModel) (models.WlanPortalTemplate, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.WlanPortalTemplate{}

	data.PortalTemplate = portalTemplateTerraformToSdk(ctx, &diags, &plan.PortalTemplate)

	return data, diags
}
