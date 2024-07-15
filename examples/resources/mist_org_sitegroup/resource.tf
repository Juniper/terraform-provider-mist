
resource "mist_org_sitegroup" "sitegroup_one" {
  org_id = mist_org.terraform_test.id
  name   = "sitegroup_one"
}