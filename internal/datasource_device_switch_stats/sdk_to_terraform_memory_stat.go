package datasource_device_switch_stats

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func memoryStatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MemoryStat) basetypes.ObjectValue {

	var usage basetypes.NumberValue

	if d != nil {
		usage = types.NumberValue(big.NewFloat(d.Usage))
	}

	data_map_attr_type := MemoryStatValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"usage": usage,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
