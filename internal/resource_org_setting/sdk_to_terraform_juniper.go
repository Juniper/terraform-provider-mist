package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func juniperAccountsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.JuniperAccount) basetypes.ListValue {

	var dataList []AccountsValue
	for _, d := range l {
		var linkedBy basetypes.StringValue
		var name basetypes.StringValue

		if d.LinkedBy != nil {
			linkedBy = types.StringValue(*d.LinkedBy)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		dataMapValue := map[string]attr.Value{
			"linked_by": linkedBy,
			"name":      name,
		}
		data, e := NewAccountsValue(AccountsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, AccountsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func juniperSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AccountJuniperInfo) JuniperValue {
	var accounts = types.ListNull(AccountsValue{}.Type(ctx))

	if d.Accounts != nil {
		accounts = juniperAccountsSdkToTerraform(ctx, diags, d.Accounts)
	}

	dataMapValue := map[string]attr.Value{
		"accounts": accounts,
	}
	data, e := NewJuniperValue(JuniperValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
