package resource_org_idpprofile

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgIdpprofileModel) (models.IdpProfile, diag.Diagnostics) {
	var diags diag.Diagnostics
	data := models.IdpProfile{}
	unset := make(map[string]interface{})

	data.Name = plan.Name.ValueStringPointer()

	if plan.BaseProfile.ValueStringPointer() != nil {
		data.BaseProfile = (*models.IdpProfileBaseProfileEnum)(plan.BaseProfile.ValueStringPointer())
	}
	if !plan.Overwrites.IsNull() && !plan.Overwrites.IsUnknown() {
		data.Overwrites = overwritesTerraformToSdk(ctx, &diags, plan.Overwrites)
	} else {
		unset["-overwrites"] = ""
	}
	data.AdditionalProperties = unset
	return data, diags
}
