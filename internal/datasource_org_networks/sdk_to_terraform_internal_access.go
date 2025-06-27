package datasource_org_networks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func InternalAccessSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.NetworkInternalAccess) basetypes.ObjectValue {
	var enabled basetypes.BoolValue
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := basetypes.NewObjectValueFrom(ctx, InternalAccessValue{}.AttributeTypes(ctx), dataMapValue)

	diags.Append(e...)
	return data
}
