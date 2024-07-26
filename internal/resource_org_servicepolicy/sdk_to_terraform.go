package resource_org_servicepolicy

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.OrgServicePolicy) (OrgServicepolicyModel, diag.Diagnostics) {
	var state OrgServicepolicyModel
	var diags diag.Diagnostics

	var action types.String
	var appqoe AppqoeValue = NewAppqoeValueNull()
	var ewf types.List = types.ListNull(EwfValue{}.Type(ctx))
	var id types.String
	var idp IdpValue = NewIdpValueNull()
	var local_routing types.Bool
	var name types.String
	var org_id types.String
	var path_preference types.String
	var services types.List = types.ListNull(types.StringType)
	var tenants types.List = types.ListNull(types.StringType)

	if d.Action != nil {
		action = types.StringValue(string(*d.Action))
	}
	if d.Appqoe != nil {
		appqoe = appQoeToTerraform(ctx, &diags, d.Appqoe)
	}
	if d.Ewf != nil {
		ewf = ewfSdkToTerraform(ctx, &diags, d.Ewf)
	}
	id = types.StringValue(d.Id.String())

	if d.Idp != nil {
		idp = idpSdkToTerraform(ctx, &diags, d.Idp)
	}
	if d.LocalRouting != nil {
		local_routing = types.BoolValue(*d.LocalRouting)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	org_id = types.StringValue(d.OrgId.String())

	if d.PathPreference != nil {
		path_preference = types.StringValue(*d.PathPreference)
	}
	if d.Services != nil {
		services = mist_transform.ListOfStringSdkToTerraform(ctx, d.Services)
	}
	if d.Tenants != nil {
		tenants = mist_transform.ListOfStringSdkToTerraform(ctx, d.Tenants)
	}

	state.Action = action
	state.Appqoe = appqoe
	state.Ewf = ewf
	state.Id = id
	state.Idp = idp
	state.LocalRouting = local_routing
	state.Name = name
	state.OrgId = org_id
	state.PathPreference = path_preference
	state.Services = services
	state.Tenants = tenants

	return state, diags

}
