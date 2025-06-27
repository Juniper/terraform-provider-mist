package resource_org_wlan

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radsecServersSkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadsecServer) basetypes.ListValue {
	var dataList []ServersValue
	for _, d := range l {
		var host basetypes.StringValue
		var port basetypes.Int64Value

		if d.Host != nil {
			host = types.StringValue(*d.Host)
		}
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}

		dataMapValue := map[string]attr.Value{
			"host": host,
			"port": port,
		}
		data, e := NewServersValue(ServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, ServersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}
func radsecSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Radsec) RadsecValue {
	var coaEnabled basetypes.BoolValue
	var enabled basetypes.BoolValue
	var idleTimeout basetypes.StringValue
	var mxclusterIds = basetypes.NewListNull(types.StringType)
	var proxyHosts = basetypes.NewListNull(types.StringType)
	var serverName basetypes.StringValue
	var servers = types.ListValueMust(ServersValue{}.Type(ctx), []attr.Value{})
	var useMxedge basetypes.BoolValue
	var useSiteMxedge basetypes.BoolValue

	if d != nil && d.CoaEnabled != nil {
		coaEnabled = types.BoolValue(*d.CoaEnabled)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d != nil && d.IdleTimeout != nil {
		idleTimeout = mistutils.RadsecIdleTimeoutAsString(d.IdleTimeout)
	}
	if d != nil && d.MxclusterIds != nil {
		mxclusterIds = mistutils.ListOfUuidSdkToTerraform(d.MxclusterIds)
	}
	if d != nil && d.ProxyHosts != nil {
		proxyHosts = mistutils.ListOfStringSdkToTerraform(d.ProxyHosts)
	}
	if d != nil && d.ServerName != nil {
		serverName = types.StringValue(*d.ServerName)
	}
	if d != nil && d.Servers != nil {
		servers = radsecServersSkToTerraform(ctx, diags, d.Servers)
	}
	if d != nil && d.UseMxedge != nil {
		useMxedge = types.BoolValue(*d.UseMxedge)
	}
	if d != nil && d.UseSiteMxedge != nil {
		useSiteMxedge = types.BoolValue(*d.UseSiteMxedge)
	}

	dataMapValue := map[string]attr.Value{
		"coa_enabled":     coaEnabled,
		"enabled":         enabled,
		"idle_timeout":    idleTimeout,
		"mxcluster_ids":   mxclusterIds,
		"proxy_hosts":     proxyHosts,
		"server_name":     serverName,
		"servers":         servers,
		"use_mxedge":      useMxedge,
		"use_site_mxedge": useSiteMxedge,
	}
	data, e := NewRadsecValue(RadsecValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
