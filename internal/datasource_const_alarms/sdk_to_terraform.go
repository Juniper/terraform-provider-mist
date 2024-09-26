package datasource_const_alarms

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstAlarmDefinition) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constAppCategorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstAlarmsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constAppCategorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstAlarmDefinition) ConstAlarmsValue {

	o, e := NewConstAlarmsValue(
		ConstAlarmsValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"display":  types.StringValue(d.Display),
			"group":    types.StringValue(d.Group),
			"key":      types.StringValue(d.Key),
			"severity": types.StringValue(d.Severity),
		},
	)
	diags.Append(e...)
	return o
}
