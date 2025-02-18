package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func appQosAppsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.WlanAppQosAppsProperties) basetypes.MapValue {

	mapAttrValues := make(map[string]attr.Value)
	for k, d := range m {

		var dscp basetypes.Int64Value
		var dstSubnet basetypes.StringValue
		var srcSubnet basetypes.StringValue

		if d.Dscp != nil {
			dscp = types.Int64Value(int64(*d.Dscp))
		}
		if d.DstSubnet != nil {
			dstSubnet = types.StringValue(*d.DstSubnet)
		}
		if d.SrcSubnet != nil {
			srcSubnet = types.StringValue(*d.SrcSubnet)
		}

		dataMapValue := map[string]attr.Value{
			"dscp":       dscp,
			"dst_subnet": dstSubnet,
			"src_subnet": srcSubnet,
		}
		data, e := NewAppsValue(AppsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		mapAttrValues[k] = data
	}

	r, e := types.MapValueFrom(ctx, AppsValue{}.Type(ctx), mapAttrValues)
	diags.Append(e...)
	return r
}

func appQosOthersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.WlanAppQosOthersItem) basetypes.ListValue {

	var dataList []OthersValue
	for _, d := range l {

		var dscp basetypes.Int64Value
		var dstSubnet basetypes.StringValue
		var portRanges basetypes.StringValue
		var protocol basetypes.StringValue
		var srcSubnet basetypes.StringValue

		if d.Dscp != nil {
			dscp = types.Int64Value(int64(*d.Dscp))
		}
		if d.DstSubnet != nil {
			dstSubnet = types.StringValue(*d.DstSubnet)
		}
		if d.PortRanges != nil {
			portRanges = types.StringValue(*d.PortRanges)
		}
		if d.Protocol != nil {
			protocol = types.StringValue(*d.Protocol)
		}
		if d.SrcSubnet != nil {
			srcSubnet = types.StringValue(*d.SrcSubnet)
		}

		dataMapValue := map[string]attr.Value{
			"dscp":        dscp,
			"dst_subnet":  dstSubnet,
			"port_ranges": portRanges,
			"protocol":    protocol,
			"src_subnet":  srcSubnet,
		}
		data, e := NewOthersValue(OthersValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, OthersValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func appQosSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.WlanAppQos) AppQosValue {
	var apps = types.MapNull(AppsValue{}.Type(ctx))
	var enabled basetypes.BoolValue
	var others = types.ListNull(OthersValue{}.Type(ctx))

	if len(d.Apps) > 0 {
		apps = appQosAppsSdkToTerraform(ctx, diags, d.Apps)
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Others != nil {
		others = appQosOthersSdkToTerraform(ctx, diags, d.Others)
	}

	dataMapValue := map[string]attr.Value{
		"apps":    apps,
		"enabled": enabled,
		"others":  others,
	}
	data, e := NewAppQosValue(AppQosValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data

}
