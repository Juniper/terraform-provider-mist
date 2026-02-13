package resource_org_mxedge

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func versionsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.MxedgeVersions) VersionsValue {

	var mxagent types.String
	var tunterm types.String

	if d.Mxagent != nil {
		mxagent = types.StringValue(*d.Mxagent)
	}
	if d.Tunterm != nil {
		tunterm = types.StringValue(*d.Tunterm)
	}

	data_map_attr_type := VersionsValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"mxagent": mxagent,
		"tunterm": tunterm,
	}
	data, e := NewVersionsValue(data_map_attr_type, data_map_value)
	diags.Append(e...)

	return data
}
