data "mist_org_nacidp_metadata" "saml_idp" {
    org_id    = mist_org.terraform_test.id
    nacidp_id = mist_org_nacidp.saml_idp_one.id
}
