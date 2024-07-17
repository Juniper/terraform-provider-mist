package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func securityTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SecurityValue) *models.OrgSettingSecurity {
	data := models.OrgSettingSecurity{}

	if d.DisableLocalSsh.ValueBoolPointer() != nil {
		data.DisableLocalSsh = d.DisableLocalSsh.ValueBoolPointer()
	}

	if d.FipsZeroizePassword.ValueStringPointer() != nil {
		data.FipsZeroizePassword = d.FipsZeroizePassword.ValueStringPointer()
	}

	if d.LimitSshAccess.ValueBoolPointer() != nil {
		data.LimitSshAccess = d.LimitSshAccess.ValueBoolPointer()
	}

	return &data
}
