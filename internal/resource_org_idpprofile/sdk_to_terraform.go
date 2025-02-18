package resource_org_idpprofile

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.IdpProfile) (OrgIdpprofileModel, diag.Diagnostics) {
	var state OrgIdpprofileModel
	var diags diag.Diagnostics

	var baseProfile types.String
	var id types.String
	var name types.String
	var orgId types.String
	var overwrites = types.ListNull(OverwritesValue{}.Type(ctx))

	if d.BaseProfile != nil {
		baseProfile = types.StringValue(string(*d.BaseProfile))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	orgId = types.StringValue(d.OrgId.String())

	//if d.Overwrites != nil {
	overwrites = overwritesSdkToTerraform(ctx, &diags, d.Overwrites)
	//}

	state.BaseProfile = baseProfile
	state.Id = id
	state.Name = name
	state.OrgId = orgId
	state.Overwrites = overwrites

	return state, diags

}
