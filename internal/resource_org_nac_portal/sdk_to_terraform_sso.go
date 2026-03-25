package resource_org_nac_portal

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func ssoRoleMatchingSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, l []models.NacPortalSsoRoleMatching) basetypes.ListValue {
	var dataList []SsoRoleMatchingValue
	for _, d := range l {
		var assigned basetypes.StringValue
		var match basetypes.StringValue

		if d.Assigned != nil {
			assigned = types.StringValue(*d.Assigned)
		}
		if d.Match != nil {
			match = types.StringValue(*d.Match)
		}

		dataMapValue := map[string]attr.Value{
			"assigned": assigned,
			"match":    match,
		}
		data, e := NewSsoRoleMatchingValue(SsoRoleMatchingValue{}.AttributeTypes(ctx), dataMapValue)
		diags.Append(e...)

		dataList = append(dataList, data)
	}

	r, e := types.ListValueFrom(ctx, SsoRoleMatchingValue{}.Type(ctx), dataList)
	diags.Append(e...)
	return r
}

func ssoSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, d *models.NacPortalSso) SsoValue {
	var idpCert basetypes.StringValue
	var idpSignAlgo basetypes.StringValue
	var idpSsoUrl basetypes.StringValue
	var issuer basetypes.StringValue
	var nameidFormat basetypes.StringValue
	var ssoRoleMatching = types.ListNull(SsoRoleMatchingValue{}.Type(ctx))
	var useSsoRoleForCert basetypes.BoolValue

	if d != nil && d.IdpCert != nil {
		idpCert = types.StringValue(*d.IdpCert)
	}
	if d != nil && d.IdpSignAlgo != nil {
		idpSignAlgo = types.StringValue(string(*d.IdpSignAlgo))
	}
	if d != nil && d.IdpSsoUrl != nil {
		idpSsoUrl = types.StringValue(*d.IdpSsoUrl)
	}
	if d != nil && d.Issuer != nil {
		issuer = types.StringValue(*d.Issuer)
	}
	if d != nil && d.NameidFormat != nil {
		nameidFormat = types.StringValue(*d.NameidFormat)
	}
	if d != nil && d.SsoRoleMatching != nil {
		ssoRoleMatching = ssoRoleMatchingSdkToTerraform(ctx, diags, d.SsoRoleMatching)
	}
	if d != nil && d.UseSsoRoleForCert != nil {
		useSsoRoleForCert = types.BoolValue(*d.UseSsoRoleForCert)
	}

	dataMapValue := map[string]attr.Value{
		"idp_cert":              idpCert,
		"idp_sign_algo":         idpSignAlgo,
		"idp_sso_url":           idpSsoUrl,
		"issuer":                issuer,
		"nameid_format":         nameidFormat,
		"sso_role_matching":     ssoRoleMatching,
		"use_sso_role_for_cert": useSsoRoleForCert,
	}
	data, e := NewSsoValue(SsoValue{}.AttributeTypes(ctx), dataMapValue)
	diags.Append(e...)

	return data
}
