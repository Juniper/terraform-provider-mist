package resource_org_deviceprofile_ap

import (
	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func meshTerraformToSdk(d MeshValue) *models.ApMesh {
	data := models.ApMesh{}

	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if !d.Bands.IsNull() && !d.Bands.IsUnknown() {
		data.Bands = mistutils.ListODot11TerraformToSdk(d.Bands)
	}
	if d.Group.ValueInt64Pointer() != nil {
		data.Group = models.NewOptional(models.ToPointer(int(d.Group.ValueInt64())))
	}
	if d.Role.ValueStringPointer() != nil {
		data.Role = models.ToPointer(models.ApMeshRoleEnum(d.Role.ValueString()))
	}

	return &data
}
