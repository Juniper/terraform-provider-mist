package resource_site_wlan

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ciscoCwaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanCiscoCwa) CiscoCwaValue {
	var allowedHostnames = mistutils.ListOfStringSdkToTerraformEmpty()
	var allowedSubnets = mistutils.ListOfStringSdkToTerraformEmpty()
	var blockedSubnets = mistutils.ListOfStringSdkToTerraformEmpty()
	var enabled basetypes.BoolValue

	if d != nil && d.AllowedHostnames != nil {
		allowedHostnames = mistutils.ListOfStringSdkToTerraform(d.AllowedHostnames)
	}
	if d != nil && d.AllowedSubnets != nil {
		allowedSubnets = mistutils.ListOfStringSdkToTerraform(d.AllowedSubnets)
	}
	if d != nil && d.BlockedSubnets != nil {
		blockedSubnets = mistutils.ListOfStringSdkToTerraform(d.BlockedSubnets)
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
