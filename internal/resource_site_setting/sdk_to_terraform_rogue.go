package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func rogueSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteRogue) RogueValue {

	var allowedVlanIds = mistutils.ListOfIntSdkToTerraformEmpty()
	var enabled basetypes.BoolValue
	var honeypotEnabled basetypes.BoolValue
	var minDuration basetypes.Int64Value
	var minRogueDuration basetypes.Int64Value
	var minRssi basetypes.Int64Value
	var minRogueRssi basetypes.Int64Value
	var whitelistedBssids = mistutils.ListOfStringSdkToTerraformEmpty()
	var whitelistedSsids = mistutils.ListOfStringSdkToTerraformEmpty()

	if d != nil && len(d.AllowedVlanIds) > 0 {
		allowedVlanIds = mistutils.ListOfIntSdkToTerraform(d.AllowedVlanIds)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.HoneypotEnabled != nil {
		honeypotEnabled = types.BoolValue(*d.HoneypotEnabled)
	}
	if d != nil && d.MinDuration != nil {
		minDuration = types.Int64Value(int64(*d.MinDuration))
	}
	if d != nil && d.MinRogueDuration != nil {
		minRogueDuration = types.Int64Value(int64(*d.MinRogueDuration))
	}
	if d != nil && d.MinRssi != nil {
		minRssi = types.Int64Value(int64(*d.MinRssi))
	}
	if d != nil && d.MinRogueRssi != nil {
		minRogueRssi = types.Int64Value(int64(*d.MinRogueRssi))
	}
	if d != nil && len(d.WhitelistedBssids) > 0 {
		whitelistedBssids = mistutils.ListOfStringSdkToTerraform(d.WhitelistedBssids)
	}
	if d != nil && len(d.WhitelistedSsids) > 0 {
		whitelistedSsids = mistutils.ListOfStringSdkToTerraform(d.WhitelistedSsids)
	}

	dataMapValue := map[string]attr.Value{
		"allowed_vlan_ids":   allowedVlanIds,
		"enabled":            enabled,
		"honeypot_enabled":   honeypotEnabled,
		"min_duration":       minDuration,
		"min_rogue_duration": minRogueDuration,
		"min_rssi":           minRssi,
		"min_rogue_rssi":     minRogueRssi,
		"whitelisted_bssids": whitelistedBssids,
		"whitelisted_ssids":  whitelistedSsids,
	}
	data, e := NewRogueValue(RogueValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
