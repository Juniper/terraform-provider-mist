package resource_org_nacidp

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgNacidpModel) (*models.Sso, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})
	data := models.Sso{}

	if !plan.IdpType.IsNull() && !plan.IdpType.IsUnknown() {
		data.IdpType = (*models.SsoIdpTypeEnum)(plan.IdpType.ValueStringPointer())
	} else {
		unset["-idp_type"] = ""
	}
	if !plan.LdapBaseDn.IsNull() && !plan.LdapBaseDn.IsUnknown() {
		data.LdapBaseDn = plan.LdapBaseDn.ValueStringPointer()
	} else {
		unset["-ldap_base_dn"] = ""
	}

	if !plan.LdapBindDn.IsNull() && !plan.LdapBindDn.IsUnknown() {
		data.LdapBindDn = plan.LdapBindDn.ValueStringPointer()
	} else {
		unset["-ldap_bind_dn"] = ""
	}

	if !plan.LdapBindPassword.IsNull() && !plan.LdapBindPassword.IsUnknown() {
		data.LdapBindPassword = plan.LdapBindPassword.ValueStringPointer()
	} else {
		unset["-ldap_bind_password"] = ""
	}

	if !plan.LdapCacerts.IsNull() && !plan.LdapCacerts.IsUnknown() {
		data.LdapCacerts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.LdapCacerts)
	} else {
		unset["-ldap_ca_certs"] = ""
	}

	if !plan.LdapClientCert.IsNull() && !plan.LdapClientCert.IsUnknown() {
		data.LdapClientCert = plan.LdapClientCert.ValueStringPointer()
	} else {
		unset["-ldap_client_cert"] = ""
	}

	if !plan.LdapClientKey.IsNull() && !plan.LdapClientKey.IsUnknown() {
		data.LdapClientKey = plan.LdapClientKey.ValueStringPointer()
	} else {
		unset["-ldap_client_key"] = ""
	}

	if !plan.LdapGroupAttr.IsNull() && !plan.LdapGroupAttr.IsUnknown() {
		data.LdapGroupAttr = plan.LdapGroupAttr.ValueStringPointer()
	} else {
		unset["-ldap_group_attr"] = ""
	}

	if !plan.LdapGroupDn.IsNull() && !plan.LdapGroupDn.IsUnknown() {
		data.LdapGroupDn = plan.LdapGroupDn.ValueStringPointer()
	} else {
		unset["-ldap_group_dn"] = ""
	}

	if !plan.GroupFilter.IsNull() && !plan.GroupFilter.IsUnknown() {
		data.GroupFilter = plan.GroupFilter.ValueStringPointer()
	} else {
		unset["-ldap_group_filter"] = ""
	}

	if !plan.MemberFilter.IsNull() && !plan.MemberFilter.IsUnknown() {
		data.MemberFilter = plan.MemberFilter.ValueStringPointer()
	} else {

		unset["-ldap_member_filter"] = ""
	}
	if !plan.LdapResolveGroups.IsNull() && !plan.LdapResolveGroups.IsUnknown() {
		data.LdapResolveGroups = plan.LdapResolveGroups.ValueBoolPointer()
	} else {
		unset["-ldap_resolve_groups"] = ""
	}

	if !plan.LdapServerHosts.IsNull() && !plan.LdapServerHosts.IsUnknown() {
		data.LdapServerHosts = mist_transform.ListOfStringTerraformToSdk(ctx, plan.LdapServerHosts)
	} else {
		unset["-ldap_server_hosts"] = ""
	}

	if !plan.LdapType.IsNull() && !plan.LdapType.IsUnknown() {
		data.LdapType = (*models.SsoLdapTypeEnum)(plan.LdapType.ValueStringPointer())
	} else {
		unset["-ldap_type"] = ""
	}

	if !plan.LdapUserFilter.IsNull() && !plan.LdapUserFilter.IsUnknown() {
		data.LdapUserFilter = plan.LdapUserFilter.ValueStringPointer()
	} else {
		unset["-ldap_user_filter"] = ""
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		data.Name = plan.Name.ValueString()
	} else {
		unset["-name"] = ""
	}

	if !plan.OauthCcClientId.IsNull() && !plan.OauthCcClientId.IsUnknown() {
		data.OauthCcClientId = plan.OauthCcClientId.ValueStringPointer()
	} else {
		unset["-oauth_cc_client_id"] = ""
	}

	if !plan.OauthCcClientSecret.IsNull() && !plan.OauthCcClientSecret.IsUnknown() {
		data.OauthCcClientSecret = plan.OauthCcClientSecret.ValueStringPointer()
	} else {
		unset["-oauth_cc_client_secret"] = ""
	}

	if !plan.OauthDiscoveryUrl.IsNull() && !plan.OauthDiscoveryUrl.IsUnknown() {
		data.OauthDiscoveryUrl = plan.OauthDiscoveryUrl.ValueStringPointer()
	} else {
		unset["-oauth_discovery_url"] = ""
	}

	if !plan.OauthRopcClientId.IsNull() && !plan.OauthRopcClientId.IsUnknown() {
		data.OauthRopcClientId = plan.OauthRopcClientId.ValueStringPointer()
	} else {
		unset["-oauth_ropc_client_id"] = ""
	}

	if !plan.OauthRopcClientSecret.IsNull() && !plan.OauthRopcClientSecret.IsUnknown() {
		data.OauthRopcClientSecret = plan.OauthRopcClientSecret.ValueStringPointer()
	} else {
		unset["-oauth_ropc_client_secret"] = ""
	}

	if !plan.OauthTenantId.IsNull() && !plan.OauthTenantId.IsUnknown() {
		data.OauthTenantId = plan.OauthTenantId.ValueStringPointer()
	} else {
		unset["-oauth_tenant_id"] = ""
	}

	if !plan.OauthType.IsNull() && !plan.OauthType.IsUnknown() {
		/// TEMP WORKAROUND
		if plan.OauthType.ValueString() == "azure" {
			data.OauthType = models.ToPointer(models.SsoOauthTypeEnum_STANDARDS)
		} else {
			data.OauthType = (*models.SsoOauthTypeEnum)(plan.OauthType.ValueStringPointer())
		}
	} else {
		unset["-oauth_type"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
