package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func mistDasSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeDas) MistDasValue {

	var coaServers = types.ListNull(CoaServersValue{}.Type(ctx))
	var enabled = types.BoolNull()

	if len(d.CoaServers) > 0 {
		var dataList []attr.Value

		for _, item := range d.CoaServers {
			var disableEventTimestampCheck = types.BoolNull()
			var itemEnabled = types.BoolNull()
			var host = types.StringNull()
			var port = types.Int64Null()
			var requireMessageAuthenticator = types.BoolNull()
			var secret = types.StringNull()

			if item.DisableEventTimestampCheck != nil {
				disableEventTimestampCheck = types.BoolValue(*item.DisableEventTimestampCheck)
			}
			if item.Enabled != nil {
				itemEnabled = types.BoolValue(*item.Enabled)
			}
			if item.Host != nil {
				host = types.StringValue(*item.Host)
			}
			if item.Port != nil {
				port = types.Int64Value(int64(*item.Port))
			}
			if item.RequireMessageAuthenticator != nil {
				requireMessageAuthenticator = types.BoolValue(*item.RequireMessageAuthenticator)
			}
			if item.Secret != nil {
				secret = types.StringValue(*item.Secret)
			}

			itemAttrType := CoaServersValue{}.AttributeTypes(ctx)
			itemValue := map[string]attr.Value{
				"disable_event_timestamp_check": disableEventTimestampCheck,
				"enabled":                       itemEnabled,
				"host":                          host,
				"port":                          port,
				"require_message_authenticator": requireMessageAuthenticator,
				"secret":                        secret,
			}
			item_o, e := NewCoaServersValue(itemAttrType, itemValue)
			diags.Append(e...)

			dataList = append(dataList, item_o)
		}

		r_list, e := types.ListValueFrom(ctx, CoaServersValue{}.Type(ctx), dataList)
		diags.Append(e...)
		coaServers = r_list
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := MistDasValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"coa_servers": coaServers,
		"enabled":     enabled,
	}
	data, e := NewMistDasValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
