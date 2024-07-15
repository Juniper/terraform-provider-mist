package datasource_device_gateway_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func dhcpdStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.DhcpdStatLan) basetypes.MapValue {
	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var num_ips basetypes.Int64Value
		var num_leased basetypes.Int64Value

		if d.NumIps != nil {
			num_ips = types.Int64Value(int64(*d.NumIps))
		}
		if d.NumLeased != nil {
			num_leased = types.Int64Value(int64(*d.NumLeased))
		}

		data_map_attr_type := DhcpdStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"num_ips":    num_ips,
			"num_leased": num_leased,
		}
		data, e := NewDhcpdStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, DhcpdStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
