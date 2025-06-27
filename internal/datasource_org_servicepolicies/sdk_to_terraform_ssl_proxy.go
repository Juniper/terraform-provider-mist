package datasource_org_servicepolicies

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func sslProxySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ServicePolicySslProxy) basetypes.ObjectValue {

	var ciphersCategory basetypes.StringValue
	var enabled basetypes.BoolValue

	if d.CiphersCategory != nil {
		ciphersCategory = types.StringValue(string(*d.CiphersCategory))
	}
	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}

	dataMapValue := map[string]attr.Value{
		"ciphers_category": ciphersCategory,
		"enabled":          enabled,
	}
	data, e := basetypes.NewObjectValue(SslProxyValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
