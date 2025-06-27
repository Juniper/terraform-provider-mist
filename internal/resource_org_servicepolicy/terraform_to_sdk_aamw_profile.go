package resource_org_servicepolicy

import (
	"github.com/google/uuid"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func aamwTerraformToSdk(diags *diag.Diagnostics, d AamwValue) *models.ServicePolicyAamw {

	var data models.ServicePolicyAamw

	if d.AamwprofileId.ValueStringPointer() != nil {
		aamwprofileId, e := uuid.Parse(d.AamwprofileId.ValueString())
		if e != nil {
			diags.AddError("Unable to parse IDP Profile ID", e.Error())
		} else {
			data.AamwprofileId = models.ToPointer(aamwprofileId)
		}
	}
	if d.Enabled.ValueBoolPointer() != nil {
		data.Enabled = d.Enabled.ValueBoolPointer()
	}
	if d.Profile.ValueStringPointer() != nil {
		data.Profile = (*models.ServicePolicyAamwProfileEnum)(d.Profile.ValueStringPointer())
	}
	return &data
}
