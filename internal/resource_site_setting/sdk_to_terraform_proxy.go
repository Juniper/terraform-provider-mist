package resource_site_setting

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func proxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Proxy) ProxyValue {

	var disabled basetypes.BoolValue
	var url basetypes.StringValue

	if d != nil && d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d != nil && d.Url != nil {
		url = types.StringValue(*d.Url)
	}

	dataMapValue := map[string]attr.Value{
		"disabled": disabled,
		"url":      url,
	}
	data, e := NewProxyValue(ProxyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
