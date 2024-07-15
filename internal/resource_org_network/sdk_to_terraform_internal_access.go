package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func InternalAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternalAccess) InternalAccessValue {
	var enabled basetypes.BoolValue
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := InternalAccessValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewInternalAccessValue(data_map_attr_type, data_map_value)

	diags.Append(e...)
	return data
}
