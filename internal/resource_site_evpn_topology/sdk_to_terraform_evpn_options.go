package resource_site_evpn_topology

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func overlayEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptionsOverlay) basetypes.ObjectValue {
	var as basetypes.Int64Value

	if d.As != nil {
		as = types.Int64Value(int64(*d.As))
	}

	dataMapValue := map[string]attr.Value{
		"as": as,
	}
	data, e := basetypes.NewObjectValue(OverlayValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func underlayEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptionsUnderlay) basetypes.ObjectValue {
	var asBase basetypes.Int64Value
	var routedIdPrefix basetypes.StringValue
	var subnet basetypes.StringValue
	var useIpv6 basetypes.BoolValue

	if d.AsBase != nil {
		asBase = types.Int64Value(int64(*d.AsBase))
	}
	if d.RoutedIdPrefix != nil {
		routedIdPrefix = types.StringValue(*d.RoutedIdPrefix)
	}
	if d.Subnet != nil {
		subnet = types.StringValue(*d.Subnet)
	}
	if d.UseIpv6 != nil {
		useIpv6 = types.BoolValue(*d.UseIpv6)
	}

	dataMapValue := map[string]attr.Value{
		"as_base":          asBase,
		"routed_id_prefix": routedIdPrefix,
		"subnet":           subnet,
		"use_ipv6":         useIpv6,
	}
	data, e := basetypes.NewObjectValue(UnderlayValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func vsInstanceEvpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.EvpnOptionsVsInstance) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {
		var networks basetypes.ListValue

		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}

		dataMapValue := map[string]attr.Value{
			"networks": networks,
		}

		data, e := NewVsInstancesValue(VsInstancesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := VsInstancesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)
	return stateResult
}
func evpnOptionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.EvpnOptions) EvpnOptionsValue {
	var autoLoopbackSubnet basetypes.StringValue
	var autoLoopbackSubnet6 basetypes.StringValue
	var autoRouterIdSubnet basetypes.StringValue
	var autoRouterIdSubnet6 basetypes.StringValue
	var coreAsBorder basetypes.BoolValue
	var enableInbandZtp basetypes.BoolValue
	var overlay = types.ObjectNull(OverlayValue{}.AttributeTypes(ctx))
	var perVlanVgaV4Mac basetypes.BoolValue
	var perVlanVgaV6Mac basetypes.BoolValue
	var routedAt basetypes.StringValue
	var underlay = types.ObjectNull(UnderlayValue{}.AttributeTypes(ctx))
	var vsInstances = types.MapNull(VsInstancesValue{}.Type(ctx))

	if d.AutoLoopbackSubnet != nil {
		autoLoopbackSubnet = types.StringValue(*d.AutoLoopbackSubnet)
	}
	if d.AutoLoopbackSubnet6 != nil {
		autoLoopbackSubnet6 = types.StringValue(*d.AutoLoopbackSubnet6)
	}
	if d.AutoRouterIdSubnet != nil {
		autoRouterIdSubnet = types.StringValue(*d.AutoRouterIdSubnet)
	}
	if d.AutoRouterIdSubnet6 != nil {
		autoRouterIdSubnet6 = types.StringValue(*d.AutoRouterIdSubnet6)
	}
	if d.CoreAsBorder != nil {
		coreAsBorder = types.BoolValue(*d.CoreAsBorder)
	}
	if d.EnableInbandZtp != nil {
		enableInbandZtp = types.BoolValue(*d.EnableInbandZtp)
	}
	if d.Overlay != nil {
		overlay = overlayEvpnOptionsSdkToTerraform(ctx, diags, d.Overlay)
	}
	if d.PerVlanVgaV4Mac != nil {
		perVlanVgaV4Mac = types.BoolValue(*d.PerVlanVgaV4Mac)
	}
	if d.PerVlanVgaV6Mac != nil {
		perVlanVgaV6Mac = types.BoolValue(*d.PerVlanVgaV6Mac)
	}
	if d.RoutedAt != nil {
		routedAt = types.StringValue(string(*d.RoutedAt))
	}
	if d.Underlay != nil {
		underlay = underlayEvpnOptionsSdkToTerraform(ctx, diags, d.Underlay)
	}
	if d.VsInstances != nil {
		vsInstances = vsInstanceEvpnOptionsSdkToTerraform(ctx, diags, d.VsInstances)
	}

	dataMapValue := map[string]attr.Value{
		"auto_loopback_subnet":   autoLoopbackSubnet,
		"auto_loopback_subnet6":  autoLoopbackSubnet6,
		"auto_router_id_subnet":  autoRouterIdSubnet,
		"auto_router_id_subnet6": autoRouterIdSubnet6,
		"core_as_border":         coreAsBorder,
		"enable_inband_ztp":      enableInbandZtp,
		"overlay":                overlay,
		"per_vlan_vga_v4_mac":    perVlanVgaV4Mac,
		"per_vlan_vga_v6_mac":    perVlanVgaV6Mac,
		"routed_at":              routedAt,
		"underlay":               underlay,
		"vs_instances":           vsInstances,
	}
	data, e := NewEvpnOptionsValue(EvpnOptionsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
