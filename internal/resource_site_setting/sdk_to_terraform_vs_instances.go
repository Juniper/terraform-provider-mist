package resource_site_setting

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vsInstanceSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VsInstanceProperty) basetypes.MapValue {
	state_value_map := make(map[string]attr.Value)
	for k, d := range m {

		var networks basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

		if d.Networks != nil {
			networks = mist_transform.ListOfStringSdkToTerraform(ctx, d.Networks)
		}

		data_map_attr_type := VsInstanceValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"networks": networks,
		}
		data, e := NewVlansValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		state_value_map[k] = data
	}
	state_type := VsInstanceValue{}.Type(ctx)
	state_result, e := types.MapValueFrom(ctx, state_type, state_value_map)
	diags.Append(e...)

	return state_result
}
