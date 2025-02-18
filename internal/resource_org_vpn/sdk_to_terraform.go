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
	var orgId types.String
	var paths = types.MapNull(PathsValue{}.Type(ctx))

	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		orgId = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, &diags, d.Paths)
	}

	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.Paths = paths

	return state, diags
}

func vpnPathsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VpnPath) basetypes.MapValue {
	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {
		var bfdProfile basetypes.StringValue
		var ip basetypes.StringValue
		var pod basetypes.Int64Value

		if d.BfdProfile != nil {
			bfdProfile = types.StringValue(string(*d.BfdProfile))
		}
		if d.Ip != nil {
			ip = types.StringValue(*d.Ip)
		}
		if d.Pod != nil {
			pod = types.Int64Value(int64(*d.Pod))
		}

		dataMapValue := map[string]attr.Value{
			"bfd_profile": bfdProfile,
			"ip":          ip,
			"pod":         pod,
		}
		data, e := NewPathsValue(PathsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapAttrValues[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, PathsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return stateResult
}
