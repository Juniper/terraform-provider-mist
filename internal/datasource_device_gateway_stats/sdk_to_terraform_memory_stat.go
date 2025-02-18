package datasource_device_gateway_stats

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

	dataMapValue := map[string]attr.Value{
		"usage": usage,
	}
	data, e := basetypes.NewObjectValue(MemoryStatValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
