package resource_org_servicepolicy

import (
	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func avTerraformToSdk(diags *diag.Diagnostics, d AntivirusValue) *models.ServicePolicyAntivirus {

	data := models.ServicePolicyAntivirus{}
	if d.AvprofileId.ValueStringPointer() != nil {
		avprofileId, e := uuid.Parse(d.AvprofileId.ValueString())
		if e != nil {
			diags.AddError("Unable to parse IDP Profile ID", e.Error())
		} else {
			data.AvprofileId = models.ToPointer(avprofileId)
		}
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Profile.ValueStringPointer() != nil {
		data.Profile = d.Profile.ValueStringPointer()
	}
	return &data
}
