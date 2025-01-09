package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func coaServersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.CoaServer) basetypes.ListValue {

	var data_list = []CoaServersValue{}
	for _, d := range l {
		var disable_event_timestamp_check basetypes.BoolValue
		var enabled basetypes.BoolValue
		var ip basetypes.StringValue
		var port basetypes.Int64Value
		var secret basetypes.StringValue

		if d.DisableEventTimestampCheck != nil {
			disable_event_timestamp_check = types.BoolValue(*d.DisableEventTimestampCheck)
		}
		if d.Enabled != nil {
			enabled = types.BoolValue(*d.Enabled)
		}
		ip = types.StringValue(string(d.Ip))
		if d.Port != nil {
			port = types.Int64Value(int64(*d.Port))
		}
		secret = types.StringValue(string(d.Secret))

		data_map_attr_type := CoaServersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"disable_event_timestamp_check": disable_event_timestamp_check,
			"enabled":                       enabled,
			"ip":                            ip,
			"port":                          port,
			"secret":                        secret,
		}
		data, e := NewCoaServersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, CoaServersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r

}
