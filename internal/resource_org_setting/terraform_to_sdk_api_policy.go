package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func apiPolicyTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d ApiPolicyValue) *models.OrgSettingApiPolicy {
	data := models.OrgSettingApiPolicy{}

	if d.NoReveal.ValueBoolPointer() != nil {
		data.NoReveal = d.NoReveal.ValueBoolPointer()
	}
	return &data
}
