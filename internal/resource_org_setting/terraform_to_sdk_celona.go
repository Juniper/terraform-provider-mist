package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func celonaTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CelonaValue) *models.OrgSettingCelona {
	data := models.OrgSettingCelona{}

	if d.ApiKey.ValueStringPointer() != nil {
		data.ApiKey = d.ApiKey.ValueStringPointer()
	}

	if d.ApiPrefix.ValueStringPointer() != nil {
		data.ApiPrefix = d.ApiPrefix.ValueStringPointer()
	}
	return &data
}
