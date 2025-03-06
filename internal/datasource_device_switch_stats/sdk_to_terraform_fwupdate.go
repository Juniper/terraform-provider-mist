package datasource_device_switch_stats

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
	var statusId basetypes.Int64Value
	var timestamp basetypes.Float64Value
	var willRetry basetypes.BoolValue

	if d.Progress.Value() != nil {
		progress = types.Int64Value(int64(*d.Progress.Value()))
	}
	if d.Status.Value() != nil {
		status = types.StringValue(string(*d.Status.Value()))
	}
	if d.StatusId.Value() != nil {
		statusId = types.Int64Value(int64(*d.StatusId.Value()))
	}
	if d.Timestamp != nil {
		timestamp = types.Float64Value(*d.Timestamp)
	}
	if d.WillRetry.Value() != nil {
		willRetry = types.BoolValue(*d.WillRetry.Value())
	}

	dataMapValue := map[string]attr.Value{
		"progress":   progress,
		"status":     status,
		"status_id":  statusId,
		"timestamp":  timestamp,
		"will_retry": willRetry,
	}
	data, e := basetypes.NewObjectValue(FwupdateValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
