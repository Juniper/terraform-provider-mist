package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMgmtTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d SwitchMgmtValue) *models.OrgSettingSwitchMgmt {
	data := models.OrgSettingSwitchMgmt{}

	if d.ApAffinityThreshold.ValueInt64Pointer() != nil {
		data.ApAffinityThreshold = models.ToPointer(int(d.ApAffinityThreshold.ValueInt64()))
	}

	return &data
}
