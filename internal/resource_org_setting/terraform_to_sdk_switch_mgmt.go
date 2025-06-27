package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func switchMgmtTerraformToSdk(d SwitchMgmtValue) *models.OrgSettingSwitchMgmt {
	data := models.OrgSettingSwitchMgmt{}

	if d.ApAffinityThreshold.ValueInt64Pointer() != nil {
		data.ApAffinityThreshold = models.ToPointer(int(d.ApAffinityThreshold.ValueInt64()))
	}

	return &data
}
