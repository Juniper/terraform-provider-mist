package datasource_org_servicepolicies

import (
	"context"

	misttransform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, l *[]models.OrgServicePolicy, elements *[]attr.Value) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, d := range *l {
		elem := servicepolicieSdkToTerraform(ctx, &diags, &d)
		*elements = append(*elements, elem)
	}

	return diags
}

func servicepolicieSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgServicePolicy) OrgServicepoliciesValue {

	var aamw = types.ObjectNull(AamwValue{}.AttributeTypes(ctx))
	var action types.String
	var antivirus = types.ObjectNull(AntivirusValue{}.AttributeTypes(ctx))
	var appqoe = types.ObjectNull(AppqoeValue{}.AttributeTypes(ctx))
	var createdTime basetypes.Float64Value
	var ewf = types.ListNull(EwfValue{}.Type(ctx))
	var id types.String
	var idp = types.ObjectNull(IdpValue{}.AttributeTypes(ctx))
	var localRouting types.Bool
	var modifiedTime basetypes.Float64Value
	var name types.String
	var orgId types.String
	var pathPreference types.String
	var services = types.ListNull(types.StringType)
	var sslProxy = types.ObjectNull(SslProxyValue{}.AttributeTypes(ctx))
	var tenants = types.ListNull(types.StringType)

	if d.Aamw != nil {
		aamw = aamwSdkToTerraform(ctx, diags, d.Aamw)
	}
	if d.Action != nil {
		action = types.StringValue(string(*d.Action))
	}
	if d.Antivirus != nil {
		antivirus = avSdkToTerraform(ctx, diags, d.Antivirus)
	}
	if d.Appqoe != nil {
		appqoe = appQoeToTerraform(ctx, diags, d.Appqoe)
	}
	if d.CreatedTime != nil {
		createdTime = types.Float64Value(*d.CreatedTime)
	}
	if d.Ewf != nil {
		ewf = ewfSdkToTerraform(ctx, diags, d.Ewf)
	}
	id = types.StringValue(d.Id.String())

	if d.Idp != nil {
		idp = idpSdkToTerraform(ctx, diags, d.Idp)
	}
	if d.ModifiedTime != nil {
		modifiedTime = types.Float64Value(*d.ModifiedTime)
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
		sslProxy = sslProxySdkToTerraform(ctx, diags, d.SslProxy)
	}
	if d.Tenants != nil {
		tenants = misttransform.ListOfStringSdkToTerraform(d.Tenants)
	}

	dataMapValue := map[string]attr.Value{
		"aamw":            aamw,
		"action":          action,
		"antivirus":       antivirus,
		"appqoe":          appqoe,
		"created_time":    createdTime,
		"ewf":             ewf,
		"id":              id,
		"idp":             idp,
		"local_routing":   localRouting,
		"modified_time":   modifiedTime,
		"name":            name,
		"org_id":          orgId,
		"path_preference": pathPreference,
		"services":        services,
		"ssl_proxy":       sslProxy,
		"tenants":         tenants,
	}
	data, e := NewOrgServicepoliciesValue(OrgServicepoliciesValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
