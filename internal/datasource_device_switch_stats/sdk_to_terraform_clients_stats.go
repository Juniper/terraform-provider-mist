package datasource_device_switch_stats

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func clientsStatsTotalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchClientsStatsTotal) basetypes.ObjectValue {

	var num_aps basetypes.ListValue = types.ListNull(types.Int64Type)
	var num_wired_clients basetypes.Int64Value

	if d.NumAps != nil {
		num_aps = mist_transform.ListOfIntSdkToTerraform(ctx, d.NumAps)
	}
	if d.NumWiredClients != nil {
		num_wired_clients = types.Int64Value(int64(*d.NumWiredClients))
	}

	data_map_attr_type := TotalValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"num_aps":           num_aps,
		"num_wired_clients": num_wired_clients,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func clientsStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchClientsStats) basetypes.ObjectValue {

	var total basetypes.ObjectValue = types.ObjectNull(TotalValue{}.AttributeTypes(ctx))

	if d.Total != nil {
		total = clientsStatsTotalSdkToTerraform(ctx, diags, d.Total)
	}

	data_map_attr_type := ClientsStatsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"total": total,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
