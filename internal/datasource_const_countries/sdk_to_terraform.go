package datasource_const_countries

import (
	"context"
	"math/big"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstCountry) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := countrySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstCountriesValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func countrySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstCountry) ConstCountriesValue {
	o, _ := NewConstCountriesValue(
		ConstCountriesValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"alpha2":    types.StringValue(d.Alpha2),
			"certified": types.BoolValue(d.Certified),
			"name":      types.StringValue(d.Name),
			"numeric":   types.NumberValue(big.NewFloat(d.Numeric)),
		},
	)
	return o
}
