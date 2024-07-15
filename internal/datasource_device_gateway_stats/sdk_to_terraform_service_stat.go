package datasource_device_gateway_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ServiceStatProperty) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var ash_version basetypes.StringValue
		var cia_version basetypes.StringValue
		var ember_version basetypes.StringValue
		var ipsec_client_version basetypes.StringValue
		var mist_agent_version basetypes.StringValue
		var package_version basetypes.StringValue
		var testing_tools_version basetypes.StringValue
		var wheeljack_version basetypes.StringValue

		if d.AshVersion != nil {
			ash_version = types.StringValue(*d.AshVersion)
		}
		if d.CiaVersion != nil {
			cia_version = types.StringValue(*d.CiaVersion)
		}
		if d.EmberVersion != nil {
			ember_version = types.StringValue(*d.EmberVersion)
		}
		if d.IpsecClientVersion != nil {
			ipsec_client_version = types.StringValue(*d.IpsecClientVersion)
		}
		if d.MistAgentVersion != nil {
			mist_agent_version = types.StringValue(*d.MistAgentVersion)
		}
		if d.PackageVersion != nil {
			package_version = types.StringValue(*d.PackageVersion)
		}
		if d.TestingToolsVersion != nil {
			testing_tools_version = types.StringValue(*d.TestingToolsVersion)
		}
		if d.WheeljackVersion != nil {
			wheeljack_version = types.StringValue(*d.WheeljackVersion)
		}

		data_map_attr_type := ServiceStatValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"ash_version":           ash_version,
			"cia_version":           cia_version,
			"ember_version":         ember_version,
			"ipsec_client_version":  ipsec_client_version,
			"mist_agent_version":    mist_agent_version,
			"package_version":       package_version,
			"testing_tools_version": testing_tools_version,
			"wheeljack_version":     wheeljack_version,
		}
		data, e := NewServiceStatValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, ServiceStatValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
