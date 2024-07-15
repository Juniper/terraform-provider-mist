package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SwitchRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStatsSwitchRedundancy) basetypes.ObjectValue {

	var num_redundant_aps basetypes.Int64Value

	if d.NumRedundantAps.Value() != nil {
		num_redundant_aps = types.Int64Value(int64(*d.NumRedundantAps.Value()))
	}

	data_map_attr_type := SwitchRedundancyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"num_redundant_aps": num_redundant_aps,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
