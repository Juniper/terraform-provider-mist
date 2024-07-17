package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func juniperAccountsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, l basetypes.ListValue) []models.JuniperAccount {
	var data_list []models.JuniperAccount
	for _, v := range l.Elements() {
		var v_interface interface{} = v
		d := v_interface.(AccountsValue)
		data := models.JuniperAccount{}
		if d.LinkedBy.ValueStringPointer() != nil {
			data.LinkedBy = d.LinkedBy.ValueStringPointer()
		}

		if d.Name.ValueStringPointer() != nil {
			data.Name = d.Name.ValueStringPointer()
		}

		data_list = append(data_list, data)
	}
	return data_list
}

func juniperTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d JuniperValue) *models.AccountJuniperInfo {
	data := models.AccountJuniperInfo{}

	if !d.Accounts.IsNull() && !d.Accounts.IsUnknown() {
		data.Accounts = juniperAccountsTerraformToSdk(ctx, diags, d.Accounts)
	}

	return &data
}
