package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

/*
************************
*
* DYNAMIC VLAN
*
/************************
*/
func dynamicVlanPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApPortConfigDynamicVlan) basetypes.ObjectValue {
	var defaultVlanId basetypes.Int64Value
	var enabled basetypes.BoolValue
	var typeDynamicVlan basetypes.StringValue
	var vlans = types.MapNull(types.StringType)

	if d.DefaultVlanId != nil {
		defaultVlanId = types.Int64Value(int64(*d.DefaultVlanId))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Type != nil {
		typeDynamicVlan = types.StringValue(*d.Type)
	}
	if d != nil && d.Vlans != nil {
		vlansAttr := make(map[string]attr.Value)
		for k, v := range d.Vlans {
			vlansAttr[k] = types.StringValue(v)
		}
		vlans = types.MapValueMust(basetypes.StringType{}, vlansAttr)
	}

	data_map_value := map[string]attr.Value{
		"default_vlan_id": defaultVlanId,
		"enabled":         enabled,
		"type":            typeDynamicVlan,
		"vlans":           vlans,
	}
	data, e := basetypes.NewObjectValue(DynamicVlanValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

/*
************************
*
* MIST NAC
*
/************************
*/
func mistNacPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanMistNac) basetypes.ObjectValue {
	var acctInterimInterval basetypes.Int64Value
	var authServersRetries basetypes.Int64Value
	var authServersTimeout basetypes.Int64Value
	var coaEnabled basetypes.BoolValue
	var coaPort basetypes.Int64Value
	var enabled basetypes.BoolValue
	var fastDot1xTimers basetypes.BoolValue
	var network basetypes.StringValue
	var sourceIp basetypes.StringValue

	if d.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*d.AcctInterimInterval))
	}
	if d.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*d.AuthServersRetries))
	}
	if d.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}
	if d.CoaEnabled != nil {
		coaEnabled = types.BoolValue(*d.CoaEnabled)
	}
	if d.CoaPort != nil {
		coaPort = types.Int64Value(int64(*d.CoaPort))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.FastDot1xTimers != nil {
		fastDot1xTimers = types.BoolValue(*d.FastDot1xTimers)
	}
	if d.Network.Value() != nil {
		network = types.StringValue(*d.Network.Value())
	}
	if d.SourceIp.Value() != nil {
		sourceIp = types.StringValue(*d.SourceIp.Value())
	}

	data_map_value := map[string]attr.Value{
		"acct_interim_interval": acctInterimInterval,
		"auth_servers_retries":  authServersRetries,
		"auth_servers_timeout":  authServersTimeout,
		"coa_enabled":           coaEnabled,
		"coa_port":              coaPort,
		"enabled":               enabled,
		"fast_dot1x_timers":     fastDot1xTimers,
		"network":               network,
		"source_ip":             sourceIp,
	}
	data, e := basetypes.NewObjectValue(MistNacValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

/*
************************
*
* RADIUS
*
/************************
*/
func radiusServersAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadiusAcctServer) basetypes.ListValue {
	var acctValueList []attr.Value
	for _, d := range l {
		var host = types.StringValue(d.Host)
		var keywrapEnabled basetypes.BoolValue
		var keywrapFormat basetypes.StringValue
		var keywrapKek basetypes.StringValue
		var keywrapMack basetypes.StringValue
		var port basetypes.StringValue
		var secret = types.StringValue(d.Secret)

		if d.KeywrapEnabled != nil {
			keywrapEnabled = types.BoolValue(*d.KeywrapEnabled)
		}
		if d.KeywrapFormat != nil {
			keywrapFormat = types.StringValue(string(*d.KeywrapFormat))
		}
		if d.KeywrapKek != nil {
			keywrapKek = types.StringValue(*d.KeywrapKek)
		}
		if d.KeywrapMack != nil {
			keywrapMack = types.StringValue(*d.KeywrapMack)
		}
		if d.Port != nil {
			port = mistutils.RadiusAcctPortAsString(d.Port)
		}

		dataMapValue := map[string]attr.Value{
			"host":            host,
			"keywrap_enabled": keywrapEnabled,
			"keywrap_format":  keywrapFormat,
			"keywrap_kek":     keywrapKek,
			"keywrap_mack":    keywrapMack,
			"port":            port,
			"secret":          secret,
		}
		data, e := NewAcctServersValue(AcctServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		acctValueList = append(acctValueList, data)
	}

	acctStateListType := AcctServersValue{}.Type(ctx)
	acctStateList, e := types.ListValueFrom(ctx, acctStateListType, acctValueList)
	diags.Append(e...)

	return acctStateList
}

func radiusServersAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadiusAuthServer) basetypes.ListValue {
	var authValueList []attr.Value
	for _, d := range l {
		var host basetypes.StringValue
		var keywrapEnabled basetypes.BoolValue
		var keywrapFormat basetypes.StringValue
		var keywrapKek basetypes.StringValue
		var keywrapMack basetypes.StringValue
		var port basetypes.StringValue
		var requireMessageAuthenticator basetypes.BoolValue
		var secret basetypes.StringValue

		host = types.StringValue(d.Host)
		if d.KeywrapEnabled != nil {
			keywrapEnabled = types.BoolValue(*d.KeywrapEnabled)
		}
		if d.KeywrapFormat != nil {
			keywrapFormat = types.StringValue(string(*d.KeywrapFormat))
		}
		if d.KeywrapKek != nil {
			keywrapKek = types.StringValue(*d.KeywrapKek)
		}
		if d.KeywrapMack != nil {
			keywrapMack = types.StringValue(*d.KeywrapMack)
		}
		if d.Port != nil {
			port = mistutils.RadiusAuthPortAsString(d.Port)
		}
		if d.RequireMessageAuthenticator != nil {
			requireMessageAuthenticator = types.BoolValue(*d.RequireMessageAuthenticator)
		}
		secret = types.StringValue(d.Secret)

		dataMapValue := map[string]attr.Value{
			"host":                          host,
			"keywrap_enabled":               keywrapEnabled,
			"keywrap_format":                keywrapFormat,
			"keywrap_kek":                   keywrapKek,
			"keywrap_mack":                  keywrapMack,
			"port":                          port,
			"require_message_authenticator": requireMessageAuthenticator,
			"secret":                        secret,
		}
		data, e := NewAuthServersValue(AuthServersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		authValueList = append(authValueList, data)
	}

	authStateListType := AuthServersValue{}.Type(ctx)
	authStateList, e := types.ListValueFrom(ctx, authStateListType, authValueList)
	diags.Append(e...)
	return authStateList
}
func radiusConfigPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RadiusConfig) basetypes.ObjectValue {
	var acctInterimInterval basetypes.Int64Value
	var acctServers = types.ListNull(AcctServersValue{}.Type(ctx))
	var authServers = types.ListNull(AcctServersValue{}.Type(ctx))
	var authServersRetries basetypes.Int64Value
	var authServersTimeout basetypes.Int64Value
	var coaEnabled basetypes.BoolValue
	var coaPort basetypes.Int64Value
	var network basetypes.StringValue
	var source_ip basetypes.StringValue

	if d.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*d.AcctInterimInterval))
	}
	if d.AcctServers != nil {
		acctServers = radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	}
	if d.AuthServers != nil {
		authServers = radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	}
	if d.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*d.AuthServersRetries))
	}
	if d.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}
	if d.CoaEnabled != nil {
		coaEnabled = types.BoolValue(*d.CoaEnabled)
	}
	if d.CoaPort != nil {
		coaPort = types.Int64Value(int64(*d.CoaPort))
	}
	if d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d.SourceIp != nil {
		source_ip = types.StringValue(*d.SourceIp)
	}

	data_map_value := map[string]attr.Value{
		"acct_interim_interval": acctInterimInterval,
		"acct_servers":          acctServers,
		"auth_servers":          authServers,
		"auth_servers_retries":  authServersRetries,
		"auth_servers_timeout":  authServersTimeout,
		"coa_enabled":           coaEnabled,
		"coa_port":              coaPort,
		"network":               network,
		"source_ip":             source_ip,
	}
	data, e := basetypes.NewObjectValue(RadiusConfigValue{}.AttributeTypes(ctx), data_map_value)
	diags.Append(e...)

	return data
}

