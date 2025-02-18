package datasource_device_switch_stats

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func serviceStatsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.ServiceStatProperty) basetypes.MapValue {

	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var ashVersion basetypes.StringValue
		var ciaVersion basetypes.StringValue
		var emberVersion basetypes.StringValue
		var ipsecClientVersion basetypes.StringValue
		var mistAgentVersion basetypes.StringValue
		var packageVersion basetypes.StringValue
		var testingToolsVersion basetypes.StringValue
		var wheeljackVersion basetypes.StringValue

		if d.AshVersion != nil {
			ashVersion = types.StringValue(*d.AshVersion)
		}
		if d.CiaVersion != nil {
			ciaVersion = types.StringValue(*d.CiaVersion)
		}
		if d.EmberVersion != nil {
			emberVersion = types.StringValue(*d.EmberVersion)
		}
		if d.IpsecClientVersion != nil {
			ipsecClientVersion = types.StringValue(*d.IpsecClientVersion)
		}
		if d.MistAgentVersion != nil {
			mistAgentVersion = types.StringValue(*d.MistAgentVersion)
		}
		if d.PackageVersion != nil {
			packageVersion = types.StringValue(*d.PackageVersion)
		}
		if d.TestingToolsVersion != nil {
			testingToolsVersion = types.StringValue(*d.TestingToolsVersion)
		}
		if d.WheeljackVersion != nil {
			wheeljackVersion = types.StringValue(*d.WheeljackVersion)
		}

		dataMapValue := map[string]attr.Value{
			"ash_version":           ashVersion,
			"cia_version":           ciaVersion,
			"ember_version":         emberVersion,
			"ipsec_client_version":  ipsecClientVersion,
			"mist_agent_version":    mistAgentVersion,
			"package_version":       packageVersion,
			"testing_tools_version": testingToolsVersion,
			"wheeljack_version":     wheeljackVersion,
		}
		data, e := NewServiceStatValue(ServiceStatValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, ServiceStatValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
