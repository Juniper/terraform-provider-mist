package resource_org_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func securitySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingSecurity) SecurityValue {

	var disable_local_ssh basetypes.BoolValue
	var fips_zeroize_password basetypes.StringValue
	var limit_ssh_access basetypes.BoolValue

	if d.DisableLocalSsh != nil {
		disable_local_ssh = types.BoolValue(*d.DisableLocalSsh)
	}
	if d.FipsZeroizePassword != nil {
		fips_zeroize_password = types.StringValue(*d.FipsZeroizePassword)
	}
	if d.LimitSshAccess != nil {
		limit_ssh_access = types.BoolValue(*d.LimitSshAccess)
	}

	data_map_attr_type := SecurityValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"disable_local_ssh":     disable_local_ssh,
		"fips_zeroize_password": fips_zeroize_password,
		"limit_ssh_access":      limit_ssh_access,
	}
	data, e := NewSecurityValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
