package resource_org_servicepolicy

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func SdkToTerraform(ctx context.Context, d *models.OrgServicePolicy) (OrgServicepolicyModel, diag.Diagnostics) {
	var state OrgServicepolicyModel
	var diags diag.Diagnostics

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

	if d.Action != nil {
		action = types.StringValue(string(*d.Action))
	}
	if d.Antivirus != nil {
		antivirus = avSdkToTerraform(ctx, &diags, d.Antivirus)
	}
	if d.Appqoe != nil {
		appqoe = appQoeToTerraform(d.Appqoe)
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
		services = misttransform.ListOfStringSdkToTerraform(d.Services)
	}
	if d.SslProxy != nil {
		sslProxy = sslProxySdkToTerraform(ctx, &diags, d.SslProxy)
	}
	if d.Tenants != nil {
		tenants = misttransform.ListOfStringSdkToTerraform(d.Tenants)
	}

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
