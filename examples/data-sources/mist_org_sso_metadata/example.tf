data "mist_org_sso_metadata" "sso_idp" {
    org_id    = mist_org.terraform_test.id
    sso_id = mist_org_sso.sso_one.id
}
