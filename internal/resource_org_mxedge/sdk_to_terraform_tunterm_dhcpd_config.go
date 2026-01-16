package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermDhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeTuntermDhcpdConfig) basetypes.MapValue {

	state_value_map_type := TuntermDhcpdConfigValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d.AdditionalProperties {
		var enabled types.Bool
		var servers = types.ListNull(types.StringType)
		var tuntermDhcpdConfigType types.String

		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}
		if v.Servers != nil {
			servers = mistutils.ListOfStringSdkToTerraform(v.Servers)
		}
		if v.Type != nil {
			tuntermDhcpdConfigType = types.StringValue(string(*v.Type))
		}

		data_map_attr_type := TuntermDhcpdConfigValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"enabled": enabled,
			"servers": servers,
			"type":    tuntermDhcpdConfigType,
		}
		data, e := NewTuntermDhcpdConfigValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}

	state_result, e := types.MapValueFrom(ctx, state_value_map_type, state_value_map)
	diags.Append(e...)
	return state_result
}
