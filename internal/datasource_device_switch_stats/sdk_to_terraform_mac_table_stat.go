package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func macTableStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MacTableStats) basetypes.ObjectValue {

	var mac_table_count basetypes.Int64Value
	var max_mac_entries_supported basetypes.Int64Value

	if d.MacTableCount != nil {
		mac_table_count = types.Int64Value(int64(*d.MacTableCount))
	}
	if d.MaxMacEntriesSupported != nil {
		max_mac_entries_supported = types.Int64Value(int64(*d.MaxMacEntriesSupported))
	}

	data_map_attr_type := MacTableStatsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"mac_table_count":           mac_table_count,
		"max_mac_entries_supported": max_mac_entries_supported,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
