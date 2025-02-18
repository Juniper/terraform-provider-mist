package datasource_const_webhooks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l []models.ConstWebhookTopic) (basetypes.SetValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var elements []attr.Value
	for _, d := range l {
		elem := constWebhookSdkToTerraform(ctx, &diags, d)
		elements = append(elements, elem)
	}

	dataSet, err := types.SetValue(ConstWebhooksValue{}.Type(ctx), elements)
	if err != nil {
		diags.Append(err...)
	}

	return dataSet, diags
}

func constWebhookSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d models.ConstWebhookTopic) ConstWebhooksValue {
	var forOrg basetypes.BoolValue
	var hasDeliveryResults basetypes.BoolValue
	var internal basetypes.BoolValue
	var key basetypes.StringValue

	if d.ForOrg != nil {
		forOrg = types.BoolValue(*d.ForOrg)
	}
	if d.HasDeliveryResults != nil {
		hasDeliveryResults = types.BoolValue(*d.HasDeliveryResults)
	}
	if d.Internal != nil {
		internal = types.BoolValue(*d.Internal)
	}
	if d.Key != nil {
		key = types.StringValue(*d.Key)
	}

	o, e := NewConstWebhooksValue(
		ConstWebhooksValue{}.AttributeTypes(ctx),
		map[string]attr.Value{
			"for_org":              forOrg,
			"has_delivery_results": hasDeliveryResults,
			"internal":             internal,
			"key":                  key,
		},
	)
	diags.Append(e...)
	return o
}
