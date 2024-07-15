package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func autoUpgradeStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApStatsAutoUpgrade) basetypes.ObjectValue {
	var lastcheck basetypes.Int64Value

	if d.Lastcheck.Value() != nil {
		lastcheck = types.Int64Value(int64(*d.Lastcheck.Value()))
	}

	data_map_attr_type := AutoUpgradeStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"lastcheck": lastcheck,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
