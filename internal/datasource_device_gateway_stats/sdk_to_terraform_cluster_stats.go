package datasource_device_gateway_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func clusterStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayStatsCluster) basetypes.ObjectValue {
	var state basetypes.StringValue

	if d.State.Value() != nil {
		state = types.StringValue(*d.State.Value())
	}

	data_map_attr_type := ClusterStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"state": state,
	}
	data, e := types.ObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
