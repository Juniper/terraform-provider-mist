package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func internalAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternalAccess) InternalAccessValue {
	var enabled basetypes.BoolValue
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewInternalAccessValue(InternalAccessValue{}.AttributeTypes(ctx), dataMapValue)

	diags.Append(e...)
	return data
}
