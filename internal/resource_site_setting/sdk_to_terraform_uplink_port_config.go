package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func uplinkPortConfigValueSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApUplinkPortConfig) UplinkPortConfigValue {
	var dot1x basetypes.BoolValue
	var keepWlansUpIfDown = types.BoolValue(true)

	if d.Dot1x != nil {
		dot1x = types.BoolValue(*d.Dot1x)
	}
	if d.KeepWlansUpIfDown != nil {
		keepWlansUpIfDown = types.BoolValue(*d.KeepWlansUpIfDown)
	}

	dataMapValue := map[string]attr.Value{
		"dot1x":                 dot1x,
		"keep_wlans_up_if_down": keepWlansUpIfDown,
	}
	data, e := NewUplinkPortConfigValue(UplinkPortConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
