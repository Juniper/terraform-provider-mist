package resource_site_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func portMirroringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchPortMirroringProperty) basetypes.MapValue {
	dataMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var inputNetworksIngress = types.ListNull(types.StringType)
		var inputPortIdsEgress = types.ListNull(types.StringType)
		var inputPortIdsIngress = types.ListNull(types.StringType)
		var outputNetwork basetypes.StringValue
		var outputIpAddress basetypes.StringValue
		var outputPortId basetypes.StringValue

		if d.InputNetworksIngress != nil {
			inputNetworksIngress = mistutils.ListOfStringSdkToTerraform(d.InputNetworksIngress)
		}
		if d.InputPortIdsEgress != nil {
			inputPortIdsEgress = mistutils.ListOfStringSdkToTerraform(d.InputPortIdsEgress)
		}
		if d.InputPortIdsIngress != nil {
			inputPortIdsIngress = mistutils.ListOfStringSdkToTerraform(d.InputPortIdsIngress)
		}
		if d.OutputIpAddress != nil {
			outputIpAddress = types.StringValue(*d.OutputIpAddress)
		}
		if d.OutputNetwork != nil {
			outputNetwork = types.StringValue(*d.OutputNetwork)
		}
		if d.OutputPortId != nil {
			outputPortId = types.StringValue(*d.OutputPortId)
		}

		itemMapValue := map[string]attr.Value{
			"input_networks_ingress": inputNetworksIngress,
			"input_port_ids_egress":  inputPortIdsEgress,
			"input_port_ids_ingress": inputPortIdsIngress,
			"output_ip_address":      outputIpAddress,
			"output_network":         outputNetwork,
			"output_port_id":         outputPortId,
		}
		data, e := NewPortMirroringValue(PortMirroringValue{}.AttributeTypes(ctx), itemMapValue)
		diags.Append(e...)

		dataMapValue[k] = data
	}
	stateType := PortMirroringValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, dataMapValue)
	diags.Append(e...)
	return stateResult
}
