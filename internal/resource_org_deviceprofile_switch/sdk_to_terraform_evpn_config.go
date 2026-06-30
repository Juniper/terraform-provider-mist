package resource_org_deviceprofile_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func evpnConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnConfig) EvpnConfigValue {
	var enabled basetypes.BoolValue
	var role basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Role != nil {
		role = types.StringValue(string(*d.Role))
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"role":    role,
	}
	data, e := NewEvpnConfigValue(EvpnConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
