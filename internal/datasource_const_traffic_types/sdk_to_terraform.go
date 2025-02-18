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
	var failoverPolicy basetypes.StringValue
	var maxJitter basetypes.Int64Value
	var maxLatency basetypes.Int64Value
	var maxLoss basetypes.Int64Value
	var name basetypes.StringValue
	var trafficClass basetypes.StringValue

	if d.Display != nil {
		display = types.StringValue(*d.Display)
	}
	if d.Dscp != nil {
		dscp = types.Int64Value(int64(*d.Dscp))
	}
	if d.FailoverPolicy != nil {
		failoverPolicy = types.StringValue(*d.FailoverPolicy)
	}
	if d.MaxJitter != nil {
		maxJitter = types.Int64Value(int64(*d.MaxJitter))
	}
	if d.MaxLatency != nil {
		maxLatency = types.Int64Value(int64(*d.MaxLatency))
	}
	if d.MaxLoss != nil {
		maxLoss = types.Int64Value(int64(*d.MaxLoss))
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.TrafficClass != nil {
		trafficClass = types.StringValue(*d.TrafficClass)
	}

	o, e := NewConstTrafficTypesValue(
		ConstTrafficTypesValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"display":         display,
			"dscp":            dscp,
			"failover_policy": failoverPolicy,
			"max_jitter":      maxJitter,
			"max_latency":     maxLatency,
			"max_loss":        maxLoss,
			"name":            name,
			"traffic_class":   trafficClass,
		},
	)
	diags.Append(e...)
	return o
}
