package resource_org_servicepolicy

import (
	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func idpTerraformToSdk(diags *diag.Diagnostics, d IdpValue) *models.IdpConfig {

	data := models.IdpConfig{}
	if d.AlertOnly.ValueBoolPointer() != nil {
		data.AlertOnly = d.AlertOnly.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.IdpprofileId.ValueStringPointer() != nil {
		idprofileId, e := uuid.Parse(d.IdpprofileId.ValueString())
		if e != nil {
			diags.AddError("Unable to parse IDP Profile ID", e.Error())
		} else {
			data.IdpprofileId = models.ToPointer(idprofileId)
		}
	}
	if d.Profile.ValueStringPointer() != nil {
		data.Profile = d.Profile.ValueStringPointer()
	}
	return &data
}
