package datasource_org_vpns

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.Vpn) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := vpnSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(OrgVpnsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func vpnSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.Vpn) OrgVpnsValue {
	var created_time basetypes.NumberValue
	var id basetypes.StringValue
	var modified_time basetypes.NumberValue
	var name types.String
	var org_id types.String
	var paths types.Map = types.MapNull(PathsValue{}.Type(ctx))

	if d.CreatedTime != nil {
		created_time = types.NumberValue(big.NewFloat(*d.CreatedTime))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.ModifiedTime != nil {
		modified_time = types.NumberValue(big.NewFloat(*d.ModifiedTime))
	}

	name = types.StringValue(d.Name)

	if d.OrgId != nil {
		org_id = types.StringValue(d.OrgId.String())
	}
	if d.Paths != nil && len(d.Paths) > 0 {
		paths = vpnPathsSdkToTerraform(ctx, diags, d.Paths)
	}

	data_map_attr_type := OrgVpnsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"created_time":  created_time,
		"id":            id,
		"modified_time": modified_time,
		"name":          name,
		"org_id":        org_id,
		"paths":         paths,
	}
	data, e := NewOrgVpnsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
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
