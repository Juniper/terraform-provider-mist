package datasource_site_wlans

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanAppQosAppsProperties) basetypes.MapValue {

	map_attr_values := make(map[string]attr.Value)
	for k, d := range m {

		var dscp basetypes.Int64Value
		var dst_subnet basetypes.StringValue
		var src_subnet basetypes.StringValue

		if d.Dscp != nil {
			dscp = types.Int64Value(int64(*d.Dscp))
		}
		if d.DstSubnet != nil {
			dst_subnet = types.StringValue(*d.DstSubnet)
		}
		if d.SrcSubnet != nil {
			src_subnet = types.StringValue(*d.SrcSubnet)
		}

		data_map_attr_type := AppsValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"dscp":       dscp,
			"dst_subnet": dst_subnet,
			"src_subnet": src_subnet,
		}
		data, e := NewAppsValue(data_map_attr_type, data_map_value)
		diags.Append(e...)
		map_attr_values[k] = data
	}

	r, e := types.MapValueFrom(ctx, AppsValue{}.Type(ctx), map_attr_values)
	diags.Append(e...)
	return r
}

func appQosOthersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.WlanAppQosOthersItem) basetypes.ListValue {

	var data_list = []OthersValue{}
	for _, d := range l {

		var dscp basetypes.Int64Value
		var dst_subnet basetypes.StringValue
		var port_ranges basetypes.StringValue
		var protocol basetypes.StringValue
		var src_subnet basetypes.StringValue

		if d.Dscp != nil {
			dscp = types.Int64Value(int64(*d.Dscp))
		}
		if d.DstSubnet != nil {
			dst_subnet = types.StringValue(*d.DstSubnet)
		}
		if d.PortRanges != nil {
			port_ranges = types.StringValue(*d.PortRanges)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}
		if d.SrcSubnet != nil {
			src_subnet = types.StringValue(*d.SrcSubnet)
		}

		data_map_attr_type := OthersValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"dscp":        dscp,
			"dst_subnet":  dst_subnet,
			"port_ranges": port_ranges,
			"protocol":    protocol,
			"src_subnet":  src_subnet,
		}
		data, e := NewOthersValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		data_list = append(data_list, data)
	}
	r, e := types.ListValueFrom(ctx, OthersValue{}.Type(ctx), data_list)
	diags.Append(e...)
	return r
}

func appQosSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanAppQos) basetypes.ObjectValue {
	var apps basetypes.MapValue = types.MapNull(AppsValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var others basetypes.ListValue = types.ListNull(OthersValue{}.Type(ctx))

	if len(d.Apps) > 0 {
		apps = appQosAppsSdkToTerraform(ctx, diags, d.Apps)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Others != nil {
		others = appQosOthersSdkToTerraform(ctx, diags, d.Others)
	}

	data_map_attr_type := AppQosValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"apps":    apps,
		"enabled": enabled,
		"others":  others,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data

}
