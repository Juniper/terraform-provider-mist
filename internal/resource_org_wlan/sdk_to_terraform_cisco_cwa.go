package resource_org_wlan

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ciscoCwaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanCiscoCwa) CiscoCwaValue {
	var allowed_hostnames basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var allowed_subnets basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var blocked_subnets basetypes.ListValue = mist_transform.ListOfIntSdkToTerraformEmpty(ctx)
	var enabled basetypes.BoolValue

	if d != nil && d.AllowedHostnames != nil {
		allowed_hostnames = mist_transform.ListOfStringSdkToTerraform(ctx, d.AllowedHostnames)
	}
	if d != nil && d.AllowedSubnets != nil {
		allowed_subnets = mist_transform.ListOfStringSdkToTerraform(ctx, d.AllowedSubnets)
	}
	if d != nil && d.BlockedSubnets != nil {
		blocked_subnets = mist_transform.ListOfStringSdkToTerraform(ctx, d.BlockedSubnets)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := CiscoCwaValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"allowed_hostnames": allowed_hostnames,
		"allowed_subnets":   allowed_subnets,
		"blocked_subnets":   blocked_subnets,
		"enabled":           enabled,
	}
	data, e := NewCiscoCwaValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
