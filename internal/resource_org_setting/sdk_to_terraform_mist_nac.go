package resource_org_setting

import (
	"context"

	mistutils "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func mistNacIdpsSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.OrgSettingMistNacIdp) basetypes.ListValue {

	var dataList []IdpsValue
	for _, d := range l {
		var excludeRealms = mistutils.ListOfStringSdkToTerraformEmpty()
		var id basetypes.StringValue
		var userRealms = types.ListNull(types.StringType)

		if d.ExcludeRealms != nil {
			excludeRealms = mistutils.ListOfStringSdkToTerraform(d.ExcludeRealms)
		}
		if d.Id != nil {
			id = types.StringValue(d.Id.String())
		}
		if d.UserRealms != nil {
			userRealms = mistutils.ListOfStringSdkToTerraform(d.UserRealms)
		}

		dataMapValue := map[string]attr.Value{
			"exclude_realms": excludeRealms,
			"id":             id,
			"user_realms":    userRealms,
		}
		data, e := NewIdpsValue(IdpsValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}
	r, e := types.ListValueFrom(ctx, IdpsValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func mistNacServerCertSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMistNacServerCert) basetypes.ObjectValue {
	var cert basetypes.StringValue
	var key basetypes.StringValue
	var password basetypes.StringValue

	if d.Cert != nil {
		cert = types.StringValue(*d.Cert)
	}
	if d.Key != nil {
		key = types.StringValue(*d.Key)
	}
	if d.Password != nil {
		password = types.StringValue(*d.Password)
	}

	dataMapValue := map[string]attr.Value{
		"cert":     cert,
		"key":      key,
		"password": password,
	}
	r, e := basetypes.NewObjectValue(ServerCertValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)
	return r
}

func mistNacSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.OrgSettingMistNac) MistNacValue {
	var cacerts = mistutils.ListOfStringSdkToTerraformEmpty()
	var defaultIdpId basetypes.StringValue
	var disableRsaeAlgorithms basetypes.BoolValue
	var eapSslSecurityLevel basetypes.Int64Value
	var euOnly basetypes.BoolValue
	var idps = types.ListNull(IdpsValue{}.Type(ctx))
	var idpMachineCertLookupField basetypes.StringValue
	var idpUserCertLookupField basetypes.StringValue
	var serverCert = types.ObjectNull(ServerCertValue{}.AttributeTypes(ctx))
	var useIpVersion basetypes.StringValue
	var useSslPort basetypes.BoolValue

	if d.Cacerts != nil {
		cacerts = mistutils.ListOfStringSdkToTerraform(d.Cacerts)
	}
	if d.DefaultIdpId != nil {
		defaultIdpId = types.StringValue(*d.DefaultIdpId)
	}
	if d.DisableRsaeAlgorithms != nil {
		disableRsaeAlgorithms = types.BoolValue(*d.DisableRsaeAlgorithms)
	}
	if d.EapSslSecurityLevel != nil {
		eapSslSecurityLevel = types.Int64Value(int64(*d.EapSslSecurityLevel))
	}
	if d.EuOnly != nil {
		euOnly = types.BoolValue(*d.EuOnly)
	}
	if d.Idps != nil {
		idps = mistNacIdpsSdkToTerraform(ctx, diags, d.Idps)
	}
	if d.IdpMachineCertLookupField != nil {
		idpMachineCertLookupField = types.StringValue(string(*d.IdpMachineCertLookupField))
	}
	if d.IdpUserCertLookupField != nil {
		idpUserCertLookupField = types.StringValue(string(*d.IdpUserCertLookupField))
	}
	if d.ServerCert != nil {
		serverCert = mistNacServerCertSdkToTerraform(ctx, diags, d.ServerCert)
	}
	if d.UseIpVersion != nil {
		useIpVersion = types.StringValue(string(*d.UseIpVersion))
	}
	if d.UseSslPort != nil {
		useSslPort = types.BoolValue(*d.UseSslPort)
	}

	dataMapValue := map[string]attr.Value{
		"cacerts":                       cacerts,
		"default_idp_id":                defaultIdpId,
		"disable_rsae_algorithms":       disableRsaeAlgorithms,
		"eap_ssl_security_level":        eapSslSecurityLevel,
		"eu_only":                       euOnly,
		"idps":                          idps,
		"idp_machine_cert_lookup_field": idpMachineCertLookupField,
		"idp_user_cert_lookup_field":    idpUserCertLookupField,
		"server_cert":                   serverCert,
		"use_ip_version":                useIpVersion,
		"use_ssl_port":                  useSslPort,
	}
	data, e := NewMistNacValue(MistNacValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
