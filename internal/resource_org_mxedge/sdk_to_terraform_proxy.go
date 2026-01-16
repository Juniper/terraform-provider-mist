package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func proxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.Proxy) ProxyValue {

	var url types.String

	if d.Url != nil {
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
