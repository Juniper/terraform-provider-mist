package datasource_device_gateway_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clusterStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsGatewayCluster) basetypes.ObjectValue {
	var state basetypes.StringValue

	if d.State.Value() != nil {
		state = types.StringValue(*d.State.Value())
	}

	dataMapValue := map[string]attr.Value{
		"state": state,
	}
	data, e := types.ObjectValue(ClusterStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
