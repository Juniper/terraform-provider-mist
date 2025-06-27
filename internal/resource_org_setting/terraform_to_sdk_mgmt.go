package resource_org_setting

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func mgmtTerraformToSdk(d MgmtValue) *models.OrgSettingMgmt {
	data := models.OrgSettingMgmt{}

	if !d.MxtunnelIds.IsNull() && !d.MxtunnelIds.IsUnknown() {
		data.MxtunnelIds = mistutils.ListOfUuidTerraformToSdk(d.MxtunnelIds)
	}

	if d.UseMxtunnel.ValueBoolPointer() != nil {
		data.UseMxtunnel = d.UseMxtunnel.ValueBoolPointer()
	}

	if d.UseWxtunnel.ValueBoolPointer() != nil {
		data.UseWxtunnel = d.UseWxtunnel.ValueBoolPointer()
	}

	return &data
}
