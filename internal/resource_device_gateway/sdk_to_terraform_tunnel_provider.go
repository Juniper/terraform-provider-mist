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

	var orgName basetypes.StringValue
	var numUsers basetypes.Int64Value
	configured := false

	if d != nil && d.Jse != nil && d.Jse.OrgName != nil {
		orgName = types.StringValue(*d.Jse.OrgName)
		configured = true
	}
	if d != nil && d.Jse != nil && d.Jse.NumUsers != nil {
		numUsers = types.Int64Value(int64(*d.Jse.NumUsers))
		configured = true
	}

	rAttrValue := map[string]attr.Value{
		"org_name":  orgName,
		"num_users": numUsers,
	}
	r, e := basetypes.NewObjectValue(JseValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderPrismaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TunnelProviderOptions) (basetypes.ObjectValue, bool) {

	var serviceAccountName basetypes.StringValue
	configured := false

	if d != nil && d.Prisma != nil && d.Prisma.ServiceAccountName != nil {
		serviceAccountName = types.StringValue(*d.Prisma.ServiceAccountName)
		configured = true
	}

	rAttrValue := map[string]attr.Value{
		"service_account_name": serviceAccountName,
	}
	r, e := basetypes.NewObjectValue(PrismaValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderZscalerSubLocationSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, t *models.TunnelProviderOptions) basetypes.ListValue {
	var dataList []SubLocationsValue
	if t != nil && t.Zscaler != nil && t.Zscaler.SubLocations != nil {
		for _, d := range t.Zscaler.SubLocations {
			var aupBlockInternetUntilAccepted basetypes.BoolValue
			var aupEnabled basetypes.BoolValue
			var aupForceSslInspection basetypes.BoolValue
			var aupTimeoutInDays basetypes.Int64Value
			var authRequired basetypes.BoolValue
			var cautionEnabled basetypes.BoolValue
			var dnBandwidth basetypes.Float64Value
			var idleTimeInMinutes basetypes.Int64Value
			var name basetypes.StringValue
			var ofwEnabled basetypes.BoolValue
			var surrogateIp basetypes.BoolValue
			var surrogateIpEnforcedForKnownBrowsers basetypes.BoolValue
			var surrogateRefreshTimeInMinutes basetypes.Int64Value
			var upBandwidth basetypes.Float64Value

			if d.AupBlockInternetUntilAccepted != nil {
				aupBlockInternetUntilAccepted = types.BoolValue(*d.AupBlockInternetUntilAccepted)
			}
			if d.AupEnabled != nil {
				aupEnabled = types.BoolValue(*d.AupEnabled)
			}
			if d.AupForceSslInspection != nil {
				aupForceSslInspection = types.BoolValue(*d.AupForceSslInspection)
			}
			if d.AupTimeoutInDays != nil {
				aupTimeoutInDays = types.Int64Value(int64(*d.AupTimeoutInDays))
			}
			if d.AuthRequired != nil {
				authRequired = types.BoolValue(*d.AuthRequired)
			}
			if d.CautionEnabled != nil {
				cautionEnabled = types.BoolValue(*d.CautionEnabled)
			}
			if d.DnBandwidth.Value() != nil {
				dnBandwidth = types.Float64Value(*d.DnBandwidth.Value())
			}
			if d.IdleTimeInMinutes != nil {
				idleTimeInMinutes = types.Int64Value(int64(*d.IdleTimeInMinutes))
			}
			if d.Name != nil {
				name = types.StringValue(*d.Name)
			}
			if d.OfwEnabled != nil {
				ofwEnabled = types.BoolValue(*d.OfwEnabled)
			}
			if d.SurrogateIP != nil {
				surrogateIp = types.BoolValue(*d.SurrogateIP)
			}
			if d.SurrogateIPEnforcedForKnownBrowsers != nil {
				surrogateIpEnforcedForKnownBrowsers = types.BoolValue(*d.SurrogateIPEnforcedForKnownBrowsers)
			}
			if d.SurrogateRefreshTimeInMinutes != nil {
				surrogateRefreshTimeInMinutes = types.Int64Value(int64(*d.SurrogateRefreshTimeInMinutes))
			}
			if d.UpBandwidth.Value() != nil {
				upBandwidth = types.Float64Value(*d.UpBandwidth.Value())
			}

			dataMapValue := map[string]attr.Value{
				"aup_block_internet_until_accepted": aupBlockInternetUntilAccepted,
				"aup_enabled":                       aupEnabled,
				"aup_force_ssl_inspection":          aupForceSslInspection,
				"aup_timeout_in_days":               aupTimeoutInDays,
				"auth_required":                     authRequired,
				"caution_enabled":                   cautionEnabled,
				"dn_bandwidth":                      dnBandwidth,
				"idle_time_in_minutes":              idleTimeInMinutes,
				"name":                              name,
				"ofw_enabled":                       ofwEnabled,
				"surrogate_ip":                      surrogateIp,
				"surrogate_ip_enforced_for_known_browsers": surrogateIpEnforcedForKnownBrowsers,
				"surrogate_refresh_time_in_minutes":        surrogateRefreshTimeInMinutes,
				"up_bandwidth":                             upBandwidth,
			}

			data, e := NewSubLocationsValue(SubLocationsValue{}.AttributeTypes(ctx), dataMapValue)
			diags.Append(e...)
			dataList = append(dataList, data)
		}
	}
	datalistType := SubLocationsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, datalistType, dataList)
	diags.Append(e...)
	return r
}
func tunnelProviderZscalerSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, t *models.TunnelProviderOptions) (basetypes.ObjectValue, bool) {

	var aupBlockInternetUntilAccepted basetypes.BoolValue
	var aupEnabled basetypes.BoolValue
	var aupForceSslInspection basetypes.BoolValue
	var aupTimeoutInDays basetypes.Int64Value
	var authRequired basetypes.BoolValue
	var cautionEnabled basetypes.BoolValue
	var dnBandwidth basetypes.Float64Value
	var idleTimeInMinutes basetypes.Int64Value
	var ofwEnabled basetypes.BoolValue
	var subLocations = types.ListNull(SubLocationsValue{}.Type(ctx))
	var surrogateIp basetypes.BoolValue
	var surrogateIpEnforcedForKnownBrowsers basetypes.BoolValue
	var surrogateRefreshTimeInMinutes basetypes.Int64Value
	var upBandwidth basetypes.Float64Value
	var xffForwardEnabled basetypes.BoolValue
	configured := false

	if t != nil && t.Zscaler != nil && t.Zscaler.AupBlockInternetUntilAccepted != nil {
		aupBlockInternetUntilAccepted = types.BoolValue(*t.Zscaler.AupBlockInternetUntilAccepted)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupEnabled != nil {
		aupEnabled = types.BoolValue(*t.Zscaler.AupEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupForceSslInspection != nil {
		aupForceSslInspection = types.BoolValue(*t.Zscaler.AupForceSslInspection)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupTimeoutInDays != nil {
		aupTimeoutInDays = types.Int64Value(int64(*t.Zscaler.AupTimeoutInDays))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AuthRequired != nil {
		authRequired = types.BoolValue(*t.Zscaler.AuthRequired)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.CautionEnabled != nil {
		cautionEnabled = types.BoolValue(*t.Zscaler.CautionEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.DnBandwidth.Value() != nil {
		dnBandwidth = types.Float64Value(*t.Zscaler.DnBandwidth.Value())
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.IdleTimeInMinutes != nil {
		idleTimeInMinutes = types.Int64Value(int64(*t.Zscaler.IdleTimeInMinutes))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.OfwEnabled != nil {
		ofwEnabled = types.BoolValue(*t.Zscaler.OfwEnabled)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SubLocations != nil {
		subLocations = tunnelProviderZscalerSubLocationSdkToTerraform(ctx, diags, t)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateIP != nil {
		surrogateIp = types.BoolValue(*t.Zscaler.SurrogateIP)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateIPEnforcedForKnownBrowsers != nil {
		surrogateIpEnforcedForKnownBrowsers = types.BoolValue(*t.Zscaler.SurrogateIPEnforcedForKnownBrowsers)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.SurrogateRefreshTimeInMinutes != nil {
		surrogateRefreshTimeInMinutes = types.Int64Value(int64(*t.Zscaler.SurrogateRefreshTimeInMinutes))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.UpBandwidth.Value() != nil {
		upBandwidth = types.Float64Value(*t.Zscaler.UpBandwidth.Value())
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.XffForwardEnabled != nil {
		xffForwardEnabled = types.BoolValue(*t.Zscaler.XffForwardEnabled)
		configured = true
	}

	rAttrValue := map[string]attr.Value{
		"aup_block_internet_until_accepted":        aupBlockInternetUntilAccepted,
		"aup_enabled":                              aupEnabled,
		"aup_force_ssl_inspection":                 aupForceSslInspection,
		"aup_timeout_in_days":                      aupTimeoutInDays,
		"auth_required":                            authRequired,
		"caution_enabled":                          cautionEnabled,
		"dn_bandwidth":                             dnBandwidth,
		"idle_time_in_minutes":                     idleTimeInMinutes,
		"ofw_enabled":                              ofwEnabled,
		"sub_locations":                            subLocations,
		"surrogate_ip":                             surrogateIp,
		"surrogate_ip_enforced_for_known_browsers": surrogateIpEnforcedForKnownBrowsers,
		"surrogate_refresh_time_in_minutes":        surrogateRefreshTimeInMinutes,
		"up_bandwidth":                             upBandwidth,
		"xff_forward_enabled":                      xffForwardEnabled,
	}
	r, e := basetypes.NewObjectValue(ZscalerValue{}.AttributeTypes(ctx), rAttrValue)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TunnelProviderOptions) (TunnelProviderOptionsValue, bool) {
	var jse = types.ObjectNull(JseValue{}.AttributeTypes(ctx))
	var prisma = types.ObjectNull(PrismaValue{}.AttributeTypes(ctx))
	var zscaler = types.ObjectNull(ZscalerValue{}.AttributeTypes(ctx))
	configured := false

	if d != nil && d.Jse != nil {
		if jseTmp, ok := tunnelProviderJseSdkToTerraform(ctx, diags, d); ok {
			jse = jseTmp
			configured = true
		}
	}
	if d != nil && d.Prisma != nil {
		if prismaTmp, ok := tunnelProviderPrismaSdkToTerraform(ctx, diags, d); ok {
			prisma = prismaTmp
			configured = true
		}
	}
	if d != nil && d.Zscaler != nil {
		if zscalerTmp, ok := tunnelProviderZscalerSdkToTerraform(ctx, diags, d); ok {
			zscaler = zscalerTmp
			configured = true
		}
	}

	dataMapValue := map[string]attr.Value{
		"jse":     jse,
		"prisma":  prisma,
		"zscaler": zscaler,
	}
	data, e := NewTunnelProviderOptionsValue(TunnelProviderOptionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data, configured
}
