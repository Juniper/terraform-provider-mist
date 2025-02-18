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

	var arpTableCount basetypes.Int64Value
	var maxEntriesSupported basetypes.Int64Value

	if d.ArpTableCount != nil {
		arpTableCount = types.Int64Value(int64(*d.ArpTableCount))
	}
	if d.MaxEntriesSupported != nil {
		maxEntriesSupported = types.Int64Value(int64(*d.MaxEntriesSupported))
	}

	dataMapValue := map[string]attr.Value{
		"arp_table_count":       arpTableCount,
		"max_entries_supported": maxEntriesSupported,
	}
	data, e := basetypes.NewObjectValue(ArpTableStatsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
