package resource_org_servicepolicy

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.OrgServicePolicy) (OrgServicepolicyModel, diag.Diagnostics) {
	var state OrgServicepolicyModel
	var diags diag.Diagnostics

	var aamw = NewAamwValueNull()
	var action types.String
	var antivirus = NewAntivirusValueNull()
	var appqoe = NewAppqoeValueNull()
	var ewf = types.ListNull(EwfValue{}.Type(ctx))
	var id types.String
	var idp = NewIdpValueNull()
	var localRouting types.Bool
	var name types.String
	var orgId types.String
	var pathPreference types.String
	var services = types.ListNull(types.StringType)
	var sslProxy = NewSslProxyValueNull()
	var tenants = types.ListNull(types.StringType)

	if d.Aamw != nil {
		aamw = aamwSdkToTerraform(ctx, &diags, d.Aamw)
	}
	if d.Action != nil {
		action = types.StringValue(string(*d.Action))
	}
	if d.Antivirus != nil {
		antivirus = avSdkToTerraform(ctx, &diags, d.Antivirus)
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
		localRouting = types.BoolValue(*d.LocalRouting)
	}
	if d.Name != nil {
		name = types.StringValue(*d.Name)
	}
	orgId = types.StringValue(d.OrgId.String())

	if d.PathPreference != nil {
		pathPreference = types.StringValue(*d.PathPreference)
	}
	if d.Services != nil {
		services = mistutils.ListOfStringSdkToTerraform(d.Services)
	}
	if d.SslProxy != nil {
		sslProxy = sslProxySdkToTerraform(ctx, &diags, d.SslProxy)
	}
	if d.Tenants != nil {
		tenants = mistutils.ListOfStringSdkToTerraform(d.Tenants)
	}

	state.Aamw = aamw
	state.Action = action
	state.Antivirus = antivirus
	state.Appqoe = appqoe
	state.Ewf = ewf
	state.Id = id
	state.Idp = idp
	state.LocalRouting = localRouting
	state.Name = name
	state.OrgId = orgId
	state.PathPreference = pathPreference
	state.Services = services
	state.SslProxy = sslProxy
	state.Tenants = tenants

	return state, diags

}
