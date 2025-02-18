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

	dataMapValue := map[string]attr.Value{
		"url": url,
	}
	data, e := NewProxyValue(ProxyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
