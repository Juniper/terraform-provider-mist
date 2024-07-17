package resource_org_servicepolicy

import (
	"context"

	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func idpTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d IdpValue) *models.IdpConfig {

	data := models.IdpConfig{}
	if d.AlertOnly.ValueBoolPointer() != nil {
		data.AlertOnly = d.AlertOnly.ValueBoolPointer()
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.IdpprofileId.ValueStringPointer() != nil {
		idprofile_id, e := uuid.Parse(d.IdpprofileId.ValueString())
		if e != nil {
			diags.AddError("Unable to parse IDP Profile ID", e.Error())
		} else {
			data.IdpprofileId = models.ToPointer(idprofile_id)
		}
	}
	if d.Profile.ValueStringPointer() != nil {
		data.Profile = d.Profile.ValueStringPointer()
	}
	return &data
}
