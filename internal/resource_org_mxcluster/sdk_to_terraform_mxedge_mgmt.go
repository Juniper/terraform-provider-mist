package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mxedgeMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeMgmt) MxedgeMgmtValue {

	var configAutoRevert = types.BoolNull()
	var fipsEnabled = types.BoolNull()
	var mistPassword = types.StringNull()
	var oobIpType = types.StringNull()
	var oobIpType6 = types.StringNull()
	var rootPassword = types.StringNull()

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

	data_map_attr_type := MxedgeMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"config_auto_revert": configAutoRevert,
		"fips_enabled":       fipsEnabled,
		"mist_password":      mistPassword,
		"oob_ip_type":        oobIpType,
		"oob_ip_type6":       oobIpType6,
		"root_password":      rootPassword,
	}
	data, e := NewMxedgeMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
