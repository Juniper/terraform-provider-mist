package resource_org_network

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkSourceNat) basetypes.ObjectValue {
	var external_ip basetypes.StringValue

	if d != nil && d.ExternalIp != nil {
		external_ip = types.StringValue(*d.ExternalIp)
	}

	r_attr_type := SourceNatValue{}.AttributeTypes(ctx)
	r_attr_value := map[string]attr.Value{
		"external_ip": external_ip,
	}

	r, e := basetypes.NewObjectValue(r_attr_type, r_attr_value)
	diags.Append(e...)
	return r
}
