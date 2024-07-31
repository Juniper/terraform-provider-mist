package datasource_const_applications

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstApplicationDefinition) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constAppCategorySdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstApplicationsValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constAppCategorySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstApplicationDefinition) ConstApplicationsValue {

	var app_id basetypes.BoolValue
	var app_image_url basetypes.StringValue
	var app_probe basetypes.BoolValue
	var category basetypes.StringValue
	var group basetypes.StringValue
	var key basetypes.StringValue
	var name basetypes.StringValue
	var signature_based basetypes.BoolValue
	var ssr_app_id basetypes.BoolValue

	if d.AppId != nil {
		app_id = types.BoolValue(*d.AppId)
	}
	if d.AppImageUrl != nil {
		app_image_url = types.StringValue(*d.AppImageUrl)
	}
	if d.AppProbe != nil {
		app_probe = types.BoolValue(*d.AppProbe)
	}
	if d.Category != nil {
		category = types.StringValue(*d.Category)
	}
	if d.Group != nil {
		group = types.StringValue(*d.Group)
	}
	if d.Key != nil {
		key = types.StringValue(*d.Key)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	if d.SignatureBased != nil {
		signature_based = types.BoolValue(*d.SignatureBased)
	}
	if d.SsrAppId != nil {
		ssr_app_id = types.BoolValue(*d.SsrAppId)
	}

	data_map_attr_type := ConstApplicationsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"app_id":          app_id,
		"app_image_url":   app_image_url,
		"app_probe":       app_probe,
		"category":        category,
		"group":           group,
		"key":             key,
		"name":            name,
		"signature_based": signature_based,
		"ssr_app_id":      ssr_app_id,
	}
	o, e := NewConstApplicationsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)
	return o
}
