package resource_org_evpn_topology

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func overlayEvpnOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.EvpnOptionsOverlay {
	data := models.EvpnOptionsOverlay{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewOverlayValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.As.IsNull() && !plan.As.IsUnknown() {
				data.As = models.ToPointer(int(plan.As.ValueInt64()))
			}
		}
	}
	return &data
}

func underlayEvpnOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.ObjectValue) *models.EvpnOptionsUnderlay {
	data := models.EvpnOptionsUnderlay{}
	if !d.IsNull() || !d.IsUnknown() {
		plan, e := NewUnderlayValue(d.AttributeTypes(ctx), d.Attributes())
		if e != nil {
			diags.Append(e...)
		} else {
			if !plan.AsBase.IsNull() && !plan.AsBase.IsUnknown() {
				data.AsBase = models.ToPointer(int(plan.AsBase.ValueInt64()))
			}
			if !plan.RoutedIdPrefix.IsNull() && !plan.RoutedIdPrefix.IsUnknown() {
				data.RoutedIdPrefix = plan.RoutedIdPrefix.ValueStringPointer()
			}
			if !plan.Subnet.IsNull() && !plan.Subnet.IsUnknown() {
				data.Subnet = plan.Subnet.ValueStringPointer()
			}
			if !plan.UseIpv6.IsNull() && !plan.UseIpv6.IsUnknown() {
				data.UseIpv6 = plan.UseIpv6.ValueBoolPointer()
			}
		}
	}
	return &data
}

func vsInstanceEvpnOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.EvpnOptionsVsInstance {
	data_map := make(map[string]models.EvpnOptionsVsInstance)
	for k, v := range d.Elements() {
		var v_interface interface{} = v
		plan := v_interface.(VsInstancesValue)
		data := models.EvpnOptionsVsInstance{}
		if !plan.Networks.IsNull() && !plan.Networks.IsUnknown() {
			data.Networks = mist_transform.ListOfStringTerraformToSdk(ctx, plan.Networks)
		}
		data_map[k] = data
	}
	return data_map
}

func evpnOptionsTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d EvpnOptionsValue) *models.EvpnOptions {
	data := models.EvpnOptions{}

	if !d.AutoLoopbackSubnet.IsNull() && !d.AutoLoopbackSubnet.IsUnknown() {
		data.AutoLoopbackSubnet = d.AutoLoopbackSubnet.ValueStringPointer()
	}
	if !d.AutoLoopbackSubnet6.IsNull() && !d.AutoLoopbackSubnet6.IsUnknown() {
		data.AutoLoopbackSubnet6 = d.AutoLoopbackSubnet6.ValueStringPointer()
	}
	if !d.AutoRouterIdSubnet.IsNull() && !d.AutoRouterIdSubnet.IsUnknown() {
		data.AutoRouterIdSubnet = d.AutoRouterIdSubnet.ValueStringPointer()
	}
	if !d.AutoRouterIdSubnet6.IsNull() && !d.AutoRouterIdSubnet6.IsUnknown() {
		data.AutoRouterIdSubnet6 = d.AutoRouterIdSubnet6.ValueStringPointer()
	}
	if !d.CoreAsBorder.IsNull() && !d.CoreAsBorder.IsUnknown() {
		data.CoreAsBorder = d.CoreAsBorder.ValueBoolPointer()
	}
	if !d.Overlay.IsNull() && !d.Overlay.IsUnknown() {
		data.Overlay = overlayEvpnOptionsTerraformToSdk(ctx, diags, d.Overlay)
	}
	if !d.PerVlanVgaV4Mac.IsNull() && !d.PerVlanVgaV4Mac.IsUnknown() {
		data.PerVlanVgaV4Mac = d.PerVlanVgaV4Mac.ValueBoolPointer()
	}
	if !d.RoutedAt.IsNull() && !d.RoutedAt.IsUnknown() {
		data.RoutedAt = (*models.EvpnOptionsRoutedAtEnum)(d.RoutedAt.ValueStringPointer())
	}
	if !d.Underlay.IsNull() && !d.Underlay.IsUnknown() {
		data.Underlay = underlayEvpnOptionsTerraformToSdk(ctx, diags, d.Underlay)
	}
	if !d.VsInstances.IsNull() && !d.VsInstances.IsUnknown() {
		data.VsInstances = vsInstanceEvpnOptionsTerraformToSdk(ctx, diags, d.VsInstances)
	}

	return &data
}
