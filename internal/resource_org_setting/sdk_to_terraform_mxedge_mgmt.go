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
	var fips_enabled basetypes.BoolValue
	var mist_password basetypes.StringValue
	var oob_ip_type basetypes.StringValue
	var oob_ip_type6 basetypes.StringValue
	var root_password basetypes.StringValue

	if d.FipsEnabled != nil {
		fips_enabled = types.BoolValue(*d.FipsEnabled)
	}
	if d.MistPassword != nil {
		mist_password = types.StringValue(*d.MistPassword)
	}
	if d.OobIpType != nil {
		oob_ip_type = types.StringValue(string(*d.OobIpType))
	}
	if d.OobIpType6 != nil {
		oob_ip_type6 = types.StringValue(string(*d.OobIpType6))
	}
	if d.RootPassword != nil {
		root_password = types.StringValue(*d.RootPassword)
	}

	data_map_attr_type := MxedgeMgmtValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"fips_enabled":  fips_enabled,
		"mist_password": mist_password,
		"oob_ip_type":   oob_ip_type,
		"oob_ip_type6":  oob_ip_type6,
		"root_password": root_password,
	}
	data, e := NewMxedgeMgmtValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
