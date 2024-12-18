package resource_org_wlan

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
		var host basetypes.StringValue = types.StringValue(string(d.Host))
		var keywrap_enabled basetypes.BoolValue
		var keywrap_format basetypes.StringValue
		var keywrap_kek basetypes.StringValue
		var keywrap_mack basetypes.StringValue
		var port basetypes.Int64Value
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
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}

		data_map_value := map[string]attr.Value{
			"host":            host,
			"keywrap_enabled": keywrap_enabled,
			"keywrap_format":  keywrap_format,
			"keywrap_kek":     keywrap_kek,
			"keywrap_mack":    keywrap_mack,
			"port":            port,
			"secret":          secret,
		}
		data, e := NewAcctServersValue(AcctServersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)

		acct_value_list = append(acct_value_list, data)
	}

	acct_state_list, e := types.ListValueFrom(ctx, AcctServersValue{}.Type(ctx), acct_value_list)
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
		var port basetypes.Int64Value
		var require_message_authenticator basetypes.BoolValue
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
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}
		if d.RequireMessageAuthenticator != nil {
			require_message_authenticator = types.BoolValue(*d.RequireMessageAuthenticator)
		}

		data_map_value := map[string]attr.Value{
			"host":                          host,
			"keywrap_enabled":               keywrap_enabled,
			"keywrap_format":                keywrap_format,
			"keywrap_kek":                   keywrap_kek,
			"keywrap_mack":                  keywrap_mack,
			"port":                          port,
			"require_message_authenticator": require_message_authenticator,
			"secret":                        secret,
		}
		data, e := NewAuthServersValue(AuthServersValue{}.AttributeTypes(ctx), data_map_value)
		diags.Append(e...)

		auth_value_list = append(auth_value_list, data)
	}

	auth_state_list, e := types.ListValueFrom(ctx, AuthServersValue{}.Type(ctx), auth_value_list)
	diags.Append(e...)

	return auth_state_list
}
