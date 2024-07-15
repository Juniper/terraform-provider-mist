package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func arpTableStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ArpTableStats) basetypes.ObjectValue {

	var arp_table_count basetypes.Int64Value
	var max_entries_supported basetypes.Int64Value

	if d.ArpTableCount != nil {
		arp_table_count = types.Int64Value(int64(*d.ArpTableCount))
	}
	if d.MaxEntriesSupported != nil {
		max_entries_supported = types.Int64Value(int64(*d.MaxEntriesSupported))
	}

	data_map_attr_type := ArpTableStatsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"arp_table_count":       arp_table_count,
		"max_entries_supported": max_entries_supported,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
