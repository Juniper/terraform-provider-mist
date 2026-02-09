package resource_org_idpprofile

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.IdpProfile) (OrgIdpprofileModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if d == nil {
		diags.AddError("Error: nil IdpProfile", "The SDK IdpProfile model is nil.")
		return OrgIdpprofileModel{}, diags
	}

	var baseProfile types.String
	if d.BaseProfile != nil {
		baseProfile = types.StringValue(string(*d.BaseProfile))
	}

	var id types.String
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}

	var name types.String
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}

	orgId := types.StringValue(d.OrgId.String())
	overwrites := overwritesSdkToTerraform(ctx, &diags, d.Overwrites)

	result := OrgIdpprofileModel{
		BaseProfile: baseProfile,
		Id:          id,
		Name:        name,
		OrgId:       orgId,
		Overwrites:  overwrites,
	}
	return result, diags
}
