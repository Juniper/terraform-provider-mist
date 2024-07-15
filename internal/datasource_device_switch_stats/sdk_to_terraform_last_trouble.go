package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func lastTroubleSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.LastTrouble) basetypes.ObjectValue {

	var code basetypes.StringValue
	var timestamp basetypes.Int64Value

	if d.Code != nil {
		code = types.StringValue(*d.Code)
	}
	if d.Timestamp != nil {
		timestamp = types.Int64Value(int64(*d.Timestamp))
	}

	data_map_attr_type := LastTroubleValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"code":      code,
		"timestamp": timestamp,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
