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

	var base_profile types.String
	var id types.String
	var name types.String
	var org_id types.String
	var overwrites types.List = types.ListNull(OverwritesValue{}.Type(ctx))

	if d.BaseProfile != nil {
		base_profile = types.StringValue(string(*d.BaseProfile))
	}
	if d.Id != nil {
		id = types.StringValue(d.Id.String())
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	org_id = types.StringValue(d.OrgId.String())

	//if d.Overwrites != nil {
	overwrites = overwritesSdkToTerraform(ctx, &diags, d.Overwrites)
	//}

	state.BaseProfile = base_profile
	state.Id = id
	state.Name = name
	state.OrgId = org_id
	state.Overwrites = overwrites

	return state, diags

}
