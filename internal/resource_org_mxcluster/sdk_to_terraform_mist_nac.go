package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxclusterNac) MistNacValue {

	var acctServerPort = types.Int64Null()
	var authServerPort = types.Int64Null()
	var clientIps = types.MapNull(ClientIpsValue{}.Type(ctx))
	var enabled = types.BoolNull()
	var secret = types.StringNull()

	if d.AcctServerPort != nil {
		acctServerPort = types.Int64Value(int64(*d.AcctServerPort))
	}
	if d.AuthServerPort != nil {
		authServerPort = types.Int64Value(int64(*d.AuthServerPort))
	}
	if len(d.ClientIps) > 0 {
		dataMap := make(map[string]attr.Value)

		for key := range d.ClientIps {
			// ClientIpsValue has no fields, it's an empty object
			itemAttrType := ClientIpsValue{}.AttributeTypes(ctx)
			itemValue := map[string]attr.Value{}
			item_o, e := NewClientIpsValue(itemAttrType, itemValue)
			diags.Append(e...)

			dataMap[key] = item_o
		}

		r_map, e := types.MapValueFrom(ctx, ClientIpsValue{}.Type(ctx), dataMap)
		diags.Append(e...)
		clientIps = r_map
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Secret != nil {
		secret = types.StringValue(*d.Secret)
	}

	data_map_attr_type := MistNacValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"acct_server_port": acctServerPort,
		"auth_server_port": authServerPort,
		"client_ips":       clientIps,
		"enabled":          enabled,
		"secret":           secret,
	}
	data, e := NewMistNacValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
