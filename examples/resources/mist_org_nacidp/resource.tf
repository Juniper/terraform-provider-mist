
// OAuth with Azure Example
resource "mist_org_nacidp" "idp_azure" {
  org_id                   = mist_org.terraform_test.id
  name                     = "idp_azure"
  idp_type                 = "oauth"
  oauth_cc_client_id       = "client_id"
  oauth_cc_client_secret   = "-----BEGIN CERTIFICATE-----MIIF0jC .../fSCGx7-----END CERTIFICATE-----"
  oauth_ropc_client_id     = "ropc_client_id"
  oauth_ropc_client_secret = "ropc_client_secret"
  oauth_tenant_id          = "tenant_id"
  oauth_type               = "azure"
}

// Custom LDAP Example
resource "mist_org_nacidp" "idp_ldap" {
  org_id             = mist_org.terraform_test.id
  name               = "idp_ldap"
  idp_type           = "ldap"
  ldap_type          = "custom"
  group_filter       = "memberOf"
  member_filter      = "memberOf"
  ldap_user_filter   = "(mail=%s)"
  ldap_server_hosts  = ["ldap.mycorp.com", "1.2.3.4"]
  ldap_base_dn       = "DC=abc,DC=com"
  ldap_bind_dn       = "CN=admin,CN=users,DC=abc,DC=com"
  ldap_bind_password = "secret!password"
  ldap_cacerts = [
    "-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----",
    "-----BEGIN CERTIFICATE-----\nBhMCRVMxFDASBgNVBAoMC1N0YXJ0Q29tIENBMSwwKgYDVn-----END CERTIFICATE-----"
  ]
  ldap_client_cert = "-----BEGIN CERTIFICATE-----\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\n-----END CERTIFICATE-----"
  ldap_client_key  = "-----BEGIN PRI..."
}
