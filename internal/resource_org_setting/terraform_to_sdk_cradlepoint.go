package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func cradlepointTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CradlepointValue) *models.AccountCradlepointConfig {
	data := models.AccountCradlepointConfig{}

	if d.CpApiId.ValueStringPointer() != nil {
		data.CpApiId = d.CpApiId.ValueStringPointer()
	}

	if d.CpApiKey.ValueStringPointer() != nil {
		data.CpApiKey = d.CpApiKey.ValueStringPointer()
	}

	if d.EcmApiId.ValueStringPointer() != nil {
		data.EcmApiId = d.EcmApiId.ValueStringPointer()
	}

	if d.EcmApiKey.ValueStringPointer() != nil {
		data.EcmApiKey = d.EcmApiKey.ValueStringPointer()
	}

	return &data
}
