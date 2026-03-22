package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func proxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Proxy) ProxyValue {

	var disabled = types.BoolNull()
	var url = types.StringNull()

	if d.Disabled != nil {
		disabled = types.BoolValue(*d.Disabled)
	}
	if d.Url != nil {
		url = types.StringValue(*d.Url)
	}

	data_map_attr_type := ProxyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"disabled": disabled,
		"url":      url,
	}
	data, e := NewProxyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
