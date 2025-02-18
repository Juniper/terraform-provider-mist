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

	var appId basetypes.BoolValue
	var appImageUrl basetypes.StringValue
	var appProbe basetypes.BoolValue
	var category basetypes.StringValue
	var group basetypes.StringValue
	var key basetypes.StringValue
	var name basetypes.StringValue
	var signatureBased basetypes.BoolValue
	var ssrAppId basetypes.BoolValue

	if d.AppId != nil {
		appId = types.BoolValue(*d.AppId)
	}
	if d.AppImageUrl != nil {
		appImageUrl = types.StringValue(*d.AppImageUrl)
	}
	if d.AppProbe != nil {
		appProbe = types.BoolValue(*d.AppProbe)
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
		signatureBased = types.BoolValue(*d.SignatureBased)
	}
	if d.SsrAppId != nil {
		ssrAppId = types.BoolValue(*d.SsrAppId)
	}

	dataMapValue := map[string]attr.Value{
		"app_id":          appId,
		"app_image_url":   appImageUrl,
		"app_probe":       appProbe,
		"category":        category,
		"group":           group,
		"key":             key,
		"name":            name,
		"signature_based": signatureBased,
		"ssr_app_id":      ssrAppId,
	}
	o, e := NewConstApplicationsValue(ConstApplicationsValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return o
}
