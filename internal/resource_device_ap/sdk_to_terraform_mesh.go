package resource_device_ap

import (
	"context"
	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func meshSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.ApMesh) MeshValue {

	var enabled basetypes.BoolValue
	bands := types.ListNull(types.StringType)
	var group basetypes.Int64Value
	var role basetypes.StringValue

	if d.Enabled != nil {
		enabled = types.BoolValue(*d.Enabled)
	}
	if d.Bands != nil {
		bands = misttransform.ListOfDot11SdkToTerraform(d.Bands)
	}
	if d.Group.Value() != nil {
		group = types.Int64Value(int64(*d.Group.Value()))
	}
	if d.Role != nil {
		role = types.StringValue(string(*d.Role))
	}

	dataMapValue := map[string]attr.Value{
		"enabled": enabled,
		"bands":   bands,
		"group":   group,
		"role":    role,
	}
	data, e := NewMeshValue(MeshValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
