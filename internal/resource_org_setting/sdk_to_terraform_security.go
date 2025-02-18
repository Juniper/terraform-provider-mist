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

	var disableLocalSsh basetypes.BoolValue
	var fipsZeroizePassword basetypes.StringValue
	var limitSshAccess basetypes.BoolValue

	if d.DisableLocalSsh != nil {
		disableLocalSsh = types.BoolValue(*d.DisableLocalSsh)
	}
	if d.FipsZeroizePassword != nil {
		fipsZeroizePassword = types.StringValue(*d.FipsZeroizePassword)
	}
	if d.LimitSshAccess != nil {
		limitSshAccess = types.BoolValue(*d.LimitSshAccess)
	}

	dataMapValue := map[string]attr.Value{
		"disable_local_ssh":     disableLocalSsh,
		"fips_zeroize_password": fipsZeroizePassword,
		"limit_ssh_access":      limitSshAccess,
	}
	data, e := NewSecurityValue(SecurityValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
