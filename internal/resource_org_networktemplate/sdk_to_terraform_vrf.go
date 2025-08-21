package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func vrfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VrfConfig) VrfConfigValue {

	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewVrfConfigValue(VrfConfigValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}

func vrfInstanceExtraRouteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VrfExtraRoute) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		dataMapValue := map[string]attr.Value{
			"via": via,
		}
		data, e := NewVrfExtraRoutesValue(VrfExtraRoutesValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}

	stateResult, e := types.MapValueFrom(ctx, VrfExtraRoutesValue{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return stateResult
}

func vrfInstanceExtraRoute6SdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VrfExtraRoute) basetypes.MapValue {
	mapItemValue := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		dataMapValue := map[string]attr.Value{
			"via": via,
		}
		data, e := NewVrfExtraRoutes6Value(VrfExtraRoutes6Value{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		mapItemValue[k] = data
	}
	stateResult, e := types.MapValueFrom(ctx, VrfExtraRoutes6Value{}.Type(ctx), mapItemValue)
	diags.Append(e...)
	return stateResult
}

func vrfInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchVrfInstance) basetypes.MapValue {

	dataMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var evpnAutoLoopbackSubnet basetypes.StringValue
		var evpnAutoLoopbackSubnet6 basetypes.StringValue
		var extraRoutes = types.MapNull(VrfExtraRoutesValue{}.Type(ctx))
		var extraRoutes6 = types.MapNull(VrfExtraRoutes6Value{}.Type(ctx))
		var networks = mistutils.ListOfStringSdkToTerraformEmpty()

		if d.EvpnAutoLoopbackSubnet != nil {
			evpnAutoLoopbackSubnet = types.StringValue(*d.EvpnAutoLoopbackSubnet)
		}
		if d.EvpnAutoLoopbackSubnet6 != nil {
			evpnAutoLoopbackSubnet6 = types.StringValue(*d.EvpnAutoLoopbackSubnet6)
		}
		if len(d.ExtraRoutes) > 0 {
			extraRoutes = vrfInstanceExtraRouteSdkToTerraform(ctx, diags, d.ExtraRoutes)
		}
		if len(d.ExtraRoutes6) > 0 {
			extraRoutes6 = vrfInstanceExtraRoute6SdkToTerraform(ctx, diags, d.ExtraRoutes6)
		}
		if d.Networks != nil {
			networks = mistutils.ListOfStringSdkToTerraform(d.Networks)
		}

		vrfMapValue := map[string]attr.Value{
			"evpn_auto_loopback_subnet":  evpnAutoLoopbackSubnet,
			"evpn_auto_loopback_subnet6": evpnAutoLoopbackSubnet6,
			"extra_routes":               extraRoutes,
			"extra_routes6":              extraRoutes6,
			"networks":                   networks,
		}
		data, e := NewVrfInstancesValue(VrfInstancesValue{}.AttributeTypes(ctx), vrfMapValue)
		diags.Append(e...)

		dataMapValue[k] = data
	}

	stateResult, e := types.MapValueFrom(ctx, VrfInstancesValue{}.Type(ctx), dataMapValue)
	diags.Append(e...)
	return stateResult
}
