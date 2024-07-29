package datasource_const_app_categories

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstAppCategoryDefinition) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constAppCategorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstAppCategoriesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constAppCategorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstAppCategoryDefinition) ConstAppCategoriesValue {
	var filter basetypes.ObjectValue = types.ObjectNull(FiltersValue{}.AttributeTypes(ctx))
	if d.Filters != nil {
		filter = constAppCategoryFiltersSdkToTerraform(ctx, diags, *d.Filters)
	}

	o, e := NewConstAppCategoriesValue(
		ConstAppCategoriesValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"display":  types.StringValue(d.Display),
			"filters":  filter,
			"includes": mist_transform.ListOfStringSdkToTerraform(ctx, d.Includes),
			"key":      types.StringValue(d.Key),
		},
	)
	diags.Append(e...)
	return o
}

func constAppCategoryFiltersSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstAppCategoryDefinitionFilters) basetypes.ObjectValue {
	var srx basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var ssr basetypes.ListValue = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)

	if d.Srx != nil {
		srx = mist_transform.ListOfStringSdkToTerraform(ctx, d.Srx)
	}
	if d.Ssr != nil {
		ssr = mist_transform.ListOfStringSdkToTerraform(ctx, d.Ssr)
	}

	data_map_attr_type := FiltersValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"srx": srx,
		"ssr": ssr,
	}
	data, e := basetypes.NewObjectValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
