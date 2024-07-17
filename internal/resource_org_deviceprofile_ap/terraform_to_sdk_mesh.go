package resource_org_deviceprofile_ap

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func meshTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d MeshValue) *models.ApMesh {
	data := models.ApMesh{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Group.ValueInt64Pointer() != nil {
		data.Group = models.NewOptional(models.ToPointer(int(d.Group.ValueInt64())))
	}
	if d.Role.ValueStringPointer() != nil {
		data.Role = models.ToPointer(models.ApMeshRoleEnum(d.Role.ValueString()))
	}

	return &data
}
