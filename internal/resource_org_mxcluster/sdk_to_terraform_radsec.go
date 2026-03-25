package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func radsecSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxclusterRadsec) RadsecValue {

	var acctServers = types.ListNull(AcctServersValue{}.Type(ctx))
	var authServers = types.ListNull(AuthServersValue{}.Type(ctx))
	var enabled = types.BoolNull()
	var matchSsid = types.BoolNull()
	var nasIpSource = types.StringNull()
	var proxyHosts = types.ListNull(types.StringType)
	var serverSelection = types.StringNull()
	var srcIpSource = types.StringNull()

	if d.AcctServers != nil {
		var dataList []attr.Value

		for _, item := range d.AcctServers {
			var host = types.StringNull()
			var port = types.Int64Null()
			var secret = types.StringNull()
			var ssids = types.ListNull(types.StringType)

			if item.Host != nil {
				host = types.StringValue(*item.Host)
			}
			if item.Port != nil {
				port = types.Int64Value(int64(*item.Port))
			}
			if item.Secret != nil {
				secret = types.StringValue(*item.Secret)
			}
			if item.Ssids != nil {
				ssids = mistutils.ListOfStringSdkToTerraform(item.Ssids)
			}

			itemAttrType := AcctServersValue{}.AttributeTypes(ctx)
			itemValue := map[string]attr.Value{
				"host":   host,
				"port":   port,
				"secret": secret,
				"ssids":  ssids,
			}
			item_o, e := NewAcctServersValue(itemAttrType, itemValue)
			diags.Append(e...)

			dataList = append(dataList, item_o)
		}

		r_list, e := types.ListValueFrom(ctx, AcctServersValue{}.Type(ctx), dataList)
		diags.Append(e...)
		acctServers = r_list
	}
	if d.AuthServers != nil {
		var dataList []attr.Value

		for _, item := range d.AuthServers {
			var host = types.StringNull()
			var inbandStatusCheck = types.BoolNull()
			var inbandStatusInterval = types.Int64Null()
			var keywrapEnabled = types.BoolNull()
			var keywrapFormat = types.StringNull()
			var keywrapKek = types.StringNull()
			var keywrapMack = types.StringNull()
			var port = types.Int64Null()
			var retry = types.Int64Null()
			var secret = types.StringNull()
			var ssids = types.ListNull(types.StringType)
			var timeout = types.Int64Null()

			if item.Host != nil {
				host = types.StringValue(*item.Host)
			}
			if item.InbandStatusCheck != nil {
				inbandStatusCheck = types.BoolValue(*item.InbandStatusCheck)
			}
			if item.InbandStatusInterval != nil {
				inbandStatusInterval = types.Int64Value(int64(*item.InbandStatusInterval))
			}
			if item.KeywrapEnabled != nil {
				keywrapEnabled = types.BoolValue(*item.KeywrapEnabled)
			}
			if item.KeywrapFormat.Value() != nil {
				keywrapFormat = types.StringValue(string(*item.KeywrapFormat.Value()))
			}
			if item.KeywrapKek != nil {
				keywrapKek = types.StringValue(*item.KeywrapKek)
			}
			if item.KeywrapMack != nil {
				keywrapMack = types.StringValue(*item.KeywrapMack)
			}
			if item.Port != nil {
				port = types.Int64Value(int64(*item.Port))
			}
			if item.Retry != nil {
				retry = types.Int64Value(int64(*item.Retry))
			}
			if item.Secret != nil {
				secret = types.StringValue(*item.Secret)
			}
			if item.Ssids != nil {
				ssids = mistutils.ListOfStringSdkToTerraform(item.Ssids)
			}
			if item.Timeout != nil {
				timeout = types.Int64Value(int64(*item.Timeout))
			}

			itemAttrType := AuthServersValue{}.AttributeTypes(ctx)
			itemValue := map[string]attr.Value{
				"host":                   host,
				"inband_status_check":    inbandStatusCheck,
				"inband_status_interval": inbandStatusInterval,
				"keywrap_enabled":        keywrapEnabled,
				"keywrap_format":         keywrapFormat,
				"keywrap_kek":            keywrapKek,
				"keywrap_mack":           keywrapMack,
				"port":                   port,
				"retry":                  retry,
				"secret":                 secret,
				"ssids":                  ssids,
				"timeout":                timeout,
			}
			item_o, e := NewAuthServersValue(itemAttrType, itemValue)
			diags.Append(e...)

			dataList = append(dataList, item_o)
		}

		r_list, e := types.ListValueFrom(ctx, AuthServersValue{}.Type(ctx), dataList)
		diags.Append(e...)
		authServers = r_list
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.MatchSsid != nil {
		matchSsid = types.BoolValue(*d.MatchSsid)
	}
	if d.NasIpSource != nil {
		nasIpSource = types.StringValue(string(*d.NasIpSource))
	}
	if d.ProxyHosts != nil {
		proxyHosts = mistutils.ListOfStringSdkToTerraform(d.ProxyHosts)
	}
	if d.ServerSelection != nil {
		serverSelection = types.StringValue(string(*d.ServerSelection))
	}
	if d.SrcIpSource != nil {
		srcIpSource = types.StringValue(string(*d.SrcIpSource))
	}

	data_map_attr_type := RadsecValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"acct_servers":     acctServers,
		"auth_servers":     authServers,
		"enabled":          enabled,
		"match_ssid":       matchSsid,
		"nas_ip_source":    nasIpSource,
		"proxy_hosts":      proxyHosts,
		"server_selection": serverSelection,
		"src_ip_source":    srcIpSource,
	}
	data, e := NewRadsecValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
