package datasource_const_traffic_types

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstTrafficType) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constAppCategorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstTrafficTypesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constAppCategorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstTrafficType) ConstTrafficTypesValue {
	var display basetypes.StringValue
	var dscp basetypes.Int64Value
	var failover_policy basetypes.StringValue
	var max_jitter basetypes.Int64Value
	var max_latency basetypes.Int64Value
	var max_loss basetypes.Int64Value
	var name basetypes.StringValue
	var traffic_class basetypes.StringValue

	if d.Display != nil {
		display = types.StringValue(*d.Display)
	}
	if d.Dscp != nil {
		dscp = types.Int64Value(int64(*d.Dscp))
	}
	if d.FailoverPolicy != nil {
		failover_policy = types.StringValue(*d.FailoverPolicy)
	}
	if d.MaxJitter != nil {
		max_jitter = types.Int64Value(int64(*d.MaxJitter))
	}
	if d.MaxLatency != nil {
		max_latency = types.Int64Value(int64(*d.MaxLatency))
	}
	if d.MaxLoss != nil {
		max_loss = types.Int64Value(int64(*d.MaxLoss))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.TrafficClass != nil {
		traffic_class = types.StringValue(*d.TrafficClass)
	}

	o, e := NewConstTrafficTypesValue(
		ConstTrafficTypesValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"display":         display,
			"dscp":            dscp,
			"failover_policy": failover_policy,
			"max_jitter":      max_jitter,
			"max_latency":     max_latency,
			"max_loss":        max_loss,
			"name":            name,
			"traffic_class":   traffic_class,
		},
	)
	diags.Append(e...)
	return o
}
