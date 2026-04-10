package resource_org_mxedge

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func tuntermPortConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TuntermPortConfig) TuntermPortConfigValue {

	var downstreamPorts = types.ListNull(types.StringType)
	var separateUpstreamDownstream = types.BoolNull()
	var upstreamPortVlanId = types.StringNull()
	var upstreamPorts = types.ListNull(types.StringType)

	if len(d.DownstreamPorts) > 0 {
		downstreamPorts = mistutils.ListOfStringSdkToTerraform(d.DownstreamPorts)
	}
	if d.SeparateUpstreamDownstream != nil {
		separateUpstreamDownstream = types.BoolValue(*d.SeparateUpstreamDownstream)
	}
	if d.UpstreamPortVlanId != nil {
		upstreamPortVlanId = mistutils.ContainerAsString(d.UpstreamPortVlanId)
	}
	if len(d.UpstreamPorts) > 0 {
		upstreamPorts = mistutils.ListOfStringSdkToTerraform(d.UpstreamPorts)
	}

	data_map_attr_type := TuntermPortConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"downstream_ports":             downstreamPorts,
		"separate_upstream_downstream": separateUpstreamDownstream,
		"upstream_port_vlan_id":        upstreamPortVlanId,
		"upstream_ports":               upstreamPorts,
	}
	data, e := NewTuntermPortConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
