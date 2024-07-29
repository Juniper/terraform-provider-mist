package datasource_const_app_sub_categories

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstAppSubcategoryDefinition) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constAppCategorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstAppSubCategoriesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constAppCategorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstAppSubcategoryDefinition) ConstAppSubCategoriesValue {

	o, e := NewConstAppSubCategoriesValue(
		ConstAppSubCategoriesValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"display":      types.StringValue(d.Display),
			"key":          types.StringValue(d.Key),
			"traffic_type": types.StringValue(d.TrafficType),
		},
	)
	diags.Append(e...)
	return o
}
