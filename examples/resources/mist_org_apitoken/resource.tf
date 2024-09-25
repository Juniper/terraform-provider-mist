resource "mist_org_apitoken" "apitoken_one" {
  org_id = mist_org.terraform_test.id
  name   = "apitoken_one"
  privileges = [
    {
      scope   = "site"
      role    = "admin"
      site_id = "d7c8364e-xxxx-xxxx-xxxx-37eff0475b03"
    },
    {
      scope   = "site"
      role    = "read"
      site_id = "08f8851b-xxxx-xxxx-xxxx-9ebb5aa62de4"
    }
  ]
  src_ips = [ "1.2.3.4/32" ]
}
