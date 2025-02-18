package resource_site_setting

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func vsInstanceSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, m map[string]models.VsInstanceProperty) basetypes.MapValue {
	stateValueMap := make(map[string]attr.Value)
	for k, d := range m {

		var networks = misttransform.ListOfStringSdkToTerraformEmpty()

		if d.Networks != nil {
			networks = misttransform.ListOfStringSdkToTerraform(d.Networks)
		}

		dataMapValue := map[string]attr.Value{
			"networks": networks,
		}
		data, e := NewVsInstanceValue(VsInstanceValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		stateValueMap[k] = data
	}
	stateType := VsInstanceValue{}.Type(ctx)
	stateResult, e := types.MapValueFrom(ctx, stateType, stateValueMap)
	diags.Append(e...)

	return stateResult
}
