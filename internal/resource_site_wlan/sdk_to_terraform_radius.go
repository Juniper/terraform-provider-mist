package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

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
		var port basetypes.Int64Value
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
			port = types.Int64Value(int64(*d.Port))
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

	acctStateList, e := types.ListValueFrom(ctx, AcctServersValue{}.Type(ctx), acctValueList)
	diags.Append(e...)

	return acctStateList
}

func radiusServersAuthSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.RadiusAuthServer) basetypes.ListValue {
	var authValueList []attr.Value
	for _, d := range l {
		var host = types.StringValue(d.Host)
		var keywrapEnabled basetypes.BoolValue
		var keywrapFormat basetypes.StringValue
		var keywrapKek basetypes.StringValue
		var keywrapMack basetypes.StringValue
		var port basetypes.Int64Value
		var requireMessageAuthenticator basetypes.BoolValue
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
			port = types.Int64Value(int64(*d.Port))
		}
		if d.RequireMessageAuthenticator != nil {
			requireMessageAuthenticator = types.BoolValue(*d.RequireMessageAuthenticator)
		}

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

	authStateList, e := types.ListValueFrom(ctx, AuthServersValue{}.Type(ctx), authValueList)
	diags.Append(e...)

	return authStateList
}
