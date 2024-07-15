package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func rogueSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteRogue) RogueValue {
	tflog.Debug(ctx, "rogueSdkToTerraform")

	var enabled basetypes.BoolValue
	var honeypot_enabled basetypes.BoolValue
	var min_duration basetypes.Int64Value
	var min_rssi basetypes.Int64Value
	var whitelisted_bssids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var whitelisted_ssids basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.HoneypotEnabled != nil {
		honeypot_enabled = types.BoolValue(*d.HoneypotEnabled)
	}
	if d != nil && d.MinDuration != nil {
		min_duration = types.Int64Value(int64(*d.MinDuration))
	}
	if d != nil && d.MinRssi != nil {
		min_rssi = types.Int64Value(int64(*d.MinRssi))
	}
	if d != nil && d.WhitelistedBssids != nil {
		whitelisted_bssids = mist_transform.ListOfStringSdkToTerraform(ctx, d.WhitelistedBssids)
	}
	if d != nil && d.WhitelistedSsids != nil {
		whitelisted_ssids = mist_transform.ListOfStringSdkToTerraform(ctx, d.WhitelistedSsids)
	}

	data_map_attr_type := RogueValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled":            enabled,
		"honeypot_enabled":   honeypot_enabled,
		"min_duration":       min_duration,
		"min_rssi":           min_rssi,
		"whitelisted_bssids": whitelisted_bssids,
		"whitelisted_ssids":  whitelisted_ssids,
	}
	data, e := NewRogueValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
