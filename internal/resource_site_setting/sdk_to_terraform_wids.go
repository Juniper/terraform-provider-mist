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

	dataMapValue := map[string]attr.Value{
		"duration":  duration,
		"threshold": threshold,
	}
	data, e := basetypes.NewObjectValue(RepeatedAuthFailuresValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func widsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.SiteWids) WidsValue {

	var repeatedAuthFailures = types.ObjectNull(RepeatedAuthFailuresValue{}.AttributeTypes(ctx))

	if d != nil && d.RepeatedAuthFailures != nil {
		repeatedAuthFailures = widsAuthFailureSdkToTerraform(ctx, diags, d.RepeatedAuthFailures)
	}

	dataMapValue := map[string]attr.Value{
		"repeated_auth_failures": repeatedAuthFailures,
	}
	data, e := NewWidsValue(WidsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
