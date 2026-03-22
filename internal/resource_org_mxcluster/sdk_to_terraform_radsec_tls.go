package resource_org_mxcluster

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func radsecTlsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxclusterRadsecTls) RadsecTlsValue {

	var keypair = types.StringNull()

	if d.Keypair != nil {
		keypair = types.StringValue(*d.Keypair)
	}

	data_map_attr_type := RadsecTlsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"keypair": keypair,
	}
	data, e := NewRadsecTlsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
