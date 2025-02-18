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

	dataMapValue := map[string]attr.Value{
		"code":      code,
		"timestamp": timestamp,
	}
	data, e := basetypes.NewObjectValue(LastTroubleValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
