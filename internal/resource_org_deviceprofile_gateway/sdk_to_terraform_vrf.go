package resource_org_deviceprofile_gateway

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
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

func vrfInstancesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.GatewayVrfInstance) basetypes.MapValue {

	data_map_value := make(map[string]attr.Value)
	for k, d := range m {
		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}

		vrf_map_attr_type := VrfInstancesValue{}.AttributeTypes(ctx)
		vrf_map_value := map[string]attr.Value{
			"networks": networks,
		}
		data, e := NewVrfInstancesValue(vrf_map_attr_type, vrf_map_value)
		diags.Append(e...)

		data_map_value[k] = data
	}
	state_type := VrfInstancesValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, data_map_value)
	diags.Append(e...)
	return state_result
}
