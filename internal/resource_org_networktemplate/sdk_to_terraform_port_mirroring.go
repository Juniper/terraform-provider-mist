package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func portMirroringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchPortMirroringProperty) basetypes.MapValue {
	data_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var input_networks_ingress basetypes.ListValue = types.ListNull(types.StringType)
		var input_port_ids_egress basetypes.ListValue = types.ListNull(types.StringType)
		var input_port_ids_ingress basetypes.ListValue = types.ListNull(types.StringType)
		var output_network basetypes.StringValue
		var output_port_id basetypes.StringValue

		if d.InputNetworksIngress != nil {
			input_networks_ingress = mist_transform.ListOfStringSdkToTerraform(ctx, d.InputNetworksIngress)
		}
		if d.InputPortIdsEgress != nil {
			input_port_ids_egress = mist_transform.ListOfStringSdkToTerraform(ctx, d.InputPortIdsEgress)
		}
		if d.InputPortIdsIngress != nil {
			input_port_ids_ingress = mist_transform.ListOfStringSdkToTerraform(ctx, d.InputPortIdsIngress)
		}
		if d.OutputNetwork != nil {
			output_network = types.StringValue(*d.OutputNetwork)
		}
		if d.OutputPortId != nil {
			output_port_id = types.StringValue(*d.OutputPortId)
		}

		data_map_attr_type := PortMirroringValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"input_networks_ingress": input_networks_ingress,
			"input_port_ids_egress":  input_port_ids_egress,
			"input_port_ids_ingress": input_port_ids_ingress,
			"output_network":         output_network,
			"output_port_id":         output_port_id,
		}
		data, e := NewPortMirroringValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_map_value[k] = data
	}
	state_type := PortMirroringValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}
