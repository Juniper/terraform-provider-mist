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
	var for_org basetypes.BoolValue
	var has_delivery_results basetypes.BoolValue
	var internal basetypes.BoolValue
	var key basetypes.StringValue

	if d.ForOrg != nil {
		for_org = types.BoolValue(*d.ForOrg)
	}
	if d.HasDeliveryResults != nil {
		has_delivery_results = types.BoolValue(*d.HasDeliveryResults)
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
			"for_org":              for_org,
			"has_delivery_results": has_delivery_results,
			"internal":             internal,
			"key":                  key,
		},
	)
	diags.Append(e...)
	return o
}
