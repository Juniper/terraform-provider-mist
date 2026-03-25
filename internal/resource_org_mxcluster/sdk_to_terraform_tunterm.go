package resource_org_mxcluster

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func tuntermDhcpdConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.TuntermDhcpdConfig) basetypes.MapValue {

	state_value_map_type := TuntermDhcpdConfigValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d.AdditionalProperties {
		var enabled = types.BoolNull()
		var servers = types.ListNull(types.StringType)
		var tuntermDhcpdConfigType = types.StringNull()

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

func tuntermExtraRoutesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.MxclusterTuntermExtraRoute) basetypes.MapValue {

	state_value_map_type := TuntermExtraRoutesValue{}.Type(ctx)
	state_value_map := make(map[string]attr.Value)

	for k, v := range d {
		var via types.String

		if v.Via != nil {
			via = types.StringValue(*v.Via)
		}

		data_map_attr_type := TuntermExtraRoutesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"via": via,
		}
		data, e := NewTuntermExtraRoutesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}

	state_result, e := types.MapValueFrom(ctx, state_value_map_type, state_value_map)
	diags.Append(e...)
	return state_result
}

func tuntermMonitoringSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d [][]models.TuntermMonitoringItem) types.List {

	if d == nil {
		return types.ListNull(types.ListType{ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"host":        types.StringType,
				"port":        types.Int64Type,
				"protocol":    types.StringType,
				"src_vlan_id": types.Int64Type,
				"timeout":     types.Int64Type,
			},
		}})
	}

	outerList := make([]attr.Value, len(d))

	for i, innerArray := range d {
		innerList := make([]attr.Value, len(innerArray))

		for j, item := range innerArray {
			var host = types.StringNull()
			var port = types.Int64Null()
			var protocol = types.StringNull()
			var srcVlanId = types.Int64Null()
			var timeout = types.Int64Null()

			if item.Host != nil {
				host = types.StringValue(*item.Host)
			}
			if item.Port != nil {
				port = types.Int64Value(int64(*item.Port))
			}
			if item.Protocol != nil {
				protocol = types.StringValue(string(*item.Protocol))
			}
			if item.SrcVlanId != nil {
				srcVlanId = types.Int64Value(int64(*item.SrcVlanId))
			}
			if item.Timeout != nil {
				timeout = types.Int64Value(int64(*item.Timeout))
			}

			itemObj, e := types.ObjectValue(
				map[string]attr.Type{
					"host":        types.StringType,
					"port":        types.Int64Type,
					"protocol":    types.StringType,
					"src_vlan_id": types.Int64Type,
					"timeout":     types.Int64Type,
				},
				map[string]attr.Value{
					"host":        host,
					"port":        port,
					"protocol":    protocol,
					"src_vlan_id": srcVlanId,
					"timeout":     timeout,
				},
			)
			diags.Append(e...)
			innerList[j] = itemObj
		}

		innerListValue, e := types.ListValue(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"host":        types.StringType,
				"port":        types.Int64Type,
				"protocol":    types.StringType,
				"src_vlan_id": types.Int64Type,
				"timeout":     types.Int64Type,
			},
		}, innerList)
		diags.Append(e...)
		outerList[i] = innerListValue
	}

	result, e := types.ListValue(types.ListType{ElemType: types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"host":        types.StringType,
			"port":        types.Int64Type,
			"protocol":    types.StringType,
			"src_vlan_id": types.Int64Type,
			"timeout":     types.Int64Type,
		},
	}}, outerList)
	diags.Append(e...)

	return result
}
