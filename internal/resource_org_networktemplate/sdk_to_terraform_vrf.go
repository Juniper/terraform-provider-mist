package resource_org_networktemplate

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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
	stateType := VrfExtraRoutesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, mapItemValue)
	diags.Append(e...)
	return stateResult
}

func vrfInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.SwitchVrfInstance) basetypes.MapValue {

	dataMapValue := make(map[string]attr.Value)
	for k, d := range m {
		var extraRoutes = types.MapNull(VrfExtraRoutesValue{}.Type(ctx))
		var networks = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.ExtraRoutes != nil && len(d.ExtraRoutes) > 0 {
			extraRoutes = vrfInstanceExtraRouteSdkToTerraform(ctx, diags, d.ExtraRoutes)
		}
		if d.Networks != nil {
			networks = misttransform.ListOfStringSdkToTerraform(d.Networks)
		}

		vrfMapAttrType := VrfInstancesValue{}.AttributeTypes(ctx)
		vrfMapValue := map[string]attr.Value{
			"extra_routes": extraRoutes,
			"networks":     networks,
		}
		data, e := NewVrfInstancesValue(vrfMapAttrType, vrfMapValue)
		diags.Append(e...)

		dataMapValue[k] = data
	}
	stateType := VrfInstancesValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, dataMapValue)
	diags.Append(e...)
	return stateResult
}
