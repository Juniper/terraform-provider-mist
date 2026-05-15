package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func varsAnnotationsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d map[string]models.VarsAnnotation) basetypes.MapValue {
	dataMap := map[string]attr.Value{}
	for k, v := range d {
		var note basetypes.StringValue
		var annotationType basetypes.StringValue

		if v.Note != nil {
			note = types.StringValue(*v.Note)
		}
		if v.Type != nil {
			annotationType = types.StringValue(*v.Type)
		}

		dataMapValue := map[string]attr.Value{
			"note": note,
			"type": annotationType,
		}
		obj, e := NewVarsAnnotationsValue(VarsAnnotationsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)
		dataMap[k] = obj
	}
	r, e := types.MapValueFrom(ctx, VarsAnnotationsValue{}.Type(ctx), dataMap)
	diags.Append(e...)
	return r
}
