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

	var url basetypes.StringValue

	if d != nil && d.Url != nil {
		url = types.StringValue(*d.Url)
	}

	data_map_attr_type := ProxyValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"url": url,
	}
	data, e := NewProxyValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
