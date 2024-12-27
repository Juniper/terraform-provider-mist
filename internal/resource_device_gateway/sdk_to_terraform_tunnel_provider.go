package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderJseSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TunnelProviderOptions) (basetypes.ObjectValue, bool) {

	var org_name basetypes.StringValue
	var num_users basetypes.Int64Value
	configured := false

	if d != nil && d.Jse != nil && d.Jse.OrgName != nil {
		org_name = types.StringValue(*d.Jse.OrgName)
		configured = true
	}
	if d != nil && d.Jse != nil && d.Jse.NumUsers != nil {
		num_users = types.Int64Value(int64(*d.Jse.NumUsers))
		configured = true
	}

	r_attr_value := map[string]attr.Value{
		"org_name":  org_name,
		"num_users": num_users,
	}
	r, e := basetypes.NewObjectValue(JseValue{}.AttributeTypes(ctx), r_attr_value)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderZscalerSubLocationSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, t *models.TunnelProviderOptions) basetypes.ListValue {
	var data_list = []SubLocationsValue{}
	if t != nil && t.Zscaler != nil && t.Zscaler.SubLocations != nil {
		for _, d := range t.Zscaler.SubLocations {
			var aup_block_internet_until_accepted basetypes.BoolValue
			var aup_enabled basetypes.BoolValue
			var aup_force_ssl_inspection basetypes.BoolValue
			var aup_timeout_in_days basetypes.Int64Value
			var auth_required basetypes.BoolValue
			var caution_enabled basetypes.BoolValue
			var dn_bandwidth basetypes.Float64Value
			var idle_time_in_minutes basetypes.Int64Value
			var name basetypes.StringValue
			var ofw_enabled basetypes.BoolValue
			var surrogate_ip basetypes.BoolValue
			var surrogate_ip_enforced_for_known_browsers basetypes.BoolValue
			var surrogate_refresh_time_in_minutes basetypes.Int64Value
			var up_bandwidth basetypes.Float64Value

			if d.AupBlockInternetUntilAccepted != nil {
				aup_block_internet_until_accepted = types.BoolValue(*d.AupBlockInternetUntilAccepted)
			}
			if d.AupEnabled != nil {
				aup_enabled = types.BoolValue(*d.AupEnabled)
			}
			if d.AupForceSslInspection != nil {
				aup_force_ssl_inspection = types.BoolValue(*d.AupForceSslInspection)
			}
			if d.AupTimeoutInDays != nil {
				aup_timeout_in_days = types.Int64Value(int64(*d.AupTimeoutInDays))
			}
			if d.AuthRequired != nil {
				auth_required = types.BoolValue(*d.AuthRequired)
			}
			if d.CautionEnabled != nil {
				caution_enabled = types.BoolValue(*d.CautionEnabled)
			}
			if d.DnBandwidth.Value() != nil {
				dn_bandwidth = types.Float64Value(*d.DnBandwidth.Value())
			}
			if d.IdleTimeInMinutes != nil {
				idle_time_in_minutes = types.Int64Value(int64(*d.IdleTimeInMinutes))
			}
			if d.Name != nil {
				name = types.StringValue(*d.Name)
			}
			if d.OfwEnabled != nil {
				ofw_enabled = types.BoolValue(*d.OfwEnabled)
			}
			if d.SurrogateIP != nil {
				surrogate_ip = types.BoolValue(*d.SurrogateIP)
			}
			if d.SurrogateIPEnforcedForKnownBrowsers != nil {
				surrogate_ip_enforced_for_known_browsers = types.BoolValue(*d.SurrogateIPEnforcedForKnownBrowsers)
			}
			if d.SurrogateRefreshTimeInMinutes != nil {
				surrogate_refresh_time_in_minutes = types.Int64Value(int64(*d.SurrogateRefreshTimeInMinutes))
			}
			if d.UpBandwidth.Value() != nil {
				up_bandwidth = types.Float64Value(*d.UpBandwidth.Value())
			}

			data_map_value := map[string]attr.Value{
				"aup_block_internet_until_accepted": aup_block_internet_until_accepted,
				"aup_enabled":                       aup_enabled,
				"aup_force_ssl_inspection":          aup_force_ssl_inspection,
				"aup_timeout_in_days":               aup_timeout_in_days,
				"auth_required":                     auth_required,
				"caution_enabled":                   caution_enabled,
				"dn_bandwidth":                      dn_bandwidth,
				"idle_time_in_minutes":              idle_time_in_minutes,
				"name":                              name,
				"ofw_enabled":                       ofw_enabled,
				"surrogate_ip":                      surrogate_ip,
				"surrogate_ip_enforced_for_known_browsers": surrogate_ip_enforced_for_known_browsers,
				"surrogate_refresh_time_in_minutes":        surrogate_refresh_time_in_minutes,
				"up_bandwidth":                             up_bandwidth,
			}

			data, e := NewSubLocationsValue(SubLocationsValue{}.AttributeTypes(ctx), data_map_value)
			diags.Append(e...)
			data_list = append(data_list, data)
		}
	}
	data_list_type := SubLocationsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func tunnelProviderZscalerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, t *models.TunnelProviderOptions) (basetypes.ObjectValue, bool) {

	var aup_block_internet_until_accepted basetypes.BoolValue
	var aup_enabled basetypes.BoolValue
	var aup_force_ssl_inspection basetypes.BoolValue
	var aup_timeout_in_days basetypes.Int64Value
	var auth_required basetypes.BoolValue
	var caution_enabled basetypes.BoolValue
	var dn_bandwidth basetypes.Float64Value
	var idle_time_in_minutes basetypes.Int64Value
	var ofw_enabled basetypes.BoolValue
	var sub_locations basetypes.ListValue = types.ListNull(SubLocationsValue{}.Type(ctx))
	var surrogate_ip basetypes.BoolValue
	var surrogate_ip_enforced_for_known_browsers basetypes.BoolValue
	var surrogate_refresh_time_in_minutes basetypes.Int64Value
	var up_bandwidth basetypes.Float64Value
	var xff_forward_enabled basetypes.BoolValue
	configured := false

	if t != nil && t.Zscaler != nil && t.Zscaler.AupBlockInternetUntilAccepted != nil {
		aup_block_internet_until_accepted = types.BoolValue(*t.Zscaler.AupBlockInternetUntilAccepted)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupEnabled != nil {
		aup_enabled = types.BoolValue(*t.Zscaler.AupEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupForceSslInspection != nil {
		aup_force_ssl_inspection = types.BoolValue(*t.Zscaler.AupForceSslInspection)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupTimeoutInDays != nil {
		aup_timeout_in_days = types.Int64Value(int64(*t.Zscaler.AupTimeoutInDays))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AuthRequired != nil {
		auth_required = types.BoolValue(*t.Zscaler.AuthRequired)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.CautionEnabled != nil {
		caution_enabled = types.BoolValue(*t.Zscaler.CautionEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.DnBandwidth.Value() != nil {
		dn_bandwidth = types.Float64Value(*t.Zscaler.DnBandwidth.Value())
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.IdleTimeInMinutes != nil {
		idle_time_in_minutes = types.Int64Value(int64(*t.Zscaler.IdleTimeInMinutes))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.OfwEnabled != nil {
		ofw_enabled = types.BoolValue(*t.Zscaler.OfwEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SubLocations != nil {
		sub_locations = tunnelProviderZscalerSubLocationSdkToTerraform(ctx, diags, t)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateIP != nil {
		surrogate_ip = types.BoolValue(*t.Zscaler.SurrogateIP)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateIPEnforcedForKnownBrowsers != nil {
		surrogate_ip_enforced_for_known_browsers = types.BoolValue(*t.Zscaler.SurrogateIPEnforcedForKnownBrowsers)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateRefreshTimeInMinutes != nil {
		surrogate_refresh_time_in_minutes = types.Int64Value(int64(*t.Zscaler.SurrogateRefreshTimeInMinutes))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.UpBandwidth.Value() != nil {
		up_bandwidth = types.Float64Value(*t.Zscaler.UpBandwidth.Value())
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.XffForwardEnabled != nil {
		xff_forward_enabled = types.BoolValue(*t.Zscaler.XffForwardEnabled)
		configured = true
	}

	r_attr_value := map[string]attr.Value{
		"aup_block_internet_until_accepted":        aup_block_internet_until_accepted,
		"aup_enabled":                              aup_enabled,
		"aup_force_ssl_inspection":                 aup_force_ssl_inspection,
		"aup_timeout_in_days":                      aup_timeout_in_days,
		"auth_required":                            auth_required,
		"caution_enabled":                          caution_enabled,
		"dn_bandwidth":                             dn_bandwidth,
		"idle_time_in_minutes":                     idle_time_in_minutes,
		"ofw_enabled":                              ofw_enabled,
		"sub_locations":                            sub_locations,
		"surrogate_ip":                             surrogate_ip,
		"surrogate_ip_enforced_for_known_browsers": surrogate_ip_enforced_for_known_browsers,
		"surrogate_refresh_time_in_minutes":        surrogate_refresh_time_in_minutes,
		"up_bandwidth":                             up_bandwidth,
		"xff_forward_enabled":                      xff_forward_enabled,
	}
	r, e := basetypes.NewObjectValue(ZscalerValue{}.AttributeTypes(ctx), r_attr_value)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TunnelProviderOptions) (TunnelProviderOptionsValue, bool) {
	var jse basetypes.ObjectValue = types.ObjectNull(JseValue{}.AttributeTypes(ctx))
	var zscaler basetypes.ObjectValue = types.ObjectNull(ZscalerValue{}.AttributeTypes(ctx))
	configured := false

	if d != nil && d.Jse != nil {
		if jse_tmp, ok := tunnelProviderJseSdkToTerraform(ctx, diags, d); ok {
			jse = jse_tmp
			configured = true
		}
	}
	if d != nil && d.Zscaler != nil {
		if zscaler_tmp, ok := tunnelProviderZscalerSdkToTerraform(ctx, diags, d); ok {
			zscaler = zscaler_tmp
			configured = true
		}
	}

	data_map_value := map[string]attr.Value{
		"jse":     jse,
		"zscaler": zscaler,
	}
	data, e := NewTunnelProviderOptionsValue(TunnelProviderOptionsValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data, configured
}
