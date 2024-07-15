package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func radiusServersAcctSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadiusAcctServer) basetypes.ListValue {
	var acct_value_list []attr.Value
	for _, d := range l {
		var host basetypes.StringValue = types.StringValue(d.Host)
		var keywrap_enabled basetypes.BoolValue
		var keywrap_format basetypes.StringValue
		var keywrap_kek basetypes.StringValue
		var keywrap_mack basetypes.StringValue
		var port basetypes.Int64Value = types.Int64Value(int64(*d.Port))
		var secret basetypes.StringValue = types.StringValue(d.Secret)

		if d.KeywrapEnabled != nil {
			keywrap_enabled = types.BoolValue(*d.KeywrapEnabled)
		}
		if d.KeywrapFormat != nil {
			keywrap_format = types.StringValue(string(*d.KeywrapFormat))
		}
		if d.KeywrapKek != nil {
			keywrap_kek = types.StringValue(*d.KeywrapKek)
		}
		if d.KeywrapMack != nil {
			keywrap_mack = types.StringValue(*d.KeywrapMack)
		}

		data_map_attr_type := AcctServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"host":            host,
			"keywrap_enabled": keywrap_enabled,
			"keywrap_format":  keywrap_format,
			"keywrap_kek":     keywrap_kek,
			"keywrap_mack":    keywrap_mack,
			"port":            port,
			"secret":          secret,
		}
		data, e := NewAcctServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, data)
	}

	acct_state_list_type := AcctServersValue{}.Type(ctx)
	acct_state_list, e := types.ListValueFrom(ctx, acct_state_list_type, acct_value_list)
	diags.Append(e...)

	return acct_state_list
}

func radiusServersAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadiusAuthServer) basetypes.ListValue {
	var auth_value_list []attr.Value
	for _, d := range l {
		var host basetypes.StringValue = types.StringValue(d.Host)
		var keywrap_enabled basetypes.BoolValue
		var keywrap_format basetypes.StringValue
		var keywrap_kek basetypes.StringValue
		var keywrap_mack basetypes.StringValue
		var port basetypes.Int64Value = types.Int64Value(int64(*d.Port))
		var secret basetypes.StringValue = types.StringValue(d.Secret)

		if d.KeywrapEnabled != nil {
			keywrap_enabled = types.BoolValue(*d.KeywrapEnabled)
		}
		if d.KeywrapFormat != nil {
			keywrap_format = types.StringValue(string(*d.KeywrapFormat))
		}
		if d.KeywrapKek != nil {
			keywrap_kek = types.StringValue(*d.KeywrapKek)
		}
		if d.KeywrapMack != nil {
			keywrap_mack = types.StringValue(*d.KeywrapMack)
		}

		data_map_attr_type := AuthServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"host":            host,
			"keywrap_enabled": keywrap_enabled,
			"keywrap_format":  keywrap_format,
			"keywrap_kek":     keywrap_kek,
			"keywrap_mack":    keywrap_mack,
			"port":            port,
			"secret":          secret,
		}
		data, e := NewAcctServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		auth_value_list = append(auth_value_list, data)
	}

	auth_state_list_type := AuthServersValue{}.Type(ctx)
	auth_state_list, e := types.ListValueFrom(ctx, auth_state_list_type, auth_value_list)
	diags.Append(e...)
	return auth_state_list
}

func radiusConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.RadiusConfig) RadiusConfigValue {
	var acct_interim_interval basetypes.Int64Value
	var acct_servers basetypes.ListValue = types.ListNull(AcctServersValue{}.Type(ctx))
	var auth_servers basetypes.ListValue = types.ListNull(AuthServersValue{}.Type(ctx))
	var auth_servers_retries basetypes.Int64Value
	var auth_servers_timeout basetypes.Int64Value
	var coa_enabled basetypes.BoolValue
	var coa_port basetypes.Int64Value
	var network basetypes.StringValue
	var source_ip basetypes.StringValue

	if d != nil && d.AcctInterimInterval != nil {
		acct_interim_interval = types.Int64Value(int64(*d.AcctInterimInterval))
	}
	if d != nil && d.AcctServers != nil {
		acct_servers = radiusServersAcctSdkToTerraform(ctx, diags, d.AcctServers)
	}
	if d != nil && d.AuthServers != nil {
		auth_servers = radiusServersAuthSdkToTerraform(ctx, diags, d.AuthServers)
	}
	if d != nil && d.AuthServersRetries != nil {
		auth_servers_retries = types.Int64Value(int64(*d.AuthServersRetries))
	}
	if d != nil && d.AuthServersTimeout != nil {
		auth_servers_timeout = types.Int64Value(int64(*d.AuthServersTimeout))
	}
	if d != nil && d.CoaEnabled != nil {
		coa_enabled = types.BoolValue(*d.CoaEnabled)
	}
	if d != nil && d.CoaPort != nil {
		coa_port = types.Int64Value(int64(*d.CoaPort))
	}
	if d != nil && d.Network != nil {
		network = types.StringValue(*d.Network)
	}
	if d != nil && d.SourceIp != nil {
		source_ip = types.StringValue(*d.SourceIp)
	}

	data_map_attr_type := RadiusConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"acct_interim_interval": acct_interim_interval,
		"acct_servers":          acct_servers,
		"auth_servers":          auth_servers,
		"auth_servers_retries":  auth_servers_retries,
		"auth_servers_timeout":  auth_servers_timeout,
		"coa_enabled":           coa_enabled,
		"coa_port":              coa_port,
		"network":               network,
		"source_ip":             source_ip,
	}
	data, e := NewRadiusConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
