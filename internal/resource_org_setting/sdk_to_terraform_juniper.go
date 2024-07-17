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

	var data_list = []AccountsValue{}
	for _, d := range l {
		var linked_by basetypes.StringValue
		var name basetypes.StringValue

		if d.LinkedBy != nil {
			linked_by = types.StringValue(*d.LinkedBy)
		}
		if d.Name != nil {
			name = types.StringValue(*d.Name)
		}

		data_map_attr_type := AccountsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"linked_by": linked_by,
			"name":      name,
		}
		data, e := NewAccountsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, AccountsValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func juniperSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.AccountJuniperInfo) JuniperValue {
	var accounts basetypes.ListValue = types.ListNull(AccountsValue{}.Type(ctx))

	if d.Accounts != nil {
		accounts = juniperAccountsSdkToTerraform(ctx, diags, d.Accounts)
	}

	data_map_attr_type := JuniperValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"accounts": accounts,
	}
	data, e := NewJuniperValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
