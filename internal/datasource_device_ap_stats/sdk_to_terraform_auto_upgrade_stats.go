package datasource_device_ap_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func autoUpgradeStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsApAutoUpgrade) basetypes.ObjectValue {
	var lastcheck basetypes.Int64Value

	if d.Lastcheck.Value() != nil {
		lastcheck = types.Int64Value(*d.Lastcheck.Value())
	}

	dataMapValue := map[string]attr.Value{
		"lastcheck": lastcheck,
	}
	data, e := types.ObjectValue(AutoUpgradeStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
