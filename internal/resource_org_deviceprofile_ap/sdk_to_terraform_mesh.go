package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func meshSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApMesh) MeshValue {

	var enabled basetypes.BoolValue
	var group basetypes.Int64Value
	var role basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Group.Value() != nil {
		group = types.Int64Value(int64(*d.Group.Value()))
	}
	if d.Role != nil {
		role = types.StringValue(string(*d.Role))
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"group":   group,
		"role":    role,
	}
	data, e := NewMeshValue(MeshValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
