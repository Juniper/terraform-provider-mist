package resource_org_idpprofile

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, data *models.IdpProfile) (OrgIdpprofileModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	if data == nil {
		diags.AddError("Error: nil IdpProfile", "The SDK IdpProfile model is nil.")
		return OrgIdpprofileModel{}, diags
	}

	var baseProfile types.String
	if data.BaseProfile != nil {
		baseProfile = types.StringValue(string(*data.BaseProfile))
	}

	var id types.String
	if data.Id != nil {
		id = types.StringValue(data.Id.String())
	}

	var name types.String
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}

	result := OrgIdpprofileModel{
		Id:          id,
		OrgId:       types.StringValue(data.OrgId.String()),
		BaseProfile: baseProfile,
		Overwrites:  overwritesSdkToTerraform(ctx, &diags, data.Overwrites),
		Name:        name,
	}
	return result, diags
}
