package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func sleThresholdsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SleThresholds) SleThresholdsValue {

	var capacity basetypes.Int64Value
	var coverage basetypes.Int64Value
	var throughput basetypes.Int64Value
	var timetoconnect basetypes.Int64Value

	if d.Capacity != nil {
		capacity = types.Int64Value(int64(*d.Capacity))
	}
	if d.Coverage != nil {
		coverage = types.Int64Value(int64(*d.Coverage))
	}
	if d.Throughput != nil {
		throughput = types.Int64Value(int64(*d.Throughput))
	}
	if d.TimeToConnect != nil {
		timetoconnect = types.Int64Value(int64(*d.TimeToConnect))
	}

	dataMapValue := map[string]attr.Value{
		"capacity":      capacity,
		"coverage":      coverage,
		"throughput":    throughput,
		"timetoconnect": timetoconnect,
	}
	data, e := NewSleThresholdsValue(SleThresholdsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
