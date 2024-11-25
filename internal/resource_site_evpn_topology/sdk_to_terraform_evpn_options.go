package resource_site_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func overlayEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptionsOverlay) basetypes.ObjectValue {
	var as basetypes.Int64Value

	if d.As != nil {
		as = types.Int64Value(int64(*d.As))
	}

	data_map_attr_type := OverlayValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"as": as,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func underlayEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptionsUnderlay) basetypes.ObjectValue {
	var as_base basetypes.Int64Value
	var routed_id_prefix basetypes.StringValue
	var subnet basetypes.StringValue
	var use_ipv6 basetypes.BoolValue

	if d.AsBase != nil {
		as_base = types.Int64Value(int64(*d.AsBase))
	}
	if d.RoutedIdPrefix != nil {
		routed_id_prefix = types.StringValue(*d.RoutedIdPrefix)
	}
	if d.Subnet != nil {
		subnet = types.StringValue(*d.Subnet)
	}
	if d.UseIpv6 != nil {
		use_ipv6 = types.BoolValue(*d.UseIpv6)
	}

	data_map_attr_type := UnderlayValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"as_base":          as_base,
		"routed_id_prefix": routed_id_prefix,
		"subnet":           subnet,
		"use_ipv6":         use_ipv6,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func vsInstanceEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.EvpnOptionsVsInstance) basetypes.MapValue {
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {
		var networks basetypes.ListValue

		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}

		data_map_attr_type := VsInstancesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"networks": networks,
		}

		data, e := NewVsInstancesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := VsInstancesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)
	return state_result
}
func evpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptions) EvpnOptionsValue {
	var auto_loopback_subnet basetypes.StringValue
	var auto_loopback_subnet6 basetypes.StringValue
	var auto_router_id_subnet basetypes.StringValue
	var auto_router_id_subnet6 basetypes.StringValue
	var core_as_border basetypes.BoolValue
	var overlay basetypes.ObjectValue = types.ObjectNull(OverlayValue{}.AttributeTypes(ctx))
	var per_vlan_vga_v4_mac basetypes.BoolValue
	var routed_at basetypes.StringValue
	var underlay basetypes.ObjectValue = types.ObjectNull(UnderlayValue{}.AttributeTypes(ctx))
	var vs_instances basetypes.MapValue = types.MapNull(VsInstancesValue{}.Type(ctx))

	if d.AutoLoopbackSubnet != nil {
		auto_loopback_subnet = types.StringValue(*d.AutoLoopbackSubnet)
	}
	if d.AutoLoopbackSubnet6 != nil {
		auto_loopback_subnet6 = types.StringValue(*d.AutoLoopbackSubnet6)
	}
	if d.AutoRouterIdSubnet != nil {
		auto_router_id_subnet = types.StringValue(*d.AutoRouterIdSubnet)
	}
	if d.AutoRouterIdSubnet6 != nil {
		auto_router_id_subnet6 = types.StringValue(*d.AutoRouterIdSubnet6)
	}
	if d.CoreAsBorder != nil {
		core_as_border = types.BoolValue(*d.CoreAsBorder)
	}
	if d.Overlay != nil {
		overlay = overlayEvpnOptionsSdkToTerraform(ctx, diags, d.Overlay)
	}
	if d.PerVlanVgaV4Mac != nil {
		per_vlan_vga_v4_mac = types.BoolValue(*d.PerVlanVgaV4Mac)
	}
	if d.RoutedAt != nil {
		routed_at = types.StringValue(string(*d.RoutedAt))
	}
	if d.Underlay != nil {
		underlay = underlayEvpnOptionsSdkToTerraform(ctx, diags, d.Underlay)
	}
	if d.VsInstances != nil {
		vs_instances = vsInstanceEvpnOptionsSdkToTerraform(ctx, diags, d.VsInstances)
	}

	data_map_attr_type := EvpnOptionsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"auto_loopback_subnet":   auto_loopback_subnet,
		"auto_loopback_subnet6":  auto_loopback_subnet6,
		"auto_router_id_subnet":  auto_router_id_subnet,
		"auto_router_id_subnet6": auto_router_id_subnet6,
		"core_as_border":         core_as_border,
		"overlay":                overlay,
		"per_vlan_vga_v4_mac":    per_vlan_vga_v4_mac,
		"routed_at":              routed_at,
		"underlay":               underlay,
		"vs_instances":           vs_instances,
	}
	data, e := NewEvpnOptionsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
