package datasource_org_networks

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func sourceNatSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NetworkSourceNat) basetypes.ObjectValue {
	var externalIp basetypes.StringValue

	if d != nil && d.ExternalIp != nil {
		externalIp = types.StringValue(*d.ExternalIp)
	}

	rAttrType := SourceNatValue{}.AttributeTypes(ctx)
	rAttrValue := map[string]attr.Value{
		"external_ip": externalIp,
	}

	r, e := basetypes.NewObjectValue(rAttrType, rAttrValue)
	diags.Append(e...)
	return r
}
