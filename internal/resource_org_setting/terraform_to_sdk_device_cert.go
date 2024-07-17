package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func deviceCertTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d DeviceCertValue) *models.OrgSettingDeviceCert {
	data := models.OrgSettingDeviceCert{}

	if d.Cert.ValueStringPointer() != nil {
		data.Cert = d.Cert.ValueStringPointer()
	}

	if d.Key.ValueStringPointer() != nil {
		data.Key = d.Key.ValueStringPointer()
	}

	return &data
}
