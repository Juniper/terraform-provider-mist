package resource_org_vpn

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, d *models.Vpn) (OrgVpnModel, diag.Diagnostics) {
	var state OrgVpnModel
	var diags diag.Diagnostics

	var id types.String
	var name types.String
	var org_id types.String
	var paths types.Map = types.MapNull(PathsValue{}.Type(ctx))

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, &diags, d.Paths)
	}

	state.Id = id
	state.Name = name
	state.OrgId = org_id
	state.Paths = paths

	return state, diags
}

func vpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPath) basetypes.MapValue {
	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {
		var bfd_profile basetypes.StringValue
		var ip basetypes.StringValue
		var pod basetypes.Int64Value

		if d.BfdProfile != nil {
			bfd_profile = types.StringValue(string(*d.BfdProfile))
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}

		data_map_attr_type := PathsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"bfd_profile": bfd_profile,
			"ip":          ip,
			"pod":         pod,
		}
		data, e := NewPathsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_attr_values[k] = data
	}
	state_result, e := types.MapValueFrom(ctx, PathsValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return state_result
}
