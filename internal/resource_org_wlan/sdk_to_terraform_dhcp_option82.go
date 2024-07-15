package resource_org_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func injectDhcpOption82dkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanInjectDhcpOption82) InjectDhcpOption82Value {
	var circuit_id basetypes.StringValue
	var enabled basetypes.BoolValue

	if d != nil && d.CircuitId != nil {
		circuit_id = types.StringValue(*d.CircuitId)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := InjectDhcpOption82Value{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"circuit_id": circuit_id,
		"enabled":    enabled,
	}
	data, e := NewInjectDhcpOption82Value(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
