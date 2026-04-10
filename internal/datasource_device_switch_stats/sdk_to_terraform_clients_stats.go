package datasource_device_switch_stats

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func clientsStatsTotalSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchClientsStatsTotal) basetypes.ObjectValue {

	var numAps = types.ListNull(types.Int64Type)
	var numWiredClients basetypes.Int64Value

	if d.NumAps != nil {
		numAps = mistutils.ListOfIntSdkToTerraform(d.NumAps)
	}
	if d.NumWiredClients != nil {
		numWiredClients = types.Int64Value(int64(*d.NumWiredClients))
	}

	dataMapValue := map[string]attr.Value{
		"num_aps":           numAps,
		"num_wired_clients": numWiredClients,
	}
	data, e := basetypes.NewObjectValue(TotalValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func clientsStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsSwitchClientsStats) basetypes.ObjectValue {

	var total = types.ObjectNull(TotalValue{}.AttributeTypes(ctx))

	if d.Total != nil {
		total = clientsStatsTotalSdkToTerraform(ctx, diags, d.Total)
	}

	dataMapValue := map[string]attr.Value{
		"total": total,
	}
	data, e := basetypes.NewObjectValue(ClientsStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
