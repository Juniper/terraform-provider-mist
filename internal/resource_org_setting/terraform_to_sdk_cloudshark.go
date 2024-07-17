package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func cloudsharkTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d CloudsharkValue) *models.OrgSettingCloudshark {
	data := models.OrgSettingCloudshark{}

	if d.Apitoken.ValueStringPointer() != nil {
		data.Apitoken = d.Apitoken.ValueStringPointer()
	}

	if d.Url.ValueStringPointer() != nil {
		data.Url = d.Url.ValueStringPointer()
	}

	return &data
}
