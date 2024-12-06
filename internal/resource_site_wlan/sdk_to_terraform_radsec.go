package resource_site_wlan

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radsecServersSkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadsecServer) basetypes.ListValue {
	var data_list = []ServersValue{}
	for _, d := range l {
		var host basetypes.StringValue
		var port basetypes.Int64Value

		if d.Host != nil {
			host = types.StringValue(*d.Host)
		}
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}

		data_map_attr_type := ServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"host": host,
			"port": port,
		}
		data, e := NewServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, ServersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}
func radsecSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Radsec) RadsecValue {
	var coa_enabled basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idle_timeout basetypes.Int64Value
	var mxcluster_ids basetypes.ListValue = mist_transform.ListOfUuidSdkToTerraformEmpty(ctx)
	var proxy_hosts basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var server_name basetypes.StringValue
	var servers basetypes.ListValue = types.ListValueMust(ServersValue{}.Type(ctx), []attr.Value{})
	var use_mxedge basetypes.BoolValue
	var use_site_mxedge basetypes.BoolValue

	if d != nil && d.CoaEnabled != nil {
		coa_enabled = types.BoolValue(*d.CoaEnabled)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.IdleTimeout != nil {
		idle_timeout = types.Int64Value(int64(*d.IdleTimeout))
	}
	if d != nil && d.MxclusterIds != nil {
		mxcluster_ids = mist_transform.ListOfUuidSdkToTerraform(ctx, d.MxclusterIds)
	}
	if d != nil && d.ProxyHosts != nil {
		proxy_hosts = mist_transform.ListOfStringSdkToTerraform(ctx, d.ProxyHosts)
	}
	if d != nil && d.ServerName != nil {
		server_name = types.StringValue(*d.ServerName)
	}
	if d != nil && d.Servers != nil {
		servers = radsecServersSkToTerraform(ctx, diags, d.Servers)
	}
	if d != nil && d.UseMxedge != nil {
		use_mxedge = types.BoolValue(*d.UseMxedge)
	}
	if d != nil && d.UseSiteMxedge != nil {
		use_site_mxedge = types.BoolValue(*d.UseSiteMxedge)
	}

	data_map_attr_type := RadsecValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"coa_enabled":     coa_enabled,
		"enabled":         enabled,
		"idle_timeout":    idle_timeout,
		"mxcluster_ids":   mxcluster_ids,
		"proxy_hosts":     proxy_hosts,
		"server_name":     server_name,
		"servers":         servers,
		"use_mxedge":      use_mxedge,
		"use_site_mxedge": use_site_mxedge,
	}
	data, e := NewRadsecValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
