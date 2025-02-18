package resource_site_wlan

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ciscoCwaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanCiscoCwa) CiscoCwaValue {
	var allowedHostnames = misttransform.ListOfStringSdkToTerraformEmpty()
	var allowedSubnets = misttransform.ListOfStringSdkToTerraformEmpty()
	var blockedSubnets = misttransform.ListOfStringSdkToTerraformEmpty()
	var enabled basetypes.BoolValue

	if d != nil && d.AllowedHostnames != nil {
		allowedHostnames = misttransform.ListOfStringSdkToTerraform(d.AllowedHostnames)
	}
	if d != nil && d.AllowedSubnets != nil {
		allowedSubnets = misttransform.ListOfStringSdkToTerraform(d.AllowedSubnets)
	}
	if d != nil && d.BlockedSubnets != nil {
		blockedSubnets = misttransform.ListOfStringSdkToTerraform(d.BlockedSubnets)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"allowed_hostnames": allowedHostnames,
		"allowed_subnets":   allowedSubnets,
		"blocked_subnets":   blockedSubnets,
		"enabled":           enabled,
	}
	data, e := NewCiscoCwaValue(CiscoCwaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
