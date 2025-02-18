package datasource_device_ap_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SwitchRedundancySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApSwitchRedundancy) basetypes.ObjectValue {

	var numRedundantAps basetypes.Int64Value

	if d.NumRedundantAps.Value() != nil {
		numRedundantAps = types.Int64Value(int64(*d.NumRedundantAps.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"num_redundant_aps": numRedundantAps,
	}
	data, e := basetypes.NewObjectValue(SwitchRedundancyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
