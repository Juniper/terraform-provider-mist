package resource_org_deviceprofile_gateway

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tunnelProviderJseSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TunnelProviderOptions) (basetypes.ObjectValue, bool) {

	var name basetypes.StringValue
	var num_users basetypes.Int64Value
	configured := false

	if d != nil && d.Jse != nil && d.Jse.Name != nil {
		name = types.StringValue(*d.Jse.Name)
		configured = true
	}
	if d != nil && d.Jse != nil && d.Jse.NumUsers != nil {
		num_users = types.Int64Value(int64(*d.Jse.NumUsers))
		configured = true
	}

	r_attr_value := map[string]attr.Value{
		"name":      name,
		"num_users": num_users,
	}
	r, e := basetypes.NewObjectValue(JseValue{}.AttributeTypes(ctx), r_attr_value)
	diags.Append(e...)
	return r, configured
}

func tunnelProviderZscalerSubLocationSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, t *models.TunnelProviderOptions) basetypes.ListValue {
	var data_list = []SubLocationsValue{}
	if t != nil && t.Zscaler != nil && t.Zscaler.SubLocations != nil {
		for _, v := range t.Zscaler.SubLocations {
			var aup_acceptance_required basetypes.BoolValue
			var aup_expire basetypes.Int64Value
			var aup_ssl_proxy basetypes.BoolValue
			var download_mbps basetypes.Int64Value
			var enable_aup basetypes.BoolValue
			var enable_caution basetypes.BoolValue
			var enforce_authentication basetypes.BoolValue
			var subnets basetypes.ListValue = mist_transform.ListOfStringSdkToTerraform(ctx, v.Subnets)
			var upload_mbps basetypes.Int64Value

			if v.AupAcceptanceRequired != nil {
				aup_acceptance_required = types.BoolValue(*v.AupAcceptanceRequired)
			}
			if v.AupExpire != nil {
				aup_expire = types.Int64Value(int64(*v.AupExpire))
			}
			if v.AupSslProxy != nil {
				aup_ssl_proxy = types.BoolValue(*v.AupSslProxy)
			}
			if v.DownloadMbps != nil {
				download_mbps = types.Int64Value(int64(*v.DownloadMbps))
			}
			if v.EnableAup != nil {
				enable_aup = types.BoolValue(*v.EnableAup)
			}
			if v.EnableCaution != nil {
				enable_caution = types.BoolValue(*v.EnableCaution)
			}
			if v.EnforceAuthentication != nil {
				enforce_authentication = types.BoolValue(*v.EnforceAuthentication)
			}
			if v.UploadMbps != nil {
				upload_mbps = types.Int64Value(int64(*v.UploadMbps))
			}

			data_map_value := map[string]attr.Value{
				"aup_acceptance_required": aup_acceptance_required,
				"aup_expire":              aup_expire,
				"aup_ssl_proxy":           aup_ssl_proxy,
				"download_mbps":           download_mbps,
				"enable_aup":              enable_aup,
				"enable_caution":          enable_caution,
				"enforce_authentication":  enforce_authentication,
				"subnets":                 subnets,
				"upload_mbps":             upload_mbps,
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

	var aup_acceptance_required basetypes.BoolValue = types.BoolValue(true)
	var aup_expire basetypes.Int64Value = types.Int64Value(1)
	var aup_ssl_proxy basetypes.BoolValue = types.BoolValue(false)
	var download_mbps basetypes.Int64Value
	var enable_aup basetypes.BoolValue = types.BoolValue(false)
	var enable_caution basetypes.BoolValue = types.BoolValue(false)
	var enforce_authentication basetypes.BoolValue = types.BoolValue(false)
	var name basetypes.StringValue
	var sub_locations basetypes.ListValue = tunnelProviderZscalerSubLocationSdkToTerraform(ctx, diags, t)
	var upload_mbps basetypes.Int64Value
	var use_xff basetypes.BoolValue
	configured := false

	if t != nil && t.Zscaler != nil && t.Zscaler.AupAcceptanceRequired != nil {
		aup_acceptance_required = types.BoolValue(*t.Zscaler.AupAcceptanceRequired)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupExpire != nil {
		aup_expire = types.Int64Value(int64(*t.Zscaler.AupExpire))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.AupSslProxy != nil {
		aup_ssl_proxy = types.BoolValue(*t.Zscaler.AupSslProxy)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.DownloadMbps != nil {
		download_mbps = types.Int64Value(int64(*t.Zscaler.DownloadMbps))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.EnableAup != nil {
		enable_aup = types.BoolValue(*t.Zscaler.EnableAup)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.EnableCaution != nil {
		enable_caution = types.BoolValue(*t.Zscaler.EnableCaution)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.EnforceAuthentication != nil {
		enforce_authentication = types.BoolValue(*t.Zscaler.EnforceAuthentication)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.Name != nil {
		name = types.StringValue(*t.Zscaler.Name)
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.UploadMbps != nil {
		upload_mbps = types.Int64Value(int64(*t.Zscaler.UploadMbps))
		configured = true
	}
	if t != nil && t.Zscaler != nil && t.Zscaler.UseXff != nil {
		use_xff = types.BoolValue(*t.Zscaler.UseXff)
		configured = true
	}

	r_attr_value := map[string]attr.Value{
		"aup_acceptance_required": aup_acceptance_required,
		"aup_expire":              aup_expire,
		"aup_ssl_proxy":           aup_ssl_proxy,
		"download_mbps":           download_mbps,
		"enable_aup":              enable_aup,
		"enable_caution":          enable_caution,
		"enforce_authentication":  enforce_authentication,
		"name":                    name,
		"sub_locations":           sub_locations,
		"upload_mbps":             upload_mbps,
		"use_xff":                 use_xff,
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
