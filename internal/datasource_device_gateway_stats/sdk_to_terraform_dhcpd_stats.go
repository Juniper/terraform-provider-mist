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
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var numIps basetypes.Int64Value
		var numLeased basetypes.Int64Value

		if d.NumIps != nil {
			numIps = types.Int64Value(int64(*d.NumIps))
		}
		if d.NumLeased != nil {
			numLeased = types.Int64Value(int64(*d.NumLeased))
		}

		dataMapValue := map[string]attr.Value{
			"num_ips":    numIps,
			"num_leased": numLeased,
		}
		data, e := NewDhcpdStatValue(DhcpdStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, DhcpdStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
