package datasource_device_ap_stats

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
	var timestamp basetypes.Float64Value

	if d.Code != nil {
		code = types.StringValue(*d.Code)
	}
	if d.Timestamp != nil {
		timestamp = types.Float64Value(*d.Timestamp)
	}

	dataMapValue := map[string]attr.Value{
		"code":      code,
		"timestamp": timestamp,
	}
	data, e := basetypes.NewObjectValue(LastTroubleValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
