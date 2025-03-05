package resource_org_setting

import (
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func junosShellAccessTerraformToSdk(d JunosShellAccessValue) *models.OrgSettingJunosShellAccess {
	var data models.OrgSettingJunosShellAccess

	if d.Admin.ValueStringPointer() != nil {
		data.Admin = (*models.OrgSettingJunosShellAccessAdminEnum)(d.Admin.ValueStringPointer())
	}
	if d.Helpdesk.ValueStringPointer() != nil {
		data.Helpdesk = (*models.OrgSettingJunosShellAccessHelpdeskEnum)(d.Helpdesk.ValueStringPointer())
	}
	if d.Read.ValueStringPointer() != nil {
		data.Read = (*models.OrgSettingJunosShellAccessReadEnum)(d.Read.ValueStringPointer())
	}
	if d.Write.ValueStringPointer() != nil {
		data.Write = (*models.OrgSettingJunosShellAccessWriteEnum)(d.Write.ValueStringPointer())
	}

	return &data
}
