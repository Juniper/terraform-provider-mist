package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
			var host types.String
			var port types.Int64
			var protocol types.String
			var srcVlanId types.Int64
			var timeout types.Int64

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
