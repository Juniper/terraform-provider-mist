package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func fwupdateSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.FwupdateStat) basetypes.ObjectValue {

	var progress basetypes.Int64Value
	var status basetypes.StringValue
	var status_id basetypes.Int64Value
	var timestamp basetypes.Float64Value
	var will_retry basetypes.BoolValue

	if d.Progress.Value() != nil {
		progress = types.Int64Value(int64(*d.Progress.Value()))
	}
	if d.Status.Value() != nil {
		status = types.StringValue(string(*d.Status.Value()))
	}
	if d.StatusId.Value() != nil {
		status_id = types.Int64Value(int64(*d.StatusId.Value()))
	}
	if d.Timestamp.Value() != nil {
		timestamp = types.Float64Value(*d.Timestamp.Value())
	}
	if d.WillRetry.Value() != nil {
		will_retry = types.BoolValue(*d.WillRetry.Value())
	}

	data_map_attr_type := FwupdateValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"progress":   progress,
		"status":     status,
		"status_id":  status_id,
		"timestamp":  timestamp,
		"will_retry": will_retry,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
