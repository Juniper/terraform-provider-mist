package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func airistaSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApAirista) AiristaValue {
	var enabled basetypes.BoolValue
	var host basetypes.StringValue
	var port basetypes.Int64Value

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Host.Value() != nil {
		host = types.StringValue(*d.Host.Value())
	}
	if d.Port.Value() != nil {
		port = types.Int64Value(int64(*d.Port.Value()))
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"host":    host,
		"port":    port,
	}
	data, e := NewAiristaValue(AiristaValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
