package resource_device_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func gatewayMgmtSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.GatewayMgmt) GatewayMgmtValue {
	var configRevertTimer basetypes.Int64Value

	if d != nil {
		if d.ConfigRevertTimer != nil {
			configRevertTimer = types.Int64Value(int64(*d.ConfigRevertTimer))
		}
	}

	dataMapValue := map[string]attr.Value{
		"config_revert_timer": configRevertTimer,
	}
	data, e := NewGatewayMgmtValue(GatewayMgmtValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
