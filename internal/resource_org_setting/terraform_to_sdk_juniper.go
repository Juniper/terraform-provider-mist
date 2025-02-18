package resource_org_setting

import (
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func juniperAccountsTerraformToSdk(l basetypes.ListValue) []models.JuniperAccount {
	var dataList []models.JuniperAccount
	for _, v := range l.Elements() {
		var vInterface interface{} = v
		d := vInterface.(AccountsValue)
		data := models.JuniperAccount{}
		if d.LinkedBy.ValueStringPointer() != nil {
			data.LinkedBy = d.LinkedBy.ValueStringPointer()
		}

		if d.Name.ValueStringPointer() != nil {
			data.Name = d.Name.ValueStringPointer()
		}

		dataList = append(dataList, data)
	}
	return dataList
}

func juniperTerraformToSdk(d JuniperValue) *models.AccountJuniperInfo {
	data := models.AccountJuniperInfo{}

	if !d.Accounts.IsNull() && !d.Accounts.IsUnknown() {
		data.Accounts = juniperAccountsTerraformToSdk(d.Accounts)
	}

	return &data
}
