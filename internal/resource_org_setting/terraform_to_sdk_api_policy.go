package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func apiPolicyTerraformToSdk(d ApiPolicyValue) *models.OrgSettingApiPolicy {
	data := models.OrgSettingApiPolicy{}

	if d.NoReveal.ValueBoolPointer() != nil {
		data.NoReveal = d.NoReveal.ValueBoolPointer()
	}
	return &data
}
