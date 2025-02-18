package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func rogueSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteRogue) RogueValue {

	var enabled basetypes.BoolValue
	var honeypotEnabled basetypes.BoolValue
	var minDuration basetypes.Int64Value
	var minRssi basetypes.Int64Value
	var whitelistedBssids = misttransform.ListOfStringSdkToTerraformEmpty()
	var whitelistedSsids = misttransform.ListOfStringSdkToTerraformEmpty()

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.HoneypotEnabled != nil {
		honeypotEnabled = types.BoolValue(*d.HoneypotEnabled)
	}
	if d != nil && d.MinDuration != nil {
		minDuration = types.Int64Value(int64(*d.MinDuration))
	}
	if d != nil && d.MinRssi != nil {
		minRssi = types.Int64Value(int64(*d.MinRssi))
	}
	if d != nil && d.WhitelistedBssids != nil {
		whitelistedBssids = misttransform.ListOfStringSdkToTerraform(d.WhitelistedBssids)
	}
	if d != nil && d.WhitelistedSsids != nil {
		whitelistedSsids = misttransform.ListOfStringSdkToTerraform(d.WhitelistedSsids)
	}

	dataMapValue := map[string]attr.Value{
		"enabled":            enabled,
		"honeypot_enabled":   honeypotEnabled,
		"min_duration":       minDuration,
		"min_rssi":           minRssi,
		"whitelisted_bssids": whitelistedBssids,
		"whitelisted_ssids":  whitelistedSsids,
	}
	data, e := NewRogueValue(RogueValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
