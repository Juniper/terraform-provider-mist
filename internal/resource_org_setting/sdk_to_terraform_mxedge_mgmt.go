package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mxedgeMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeMgmt) MxedgeMgmtValue {
	var configAutoRevert basetypes.BoolValue
	var fipsEnabled basetypes.BoolValue
	var mistPassword basetypes.StringValue
	var oobIpType basetypes.StringValue
	var oobIpType6 basetypes.StringValue
	var rootPassword basetypes.StringValue

	if d.ConfigAutoRevert != nil {
		configAutoRevert = types.BoolValue(*d.ConfigAutoRevert)
	}
	if d.FipsEnabled != nil {
		fipsEnabled = types.BoolValue(*d.FipsEnabled)
	}
	if d.MistPassword != nil {
		mistPassword = types.StringValue(*d.MistPassword)
	}
	if d.OobIpType != nil {
		oobIpType = types.StringValue(string(*d.OobIpType))
	}
	if d.OobIpType6 != nil {
		oobIpType6 = types.StringValue(string(*d.OobIpType6))
	}
	if d.RootPassword != nil {
		rootPassword = types.StringValue(*d.RootPassword)
	}

	dataMapValue := map[string]attr.Value{
		"config_auto_revert": configAutoRevert,
		"fips_enabled":       fipsEnabled,
		"mist_password":      mistPassword,
		"oob_ip_type":        oobIpType,
		"oob_ip_type6":       oobIpType6,
		"root_password":      rootPassword,
	}
	data, e := NewMxedgeMgmtValue(MxedgeMgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
