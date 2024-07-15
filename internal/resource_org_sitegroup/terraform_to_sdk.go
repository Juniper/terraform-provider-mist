package resource_org_sitegroup

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(plan *OrgSitegroupModel) (*models.Sitegroup, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := models.Sitegroup{}
	data.Name = plan.Name.ValueString()

	return &data, diags
}
