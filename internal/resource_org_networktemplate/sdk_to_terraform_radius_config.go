package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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

func radiusConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SwitchRadiusConfig) RadiusConfigValue {
	var acctImmediateUpdate basetypes.BoolValue
	var acctInterimInterval basetypes.Int64Value
	var acctServers = types.ListNull(AcctServersValue{}.Type(ctx))
	var authServers = types.ListNull(AuthServersValue{}.Type(ctx))
	var authServersRetries basetypes.Int64Value
	var authServerSelection basetypes.StringValue
	var authServersTimeout basetypes.Int64Value
	var coaEnabled basetypes.BoolValue
	var coaPort basetypes.StringValue
	var fastDot1xTimers basetypes.BoolValue
	var network basetypes.StringValue
	var sourceIp basetypes.StringValue

	if d != nil && d.AcctImmediateUpdate != nil {
		acctImmediateUpdate = types.BoolValue(*d.AcctImmediateUpdate)
	}
	if d != nil && d.AcctInterimInterval != nil {
		acctInterimInterval = types.Int64Value(int64(*d.AcctInterimInterval))
	}
	if d != nil && d.AcctServers != nil {
		acctServers = radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	}
	if d != nil && d.AuthServers != nil {
		authServers = radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	}
	if d != nil && d.AuthServersRetries != nil {
		authServersRetries = types.Int64Value(int64(*d.AuthServersRetries))
	}
	if d != nil && d.AuthServerSelection != nil {
		authServerSelection = types.StringValue(string(*d.AuthServerSelection))
	}
	if d != nil && d.AuthServersTimeout != nil {
		authServersTimeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}
	if d != nil && d.CoaEnabled != nil {
		coaEnabled = types.BoolValue(*d.CoaEnabled)
	}
	if d != nil && d.CoaPort != nil {
		coaPort = mistutils.RadiusCoaPortAsString(d.CoaPort)
	}
	if d != nil && d.FastDot1xTimers != nil {
		fastDot1xTimers = types.BoolValue(*d.FastDot1xTimers)
	}
	if d != nil && d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d != nil && d.SourceIp != nil {
		sourceIp = types.StringValue(*d.SourceIp)
	}

	dataMapValue := map[string]attr.Value{
		"acct_immediate_update": acctImmediateUpdate,
		"acct_interim_interval": acctInterimInterval,
		"acct_servers":          acctServers,
		"auth_servers":          authServers,
		"auth_servers_retries":  authServersRetries,
		"auth_servers_timeout":  authServersTimeout,
		"auth_server_selection": authServerSelection,
		"coa_enabled":           coaEnabled,
		"coa_port":              coaPort,
		"fast_dot1x_timers":     fastDot1xTimers,
		"network":               network,
		"source_ip":             sourceIp,
	}
	data, e := NewRadiusConfigValue(RadiusConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
