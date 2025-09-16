name = "test-nacidp-ldap"
idp_type = "ldap"
ldap_base_dn = "dc=example,dc=com"
ldap_bind_dn = "cn=admin,dc=example,dc=com"
ldap_bind_password = "password123"
ldap_cacerts = ["-----BEGIN CERTIFICATE-----\nMIIC...test...cert...\n-----END CERTIFICATE-----"]
ldap_client_cert = "-----BEGIN CERTIFICATE-----\nMIIC...client...cert...\n-----END CERTIFICATE-----"
ldap_client_key = "-----BEGIN RSA PRIVATE KEY-----\nMIIE...client...key...\n-----END RSA PRIVATE KEY-----"
ldap_group_attr = "memberOf"
ldap_group_dn = "cn=groups,dc=example,dc=com"
ldap_resolve_groups = true
ldap_server_hosts = ["ldap.example.com", "ldap2.example.com"]
ldap_type = "custom"
ldap_user_filter = "(objectClass=person)"
member_filter = "(objectClass=group)"
group_filter = "(objectClass=groupOfNames)"
␞
name = "test-nacidp-oauth"
idp_type = "oauth"
oauth_type = "azure"
oauth_cc_client_id = "12345678-1234-1234-1234-123456789abc"
oauth_cc_client_secret = "-----BEGIN RSA PRIVATE KEY-----\nMIIE...oauth...key...\n-----END RSA PRIVATE KEY-----"
oauth_discovery_url = "https://login.microsoftonline.com/tenant-id/.well-known/openid-configuration"
oauth_ping_identity_region = "us"
oauth_ropc_client_id = "87654321-4321-4321-4321-cba987654321"
oauth_ropc_client_secret = "oauth-ropc-secret"
oauth_tenant_id = "tenant-12345678-1234-1234-1234-123456789abc"
scim_enabled = true
scim_secret_token = "scim-secret-token-123"
␞
name = "test-nacidp-mxedge"
idp_type = "mxedge_proxy"