package resource_org_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func apiPolicyTerraformToSdk(d ApiPolicyValue) *models.OrgSettingApiPolicy {
	data := models.OrgSettingApiPolicy{}

	if d.NoReveal.ValueBoolPointer() != nil {
		data.NoReveal = d.NoReveal.ValueBoolPointer()
	}
	if !d.SrcIps.IsNull() && !d.SrcIps.IsUnknown() {
		data.SrcIps = mistutils.ListOfStringTerraformToSdk(d.SrcIps)
	}
	return &data
}
