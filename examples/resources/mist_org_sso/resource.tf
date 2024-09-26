resource "mist_org_sso" "sso_admin_one" {
  org_id            = mist_org.terraform_test.id
  name              = "sso_admin_one"
  custom_logout_url = "https://idp.com/logout"
  idp_cert          = "-----BEGIN CERTIFICATE-----MIIF0jC .../fSCGx7-----END CERTIFICATE-----"
  idp_sign_algo     = "sha512"
  idp_sso_url       = "https://idp.com/login"
  issuer            = "my_idp_issuer"
  nameid_format     = "email"
}
