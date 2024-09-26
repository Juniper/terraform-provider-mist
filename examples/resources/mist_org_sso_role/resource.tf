resource "mist_org_sso_role" "sso_role_one" {
  org_id = mist_org.terraform_test.id
  name   = "admin_sso"
  privileges = [
    {
      scope   = "site"
      role    = "read"
      site_id = mist_site.terraform_site.id
    }
  ]
}
