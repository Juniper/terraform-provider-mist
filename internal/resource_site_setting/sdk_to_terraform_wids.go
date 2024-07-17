package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func widsAuthFailureSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteWidsRepeatedAuthFailures) basetypes.ObjectValue {

	var duration basetypes.Int64Value
	var threshold basetypes.Int64Value

	if d.Duration != nil {
		duration = types.Int64Value(int64(*d.Duration))
	}
	if d.Threshold != nil {
		threshold = types.Int64Value(int64(*d.Threshold))
	}

	data_map_attr_type := RepeatedAuthFailuresValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"duration":  duration,
		"threshold": threshold,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func widsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteWids) WidsValue {

	var repeated_auth_failures basetypes.ObjectValue = types.ObjectNull(RepeatedAuthFailuresValue{}.AttributeTypes(ctx))

	if d != nil && d.RepeatedAuthFailures != nil {
		repeated_auth_failures = widsAuthFailureSdkToTerraform(ctx, diags, d.RepeatedAuthFailures)
	}

	data_map_attr_type := WidsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"repeated_auth_failures": repeated_auth_failures,
	}
	data, e := NewWidsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
