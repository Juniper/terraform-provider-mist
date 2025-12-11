package datasource_device_gateway_stats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func macTableStatStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.StatsGatewayMacTableStats) basetypes.ObjectValue {
	var macTableCount basetypes.Int64Value
	var maxMacEntriesSupported basetypes.Int64Value

	if d.MacTableCount != nil {
		macTableCount = types.Int64Value(int64(*d.MacTableCount))
	}
	if d.MaxMacEntriesSupported != nil {
		maxMacEntriesSupported = types.Int64Value(int64(*d.MaxMacEntriesSupported))
	}

	dataMapValue := map[string]attr.Value{
		"mac_table_count":           macTableCount,
		"max_mac_entries_supported": maxMacEntriesSupported,
	}
	data, e := types.ObjectValue(MacTableStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
