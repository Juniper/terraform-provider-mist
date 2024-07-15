package resource_org_networktemplate

import (
	"context"
	mist_transform "terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func aclTagSpecsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.AclTagSpec) basetypes.ListValue {

	var data_list = []SpecsValue{}

	for _, d := range l {

		var port_range basetypes.StringValue
		var protocol basetypes.StringValue

		if d.PortRange != nil {
			port_range = types.StringValue(*d.PortRange)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}

		data_map_attr_type := SpecsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"port_range": port_range,
			"protocol":   protocol,
		}
		data, e := NewSpecsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	data_list_type := SpecsValue{}.Type(ctx)
	r, e := types.ListValueFrom(ctx, data_list_type, data_list)
	diags.Append(e...)
	return r
}
func aclTagsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.AclTag) basetypes.MapValue {

	state_value_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var gbp_tag basetypes.Int64Value
		var macs basetypes.ListValue = types.ListNull(types.StringType)
		var network basetypes.StringValue
		var radius_group basetypes.StringValue
		var specs basetypes.ListValue = types.ListNull(SpecsValue{}.Type(ctx))
		var subnets basetypes.ListValue = types.ListNull(types.StringType)
		var tag_type basetypes.StringValue = types.StringValue(string(d.Type))

		if d.GbpTag != nil {
			gbp_tag = types.Int64Value(int64(*d.GbpTag))
		}
		if d.Macs != nil {
			macs = mist_transform.ListOfStringSdkToTerraform(ctx, d.Macs)
		}
		if d.Network != nil {
			network = types.StringValue(*d.Network)
		}
		if d.RadiusGroup != nil {
			radius_group = types.StringValue(*d.RadiusGroup)
		}
		if d.Specs != nil {
			specs = aclTagSpecsSdkToTerraform(ctx, diags, d.Specs)
		}
		if d.Subnets != nil {
			subnets = mist_transform.ListOfStringSdkToTerraform(ctx, d.Subnets)
		}

		data_map_attr_type := AclTagsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"gbp_tag":      gbp_tag,
			"macs":         macs,
			"network":      network,
			"radius_group": radius_group,
			"specs":        specs,
			"subnets":      subnets,
			"type":         tag_type,
		}
		data, e := NewAclTagsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map_value[k] = data
	}
	state_result_map_type := AclTagsValue{}.Type(ctx)
	state_result_map, e := types.MapValueFrom(ctx, state_result_map_type, state_value_map_value)
	diags.Append(e...)
	return state_result_map
}
