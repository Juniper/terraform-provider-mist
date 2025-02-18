package datasource_org_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func injectDhcpOption82dkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.WlanInjectDhcpOption82) basetypes.ObjectValue {
	var circuitId basetypes.StringValue
	var enabled basetypes.BoolValue

	if d != nil && d.CircuitId != nil {
		circuitId = types.StringValue(*d.CircuitId)
	}
	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"circuit_id": circuitId,
		"enabled":    enabled,
	}
	data, e := basetypes.NewObjectValue(InjectDhcpOption82Value{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