/*
************************
*
* RADSEC
*
/************************
*/
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
func radsecPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Radsec) basetypes.ObjectValue {
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
	data, e := basetypes.NewObjectValue(RadsecValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func portConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ApPortConfig) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var disabled basetypes.BoolValue
		var dynamicVlan basetypes.ObjectValue
		var enableMacAuth basetypes.BoolValue
		var forwarding basetypes.StringValue
		var macAuthPreferred basetypes.BoolValue
		var macAuthProtocol basetypes.StringValue
		var mistNac basetypes.ObjectValue
		var mxTunnelId basetypes.StringValue
		var mxtunnelName basetypes.StringValue
		var portAuth basetypes.StringValue
		var portVlanId basetypes.Int64Value
		var radiusConfig basetypes.ObjectValue
		var radsec basetypes.ObjectValue
		var vlanId basetypes.Int64Value
		var vlanIds basetypes.StringValue
		var wxtunnelId basetypes.StringValue
		var wxtunnelRemoteId basetypes.StringValue

		if d.Disabled != nil {
			disabled = types.BoolValue(*d.Disabled)
		}
		if d.DynamicVlan != nil {
			dynamicVlan = dynamicVlanPortConfigSdkToTerraform(ctx, diags, d.DynamicVlan)
		}
		if d.EnableMacAuth != nil {
			enableMacAuth = types.BoolValue(*d.EnableMacAuth)
		}
		if d.Forwarding != nil {
			forwarding = types.StringValue(string(*d.Forwarding))
		}
		if d.MacAuthPreferred != nil {
			macAuthPreferred = types.BoolValue(*d.MacAuthPreferred)
		}
		if d.MacAuthProtocol != nil {
			macAuthProtocol = types.StringValue(string(*d.MacAuthProtocol))
		}
		if d.MistNac != nil {
			mistNac = mistNacPortConfigSdkToTerraform(ctx, diags, d.MistNac)
		}
		if d.MxTunnelId != nil {
			mxTunnelId = types.StringValue(d.MxTunnelId.String())
		}
		if d.MxtunnelName != nil {
			mxtunnelName = types.StringValue(*d.MxtunnelName)
		}
		if d.PortAuth != nil {
			portAuth = types.StringValue(string(*d.PortAuth))
		}
		if d.PortVlanId != nil {
			portVlanId = types.Int64Value(int64(*d.PortVlanId))
		}
		if d.RadiusConfig != nil {
			radiusConfig = radiusConfigPortConfigSdkToTerraform(ctx, diags, d.RadiusConfig)
		}
		if d.Radsec != nil {
			radsec = radsecPortConfigSdkToTerraform(ctx, diags, d.Radsec)
		}
		if d.VlanId != nil {
			vlanId = types.Int64Value(int64(*d.VlanId))
		}
		if d.VlanIds != nil {
			vlanIds = types.StringValue(*d.VlanIds)
		}
		if d.WxtunnelId != nil {
			wxtunnelId = types.StringValue(d.WxtunnelId.String())
		}
		if d.WxtunnelRemoteId != nil {
			wxtunnelRemoteId = types.StringValue(*d.WxtunnelRemoteId)
		}

		data_map_value := map[string]attr.Value{
			"disabled":           disabled,
			"dynamic_vlan":       dynamicVlan,
			"enable_mac_auth":    enableMacAuth,
			"forwarding":         forwarding,
			"mac_auth_preferred": macAuthPreferred,
			"mac_auth_protocol":  macAuthProtocol,
			"mist_nac":           mistNac,
			"mx_tunnel_id":       mxTunnelId,
			"mxtunnel_name":      mxtunnelName,
			"port_auth":          portAuth,
			"port_vlan_id":       portVlanId,
			"radius_config":      radiusConfig,
			"radsec":             radsec,
			"vlan_id":            vlanId,
			"vlan_ids":           vlanIds,
			"wxtunnel_id":        wxtunnelId,
			"wxtunnel_remote_id": wxtunnelRemoteId,
		}
		data, e := NewPortConfigValue(PortConfigValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PortConfigValue{}.Type(ctx), stateValueMap)
	diags.Append(e...)
	return stateResult
}
