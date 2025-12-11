package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mistNacSkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanMistNac) basetypes.ObjectValue {
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

	dataMapValue := map[string]attr.Value{
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
	data, e := basetypes.NewObjectValue(MistNacValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
