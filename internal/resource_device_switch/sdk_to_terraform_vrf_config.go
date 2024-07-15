package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "terraform-provider-mist/internal/commons/utils"
)

func vrfConfigSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.VrfConfig) VrfConfigValue {

	var enabled basetypes.BoolValue

	if d != nil && d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	data_map_attr_type := VrfConfigValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"enabled": enabled,
	}
	data, e := NewVrfConfigValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}

func vrfInstanceExtraRouteSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VrfExtraRoute) basetypes.MapValue {
	map_item_type := make(map[string]attr.Value)
	for k, d := range m {
		var via basetypes.StringValue

		if d.Via != nil {
			via = types.StringValue(*d.Via)
		}

		data_map_attr_type := VrfExtraRoutesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"via": via,
		}
		data, e := NewVrfExtraRoutesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_item_type[k] = data
	}
	state_type := ExtraRoutesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, map_item_type)
	diags.Append(e...)
	return state_result
}

func vrfInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VrfInstance) basetypes.MapValue {

	map_item_type := make(map[string]attr.Value)
	for k, d := range m {
		var extra_routes basetypes.MapValue = types.MapNull(VrfExtraRoutesValue{}.Type(ctx))
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.ExtraRoutes != nil && len(d.ExtraRoutes) > 0 {
			extra_routes = vrfInstanceExtraRouteSdkToTerraform(ctx, diags, d.ExtraRoutes)
		}
		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}

		data_map_attr_type := VrfInstancesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"vrf_extra_routes": extra_routes,
			"networks":         networks,
		}
		data, e := NewVrfInstancesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		map_item_type[k] = data
	}
	state_type := VrfInstancesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, map_item_type)
	diags.Append(e...)
	return state_result
}
