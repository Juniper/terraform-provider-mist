package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func stpConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.SwitchStpConfig) StpConfigValue {

	var bridgePriority basetypes.StringValue

	if d.BridgePriority != nil {
		bridgePriority = types.StringValue(*d.BridgePriority)
	}

	dataMapValue := map[string]attr.Value{
		"bridge_priority": bridgePriority,
	}
	data, e := NewStpConfigValue(StpConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
